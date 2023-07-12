/*
 Navicat Premium Data Transfer

 Source Server         : mysql
 Source Server Type    : MySQL
 Source Server Version : 80033
 Source Host           : 60.204.144.40:13306
 Source Schema         : imlogic

 Target Server Type    : MySQL
 Target Server Version : 80033
 File Encoding         : 65001

 Date: 18/06/2023 02:34:33
*/

SET NAMES utf8mb4;
SET
FOREIGN_KEY_CHECKS = 0;

-- ----------------------------
-- Table structure for user
-- ----------------------------
DROP TABLE IF EXISTS `user`;
CREATE TABLE `user`
(
    `id`           bigint unsigned NOT NULL AUTO_INCREMENT COMMENT '主键',
    `created_time` bigint                                                 NOT NULL DEFAULT '0' COMMENT '创建时间',
    `updated_time` bigint                                                 NOT NULL DEFAULT '0' COMMENT '更新时间',
    `delete_time`  bigint                                                 NOT NULL DEFAULT '0' COMMENT '删除时间',
    `del_stale`    tinyint                                                NOT NULL DEFAULT '0' COMMENT '删除状态',
    `username`     varchar(255) CHARACTER SET utf8mb4 COLLATE utf8mb4_bin          DEFAULT 'user' COMMENT '用户名',
    `password`     varchar(20) CHARACTER SET utf8mb4 COLLATE utf8mb4_bin           DEFAULT '123456' COMMENT '密码',
    `mobile`       varchar(255) CHARACTER SET utf8mb4 COLLATE utf8mb4_bin NOT NULL COMMENT '手机号',
    `avatar`       varchar(255) CHARACTER SET utf8mb4 COLLATE utf8mb4_bin NOT NULL DEFAULT 'https://file.xu756.top/avatar.png' COMMENT '头像',
    `role`         bigint                                                 NOT NULL DEFAULT '1' COMMENT '角色',
    `status`       bigint                                                 NOT NULL DEFAULT '1' COMMENT '状态',
    `open_id`      varchar(50) CHARACTER SET utf8mb4 COLLATE utf8mb4_bin  NOT NULL COMMENT '用户唯一标志',
    `balance`      float                                                  NOT NULL DEFAULT '0' COMMENT '余额',
    `use_count`    int                                                    NOT NULL DEFAULT '0' COMMENT '使用次数',
    PRIMARY KEY (`id`) USING BTREE,
    UNIQUE KEY `mobile` (`mobile`),
    UNIQUE KEY `open_id` (`open_id`)
) ENGINE=InnoDB AUTO_INCREMENT=28 DEFAULT CHARSET=utf8mb4 COLLATE=utf8mb4_bin;

SET FOREIGN_KEY_CHECKS = 1;
