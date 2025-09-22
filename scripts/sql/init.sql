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
  KEY `idx_pid`(`pid`),
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

CREATE TABLE IF NOT EXISTS `t_casbin_rule`  (
  `ptype` VARCHAR(10) NOT NULL DEFAULT '',
  `v0` VARCHAR(256) NOT NULL DEFAULT '',
  `v1` VARCHAR(256) NOT NULL DEFAULT '',
  `v2` VARCHAR(256) NOT NULL DEFAULT '',
  `v3` VARCHAR(256) NOT NULL DEFAULT '',
  `v4` VARCHAR(256) NOT NULL DEFAULT '',
  `v5` VARCHAR(256) NOT NULL DEFAULT '',
  KEY `idx_ptype_v0` (`ptype`, `v0`),
  KEY `idx_ptype_v1` (`ptype`, `v1`)
) ENGINE = InnoDB COMMENT = 'casbin规则表';

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

INSERT INTO `t_org` (`id`, `pid`, `name`, `manager_id`, `manager_name`, `status`)
VALUES ("00000000-0000-0000-0000-000000000000", '', 'Org-Background-Default', '11111111-0000-0000-0000-000000000000', 'admin', 1);

INSERT INTO `t_user` (`id`, `name`, `nickname`, `password`, `salt`, `status`, `org_id`)
VALUES ('11111111-0000-0000-0000-000000000000', 'admin', 'admin', 'c567ae329f9929b518759d3bea13f492', 'f9aZTAa8yz', 1, '00000000-0000-0000-0000-000000000000');


SET FOREIGN_KEY_CHECKS = 1;

