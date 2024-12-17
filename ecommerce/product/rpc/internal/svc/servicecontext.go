package svc

import (
	"golang.org/x/sync/singleflight"
	"time"

	"github.com/fyerfyer/go-zero-ecommerce/ecommerce/product/rpc/internal/config"
	"github.com/fyerfyer/go-zero-ecommerce/ecommerce/product/rpc/model"
	"github.com/jinzhu/gorm"
	_ "github.com/lib/pq"
	"github.com/zeromicro/go-zero/core/collection"
	"github.com/zeromicro/go-zero/core/stores/redis"
	"github.com/zeromicro/go-zero/core/stores/sqlx"
)

type ServiceContext struct {
	Config         config.Config
	ProductModel   model.ProductModel
	CategoryModel  model.CategoryModel
	OperationModel model.ProductOperationModel
	ProductRedis   *redis.Redis
	SingleGroup    singleflight.Group
	LocalCache     *collection.Cache
	orm            *gorm.DB
}

func NewServiceContext(c config.Config) *ServiceContext {
	// to do: add the constant into config
	conn := sqlx.NewSqlConn("postgres", c.DataSource)
	ormDB, err := gorm.Open("postgres", c.DataSource)
	if err != nil {
		panic("failed to connect to PostgreSQL:" + err.Error())
	}

	// set connection pool
	ormDB.DB().SetMaxOpenConns(20)
	ormDB.DB().SetMaxIdleConns(10)
	ormDB.DB().SetConnMaxIdleTime(time.Hour * 24)
	localCache, err := collection.NewCache(time.Second * 60)
	if err != nil {
		panic("failed to set up local cache:" + err.Error())
	}

	return &ServiceContext{
		Config:         c,
		ProductModel:   model.NewProductModel(conn, c.CacheRedis),
		CategoryModel:  model.NewCategoryModel(conn, c.CacheRedis),
		OperationModel: model.NewProductOperationModel(conn, c.CacheRedis),
		ProductRedis: redis.New(c.ProductRedis.Host,
			redis.WithPass(c.ProductRedis.Pass)),
		LocalCache: localCache,
		orm:        ormDB,
	}
}
