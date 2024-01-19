CREATE TABLE `employee` (
    `id` INT(11) NOT NULL AUTO_INCREMENT,
    `name` VARCHAR(64) NOT NULL DEFAULT '' COMMENT '办公室名称' COLLATE 'utf8_general_ci',
    `age` INT(11) NOT NULL DEFAULT '0' COMMENT '员工年龄',
    `company` VARCHAR(255) NOT NULL DEFAULT '' COMMENT '公司' COLLATE 'utf8_general_ci',
    `phone` VARCHAR(64) NOT NULL DEFAULT '' COMMENT '电话' COLLATE 'utf8_general_ci',
    `office_id` INT(64) NOT NULL DEFAULT '0' COMMENT '办公室id',
    `created_at` TIMESTAMP NOT NULL DEFAULT CURRENT_TIMESTAMP,
    `updated_at` TIMESTAMP NOT NULL DEFAULT CURRENT_TIMESTAMP,
    `creator` VARCHAR(64) NOT NULL DEFAULT '' COMMENT '创建人' COLLATE 'utf8_general_ci',
    `updater` VARCHAR(64) NOT NULL DEFAULT '' COMMENT '更新人' COLLATE 'utf8_general_ci',
    `version` TIMESTAMP NOT NULL DEFAULT CURRENT_TIMESTAMP COMMENT '版本号',
    PRIMARY KEY (`id`) USING BTREE
)
COLLATE='utf8_general_ci'
ENGINE=InnoDB;