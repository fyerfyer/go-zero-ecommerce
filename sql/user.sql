-- 数据库: users
CREATE DATABASE users;
\c users;

CREATE TABLE "users" (
    id BIGSERIAL PRIMARY KEY,                  -- 用户ID
    username VARCHAR(50) NOT NULL DEFAULT '',  -- 用户名
    password VARCHAR(50) NOT NULL DEFAULT '',  -- 用户密码，MD5加密
    phone VARCHAR(20) NOT NULL DEFAULT '',     -- 手机号
    question VARCHAR(100) NOT NULL DEFAULT '', -- 找回密码问题
    answer VARCHAR(100) NOT NULL DEFAULT '',   -- 找回密码答案
    create_time TIMESTAMPTZ NOT NULL DEFAULT CURRENT_TIMESTAMP, -- 创建时间
    update_time TIMESTAMPTZ NOT NULL DEFAULT CURRENT_TIMESTAMP -- 更新时间
);

CREATE UNIQUE INDEX uniq_phone ON "users" (phone);
CREATE UNIQUE INDEX uniq_username ON "users" (username);
CREATE INDEX ix_update_time_users ON "users" (update_time);

COMMENT ON TABLE "users" IS '用户表';

CREATE TABLE users_receive_address (
    id BIGSERIAL PRIMARY KEY,
    uid BIGINT NOT NULL DEFAULT 0, -- 用户id
    name VARCHAR(64) NOT NULL DEFAULT '', -- 收货人名称
    phone VARCHAR(20) NOT NULL DEFAULT '', -- 手机号
    is_default BOOLEAN NOT NULL DEFAULT FALSE, -- 是否为默认地址
    post_code VARCHAR(100) NOT NULL DEFAULT '', -- 邮政编码
    province VARCHAR(100) NOT NULL DEFAULT '', -- 省份/直辖市
    city VARCHAR(100) NOT NULL DEFAULT '', -- 城市
    region VARCHAR(100) NOT NULL DEFAULT '', -- 区
    detail_address VARCHAR(128) NOT NULL DEFAULT '', -- 详细地址(街道)
    is_delete BOOLEAN NOT NULL DEFAULT FALSE, -- 是否删除
    create_time TIMESTAMPTZ NOT NULL DEFAULT CURRENT_TIMESTAMP, -- 数据创建时间
    update_time TIMESTAMPTZ NOT NULL DEFAULT CURRENT_TIMESTAMP -- 数据更新时间
);

CREATE INDEX idx_uid ON users_receive_address (uid);
COMMENT ON TABLE users_receive_address IS '用户收货地址表';

CREATE TABLE users_collection (
    id BIGSERIAL PRIMARY KEY, -- 收藏Id
    uid BIGINT NOT NULL DEFAULT 0, -- 用户id
    product_id BIGINT NOT NULL DEFAULT 0, -- 商品id
    is_delete BOOLEAN NOT NULL DEFAULT FALSE, -- 是否删除
    create_time TIMESTAMPTZ NOT NULL DEFAULT CURRENT_TIMESTAMP, -- 数据创建时间[禁止在代码中赋值]
    update_time TIMESTAMPTZ NOT NULL DEFAULT CURRENT_TIMESTAMP, -- 数据更新时间[禁止在代码中赋值]
    CONSTRAINT uniq_uid_product_id UNIQUE (uid, product_id)
);

COMMENT ON TABLE users_collection IS '用户收藏表';