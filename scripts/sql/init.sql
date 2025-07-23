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

INSERT INTO `t_org` 
(`id`,                                   `pid`,  `name`,                   `manager_id`,                           `manager_name`, `status`)
VALUES 
("00000000-0000-0000-0000-000000000000", '0',    'Org-Background-Default', '00000000-0000-0000-0000-000000000001', 'admin',         1),
("11111111-0000-0000-0000-000000000000", '0',    'Org-Front-Default',      '00000000-0000-0000-0000-000000000001', 'admin',         1);


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
  INDEX `idx_ord_id`(`org_id`)
) ENGINE = InnoDB COMMENT = '用户表';

INSERT INTO `t_user` 
(`id`,                                   `name`,  `nickname`, `password`,                         `salt`,      `status`, `org_id`)
VALUES 
('00000000-0000-0000-0000-000000000001', 'admin', 'admin',    'c567ae329f9929b518759d3bea13f492', 'f9aZTAa8yz', 1,       '00000000-0000-0000-0000-000000000000');

CREATE TABLE IF NOT EXISTS `t_role`  (
  `id` INT(10) NOT NULL AUTO_INCREMENT,
  `org_id` VARCHAR(40) NOT NULL COMMENT '组织ID',
  `pid` INT(10) NOT NULL COMMENT '父级ID',
  `name` VARCHAR(20) NOT NULL COMMENT '角色名称',
  `status` TINYINT(4) NOT NULL COMMENT '状态',
  `creator_id` VARCHAR(40) NOT NULL COMMENT '创建人ID',
  `created_at` DATETIME DEFAULT CURRENT_TIMESTAMP COMMENT '创建时间',
  `updated_at` DATETIME DEFAULT CURRENT_TIMESTAMP ON UPDATE CURRENT_TIMESTAMP COMMENT '更新时间',
  PRIMARY KEY (`id`),
  INDEX `idx_pid`(`pid`),
  INDEX `idx_name`(`name`),
  INDEX `idx_creator_id`(`creator_id`)
) ENGINE = InnoDB COMMENT = '角色表';

INSERT INTO `t_role`
(`id`, `pid`, `org_id`,                               `name`,          `status`, `creator_id`) 
VALUES 
(1,    0,     '00000000-0000-0000-0000-000000000000', 'SuperAdmin',    1,        '00000000-0000-0000-0000-000000000001'),
(2,    0,     '11111111-0000-0000-0000-000000000000', 'FrontOrgAdmin', 1,        '00000000-0000-0000-0000-000000000001');


CREATE TABLE IF NOT EXISTS `t_auth_rule`  (
  `id` BIGINT(20) NOT NULL AUTO_INCREMENT,
  `pid` BIGINT(20) NOT NULL DEFAULT 0,
  `name` VARCHAR(100) NOT NULL DEFAULT '',
  `type` TINYINT(1) NOT NULL DEFAULT 0,
  `path` VARCHAR(100) NOT NULL DEFAULT '',
  `component` VARCHAR(100) NOT NULL DEFAULT '',
  
  `title` VARCHAR(50) NOT NULL DEFAULT '',
  `icon` VARCHAR(300) NOT NULL DEFAULT '',
  `active_icon` VARCHAR(300) NOT NULL DEFAULT '',
  `keep_alive` TINYINT(4) NOT NULL DEFAULT 0,
  `hide_in_menu` TINYINT(4) NOT NULL DEFAULT 0,
  `hide_in_tab` TINYINT(4) NOT NULL DEFAULT 0,
  `hide_in_breadcrumb` TINYINT(4) NOT NULL DEFAULT 0,
  `hide_children_in_menu` TINYINT(4) NOT NULL DEFAULT 0,
  `authority` VARCHAR(255) NOT NULL DEFAULT '',
  `badge` VARCHAR(255) NOT NULL DEFAULT '',
  `badge_type` VARCHAR(255) NOT NULL DEFAULT 'normal',
  `badge_variants` VARCHAR(255) NOT NULL DEFAULT '',
  `full_path_key` TINYINT(4)  NOT NULL DEFAULT 0,
  `active_path` VARCHAR(255) NOT NULL DEFAULT '',
  `affix_tab` TINYINT(4)  NOT NULL DEFAULT 0,
  `affix_tab_order` TINYINT(4) NOT NULL DEFAULT 0,
  `iframe_src` VARCHAR(255) NOT NULL DEFAULT '',
  `ignore_access` TINYINT(4) NOT NULL DEFAULT 0,
  `link` VARCHAR(255) NOT NULL DEFAULT '',
  `max_num_of_open_tab` TINYINT(4) NOT NULL DEFAULT -1,
  `menu_visible_with_forbidden` TINYINT(4) NOT NULL DEFAULT 0,
  `open_in_new_window` TINYINT(1) NOT NULL DEFAULT 0,
  `order` INT NOT NULL DEFAULT 0,
  `query` TEXT,
  `no_basic_layout` TINYINT(4) NOT NULL DEFAULT 0,
  
  `created_at` DATETIME NOT NULL DEFAULT CURRENT_TIMESTAMP,
  `updated_at` DATETIME NOT NULL DEFAULT CURRENT_TIMESTAMP ON UPDATE CURRENT_TIMESTAMP,
  PRIMARY KEY (`id`),
  UNIQUE INDEX `idx_name`(`name`),
  INDEX `pid`(`pid`)
) ENGINE = InnoDB COMMENT = '菜单表';

INSERT INTO `t_auth_rule` 
VALUES 
(1,0,'Dashboard',0,'/dashboard','','page.dashboard.title','lucide:layout-dashboard','',0,0,0,0,0,'null','','','',0,'',0,0,'',0,'',0,0,0,-1,NULL,0,'2025-07-30 13:40:19','2025-07-30 13:40:19'),
(2,1,'Analytics',1,'/analytics','/dashboard/analytics/index','page.dashboard.analytics','lucide:area-chart','',0,0,0,0,0,'[\"null\"]','','','',0,'/analytics',1,0,'',0,'',0,0,0,0,'',0,'2025-07-30 13:46:41','2025-07-31 14:30:36'),
(3,0,'System',0,'/system','','system.title','ion:settings-outline','',0,0,0,0,0,'null','','','',0,'',0,0,'',0,'',0,0,0,9997,NULL,0,'2025-07-30 13:50:43','2025-07-30 13:50:43'),
(4,3,'SystemMenu',1,'/system/menu','/system/menu/list','system.menu.title','mdi:menu','',0,0,0,0,0,'[\"null\"]','','','',0,'/system/menu',0,0,'',0,'',0,0,0,0,'',0,'2025-07-30 13:53:06','2025-07-31 14:29:58'),
(5,3,'SystemDept',1,'/system/dept','../../views/system/dept/list.vue','system.dept.title','charm:organisation','',0,0,0,0,0,'null','','','',0,'',0,0,'',0,'',0,0,0,0,NULL,0,'2025-07-30 13:58:16','2025-07-30 13:58:16'),
(6,3,'SystemRole',1,'/system/role','../../views/system/role/list.vue','system.role.title','mdi:account-group','',0,0,0,0,0,'[\"null\"]','','','',0,'/system/role',0,0,'',0,'',0,0,0,0,'',0,'2025-07-30 13:58:44','2025-07-31 14:30:18'),
(7,3,'SystemUser',1,'/system/user','../../views/system/user/list.vue','system.user.title','carbon:user','',0,0,0,0,0,'[\"null\"]','','','',0,'/system/user',0,0,'',0,'',0,0,0,0,'',0,'2025-07-30 14:00:55','2025-07-31 14:30:11'),
(8,0,'Device',1,'/device','../../views/device/list.vue','device.title','carbon:tool-kit','',0,0,0,0,0,'null','','','',0,'/device',0,0,'',0,'',0,0,0,0,NULL,0,'2025-07-31 16:05:27','2025-07-31 16:05:27'),
(9,0,'LogManagement',0,'/logmanagement','','log.title','carbon:account','',0,0,0,0,0,'null','','','',0,'',0,0,'',0,'',0,0,0,0,NULL,0,'2025-08-02 13:55:01','2025-08-02 13:55:01'),
(10,0,'configration',0,'/configration','','config.title','carbon:task-settings','',0,0,0,0,0,'[\"null\"]','','','',0,'',0,0,'',0,'',0,0,0,0,'',0,'2025-08-02 14:05:42','2025-08-02 14:07:27'),
(11,10,'AlarmConfiguration',1,'/alarmconfiguration','../../views/configration/alarm/list.vue','config.Alarm.title','','',0,0,0,0,0,'null','','','',0,'/alarmconfiguration',0,0,'',0,'',0,0,0,0,NULL,0,'2025-08-02 16:21:45','2025-08-02 16:21:45'),
(12,9,'AlarmLog',1,'/alarmlog','../../views/log/alarm/list.vue','log.alarm.title','carbon:notification','',0,0,0,0,0,'null','','','',0,'/alarmlog',0,0,'',0,'',0,0,0,0,NULL,0,'2025-08-04 14:17:55','2025-08-04 14:17:55'),
(13,9,'ActionLog',1,'/actionlog','../../views/log/action/list.vue','log.action.title','carbon:touch-interaction','',0,0,0,0,0,'null','','','',0,'/actionlog',0,0,'',0,'',0,0,0,0,NULL,0,'2025-08-04 17:10:20','2025-08-04 17:10:20'),
(14,9,'LoginLog',1,'/loginLog','../../views/log/login/list.vue','log.login.title','carbon:login','',0,0,0,0,0,'null','','','',0,'/loginLog',0,0,'',0,'',0,0,0,0,NULL,0,'2025-08-04 17:18:41','2025-08-04 17:18:41');


-- ptype: 策略的类型标识
-- p: 基本的权限策略，定义 主体 对 资源 的 权限。p, tom, resource1, read 表示 tom 对 resource1 有 read 权限
-- g: 角色继承关系。基础角色的稳定性。
CREATE TABLE IF NOT EXISTS `t_casbin_rule`  (
  `ptype` VARCHAR(10) NOT NULL DEFAULT '',
  `v0` VARCHAR(256) NOT NULL DEFAULT '',
  `v1` VARCHAR(256) NOT NULL DEFAULT '',
  `v2` VARCHAR(256) NOT NULL DEFAULT '',
  `v3` VARCHAR(256) NOT NULL DEFAULT '',
  `v4` VARCHAR(256) NOT NULL DEFAULT '',
  `v5` VARCHAR(256) NOT NULL DEFAULT ''
) ENGINE = InnoDB COMMENT = 'casbin规则表';

INSERT INTO `t_casbin_rule`
(`ptype`, `v0`, `v1`, `v2`)
VALUES 
('p', '2', '1', 'ALL'), 
('p', '2', '2', 'ALL'), 
('p', '2', '3', 'ALL'),
('p', '2', '4', 'ALL'),
('p', '2', '6', 'ALL'), 
('p', '2', '7', 'ALL'), 
('p', '2', '8', 'ALL'),
('p', '2', '9', 'ALL'),
('p', '2', '10', 'ALL'),
('p', '2', '11', 'ALL'), 
('p', '2', '12', 'ALL'), 
('p', '2', '13', 'ALL'),
('p', '2', '14', 'ALL');

INSERT INTO `t_casbin_rule`
(`ptype`, `v0`, `v1`)
VALUES 
('g', 'u_00000000-0000-0000-0000-000000000001', '1');





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

CREATE TABLE IF NOT EXISTS `t_login_log`  (
  `id` BIGINT(20) NOT NULL AUTO_INCREMENT,
  `org_id` VARCHAR(40) NOT NULL COMMENT '组织ID',
  `login_name` VARCHAR(50) NOT NULL DEFAULT '' COMMENT '登录账号',
  `ip` VARCHAR(50) NOT NULL DEFAULT '' COMMENT '登录IP地址',
  `browser` VARCHAR(512) NOT NULL DEFAULT '' COMMENT '浏览器类型',
  `status` TINYINT(4) NOT NULL DEFAULT 0 COMMENT '登录状态',
  `message` VARCHAR(255) NOT NULL DEFAULT '' COMMENT '提示消息',
  `login_time` DATETIME NOT NULL COMMENT '登录时间',
  `created_at` DATETIME NOT NULL DEFAULT CURRENT_TIMESTAMP COMMENT '创建时间',
  PRIMARY KEY (`id`),
  INDEX `idx_org_id_name_time` (`org_id`, `login_name`, `login_time`),
  INDEX `idx_org_id_time`(`org_id`, `login_time`)
) ENGINE = InnoDB COMMENT = '登录日志表';

CREATE TABLE IF NOT EXISTS `t_oper_log`  (
  `id` BIGINT(20) UNSIGNED NOT NULL AUTO_INCREMENT COMMENT '日志主键',
  `org_id` VARCHAR(50) NOT NULL DEFAULT '' COMMENT '组织ID',
  `oper_name` VARCHAR(50) NOT NULL DEFAULT '' COMMENT '操作人员',
  `oper_url` VARCHAR(255) NOT NULL DEFAULT '' COMMENT '操作URL',
  `oper_method` VARCHAR(100) NOT NULL DEFAULT '' COMMENT '操作方法',
  `oper_ip` VARCHAR(50) NOT NULL DEFAULT '' COMMENT '操作IP',
  `oper_time` DATETIME NOT NULL DEFAULT CURRENT_TIMESTAMP COMMENT '操作时间',
  `created_at` DATETIME NOT NULL DEFAULT CURRENT_TIMESTAMP COMMENT '创建时间',
  PRIMARY KEY (`id`),
  INDEX `idx_org_id_name_time`(`org_id`, `oper_name`, `oper_time`),
  INDEX `idx_org_id_method_name`(`org_id`, `oper_method`, `oper_name`),
  INDEX `idx_org_id_time`(`org_id`, `oper_time`)
) ENGINE = InnoDB COMMENT = '操作日志记录';

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
