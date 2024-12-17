-- 数据库: cart
CREATE DATABASE cart;
\c cart;

CREATE TABLE cart (
    id BIGSERIAL PRIMARY KEY, -- 购物车ID
    userid BIGINT NOT NULL DEFAULT 0, -- 用户ID
    proid BIGINT NOT NULL DEFAULT 0, -- 商品ID
    quantity INTEGER NOT NULL DEFAULT 0, -- 数量
    checked BOOLEAN NOT NULL DEFAULT FALSE, -- 是否勾选
    create_time TIMESTAMPTZ NOT NULL DEFAULT CURRENT_TIMESTAMP, -- 创建时间
    update_time TIMESTAMPTZ NOT NULL DEFAULT CURRENT_TIMESTAMP -- 更新时间
);

CREATE INDEX ix_userid ON cart (userid);
CREATE INDEX ix_proid ON cart (proid);
COMMENT ON TABLE cart IS '购物车表';