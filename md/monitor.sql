/*
 Navicat MySQL Data Transfer

 Source Server         : 大杂烩
 Source Server Type    : MySQL
 Source Server Version : 80019
 Source Host           : 192.168.0.125:3306
 Source Schema         : monitor

 Target Server Type    : MySQL
 Target Server Version : 80019
 File Encoding         : 65001

 Date: 30/09/2022 13:36:07
*/

SET NAMES utf8mb4;
SET FOREIGN_KEY_CHECKS = 0;

-- ----------------------------
-- Table structure for mon_admin
-- ----------------------------
DROP TABLE IF EXISTS `mon_admin`;
CREATE TABLE `mon_admin` (
  `id` bigint NOT NULL AUTO_INCREMENT COMMENT '主键id',
  `username` varchar(255) COLLATE utf8mb4_general_ci NOT NULL COMMENT '用户昵称',
  `real_name` varchar(255) COLLATE utf8mb4_general_ci NOT NULL COMMENT '真实姓名',
  `password` varchar(32) COLLATE utf8mb4_general_ci NOT NULL COMMENT '密码',
  `phone` varchar(11) COLLATE utf8mb4_general_ci DEFAULT NULL COMMENT '手机号',
  `role_id` varchar(255) CHARACTER SET utf8mb4 COLLATE utf8mb4_general_ci DEFAULT NULL COMMENT '角色id',
  `department` varchar(255) CHARACTER SET utf8mb4 COLLATE utf8mb4_general_ci DEFAULT NULL COMMENT '部门',
  `office_post` varchar(255) CHARACTER SET utf8mb4 COLLATE utf8mb4_general_ci DEFAULT NULL COMMENT '职位',
  `state` int NOT NULL DEFAULT '2' COMMENT '状态 1: 正常 2: 禁用',
  `last_login_time` timestamp NULL DEFAULT NULL COMMENT '上次登陆时间',
  `created_at` timestamp NOT NULL DEFAULT CURRENT_TIMESTAMP COMMENT '创建时间',
  `updated_at` timestamp NOT NULL DEFAULT CURRENT_TIMESTAMP ON UPDATE CURRENT_TIMESTAMP COMMENT '更新时间',
  `version` bigint NOT NULL COMMENT '版本号',
  PRIMARY KEY (`id`)
) ENGINE=InnoDB AUTO_INCREMENT=29 DEFAULT CHARSET=utf8mb4 COLLATE=utf8mb4_general_ci;
INSERT INTO `monitor`.`mon_admin`(`id`, `username`, `real_name`, `password`, `phone`, `role_id`, `department`, `office_post`, `state`, `last_login_time`, `created_at`, `updated_at`, `version`) VALUES (1, 'admin', 'admin', '37029f3c646ca4d471884ac903754946', '18727361635', '1', '技术部', '运维工程师', 1, '2022-09-19 13:46:51', '2022-09-01 18:56:53', '2022-09-19 13:46:51', 21);


-- ----------------------------
-- Table structure for mon_machine
-- ----------------------------
DROP TABLE IF EXISTS `mon_machine`;
CREATE TABLE `mon_machine` (
  `id` bigint NOT NULL AUTO_INCREMENT COMMENT '机器主键id',
  `machine_code` varchar(255) COLLATE utf8mb4_general_ci NOT NULL COMMENT '机器码',
  `ip` varchar(255) COLLATE utf8mb4_general_ci NOT NULL COMMENT '服务器ip',
  `hostname` varchar(255) COLLATE utf8mb4_general_ci DEFAULT NULL COMMENT '系统名',
  `remark` varchar(255) COLLATE utf8mb4_general_ci DEFAULT NULL COMMENT '备注',
  `created_at` timestamp NOT NULL DEFAULT CURRENT_TIMESTAMP COMMENT '创建时间',
  `updated_at` timestamp NOT NULL DEFAULT CURRENT_TIMESTAMP ON UPDATE CURRENT_TIMESTAMP COMMENT '更新时间',
  `version` bigint NOT NULL COMMENT '版本号',
  PRIMARY KEY (`id`)
) ENGINE=InnoDB AUTO_INCREMENT=25 DEFAULT CHARSET=utf8mb4 COLLATE=utf8mb4_general_ci;

-- ----------------------------
-- Table structure for mon_machine_check_record
-- ----------------------------
DROP TABLE IF EXISTS `mon_machine_check_record`;
CREATE TABLE `mon_machine_check_record` (
  `id` bigint NOT NULL AUTO_INCREMENT COMMENT '主键Id',
  `machine_ip` varchar(255) COLLATE utf8mb4_general_ci NOT NULL COMMENT '机器ip',
  `machine_name` varchar(255) CHARACTER SET utf8mb4 COLLATE utf8mb4_general_ci DEFAULT NULL COMMENT '机器名字',
  `category` varchar(255) CHARACTER SET utf8mb4 COLLATE utf8mb4_general_ci NOT NULL COMMENT '类型 ',
  `percent` varchar(255) COLLATE utf8mb4_general_ci NOT NULL COMMENT '百分比',
  `level` varchar(255) COLLATE utf8mb4_general_ci NOT NULL COMMENT '等级 低危: 60% 中危: 70% 高危: 80%',
  `created_at` timestamp NOT NULL COMMENT '创建时间',
  `updated_at` timestamp NOT NULL COMMENT '更新时间',
  `version` bigint NOT NULL COMMENT '版本号',
  PRIMARY KEY (`id`)
) ENGINE=InnoDB AUTO_INCREMENT=31 DEFAULT CHARSET=utf8mb4 COLLATE=utf8mb4_general_ci;

-- ----------------------------
-- Table structure for mon_menus
-- ----------------------------
DROP TABLE IF EXISTS `mon_menus`;
CREATE TABLE `mon_menus` (
  `id` bigint NOT NULL AUTO_INCREMENT COMMENT '主键id',
  `parent_id` varchar(10) CHARACTER SET utf8mb4 COLLATE utf8mb4_general_ci NOT NULL COMMENT '父级菜单',
  `menu_name` varchar(255) CHARACTER SET utf8mb4 COLLATE utf8mb4_general_ci NOT NULL COMMENT '权限名字',
  `icons` varchar(255) CHARACTER SET utf8mb4 COLLATE utf8mb4_general_ci DEFAULT NULL COMMENT 'icons图标',
  `url` varchar(255) CHARACTER SET utf8mb4 COLLATE utf8mb4_general_ci DEFAULT NULL COMMENT '路径',
  `front_url` varchar(255) CHARACTER SET utf8mb4 COLLATE utf8mb4_general_ci DEFAULT NULL COMMENT '前端地址url',
  `method` varchar(20) CHARACTER SET utf8mb4 COLLATE utf8mb4_general_ci DEFAULT '' COMMENT '方法',
  `created_at` timestamp NOT NULL DEFAULT CURRENT_TIMESTAMP COMMENT '创建时间',
  `updated_at` timestamp NOT NULL DEFAULT CURRENT_TIMESTAMP ON UPDATE CURRENT_TIMESTAMP COMMENT '更新时间',
  `version` bigint NOT NULL COMMENT '版本号',
  PRIMARY KEY (`id`)
) ENGINE=InnoDB AUTO_INCREMENT=12 DEFAULT CHARSET=utf8mb4 COLLATE=utf8mb4_general_ci;

INSERT INTO `monitor`.`mon_menus`(`id`, `parent_id`, `menu_name`, `icons`, `url`, `front_url`, `method`, `created_at`, `updated_at`, `version`) VALUES (1, '0', '首页', 'HomeFilled', NULL, '/monitor/board', '', '2022-08-30 09:44:08', '2022-09-02 17:20:16', 1);
INSERT INTO `monitor`.`mon_menus`(`id`, `parent_id`, `menu_name`, `icons`, `url`, `front_url`, `method`, `created_at`, `updated_at`, `version`) VALUES (2, '0', '账号管理', 'User', NULL, NULL, '', '2022-08-30 09:44:41', '2022-09-02 17:11:09', 1);
INSERT INTO `monitor`.`mon_menus`(`id`, `parent_id`, `menu_name`, `icons`, `url`, `front_url`, `method`, `created_at`, `updated_at`, `version`) VALUES (3, '2', '管理员列表', NULL, NULL, '/monitor/admin/list', '', '2022-08-30 09:45:07', '2022-09-06 11:49:43', 1);
INSERT INTO `monitor`.`mon_menus`(`id`, `parent_id`, `menu_name`, `icons`, `url`, `front_url`, `method`, `created_at`, `updated_at`, `version`) VALUES (4, '0', '服务检测', 'Monitor', NULL, NULL, '', '2022-09-02 17:04:08', '2022-09-02 17:07:15', 1);
INSERT INTO `monitor`.`mon_menus`(`id`, `parent_id`, `menu_name`, `icons`, `url`, `front_url`, `method`, `created_at`, `updated_at`, `version`) VALUES (5, '4', '服务检测列表', NULL, NULL, '/monitor/serve/list', '', '2022-09-02 17:04:26', '2022-09-02 17:20:04', 1);
INSERT INTO `monitor`.`mon_menus`(`id`, `parent_id`, `menu_name`, `icons`, `url`, `front_url`, `method`, `created_at`, `updated_at`, `version`) VALUES (6, '0', '机器', 'Monitor', NULL, NULL, '', '2022-09-05 10:56:19', '2022-09-05 10:56:21', 1);
INSERT INTO `monitor`.`mon_menus`(`id`, `parent_id`, `menu_name`, `icons`, `url`, `front_url`, `method`, `created_at`, `updated_at`, `version`) VALUES (7, '6', '客户端机器', NULL, NULL, '/monitor/machine/list', '', '2022-09-05 10:58:48', '2022-09-05 10:58:58', 1);
INSERT INTO `monitor`.`mon_menus`(`id`, `parent_id`, `menu_name`, `icons`, `url`, `front_url`, `method`, `created_at`, `updated_at`, `version`) VALUES (8, '1', '监控面板', NULL, NULL, '/monitor/board', '', '2022-09-05 13:23:47', '2022-09-05 13:23:47', 1);
INSERT INTO `monitor`.`mon_menus`(`id`, `parent_id`, `menu_name`, `icons`, `url`, `front_url`, `method`, `created_at`, `updated_at`, `version`) VALUES (9, '0', '操作日志', 'Reading', NULL, NULL, '', '2022-09-07 13:15:51', '2022-09-07 13:15:54', 1);
INSERT INTO `monitor`.`mon_menus`(`id`, `parent_id`, `menu_name`, `icons`, `url`, `front_url`, `method`, `created_at`, `updated_at`, `version`) VALUES (11, '9', '操作日志', NULL, NULL, '/monitor/log/list', '', '2022-09-07 13:16:53', '2022-09-07 13:16:55', 1);


-- ----------------------------
-- Table structure for mon_operate_record
-- ----------------------------
DROP TABLE IF EXISTS `mon_operate_record`;
CREATE TABLE `mon_operate_record` (
  `id` bigint NOT NULL AUTO_INCREMENT COMMENT '主键id',
  `admin_id` bigint NOT NULL COMMENT '管理id',
  `admin_username` varchar(255) COLLATE utf8mb4_general_ci NOT NULL COMMENT '管理员名字',
  `admin_real_name` varchar(255) COLLATE utf8mb4_general_ci NOT NULL COMMENT '管理员真实名字',
  `url` varchar(255) COLLATE utf8mb4_general_ci NOT NULL COMMENT 'api路径',
  `method` varchar(255) COLLATE utf8mb4_general_ci NOT NULL COMMENT '方法',
  `content` text COLLATE utf8mb4_general_ci NOT NULL COMMENT '操作内容',
  `created_at` timestamp NOT NULL DEFAULT CURRENT_TIMESTAMP COMMENT '操作时间',
  `updated_at` timestamp NOT NULL DEFAULT CURRENT_TIMESTAMP ON UPDATE CURRENT_TIMESTAMP COMMENT '更新时间',
  `version` bigint NOT NULL COMMENT '版本号',
  PRIMARY KEY (`id`)
) ENGINE=InnoDB AUTO_INCREMENT=7054 DEFAULT CHARSET=utf8mb4 COLLATE=utf8mb4_general_ci;

-- ----------------------------
-- Table structure for mon_serve
-- ----------------------------
DROP TABLE IF EXISTS `mon_serve`;
CREATE TABLE `mon_serve` (
  `id` int NOT NULL AUTO_INCREMENT COMMENT '服务检测主键id',
  `serve_name` varchar(255) COLLATE utf8mb4_general_ci NOT NULL COMMENT '服务名称',
  `serve_address` varchar(255) COLLATE utf8mb4_general_ci NOT NULL COMMENT '服务地址',
  `last_check_time` timestamp NOT NULL COMMENT '上次检测时间',
  `serve_state` int NOT NULL COMMENT '服务状态 1: 正常  2: 异常',
  `created_at` timestamp NOT NULL DEFAULT CURRENT_TIMESTAMP COMMENT '服务创建时间',
  `updated_at` timestamp NOT NULL DEFAULT CURRENT_TIMESTAMP ON UPDATE CURRENT_TIMESTAMP COMMENT '服务更新时间',
  `version` bigint NOT NULL COMMENT '版本号',
  PRIMARY KEY (`id`)
) ENGINE=InnoDB AUTO_INCREMENT=24 DEFAULT CHARSET=utf8mb4 COLLATE=utf8mb4_general_ci;

-- ----------------------------
-- Table structure for mon_serve_check_record
-- ----------------------------
DROP TABLE IF EXISTS `mon_serve_check_record`;
CREATE TABLE `mon_serve_check_record` (
  `id` bigint NOT NULL AUTO_INCREMENT COMMENT '主键id',
  `serve_id` bigint NOT NULL COMMENT '服务id',
  `serve_name` varchar(255) COLLATE utf8mb4_general_ci NOT NULL COMMENT '服务名称',
  `serve_address` varchar(255) COLLATE utf8mb4_general_ci NOT NULL COMMENT '服务地址',
  `last_check_time` timestamp NOT NULL COMMENT '上次检测时间',
  `created_at` timestamp NOT NULL COMMENT '创建时间',
  `updated_at` timestamp NOT NULL COMMENT '更新时间',
  `version` bigint NOT NULL COMMENT '版本号',
  PRIMARY KEY (`id`)
) ENGINE=InnoDB AUTO_INCREMENT=59 DEFAULT CHARSET=utf8mb4 COLLATE=utf8mb4_general_ci;

-- ----------------------------
-- Table structure for mon_upgrade_machine_record
-- ----------------------------
DROP TABLE IF EXISTS `mon_upgrade_machine_record`;
CREATE TABLE `mon_upgrade_machine_record` (
  `id` bigint NOT NULL AUTO_INCREMENT COMMENT '主键id',
  `machine_code` varchar(255) COLLATE utf8mb4_general_ci NOT NULL COMMENT '机器码',
  `machine_ip` varchar(255) COLLATE utf8mb4_general_ci NOT NULL COMMENT '机器ip',
  `machine_hostname` varchar(255) COLLATE utf8mb4_general_ci NOT NULL COMMENT '机器主机名',
  `machine_remark` varchar(255) COLLATE utf8mb4_general_ci NOT NULL COMMENT '机器备注',
  `package_name` varchar(255) COLLATE utf8mb4_general_ci NOT NULL COMMENT '客户端包名',
  `upgrade_version` varchar(255) COLLATE utf8mb4_general_ci NOT NULL COMMENT '版本号',
  `md5_sum` varchar(255) CHARACTER SET utf8mb4 COLLATE utf8mb4_general_ci DEFAULT NULL COMMENT '文件md5',
  `created_at` timestamp NOT NULL DEFAULT CURRENT_TIMESTAMP COMMENT '创建时间',
  `updated_at` timestamp NOT NULL COMMENT '更新时间',
  `version` bigint NOT NULL COMMENT '版本号',
  PRIMARY KEY (`id`)
) ENGINE=InnoDB AUTO_INCREMENT=62 DEFAULT CHARSET=utf8mb4 COLLATE=utf8mb4_general_ci;

-- ----------------------------
-- Table structure for mon_upgrade_serve_record
-- ----------------------------
DROP TABLE IF EXISTS `mon_upgrade_serve_record`;
CREATE TABLE `mon_upgrade_serve_record` (
  `id` bigint NOT NULL AUTO_INCREMENT COMMENT '主键id',
  `serve_id` bigint NOT NULL COMMENT '服务id',
  `serve_name` varchar(255) COLLATE utf8mb4_general_ci NOT NULL COMMENT '服务名字',
  `serve_address` varchar(255) COLLATE utf8mb4_general_ci NOT NULL COMMENT '服务地址',
  `serve_created_at` timestamp NULL DEFAULT NULL COMMENT '服务创建时间',
  `serve_state` int NOT NULL COMMENT '服务状态',
  `package_name` varchar(255) COLLATE utf8mb4_general_ci NOT NULL COMMENT '包名',
  `upgrade_version` varchar(255) CHARACTER SET utf8mb4 COLLATE utf8mb4_general_ci NOT NULL COMMENT '服务版本号',
  `created_at` timestamp NOT NULL DEFAULT CURRENT_TIMESTAMP COMMENT '升级时间',
  `updated_at` timestamp NOT NULL DEFAULT CURRENT_TIMESTAMP COMMENT '更新时间',
  `version` bigint NOT NULL COMMENT '版本号',
  PRIMARY KEY (`id`)
) ENGINE=InnoDB AUTO_INCREMENT=6 DEFAULT CHARSET=utf8mb4 COLLATE=utf8mb4_general_ci;

SET FOREIGN_KEY_CHECKS = 1;
