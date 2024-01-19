CREATE TABLE `office` (
    `id` INT(11) NOT NULL AUTO_INCREMENT,
    `name` VARCHAR(64) NOT NULL DEFAULT '' COMMENT '办公室名称' COLLATE 'utf8_general_ci',
    `address` VARCHAR(255) NOT NULL DEFAULT '' COMMENT '办公室地址' COLLATE 'utf8_general_ci',
    `created_at` TIMESTAMP NOT NULL DEFAULT CURRENT_TIMESTAMP,
    `updated_at` TIMESTAMP NOT NULL DEFAULT CURRENT_TIMESTAMP,
    `creator` VARCHAR(64) NOT NULL DEFAULT '' COMMENT '创建人' COLLATE 'utf8_general_ci',
    `updater` VARCHAR(64) NOT NULL DEFAULT '' COMMENT '更新人' COLLATE 'utf8_general_ci',
    `version` TIMESTAMP NOT NULL DEFAULT CURRENT_TIMESTAMP COMMENT '版本号',
    PRIMARY KEY (`id`) USING BTREE
)
COLLATE='utf8_general_ci'
ENGINE=InnoDB;