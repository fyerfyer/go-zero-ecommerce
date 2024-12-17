-- 数据库: payment
CREATE DATABASE payment;
\c payment;

CREATE TABLE payinfo (
    id BIGSERIAL PRIMARY KEY, -- 支付信息表ID
    orderid UUID NOT NULL, -- 订单ID
    userid BIGINT NOT NULL DEFAULT 0, -- 用户ID
    payplatform SMALLINT NOT NULL DEFAULT 0, -- 支付平台
    platformnumber VARCHAR(200) NOT NULL DEFAULT '', -- 支付流水号
    platformstatus VARCHAR(20) NOT NULL DEFAULT '', -- 支付状态
    create_time TIMESTAMPTZ NOT NULL DEFAULT CURRENT_TIMESTAMP, -- 创建时间
    update_time TIMESTAMPTZ NOT NULL DEFAULT CURRENT_TIMESTAMP -- 更新时间
);

CREATE INDEX ix_orderid_pay ON payinfo (orderid);
CREATE INDEX ix_userid_pay ON payinfo (userid);
COMMENT ON TABLE payinfo IS '支付信息表';