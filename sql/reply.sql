-- 数据库: reply
CREATE DATABASE reply;
\c reply;

CREATE TABLE reply (
    id BIGSERIAL PRIMARY KEY, -- 评论表ID
    business VARCHAR(64) NOT NULL DEFAULT '', -- 评论业务类型
    targetid BIGINT NOT NULL DEFAULT 0, -- 评论目标ID
    reply_userid BIGINT NOT NULL DEFAULT 0, -- 回复用户ID
    be_reply_userid BIGINT NOT NULL DEFAULT 0, -- 被回复用户ID
    parentid BIGINT NOT NULL DEFAULT 0, -- 父评论ID
    content VARCHAR(255) NOT NULL DEFAULT '', -- 评论内容
    image VARCHAR(255) NOT NULL DEFAULT '', -- 评论图片
    create_time TIMESTAMPTZ NOT NULL DEFAULT CURRENT_TIMESTAMP, -- 创建时间
    update_time TIMESTAMPTZ NOT NULL DEFAULT CURRENT_TIMESTAMP -- 更新时间
);

CREATE INDEX ix_targetid_reply ON reply (targetid);
COMMENT ON TABLE reply IS '评论列表';