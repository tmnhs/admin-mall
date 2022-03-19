/*
SQLyog Ultimate v12.09 (64 bit)
MySQL - 8.0.27 : Database - hust_mall
*********************************************************************
*/

/*!40101 SET NAMES utf8 */;

/*!40101 SET SQL_MODE=''*/;

/*!40014 SET @OLD_UNIQUE_CHECKS=@@UNIQUE_CHECKS, UNIQUE_CHECKS=0 */;
/*!40014 SET @OLD_FOREIGN_KEY_CHECKS=@@FOREIGN_KEY_CHECKS, FOREIGN_KEY_CHECKS=0 */;
/*!40101 SET @OLD_SQL_MODE=@@SQL_MODE, SQL_MODE='NO_AUTO_VALUE_ON_ZERO' */;
/*!40111 SET @OLD_SQL_NOTES=@@SQL_NOTES, SQL_NOTES=0 */;
CREATE DATABASE /*!32312 IF NOT EXISTS*/`hust_mall` /*!40100 DEFAULT CHARACTER SET utf8mb4 COLLATE utf8mb4_0900_ai_ci */ /*!80016 DEFAULT ENCRYPTION='N' */;

USE `hust_mall`;

/*Table structure for table `m_ad` */

CREATE TABLE `m_ad` (
  `id` bigint NOT NULL AUTO_INCREMENT,
  `name` varchar(32) DEFAULT NULL COMMENT '广告名称',
  `type` int DEFAULT NULL COMMENT '广告类型',
  `image` varchar(500) DEFAULT NULL COMMENT '图片地址',
  `width` int DEFAULT NULL COMMENT '图片宽度 ，单位：px',
  `height` int DEFAULT NULL COMMENT '图片高度',
  `introduce` varchar(255) DEFAULT NULL COMMENT '广告介绍',
  `url` varchar(500) DEFAULT NULL COMMENT '链接地址',
  `status` int DEFAULT NULL COMMENT '是否激活，0->未激活，1->已激活',
  `created_time` bigint DEFAULT '0' COMMENT '创建时间',
  `expiration_time` bigint DEFAULT '0' COMMENT '过期时间',
  PRIMARY KEY (`id`)
) ENGINE=InnoDB AUTO_INCREMENT=2 DEFAULT CHARSET=utf8mb4 COLLATE=utf8mb4_0900_ai_ci;

/*Table structure for table `m_address` */

CREATE TABLE `m_address` (
  `id` bigint NOT NULL AUTO_INCREMENT,
  `region` int DEFAULT NULL COMMENT '地址区域，1表示韵苑，2表示沁苑，3表示紫菘，4表示其他',
  `other` varchar(80) DEFAULT NULL COMMENT '区域为其他时的地址',
  `unit` varchar(32) DEFAULT NULL COMMENT '单元，比如7栋、十二舍',
  PRIMARY KEY (`id`)
) ENGINE=InnoDB AUTO_INCREMENT=3 DEFAULT CHARSET=utf8mb4 COLLATE=utf8mb4_0900_ai_ci;

/*Table structure for table `m_feedback` */

CREATE TABLE `m_feedback` (
  `id` bigint NOT NULL AUTO_INCREMENT,
  `uid` bigint DEFAULT NULL COMMENT '用户UID',
  `content` text COMMENT '反馈内容',
  `status` int DEFAULT NULL COMMENT '状态，0->未读，1表示已阅',
  `created_time` bigint DEFAULT '0' COMMENT '创建时间',
  PRIMARY KEY (`id`)
) ENGINE=InnoDB AUTO_INCREMENT=5 DEFAULT CHARSET=utf8mb4 COLLATE=utf8mb4_0900_ai_ci;

/*Table structure for table `m_follow` */

CREATE TABLE `m_follow` (
  `id` bigint NOT NULL AUTO_INCREMENT,
  `uid` bigint DEFAULT NULL COMMENT '用户id',
  `follow_ids` varchar(500) DEFAULT NULL COMMENT '所有关注者id',
  `created_time` bigint DEFAULT '0' COMMENT '创建时间',
  PRIMARY KEY (`id`)
) ENGINE=InnoDB DEFAULT CHARSET=utf8mb4 COLLATE=utf8mb4_0900_ai_ci;

/*Table structure for table `m_idle_item` */

CREATE TABLE `m_idle_item` (
  `id` bigint NOT NULL AUTO_INCREMENT,
  `uid` bigint DEFAULT NULL COMMENT '用户编号id',
  `name` varchar(64) NOT NULL COMMENT '物品名称',
  `category_id` bigint DEFAULT NULL COMMENT '分类编号id',
  `brand` varchar(32) DEFAULT NULL COMMENT '品牌名称',
  `album_images` varchar(2000) CHARACTER SET utf8mb4 COLLATE utf8mb4_0900_ai_ci DEFAULT NULL COMMENT '画册图片，连产品图片限制为5张，以逗号分割，默认第一张为主图',
  `purchase_channel` int DEFAULT NULL COMMENT '购买渠道,1->网上购买，2->实体店购买',
  `quality` int DEFAULT NULL COMMENT '几成新，范围1-10，其中10为全新未拆封',
  `shelf_life` varchar(32) DEFAULT NULL COMMENT '剩余保质期',
  `sort` int DEFAULT NULL COMMENT '排序权值',
  `origin_price` decimal(10,2) DEFAULT NULL COMMENT '原价',
  `price` decimal(10,2) DEFAULT NULL COMMENT '预售价格',
  `description` varchar(255) DEFAULT NULL COMMENT '物品描述',
  `stock` int DEFAULT NULL COMMENT '库存（数量）',
  `trade_way` int DEFAULT NULL COMMENT '交易方式,1为自取件，2为派送件，3为自定义交易地点',
  `sale_status` int DEFAULT NULL COMMENT '销售状态，1表示已售，0表示未销售',
  `publish_status` int DEFAULT NULL COMMENT '上架状态：0->下架；1->上架',
  `created_time` bigint DEFAULT '0' COMMENT '创建时间',
  `updated_time` bigint DEFAULT '0' COMMENT '更新时间',
  PRIMARY KEY (`id`),
  KEY `idx_purchase_name_category_trade` (`name`,`category_id`,`trade_way`),
  KEY `idx_idle_name_category_price_trade` (`name`,`category_id`,`price`,`trade_way`)
) ENGINE=InnoDB AUTO_INCREMENT=5 DEFAULT CHARSET=utf8mb4 COLLATE=utf8mb4_0900_ai_ci;

/*Table structure for table `m_item_category` */

CREATE TABLE `m_item_category` (
  `id` bigint NOT NULL AUTO_INCREMENT,
  `name` varchar(64) CHARACTER SET utf8mb4 COLLATE utf8mb4_0900_ai_ci NOT NULL COMMENT '分类名称',
  `parent_id` bigint DEFAULT '0' COMMENT '上级分类的编号：0表示一级分类',
  `level` int DEFAULT NULL COMMENT '分类级别：1->1级；2->2级',
  `item_count` int DEFAULT NULL COMMENT '商品数量',
  `show_status` int DEFAULT '1' COMMENT '显示状态：0->不显示；1->显示',
  `sort` int DEFAULT NULL COMMENT '排序权重',
  `icon` varchar(255) DEFAULT NULL COMMENT '图标',
  `description` text COMMENT '描述',
  PRIMARY KEY (`id`)
) ENGINE=InnoDB AUTO_INCREMENT=4 DEFAULT CHARSET=utf8mb4 COLLATE=utf8mb4_0900_ai_ci;

/*Table structure for table `m_item_list` */

CREATE TABLE `m_item_list` (
  `id` bigint NOT NULL AUTO_INCREMENT,
  `uid` bigint DEFAULT NULL COMMENT '用户id',
  `item_type` int DEFAULT NULL COMMENT '列表类型：0表示闲置，1表示求购，2表示收藏',
  `item_ids` varchar(500) DEFAULT NULL COMMENT '所有物品id',
  PRIMARY KEY (`id`)
) ENGINE=InnoDB DEFAULT CHARSET=utf8mb4 COLLATE=utf8mb4_0900_ai_ci;

/*Table structure for table `m_message` */

CREATE TABLE `m_message` (
  `id` bigint NOT NULL AUTO_INCREMENT,
  `msg_type` int DEFAULT '1' COMMENT '类型 1：文字；47：emoji；43：音频；436207665：红包；49：文件；48：位置；3：图片',
  `from_uid` bigint DEFAULT NULL COMMENT '存储的消息UID,0表示系统消息',
  `item_id` bigint DEFAULT NULL COMMENT '针对的物品id',
  `to_uid` bigint DEFAULT NULL COMMENT '消息接收UID',
  `content` text COMMENT '消息内容',
  `created_time` bigint DEFAULT '0' COMMENT '创建时间',
  PRIMARY KEY (`id`)
) ENGINE=InnoDB DEFAULT CHARSET=utf8mb4 COLLATE=utf8mb4_0900_ai_ci;

/*Table structure for table `m_order` */

CREATE TABLE `m_order` (
  `id` bigint NOT NULL AUTO_INCREMENT,
  `order_sn` varchar(64) DEFAULT NULL COMMENT '订单编号',
  `receiver_id` bigint DEFAULT NULL COMMENT '买家id',
  `type` int DEFAULT NULL COMMENT '订单类型，0->闲置，1->求购',
  `item_id` bigint DEFAULT NULL COMMENT '物品id',
  `total_amount` decimal(10,2) DEFAULT NULL COMMENT '订单总金额',
  `address` varchar(32) DEFAULT NULL COMMENT '交易地点',
  `trade_time` bigint DEFAULT NULL COMMENT '交易时间',
  `status` int DEFAULT NULL COMMENT '订单状态：1->已完成；2->进行中;3->已关闭；4->无效订单',
  `note` varchar(500) DEFAULT NULL COMMENT '订单备注',
  `confirm_status` int DEFAULT NULL COMMENT '确认收货状态：0->未确认；1->已确认',
  `delete_status` int NOT NULL DEFAULT '0' COMMENT '删除状态：0->未删除；1->已删除',
  `created_time` bigint DEFAULT '0' COMMENT '创建时间',
  `updated_time` bigint DEFAULT '0' COMMENT '更新时间',
  PRIMARY KEY (`id`),
  KEY `idx_order_item_status_address` (`item_id`,`confirm_status`,`address`)
) ENGINE=InnoDB AUTO_INCREMENT=3 DEFAULT CHARSET=utf8mb4 COLLATE=utf8mb4_0900_ai_ci;

/*Table structure for table `m_order_operate_history` */

CREATE TABLE `m_order_operate_history` (
  `id` bigint NOT NULL AUTO_INCREMENT,
  `order_id` bigint DEFAULT NULL COMMENT '订单id',
  `type` int DEFAULT NULL COMMENT '操作人类型：0->普通用户，1->后台管理人员',
  `uid` bigint DEFAULT NULL COMMENT '操作人id',
  `order_status` int DEFAULT NULL COMMENT '订单状态：1->已完成；2->进行中;3->已关闭；4->无效订单',
  `note` varchar(500) DEFAULT NULL COMMENT '备注',
  `created_time` bigint DEFAULT NULL COMMENT '操作时间',
  PRIMARY KEY (`id`)
) ENGINE=InnoDB AUTO_INCREMENT=2 DEFAULT CHARSET=utf8mb4 COLLATE=utf8mb4_0900_ai_ci;

/*Table structure for table `m_order_return` */

CREATE TABLE `m_order_return` (
  `id` bigint NOT NULL AUTO_INCREMENT,
  `order_id` bigint NOT NULL COMMENT '订单id',
  `reason` varchar(100) DEFAULT NULL COMMENT '退货原因',
  `sort` int DEFAULT '0' COMMENT '排序',
  `status` int DEFAULT '0' COMMENT '是否已退货，1->退货，0->不退货',
  `created_time` bigint DEFAULT '0' COMMENT '创建时间',
  PRIMARY KEY (`id`)
) ENGINE=InnoDB AUTO_INCREMENT=3 DEFAULT CHARSET=utf8mb4 COLLATE=utf8mb4_0900_ai_ci;

/*Table structure for table `m_profile` */

CREATE TABLE `m_profile` (
  `uid` bigint NOT NULL COMMENT '用户id',
  `username` varchar(50) NOT NULL COMMENT '用户名称',
  `gender` int DEFAULT NULL COMMENT '性别',
  `status` int NOT NULL DEFAULT '1' COMMENT '用户状态，1表示正常，0表示封禁',
  `address_id` bigint DEFAULT NULL COMMENT '用户地址编号',
  `phone` varchar(20) DEFAULT NULL COMMENT '电话号码',
  `email` varchar(32) DEFAULT NULL COMMENT '电子邮箱',
  `qq` varchar(32) DEFAULT NULL COMMENT 'QQ号',
  `wechat` varchar(32) DEFAULT NULL COMMENT '微信',
  `avatar` varchar(32) DEFAULT NULL COMMENT '头像',
  `bg_image` varchar(32) DEFAULT NULL COMMENT '背景图像',
  `is_online` int DEFAULT '0' COMMENT '用户是否在线，1表示在线，0表示离线',
  `finish_order_count` int DEFAULT NULL COMMENT '交易完成次数',
  `created_time` bigint DEFAULT '0' COMMENT '创建时间',
  `updated_time` bigint DEFAULT '0' COMMENT '更新时间',
  PRIMARY KEY (`uid`),
  KEY `idx_user_name_statu` (`username`,`status`)
) ENGINE=InnoDB DEFAULT CHARSET=utf8mb4 COLLATE=utf8mb4_0900_ai_ci;

/*Table structure for table `m_purchase_item` */

CREATE TABLE `m_purchase_item` (
  `id` bigint NOT NULL AUTO_INCREMENT,
  `uid` bigint DEFAULT NULL COMMENT '用户编号id',
  `name` varchar(64) NOT NULL COMMENT '求购物品名称',
  `category_id` bigint DEFAULT NULL COMMENT '分类编号id',
  `sort` int DEFAULT NULL COMMENT '排序权值',
  `price` decimal(10,2) DEFAULT NULL COMMENT '预估价格',
  `description` varchar(255) DEFAULT NULL COMMENT '求购物品描述',
  `trade_way` int DEFAULT NULL COMMENT '交易方式,1为自取件，2为派送件，3为自定义交易地点',
  `status` int DEFAULT NULL COMMENT '状态，1表示已购，0表示未购得',
  `delete_status` int DEFAULT NULL COMMENT '删除状态：0->未删除；1->已删除',
  `created_time` bigint DEFAULT '0' COMMENT '创建时间',
  `updated_time` bigint DEFAULT '0' COMMENT '更新时间',
  PRIMARY KEY (`id`)
) ENGINE=InnoDB AUTO_INCREMENT=3 DEFAULT CHARSET=utf8mb4 COLLATE=utf8mb4_0900_ai_ci;

/*Table structure for table `m_report` */

CREATE TABLE `m_report` (
  `id` bigint NOT NULL AUTO_INCREMENT,
  `uid` bigint NOT NULL COMMENT '用户UID',
  `item_id` bigint NOT NULL COMMENT '举报的物品对象id',
  `type` int NOT NULL COMMENT '举报类型，1->广告或垃圾信息、2->色情类信息、3->违法或政治敏感信息、4->欺诈类信息、5->其他',
  `other` varchar(255) DEFAULT NULL COMMENT '举报类型type为其他时的举报内容',
  `status` int DEFAULT '0' COMMENT '审核状态，0->未审核通过，1->审核通过',
  `created_time` bigint DEFAULT '0' COMMENT '创建时间',
  PRIMARY KEY (`id`)
) ENGINE=InnoDB AUTO_INCREMENT=3 DEFAULT CHARSET=utf8mb4 COLLATE=utf8mb4_0900_ai_ci;

/*Table structure for table `m_user` */

CREATE TABLE `m_user` (
  `id` bigint NOT NULL AUTO_INCREMENT,
  `username` varchar(50) NOT NULL COMMENT '用户名称',
  `PASSWORD` varchar(88) NOT NULL COMMENT '用户密码',
  `created_time` bigint DEFAULT '0' COMMENT '创建时间',
  `updated_time` bigint DEFAULT '0' COMMENT '更新时间',
  PRIMARY KEY (`id`)
) ENGINE=InnoDB AUTO_INCREMENT=3 DEFAULT CHARSET=utf8mb4 COLLATE=utf8mb4_0900_ai_ci;

/*Table structure for table `m_user_login_log` */

CREATE TABLE `m_user_login_log` (
  `id` bigint NOT NULL AUTO_INCREMENT,
  `uid` bigint NOT NULL COMMENT '用户id',
  `username` varchar(32) DEFAULT NULL COMMENT '用户名',
  `ip` varchar(32) DEFAULT NULL COMMENT '登录ip地址',
  `login_type` int DEFAULT NULL COMMENT '登录类型',
  `city` varchar(32) DEFAULT NULL COMMENT '登录城市',
  `login_time` bigint DEFAULT NULL COMMENT '登录时间',
  PRIMARY KEY (`id`)
) ENGINE=InnoDB AUTO_INCREMENT=3 DEFAULT CHARSET=utf8mb4 COLLATE=utf8mb4_0900_ai_ci;

/*Table structure for table `m_user_statistics_info` */

CREATE TABLE `m_user_statistics_info` (
  `id` bigint NOT NULL AUTO_INCREMENT,
  `uid` bigint NOT NULL COMMENT '用户id',
  `finish_order_count` int DEFAULT NULL COMMENT '交易完成次数',
  `login_count` int DEFAULT NULL COMMENT '登录总数',
  `last_login_time` bigint DEFAULT '0' COMMENT '最近登录时间',
  `follow_count` int DEFAULT NULL COMMENT '好友数量',
  PRIMARY KEY (`id`)
) ENGINE=InnoDB AUTO_INCREMENT=3 DEFAULT CHARSET=utf8mb4 COLLATE=utf8mb4_0900_ai_ci;

/*!40101 SET SQL_MODE=@OLD_SQL_MODE */;
/*!40014 SET FOREIGN_KEY_CHECKS=@OLD_FOREIGN_KEY_CHECKS */;
/*!40014 SET UNIQUE_CHECKS=@OLD_UNIQUE_CHECKS */;
/*!40111 SET SQL_NOTES=@OLD_SQL_NOTES */;
