create table `users`
(
    `id`          bigint auto_increment,
    `name`        varchar(100) comment '用户名',
    `avatar`      varchar(100) comment '用户头像',
    `gender`      enum ('保密', '女', '男') default '保密' comment '性别',
    `state`       enum ('0', '1')       default '1' comment '账户状态 0未激活 1激活',
    `create_time` timestamp NULL        DEFAULT CURRENT_TIMESTAMP,
    `update_time` timestamp NULL        DEFAULT CURRENT_TIMESTAMP ON UPDATE CURRENT_TIMESTAMP,
    PRIMARY KEY (`id`),
    UNIQUE KEY `name_index` (`name`)
) ENGINE = InnoDB
  DEFAULT CHARSET = utf8mb4 comment '用户信息表';