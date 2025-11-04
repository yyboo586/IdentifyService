USE `exhibition-admin`;

CREATE TABLE IF NOT EXISTS `t_sms_code` (
    `id` BIGINT(20) NOT NULL AUTO_INCREMENT COMMENT 'ID',
    `phone` VARCHAR(20) NOT NULL COMMENT '手机号',
    `business_type` TINYINT(1) NOT NULL COMMENT '业务类型(1:验证码登录)',
    `code` VARCHAR(10) NOT NULL COMMENT '验证码',
    `status` TINYINT(1) NOT NULL DEFAULT 0 COMMENT '状态(0:未使用 1:已使用)',
    `created_at` BIGINT(20) NOT NULL COMMENT '创建时间',
    `expired_at` BIGINT(20) NOT NULL COMMENT '过期时间',
    `updated_at` BIGINT(20) NOT NULL COMMENT '更新时间',
    PRIMARY KEY (`id`),
    KEY `idx_phone_business_type` (`phone`, `business_type`)
) ENGINE = InnoDB DEFAULT COMMENT = '手机验证码';