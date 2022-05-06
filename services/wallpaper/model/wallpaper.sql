/*
 Data Transfer
*/

SET NAMES utf8mb4;
SET FOREIGN_KEY_CHECKS = 0;

-- ----------------------------
-- Table structure for wallpaper_tab
-- ----------------------------
DROP TABLE IF EXISTS `wallpaper_tab`;
CREATE TABLE `wallpaper_tab` (
    `id` bigint unsigned NOT NULL AUTO_INCREMENT COMMENT 'id',
    `wid` varchar(128) NOT NULL COMMENT 'wallpaper id',
    `name` varchar(128) NOT NULL DEFAULT '' COMMENT 'name',
    `tp` varchar(128) NOT NULL DEFAULT '' COMMENT 'type',
    `category` varchar(128) NOT NULL DEFAULT '' COMMENT 'category',
    `image_url` varchar(256) NOT NULL DEFAULT '' COMMENT 'image url',
    `author` varchar(128) NOT NULL DEFAULT '' COMMENT 'author',
    `desc` varchar(512) NOT NULL DEFAULT '' COMMENT 'desc',
    `del_flag` char(1) NOT NULL DEFAULT '0' COMMENT 'del flag（0-normal 1-delete)',
    `create_time` bigint unsigned NOT NULL DEFAULT 0 COMMENT 'create time',
    `update_time` bigint unsigned NOT NULL DEFAULT 0 COMMENT 'update time',
    PRIMARY KEY (`id`),
    UNIQUE KEY `uniq_wid` (`wid`) USING BTREE COMMENT 'uniq_wid',
    KEY `idx_tp` (`tp`) USING BTREE COMMENT 'idx_tp',
    KEY `idx_category` (`category`) USING BTREE COMMENT 'idx_category',
    KEY `idx_author_desc` (`author`,`desc`) USING BTREE COMMENT 'idx_author_desc'
) ENGINE=InnoDB AUTO_INCREMENT=1 DEFAULT CHARSET=utf8mb4 COLLATE=utf8mb4_unicode_ci COMMENT='wallpaper table';

-- ----------------------------
-- Table structure for wallpaper_category_tab
-- ----------------------------
DROP TABLE IF EXISTS `wallpaper_category_tab`;
CREATE TABLE `wallpaper_category_tab` (
     `id` bigint unsigned NOT NULL AUTO_INCREMENT COMMENT 'id',
     `category` varchar(128) NOT NULL DEFAULT '' COMMENT 'category',
     `desc` varchar(512) NOT NULL DEFAULT '' COMMENT 'desc',
     `del_flag` char(1) NOT NULL DEFAULT '0' COMMENT 'del flag（0-normal 1-delete)',
     `create_time` bigint unsigned NOT NULL DEFAULT 0 COMMENT 'create time',
     `update_time` bigint unsigned NOT NULL DEFAULT 0 COMMENT 'update time',
     PRIMARY KEY (`id`),
     KEY `idx_category` (`category`) USING BTREE COMMENT 'idx_category'
) ENGINE=InnoDB AUTO_INCREMENT=1 DEFAULT CHARSET=utf8mb4 COLLATE=utf8mb4_unicode_ci COMMENT='wallpaper category table';
insert into wallpaper_category_tab (category, `desc`, del_flag, create_time, update_time) VALUES ('治愈', '', 0, unix_timestamp(now()), unix_timestamp(now()));
insert into wallpaper_category_tab (category, `desc`, del_flag, create_time, update_time) VALUES ('科技', '', 0, unix_timestamp(now()), unix_timestamp(now()));
insert into wallpaper_category_tab (category, `desc`, del_flag, create_time, update_time) VALUES ('风景', '', 0, unix_timestamp(now()), unix_timestamp(now()));
insert into wallpaper_category_tab (category, `desc`, del_flag, create_time, update_time) VALUES ('人物', '', 0, unix_timestamp(now()), unix_timestamp(now()));
insert into wallpaper_category_tab (category, `desc`, del_flag, create_time, update_time) VALUES ('游戏', '', 0, unix_timestamp(now()), unix_timestamp(now()));
insert into wallpaper_category_tab (category, `desc`, del_flag, create_time, update_time) VALUES ('汽车', '', 0, unix_timestamp(now()), unix_timestamp(now()));
insert into wallpaper_category_tab (category, `desc`, del_flag, create_time, update_time) VALUES ('艺术', '', 0, unix_timestamp(now()), unix_timestamp(now()));
insert into wallpaper_category_tab (category, `desc`, del_flag, create_time, update_time) VALUES ('文字', '', 0, unix_timestamp(now()), unix_timestamp(now()));
insert into wallpaper_category_tab (category, `desc`, del_flag, create_time, update_time) VALUES ('宠物', '', 0, unix_timestamp(now()), unix_timestamp(now()));
insert into wallpaper_category_tab (category, `desc`, del_flag, create_time, update_time) VALUES ('机车', '', 0, unix_timestamp(now()), unix_timestamp(now()));
insert into wallpaper_category_tab (category, `desc`, del_flag, create_time, update_time) VALUES ('创意', '', 0, unix_timestamp(now()), unix_timestamp(now()));
insert into wallpaper_category_tab (category, `desc`, del_flag, create_time, update_time) VALUES ('插画', '', 0, unix_timestamp(now()), unix_timestamp(now()));
insert into wallpaper_category_tab (category, `desc`, del_flag, create_time, update_time) VALUES ('运动', '', 0, unix_timestamp(now()), unix_timestamp(now()));
insert into wallpaper_category_tab (category, `desc`, del_flag, create_time, update_time) VALUES ('城市', '', 0, unix_timestamp(now()), unix_timestamp(now()));
insert into wallpaper_category_tab (category, `desc`, del_flag, create_time, update_time) VALUES ('情侣', '', 0, unix_timestamp(now()), unix_timestamp(now()));
insert into wallpaper_category_tab (category, `desc`, del_flag, create_time, update_time) VALUES ('明星', '', 0, unix_timestamp(now()), unix_timestamp(now()));
insert into wallpaper_category_tab (category, `desc`, del_flag, create_time, update_time) VALUES ('星空', '', 0, unix_timestamp(now()), unix_timestamp(now()));
insert into wallpaper_category_tab (category, `desc`, del_flag, create_time, update_time) VALUES ('动漫', '', 0, unix_timestamp(now()), unix_timestamp(now()));

-- ----------------------------
-- Table structure for wallpaper_type_tab
-- ----------------------------
DROP TABLE IF EXISTS `wallpaper_type_tab`;
CREATE TABLE `wallpaper_type_tab` (
    `id` bigint unsigned NOT NULL AUTO_INCREMENT COMMENT 'id',
    `tp` varchar(128) NOT NULL DEFAULT '' COMMENT 'type',
    `desc` varchar(512) NOT NULL DEFAULT '' COMMENT 'desc',
    `del_flag` char(1) NOT NULL DEFAULT '0' COMMENT 'del flag（0-normal 1-delete)',
    `create_time` bigint unsigned NOT NULL DEFAULT 0 COMMENT 'create time',
    `update_time` bigint unsigned NOT NULL DEFAULT 0 COMMENT 'update time',
    PRIMARY KEY (`id`),
    KEY `idx_tp` (`tp`) USING BTREE COMMENT 'idx_tp'
) ENGINE=InnoDB AUTO_INCREMENT=1 DEFAULT CHARSET=utf8mb4 COLLATE=utf8mb4_unicode_ci COMMENT='wallpaper type table';
insert into wallpaper_type_tab (tp, `desc`, del_flag, create_time, update_time) VALUES ('手机壁纸', '', 0, unix_timestamp(now()), unix_timestamp(now()));
insert into wallpaper_type_tab (tp, `desc`, del_flag, create_time, update_time) VALUES ('动态壁纸', '', 0, unix_timestamp(now()), unix_timestamp(now()));
insert into wallpaper_type_tab (tp, `desc`, del_flag, create_time, update_time) VALUES ('背景图', '', 0, unix_timestamp(now()), unix_timestamp(now()));
insert into wallpaper_type_tab (tp, `desc`, del_flag, create_time, update_time) VALUES ('头像', '', 0, unix_timestamp(now()), unix_timestamp(now()));
insert into wallpaper_type_tab (tp, `desc`, del_flag, create_time, update_time) VALUES ('表情包', '', 0, unix_timestamp(now()), unix_timestamp(now()));

BEGIN;
COMMIT;

SET FOREIGN_KEY_CHECKS = 1;