/*
 Navicat Premium Data Transfer

 Source Server         : 本地开发数据库
 Source Server Type    : MySQL
 Source Server Version : 80033 (8.0.33)
 Source Host           : localhost:13306
 Source Schema         : imlogic

 Target Server Type    : MySQL
 Target Server Version : 80033 (8.0.33)
 File Encoding         : 65001

 Date: 16/07/2023 11:00:02
*/

SET NAMES utf8mb4;
SET
FOREIGN_KEY_CHECKS = 0;

-- ----------------------------
-- Table structure for account
-- ----------------------------
DROP TABLE IF EXISTS `account`;
CREATE TABLE `account`
(
    `id`        bigint                                                        NOT NULL AUTO_INCREMENT COMMENT '账号ID',
    `user_id`   bigint                                                        NOT NULL COMMENT '用户ID',
    `open_code` varchar(255) CHARACTER SET utf8mb3 COLLATE utf8mb3_general_ci NOT NULL COMMENT '登录账号,如手机号等',
    `category`  tinyint(1) DEFAULT NULL COMMENT '账号类别',
    `created`   bigint                                                        NOT NULL COMMENT '创建时间',
    `creator`   bigint                                                        NOT NULL COMMENT '创建人',
    `edited`    bigint                                                        NOT NULL COMMENT '修改时间',
    `editor`    bigint                                                        NOT NULL COMMENT '修改人',
    `deleted`   double(1, 0
) unsigned zerofill NOT NULL COMMENT '逻辑删除:0=未删除,1=已删除',
  PRIMARY KEY (`id`) USING BTREE,
  KEY `idx_member_id` (`user_id`) USING BTREE COMMENT '普通索引'
) ENGINE=InnoDB DEFAULT CHARSET=utf8mb4 COLLATE=utf8mb4_general_ci ROW_FORMAT=DYNAMIC COMMENT='账号';

-- ----------------------------
-- Table structure for role
-- ----------------------------
DROP TABLE IF EXISTS `role`;
CREATE TABLE `role`
(
    `id`        bigint                                                        NOT NULL AUTO_INCREMENT COMMENT '角色ID',
    `parent_id` bigint                                                        NOT NULL COMMENT '所属父级角色ID',
    `code`      varchar(255) CHARACTER SET utf8mb3 COLLATE utf8mb3_general_ci NOT NULL COMMENT '角色唯一CODE代码',
    `name`      varchar(255) CHARACTER SET utf8mb3 COLLATE utf8mb3_general_ci NOT NULL COMMENT '角色名称',
    `intro`     varchar(255) CHARACTER SET utf8mb3 COLLATE utf8mb3_general_ci NOT NULL COMMENT '角色介绍',
    `created`   bigint                                                        NOT NULL COMMENT '创建时间',
    `creator`   bigint                                                        NOT NULL COMMENT '创建人',
    `edited`    bigint                                                        NOT NULL COMMENT '修改时间',
    `editor`    bigint                                                        NOT NULL COMMENT '修改人',
    `deleted`   tinyint(1) unsigned zerofill NOT NULL COMMENT '逻辑删除:0=未删除,1=已删除',
    PRIMARY KEY (`id`) USING BTREE,
    KEY         `parent_id` (`parent_id`) USING BTREE COMMENT '父级权限ID',
    KEY         `code` (`code`) USING BTREE COMMENT '权限CODE代码'
) ENGINE=InnoDB DEFAULT CHARSET=utf8mb3 ROW_FORMAT=DYNAMIC COMMENT='角色';

-- ----------------------------
-- Table structure for user
-- ----------------------------
DROP TABLE IF EXISTS `user`;
CREATE TABLE `user`
(
    `id`           bigint                                                        NOT NULL AUTO_INCREMENT COMMENT '用户ID',
    `state`        tinyint(1) NOT NULL COMMENT '用户状态:0=正常,1=禁用',
    `name`         varchar(255) CHARACTER SET utf8mb3 COLLATE utf8mb3_general_ci NOT NULL COMMENT '姓名',
    `head_img_url` varchar(255) CHARACTER SET utf8mb3 COLLATE utf8mb3_general_ci NOT NULL COMMENT '头像图片地址',
    `mobile`       varchar(11) CHARACTER SET utf8mb3 COLLATE utf8mb3_general_ci  NOT NULL COMMENT '手机号码',
    `salt`         varchar(64) CHARACTER SET utf8mb3 COLLATE utf8mb3_general_ci  NOT NULL COMMENT '密码加盐',
    `password`     varchar(64) CHARACTER SET utf8mb3 COLLATE utf8mb3_general_ci  NOT NULL COMMENT '登录密码',
    `created`      bigint                                                        NOT NULL COMMENT '创建时间',
    `creator`      bigint                                                        NOT NULL COMMENT '创建人',
    `edited`       bigint                                                        NOT NULL COMMENT '修改时间',
    `editor`       bigint                                                        NOT NULL COMMENT '修改人',
    `deleted`      tinyint(1) unsigned zerofill NOT NULL COMMENT '逻辑删除:0=未删除,1=已删除',
    PRIMARY KEY (`id`) USING BTREE
) ENGINE=InnoDB DEFAULT CHARSET=utf8mb3 ROW_FORMAT=DYNAMIC COMMENT='用户';

-- ----------------------------
-- Table structure for user_group
-- ----------------------------
DROP TABLE IF EXISTS `user_group`;
CREATE TABLE `user_group`
(
    `id`        bigint                                                        NOT NULL AUTO_INCREMENT COMMENT 'ID',
    `parent_id` bigint                                                        NOT NULL COMMENT '所属父级用户组ID',
    `name`      varchar(255) CHARACTER SET utf8mb3 COLLATE utf8mb3_general_ci NOT NULL COMMENT '用户组名称',
    `code`      varchar(255) CHARACTER SET utf8mb3 COLLATE utf8mb3_general_ci NOT NULL COMMENT '用户组CODE唯一代码',
    `intro`     varchar(255) CHARACTER SET utf8mb3 COLLATE utf8mb3_general_ci NOT NULL COMMENT '用户组介绍',
    `created`   bigint                                                        NOT NULL COMMENT '创建时间',
    `creator`   bigint                                                        NOT NULL COMMENT '创建人',
    `edited`    bigint                                                        NOT NULL COMMENT '修改时间',
    `editor`    bigint                                                        NOT NULL COMMENT '修改人',
    `deleted`   tinyint(1) unsigned zerofill NOT NULL COMMENT '逻辑删除:0=未删除,1=已删除',
    PRIMARY KEY (`id`) USING BTREE,
    KEY         `parent_id` (`parent_id`) USING BTREE COMMENT '父级用户组ID',
    KEY         `code` (`code`) USING BTREE COMMENT '用户组CODE代码'
) ENGINE=InnoDB DEFAULT CHARSET=utf8mb3 ROW_FORMAT=DYNAMIC COMMENT='用户组';

-- ----------------------------
-- Table structure for user_group_user
-- ----------------------------
DROP TABLE IF EXISTS `user_group_user`;
CREATE TABLE `user_group_user`
(
    `id`            bigint NOT NULL AUTO_INCREMENT COMMENT 'ID说',
    `user_group_id` bigint NOT NULL COMMENT '用户组ID',
    `user_id`       bigint NOT NULL COMMENT '用户ID',
    `created`       bigint NOT NULL COMMENT '创建时间',
    `creator`       bigint NOT NULL COMMENT '创建人',
    `edited`        bigint NOT NULL COMMENT '修改时间',
    `editor`        bigint NOT NULL COMMENT '修改人',
    `deleted`       bigint unsigned zerofill NOT NULL COMMENT '逻辑删除:0=未删除,1=已删除',
    PRIMARY KEY (`id`) USING BTREE,
    KEY             `member_group_id` (`user_group_id`) USING BTREE COMMENT '用户组ID',
    KEY             `member_id` (`user_id`) USING BTREE COMMENT '用户ID'
) ENGINE=InnoDB DEFAULT CHARSET=utf8mb3 ROW_FORMAT=DYNAMIC COMMENT='用户组成员';

-- ----------------------------
-- Table structure for user_role
-- ----------------------------
DROP TABLE IF EXISTS `user_role`;
CREATE TABLE `user_role`
(
    `id`      bigint NOT NULL AUTO_INCREMENT COMMENT 'ID',
    `user_id` bigint NOT NULL COMMENT '用户ID',
    `role_id` bigint NOT NULL COMMENT '角色ID',
    `created` bigint NOT NULL COMMENT '创建时间',
    `creator` bigint NOT NULL COMMENT '创建人',
    `edited`  bigint NOT NULL COMMENT '修改时间',
    `editor`  bigint NOT NULL COMMENT '修改人',
    `deleted` tinyint(1) unsigned zerofill NOT NULL COMMENT '逻辑删除:0=未删除,1=已删除',
    PRIMARY KEY (`id`) USING BTREE,
    KEY       `member_id` (`user_id`) USING BTREE COMMENT '用户ID',
    KEY       `role_id` (`role_id`) USING BTREE COMMENT '角色ID'
) ENGINE=InnoDB DEFAULT CHARSET=utf8mb3 ROW_FORMAT=DYNAMIC COMMENT='用户角色';

SET
FOREIGN_KEY_CHECKS = 1;
