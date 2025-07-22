SET NAMES utf8mb4;
SET FOREIGN_KEY_CHECKS = 0;

CREATE DATABASE IF NOT EXISTS `identify_service` DEFAULT CHARACTER SET utf8mb4 COLLATE utf8mb4_unicode_ci;
USE identify_service;

-- ----------------------------
-- Table structure for casbin_rule
-- ----------------------------
DROP TABLE IF EXISTS `casbin_rule`;
CREATE TABLE `casbin_rule`  (
  `ptype` varchar(10) CHARACTER SET utf8mb4 COLLATE utf8mb4_general_ci NULL DEFAULT NULL,
  `v0` varchar(256) CHARACTER SET utf8mb4 COLLATE utf8mb4_general_ci NULL DEFAULT NULL,
  `v1` varchar(256) CHARACTER SET utf8mb4 COLLATE utf8mb4_general_ci NULL DEFAULT NULL,
  `v2` varchar(256) CHARACTER SET utf8mb4 COLLATE utf8mb4_general_ci NULL DEFAULT NULL,
  `v3` varchar(256) CHARACTER SET utf8mb4 COLLATE utf8mb4_general_ci NULL DEFAULT NULL,
  `v4` varchar(256) CHARACTER SET utf8mb4 COLLATE utf8mb4_general_ci NULL DEFAULT NULL,
  `v5` varchar(256) CHARACTER SET utf8mb4 COLLATE utf8mb4_general_ci NULL DEFAULT NULL
) ENGINE = InnoDB CHARACTER SET = utf8mb4 COLLATE = utf8mb4_general_ci ROW_FORMAT = COMPACT;

-- ----------------------------
-- Records of casbin_rule
-- ----------------------------
INSERT INTO `casbin_rule` VALUES ('g', 'u_1', '1', '', '', '', '');
INSERT INTO `casbin_rule` VALUES ('g', 'u_43', '1', '', '', '', '');
INSERT INTO `casbin_rule` VALUES ('p', '1', '27', 'All', '', '', '');
INSERT INTO `casbin_rule` VALUES ('p', '1', '28', 'All', '', '', '');
INSERT INTO `casbin_rule` VALUES ('p', '1', '29', 'All', '', '', '');
INSERT INTO `casbin_rule` VALUES ('p', '1', '30', 'All', '', '', '');
INSERT INTO `casbin_rule` VALUES ('p', '1', '1', 'All', '', '', '');
INSERT INTO `casbin_rule` VALUES ('p', '1', '2', 'All', '', '', '');
INSERT INTO `casbin_rule` VALUES ('p', '1', '3', 'All', '', '', '');
INSERT INTO `casbin_rule` VALUES ('p', '1', '4', 'All', '', '', '');
INSERT INTO `casbin_rule` VALUES ('p', '1', '11', 'All', '', '', '');
INSERT INTO `casbin_rule` VALUES ('p', '1', '10', 'All', '', '', '');
INSERT INTO `casbin_rule` VALUES ('p', '1', '12', 'All', '', '', '');
INSERT INTO `casbin_rule` VALUES ('p', '1', '13', 'All', '', '', '');
INSERT INTO `casbin_rule` VALUES ('p', '1', '14', 'All', '', '', '');
INSERT INTO `casbin_rule` VALUES ('p', '1', '15', 'All', '', '', '');
INSERT INTO `casbin_rule` VALUES ('p', '1', '19', 'All', '', '', '');
INSERT INTO `casbin_rule` VALUES ('p', '1', '20', 'All', '', '', '');
INSERT INTO `casbin_rule` VALUES ('p', '1', '21', 'All', '', '', '');
INSERT INTO `casbin_rule` VALUES ('p', '1', '22', 'All', '', '', '');
INSERT INTO `casbin_rule` VALUES ('p', '1', '23', 'All', '', '', '');
INSERT INTO `casbin_rule` VALUES ('p', '1', '24', 'All', '', '', '');
INSERT INTO `casbin_rule` VALUES ('p', '1', '25', 'All', '', '', '');
INSERT INTO `casbin_rule` VALUES ('p', '1', '26', 'All', '', '', '');
INSERT INTO `casbin_rule` VALUES ('p', '1', '31', 'All', '', '', '');
INSERT INTO `casbin_rule` VALUES ('p', '1', '32', 'All', '', '', '');
INSERT INTO `casbin_rule` VALUES ('p', '1', '34', 'All', '', '', '');
INSERT INTO `casbin_rule` VALUES ('p', '1', '38', 'All', '', '', '');
INSERT INTO `casbin_rule` VALUES ('p', '1', '39', 'All', '', '', '');
INSERT INTO `casbin_rule` VALUES ('p', '1', '35', 'All', '', '', '');
INSERT INTO `casbin_rule` VALUES ('p', '1', '33', 'All', '', '', '');
INSERT INTO `casbin_rule` VALUES ('p', '1', '36', 'All', '', '', '');
INSERT INTO `casbin_rule` VALUES ('p', '1', '37', 'All', '', '', '');
INSERT INTO `casbin_rule` VALUES ('p', '1', '53', 'All', '', '', '');
INSERT INTO `casbin_rule` VALUES ('p', '1', '54', 'All', '', '', '');
INSERT INTO `casbin_rule` VALUES ('p', '1', '55', 'All', '', '', '');
INSERT INTO `casbin_rule` VALUES ('p', '1', '56', 'All', '', '', '');
INSERT INTO `casbin_rule` VALUES ('p', '1', '57', 'All', '', '', '');
INSERT INTO `casbin_rule` VALUES ('p', '1', '58', 'All', '', '', '');
INSERT INTO `casbin_rule` VALUES ('p', '1', '17', 'All', '', '', '');
INSERT INTO `casbin_rule` VALUES ('p', '1', '16', 'All', '', '', '');
INSERT INTO `casbin_rule` VALUES ('p', '1', '18', 'All', '', '', '');
INSERT INTO `casbin_rule` VALUES ('p', '9', '27', 'All', '', '', '');
INSERT INTO `casbin_rule` VALUES ('p', '9', '28', 'All', '', '', '');
INSERT INTO `casbin_rule` VALUES ('p', '9', '29', 'All', '', '', '');
INSERT INTO `casbin_rule` VALUES ('p', '9', '30', 'All', '', '', '');
INSERT INTO `casbin_rule` VALUES ('p', '9', '1', 'All', '', '', '');
INSERT INTO `casbin_rule` VALUES ('p', '9', '2', 'All', '', '', '');
INSERT INTO `casbin_rule` VALUES ('p', '9', '3', 'All', '', '', '');
INSERT INTO `casbin_rule` VALUES ('p', '9', '4', 'All', '', '', '');
INSERT INTO `casbin_rule` VALUES ('p', '9', '11', 'All', '', '', '');
INSERT INTO `casbin_rule` VALUES ('p', '9', '10', 'All', '', '', '');
INSERT INTO `casbin_rule` VALUES ('p', '9', '12', 'All', '', '', '');
INSERT INTO `casbin_rule` VALUES ('p', '9', '13', 'All', '', '', '');
INSERT INTO `casbin_rule` VALUES ('p', '9', '14', 'All', '', '', '');
INSERT INTO `casbin_rule` VALUES ('p', '9', '15', 'All', '', '', '');
INSERT INTO `casbin_rule` VALUES ('p', '9', '19', 'All', '', '', '');
INSERT INTO `casbin_rule` VALUES ('p', '9', '20', 'All', '', '', '');
INSERT INTO `casbin_rule` VALUES ('p', '9', '21', 'All', '', '', '');
INSERT INTO `casbin_rule` VALUES ('p', '9', '22', 'All', '', '', '');
INSERT INTO `casbin_rule` VALUES ('p', '9', '23', 'All', '', '', '');
INSERT INTO `casbin_rule` VALUES ('p', '9', '24', 'All', '', '', '');
INSERT INTO `casbin_rule` VALUES ('p', '9', '25', 'All', '', '', '');
INSERT INTO `casbin_rule` VALUES ('p', '9', '26', 'All', '', '', '');
INSERT INTO `casbin_rule` VALUES ('g', 'u_5', '2', '', '', '', '');
INSERT INTO `casbin_rule` VALUES ('g', 'u_31', '2', '', '', '', '');
INSERT INTO `casbin_rule` VALUES ('g', 'u_6', '2', '', '', '', '');
INSERT INTO `casbin_rule` VALUES ('g', 'u_16', '2', '', '', '', '');
INSERT INTO `casbin_rule` VALUES ('g', 'u_3', '2', '', '', '', '');
INSERT INTO `casbin_rule` VALUES ('p', '2', '1', 'All', '', '', '');
INSERT INTO `casbin_rule` VALUES ('p', '2', '2', 'All', '', '', '');
INSERT INTO `casbin_rule` VALUES ('p', '2', '3', 'All', '', '', '');
INSERT INTO `casbin_rule` VALUES ('p', '2', '4', 'All', '', '', '');
INSERT INTO `casbin_rule` VALUES ('p', '2', '11', 'All', '', '', '');
INSERT INTO `casbin_rule` VALUES ('p', '2', '10', 'All', '', '', '');
INSERT INTO `casbin_rule` VALUES ('p', '2', '12', 'All', '', '', '');
INSERT INTO `casbin_rule` VALUES ('p', '2', '13', 'All', '', '', '');
INSERT INTO `casbin_rule` VALUES ('p', '2', '14', 'All', '', '', '');
INSERT INTO `casbin_rule` VALUES ('p', '2', '114', 'All', '', '', '');
INSERT INTO `casbin_rule` VALUES ('p', '2', '115', 'All', '', '', '');
INSERT INTO `casbin_rule` VALUES ('p', '2', '15', 'All', '', '', '');
INSERT INTO `casbin_rule` VALUES ('p', '2', '19', 'All', '', '', '');
INSERT INTO `casbin_rule` VALUES ('p', '2', '20', 'All', '', '', '');
INSERT INTO `casbin_rule` VALUES ('p', '2', '21', 'All', '', '', '');
INSERT INTO `casbin_rule` VALUES ('p', '2', '22', 'All', '', '', '');
INSERT INTO `casbin_rule` VALUES ('p', '2', '23', 'All', '', '', '');
INSERT INTO `casbin_rule` VALUES ('p', '2', '24', 'All', '', '', '');
INSERT INTO `casbin_rule` VALUES ('p', '2', '25', 'All', '', '', '');
INSERT INTO `casbin_rule` VALUES ('p', '2', '26', 'All', '', '', '');
INSERT INTO `casbin_rule` VALUES ('p', '2', '116', 'All', '', '', '');
INSERT INTO `casbin_rule` VALUES ('p', '2', '117', 'All', '', '', '');
INSERT INTO `casbin_rule` VALUES ('p', '2', '118', 'All', '', '', '');
INSERT INTO `casbin_rule` VALUES ('p', '2', '119', 'All', '', '', '');
INSERT INTO `casbin_rule` VALUES ('p', '2', '31', 'All', '', '', '');
INSERT INTO `casbin_rule` VALUES ('p', '2', '32', 'All', '', '', '');
INSERT INTO `casbin_rule` VALUES ('p', '2', '34', 'All', '', '', '');
INSERT INTO `casbin_rule` VALUES ('p', '2', '33', 'All', '', '', '');
INSERT INTO `casbin_rule` VALUES ('g', 'u_10', '2', '', '', '', '');


-- ----------------------------
-- Table structure for sys_auth_rule
-- ----------------------------
DROP TABLE IF EXISTS `sys_auth_rule`;
CREATE TABLE `sys_auth_rule`  (
  `id` int(10) UNSIGNED NOT NULL AUTO_INCREMENT,
  `pid` int(10) UNSIGNED NOT NULL DEFAULT 0 COMMENT '父ID',
  `name` varchar(100) CHARACTER SET utf8mb4 COLLATE utf8mb4_general_ci NOT NULL DEFAULT '' COMMENT '规则名称',
  `title` varchar(50) CHARACTER SET utf8mb4 COLLATE utf8mb4_general_ci NOT NULL DEFAULT '' COMMENT '规则名称',
  `icon` varchar(300) CHARACTER SET utf8mb4 COLLATE utf8mb4_general_ci NOT NULL DEFAULT '' COMMENT '图标',
  `condition` varchar(255) CHARACTER SET utf8mb4 COLLATE utf8mb4_general_ci NOT NULL DEFAULT '' COMMENT '条件',
  `remark` varchar(255) CHARACTER SET utf8mb4 COLLATE utf8mb4_general_ci NOT NULL DEFAULT '' COMMENT '备注',
  `menu_type` tinyint(1) UNSIGNED NOT NULL DEFAULT 0 COMMENT '类型 0目录 1菜单 2按钮',
  `weigh` int(10) NOT NULL DEFAULT 0 COMMENT '权重',
  `is_hide` tinyint(3) UNSIGNED NOT NULL DEFAULT 0 COMMENT '显示状态',
  `path` varchar(100) CHARACTER SET utf8mb4 COLLATE utf8mb4_general_ci NOT NULL DEFAULT '' COMMENT '路由地址',
  `component` varchar(100) CHARACTER SET utf8mb4 COLLATE utf8mb4_general_ci NOT NULL DEFAULT '' COMMENT '组件路径',
  `is_link` tinyint(3) UNSIGNED NOT NULL DEFAULT 0 COMMENT '是否外链 1是 0否',
  `module_type` varchar(30) CHARACTER SET utf8mb4 COLLATE utf8mb4_general_ci NOT NULL DEFAULT '' COMMENT '所属模块',
  `model_id` int(10) UNSIGNED NOT NULL DEFAULT 0 COMMENT '模型ID',
  `is_iframe` tinyint(3) UNSIGNED NOT NULL DEFAULT 0 COMMENT '是否内嵌iframe',
  `is_cached` tinyint(3) UNSIGNED NOT NULL DEFAULT 0 COMMENT '是否缓存',
  `redirect` varchar(255) CHARACTER SET utf8mb4 COLLATE utf8mb4_general_ci NOT NULL DEFAULT '' COMMENT '路由重定向地址',
  `is_affix` tinyint(3) UNSIGNED NOT NULL DEFAULT 0 COMMENT '是否固定',
  `link_url` varchar(500) CHARACTER SET utf8mb4 COLLATE utf8mb4_general_ci NOT NULL DEFAULT '' COMMENT '链接地址',
  `created_at` datetime NULL DEFAULT NULL COMMENT '创建日期',
  `updated_at` datetime NULL DEFAULT NULL COMMENT '修改日期',
  PRIMARY KEY (`id`) USING BTREE,
  UNIQUE INDEX `name`(`name`) USING BTREE,
  INDEX `pid`(`pid`) USING BTREE,
  INDEX `weigh`(`weigh`) USING BTREE
) ENGINE = InnoDB AUTO_INCREMENT = 152 CHARACTER SET = utf8mb4 COLLATE = utf8mb4_general_ci COMMENT = '菜单节点表' ROW_FORMAT = COMPACT;

-- ----------------------------
-- Records of sys_auth_rule
-- ----------------------------
INSERT INTO `sys_auth_rule` VALUES (1, 0, 'api/v1/system/auth', '权限管理', 'ele-Stamp', '', '', 0, 30, 0, '/system/auth', 'layout/routerView/parent', 0, '', 0, 0, 1, '0', 0, '', '2022-03-24 15:03:37', '2022-04-14 16:29:19');
INSERT INTO `sys_auth_rule` VALUES (2, 1, 'api/v1/system/auth/menuList', '菜单管理', 'ele-Calendar', '', '', 1, 0, 0, '/system/auth/menuList', 'system/menu/index', 0, '', 0, 0, 1, '', 0, '', '2022-03-24 17:24:13', '2022-03-29 10:54:49');
INSERT INTO `sys_auth_rule` VALUES (3, 2, 'api/v1/system/menu/add', '添加菜单', '', '', '', 2, 0, 0, '', '', 0, '', 0, 0, 1, '', 0, '', '2022-03-29 16:48:43', '2022-03-29 17:05:19');
INSERT INTO `sys_auth_rule` VALUES (4, 2, 'api/v1/system/menu/update', '修改菜单', '', '', '', 2, 0, 0, '', '', 0, '', 0, 0, 1, '', 0, '', '2022-03-29 17:04:25', '2022-03-29 18:11:36');
INSERT INTO `sys_auth_rule` VALUES (10, 1, 'api/v1/system/role/list', '角色管理', 'iconfont icon-juxingkaobei', '', '', 1, 0, 0, '/system/auth/roleList', 'system/role/index', 0, '', 0, 0, 1, '', 0, '', '2022-03-29 18:15:03', '2022-03-30 10:25:34');
INSERT INTO `sys_auth_rule` VALUES (11, 2, 'api/v1/system/menu/delete', '删除菜单', '', '', '', 2, 0, 0, '', '', 0, '', 0, 0, 1, '', 0, '', '2022-04-06 14:49:10', '2022-04-06 14:49:17');
INSERT INTO `sys_auth_rule` VALUES (12, 10, 'api/v1/system/role/add', '添加角色', '', '', '', 2, 0, 0, '', '', 0, '', 0, 0, 1, '', 0, '', '2022-04-06 14:49:46', '2022-04-06 14:49:46');
INSERT INTO `sys_auth_rule` VALUES (13, 10, '/api/v1/system/role/edit', '修改角色', '', '', '', 2, 0, 0, '', '', 0, '', 0, 0, 1, '', 0, '', '2022-04-06 14:50:08', '2022-04-06 14:50:08');
INSERT INTO `sys_auth_rule` VALUES (14, 10, '/api/v1/system/role/delete', '删除角色', '', '', '', 2, 0, 0, '', '', 0, '', 0, 0, 1, '', 0, '', '2022-04-06 14:50:22', '2022-04-06 14:50:22');
INSERT INTO `sys_auth_rule` VALUES (15, 1, 'api/v1/system/dept/list', '部门管理', 'iconfont icon-siweidaotu', '', '', 1, 0, 0, '/system/auth/deptList', 'system/dept/index', 0, '', 0, 0, 1, '', 0, '', '2022-04-06 14:52:23', '2022-04-07 22:59:20');
INSERT INTO `sys_auth_rule` VALUES (16, 17, 'aliyun', '阿里云-iframe', 'iconfont icon-diannao1', '', '', 1, 0, 0, '/demo/outLink/aliyun', 'layout/routerView/iframes', 1, '', 0, 1, 1, '', 0, 'https://www.aliyun.com/daily-act/ecs/activity_selection?spm=5176.8789780.J_3965641470.5.568845b58KHj51', '2022-04-06 17:26:29', '2022-04-07 15:27:17');
INSERT INTO `sys_auth_rule` VALUES (17, 0, 'outLink', '外链测试', 'iconfont icon-zhongduancanshu', '', '', 0, 20, 0, '/demo/outLink', 'layout/routerView/parent', 0, '', 0, 0, 1, '', 0, '', '2022-04-07 15:20:51', '2022-04-14 16:29:07');
INSERT INTO `sys_auth_rule` VALUES (18, 17, 'tenyun', '腾讯云-外链', 'iconfont icon-shouye_dongtaihui', '', '', 1, 0, 0, '/demo/outLink/tenyun', 'layout/routerView/link', 1, '', 0, 0, 1, '', 0, 'https://cloud.tencent.com/act/new?cps_key=20b1c3842f74986b2894e2c5fcde7ea2&fromSource=gwzcw.3775555.3775555.3775555&utm_id=gwzcw.3775555.3775555.3775555&utm_medium=cpc', '2022-04-07 15:23:52', '2023-09-28 10:27:41');
INSERT INTO `sys_auth_rule` VALUES (19, 15, 'api/v1/system/dept/add', '添加部门', '', '', '', 2, 0, 0, '', '', 0, '', 0, 0, 1, '', 0, '', '2022-04-07 22:56:39', '2022-04-07 22:56:39');
INSERT INTO `sys_auth_rule` VALUES (20, 15, 'api/v1/system/dept/edit', '修改部门', '', '', '', 2, 0, 0, '', '', 0, '', 0, 0, 1, '', 0, '', '2022-04-07 22:57:00', '2022-04-07 22:57:00');
INSERT INTO `sys_auth_rule` VALUES (21, 15, 'api/v1/system/dept/delete', '删除部门', '', '', '', 2, 0, 0, '', '', 0, '', 0, 0, 1, '', 0, '', '2022-04-07 22:57:30', '2022-04-07 22:57:30');
-- INSERT INTO `sys_auth_rule` VALUES (22, 1, 'api/v1/system/post/list', '岗位管理', 'iconfont icon-neiqianshujuchucun', '', '', 1, 0, 0, '/system/auth/postList', 'system/post/index', 0, '', 0, 0, 1, '', 0, '', '2022-04-07 22:58:46', '2022-04-09 14:26:15');
-- INSERT INTO `sys_auth_rule` VALUES (23, 22, 'api/v1/system/post/add', '添加岗位', '', '', '', 2, 0, 0, '', '', 0, '', 0, 0, 1, '', 0, '', '2022-04-09 14:14:49', '2022-04-09 14:14:49');
-- INSERT INTO `sys_auth_rule` VALUES (24, 22, 'api/v1/system/post/edit', '修改岗位', '', '', '', 2, 0, 0, '', '', 0, '', 0, 0, 1, '', 0, '', '2022-04-09 14:15:25', '2022-04-09 14:15:25');
-- INSERT INTO `sys_auth_rule` VALUES (25, 22, 'api/v1/system/post/delete', '删除岗位', '', '', '', 2, 0, 0, '', '', 0, '', 0, 0, 1, '', 0, '', '2022-04-09 14:15:47', '2022-04-09 14:15:47');
INSERT INTO `sys_auth_rule` VALUES (26, 1, 'api/v1/system/user/list', '用户管理', 'ele-User', '', '', 1, 0, 0, '/system/auth/user/list', 'system/user/index', 0, '', 0, 0, 1, '', 0, '', '2022-04-09 14:19:10', '2022-04-09 14:19:58');
INSERT INTO `sys_auth_rule` VALUES (27, 0, 'api/v1/system/dict', '系统配置', 'iconfont icon-shuxingtu', '', '', 0, 40, 0, '/system/dict', 'layout/routerView/parent', 0, '', 0, 0, 1, '654', 0, '', '2022-04-14 16:28:51', '2022-04-18 14:40:56');
INSERT INTO `sys_auth_rule` VALUES (28, 27, 'api/v1/system/dict/type/list', '字典管理', 'iconfont icon-crew_feature', '', '', 1, 0, 0, '/system/dict/type/list', 'system/dict/index', 0, '', 0, 0, 1, '', 0, '', '2022-04-14 16:32:10', '2022-04-16 17:02:50');
INSERT INTO `sys_auth_rule` VALUES (29, 27, 'api/v1/system/dict/dataList', '字典数据管理', 'iconfont icon-putong', '', '', 1, 0, 1, '/system/dict/data/list/:dictType', 'system/dict/dataList', 0, '', 0, 0, 1, '', 0, '', '2022-04-18 12:04:17', '2022-04-18 14:58:43');
INSERT INTO `sys_auth_rule` VALUES (30, 27, 'api/v1/system/config/list', '参数管理', 'ele-Cherry', '', '', 1, 0, 0, '/system/config/list', 'system/config/index', 0, '', 0, 0, 1, '', 0, '', '2022-04-18 21:05:20', '2022-04-18 21:13:19');
INSERT INTO `sys_auth_rule` VALUES (31, 0, 'api/v1/system/monitor', '系统监控', 'iconfont icon-xuanzeqi', '', '', 0, 30, 0, '/system/monitor', 'layout/routerView/parent', 0, '', 0, 0, 1, '', 0, '', '2022-04-19 10:40:19', '2022-04-19 10:44:38');
INSERT INTO `sys_auth_rule` VALUES (32, 31, 'api/v1/system/monitor/server', '服务监控', 'iconfont icon-shuju', '', '', 1, 0, 0, '/system/monitor/server', 'system/monitor/server/index', 0, '', 0, 0, 1, '', 0, '', '2022-04-19 10:43:32', '2022-04-19 10:44:47');
-- INSERT INTO `sys_auth_rule` VALUES (33, 35, 'api/swagger', 'api文档', 'iconfont icon--chaifenlie', '', '', 1, 0, 0, '/system/swagger', 'layout/routerView/iframes', 1, '', 0, 1, 1, '', 0, 'http://localhost:8808/swagger', '2022-04-21 09:23:43', '2022-11-29 17:10:35');
INSERT INTO `sys_auth_rule` VALUES (34, 31, 'api/v1/system/loginLog/list', '登录日志', 'ele-Finished', '', '', 1, 0, 0, '/system/monitor/loginLog', 'system/monitor/loginLog/index', 0, '', 0, 0, 1, '', 0, '', '2022-04-28 09:59:47', '2022-04-28 09:59:47');
INSERT INTO `sys_auth_rule` VALUES (35, 0, 'api/v1/system/tools', '系统工具', 'iconfont icon-zujian', '', '', 0, 25, 0, '/system/tools', 'layout/routerView/parent', 0, '', 0, 0, 1, '', 0, '', '2022-10-26 09:29:08', '2022-10-26 10:11:25');
INSERT INTO `sys_auth_rule` VALUES (36, 35, 'api/v1/system/tools/gen/tableList', '代码生成', 'iconfont icon-step', '', '', 1, 0, 0, '/system/tools/gen', 'system/tools/gen/index', 0, '', 0, 0, 1, '', 0, '', '2022-10-26 09:31:08', '2022-10-31 10:17:23');
INSERT INTO `sys_auth_rule` VALUES (37, 36, 'api/v1/system/tools/gen/columnList', '代码生成配置', 'ele-Edit', '', '', 1, 0, 1, '/system/tools/gen/edit', 'system/tools/gen/component/edit', 0, '', 0, 0, 1, '', 0, '', '2022-10-31 10:11:12', '2022-10-31 10:19:19');
INSERT INTO `sys_auth_rule` VALUES (38, 31, 'api/v1/system/operLog/list', '操作日志', 'iconfont icon-bolangnengshiyanchang', '', '', 1, 0, 0, '/system/monitor/operLog', 'system/monitor/operLog/index', 0, '', 0, 0, 1, '', 0, '', '2022-12-23 16:19:05', '2022-12-23 16:21:50');
INSERT INTO `sys_auth_rule` VALUES (39, 31, 'api/v1/system/online/list', '在线用户', 'iconfont icon-skin', '', '', 1, 0, 0, '/system/monitor/userOnlie', 'system/monitor/userOnline/index', 0, '', 0, 0, 1, '', 0, '', '2023-01-11 15:48:06', '2023-01-11 17:02:39');
INSERT INTO `sys_auth_rule` VALUES (53, 35, 'api/v1/system/sysJob/list', '定时任务', 'fa fa-superpowers', '', '', 1, 0, 0, '/system/sysJob/list', 'system/sysJob/list/index', 0, 'sys_admin', 0, 0, 1, '', 0, '', NULL, '2023-01-12 17:51:27');
INSERT INTO `sys_auth_rule` VALUES (54, 53, 'api/v1/system/sysJob/get', '定时任务查询', '', '', '定时任务查询', 2, 0, 0, '', '', 0, 'sys_admin', 0, 0, 1, '', 0, '', NULL, NULL);
INSERT INTO `sys_auth_rule` VALUES (55, 53, 'api/v1/system/sysJob/add', '定时任务添加', '', '', '定时任务添加', 2, 0, 0, '', '', 0, 'sys_admin', 0, 0, 1, '', 0, '', NULL, NULL);
INSERT INTO `sys_auth_rule` VALUES (56, 53, 'api/v1/system/sysJob/edit', '定时任务修改', '', '', '定时任务修改', 2, 0, 0, '', '', 0, 'sys_admin', 0, 0, 1, '', 0, '', NULL, NULL);
INSERT INTO `sys_auth_rule` VALUES (57, 53, 'api/v1/system/sysJob/delete', '定时任务删除', '', '', '定时任务删除', 2, 0, 0, '', '', 0, 'sys_admin', 0, 0, 1, '', 0, '', NULL, NULL);
INSERT INTO `sys_auth_rule` VALUES (58, 53, 'api/v1/system/sysJob/run', '执行一次', '', '', '', 2, 0, 0, '', '', 0, '', 0, 0, 1, '', 0, '', '2023-01-12 18:20:13', '2023-01-12 18:20:13');
INSERT INTO `sys_auth_rule` VALUES (59, 0, 'api/v1/system/sysNotice', '通知公告', 'iconfont icon-fuwenbenkuang', '', '', 0, 0, 0, '/system/sysNotice', 'layout/routerView/parent', 0, 'sys_admin', 0, 0, 1, '', 0, '', NULL, '2023-11-09 15:40:55');
INSERT INTO `sys_auth_rule` VALUES (60, 59, 'api/v1/system/sysNotice/list', '通知公告管理', 'ele-Fold', '', '', 1, 0, 0, '/system/sysNotice/list', 'system/sysNotice/list/index', 0, 'sys_admin', 0, 0, 1, '', 0, '', NULL, '2023-11-09 15:41:13');
INSERT INTO `sys_auth_rule` VALUES (61, 60, 'api/v1/system/sysNotice/get', '通知公告查询', '', '', '通知公告查询', 2, 0, 0, '', '', 0, 'sys_admin', 0, 0, 1, '', 0, '', NULL, NULL);
INSERT INTO `sys_auth_rule` VALUES (62, 60, 'api/v1/system/sysNotice/add', '通知公告添加', '', '', '通知公告添加', 2, 0, 0, '', '', 0, 'sys_admin', 0, 0, 1, '', 0, '', NULL, NULL);
INSERT INTO `sys_auth_rule` VALUES (63, 60, 'api/v1/system/sysNotice/edit', '通知公告修改', '', '', '通知公告修改', 2, 0, 0, '', '', 0, 'sys_admin', 0, 0, 1, '', 0, '', NULL, NULL);
INSERT INTO `sys_auth_rule` VALUES (64, 60, 'api/v1/system/sysNotice/delete', '通知公告删除', '', '', '通知公告删除', 2, 0, 0, '', '', 0, 'sys_admin', 0, 0, 1, '', 0, '', NULL, NULL);
INSERT INTO `sys_auth_rule` VALUES (65, 59, 'api/v1/system/sysNotice/show', '通知公告展示', 'iconfont icon-tongzhi', '', '', 0, 0, 0, '/system/sysNotice/show', 'system/sysNotice/show/index', 0, '', 0, 0, 1, '', 0, '', '2023-12-25 10:34:32', '2024-01-03 10:09:12');
INSERT INTO `sys_auth_rule` VALUES (114, 10, 'api/v1/system/role/dataScope', '授权数据权限', '', '', '', 2, 0, 0, '', '', 0, '', 0, 0, 1, '', 0, '', '2024-04-11 11:18:23', '2024-04-11 11:19:17');
INSERT INTO `sys_auth_rule` VALUES (115, 10, 'api/v1/system/role/setRoleUser', '授权用户权限', '', '', '', 2, 0, 0, '', '', 0, '', 0, 0, 1, '', 0, '', '2024-04-11 11:18:59', '2024-04-11 11:18:59');
INSERT INTO `sys_auth_rule` VALUES (116, 26, 'api/v1/system/user/add', '新增用户', '', '', '', 2, 0, 0, '', '', 0, '', 0, 0, 1, '', 0, '', '2024-04-11 11:16:14', '2024-04-11 11:16:14');
INSERT INTO `sys_auth_rule` VALUES (117, 26, 'api/v1/system/user/edit', '修改用户', '', '', '', 2, 0, 0, '', '', 0, '', 0, 0, 1, '', 0, '', '2024-04-11 11:16:45', '2024-04-11 11:19:25');
INSERT INTO `sys_auth_rule` VALUES (118, 26, 'api/v1/system/user/getEdit', '用户信息', '', '', '', 2, 0, 0, '', '', 0, '', 0, 0, 1, '', 0, '', '2024-04-11 11:17:21', '2024-04-11 11:19:30');
INSERT INTO `sys_auth_rule` VALUES (119, 26, 'api/v1/system/user/delete', '删除用户', '', '', '', 2, 0, 0, '', '', 0, '', 0, 0, 1, '', 0, '', '2024-04-11 11:17:39', '2024-04-11 11:19:34');
INSERT INTO `sys_auth_rule` VALUES (141, 26, 'api/v1/system/user/all', '管理所有', '', '', '', 2, 0, 0, '', '', 0, '', 0, 0, 1, '', 0, '', '2024-09-13 16:57:13', '2024-09-13 16:57:13');
INSERT INTO `sys_auth_rule` VALUES (142, 15, 'api/v1/system/dept/all', '管理所有', '', '', '', 2, 0, 0, '', '', 0, '', 0, 0, 1, '', 0, '', '2024-09-13 16:57:49', '2024-09-13 16:57:49');
INSERT INTO `sys_auth_rule` VALUES (143, 35, 'api/v1/system/sysAttachment/list', '附件管理', 'ele-Folder', '', '', 1, 0, 0, '/system/sysAttachment/list', 'system/sysAttachment/list/index', 0, 'sys_admin', 0, 0, 1, '', 0, '', NULL, '2024-11-04 14:39:02');
INSERT INTO `sys_auth_rule` VALUES (144, 143, 'api/v1/system/sysAttachment/delete', '附件管理删除', '', '', '附件管理删除', 2, 0, 0, '', '', 0, 'sys_admin', 0, 0, 1, '', 0, '', NULL, NULL);
-- INSERT INTO `sys_auth_rule` VALUES (145, 0, 'api/v1/demo/demoSnowId', '雪花ID测试管理', 'iconfont icon-fuwenbenkuang', '', '雪花ID测试管理', 0, 0, 0, '/demo/demoSnowId', 'layout/routerView/parent', 0, 'sys_admin', 0, 0, 1, '', 0, '', NULL, NULL);
-- INSERT INTO `sys_auth_rule` VALUES (146, 145, 'api/v1/demo/demoSnowId/list', '雪花ID测试列表', 'ele-Fold', '', '雪花ID测试列表', 1, 0, 0, '/demo/demoSnowId/list', 'demo/demoSnowId/list/index', 0, 'sys_admin', 0, 0, 1, '', 0, '', NULL, NULL);
-- INSERT INTO `sys_auth_rule` VALUES (147, 146, 'api/v1/demo/demoSnowId/get', '雪花ID测试查询', '', '', '雪花ID测试查询', 2, 0, 0, '', '', 0, 'sys_admin', 0, 0, 1, '', 0, '', NULL, NULL);
-- INSERT INTO `sys_auth_rule` VALUES (148, 146, 'api/v1/demo/demoSnowId/add', '雪花ID测试添加', '', '', '雪花ID测试添加', 2, 0, 0, '', '', 0, 'sys_admin', 0, 0, 1, '', 0, '', NULL, NULL);
-- INSERT INTO `sys_auth_rule` VALUES (149, 146, 'api/v1/demo/demoSnowId/edit', '雪花ID测试修改', '', '', '雪花ID测试修改', 2, 0, 0, '', '', 0, 'sys_admin', 0, 0, 1, '', 0, '', NULL, NULL);
-- INSERT INTO `sys_auth_rule` VALUES (150, 146, 'api/v1/demo/demoSnowId/delete', '雪花ID测试删除', '', '', '雪花ID测试删除', 2, 0, 0, '', '', 0, 'sys_admin', 0, 0, 1, '', 0, '', NULL, NULL);
-- INSERT INTO `sys_auth_rule` VALUES (151, 146, 'api/v1/demo/demoSnowId/export', '雪花ID测试导出', '', '', '雪花ID测试导出', 2, 0, 0, '', '', 0, 'sys_admin', 0, 0, 1, '', 0, '', NULL, NULL);
INSERT INTO `sys_auth_rule` VALUES (152,0,'api/v1/system/tStore','门店管理','iconfont icon-fuwenbenkuang','','门店管理',0,0,0,'/system/tStore','layout/routerView/parent',0,'sys_admin',0,0,1,'',0,'',NULL,NULL);
INSERT INTO `sys_auth_rule` VALUES (153,152,'api/v1/system/tStore/list','门店列表','ele-Fold','','门店列表',1,0,0,'/system/tStore/list','system/tStore/list/index',0,'sys_admin',0,0,1,'',0,'',NULL,NULL);
INSERT INTO `sys_auth_rule` VALUES (154,153,'api/v1/system/tStore/get','门店查询','','','门店查询',2,0,0,'','',0,'sys_admin',0,0,1,'',0,'',NULL,NULL);
INSERT INTO `sys_auth_rule` VALUES (155,153,'api/v1/system/tStore/add','门店添加','','','门店添加',2,0,0,'','',0,'sys_admin',0,0,1,'',0,'',NULL,NULL);
INSERT INTO `sys_auth_rule` VALUES (156,153,'api/v1/system/tStore/edit','门店修改','','','门店修改',2,0,0,'','',0,'sys_admin',0,0,1,'',0,'',NULL,NULL);
INSERT INTO `sys_auth_rule` VALUES (157,153,'api/v1/system/tStore/delete','门店删除','','','门店删除',2,0,0,'','',0,'sys_admin',0,0,1,'',0,'',NULL,NULL);

-- ----------------------------
-- Table structure for sys_user
-- ----------------------------
DROP TABLE IF EXISTS `sys_user`;
CREATE TABLE `sys_user`  (
  `id` bigint(20) UNSIGNED NOT NULL AUTO_INCREMENT,
  `user_name` varchar(60) CHARACTER SET utf8mb4 COLLATE utf8mb4_unicode_ci NOT NULL DEFAULT '' COMMENT '用户名',
  `mobile` varchar(20) CHARACTER SET utf8mb4 COLLATE utf8mb4_general_ci NOT NULL DEFAULT '' COMMENT '中国手机不带国家代码，国际手机号格式为：国家代码-手机号',
  `user_nickname` varchar(50) CHARACTER SET utf8mb4 COLLATE utf8mb4_unicode_ci NOT NULL DEFAULT '' COMMENT '用户昵称',
  `birthday` int(11) NOT NULL DEFAULT 0 COMMENT '生日',
  `user_password` varchar(255) CHARACTER SET utf8mb4 COLLATE utf8mb4_general_ci NOT NULL DEFAULT '' COMMENT '登录密码;cmf_password加密',
  `user_salt` char(10) CHARACTER SET utf8mb4 COLLATE utf8mb4_general_ci NOT NULL COMMENT '加密盐',
  `user_status` tinyint(3) UNSIGNED NOT NULL DEFAULT 1 COMMENT '用户状态;0:禁用,1:正常,2:未验证',
  `user_email` varchar(100) CHARACTER SET utf8mb4 COLLATE utf8mb4_general_ci NOT NULL DEFAULT '' COMMENT '用户登录邮箱',
  `sex` tinyint(2) NOT NULL DEFAULT 0 COMMENT '性别;0:保密,1:男,2:女',
  `avatar` varchar(255) CHARACTER SET utf8mb4 COLLATE utf8mb4_general_ci NOT NULL DEFAULT '' COMMENT '用户头像',
  `dept_id` bigint(20) UNSIGNED NOT NULL DEFAULT 0 COMMENT '部门id',
  `remark` varchar(255) CHARACTER SET utf8mb4 COLLATE utf8mb4_general_ci NOT NULL DEFAULT '' COMMENT '备注',
  `is_admin` tinyint(4) NOT NULL DEFAULT 1 COMMENT '是否后台管理员 1 是  0   否',
  `address` varchar(255) CHARACTER SET utf8mb4 COLLATE utf8mb4_general_ci NOT NULL DEFAULT '' COMMENT '联系地址',
  `describe` varchar(255) CHARACTER SET utf8mb4 COLLATE utf8mb4_general_ci NOT NULL DEFAULT '' COMMENT ' 描述信息',
  `last_login_ip` varchar(15) CHARACTER SET utf8mb4 COLLATE utf8mb4_general_ci NOT NULL DEFAULT '' COMMENT '最后登录ip',
  `last_login_time` datetime NULL DEFAULT NULL COMMENT '最后登录时间',
  `created_at` datetime NULL DEFAULT NULL COMMENT '创建时间',
  `updated_at` datetime NULL DEFAULT NULL COMMENT '更新时间',
  `deleted_at` datetime NULL DEFAULT NULL COMMENT '删除时间',
  `open_id` varchar(50) CHARACTER SET utf8mb4 COLLATE utf8mb4_general_ci NOT NULL DEFAULT '' COMMENT '微信open id',
  PRIMARY KEY (`id`) USING BTREE,
  UNIQUE INDEX `user_login`(`user_name`, `deleted_at`) USING BTREE,
  UNIQUE INDEX `mobile`(`mobile`, `deleted_at`) USING BTREE,
  INDEX `user_nickname`(`user_nickname`) USING BTREE,
  INDEX `open_id`(`open_id`) USING BTREE
) ENGINE = InnoDB AUTO_INCREMENT = 100000 CHARACTER SET = utf8mb4 COLLATE = utf8mb4_general_ci COMMENT = '用户表' ROW_FORMAT = COMPACT;

-- ----------------------------
-- Records of sys_user
-- ----------------------------
INSERT INTO `sys_user` VALUES (1, 'admin', '13578342363', '超级管理员', 0, 'c567ae329f9929b518759d3bea13f492', 'f9aZTAa8yz', 1, 'yxh669@qq.com', 1, 'https://yxh-1301841944.cos.ap-chongqing.myqcloud.com/gfast/2021-07-19/ccwpeuqz1i2s769hua.jpeg', 101, '', 1, 'asdasfdsaf大发放打发士大夫发按时', '描述信息', '::1', '2023-10-31 11:22:06', '2021-06-22 17:58:00', '2023-04-22 14:39:18', NULL, '');

-- ----------------------------
-- Table structure for sys_role
-- ----------------------------
DROP TABLE IF EXISTS `sys_role`;
CREATE TABLE `sys_role`  (
  `id` int(10) UNSIGNED NOT NULL AUTO_INCREMENT,
  `pid` int(10) UNSIGNED NOT NULL DEFAULT 0 COMMENT '父级ID',
  `status` tinyint(3) UNSIGNED NOT NULL DEFAULT 0 COMMENT '状态;0:禁用;1:正常',
  `list_order` int(10) UNSIGNED NOT NULL DEFAULT 0 COMMENT '排序',
  `name` varchar(20) CHARACTER SET utf8mb4 COLLATE utf8mb4_unicode_ci NOT NULL DEFAULT '' COMMENT '角色名称',
  `remark` varchar(255) CHARACTER SET utf8mb4 COLLATE utf8mb4_general_ci NOT NULL DEFAULT '' COMMENT '备注',
  `data_scope` tinyint(3) UNSIGNED NOT NULL DEFAULT 3 COMMENT '数据范围（1：全部数据权限 2：自定数据权限 3：本部门数据权限 4：本部门及以下数据权限）',
  `created_at` datetime NULL DEFAULT NULL COMMENT '创建时间',
  `updated_at` datetime NULL DEFAULT NULL COMMENT '更新时间',
  `created_by` bigint(20) UNSIGNED NOT NULL DEFAULT 0 COMMENT '添加人',
  `effectiveTime` text CHARACTER SET utf8mb4 COLLATE utf8mb4_general_ci NULL COMMENT '角色有效日期',
  PRIMARY KEY (`id`) USING BTREE,
  INDEX `status`(`status`) USING BTREE
) ENGINE = InnoDB AUTO_INCREMENT = 10 CHARACTER SET = utf8mb4 COLLATE = utf8mb4_general_ci COMMENT = '角色表' ROW_FORMAT = COMPACT;

-- ----------------------------
-- Records of sys_role
-- ----------------------------
INSERT INTO `sys_role` VALUES (1, 0, 1, 0, '超级管理员', '备注', 2, '2022-04-01 11:38:39', '2023-09-28 10:27:55', 0, NULL);
INSERT INTO `sys_role` VALUES (2, 1, 1, 0, '普通管理员', '备注', 5, '2022-04-01 11:38:39', '2024-09-14 09:10:55', 0, '{\"effectiveType\":0,\"weekDay\":[1,2,3,4,5],\"dayRange\":[\"2024-04-12 08:00:00\",\"2024-04-12 18:00:00\"],\"dateRange\":null}');
INSERT INTO `sys_role` VALUES (3, 0, 1, 0, '站点管理员', '站点管理人员', 3, '2022-04-01 11:38:39', '2022-04-01 11:38:39', 0, NULL);
INSERT INTO `sys_role` VALUES (4, 5, 1, 0, '初级管理员', '初级管理员', 3, '2022-04-01 11:38:39', '2024-03-18 10:16:15', 0, '{\"effectiveType\":0,\"weekDay\":null,\"dayRange\":null,\"dateRange\":null}');
INSERT INTO `sys_role` VALUES (5, 0, 1, 0, '高级管理员', '高级管理员', 2, '2022-04-01 11:38:39', '2022-04-01 11:38:39', 0, NULL);
INSERT INTO `sys_role` VALUES (8, 0, 1, 0, '区级管理员', '', 2, '2022-04-01 11:38:39', '2022-04-06 09:53:40', 0, NULL);
INSERT INTO `sys_role` VALUES (9, 0, 1, 0, '测试', '', 3, '2023-04-22 12:39:13', '2023-09-28 15:48:56', 3, NULL);


-- ----------------------------
-- Table structure for sys_job
-- ----------------------------
DROP TABLE IF EXISTS `sys_job`;
CREATE TABLE `sys_job`  (
  `job_id` bigint(20) UNSIGNED NOT NULL AUTO_INCREMENT COMMENT '任务ID',
  `job_name` varchar(64) CHARACTER SET utf8mb4 COLLATE utf8mb4_general_ci NOT NULL DEFAULT '' COMMENT '任务名称',
  `job_params` varchar(200) CHARACTER SET utf8mb4 COLLATE utf8mb4_general_ci NULL DEFAULT '' COMMENT '参数',
  `job_group` varchar(64) CHARACTER SET utf8mb4 COLLATE utf8mb4_general_ci NOT NULL DEFAULT 'DEFAULT' COMMENT '任务组名',
  `invoke_target` varchar(100) CHARACTER SET utf8mb4 COLLATE utf8mb4_general_ci NOT NULL COMMENT '调用目标字符串',
  `cron_expression` varchar(255) CHARACTER SET utf8mb4 COLLATE utf8mb4_general_ci NULL DEFAULT '' COMMENT 'cron执行表达式',
  `misfire_policy` tinyint(4) NULL DEFAULT 1 COMMENT '计划执行策略（1多次执行 2执行一次）',
  `concurrent` tinyint(4) NULL DEFAULT 1 COMMENT '是否并发执行（0允许 1禁止）',
  `status` tinyint(4) NULL DEFAULT 0 COMMENT '状态（0正常 1暂停）',
  `created_by` bigint(64) UNSIGNED NULL DEFAULT 0 COMMENT '创建者',
  `updated_by` bigint(64) UNSIGNED NULL DEFAULT 0 COMMENT '更新者',
  `remark` varchar(500) CHARACTER SET utf8mb4 COLLATE utf8mb4_general_ci NULL DEFAULT '' COMMENT '备注信息',
  `created_at` datetime NULL DEFAULT NULL COMMENT '创建时间',
  `updated_at` datetime NULL DEFAULT NULL COMMENT '更新时间',
  PRIMARY KEY (`job_id`) USING BTREE,
  UNIQUE INDEX `invoke_target`(`invoke_target`) USING BTREE
) ENGINE = InnoDB AUTO_INCREMENT = 9 CHARACTER SET = utf8mb4 COLLATE = utf8mb4_general_ci COMMENT = '定时任务调度表' ROW_FORMAT = COMPACT;


-- ----------------------------
-- Table structure for sys_attachment
-- ----------------------------
DROP TABLE IF EXISTS `sys_attachment`;
CREATE TABLE `sys_attachment`  (
  `id` bigint(20) NOT NULL AUTO_INCREMENT COMMENT '文件ID',
  `app_id` varchar(64) CHARACTER SET utf8mb4 COLLATE utf8mb4_general_ci NOT NULL COMMENT '应用ID',
  `drive` varchar(64) CHARACTER SET utf8mb4 COLLATE utf8mb4_general_ci NULL DEFAULT NULL COMMENT '上传驱动',
  `name` varchar(1000) CHARACTER SET utf8mb4 COLLATE utf8mb4_general_ci NULL DEFAULT NULL COMMENT '文件原始名',
  `kind` varchar(16) CHARACTER SET utf8mb4 COLLATE utf8mb4_general_ci NULL DEFAULT NULL COMMENT '上传类型',
  `mime_type` varchar(128) CHARACTER SET utf8mb4 COLLATE utf8mb4_general_ci NOT NULL DEFAULT '' COMMENT '扩展类型',
  `path` varchar(1000) CHARACTER SET utf8mb4 COLLATE utf8mb4_general_ci NULL DEFAULT NULL COMMENT '本地路径',
  `size` bigint(20) NULL DEFAULT 0 COMMENT '文件大小',
  `ext` varchar(50) CHARACTER SET utf8mb4 COLLATE utf8mb4_general_ci NULL DEFAULT NULL COMMENT '扩展名',
  `md5` varchar(32) CHARACTER SET utf8mb4 COLLATE utf8mb4_general_ci NULL DEFAULT NULL COMMENT 'md5校验码',
  `created_by` bigint(20) NULL DEFAULT 0 COMMENT '上传人ID',
  `status` tinyint(1) NOT NULL DEFAULT 1 COMMENT '状态',
  `created_at` datetime NULL DEFAULT NULL COMMENT '创建时间',
  `updated_at` datetime NULL DEFAULT NULL COMMENT '修改时间',
  PRIMARY KEY (`id`) USING BTREE,
  INDEX `md5`(`md5`) USING BTREE
) ENGINE = InnoDB AUTO_INCREMENT = 4 CHARACTER SET = utf8mb4 COLLATE = utf8mb4_general_ci COMMENT = '附件管理' ROW_FORMAT = DYNAMIC;

-- ----------------------------
-- Records of sys_attachment
-- ----------------------------
INSERT INTO `sys_attachment` VALUES (1, 'system', '0', 'a333.jpg', 'image', 'image/jpeg', 'upload_file/2024-11-18/d5p41ldap0bwdvsdee.jpg', 18381, 'jpg', 'b1f4daf376347c18666c4670325cebd7', 31, 1, '2024-11-18 15:02:30', '2024-11-18 15:02:30');
INSERT INTO `sys_attachment` VALUES (2, 'system', '0', 'a222.jpg', 'image', 'image/jpeg', 'upload_file/2024-11-18/d5p41p66q7fctdzicz.jpg', 25947, 'jpg', 'f4378e34be8f7a6c311ad068b3498405', 31, 1, '2024-11-18 15:02:39', '2024-11-18 15:02:39');
INSERT INTO `sys_attachment` VALUES (3, 'system', '0', 'a111.jpg', 'image', 'application/octet-stream', 'upload_file/2024-11-18/d5p41ssy6gg4beecej.jpg', 104079, 'jpg', '11c14ce4809ff610e44b1da373500703', 31, 1, '2024-11-18 15:02:47', '2024-11-18 15:02:47');


-- ----------------------------
-- Table structure for sys_config
-- ----------------------------
DROP TABLE IF EXISTS `sys_config`;
CREATE TABLE `sys_config`  (
  `config_id` int(5) UNSIGNED NOT NULL AUTO_INCREMENT COMMENT '参数主键',
  `config_name` varchar(100) CHARACTER SET utf8mb4 COLLATE utf8mb4_general_ci NULL DEFAULT '' COMMENT '参数名称',
  `config_key` varchar(100) CHARACTER SET utf8mb4 COLLATE utf8mb4_general_ci NULL DEFAULT '' COMMENT '参数键名',
  `config_value` varchar(500) CHARACTER SET utf8mb4 COLLATE utf8mb4_general_ci NULL DEFAULT '' COMMENT '参数键值',
  `config_type` tinyint(1) NULL DEFAULT 0 COMMENT '系统内置（Y是 N否）',
  `create_by` int(64) UNSIGNED NULL DEFAULT 0 COMMENT '创建者',
  `update_by` int(64) UNSIGNED NULL DEFAULT 0 COMMENT '更新者',
  `remark` varchar(500) CHARACTER SET utf8mb4 COLLATE utf8mb4_general_ci NULL DEFAULT NULL COMMENT '备注',
  `created_at` datetime NULL DEFAULT NULL COMMENT '创建时间',
  `updated_at` datetime NULL DEFAULT NULL COMMENT '修改时间',
  PRIMARY KEY (`config_id`) USING BTREE,
  UNIQUE INDEX `uni_config_key`(`config_key`) USING BTREE
) ENGINE = InnoDB AUTO_INCREMENT = 12 CHARACTER SET = utf8mb4 COLLATE = utf8mb4_general_ci ROW_FORMAT = COMPACT;

-- ----------------------------
-- Records of sys_config
-- ----------------------------
INSERT INTO `sys_config` VALUES (1, '文件上传-文件大小', 'sys.uploadFile.fileSize', '50M', 1, 31, 31, '文件上传大小限制', NULL, '2021-07-06 14:57:35');
INSERT INTO `sys_config` VALUES (2, '文件上传-文件类型', 'sys.uploadFile.fileType', 'doc,docx,zip,xls,xlsx,rar,jpg,jpeg,gif,npm,png,mp4', 1, 31, 31, '文件上传后缀类型限制', NULL, '2022-12-16 09:52:45');
INSERT INTO `sys_config` VALUES (3, '图片上传-图片类型', 'sys.uploadFile.imageType', 'jpg,jpeg,gif,npm,png', 1, 31, 0, '图片上传后缀类型限制', NULL, NULL);
INSERT INTO `sys_config` VALUES (4, '图片上传-图片大小', 'sys.uploadFile.imageSize', '50M', 1, 31, 31, '图片上传大小限制', NULL, NULL);
INSERT INTO `sys_config` VALUES (11, '静态资源', 'static.resource', '/', 1, 2, 0, '', NULL, NULL);

-- ----------------------------
-- Table structure for sys_dept
-- ----------------------------
DROP TABLE IF EXISTS `sys_dept`;
CREATE TABLE `sys_dept`  (
  `dept_id` bigint(20) NOT NULL AUTO_INCREMENT COMMENT '部门id',
  `parent_id` bigint(20) NULL DEFAULT 0 COMMENT '父部门id',
  `ancestors` varchar(50) CHARACTER SET utf8mb4 COLLATE utf8mb4_general_ci NULL DEFAULT '' COMMENT '祖级列表',
  `dept_name` varchar(30) CHARACTER SET utf8mb4 COLLATE utf8mb4_general_ci NULL DEFAULT '' COMMENT '部门名称',
  `order_num` int(4) NULL DEFAULT 0 COMMENT '显示顺序',
  `leader` varchar(20) CHARACTER SET utf8mb4 COLLATE utf8mb4_general_ci NULL DEFAULT NULL COMMENT '负责人',
  `phone` varchar(11) CHARACTER SET utf8mb4 COLLATE utf8mb4_general_ci NULL DEFAULT NULL COMMENT '联系电话',
  `email` varchar(50) CHARACTER SET utf8mb4 COLLATE utf8mb4_general_ci NULL DEFAULT NULL COMMENT '邮箱',
  `status` tinyint(3) UNSIGNED NULL DEFAULT 0 COMMENT '部门状态（0正常 1停用）',
  `created_by` bigint(20) UNSIGNED NULL DEFAULT 0 COMMENT '创建人',
  `updated_by` bigint(20) NULL DEFAULT NULL COMMENT '修改人',
  `created_at` datetime NULL DEFAULT NULL COMMENT '创建时间',
  `updated_at` datetime NULL DEFAULT NULL COMMENT '修改时间',
  `deleted_at` datetime NULL DEFAULT NULL COMMENT '删除时间',
  PRIMARY KEY (`dept_id`) USING BTREE
) ENGINE = InnoDB AUTO_INCREMENT = 204 CHARACTER SET = utf8mb4 COLLATE = utf8mb4_general_ci COMMENT = '部门表' ROW_FORMAT = COMPACT;

-- ----------------------------
-- Records of sys_dept
-- ----------------------------
INSERT INTO `sys_dept` VALUES (100, 0, '0', '奇讯科技', 0, '[1,2,3]', '15888888888', 'ry@qq.com', 1, 0, 31, '2021-07-13 15:56:52', '2024-01-29 16:00:28', NULL);
INSERT INTO `sys_dept` VALUES (101, 100, '0,100', '深圳总公司', 1, NULL, '15888888888', 'ry@qq.com', 1, 0, NULL, '2021-07-13 15:56:52', '2021-07-13 15:56:52', NULL);
INSERT INTO `sys_dept` VALUES (102, 100, '0,100', '长沙分公司', 2, NULL, '15888888888', 'ry@qq.com', 1, 0, NULL, '2021-07-13 15:56:52', '2021-07-13 15:56:52', NULL);
INSERT INTO `sys_dept` VALUES (103, 101, '0,100,101', '研发部门', 1, NULL, '15888888888', 'ry@qq.com', 1, 0, NULL, '2021-07-13 15:56:52', '2021-07-13 15:56:52', NULL);
INSERT INTO `sys_dept` VALUES (104, 101, '0,100,101', '市场部门', 2, NULL, '15888888888', 'ry@qq.com', 1, 0, 31, '2021-07-13 15:56:52', '2021-11-04 09:16:38', NULL);
INSERT INTO `sys_dept` VALUES (105, 101, '0,100,101', '测试部门', 3, NULL, '15888888888', 'ry@qq.com', 1, 0, NULL, '2021-07-13 15:56:52', '2021-07-13 15:56:52', NULL);
INSERT INTO `sys_dept` VALUES (106, 101, '0,100,101', '财务部门', 4, NULL, '15888888888', 'ry@qq.com', 1, 0, NULL, '2021-07-13 15:56:52', '2021-07-13 15:56:52', NULL);
INSERT INTO `sys_dept` VALUES (107, 101, '0,100,101', '运维部门', 5, NULL, '15888888888', 'ry@qq.com', 1, 0, NULL, '2021-07-13 15:56:52', '2021-07-13 15:56:52', NULL);
INSERT INTO `sys_dept` VALUES (108, 102, '0,100,102', '市场部门', 1, NULL, '15888888888', 'ry@qq.com', 1, 0, NULL, '2021-07-13 15:56:52', '2021-07-13 15:56:52', NULL);
INSERT INTO `sys_dept` VALUES (109, 102, '0,100,102', '财务部门', 2, NULL, '15888888888', 'ry@qq.com', 1, 0, NULL, '2021-07-13 15:56:52', '2021-07-13 15:56:52', NULL);
INSERT INTO `sys_dept` VALUES (200, 100, '', '大数据', 1, '', '18888888888', 'liou@qq.com', 0, 0, 31, '2021-07-13 15:56:52', '2022-09-16 16:46:57', NULL);
INSERT INTO `sys_dept` VALUES (201, 100, '', '开发', 1, NULL, '18888888888', 'li@qq.com', 0, 31, NULL, '2021-07-13 15:56:52', '2022-04-07 22:35:21', NULL);
INSERT INTO `sys_dept` VALUES (202, 108, '', '外勤', 1, NULL, '18888888888', 'aa@qq.com', 1, 0, NULL, '2021-07-13 15:56:52', '2021-07-13 15:56:52', NULL);
INSERT INTO `sys_dept` VALUES (203, 108, '', '行政', 0, '', '18888888888', 'aa@qq.com', 1, 0, 31, '2021-07-13 15:56:52', '2022-09-16 16:46:47', NULL);

-- ----------------------------
-- Table structure for sys_dict_data
-- ----------------------------
DROP TABLE IF EXISTS `sys_dict_data`;
CREATE TABLE `sys_dict_data`  (
  `dict_code` bigint(20) NOT NULL AUTO_INCREMENT COMMENT '字典编码',
  `dict_sort` int(4) NULL DEFAULT 0 COMMENT '字典排序',
  `dict_label` varchar(100) CHARACTER SET utf8mb4 COLLATE utf8mb4_general_ci NULL DEFAULT '' COMMENT '字典标签',
  `dict_value` varchar(100) CHARACTER SET utf8mb4 COLLATE utf8mb4_general_ci NULL DEFAULT '' COMMENT '字典键值',
  `dict_type` varchar(100) CHARACTER SET utf8mb4 COLLATE utf8mb4_general_ci NULL DEFAULT '' COMMENT '字典类型',
  `css_class` varchar(100) CHARACTER SET utf8mb4 COLLATE utf8mb4_general_ci NULL DEFAULT NULL COMMENT '样式属性（其他样式扩展）',
  `list_class` varchar(100) CHARACTER SET utf8mb4 COLLATE utf8mb4_general_ci NULL DEFAULT NULL COMMENT '表格回显样式',
  `is_default` tinyint(1) NULL DEFAULT 0 COMMENT '是否默认（1是 0否）',
  `status` tinyint(1) NULL DEFAULT 0 COMMENT '状态（1正常 0停用）',
  `create_by` bigint(64) UNSIGNED NULL DEFAULT 0 COMMENT '创建者',
  `update_by` bigint(64) UNSIGNED NULL DEFAULT 0 COMMENT '更新者',
  `remark` varchar(500) CHARACTER SET utf8mb4 COLLATE utf8mb4_general_ci NULL DEFAULT NULL COMMENT '备注',
  `created_at` datetime NULL DEFAULT NULL COMMENT '创建时间',
  `updated_at` datetime NULL DEFAULT NULL COMMENT '修改时间',
  PRIMARY KEY (`dict_code`) USING BTREE
) ENGINE = InnoDB AUTO_INCREMENT = 122 CHARACTER SET = utf8mb4 COLLATE = utf8mb4_general_ci COMMENT = '字典数据表' ROW_FORMAT = COMPACT;

-- ----------------------------
-- Records of sys_dict_data
-- ----------------------------
INSERT INTO `sys_dict_data` VALUES (1, 0, '男', '1', 'sys_user_sex', '', '', 0, 1, 31, 2, '备注信息', '2022-04-18 16:46:22', '2023-08-21 15:07:21');
INSERT INTO `sys_dict_data` VALUES (2, 0, '女', '2', 'sys_user_sex', '', '', 0, 1, 31, 31, '备注信息', NULL, '2023-08-21 15:07:21');
INSERT INTO `sys_dict_data` VALUES (3, 0, '保密', '0', 'sys_user_sex', '', '', 1, 1, 31, 31, '备注信息', NULL, '2023-08-21 15:10:28');
INSERT INTO `sys_dict_data` VALUES (24, 0, '频道页', '1', 'cms_category_type', '', '', 0, 1, 31, 31, '作为频道页，不可作为栏目发布文章，可添加下级分类', NULL, '2021-07-21 10:54:22');
INSERT INTO `sys_dict_data` VALUES (25, 0, '发布栏目', '2', 'cms_category_type', '', '', 0, 1, 31, 31, '作为发布栏目，可添加文章', NULL, '2021-07-21 10:54:22');
INSERT INTO `sys_dict_data` VALUES (26, 0, '跳转栏目', '3', 'cms_category_type', '', '', 0, 1, 31, 31, '不直接发布内容，用于跳转页面', NULL, '2021-07-21 10:54:22');
INSERT INTO `sys_dict_data` VALUES (27, 0, '单页栏目', '4', 'cms_category_type', '', '', 0, 1, 31, 31, '单页面模式，分类直接显示为文章', NULL, '2021-07-21 10:54:22');
INSERT INTO `sys_dict_data` VALUES (28, 0, '正常', '0', 'sys_job_status', '', 'default', 1, 1, 31, 0, '', NULL, NULL);
INSERT INTO `sys_dict_data` VALUES (29, 0, '暂停', '1', 'sys_job_status', '', 'default', 0, 1, 31, 31, '', NULL, NULL);
INSERT INTO `sys_dict_data` VALUES (30, 0, '默认', 'DEFAULT', 'sys_job_group', '', 'default', 1, 1, 31, 0, '', NULL, NULL);
INSERT INTO `sys_dict_data` VALUES (31, 0, '系统', 'SYSTEM', 'sys_job_group', '', 'default', 0, 1, 31, 0, '', NULL, NULL);
INSERT INTO `sys_dict_data` VALUES (32, 0, '成功', '1', 'admin_login_status', '', 'default', 0, 1, 31, 31, '', NULL, '2022-09-16 15:26:01');
INSERT INTO `sys_dict_data` VALUES (33, 0, '失败', '0', 'admin_login_status', '', 'default', 0, 1, 31, 0, '', NULL, '2022-09-16 15:26:01');
INSERT INTO `sys_dict_data` VALUES (34, 0, '成功', '1', 'sys_oper_log_status', '', 'default', 0, 1, 31, 0, '', NULL, NULL);
INSERT INTO `sys_dict_data` VALUES (35, 0, '失败', '0', 'sys_oper_log_status', '', 'default', 0, 1, 31, 0, '', NULL, NULL);
INSERT INTO `sys_dict_data` VALUES (36, 0, '重复执行', '1', 'sys_job_policy', '', 'default', 1, 1, 31, 0, '', NULL, NULL);
INSERT INTO `sys_dict_data` VALUES (37, 0, '执行一次', '2', 'sys_job_policy', '', 'default', 1, 1, 31, 0, '', NULL, NULL);
INSERT INTO `sys_dict_data` VALUES (38, 0, '显示', '0', 'sys_show_hide', NULL, 'default', 1, 1, 31, 0, NULL, NULL, NULL);
INSERT INTO `sys_dict_data` VALUES (39, 0, '隐藏', '1', 'sys_show_hide', NULL, 'default', 0, 1, 31, 0, NULL, NULL, NULL);
INSERT INTO `sys_dict_data` VALUES (40, 0, '正常', '1', 'sys_normal_disable', '', 'default', 1, 1, 31, 0, '', NULL, NULL);
INSERT INTO `sys_dict_data` VALUES (41, 0, '停用', '0', 'sys_normal_disable', '', 'default', 0, 1, 31, 0, '', NULL, NULL);
INSERT INTO `sys_dict_data` VALUES (49, 0, '是', '1', 'sys_yes_no', '', '', 1, 1, 31, 0, '', NULL, NULL);
INSERT INTO `sys_dict_data` VALUES (50, 0, '否', '0', 'sys_yes_no', '', '', 0, 1, 31, 0, '', NULL, NULL);
INSERT INTO `sys_dict_data` VALUES (51, 0, '已发布', '1', 'cms_article_pub_type', '', '', 1, 1, 31, 31, '', NULL, NULL);
INSERT INTO `sys_dict_data` VALUES (54, 0, '未发布', '0', 'cms_article_pub_type', '', '', 0, 1, 31, 0, '', NULL, NULL);
INSERT INTO `sys_dict_data` VALUES (55, 0, '置顶', '1', 'cms_article_attr', '', '', 0, 1, 31, 0, '', NULL, NULL);
INSERT INTO `sys_dict_data` VALUES (56, 0, '推荐', '2', 'cms_article_attr', '', '', 0, 1, 31, 0, '', NULL, NULL);
INSERT INTO `sys_dict_data` VALUES (57, 0, '普通文章', '0', 'cms_article_type', '', '', 0, 1, 31, 31, '', NULL, NULL);
INSERT INTO `sys_dict_data` VALUES (58, 0, '跳转链接', '1', 'cms_article_type', '', '', 0, 1, 31, 31, '', NULL, NULL);
INSERT INTO `sys_dict_data` VALUES (59, 0, 'cms模型', '6', 'cms_cate_models', '', '', 0, 1, 1, 1, '', NULL, NULL);
INSERT INTO `sys_dict_data` VALUES (61, 0, '政府工作目标', '1', 'gov_cate_models', '', '', 0, 1, 2, 0, '', NULL, NULL);
INSERT INTO `sys_dict_data` VALUES (62, 0, '系统后台', 'sys_admin', 'menu_module_type', '', '', 1, 1, 2, 0, '', NULL, NULL);
INSERT INTO `sys_dict_data` VALUES (63, 0, '政务工作', 'gov_work', 'menu_module_type', '', '', 0, 1, 2, 0, '', NULL, NULL);
INSERT INTO `sys_dict_data` VALUES (64, 0, '幻灯', '3', 'cms_article_attr', '', '', 0, 1, 31, 0, '', NULL, NULL);
INSERT INTO `sys_dict_data` VALUES (65, 0, '[work]测试业务表', 'wf_news', 'flow_type', '', '', 0, 1, 2, 2, '', NULL, NULL);
INSERT INTO `sys_dict_data` VALUES (66, 0, '回退修改', '-1', 'flow_status', '', '', 0, 1, 31, 0, '', NULL, NULL);
INSERT INTO `sys_dict_data` VALUES (67, 0, '保存中', '0', 'flow_status', '', '', 0, 1, 31, 0, '', NULL, NULL);
INSERT INTO `sys_dict_data` VALUES (68, 0, '流程中', '1', 'flow_status', '', '', 0, 1, 31, 0, '', NULL, NULL);
INSERT INTO `sys_dict_data` VALUES (69, 0, '审批通过', '2', 'flow_status', '', '', 0, 1, 31, 2, '', NULL, NULL);
INSERT INTO `sys_dict_data` VALUES (70, 2, '发布栏目', '2', 'sys_blog_sign', '', '', 0, 1, 31, 31, '', NULL, NULL);
INSERT INTO `sys_dict_data` VALUES (71, 3, '跳转栏目', '3', 'sys_blog_sign', '', '', 0, 1, 31, 31, '', NULL, NULL);
INSERT INTO `sys_dict_data` VALUES (72, 4, '单页栏目', '4', 'sys_blog_sign', '', '', 0, 1, 31, 31, '', NULL, NULL);
INSERT INTO `sys_dict_data` VALUES (73, 2, '置顶', '1', 'sys_log_sign', '', '', 0, 1, 31, 31, '', NULL, NULL);
INSERT INTO `sys_dict_data` VALUES (74, 3, '幻灯', '2', 'sys_log_sign', '', '', 0, 1, 31, 31, '', NULL, NULL);
INSERT INTO `sys_dict_data` VALUES (75, 4, '推荐', '3', 'sys_log_sign', '', '', 0, 1, 31, 31, '', NULL, NULL);
INSERT INTO `sys_dict_data` VALUES (76, 1, '一般', '0', 'sys_log_sign', '', '', 0, 1, 31, 31, '', NULL, NULL);
INSERT INTO `sys_dict_data` VALUES (77, 1, '频道页', '1', 'sys_blog_sign', '', '', 0, 1, 31, 31, '', NULL, NULL);
INSERT INTO `sys_dict_data` VALUES (78, 0, '普通', '0', 'flow_level', '', '', 0, 1, 31, 0, '', NULL, '2021-07-20 08:55:20');
INSERT INTO `sys_dict_data` VALUES (79, 0, '加急', '1', 'flow_level', '', '', 0, 1, 31, 0, '', NULL, '2021-07-20 08:55:20');
INSERT INTO `sys_dict_data` VALUES (80, 0, '紧急', '2', 'flow_level', '', '', 0, 1, 31, 0, '', NULL, '2021-07-20 08:55:20');
INSERT INTO `sys_dict_data` VALUES (81, 0, '特急', '3', 'flow_level', '', '', 0, 1, 31, 31, '', NULL, '2021-07-20 08:55:25');
INSERT INTO `sys_dict_data` VALUES (82, 0, '频道页', '1', 'sys_blog_type', '', '', 0, 1, 31, 0, '', NULL, NULL);
INSERT INTO `sys_dict_data` VALUES (83, 0, '发布栏目', '2', 'sys_blog_type', '', '', 0, 1, 31, 0, '', NULL, NULL);
INSERT INTO `sys_dict_data` VALUES (84, 0, '跳转栏目', '3', 'sys_blog_type', '', '', 0, 1, 31, 31, '', NULL, NULL);
INSERT INTO `sys_dict_data` VALUES (85, 0, '单页栏目', '4', 'sys_blog_type', '', '', 0, 1, 31, 0, '', NULL, NULL);
INSERT INTO `sys_dict_data` VALUES (87, 0, '[cms]文章表', 'cms_news', 'flow_type', '', '', 0, 1, 31, 0, '', NULL, NULL);
INSERT INTO `sys_dict_data` VALUES (91, 0, '测试一下', '666', 'cms_article_type', '', '', 0, 1, 31, 0, '', '2021-08-03 17:04:12', '2021-08-03 17:04:12');
INSERT INTO `sys_dict_data` VALUES (92, 0, '缓存测试222', '33333', 'cms_article_type', '', '', 0, 1, 31, 31, '', '2021-08-03 17:16:45', '2021-08-03 17:19:41');
INSERT INTO `sys_dict_data` VALUES (93, 0, '缓存测试222', '11111', 'cms_article_type', '', '', 0, 1, 31, 31, '', '2021-08-03 17:26:14', '2021-08-03 17:26:26');
INSERT INTO `sys_dict_data` VALUES (94, 0, '1折', '10', 'plugin_store_discount', '', '', 0, 1, 31, 0, '', '2021-08-14 11:59:38', '2021-08-14 11:59:38');
INSERT INTO `sys_dict_data` VALUES (95, 0, '5折', '50', 'plugin_store_discount', '', '', 0, 1, 31, 0, '', '2021-08-14 11:59:49', '2021-08-14 11:59:49');
INSERT INTO `sys_dict_data` VALUES (96, 0, '8折', '80', 'plugin_store_discount', '', '', 0, 1, 31, 0, '', '2021-08-14 12:00:00', '2021-08-14 12:00:00');
INSERT INTO `sys_dict_data` VALUES (97, 0, '9折', '90', 'plugin_store_discount', '', '', 0, 1, 31, 0, '', '2021-08-14 12:00:07', '2021-08-14 12:00:07');
INSERT INTO `sys_dict_data` VALUES (98, 0, '无折扣', '100', 'plugin_store_discount', '', '', 0, 1, 31, 0, '', '2021-08-14 12:00:16', '2021-08-14 12:00:16');
INSERT INTO `sys_dict_data` VALUES (99, 0, '不显示', 'none', 'cms_nav_position', '', '', 1, 1, 22, 0, '', '2021-08-31 15:37:35', '2021-08-31 15:37:35');
INSERT INTO `sys_dict_data` VALUES (100, 0, '顶部导航', 'top', 'cms_nav_position', '', '', 0, 1, 22, 0, '', '2021-08-31 15:37:57', '2021-08-31 15:37:57');
INSERT INTO `sys_dict_data` VALUES (101, 0, '底部导航', 'bottom', 'cms_nav_position', '', '', 0, 1, 22, 0, '', '2021-08-31 15:38:08', '2021-08-31 15:38:08');
INSERT INTO `sys_dict_data` VALUES (102, 0, '读取', 'GET', 'sys_oper_log_type', '', '', 0, 1, 31, 31, '', '2022-12-21 11:59:10', '2022-12-23 19:03:02');
INSERT INTO `sys_dict_data` VALUES (103, 0, '新增', 'POST', 'sys_oper_log_type', '', '', 0, 1, 31, 31, '', '2022-12-21 11:59:22', '2022-12-23 19:03:10');
INSERT INTO `sys_dict_data` VALUES (104, 0, '修改', 'PUT', 'sys_oper_log_type', '', '', 0, 1, 31, 31, '', '2022-12-21 11:59:32', '2022-12-23 19:03:19');
INSERT INTO `sys_dict_data` VALUES (105, 0, '删除', 'DELETE', 'sys_oper_log_type', '', '', 0, 1, 31, 31, '', '2022-12-21 11:59:44', '2022-12-23 19:03:27');
INSERT INTO `sys_dict_data` VALUES (106, 0, '无标签', '0', 'notice_tag', '', '', 0, 1, 31, 31, '', '2023-12-28 15:48:45', '2023-12-28 15:52:14');
INSERT INTO `sys_dict_data` VALUES (107, 0, '提醒', '1', 'notice_tag', '', '', 0, 1, 31, 31, '', '2023-12-28 15:48:54', '2023-12-28 15:52:24');
INSERT INTO `sys_dict_data` VALUES (108, 0, '一般', '2', 'notice_tag', '', '', 0, 1, 31, 0, '', '2023-12-28 15:52:35', '2023-12-28 15:52:35');
INSERT INTO `sys_dict_data` VALUES (109, 0, '次要', '3', 'notice_tag', '', '', 0, 1, 31, 0, '', '2023-12-28 15:52:44', '2023-12-28 15:52:44');
INSERT INTO `sys_dict_data` VALUES (110, 0, '重要', '4', 'notice_tag', '', '', 0, 1, 31, 0, '', '2023-12-28 15:52:53', '2023-12-28 15:52:53');
INSERT INTO `sys_dict_data` VALUES (111, 0, '紧急', '5', 'notice_tag', '', '', 0, 1, 31, 0, '', '2023-12-28 15:53:01', '2023-12-28 15:53:01');
INSERT INTO `sys_dict_data` VALUES (112, 0, '本地上传', '0', 'sys_upload_drive', '', '', 0, 1, 31, 0, '', '2024-10-23 14:37:27', '2024-10-23 14:37:27');
INSERT INTO `sys_dict_data` VALUES (113, 0, '腾讯云', '1', 'sys_upload_drive', '', '', 0, 1, 31, 31, '', '2024-10-23 14:37:38', '2024-10-23 14:38:05');
INSERT INTO `sys_dict_data` VALUES (114, 0, '七牛云', '2', 'sys_upload_drive', '', '', 0, 1, 31, 31, '', '2024-10-23 14:37:52', '2024-10-23 14:38:43');
INSERT INTO `sys_dict_data` VALUES (115, 0, '阿里云', '3', 'sys_upload_drive', '', '', 0, 1, 31, 31, '', '2024-10-23 14:38:11', '2024-10-23 14:38:49');
INSERT INTO `sys_dict_data` VALUES (116, 0, '图片', 'image', 'sys_upload_file_type', '', '', 0, 1, 31, 0, '', '2024-10-23 14:54:18', '2024-10-23 14:54:18');
INSERT INTO `sys_dict_data` VALUES (117, 0, '文档', 'doc', 'sys_upload_file_type', '', '', 0, 1, 31, 0, '', '2024-10-23 14:54:32', '2024-10-23 14:54:32');
INSERT INTO `sys_dict_data` VALUES (118, 0, '音频', 'audio', 'sys_upload_file_type', '', '', 0, 1, 31, 0, '', '2024-10-23 14:54:55', '2024-10-23 14:54:55');
INSERT INTO `sys_dict_data` VALUES (119, 0, '视频', 'video', 'sys_upload_file_type', '', '', 0, 1, 31, 0, '', '2024-10-23 14:55:09', '2024-10-23 14:55:09');
INSERT INTO `sys_dict_data` VALUES (120, 0, '压缩包', 'zip', 'sys_upload_file_type', '', '', 0, 1, 31, 0, '', '2024-10-23 14:55:22', '2024-10-23 14:55:22');
INSERT INTO `sys_dict_data` VALUES (121, 0, '其它', 'other', 'sys_upload_file_type', '', '', 0, 1, 31, 0, '', '2024-10-23 14:55:40', '2024-10-23 14:55:40');

-- ----------------------------
-- Table structure for sys_dict_type
-- ----------------------------
DROP TABLE IF EXISTS `sys_dict_type`;
CREATE TABLE `sys_dict_type`  (
  `dict_id` bigint(20) UNSIGNED NOT NULL AUTO_INCREMENT COMMENT '字典主键',
  `dict_name` varchar(100) CHARACTER SET utf8mb4 COLLATE utf8mb4_general_ci NULL DEFAULT '' COMMENT '字典名称',
  `dict_type` varchar(100) CHARACTER SET utf8mb4 COLLATE utf8mb4_general_ci NULL DEFAULT '' COMMENT '字典类型',
  `status` tinyint(1) UNSIGNED NULL DEFAULT 0 COMMENT '状态（0正常 1停用）',
  `create_by` int(64) UNSIGNED NULL DEFAULT 0 COMMENT '创建者',
  `update_by` int(64) UNSIGNED NULL DEFAULT 0 COMMENT '更新者',
  `remark` varchar(500) CHARACTER SET utf8mb4 COLLATE utf8mb4_general_ci NULL DEFAULT NULL COMMENT '备注',
  `created_at` datetime NULL DEFAULT NULL COMMENT '创建日期',
  `updated_at` datetime NULL DEFAULT NULL COMMENT '修改日期',
  PRIMARY KEY (`dict_id`) USING BTREE,
  UNIQUE INDEX `dict_type`(`dict_type`) USING BTREE
) ENGINE = InnoDB AUTO_INCREMENT = 54 CHARACTER SET = utf8mb4 COLLATE = utf8mb4_general_ci COMMENT = '字典类型表' ROW_FORMAT = COMPACT;

-- ----------------------------
-- Records of sys_dict_type
-- ----------------------------
INSERT INTO `sys_dict_type` VALUES (1, '用户性别', 'sys_user_sex', 1, 31, 31, '用于选择用户性别', NULL, '2023-08-21 15:07:21');
INSERT INTO `sys_dict_type` VALUES (2, '分类类型', 'cms_category_type', 1, 31, 3, '文章分类类型', NULL, '2021-07-21 10:54:22');
INSERT INTO `sys_dict_type` VALUES (3, '任务状态', 'sys_job_status', 1, 31, 31, '任务状态列表', NULL, NULL);
INSERT INTO `sys_dict_type` VALUES (13, '任务分组', 'sys_job_group', 1, 31, 0, '', NULL, NULL);
INSERT INTO `sys_dict_type` VALUES (14, '管理员登录状态', 'admin_login_status', 1, 31, 31, '', NULL, '2022-09-16 15:26:01');
INSERT INTO `sys_dict_type` VALUES (15, '操作日志状态', 'sys_oper_log_status', 1, 31, 0, '', NULL, NULL);
INSERT INTO `sys_dict_type` VALUES (16, '任务策略', 'sys_job_policy', 1, 31, 0, '', NULL, NULL);
INSERT INTO `sys_dict_type` VALUES (17, '菜单状态', 'sys_show_hide', 1, 31, 0, '菜单状态', NULL, NULL);
INSERT INTO `sys_dict_type` VALUES (18, '系统开关', 'sys_normal_disable', 1, 31, 31, '系统开关', NULL, NULL);
INSERT INTO `sys_dict_type` VALUES (24, '系统内置', 'sys_yes_no', 1, 31, 0, '', NULL, NULL);
INSERT INTO `sys_dict_type` VALUES (25, '文章发布状态', 'cms_article_pub_type', 1, 31, 0, '', NULL, NULL);
INSERT INTO `sys_dict_type` VALUES (26, '文章附加状态', 'cms_article_attr', 1, 31, 0, '', NULL, NULL);
INSERT INTO `sys_dict_type` VALUES (27, '文章类型', 'cms_article_type', 1, 31, 0, '', NULL, NULL);
INSERT INTO `sys_dict_type` VALUES (28, '文章栏目模型分类', 'cms_cate_models', 1, 1, 0, '', NULL, NULL);
INSERT INTO `sys_dict_type` VALUES (29, '政务工作模型分类', 'gov_cate_models', 1, 2, 0, '', NULL, NULL);
INSERT INTO `sys_dict_type` VALUES (30, '菜单模块类型', 'menu_module_type', 1, 2, 0, '', NULL, NULL);
INSERT INTO `sys_dict_type` VALUES (31, '工作流程类型', 'flow_type', 1, 2, 0, '', NULL, NULL);
INSERT INTO `sys_dict_type` VALUES (32, '工作流程审批状态', 'flow_status', 1, 31, 0, '工作流程审批状态', NULL, NULL);
INSERT INTO `sys_dict_type` VALUES (33, '博客分类类型', 'sys_blog_type', 1, 31, 31, '博客分类中的标志', NULL, NULL);
INSERT INTO `sys_dict_type` VALUES (34, '博客日志标志', 'sys_log_sign', 1, 31, 0, '博客日志管理中的标志数据字典', NULL, NULL);
INSERT INTO `sys_dict_type` VALUES (35, '工作流紧急状态', 'flow_level', 1, 31, 31, '', NULL, '2021-07-20 08:55:20');
INSERT INTO `sys_dict_type` VALUES (48, '插件商城折扣', 'plugin_store_discount', 1, 31, 0, '', '2021-08-14 11:59:26', '2021-08-14 11:59:26');
INSERT INTO `sys_dict_type` VALUES (49, 'CMS栏目导航位置', 'cms_nav_position', 1, 22, 0, '', '2021-08-31 15:37:04', '2021-08-31 15:37:04');
INSERT INTO `sys_dict_type` VALUES (50, '操作日志类型', 'sys_oper_log_type', 1, 31, 0, '', '2022-12-21 11:55:02', '2022-12-21 11:55:02');
INSERT INTO `sys_dict_type` VALUES (51, '系统公告标签', 'notice_tag', 1, 31, 0, '', '2023-12-28 15:48:03', '2023-12-28 15:48:03');
INSERT INTO `sys_dict_type` VALUES (52, '附件上传驱动', 'sys_upload_drive', 1, 31, 31, '', '2024-10-23 14:36:17', '2024-10-23 14:36:30');
INSERT INTO `sys_dict_type` VALUES (53, '上传文件类型', 'sys_upload_file_type', 1, 31, 0, '', '2024-10-23 14:53:50', '2024-10-23 14:53:50');


-- ----------------------------
-- Records of sys_job
-- ----------------------------

-- ----------------------------
-- Table structure for sys_job_log
-- ----------------------------
DROP TABLE IF EXISTS `sys_job_log`;
CREATE TABLE `sys_job_log`  (
  `id` bigint(20) UNSIGNED NOT NULL AUTO_INCREMENT COMMENT '主键',
  `target_name` varchar(100) CHARACTER SET utf8mb4 COLLATE utf8mb4_general_ci NOT NULL COMMENT '方法名',
  `created_at` datetime NULL DEFAULT NULL COMMENT '执行日期',
  `result` varchar(100) CHARACTER SET utf8mb4 COLLATE utf8mb4_general_ci NULL DEFAULT NULL COMMENT '执行结果',
  PRIMARY KEY (`id`) USING BTREE
) ENGINE = InnoDB AUTO_INCREMENT = 335 CHARACTER SET = utf8mb4 COLLATE = utf8mb4_general_ci COMMENT = '任务日志表' ROW_FORMAT = DYNAMIC;

-- ----------------------------
-- Records of sys_job_log
-- ----------------------------
INSERT INTO `sys_job_log` VALUES (1, 'checkUserOnline', '2024-04-12 17:20:05', '在线用户定时更新，执行成功');

-- ----------------------------
-- Table structure for sys_login_log
-- ----------------------------
DROP TABLE IF EXISTS `sys_login_log`;
CREATE TABLE `sys_login_log`  (
  `info_id` bigint(20) NOT NULL AUTO_INCREMENT COMMENT '访问ID',
  `login_name` varchar(50) CHARACTER SET utf8mb4 COLLATE utf8mb4_general_ci NULL DEFAULT '' COMMENT '登录账号',
  `ipaddr` varchar(50) CHARACTER SET utf8mb4 COLLATE utf8mb4_general_ci NULL DEFAULT '' COMMENT '登录IP地址',
  `login_location` varchar(255) CHARACTER SET utf8mb4 COLLATE utf8mb4_general_ci NULL DEFAULT '' COMMENT '登录地点',
  `browser` varchar(50) CHARACTER SET utf8mb4 COLLATE utf8mb4_general_ci NULL DEFAULT '' COMMENT '浏览器类型',
  `os` varchar(50) CHARACTER SET utf8mb4 COLLATE utf8mb4_general_ci NULL DEFAULT '' COMMENT '操作系统',
  `status` tinyint(4) NULL DEFAULT 0 COMMENT '登录状态（0成功 1失败）',
  `msg` varchar(255) CHARACTER SET utf8mb4 COLLATE utf8mb4_general_ci NULL DEFAULT '' COMMENT '提示消息',
  `login_time` datetime NULL DEFAULT NULL COMMENT '登录时间',
  `module` varchar(30) CHARACTER SET utf8mb4 COLLATE utf8mb4_general_ci NULL DEFAULT '' COMMENT '登录模块',
  PRIMARY KEY (`info_id`) USING BTREE
) ENGINE = InnoDB AUTO_INCREMENT = 8 CHARACTER SET = utf8mb4 COLLATE = utf8mb4_general_ci COMMENT = '系统访问记录' ROW_FORMAT = COMPACT;

-- ----------------------------
-- Records of sys_login_log
-- ----------------------------

-- ----------------------------
-- Table structure for sys_notice
-- ----------------------------
DROP TABLE IF EXISTS `sys_notice`;
CREATE TABLE `sys_notice`  (
  `id` bigint(20) NOT NULL AUTO_INCREMENT COMMENT 'ID',
  `title` varchar(64) CHARACTER SET utf8 COLLATE utf8_general_ci NOT NULL COMMENT '标题',
  `type` bigint(20) NOT NULL COMMENT '类型',
  `tag` int(11) NULL DEFAULT NULL COMMENT '标签',
  `content` longtext CHARACTER SET utf8 COLLATE utf8_general_ci NOT NULL COMMENT '内容',
  `remark` varchar(255) CHARACTER SET utf8 COLLATE utf8_general_ci NULL DEFAULT NULL COMMENT '备注',
  `sort` int(11) NOT NULL DEFAULT 0 COMMENT '排序',
  `status` tinyint(1) NULL DEFAULT 1 COMMENT '状态',
  `created_by` bigint(20) NULL DEFAULT NULL COMMENT '发送人',
  `updated_by` bigint(20) NULL DEFAULT 0 COMMENT '修改人',
  `created_at` datetime NULL DEFAULT NULL COMMENT '创建时间',
  `updated_at` datetime NULL DEFAULT NULL COMMENT '更新时间',
  `deleted_at` datetime NULL DEFAULT NULL COMMENT '删除时间',
  `receiver` json NULL COMMENT '接收者（私信）',
  PRIMARY KEY (`id`) USING BTREE
) ENGINE = InnoDB AUTO_INCREMENT = 3 CHARACTER SET = utf8 COLLATE = utf8_general_ci COMMENT = '通知公告' ROW_FORMAT = Dynamic;

-- ----------------------------
-- Records of sys_notice
-- ----------------------------

-- ----------------------------
-- Table structure for sys_notice_read
-- ----------------------------
DROP TABLE IF EXISTS `sys_notice_read`;
CREATE TABLE `sys_notice_read`  (
  `id` bigint(20) NOT NULL AUTO_INCREMENT COMMENT 'id',
  `notice_id` bigint(20) NOT NULL COMMENT '信息id',
  `user_id` bigint(20) NOT NULL COMMENT '用户id',
  `clicks` int(11) NULL DEFAULT NULL COMMENT '点击次数',
  `updated_at` datetime NULL DEFAULT NULL COMMENT '更新时间',
  `created_at` datetime NULL DEFAULT NULL COMMENT '阅读时间',
  PRIMARY KEY (`id`) USING BTREE,
  UNIQUE INDEX `notice_id`(`notice_id`, `user_id`) USING BTREE
) ENGINE = InnoDB AUTO_INCREMENT = 2 CHARACTER SET = utf8 COLLATE = utf8_general_ci COMMENT = '已读记录' ROW_FORMAT = Dynamic;

-- ----------------------------
-- Records of sys_notice_read
-- ----------------------------
INSERT INTO `sys_notice_read` VALUES (1, 1, 31, 0, '2024-01-02 17:47:05', '2024-01-02 17:47:05');

-- ----------------------------
-- Table structure for sys_oper_log
-- ----------------------------
DROP TABLE IF EXISTS `sys_oper_log`;
CREATE TABLE `sys_oper_log`  (
  `oper_id` bigint(20) UNSIGNED NOT NULL AUTO_INCREMENT COMMENT '日志主键',
  `title` varchar(50) CHARACTER SET utf8mb4 COLLATE utf8mb4_general_ci NULL DEFAULT '' COMMENT '模块标题',
  `business_type` int(2) NULL DEFAULT 0 COMMENT '业务类型（0其它 1新增 2修改 3删除）',
  `method` varchar(100) CHARACTER SET utf8mb4 COLLATE utf8mb4_general_ci NULL DEFAULT '' COMMENT '方法名称',
  `request_method` varchar(10) CHARACTER SET utf8mb4 COLLATE utf8mb4_general_ci NULL DEFAULT '' COMMENT '请求方式',
  `operator_type` int(1) NULL DEFAULT 0 COMMENT '操作类别（0其它 1后台用户 2手机端用户）',
  `oper_name` varchar(50) CHARACTER SET utf8mb4 COLLATE utf8mb4_general_ci NULL DEFAULT '' COMMENT '操作人员',
  `dept_name` varchar(50) CHARACTER SET utf8mb4 COLLATE utf8mb4_general_ci NULL DEFAULT '' COMMENT '部门名称',
  `oper_url` varchar(500) CHARACTER SET utf8mb4 COLLATE utf8mb4_general_ci NULL DEFAULT '' COMMENT '请求URL',
  `oper_ip` varchar(50) CHARACTER SET utf8mb4 COLLATE utf8mb4_general_ci NULL DEFAULT '' COMMENT '主机地址',
  `oper_location` varchar(255) CHARACTER SET utf8mb4 COLLATE utf8mb4_general_ci NULL DEFAULT '' COMMENT '操作地点',
  `oper_param` text CHARACTER SET utf8mb4 COLLATE utf8mb4_general_ci NULL COMMENT '请求参数',
  `error_msg` text CHARACTER SET utf8mb4 COLLATE utf8mb4_general_ci NULL COMMENT '错误消息',
  `oper_time` datetime NULL DEFAULT NULL COMMENT '操作时间',
  PRIMARY KEY (`oper_id`) USING BTREE
) ENGINE = InnoDB AUTO_INCREMENT = 133 CHARACTER SET = utf8mb4 COLLATE = utf8mb4_general_ci COMMENT = '操作日志记录' ROW_FORMAT = COMPACT;

-- ----------------------------
-- Records of sys_oper_log
-- ----------------------------
-- INSERT INTO `sys_oper_log` VALUES (1, '', 0, '/api/v1/system/operLog/clear', 'DELETE', 1, 'demo', '财务部门', '/api/v1/system/operLog/clear', '::1', '内网IP', '{}', NULL, '2024-09-30 17:53:53');

-- ----------------------------
-- Table structure for sys_post
-- ----------------------------
DROP TABLE IF EXISTS `sys_post`;
CREATE TABLE `sys_post`  (
  `post_id` bigint(20) UNSIGNED NOT NULL AUTO_INCREMENT COMMENT '岗位ID',
  `post_code` varchar(64) CHARACTER SET utf8mb4 COLLATE utf8mb4_general_ci NOT NULL COMMENT '岗位编码',
  `post_name` varchar(50) CHARACTER SET utf8mb4 COLLATE utf8mb4_general_ci NOT NULL COMMENT '岗位名称',
  `post_sort` int(4) NOT NULL COMMENT '显示顺序',
  `status` tinyint(1) UNSIGNED NOT NULL DEFAULT 0 COMMENT '状态（0正常 1停用）',
  `remark` varchar(500) CHARACTER SET utf8mb4 COLLATE utf8mb4_general_ci NULL DEFAULT NULL COMMENT '备注',
  `created_by` bigint(20) UNSIGNED NOT NULL DEFAULT 0 COMMENT '创建人',
  `updated_by` bigint(20) UNSIGNED NOT NULL DEFAULT 0 COMMENT '修改人',
  `created_at` datetime NULL DEFAULT NULL COMMENT '创建时间',
  `updated_at` datetime NULL DEFAULT NULL COMMENT '修改时间',
  `deleted_at` datetime NULL DEFAULT NULL COMMENT '删除时间',
  PRIMARY KEY (`post_id`) USING BTREE
) ENGINE = InnoDB AUTO_INCREMENT = 10 CHARACTER SET = utf8mb4 COLLATE = utf8mb4_general_ci COMMENT = '岗位信息表' ROW_FORMAT = COMPACT;

-- ----------------------------
-- Records of sys_post
-- ----------------------------


-- ----------------------------
-- Table structure for sys_role_dept
-- ----------------------------
DROP TABLE IF EXISTS `sys_role_dept`;
CREATE TABLE `sys_role_dept`  (
  `role_id` bigint(20) NOT NULL COMMENT '角色ID',
  `dept_id` bigint(20) NOT NULL COMMENT '部门ID',
  PRIMARY KEY (`role_id`, `dept_id`) USING BTREE
) ENGINE = InnoDB CHARACTER SET = utf8mb4 COLLATE = utf8mb4_general_ci COMMENT = '角色和部门关联表' ROW_FORMAT = COMPACT;

-- ----------------------------
-- Records of sys_role_dept
-- ----------------------------
INSERT INTO `sys_role_dept` VALUES (1, 101);
INSERT INTO `sys_role_dept` VALUES (1, 103);
INSERT INTO `sys_role_dept` VALUES (1, 104);
INSERT INTO `sys_role_dept` VALUES (1, 105);
INSERT INTO `sys_role_dept` VALUES (1, 106);
INSERT INTO `sys_role_dept` VALUES (1, 107);
INSERT INTO `sys_role_dept` VALUES (5, 103);
INSERT INTO `sys_role_dept` VALUES (5, 104);
INSERT INTO `sys_role_dept` VALUES (5, 105);
INSERT INTO `sys_role_dept` VALUES (8, 105);
INSERT INTO `sys_role_dept` VALUES (8, 106);

-- ----------------------------
-- Table structure for sys_role_scope
-- ----------------------------
DROP TABLE IF EXISTS `sys_role_scope`;
CREATE TABLE `sys_role_scope`  (
  `id` bigint(20) UNSIGNED NOT NULL AUTO_INCREMENT COMMENT 'ID',
  `role_id` bigint(20) UNSIGNED NOT NULL DEFAULT 0 COMMENT '角色id',
  `menu_id` int(11) NOT NULL COMMENT 'api接口id',
  `data_scope` int(11) NOT NULL COMMENT '数据范围（1：全部数据权限 2：自定数据权限 3：本部门数据权限 4：本部门及以下数据权限）',
  `dept_ids` json NULL COMMENT '扩展数据',
  PRIMARY KEY (`id`) USING BTREE,
  UNIQUE INDEX `role_id`(`role_id`, `menu_id`) USING BTREE
) ENGINE = InnoDB AUTO_INCREMENT = 99 CHARACTER SET = utf8mb4 COLLATE = utf8mb4_general_ci COMMENT = '角色数据权限' ROW_FORMAT = DYNAMIC;

-- ----------------------------
-- Records of sys_role_scope
-- ----------------------------
INSERT INTO `sys_role_scope` VALUES (80, 2, 59, 5, '[]');
INSERT INTO `sys_role_scope` VALUES (81, 2, 60, 5, '[]');
INSERT INTO `sys_role_scope` VALUES (82, 2, 61, 5, '[101, 103, 104, 105, 106, 107]');
INSERT INTO `sys_role_scope` VALUES (83, 2, 62, 5, '[]');
INSERT INTO `sys_role_scope` VALUES (84, 2, 63, 5, '[]');
INSERT INTO `sys_role_scope` VALUES (85, 2, 64, 5, '[]');
INSERT INTO `sys_role_scope` VALUES (86, 2, 65, 5, '[]');
INSERT INTO `sys_role_scope` VALUES (87, 2, 120, 1, '[]');
INSERT INTO `sys_role_scope` VALUES (88, 2, 121, 1, '[]');
INSERT INTO `sys_role_scope` VALUES (89, 2, 122, 1, '[]');
INSERT INTO `sys_role_scope` VALUES (90, 2, 123, 1, '[]');
INSERT INTO `sys_role_scope` VALUES (91, 2, 124, 1, '[]');
INSERT INTO `sys_role_scope` VALUES (92, 2, 125, 1, '[]');
INSERT INTO `sys_role_scope` VALUES (93, 2, 26, 4, '[]');
INSERT INTO `sys_role_scope` VALUES (94, 2, 116, 4, '[]');
INSERT INTO `sys_role_scope` VALUES (95, 2, 117, 4, '[]');
INSERT INTO `sys_role_scope` VALUES (96, 2, 118, 4, '[]');
INSERT INTO `sys_role_scope` VALUES (97, 2, 119, 4, '[]');
INSERT INTO `sys_role_scope` VALUES (98, 2, 141, 4, '[]');


-- ----------------------------
-- Table structure for sys_user_online
-- ----------------------------
DROP TABLE IF EXISTS `sys_user_online`;
CREATE TABLE `sys_user_online`  (
  `id` bigint(20) UNSIGNED NOT NULL AUTO_INCREMENT,
  `uuid` char(32) CHARACTER SET latin1 COLLATE latin1_general_ci NOT NULL DEFAULT '' COMMENT '用户标识',
  `token` varchar(255) CHARACTER SET latin1 COLLATE latin1_general_ci NOT NULL DEFAULT '' COMMENT '用户token',
  `create_time` datetime NULL DEFAULT NULL COMMENT '登录时间',
  `user_name` varchar(255) CHARACTER SET utf8mb4 COLLATE utf8mb4_general_ci NOT NULL COMMENT '用户名',
  `ip` varchar(120) CHARACTER SET utf8mb4 COLLATE utf8mb4_general_ci NOT NULL DEFAULT '' COMMENT '登录ip',
  `explorer` varchar(30) CHARACTER SET utf8mb4 COLLATE utf8mb4_general_ci NOT NULL DEFAULT '' COMMENT '浏览器',
  `os` varchar(150) CHARACTER SET utf8mb4 COLLATE utf8mb4_general_ci NOT NULL DEFAULT '' COMMENT '操作系统',
  PRIMARY KEY (`id`) USING BTREE,
  UNIQUE INDEX `uni_token`(`token`) USING BTREE
) ENGINE = InnoDB AUTO_INCREMENT = 5 CHARACTER SET = utf8mb4 COLLATE = utf8mb4_general_ci COMMENT = '用户在线状态表' ROW_FORMAT = COMPACT;

-- ----------------------------
-- Table structure for sys_user_post
-- ----------------------------
DROP TABLE IF EXISTS `sys_user_post`;
CREATE TABLE `sys_user_post`  (
  `user_id` bigint(20) NOT NULL COMMENT '用户ID',
  `post_id` bigint(20) NOT NULL COMMENT '岗位ID',
  PRIMARY KEY (`user_id`, `post_id`) USING BTREE
) ENGINE = InnoDB CHARACTER SET = utf8mb4 COLLATE = utf8mb4_general_ci COMMENT = '用户与岗位关联表' ROW_FORMAT = COMPACT;

-- ----------------------------
-- Records of sys_user_post
-- ----------------------------
INSERT INTO `sys_user_post` VALUES (1, 2);
INSERT INTO `sys_user_post` VALUES (1, 3);
INSERT INTO `sys_user_post` VALUES (2, 1);
INSERT INTO `sys_user_post` VALUES (2, 2);
INSERT INTO `sys_user_post` VALUES (3, 2);
INSERT INTO `sys_user_post` VALUES (4, 1);
INSERT INTO `sys_user_post` VALUES (5, 2);
INSERT INTO `sys_user_post` VALUES (10, 1);
INSERT INTO `sys_user_post` VALUES (10, 2);
INSERT INTO `sys_user_post` VALUES (10, 3);
INSERT INTO `sys_user_post` VALUES (10, 4);
INSERT INTO `sys_user_post` VALUES (10, 5);
INSERT INTO `sys_user_post` VALUES (15, 4);
INSERT INTO `sys_user_post` VALUES (16, 2);
INSERT INTO `sys_user_post` VALUES (22, 1);
INSERT INTO `sys_user_post` VALUES (22, 2);
INSERT INTO `sys_user_post` VALUES (31, 2);
INSERT INTO `sys_user_post` VALUES (34, 1);
INSERT INTO `sys_user_post` VALUES (35, 2);
INSERT INTO `sys_user_post` VALUES (35, 3);
INSERT INTO `sys_user_post` VALUES (36, 1);
INSERT INTO `sys_user_post` VALUES (37, 3);
INSERT INTO `sys_user_post` VALUES (38, 2);
INSERT INTO `sys_user_post` VALUES (38, 3);
INSERT INTO `sys_user_post` VALUES (42, 2);
INSERT INTO `sys_user_post` VALUES (42, 3);


SET FOREIGN_KEY_CHECKS = 1;
