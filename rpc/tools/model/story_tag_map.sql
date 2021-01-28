create table `story_tag_map`
(
    `id`           bigint AUTO_INCREMENT,
    `story_id`         bigint      comment '睡前故事ID',
    `tag_id`         bigint      comment '标签ID',
    `created_on`   timestamp NULL DEFAULT CURRENT_TIMESTAMP comment '新建时间',
    `modified_on`  timestamp NULL DEFAULT CURRENT_TIMESTAMP ON UPDATE CURRENT_TIMESTAMP comment '修改时间',
    `is_del`       enum ('0', '1')       default '1' comment '是否删除 0未删除 1删除',
    `created_by`   varchar(100)      comment '创建人',
    `modified_by`  varchar(100)      comment '修改人',
    primary key(`id`)
)ENGINE=InnoDB DEFAULT CHARSET=utf8mb4 COLLATE=utf8mb4_general_ci comment '睡前故事标签关联表';
