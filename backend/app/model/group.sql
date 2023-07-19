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

 Date: 19/07/2023 21:46:26
*/

SET NAMES utf8mb4;
SET FOREIGN_KEY_CHECKS = 0;

-- ----------------------------
-- Table structure for group
-- ----------------------------
DROP TABLE IF EXISTS `group`;
CREATE TABLE `group`
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
    `deleted`   tinyint(1) unsigned zerofill                                  NOT NULL COMMENT '逻辑删除:0=未删除,1=已删除',
    PRIMARY KEY (`id`) USING BTREE,
    KEY `parent_id` (`parent_id`) USING BTREE COMMENT '父级用户组ID',
    KEY `code` (`code`) USING BTREE COMMENT '用户组CODE代码'
) ENGINE = InnoDB
  AUTO_INCREMENT = 2
  DEFAULT CHARSET = utf8mb3
  ROW_FORMAT = DYNAMIC COMMENT ='用户组';

SET FOREIGN_KEY_CHECKS = 1;
