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
    `name` varchar(128) NOT NULL COMMENT 'name',
    `image_url` varchar(256) NOT NULL DEFAULT '' COMMENT 'image url',
    `author` varchar(128) NOT NULL DEFAULT '' COMMENT 'author',
    `desc` varchar(512) NOT NULL DEFAULT '' COMMENT 'desc',
    `del_flag` char(1) NOT NULL DEFAULT '0' COMMENT 'del flag（0-normal 1-delete)',
    `create_time` bigint unsigned NOT NULL DEFAULT 0 COMMENT 'create time',
    `update_time` bigint unsigned NOT NULL DEFAULT 0 COMMENT 'update time',
    PRIMARY KEY (`id`),
    UNIQUE KEY `uniq_wid` (`wid`) USING BTREE COMMENT 'uniq_wid',
    KEY `idx_ctime` (`create_time`) USING BTREE COMMENT 'idx_ctime'
) ENGINE=InnoDB AUTO_INCREMENT=1 DEFAULT CHARSET=utf8mb4 COLLATE=utf8mb4_unicode_ci COMMENT='wallpaper table';

BEGIN;
COMMIT;

SET FOREIGN_KEY_CHECKS = 1;