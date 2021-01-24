create table `users`
(
    `id`       int auto_increment,
    `mobile`   varchar(100) null comment '手机号',
    `name`     varchar(100) null comment '用户名',
    `password` varchar(100) null comment '加密后的密码',
    PRIMARY KEY(`id`)
)ENGINE=InnoDB DEFAULT CHARSET=utf8mb4 comment '用户表';