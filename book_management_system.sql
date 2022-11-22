/*
 Navicat Premium Data Transfer

 Source Server         : 172.0.0.1
 Source Server Type    : MySQL
 Source Server Version : 80028 (8.0.28)
 Source Host           : localhost:3306
 Source Schema         : book_management_system

 Target Server Type    : MySQL
 Target Server Version : 80028 (8.0.28)
 File Encoding         : 65001

 Date: 01/11/2022 20:41:28
*/

SET NAMES utf8mb4;
SET FOREIGN_KEY_CHECKS = 0;

-- ----------------------------
-- Table structure for book_type
-- ----------------------------
DROP TABLE IF EXISTS `book_type`;
CREATE TABLE `book_type` (
  `id` int NOT NULL AUTO_INCREMENT,
  `typeName` varchar(40) NOT NULL COMMENT '分类名称',
  `level` varchar(8) DEFAULT NULL COMMENT '层级',
  `pId` int DEFAULT NULL COMMENT '父级ID',
  `remake` text COMMENT '备注',
  PRIMARY KEY (`id`)
) ENGINE=InnoDB AUTO_INCREMENT=28 DEFAULT CHARSET=utf8mb4 COLLATE=utf8mb4_0900_ai_ci;

-- ----------------------------
-- Table structure for borrow
-- ----------------------------
DROP TABLE IF EXISTS `borrow`;
CREATE TABLE `borrow` (
  `id` int NOT NULL AUTO_INCREMENT COMMENT '借书编号',
  `borrow_reader_id` bigint DEFAULT NULL COMMENT '借书人id',
  `borrow_book_isbn` varchar(30) DEFAULT NULL COMMENT '被借的书的ISBN',
  `is_borrow` tinyint(1) DEFAULT NULL COMMENT '是否同意借出',
  `borrow_time` varchar(20) DEFAULT NULL COMMENT '借书时间',
  `should_return_time` varchar(30) DEFAULT NULL COMMENT '应当还书时间',
  `is_return` smallint DEFAULT NULL COMMENT '是否归还',
  `really_return_time` varchar(30) DEFAULT NULL COMMENT '实际归还时间',
  `agree_borrow_time` varchar(30) DEFAULT NULL COMMENT '同意借书的时间',
  PRIMARY KEY (`id`)
) ENGINE=InnoDB AUTO_INCREMENT=39 DEFAULT CHARSET=utf8mb4 COLLATE=utf8mb4_0900_ai_ci;

-- ----------------------------
-- Table structure for g_book_info
-- ----------------------------
DROP TABLE IF EXISTS `g_book_info`;
CREATE TABLE `g_book_info` (
  `isbn` varchar(191) NOT NULL COMMENT '图书唯一编号',
  `bookName` varchar(191) NOT NULL,
  `author` varchar(191) DEFAULT NULL,
  `publisher` varchar(191) DEFAULT NULL,
  `publishTime` varchar(191) DEFAULT NULL,
  `bookStock` bigint unsigned DEFAULT NULL,
  `price` decimal(8,2) DEFAULT NULL,
  `typeId` int DEFAULT NULL,
  `context` text,
  `pageNum` varchar(191) DEFAULT NULL,
  `translator` varchar(191) DEFAULT NULL,
  `photo` varchar(191) DEFAULT NULL,
  PRIMARY KEY (`isbn`),
  KEY `FK_book_type` (`typeId`),
  CONSTRAINT `FK_book_type` FOREIGN KEY (`typeId`) REFERENCES `book_type` (`id`) ON DELETE RESTRICT ON UPDATE CASCADE
) ENGINE=InnoDB DEFAULT CHARSET=utf8mb4 COLLATE=utf8mb4_0900_ai_ci;

-- ----------------------------
-- Table structure for g_book_list
-- ----------------------------
DROP TABLE IF EXISTS `g_book_list`;
CREATE TABLE `g_book_list` (
  `id` bigint unsigned NOT NULL AUTO_INCREMENT,
  `name` varchar(191) NOT NULL,
  `user_id` bigint unsigned NOT NULL,
  `photo` varchar(191) DEFAULT NULL,
  `remake` varchar(191) DEFAULT NULL,
  `time` datetime(3) DEFAULT NULL,
  PRIMARY KEY (`id`),
  UNIQUE KEY `name` (`name`)
) ENGINE=InnoDB AUTO_INCREMENT=10 DEFAULT CHARSET=utf8mb4 COLLATE=utf8mb4_0900_ai_ci;

-- ----------------------------
-- Table structure for g_book_list_map
-- ----------------------------
DROP TABLE IF EXISTS `g_book_list_map`;
CREATE TABLE `g_book_list_map` (
  `book_info_isbn` varchar(191) NOT NULL COMMENT '图书唯一编号',
  `book_list_id` bigint unsigned NOT NULL,
  PRIMARY KEY (`book_info_isbn`,`book_list_id`)
) ENGINE=InnoDB DEFAULT CHARSET=utf8mb4 COLLATE=utf8mb4_0900_ai_ci;

-- ----------------------------
-- Table structure for g_book_user_map
-- ----------------------------
DROP TABLE IF EXISTS `g_book_user_map`;
CREATE TABLE `g_book_user_map` (
  `user_id` bigint NOT NULL,
  `book_list_id` bigint unsigned NOT NULL,
  PRIMARY KEY (`user_id`,`book_list_id`)
) ENGINE=InnoDB DEFAULT CHARSET=utf8mb4 COLLATE=utf8mb4_0900_ai_ci;

-- ----------------------------
-- Table structure for g_user
-- ----------------------------
DROP TABLE IF EXISTS `g_user`;
CREATE TABLE `g_user` (
  `id` bigint NOT NULL AUTO_INCREMENT,
  `user_name` varchar(191) DEFAULT NULL,
  `user_password` varchar(191) DEFAULT NULL,
  `sex` tinyint DEFAULT NULL,
  `birthday` varchar(191) DEFAULT NULL,
  `borrow_book_count` int DEFAULT NULL,
  `phone` varchar(191) DEFAULT NULL,
  `remake` varchar(191) DEFAULT NULL,
  `email` varchar(191) DEFAULT NULL,
  `role` int DEFAULT NULL,
  PRIMARY KEY (`id`),
  KEY `FK_USER_ROLE` (`role`),
  CONSTRAINT `FK_USER_ROLE` FOREIGN KEY (`role`) REFERENCES `role` (`id`) ON DELETE RESTRICT ON UPDATE CASCADE
) ENGINE=InnoDB AUTO_INCREMENT=3 DEFAULT CHARSET=utf8mb4 COLLATE=utf8mb4_0900_ai_ci;

-- ----------------------------
-- Table structure for role
-- ----------------------------
DROP TABLE IF EXISTS `role`;
CREATE TABLE `role` (
  `id` int NOT NULL AUTO_INCREMENT,
  `role_name` varchar(20) DEFAULT NULL COMMENT '角色名',
  `role_weight` int DEFAULT NULL COMMENT '角色权重，用户为1，管理员为2',
  PRIMARY KEY (`id`)
) ENGINE=InnoDB AUTO_INCREMENT=12 DEFAULT CHARSET=utf8mb4 COLLATE=utf8mb4_0900_ai_ci;

-- ----------------------------
-- View structure for book_type_with_pname
-- ----------------------------
DROP VIEW IF EXISTS `book_type_with_pname`;
CREATE ALGORITHM = UNDEFINED SQL SECURITY DEFINER VIEW `book_type_with_pname` AS select `t1`.`id` AS `id`,`t1`.`typeName` AS `typeName`,`t1`.`level` AS `level`,`t1`.`remake` AS `remake`,`t1`.`pId` AS `pId`,ifnull(`t2`.`typeName`,'') AS `pName` from (`book_type` `t1` left join `book_type` `t2` on((`t1`.`pId` = `t2`.`id`)));

-- ----------------------------
-- View structure for borrow_with_name
-- ----------------------------
DROP VIEW IF EXISTS `borrow_with_name`;
CREATE ALGORITHM = UNDEFINED SQL SECURITY DEFINER VIEW `borrow_with_name` AS select `b`.`id` AS `id`,`b`.`borrow_book_isbn` AS `borrow_book_isbn`,`bi`.`bookName` AS `borrow_book_name`,`b`.`borrow_reader_id` AS `borrow_reader_id`,`u`.`user_name` AS `borrow_reader_name`,`u`.`email` AS `email`,`b`.`borrow_time` AS `borrow_time`,`b`.`agree_borrow_time` AS `agree_borrow_time`,`b`.`is_borrow` AS `is_borrow`,`b`.`is_return` AS `is_return`,`b`.`really_return_time` AS `really_return_time`,`b`.`should_return_time` AS `should_return_time` from ((`borrow` `b` left join `g_book_info` `bi` on((`b`.`borrow_book_isbn` = `bi`.`isbn`))) left join `g_user` `u` on((`b`.`borrow_reader_id` = `u`.`id`))) where ((`b`.`borrow_book_isbn` = `bi`.`isbn`) and (`b`.`borrow_reader_id` = `u`.`id`));

SET FOREIGN_KEY_CHECKS = 1;

INSERT INTO `book_management_system`.`role` (`id`, `role_name`, `role_weight`) VALUES (1, 'admin', 99);
INSERT INTO `book_management_system`.`g_user` (`id`, `user_name`, `user_password`, `sex`, `birthday`, `borrow_book_count`, `phone`, `remake`, `email`, `role`) VALUES (1, 'admin', '21232f297a57a5a743894a0e4a801fc3', 1, '1996-09-10', 0, '18136102555', '这是管理员账号', '1481410897@qq.com', 1);
