CREATE TABLE `ms_project`
(
    `id`                   bigint(0) UNSIGNED                                      NOT NULL AUTO_INCREMENT,
    `cover`                varchar(255) CHARACTER SET utf8 COLLATE utf8_general_ci NULL DEFAULT NULL COMMENT '封面',
    `name`                 varchar(90) CHARACTER SET utf8 COLLATE utf8_general_ci  NULL DEFAULT NULL COMMENT '名称',
    `description`          text CHARACTER SET utf8 COLLATE utf8_general_ci         NULL COMMENT '描述',
    `access_control_type`  tinyint(0)                                              NULL DEFAULT 0 COMMENT '访问控制l类型',
    `white_list`           varchar(255) CHARACTER SET utf8 COLLATE utf8_general_ci NULL DEFAULT NULL COMMENT '可以访问项目的权限组（白名单）',
    `order`                int(0) UNSIGNED                                         NULL DEFAULT 0 COMMENT '排序',
    `deleted`              tinyint(1)                                              NULL DEFAULT 0 COMMENT '删除标记',
    `template_code`        varchar(30) CHARACTER SET utf8 COLLATE utf8_general_ci  NULL DEFAULT '' COMMENT '项目类型',
    `schedule`             double(5, 2)                                            NULL DEFAULT 0.00 COMMENT '进度',
    `create_time`          varchar(255) CHARACTER SET utf8 COLLATE utf8_general_ci NULL DEFAULT NULL COMMENT '创建时间',
    `organization_code`    bigint(0)                                               NULL DEFAULT NULL COMMENT '组织id',
    `deleted_time`         varchar(30) CHARACTER SET utf8 COLLATE utf8_general_ci  NULL DEFAULT NULL COMMENT '删除时间',
    `private`              tinyint(1)                                              NULL DEFAULT 1 COMMENT '是否私有',
    `prefix`               varchar(10) CHARACTER SET utf8 COLLATE utf8_general_ci  NULL DEFAULT NULL COMMENT '项目前缀',
    `open_prefix`          tinyint(1)                                              NULL DEFAULT 0 COMMENT '是否开启项目前缀',
    `archive`              tinyint(1)                                              NULL DEFAULT 0 COMMENT '是否归档',
    `archive_time`         bigint(0)                                               NULL DEFAULT NULL COMMENT '归档时间',
    `open_begin_time`      tinyint(1)                                              NULL DEFAULT 0 COMMENT '是否开启任务开始时间',
    `open_task_private`    tinyint(1)                                              NULL DEFAULT 0 COMMENT '是否开启新任务默认开启隐私模式',
    `task_board_theme`     varchar(255) CHARACTER SET utf8 COLLATE utf8_general_ci NULL DEFAULT 'default' COMMENT '看板风格',
    `begin_time`           bigint(0)                                               NULL DEFAULT NULL COMMENT '项目开始日期',
    `end_time`             bigint(0)                                               NULL DEFAULT NULL COMMENT '项目截止日期',
    `auto_update_schedule` tinyint(1)                                              NULL DEFAULT 0 COMMENT '自动更新项目进度',
    PRIMARY KEY (`id`) USING BTREE,
    INDEX `project` (`order`) USING BTREE
) ENGINE = InnoDB
  AUTO_INCREMENT = 13043
  CHARACTER SET = utf8
  COLLATE = utf8_general_ci COMMENT = '项目表'
  ROW_FORMAT = COMPACT;