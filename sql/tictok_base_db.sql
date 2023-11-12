/*
 Navicat MySQL Data Transfer

 Source Server         : root
 Source Server Type    : MySQL
 Source Server Version : 80032
 Source Host           : localhost:3306
 Source Schema         : tictok_base_db

 Target Server Type    : MySQL
 Target Server Version : 80032
 File Encoding         : 65001

 Date: 30/10/2023 22:16:10
*/

SET NAMES utf8mb4;
SET FOREIGN_KEY_CHECKS = 0;

-- ----------------------------
-- Table structure for comments
-- ----------------------------
DROP TABLE IF EXISTS `comments`;
CREATE TABLE `comments`  (
  `id` bigint(0) NOT NULL AUTO_INCREMENT,
  `video_id` bigint(0) NULL DEFAULT NULL,
  `user_id` bigint(0) NULL DEFAULT NULL,
  `content` varchar(255) CHARACTER SET utf8mb4 COLLATE utf8mb4_0900_ai_ci NULL DEFAULT NULL,
  `create_date` datetime(0) NULL DEFAULT NULL,
  PRIMARY KEY (`id`) USING BTREE
) ENGINE = InnoDB AUTO_INCREMENT = 1 CHARACTER SET = utf8mb4 COLLATE = utf8mb4_0900_ai_ci ROW_FORMAT = Dynamic;

-- ----------------------------
-- Records of comments
-- ----------------------------
INSERT INTO `comments` VALUES (1, 1, 2, '我好喜欢！', '2023-09-02 20:10:43');

-- ----------------------------
-- Table structure for favorite_actions
-- ----------------------------
DROP TABLE IF EXISTS `favorite_actions`;
CREATE TABLE `favorite_actions`  (
  `id` bigint(0) NOT NULL AUTO_INCREMENT,
  `video_id` bigint(0) NULL DEFAULT NULL,
  `user_id` bigint(0) NULL DEFAULT NULL,
  PRIMARY KEY (`id`) USING BTREE
) ENGINE = InnoDB AUTO_INCREMENT = 1 CHARACTER SET = utf8mb4 COLLATE = utf8mb4_0900_ai_ci ROW_FORMAT = Dynamic;

-- ----------------------------
-- Records of favorite_actions
-- ----------------------------
INSERT INTO `favorite_actions` VALUES (1, 1, 2);

-- ----------------------------
-- Table structure for messages
-- ----------------------------
DROP TABLE IF EXISTS `messages`;
CREATE TABLE `messages`  (
  `id` bigint(0) NOT NULL AUTO_INCREMENT,
  `to_user_id` bigint(0) NULL DEFAULT NULL,
  `from_user_id` bigint(0) NULL DEFAULT NULL,
  `content` varchar(255) CHARACTER SET utf8mb4 COLLATE utf8mb4_0900_ai_ci NULL DEFAULT NULL,
  `create_time` bigint(0) NULL DEFAULT NULL,
  PRIMARY KEY (`id`) USING BTREE
) ENGINE = InnoDB AUTO_INCREMENT = 1 CHARACTER SET = utf8mb4 COLLATE = utf8mb4_0900_ai_ci ROW_FORMAT = Dynamic;

-- ----------------------------
-- Records of messages
-- ----------------------------
INSERT INTO `messages` VALUES (1, 1, 2, '在干嘛？', 1693670761);

-- ----------------------------
-- Table structure for relation_actions
-- ----------------------------
DROP TABLE IF EXISTS `relation_actions`;
CREATE TABLE `relation_actions`  (
  `id` bigint(0) NOT NULL AUTO_INCREMENT,
  `user_id` bigint(0) NULL DEFAULT NULL,
  `to_user_id` bigint(0) NULL DEFAULT NULL,
  PRIMARY KEY (`id`) USING BTREE
) ENGINE = InnoDB AUTO_INCREMENT = 1 CHARACTER SET = utf8mb4 COLLATE = utf8mb4_0900_ai_ci ROW_FORMAT = Dynamic;

-- ----------------------------
-- Records of relation_actions
-- ----------------------------
INSERT INTO `relation_actions` VALUES (1, 2, 1);

-- ----------------------------
-- Table structure for users
-- ----------------------------
DROP TABLE IF EXISTS `users`;
CREATE TABLE `users`  (
  `id` bigint(0) NOT NULL AUTO_INCREMENT,
  `name` varchar(255) CHARACTER SET utf8mb4 COLLATE utf8mb4_0900_ai_ci NULL DEFAULT NULL,
  `password` varchar(255) CHARACTER SET utf8mb4 COLLATE utf8mb4_0900_ai_ci NULL DEFAULT NULL,
  `follow_count` bigint(0) NULL DEFAULT NULL,
  `follower_count` bigint(0) NULL DEFAULT NULL,
  `avatar` varchar(255) CHARACTER SET utf8mb4 COLLATE utf8mb4_0900_ai_ci NULL DEFAULT NULL,
  `background_image` varchar(255) CHARACTER SET utf8mb4 COLLATE utf8mb4_0900_ai_ci NULL DEFAULT NULL,
  `signature` varchar(255) CHARACTER SET utf8mb4 COLLATE utf8mb4_0900_ai_ci NULL DEFAULT NULL,
  `total_favorited` bigint(0) NULL DEFAULT NULL,
  `work_count` bigint(0) NULL DEFAULT NULL,
  `favorite_count` bigint(0) NULL DEFAULT NULL,
  PRIMARY KEY (`id`) USING BTREE
) ENGINE = InnoDB AUTO_INCREMENT = 4 CHARACTER SET = utf8mb4 COLLATE = utf8mb4_0900_ai_ci ROW_FORMAT = Dynamic;

-- ----------------------------
-- Records of users
-- ----------------------------
INSERT INTO `users` VALUES (1, 'helianthus', 'a12345', 0, 0, '/image/helianthus.jpg', '/image/background.jpg', 'Hello, I\'m Helianthus!', 0, 0, 0);
INSERT INTO `users` VALUES (2, 'cyo', 'cyo520', 0, 0, '/image/default.jpg', '/image/background.jpg', '', 0, 0, 0);
INSERT INTO `users` VALUES (3, 'gyci', 'gyci666', 0, 0, '/image/default.jpg', '/image/background.jpg', '', 0, 0, 0);
INSERT INTO `users` VALUES (4, 'wxs', 'a123456', 0, 0, '/image/default.jpg', '/image/background.jpg', '', 0, 0, 0);

-- ----------------------------
-- Table structure for videos
-- ----------------------------
DROP TABLE IF EXISTS `videos`;
CREATE TABLE `videos`  (
  `id` bigint(0) NOT NULL AUTO_INCREMENT,
  `author_id` bigint(0) NULL DEFAULT NULL,
  `play_url` varchar(255) CHARACTER SET utf8mb4 COLLATE utf8mb4_0900_ai_ci NULL DEFAULT NULL,
  `cover_url` varchar(255) CHARACTER SET utf8mb4 COLLATE utf8mb4_0900_ai_ci NULL DEFAULT NULL,
  `favorite_count` bigint(0) NULL DEFAULT NULL,
  `comment_count` bigint(0) NULL DEFAULT NULL,
  `title` varchar(255) CHARACTER SET utf8mb4 COLLATE utf8mb4_0900_ai_ci NULL DEFAULT NULL,
  `create_date` bigint(0) NULL DEFAULT NULL,
  PRIMARY KEY (`id`) USING BTREE
) ENGINE = InnoDB AUTO_INCREMENT = 13 CHARACTER SET = utf8mb4 COLLATE = utf8mb4_0900_ai_ci ROW_FORMAT = Dynamic;

-- ----------------------------
-- Records of videos
-- ----------------------------
INSERT INTO `videos` VALUES (1, 1, '/video/helianthus.mp4', '/image/helianthus.jpg', 0, 0, 'Helianthus\'s First vlog ！', 1693721358);
INSERT INTO `videos` VALUES (2, 1, '/video/heliantus_9_3_1533.mp4', '/image/heliantus_9_3_1533.jpg', 0, 0, 'Helianthus\'s Second vlog ！', 1693726558);
INSERT INTO `videos` VALUES (12, 2, '/video/2_share_267c185986d9b7afb95f6eb8a96a8157_1695837457.mp4', '/image/default.jpg', 0, 0, '橙橙跳舞', 1695837457);
INSERT INTO `videos` VALUES (13, 2, '/video/2_share_a17d5001853675a9ab5b6b381ccb439f_1695907553.mp4', '/image/default.jpg', 0, 0, '橙橙国庆跳舞💃🏻', 1695907553);

SET FOREIGN_KEY_CHECKS = 1;
