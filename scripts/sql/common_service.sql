CREATE DATABASE IF NOT EXISTS `identify_service` DEFAULT CHARACTER SET utf8mb4 COLLATE utf8mb4_unicode_ci;

USE identify_service;

CREATE TABLE IF NOT EXISTS `t_token_blacklist` (
  `id` VARCHAR(40) NOT NULL COMMENT '令牌ID',
  `operator_id` VARCHAR(40) DEFAULT '' COMMENT '操作者ID',
  `created_at` DATETIME NOT NULL DEFAULT CURRENT_TIMESTAMP COMMENT '创建时间',
  PRIMARY KEY (`id`)
)ENGINE = InnoDB COMMENT = '令牌黑名单';

CREATE TABLE IF NOT EXISTS `t_log` (
  `id` BIGINT(20) NOT NULL AUTO_INCREMENT COMMENT '日志主键',
  `org_id` VARCHAR(50) NOT NULL DEFAULT '' COMMENT '组织ID',
  `user_id` VARCHAR(50) NOT NULL DEFAULT '' COMMENT '用户ID',
  `user_name` VARCHAR(50) NOT NULL DEFAULT '' COMMENT '用户名',
  `ip` VARCHAR(50) NOT NULL DEFAULT '' COMMENT 'IP地址',
  `type` TINYINT(4) NOT NULL DEFAULT 0 COMMENT '日志类型',
  `content` TEXT NOT NULL COMMENT '日志内容',
  `created_at` DATETIME NOT NULL DEFAULT CURRENT_TIMESTAMP COMMENT '创建时间',
  PRIMARY KEY (`id`),
  INDEX `idx_org_id_type`(`org_id`, `type`)
) ENGINE = InnoDB COMMENT = '日志表';

