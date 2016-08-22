/*
SQLyog 企业版 - MySQL GUI v8.14 
MySQL - 5.7.9 : Database - wx
*********************************************************************
*/

/*!40101 SET NAMES utf8 */;

/*!40101 SET SQL_MODE=''*/;

/*!40014 SET @OLD_UNIQUE_CHECKS=@@UNIQUE_CHECKS, UNIQUE_CHECKS=0 */;
/*!40014 SET @OLD_FOREIGN_KEY_CHECKS=@@FOREIGN_KEY_CHECKS, FOREIGN_KEY_CHECKS=0 */;
/*!40101 SET @OLD_SQL_MODE=@@SQL_MODE, SQL_MODE='NO_AUTO_VALUE_ON_ZERO' */;
/*!40111 SET @OLD_SQL_NOTES=@@SQL_NOTES, SQL_NOTES=0 */;
CREATE DATABASE /*!32312 IF NOT EXISTS*/`wx` /*!40100 DEFAULT CHARACTER SET utf8 */;

/*Table structure for table `artices` */

CREATE TABLE `artices` (
  `id` varchar(50) NOT NULL COMMENT '表主键',
  `cTime` datetime NOT NULL COMMENT '创建数据时间',
  `version` bigint(20) NOT NULL COMMENT '版本号',
  `updateTime` datetime DEFAULT NULL COMMENT '更新时间',
  `title` varchar(200) NOT NULL COMMENT '文章标题',
  `thumbMediaID` varchar(50) NOT NULL COMMENT '图文消息的封面图片素材id（必须是永久mediaID）',
  `author` varchar(50) NOT NULL COMMENT '作者',
  `digest` varchar(200) NOT NULL COMMENT '图文消息的摘要，仅有单图文消息才有摘要，多图文此处为空',
  `show_cover_pic` char(2) DEFAULT NULL COMMENT '是否显示封面，0为false，即不显示，1为true，即显示',
  `content` blob NOT NULL COMMENT '内容',
  `content_source_url` varchar(200) NOT NULL COMMENT '原文链接地址',
  PRIMARY KEY (`id`)
) ENGINE=MyISAM DEFAULT CHARSET=utf8 CHECKSUM=1 DELAY_KEY_WRITE=1 ROW_FORMAT=DYNAMIC;

/*Data for the table `artices` */

/*Table structure for table `subscribe` */

CREATE TABLE `subscribe` (
  `id` varchar(50) NOT NULL COMMENT '主键',
  `cTime` datetime NOT NULL COMMENT '创建时间',
  `updateTime` datetime NOT NULL COMMENT '更新时间',
  `version` int(20) NOT NULL COMMENT '版本号',
  `toUserName` varchar(200) NOT NULL COMMENT '开发者微信号',
  `fromUserName` varchar(200) NOT NULL COMMENT '关注者微信号',
  `createTime` int(11) NOT NULL COMMENT '微信创建时间',
  `msgType` varchar(20) NOT NULL COMMENT '消息类型',
  `event` varchar(10) NOT NULL DEFAULT 'subscribe' COMMENT '事件类型，subscribe(订阅)、unsubscribe(取消订阅)',
  PRIMARY KEY (`id`)
) ENGINE=MyISAM DEFAULT CHARSET=utf8 CHECKSUM=1 DELAY_KEY_WRITE=1 ROW_FORMAT=DYNAMIC COMMENT='关注时间表';

/*Data for the table `subscribe` */

/*!40101 SET SQL_MODE=@OLD_SQL_MODE */;
/*!40014 SET FOREIGN_KEY_CHECKS=@OLD_FOREIGN_KEY_CHECKS */;
/*!40014 SET UNIQUE_CHECKS=@OLD_UNIQUE_CHECKS */;
/*!40111 SET SQL_NOTES=@OLD_SQL_NOTES */;
