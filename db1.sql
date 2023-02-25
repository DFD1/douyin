/*
 Navicat Premium Data Transfer

 Source Server         : test
 Source Server Type    : MySQL
 Source Server Version : 50710
 Source Host           : localhost:3306
 Source Schema         : db1

 Target Server Type    : MySQL
 Target Server Version : 50710
 File Encoding         : 65001

 Date: 25/02/2023 16:12:05
*/

SET NAMES utf8mb4;
SET FOREIGN_KEY_CHECKS = 0;

-- ----------------------------
-- Table structure for comments
-- ----------------------------
DROP TABLE IF EXISTS `comments`;
CREATE TABLE `comments`  (
  `id` int(255) NOT NULL AUTO_INCREMENT COMMENT '自增主键',
  `user_id` int(255) NULL DEFAULT NULL COMMENT '评论发布用户id',
  `video_id` int(255) NULL DEFAULT NULL COMMENT '评论视频id',
  `comment_text` varchar(255) CHARACTER SET utf8 COLLATE utf8_general_ci NULL DEFAULT NULL COMMENT '评论内容',
  `create_date` datetime(0) NULL DEFAULT NULL COMMENT '评论发布时间',
  `cancel` int(10) NULL DEFAULT NULL COMMENT '正常为0，取消后为1',
  PRIMARY KEY (`id`) USING BTREE
) ENGINE = InnoDB AUTO_INCREMENT = 7 CHARACTER SET = utf8 COLLATE = utf8_general_ci ROW_FORMAT = Dynamic;

-- ----------------------------
-- Records of comments
-- ----------------------------
INSERT INTO `comments` VALUES (2, 1, 1, '1111', '2023-02-23 16:44:32', 0);
INSERT INTO `comments` VALUES (4, 1, 14, 'ddd', '2023-02-23 16:52:42', 0);
INSERT INTO `comments` VALUES (5, 1, 14, '324324', '2023-02-23 17:18:04', 0);
INSERT INTO `comments` VALUES (6, 2, 14, '3432432', '2023-02-23 17:21:35', 0);

-- ----------------------------
-- Table structure for likes
-- ----------------------------
DROP TABLE IF EXISTS `likes`;
CREATE TABLE `likes`  (
  `id` int(255) NOT NULL AUTO_INCREMENT COMMENT '自增主键',
  `user_id` int(255) NULL DEFAULT NULL COMMENT '点赞用户的id',
  `video_id` int(255) NULL DEFAULT NULL COMMENT '被点赞视频的id',
  `cancel` int(10) NULL DEFAULT NULL COMMENT '默认为0，取消后为1',
  PRIMARY KEY (`id`) USING BTREE
) ENGINE = InnoDB AUTO_INCREMENT = 8 CHARACTER SET = utf8 COLLATE = utf8_general_ci ROW_FORMAT = Dynamic;

-- ----------------------------
-- Records of likes
-- ----------------------------
INSERT INTO `likes` VALUES (5, 1, 9, 2);
INSERT INTO `likes` VALUES (6, 1, 13, 2);
INSERT INTO `likes` VALUES (7, 1, 12, 1);

-- ----------------------------
-- Table structure for users
-- ----------------------------
DROP TABLE IF EXISTS `users`;
CREATE TABLE `users`  (
  `id` int(11) NOT NULL AUTO_INCREMENT,
  `name` varchar(255) CHARACTER SET utf8 COLLATE utf8_general_ci NULL DEFAULT NULL,
  `password` varchar(255) CHARACTER SET utf8 COLLATE utf8_general_ci NULL DEFAULT NULL,
  `token` varchar(255) CHARACTER SET utf8 COLLATE utf8_general_ci NULL DEFAULT NULL,
  PRIMARY KEY (`id`) USING BTREE
) ENGINE = InnoDB AUTO_INCREMENT = 3 CHARACTER SET = utf8 COLLATE = utf8_general_ci ROW_FORMAT = Dynamic;

-- ----------------------------
-- Records of users
-- ----------------------------
INSERT INTO `users` VALUES (1, 'xiaoming', '123456', 'xiaoming123456');
INSERT INTO `users` VALUES (2, '1240047299', '123456', '1240047299123456');

-- ----------------------------
-- Table structure for videos
-- ----------------------------
DROP TABLE IF EXISTS `videos`;
CREATE TABLE `videos`  (
  `id` int(255) NOT NULL AUTO_INCREMENT COMMENT '视频id',
  `author_id` int(255) NULL DEFAULT NULL COMMENT ' 作者id',
  `play_url` varchar(255) CHARACTER SET utf8 COLLATE utf8_general_ci NULL DEFAULT NULL COMMENT ' 播放url',
  `cover_url` varchar(255) CHARACTER SET utf8 COLLATE utf8_general_ci NULL DEFAULT NULL COMMENT ' 封面url',
  `publish_time` datetime(0) NULL DEFAULT NULL COMMENT ' 发布时间戳',
  `title` varchar(255) CHARACTER SET utf8 COLLATE utf8_general_ci NULL DEFAULT NULL COMMENT ' 视频标题',
  `favorite_count` int(255) NULL DEFAULT NULL COMMENT '视频的点赞总数',
  `comment_count` int(255) NULL DEFAULT NULL COMMENT '视频的评论总数',
  `is_favorite` tinyint(1) NULL DEFAULT NULL COMMENT '是否点赞，true为已点赞，false为未点赞',
  PRIMARY KEY (`id`) USING BTREE
) ENGINE = InnoDB AUTO_INCREMENT = 15 CHARACTER SET = utf8 COLLATE = utf8_general_ci ROW_FORMAT = Dynamic;

-- ----------------------------
-- Records of videos
-- ----------------------------
INSERT INTO `videos` VALUES (1, 1, 'http://192.168.124.24:8080/public/bear.mp4', 'http://192.168.124.24:8080/public/bear-1283347_1280.jpg', '2023-02-19 00:00:00', 'fsdfsdf', NULL, NULL, NULL);
INSERT INTO `videos` VALUES (12, 1, 'http://192.168.124.24:8080/public/1_test.mp4', 'http://192.168.124.24:8080/public/1_test.png', '2023-02-22 16:45:59', 'test', 1, 0, 1);
INSERT INTO `videos` VALUES (13, 1, 'http://192.168.124.24:8080/public/1_oceans.mp4', 'http://192.168.124.24:8080/public/1_oceans.png', '2023-02-22 16:46:36', 'ocean', 0, 0, 0);
INSERT INTO `videos` VALUES (14, 1, 'http://192.168.124.24:8080/public/1_oceans.mp4', 'http://192.168.124.24:8080/public/1_oceans.png', '2023-02-22 16:52:57', 'ocean1', 0, 0, 0);

SET FOREIGN_KEY_CHECKS = 1;
