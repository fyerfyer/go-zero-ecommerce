CREATE DATABASE reply;
\c reply;

CREATE TABLE replies (
    id BIGSERIAL PRIMARY KEY,
    comment_id BIGINT NOT NULL,
    user_id BIGINT NOT NULL,
    content TEXT NOT NULL,
    created_time TIMESTAMP DEFAULT CURRENT_TIMESTAMP
);

CREATE INDEX idx_replies_comment_id ON replies(comment_id);
CREATE INDEX idx_replies_user_id ON replies(user_id);
