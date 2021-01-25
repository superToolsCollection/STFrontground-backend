create table `tools`
(
    `id`           bigint AUTO_INCREMENT,
    `name`         varchar(50)       comment '工具名称',
    `api`          varchar(100)      comment '工具api链接',
    `api_describe` longtext          comment '工具描述',
    `created_on`   timestamp NULL DEFAULT CURRENT_TIMESTAMP comment '新建时间',
    `modified_on`  timestamp NULL DEFAULT CURRENT_TIMESTAMP ON UPDATE CURRENT_TIMESTAMP comment '修改时间',
    `is_del`       tinyint default 0 comment '是否删除 0为未删除 1为已删除',
    `state`        tinyint default 0 comment '状态 0为未上线 1为上线',
    `created_by`   varchar(100)      comment '创建人',
    `modified_by`  varchar(100)      comment '修改人',
    primary key(`id`),
    unique key `api_index` (`api`)
)ENGINE=InnoDB DEFAULT CHARSET=utf8mb4 COLLATE=utf8mb4_general_ci comment '工具列表';
