create table `users_auth`
(
    `id`             bigint auto_increment,
    `user_id`        bigint comment '用户id',
    `name`           varchar(100) comment '用户名',
    `password`       varchar(100) comment '用户密码',
    `token`          varchar(100) comment '用户token',
    `phone`          varchar(13) comment '手机号',
    `email`          varchar(50) comment '邮箱',
    `email_activate` enum ('0', '1') default '0' comment '邮箱是否激活 0未激活 1激活',
    `wechat_id`      varchar(50) comment '微信id',
    `wechat_token`   varchar(100) comment '微信token',
    `wechat_expires` varchar(50) comment '微信token',
    `create_time`    timestamp NULL  DEFAULT CURRENT_TIMESTAMP,
    `update_time`    timestamp NULL  DEFAULT CURRENT_TIMESTAMP ON UPDATE CURRENT_TIMESTAMP,
    PRIMARY KEY (`id`),
    UNIQUE KEY `phone_index` (`phone`),
) ENGINE = InnoDB
  DEFAULT CHARSET = utf8mb4 comment '用户验证表';