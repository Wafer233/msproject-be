CREATE TABLE `ms_project_member`
(
    `id`           bigint(0)                                               NOT NULL AUTO_INCREMENT,
    `project_code` bigint(0)                                               NULL DEFAULT NULL COMMENT '项目id',
    `member_code`  bigint(0)                                               NULL DEFAULT NULL COMMENT '成员id',
    `join_time`    bigint(0)                                               NULL DEFAULT NULL COMMENT '加入时间',
    `is_owner`     bigint(0)                                               NULL DEFAULT 0 COMMENT '拥有者',
    `authorize`    varchar(255) CHARACTER SET utf8 COLLATE utf8_general_ci NULL DEFAULT NULL COMMENT '角色',
    PRIMARY KEY (`id`) USING BTREE,
    UNIQUE INDEX `unique` (`project_code`, `member_code`) USING BTREE
) ENGINE = InnoDB
  AUTO_INCREMENT = 37
  CHARACTER SET = utf8
  COLLATE = utf8_general_ci COMMENT = '项目-成员表'
  ROW_FORMAT = COMPACT;