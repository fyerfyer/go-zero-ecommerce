package batcher

import (
	"context"
	"errors"
	"log"
	"sync"
	"time"
)

var ErrFull = errors.New("channel is full")

type Option func(*options)

type options struct {
	size     int
	buffer   int
	worker   int
	interval time.Duration
}

func defaultOptions() options {
	return options{
		size:     100,
		buffer:   100,
		worker:   5,
		interval: time.Second,
	}
}

func WithSize(size int) Option {
	return func(o *options) { o.size = size }
}

func WithBuffer(buffer int) Option {
	return func(o *options) { o.buffer = buffer }
}

func WithWorker(worker int) Option {
	return func(o *options) { o.worker = worker }
}

func WithInterval(interval time.Duration) Option {
	return func(o *options) { o.interval = interval }
}

func (o *options) validate() {
	if o.size <= 0 {
		o.size = 100
	}
	if o.buffer <= 0 {
		o.buffer = 100
	}
	if o.worker <= 0 {
		o.worker = 5
	}
	if o.interval <= 0 {
		o.interval = time.Second
	}
}

type msg struct {
	key string
	val interface{}
}

type Batcher struct {
	opts      options
	Do        func(ctx context.Context, val map[string][]interface{})
	Sharding  func(key string) int
	chans     []chan *msg
	waitGroup sync.WaitGroup
}

func New(opts ...Option) *Batcher {
	options := defaultOptions()
	for _, opt := range opts {
		opt(&options)
	}

	options.validate()

	b := &Batcher{
		opts:  options,
		chans: make([]chan *msg, options.worker),
	}
	for i := range b.chans {
		b.chans[i] = make(chan *msg, options.buffer)
	}
	return b
}

// Start启动所有通道的消息处理方法merge
func (b *Batcher) Start() {
	if b.Do == nil || b.Sharding == nil {
		log.Fatal("Batcher: Do or Sharding func is nil")
	}

	b.waitGroup.Add(len(b.chans))
	for i, ch := range b.chans {
		go b.merge(i, ch)
	}
}

// Add添加消息到对应通道中
func (b *Batcher) Add(key string, val interface{}) error {
	ch, msg := b.getChannelAndMessage(key, val)
	select {
	case ch <- msg:
		return nil
	default:
		return ErrFull
	}
}

// getChannelAndMessage决定发送消息的通道
func (b *Batcher) getChannelAndMessage(key string, val interface{}) (chan *msg, *msg) {
	idx := b.Sharding(key) % b.opts.worker
	ch := b.chans[idx]
	msg := &msg{key: key, val: val}
	return ch, msg
}

// merge方法处理batcher通道中的消息
func (b *Batcher) merge(idx int, ch <-chan *msg) {
	defer b.waitGroup.Done()

	vals := make(map[string][]interface{})
	count := 0
	ticker := time.NewTicker(b.opts.interval)
	defer ticker.Stop()

	if idx > 0 {
		// 避免时间冲突
		ticker = time.NewTicker(time.Duration(int64(idx) * (int64(b.opts.interval) / int64(b.opts.worker))))
		defer ticker.Stop()
	}

	for {
		select {
		case msg, ok := <-ch:
			// 通道关闭，处理通道内剩余值
			if !ok {
				b.flush(vals)
				return
			}
			vals[msg.key] = append(vals[msg.key], msg.val)
			count++

			// 通道已满，处理当前通道值
			if count >= b.opts.size {
				b.flush(vals)
				count = 0
			}

		// 计时器定时调用处理函数
		case <-ticker.C:
			if len(vals) > 0 {
				b.flush(vals)
				count = 0
			}
		}
	}
}

// flush方法用于批量处理vals，对所有待处理值应用batcher的Do函数
func (b *Batcher) flush(vals map[string][]interface{}) {
	if len(vals) > 0 {
		b.Do(context.Background(), vals)
		// Clear the map after processing
		for k := range vals {
			delete(vals, k)
		}
	}
}

// Close关闭所有通道，关闭后的通道会触发merge方法中的ok为false，从而停止batcher
func (b *Batcher) Close() {
	for _, ch := range b.chans {
		close(ch)
	}
	b.waitGroup.Wait()
}
