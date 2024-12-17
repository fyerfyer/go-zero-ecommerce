-- 数据库: orders
CREATE DATABASE orders;
\c orders;

CREATE TABLE orders (
    id UUID PRIMARY KEY DEFAULT gen_random_uuid(), -- 订单ID
    userid BIGINT NOT NULL DEFAULT 0, -- 用户ID
    shoppingid BIGINT NOT NULL DEFAULT 0, -- 收货信息ID
    payment NUMERIC(20,2) NOT NULL DEFAULT 0, -- 实际付款金额
    paymenttype SMALLINT NOT NULL DEFAULT 1, -- 支付类型
    postage INTEGER NOT NULL DEFAULT 0, -- 运费
    status SMALLINT NOT NULL DEFAULT 10, -- 订单状态
    create_time TIMESTAMPTZ NOT NULL DEFAULT CURRENT_TIMESTAMP, -- 创建时间
    update_time TIMESTAMPTZ NOT NULL DEFAULT CURRENT_TIMESTAMP -- 更新时间
);

CREATE INDEX ix_userid_orders ON orders (userid);
COMMENT ON TABLE orders IS '订单表';

CREATE TABLE orderitem (
    id BIGSERIAL PRIMARY KEY, -- 订单子表ID
    order_id UUID NOT NULL, -- 订单ID
    user_id BIGINT NOT NULL DEFAULT 0, -- 用户ID
    product_id BIGINT NOT NULL DEFAULT 0, -- 商品ID
    product_name VARCHAR(100) NOT NULL DEFAULT '', -- 商品名称
    product_image VARCHAR(500) NOT NULL DEFAULT '', -- 商品图片
    current_price NUMERIC(20,2) NOT NULL DEFAULT 0, -- 商品单价
    quantity INTEGER NOT NULL DEFAULT 0, -- 商品数量
    total_price NUMERIC(20,2) NOT NULL DEFAULT 0, -- 商品总价
    create_time TIMESTAMPTZ NOT NULL DEFAULT CURRENT_TIMESTAMP, -- 创建时间
    update_time TIMESTAMPTZ NOT NULL DEFAULT CURRENT_TIMESTAMP -- 更新时间
);

CREATE INDEX ix_orderid ON orderitem (order_id);
CREATE INDEX ix_userid_orderitem ON orderitem (user_id);
CREATE INDEX ix_proid_orderitem ON orderitem (product_id);
COMMENT ON TABLE orderitem IS '订单明细表';

CREATE TABLE shipping (
    id BIGSERIAL PRIMARY KEY, -- 收货信息ID
    orderid UUID NOT NULL, -- 订单ID
    userid BIGINT NOT NULL DEFAULT 0, -- 用户ID
    receiver_name VARCHAR(20) NOT NULL DEFAULT '', -- 收货姓名
    receiver_phone VARCHAR(20) NOT NULL DEFAULT '', -- 固定电话
    receiver_mobile VARCHAR(20) NOT NULL DEFAULT '', -- 移动电话
    receiver_province VARCHAR(20) NOT NULL DEFAULT '', -- 省份
    receiver_city VARCHAR(20) NOT NULL DEFAULT '', -- 城市
    receiver_district VARCHAR(20) NOT NULL DEFAULT '', -- 区县
    receiver_address VARCHAR(200) NOT NULL DEFAULT '', -- 详细地址
    create_time TIMESTAMPTZ NOT NULL DEFAULT CURRENT_TIMESTAMP, -- 创建时间
    update_time TIMESTAMPTZ NOT NULL DEFAULT CURRENT_TIMESTAMP -- 更新时间
);

CREATE INDEX ix_orderid_shipping ON shipping (orderid);
CREATE INDEX ix_userid_shipping ON shipping (userid);
COMMENT ON TABLE shipping IS '收货信息表';