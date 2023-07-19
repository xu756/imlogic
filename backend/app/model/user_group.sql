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

 Date: 19/07/2023 21:46:47
*/

SET NAMES utf8mb4;
SET FOREIGN_KEY_CHECKS = 0;

-- ----------------------------
-- Table structure for user_group
-- ----------------------------
DROP TABLE IF EXISTS `user_group`;
CREATE TABLE `user_group`
(
    `id`            bigint                       NOT NULL AUTO_INCREMENT COMMENT 'ID说',
    `user_group_id` bigint                       NOT NULL COMMENT '用户组ID',
    `user_id`       bigint                       NOT NULL COMMENT '用户ID',
    `created`       bigint                       NOT NULL COMMENT '创建时间',
    `creator`       bigint                       NOT NULL COMMENT '创建人',
    `edited`        bigint                       NOT NULL COMMENT '修改时间',
    `editor`        bigint                       NOT NULL COMMENT '修改人',
    `deleted`       bigint(20) unsigned zerofill NOT NULL COMMENT '逻辑删除:0=未删除,1=已删除',
    PRIMARY KEY (`id`) USING BTREE,
    KEY `member_group_id` (`user_group_id`) USING BTREE COMMENT '用户组ID',
    KEY `member_id` (`user_id`) USING BTREE COMMENT '用户ID'
) ENGINE = InnoDB
  DEFAULT CHARSET = utf8mb3
  ROW_FORMAT = DYNAMIC COMMENT ='用户组成员';

SET FOREIGN_KEY_CHECKS = 1;
