USE `exhibition-admin`;

CREATE TABLE IF NOT EXISTS `t_sms_code` (
    `id` BIGINT(20) NOT NULL AUTO_INCREMENT COMMENT 'ID',
    `phone` VARCHAR(20) NOT NULL COMMENT '手机号',
    `business_type` TINYINT(1) NOT NULL COMMENT '业务类型(1:验证码登录)',
    `code` VARCHAR(10) NOT NULL COMMENT '验证码',
    `status` TINYINT(1) NOT NULL DEFAULT 0 COMMENT '状态(0:未使用 1:已使用)',
    `created_at` BIGINT(20) NOT NULL COMMENT '创建时间',
    `expired_at` BIGINT(20) NOT NULL COMMENT '过期时间',
    `updated_at` BIGINT(20) NOT NULL COMMENT '更新时间',
    PRIMARY KEY (`id`),
    KEY `idx_phone_business_type_code` (`phone`, `business_type`, `code`)
) ENGINE = InnoDB DEFAULT COMMENT = '手机验证码';

DROP TABLE IF EXISTS `sys_user`;
CREATE TABLE `sys_user`  (
  `id` VARCHAR(40) NOT NULL,
  `dept_id` BIGINT(20) UNSIGNED NOT NULL DEFAULT 0 COMMENT '部门id',
  `user_name` VARCHAR(60) NOT NULL DEFAULT '' COMMENT '用户名',
  `user_nickname` VARCHAR(50) NOT NULL DEFAULT '' COMMENT '用户昵称',
  `user_password` VARCHAR(255) NOT NULL DEFAULT '' COMMENT '登录密码;cmf_password加密',
  `user_salt` CHAR(10) NOT NULL COMMENT '加密盐',
  `user_status` TINYINT(3) UNSIGNED NOT NULL DEFAULT 1 COMMENT '用户状态;0:禁用,1:正常,2:未验证',
  `is_admin` TINYINT(4) NOT NULL DEFAULT 1 COMMENT '是否后台管理员 1 是  0   否',

  `mobile` VARCHAR(20) NOT NULL DEFAULT '' COMMENT '中国手机不带国家代码，国际手机号格式为：国家代码-手机号',  
  `user_email` VARCHAR(100) NOT NULL DEFAULT '' COMMENT '用户登录邮箱',
  `sex` TINYINT(2) NOT NULL DEFAULT 0 COMMENT '性别;0:保密,1:男,2:女',
  `avatar` VARCHAR(255) NOT NULL DEFAULT '' COMMENT '用户头像',
  `city` VARCHAR(255) NOT NULL DEFAULT '' COMMENT '城市',
  `birthday` DATE NOT NULL DEFAULT '1900-01-01' COMMENT '生日',
  `address` VARCHAR(255) NOT NULL DEFAULT '' COMMENT '联系地址',
  `describe` VARCHAR(255) NOT NULL DEFAULT '' COMMENT ' 描述信息',
  `remark` VARCHAR(255) NOT NULL DEFAULT '' COMMENT '备注',

  `open_id` VARCHAR(50) NOT NULL DEFAULT '' COMMENT '微信open id',
  `iuqt_id` VARCHAR(50) DEFAULT '' COMMENT 'IUQT ID',
  `user_type` TINYINT(4) DEFAULT 0 COMMENT '用户类型(1:服务提供商,2:展商)',

  `card_type` VARCHAR(20) NOT NULL DEFAULT '' COMMENT '证件类型(居民身份证,港澳居民往来内地通行证,台湾居民往来大陆通行证)',
  `id_card` VARCHAR(512) NOT NULL DEFAULT '' COMMENT '证件号',
  `real_name` VARCHAR(255) NOT NULL DEFAULT '' COMMENT '真实姓名',

  `last_login_ip` VARCHAR(15) NOT NULL DEFAULT '' COMMENT '最后登录ip',
  `last_login_time` DATETIME NULL DEFAULT NULL COMMENT '最后登录时间',

  `created_at` DATETIME NULL DEFAULT NULL COMMENT '创建时间',
  `updated_at` DATETIME NULL DEFAULT NULL COMMENT '更新时间',
  `deleted_at` DATETIME NULL DEFAULT NULL COMMENT '删除时间',
  PRIMARY KEY (`id`) USING BTREE,
  UNIQUE INDEX `user_login`(`user_name`, `deleted_at`) USING BTREE,
  UNIQUE INDEX `mobile`(`mobile`, `deleted_at`) USING BTREE,
  INDEX `user_nickname`(`user_nickname`) USING BTREE,
  INDEX `open_id`(`open_id`) USING BTREE,
  INDEX `iuqt_id`(`iuqt_id`) USING BTREE
) ENGINE = InnoDB COMMENT = '用户表';

ALTER TABLE `sys_user` DROP INDEX `iuqt_id`;
ALTER TABLE `sys_user` ADD INDEX `iuqt_id`(`iuqt_id`) USING BTREE;
ALTER TABLE `sys_user` ADD COLUMN `city` VARCHAR(255) NOT NULL DEFAULT '' COMMENT '城市';
ALTER TABLE `sys_user` ADD COLUMN `birthday` DATE NOT NULL DEFAULT '1900-01-01' COMMENT '生日';
ALTER TABLE `sys_user` ADD COLUMN `id_card` VARCHAR(512) NOT NULL DEFAULT '' COMMENT '身份证号';
ALTER TABLE `sys_user` ADD COLUMN `card_type` VARCHAR(20) NOT NULL DEFAULT '' COMMENT '证件类型(居民身份证,港澳居民往来内地通行证,台湾居民往来大陆通行证)';
ALTER TABLE `sys_user` ADD COLUMN `real_name` VARCHAR(255) NOT NULL DEFAULT '' COMMENT '真实姓名';

-- ----------------------------
-- Records of sys_user
-- ----------------------------
INSERT INTO `sys_user` VALUES 
('1', 'admin', '13578342363', '超级管理员', 'c567ae329f9929b518759d3bea13f492', 'f9aZTAa8yz', 1, 'yxh669@qq.com', 1, 
'https://yxh-1301841944.cos.ap-chongqing.myqcloud.com/gfast/2021-07-19/ccwpeuqz1i2s769hua.jpeg', 101, '', 1, 'asdasfdsaf大发放打发士大夫发按时', '描述信息', 
'::1', '2023-10-31 11:22:06', '2021-06-22 17:58:00', '2023-04-22 14:39:18', NULL, '', '', 0);


-- 协议主表
CREATE TABLE IF NOT EXISTS `t_agreement` (
  `id` BIGINT(20) NOT NULL AUTO_INCREMENT,
  `name` VARCHAR(255) NOT NULL COMMENT '协议名称',
  `major_version` VARCHAR(255) NOT NULL DEFAULT '' COMMENT '主版本号(如1.0.0)',
  `minor_version` VARCHAR(255) NOT NULL DEFAULT '' COMMENT '次版本号(如1.1.0)',
  `patch_version` VARCHAR(255) NOT NULL DEFAULT '' COMMENT '补丁版本号(如1.1.1)',
  `content` MEDIUMTEXT NOT NULL COMMENT 'HTML格式协议内容',
  `version` INT(11) NOT NULL DEFAULT 0 COMMENT '并发版本控制',
  `created_at` BIGINT(20) NOT NULL COMMENT '创建时间',
  `updated_at` BIGINT(20) NOT NULL COMMENT '更新时间',
  PRIMARY KEY (`id`),
  UNIQUE INDEX `idx_name_major_version`(`name`, `major_version`),
  UNIQUE INDEX `idx_name_minor_version`(`name`, `minor_version`),
  UNIQUE INDEX `idx_name_patch_version`(`name`, `patch_version`)
) ENGINE=InnoDB  COMMENT = '协议主表';

-- 用户同意记录表
CREATE TABLE IF NOT EXISTS `t_user_agreement` (
  `id` BIGINT(20) NOT NULL AUTO_INCREMENT ,
  `user_id` VARCHAR(40) NOT NULL COMMENT '用户ID',
  `agreement_id` BIGINT(20) NOT NULL COMMENT '协议ID',
  `agreement_name` VARCHAR(255) NOT NULL COMMENT '协议名称',
  `agreed` TINYINT(1) NOT NULL DEFAULT 0 COMMENT '是否同意(0:不同意 1:同意)',
  `created_at` BIGINT(20) NOT NULL COMMENT '同意时间',
  PRIMARY KEY (`id`),
  INDEX `idx_user_id` (`user_id`)
) ENGINE=InnoDB COMMENT = '用户同意记录表';

CREATE TABLE IF NOT EXISTS `t_user_device` (
  `id` BIGINT(20) NOT NULL AUTO_INCREMENT,
  `user_id` VARCHAR(40) NOT NULL COMMENT '用户ID',
  `device_id` VARCHAR(255) NOT NULL COMMENT '设备ID',
  `device_name` VARCHAR(255) NOT NULL COMMENT '设备名称',
  `device_ip` VARCHAR(255) NOT NULL COMMENT '设备IP',
  `login_type` VARCHAR(255) NOT NULL COMMENT '登录类型(手机密码, 手机验证码)',
  `created_at` BIGINT(20) NOT NULL COMMENT '创建时间',
  `updated_at` BIGINT(20) NOT NULL COMMENT '更新时间',
  PRIMARY KEY (`id`),
  KEY `idx_user_id` (`user_id`),
  KEY `idx_device_id` (`device_id`),
  UNIQUE KEY `idx_user_id_device_id` (`user_id`, `device_id`)
)ENGINE=InnoDB COMMENT = '用户登录设备表';