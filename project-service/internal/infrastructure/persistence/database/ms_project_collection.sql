CREATE TABLE `ms_project_collection`
(
    `id`           bigint(20) NOT NULL AUTO_INCREMENT,
    `project_code` bigint(20) NULL DEFAULT 0 COMMENT '项目id',
    `member_code`  bigint(20) NULL DEFAULT 0 COMMENT '成员id',
    `create_time`  bigint(20) NULL DEFAULT 0 COMMENT '加入时间',
    PRIMARY KEY (`id`) USING BTREE
) ENGINE = InnoDB
  AUTO_INCREMENT = 46
  CHARACTER SET = utf8
  COLLATE = utf8_general_ci COMMENT = '项目-收藏表'
  ROW_FORMAT = COMPACT;