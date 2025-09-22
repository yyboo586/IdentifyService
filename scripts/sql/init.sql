SET FOREIGN_KEY_CHECKS = 0;

CREATE DATABASE IF NOT EXISTS `identify_service` DEFAULT CHARACTER SET utf8mb4 COLLATE utf8mb4_unicode_ci;

USE identify_service;

CREATE TABLE IF NOT EXISTS `t_org`  (
  `id` VARCHAR(40) NOT NULL COMMENT '组织id',
  `pid` VARCHAR(40) NOT NULL COMMENT '父级ID',
  `name` VARCHAR(255) NOT NULL COMMENT '组织名称',
  `manager_id` VARCHAR(40) NOT NULL DEFAULT '' COMMENT '负责人ID',
  `manager_name` VARCHAR(255) NOT NULL DEFAULT '' COMMENT '负责人名称',
  `status` TINYINT(1) NOT NULL DEFAULT 0 COMMENT '状态',
  `created_at` DATETIME NOT NULL DEFAULT CURRENT_TIMESTAMP COMMENT '创建时间',
  `updated_at` DATETIME NOT NULL DEFAULT CURRENT_TIMESTAMP ON UPDATE CURRENT_TIMESTAMP COMMENT '修改时间',
  PRIMARY KEY (`id`),
  INDEX `idx_pid`(`pid`),
  KEY `idx_manager_id`(`manager_id`)
) ENGINE = InnoDB COMMENT = '组织表';

CREATE TABLE IF NOT EXISTS `t_user`  (
  `id` VARCHAR(40) NOT NULL,
  `name` VARCHAR(255) NOT NULL COMMENT '账户名',
  `nickname` VARCHAR(50) NOT NULL COMMENT '账户昵称',
  `password` VARCHAR(255) NOT NULL COMMENT '账户密码',
  `salt` CHAR(10) NOT NULL COMMENT '盐值',
  `status` TINYINT(4) NOT NULL COMMENT '账户状态',
  `org_id` VARCHAR(40) NOT NULL COMMENT '组织ID',
  `sex` TINYINT(4) NOT NULL DEFAULT 0 COMMENT '性别',
  `email` VARCHAR(100) NOT NULL DEFAULT '' COMMENT '邮箱',  
  `avatar` VARCHAR(255) NOT NULL DEFAULT '' COMMENT '头像',   
  `mobile` VARCHAR(20) NOT NULL DEFAULT '' COMMENT '手机号码',
  `address` VARCHAR(255) NOT NULL DEFAULT '' COMMENT '地址',
  `describe` VARCHAR(255) NOT NULL DEFAULT '' COMMENT '描述',
  `is_admin` TINYINT(4) NOT NULL DEFAULT 0 COMMENT '是否管理员',
  `created_at` DATETIME NOT NULL DEFAULT CURRENT_TIMESTAMP,
  `updated_at` DATETIME NOT NULL DEFAULT CURRENT_TIMESTAMP ON UPDATE CURRENT_TIMESTAMP,
  PRIMARY KEY (`id`),
  UNIQUE INDEX `idx_name`(`name`),
  INDEX `idx_ord_status`(`org_id`, `status`)
) ENGINE = InnoDB COMMENT = '用户表';

CREATE TABLE IF NOT EXISTS `t_role`  (
  `id` INT(10) NOT NULL AUTO_INCREMENT,
  `pid` INT(10) NOT NULL COMMENT '父级ID',
  `org_id` VARCHAR(40) NOT NULL COMMENT '组织ID',
  `name` VARCHAR(20) NOT NULL COMMENT '角色名称',
  `creator_id` VARCHAR(40) NOT NULL COMMENT '创建人ID',
  `deletor_id` VARCHAR(40) DEFAULT '' COMMENT '删除人ID',
  `created_at` DATETIME DEFAULT CURRENT_TIMESTAMP COMMENT '创建时间',
  `updated_at` DATETIME DEFAULT CURRENT_TIMESTAMP ON UPDATE CURRENT_TIMESTAMP COMMENT '更新时间',
  `deleted_at` DATETIME DEFAULT NULL COMMENT '删除时间',
  PRIMARY KEY (`id`),
  UNIQUE INDEX `idx_org_id_name`(`org_id`, `name`),
  INDEX `idx_pid`(`pid`),
  INDEX `idx_name_deleted_at`(`name`,`deleted_at`)
) ENGINE = InnoDB COMMENT = '角色表';

CREATE TABLE IF NOT EXISTS `t_resource` (
  `id` BIGINT(20) NOT NULL AUTO_INCREMENT,
  `type` VARCHAR(20) NOT NULL COMMENT '资源类型',
  `code` VARCHAR(100) NOT NULL COMMENT '资源代码',
  `created_at` DATETIME NOT NULL DEFAULT CURRENT_TIMESTAMP,
  `updated_at` DATETIME NOT NULL DEFAULT CURRENT_TIMESTAMP ON UPDATE CURRENT_TIMESTAMP,
  PRIMARY KEY (`id`),
  UNIQUE INDEX `idx_type_code`(`type`, `code`)
)ENGINE = InnoDB COMMENT = '资源表';

CREATE TABLE IF NOT EXISTS `t_permission` (
  `id` BIGINT(20) NOT NULL AUTO_INCREMENT,
  `code` VARCHAR(100) NOT NULL COMMENT '权限点代码',
  `code_name` VARCHAR(100) NOT NULL COMMENT '权限点名称',
  `resource_id` BIGINT(20) NOT NULL COMMENT '资源ID',
  `description` TEXT COMMENT '权限点描述',
  `created_at` DATETIME NOT NULL DEFAULT CURRENT_TIMESTAMP,
  `updated_at` DATETIME NOT NULL DEFAULT CURRENT_TIMESTAMP ON UPDATE CURRENT_TIMESTAMP,
  PRIMARY KEY (`id`),
  UNIQUE INDEX `idx_code_resource_id`(`code`,`resource_id`)
) ENGINE = InnoDB COMMENT = '权限点表';

CREATE TABLE IF NOT EXISTS `t_permission_resource` (
  `id` BIGINT(20) NOT NULL AUTO_INCREMENT,
  `permission_id` BIGINT(20) NOT NULL COMMENT '权限点ID',
  `resource_id` BIGINT(20) NOT NULL COMMENT '资源ID',
  `created_at` DATETIME NOT NULL DEFAULT CURRENT_TIMESTAMP,
  PRIMARY KEY (`id`),
  UNIQUE INDEX `idx_permission_id_resource_id`(`permission_id`, `resource_id`)
) ENGINE = InnoDB COMMENT = '权限点资源表';

INSERT INTO `t_org` (`id`, `pid`, `name`, `manager_id`, `manager_name`, `status`)
VALUES ("00000000-0000-0000-0000-000000000000", '', 'Org-Background-Default', '11111111-0000-0000-0000-000000000000', 'admin', 1);

INSERT INTO `t_user` (`id`, `name`, `nickname`, `password`, `salt`, `status`, `org_id`)
VALUES ('11111111-0000-0000-0000-000000000000', 'admin', 'admin', 'c567ae329f9929b518759d3bea13f492', 'f9aZTAa8yz', 1, '00000000-0000-0000-0000-000000000000');

INSERT INTO `t_role` (`id`, `pid`, `org_id`, `name`, `creator_id`) 
VALUES (1, 0, '00000000-0000-0000-0000-000000000000', 'SuperAdmin', '11111111-0000-0000-0000-000000000000');

INSERT INTO `t_casbin_rule` (`ptype`, `v0`, `v1`)
VALUES ('g', 'u_11111111-0000-0000-0000-000000000000', '1');


CREATE TABLE IF NOT EXISTS `sys_role_scope`  (
  `id` BIGINT(20) UNSIGNED NOT NULL AUTO_INCREMENT COMMENT 'ID',
  `role_id` BIGINT(20) UNSIGNED NOT NULL DEFAULT 0 COMMENT '角色id',
  `menu_id` INT(11) NOT NULL COMMENT 'api接口id',
  `data_scope` INT(11) NOT NULL COMMENT '数据范围(1:全部数据权限 2:自定数据权限 3:本部门数据权限 4:本部门及以下数据权限)',
  `dept_ids` json NULL COMMENT '扩展数据',
  PRIMARY KEY (`id`),
  UNIQUE INDEX `role_id`(`role_id`, `menu_id`)
) ENGINE = InnoDB COMMENT = '角色数据权限';

CREATE TABLE IF NOT EXISTS `sys_config`  (
  `config_id` INT(5) UNSIGNED NOT NULL AUTO_INCREMENT COMMENT '参数主键',
  `config_name` VARCHAR(100) NOT NULL DEFAULT '' COMMENT '参数名称',
  `config_key` VARCHAR(100) NOT NULL DEFAULT '' COMMENT '参数键名',
  `config_value` VARCHAR(500) NOT NULL DEFAULT '' COMMENT '参数键值',
  `config_type` TINYINT(1) NOT NULL DEFAULT 0 COMMENT '系统内置(Y是 N否)',
  `create_by` INT(64) UNSIGNED NOT NULL DEFAULT 0 COMMENT '创建者',
  `update_by` INT(64) UNSIGNED NOT NULL DEFAULT 0 COMMENT '更新者',
  `remark` VARCHAR(500) NOT NULL DEFAULT '' COMMENT '备注',
  `created_at` DATETIME NOT NULL DEFAULT CURRENT_TIMESTAMP COMMENT '创建时间',
  `updated_at` DATETIME NOT NULL DEFAULT CURRENT_TIMESTAMP ON UPDATE CURRENT_TIMESTAMP COMMENT '修改时间',
  PRIMARY KEY (`config_id`),
  UNIQUE INDEX `uni_config_key`(`config_key`)
) ENGINE = InnoDB COMMENT = '参数配置表';

CREATE TABLE IF NOT EXISTS `sys_job`  (
  `job_id` BIGINT(20) UNSIGNED NOT NULL AUTO_INCREMENT COMMENT '任务ID',
  `job_name` VARCHAR(64) NOT NULL DEFAULT '' COMMENT '任务名称',
  `job_params` VARCHAR(200) NOT NULL DEFAULT '' COMMENT '参数',
  `job_group` VARCHAR(64) NOT NULL DEFAULT 'DEFAULT' COMMENT '任务组名',
  `invoke_target` VARCHAR(100) NOT NULL COMMENT '调用目标字符串',
  `cron_expression` VARCHAR(255) NOT NULL DEFAULT '' COMMENT 'cron执行表达式',
  `misfire_policy` TINYINT(4) NOT NULL DEFAULT 1 COMMENT '计划执行策略(1多次执行 2执行一次)',
  `concurrent` TINYINT(4) NOT NULL DEFAULT 1 COMMENT '是否并发执行(0允许 1禁止)',
  `status` TINYINT(4) NOT NULL DEFAULT 0 COMMENT '状态(0正常 1暂停)',
  `created_by` BIGINT(64) UNSIGNED NOT NULL DEFAULT 0 COMMENT '创建者',
  `updated_by` BIGINT(64) UNSIGNED NOT NULL DEFAULT 0 COMMENT '更新者',
  `remark` VARCHAR(500) NOT NULL DEFAULT '' COMMENT '备注信息',
  `created_at` DATETIME NOT NULL DEFAULT CURRENT_TIMESTAMP COMMENT '创建时间',
  `updated_at` DATETIME NOT NULL DEFAULT CURRENT_TIMESTAMP ON UPDATE CURRENT_TIMESTAMP COMMENT '更新时间',
  PRIMARY KEY (`job_id`),
  UNIQUE INDEX `invoke_target`(`invoke_target`)
) ENGINE = InnoDB COMMENT = '定时任务调度表';

CREATE TABLE IF NOT EXISTS `sys_job_log`  (
  `id` BIGINT(20) UNSIGNED NOT NULL AUTO_INCREMENT COMMENT '主键',
  `target_name` VARCHAR(100) NOT NULL COMMENT '方法名',
  `created_at` DATETIME NOT NULL DEFAULT CURRENT_TIMESTAMP COMMENT '执行日期',
  `result` VARCHAR(100) NOT NULL DEFAULT '' COMMENT '执行结果',
  PRIMARY KEY (`id`)
) ENGINE = InnoDB COMMENT = '任务日志表';


CREATE TABLE IF NOT EXISTS `sys_notice`  (
  `id` BIGINT(20) NOT NULL AUTO_INCREMENT COMMENT 'ID',
  `title` VARCHAR(64) NOT NULL COMMENT '标题',
  `type` BIGINT(20) NOT NULL COMMENT '类型',
  `tag` INT(11) NULL DEFAULT NULL COMMENT '标签',
  `content` LONGTEXT NOT NULL COMMENT '内容',
  `remark` VARCHAR(255) NULL DEFAULT NULL COMMENT '备注',
  `sort` INT(11) NOT NULL DEFAULT 0 COMMENT '排序',
  `status` TINYINT(1) NOT NULL DEFAULT 1 COMMENT '状态',
  `created_by` BIGINT(20) NULL DEFAULT NULL COMMENT '发送人',
  `updated_by` BIGINT(20) NULL DEFAULT 0 COMMENT '修改人',
  `created_at` DATETIME NOT NULL DEFAULT CURRENT_TIMESTAMP COMMENT '创建时间',
  `updated_at` DATETIME NOT NULL DEFAULT CURRENT_TIMESTAMP ON UPDATE CURRENT_TIMESTAMP COMMENT '更新时间',
  `deleted_at` DATETIME NULL DEFAULT NULL COMMENT '删除时间',
  `receiver` JSON NULL COMMENT '接收者（私信）',
  PRIMARY KEY (`id`) USING BTREE
) ENGINE = InnoDB COMMENT = '通知公告';

CREATE TABLE IF NOT EXISTS `sys_notice_read`  (
  `id` BIGINT(20) NOT NULL AUTO_INCREMENT COMMENT 'id',
  `notice_id` BIGINT(20) NOT NULL COMMENT '信息id',
  `user_id` VARCHAR(40) NOT NULL COMMENT '用户id',
  `clicks` INT(11) NULL DEFAULT NULL COMMENT '点击次数',
  `created_at` DATETIME NOT NULL DEFAULT CURRENT_TIMESTAMP COMMENT '阅读时间',
  `updated_at` DATETIME NOT NULL DEFAULT CURRENT_TIMESTAMP ON UPDATE CURRENT_TIMESTAMP COMMENT '更新时间',
  PRIMARY KEY (`id`),
  UNIQUE INDEX `notice_id`(`notice_id`, `user_id`)
) ENGINE = InnoDB COMMENT = '已读记录';

SET FOREIGN_KEY_CHECKS = 1;

