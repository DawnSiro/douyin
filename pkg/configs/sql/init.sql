/*
 Navicat Premium Data Transfer

 Source Server         : localhost_3306
 Source Server Type    : MySQL
 Source Server Version : 50732
 Source Host           : localhost:3306
 Source Schema         : douyin

 Target Server Type    : MySQL
 Target Server Version : 50732
 File Encoding         : 65001

 Date: 10/02/2023 15:15:52
*/

SET NAMES utf8mb4;
SET FOREIGN_KEY_CHECKS = 0;

-- ----------------------------
-- Table structure for comment
-- ----------------------------
DROP TABLE IF EXISTS `comment`;
CREATE TABLE `comment`  (
                            `id` bigint(20) UNSIGNED NOT NULL AUTO_INCREMENT,
                            `created_at` datetime(3) NULL DEFAULT NULL,
                            `updated_at` datetime(3) NULL DEFAULT NULL,
                            `deleted_at` datetime(3) NULL DEFAULT NULL,
                            `video_id` bigint(20) UNSIGNED NOT NULL,
                            `user_id` bigint(20) UNSIGNED NOT NULL,
                            `content` varchar(255) CHARACTER SET utf8mb4 COLLATE utf8mb4_general_ci NOT NULL,
                            PRIMARY KEY (`id`) USING BTREE,
                            INDEX `idx_comment_deleted_at`(`deleted_at`) USING BTREE,
                            INDEX `idx_videoID`(`video_id`) USING BTREE,
                            INDEX `idx_userid`(`user_id`) USING BTREE,
                            CONSTRAINT `fk_comment_user` FOREIGN KEY (`user_id`) REFERENCES `user` (`id`) ON DELETE RESTRICT ON UPDATE RESTRICT,
                            CONSTRAINT `fk_comment_video` FOREIGN KEY (`video_id`) REFERENCES `video` (`id`) ON DELETE RESTRICT ON UPDATE RESTRICT
) ENGINE = InnoDB AUTO_INCREMENT = 1 CHARACTER SET = utf8mb4 COLLATE = utf8mb4_general_ci ROW_FORMAT = Dynamic;

-- ----------------------------
-- Table structure for message
-- ----------------------------
DROP TABLE IF EXISTS `message`;
CREATE TABLE `message`  (
                            `id` bigint(20) UNSIGNED NOT NULL AUTO_INCREMENT,
                            `created_at` datetime(3) NULL DEFAULT NULL,
                            `updated_at` datetime(3) NULL DEFAULT NULL,
                            `deleted_at` datetime(3) NULL DEFAULT NULL,
                            `to_user_id` bigint(20) UNSIGNED NOT NULL,
                            `from_user_id` bigint(20) UNSIGNED NOT NULL,
                            `content` varchar(255) CHARACTER SET utf8mb4 COLLATE utf8mb4_general_ci NOT NULL,
                            PRIMARY KEY (`id`) USING BTREE,
                            UNIQUE INDEX `idx_userid`(`to_user_id`, `from_user_id`) USING BTREE,
                            INDEX `idx_userid_to`(`to_user_id`) USING BTREE,
                            INDEX `idx_userid_from`(`from_user_id`) USING BTREE,
                            INDEX `idx_message_deleted_at`(`deleted_at`) USING BTREE,
                            CONSTRAINT `fk_message_from_user` FOREIGN KEY (`from_user_id`) REFERENCES `user` (`id`) ON DELETE RESTRICT ON UPDATE RESTRICT,
                            CONSTRAINT `fk_message_to_user` FOREIGN KEY (`to_user_id`) REFERENCES `user` (`id`) ON DELETE RESTRICT ON UPDATE RESTRICT
) ENGINE = InnoDB AUTO_INCREMENT = 1 CHARACTER SET = utf8mb4 COLLATE = utf8mb4_general_ci ROW_FORMAT = Dynamic;

-- ----------------------------
-- Table structure for relation
-- ----------------------------
DROP TABLE IF EXISTS `relation`;
CREATE TABLE `relation`  (
                             `id` bigint(20) UNSIGNED NOT NULL AUTO_INCREMENT,
                             `created_at` datetime(3) NULL DEFAULT NULL,
                             `updated_at` datetime(3) NULL DEFAULT NULL,
                             `deleted_at` datetime(3) NULL DEFAULT NULL,
                             `user_id` bigint(20) UNSIGNED NOT NULL,
                             `to_user_id` bigint(20) UNSIGNED NOT NULL,
                             PRIMARY KEY (`id`) USING BTREE,
                             UNIQUE INDEX `idx_userid`(`user_id`, `to_user_id`) USING BTREE,
                             INDEX `idx_relation_deleted_at`(`deleted_at`) USING BTREE,
                             INDEX `idx_userid_to`(`to_user_id`) USING BTREE,
                             CONSTRAINT `fk_relation_to_user` FOREIGN KEY (`to_user_id`) REFERENCES `user` (`id`) ON DELETE RESTRICT ON UPDATE RESTRICT,
                             CONSTRAINT `fk_relation_user` FOREIGN KEY (`user_id`) REFERENCES `user` (`id`) ON DELETE RESTRICT ON UPDATE RESTRICT
) ENGINE = InnoDB AUTO_INCREMENT = 1 CHARACTER SET = utf8mb4 COLLATE = utf8mb4_general_ci ROW_FORMAT = Dynamic;

-- ----------------------------
-- Table structure for user
-- ----------------------------
DROP TABLE IF EXISTS `user`;
CREATE TABLE `user`  (
                         `id` bigint(20) UNSIGNED NOT NULL AUTO_INCREMENT,
                         `created_at` datetime(3) NULL DEFAULT NULL,
                         `updated_at` datetime(3) NULL DEFAULT NULL,
                         `deleted_at` datetime(3) NULL DEFAULT NULL,
                         `username` varchar(64) CHARACTER SET utf8mb4 COLLATE utf8mb4_general_ci NOT NULL,
                         `password` varchar(256) CHARACTER SET utf8mb4 COLLATE utf8mb4_general_ci NOT NULL,
                         `following_count` bigint(20) NULL DEFAULT 0,
                         `follower_count` bigint(20) NULL DEFAULT 0,
                         `avatar` varchar(256) CHARACTER SET utf8mb4 COLLATE utf8mb4_general_ci NULL DEFAULT 'https://picture-bucket-01.oss-cn-beijing.aliyuncs.com/avatar/default.png',
                         PRIMARY KEY (`id`) USING BTREE,
                         UNIQUE INDEX `idx_username`(`username`) USING BTREE,
                         INDEX `idx_user_deleted_at`(`deleted_at`) USING BTREE
) ENGINE = InnoDB AUTO_INCREMENT = 2 CHARACTER SET = utf8mb4 COLLATE = utf8mb4_general_ci ROW_FORMAT = Dynamic;

-- ----------------------------
-- Table structure for user_favorite_videos
-- ----------------------------
DROP TABLE IF EXISTS `user_favorite_videos`;
CREATE TABLE `user_favorite_videos`  (
                                         `user_id` bigint(20) UNSIGNED NOT NULL,
                                         `video_id` bigint(20) UNSIGNED NOT NULL,
                                         PRIMARY KEY (`user_id`, `video_id`) USING BTREE,
                                         INDEX `fk_user_favorite_videos_video`(`video_id`) USING BTREE,
                                         CONSTRAINT `fk_user_favorite_videos_user` FOREIGN KEY (`user_id`) REFERENCES `user` (`id`) ON DELETE RESTRICT ON UPDATE RESTRICT,
                                         CONSTRAINT `fk_user_favorite_videos_video` FOREIGN KEY (`video_id`) REFERENCES `video` (`id`) ON DELETE RESTRICT ON UPDATE RESTRICT
) ENGINE = InnoDB CHARACTER SET = utf8mb4 COLLATE = utf8mb4_general_ci ROW_FORMAT = Dynamic;

-- ----------------------------
-- Table structure for video
-- ----------------------------
DROP TABLE IF EXISTS `video`;
CREATE TABLE `video`  (
                          `id` bigint(20) UNSIGNED NOT NULL AUTO_INCREMENT,
                          `created_at` datetime(3) NULL DEFAULT NULL,
                          `updated_at` datetime(3) NULL DEFAULT NULL,
                          `deleted_at` datetime(3) NULL DEFAULT NULL,
                          `update_time` datetime(3) NOT NULL,
                          `author_id` bigint(20) UNSIGNED NOT NULL,
                          `play_url` varchar(255) CHARACTER SET utf8mb4 COLLATE utf8mb4_general_ci NOT NULL,
                          `cover_url` varchar(255) CHARACTER SET utf8mb4 COLLATE utf8mb4_general_ci NULL DEFAULT NULL,
                          `favorite_count` bigint(20) NULL DEFAULT 0,
                          `comment_count` bigint(20) NULL DEFAULT 0,
                          `title` varchar(50) CHARACTER SET utf8mb4 COLLATE utf8mb4_general_ci NOT NULL,
                          PRIMARY KEY (`id`) USING BTREE,
                          INDEX `idx_video_deleted_at`(`deleted_at`) USING BTREE,
                          INDEX `idx_update`(`update_time`) USING BTREE,
                          INDEX `idx_authorID`(`author_id`) USING BTREE,
                          CONSTRAINT `fk_video_author` FOREIGN KEY (`author_id`) REFERENCES `user` (`id`) ON DELETE RESTRICT ON UPDATE RESTRICT
) ENGINE = InnoDB AUTO_INCREMENT = 1 CHARACTER SET = utf8mb4 COLLATE = utf8mb4_general_ci ROW_FORMAT = Dynamic;

SET FOREIGN_KEY_CHECKS = 1;
