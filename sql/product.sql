-- 数据库: product
CREATE DATABASE product;
\c product;

CREATE TABLE product (
    id BIGSERIAL PRIMARY KEY, -- 商品ID
    cateid SMALLINT NOT NULL DEFAULT 0, -- 类别ID
    name VARCHAR(100) NOT NULL DEFAULT '', -- 商品名称
    subtitle VARCHAR(200) NOT NULL DEFAULT '', -- 商品副标题
    images VARCHAR(1024) NOT NULL DEFAULT '', -- 图片地址,逗号分隔
    detail VARCHAR(1024) NOT NULL DEFAULT '', -- 商品详情
    price NUMERIC(20, 2) NOT NULL DEFAULT 0, -- 价格
    stock INTEGER NOT NULL DEFAULT 0, -- 库存数量
    status SMALLINT NOT NULL DEFAULT 1, -- 商品状态
    create_time TIMESTAMPTZ NOT NULL DEFAULT CURRENT_TIMESTAMP, -- 创建时间
    update_time TIMESTAMPTZ NOT NULL DEFAULT CURRENT_TIMESTAMP -- 更新时间
);

CREATE INDEX ix_cateid ON product (cateid);
CREATE INDEX ix_update_time_product ON product (update_time);
COMMENT ON TABLE product IS '商品表';

-- 表: category
CREATE TABLE category (
    id SMALLSERIAL PRIMARY KEY, -- 分类ID
    parentid SMALLINT NOT NULL DEFAULT 0, -- 父类别ID
    name VARCHAR(50) NOT NULL DEFAULT '', -- 类别名称
    status SMALLINT NOT NULL DEFAULT 1, -- 类别状态
    create_time TIMESTAMPTZ NOT NULL DEFAULT CURRENT_TIMESTAMP, -- 创建时间
    update_time TIMESTAMPTZ NOT NULL DEFAULT CURRENT_TIMESTAMP -- 更新时间
);

COMMENT ON TABLE category IS '商品类别表';

-- 表: product_operation
CREATE TABLE product_operation (
    id BIGSERIAL PRIMARY KEY, -- 商品运营ID
    product_id BIGINT NOT NULL DEFAULT 0, -- 商品ID
    status SMALLINT NOT NULL DEFAULT 1, -- 商品运营状态
    create_time TIMESTAMPTZ NOT NULL DEFAULT CURRENT_TIMESTAMP, -- 创建时间
    update_time TIMESTAMPTZ NOT NULL DEFAULT CURRENT_TIMESTAMP -- 更新时间
);

CREATE INDEX ix_update_time_product_op ON product_operation (update_time);
COMMENT ON TABLE product_operation IS '商品运营表';