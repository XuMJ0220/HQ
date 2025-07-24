CREATE TABLE IF NOT EXISTS users (
    id BIGINT(20) PRIMARY KEY AUTO_INCREMENT,
    user_id BIGINT(20) NOT NULL,
    username VARCHAR(64) COLLATE utf8mb4_general_ci NOT NULL,
    password VARCHAR(64) COLLATE utf8mb4_general_ci NOT NULL,
    email VARCHAR(64) COLLATE utf8mb4_general_ci NOT NULL,
    gender TINYINT(4) NOT NULL DEFAULT '0',
    `create_time` timestamp NULL DEFAULT CURRENT_TIMESTAMP,
    `update_time` timestamp NULL DEFAULT CURRENT_TIMESTAMP ON UPDATE CURRENT_TIMESTAMP,
    `delete_time` timestamp NULL DEFAULT NULL,
    UNIQUE KEY `idx_username` (`username`) USING BTREE,
    UNIQUE KEY `idx_user_id` (`user_id`) USING BTREE
) ENGINE=InnoDB DEFAULT CHARSET=utf8mb4 COLLATE=utf8mb4_general_ci COMMENT='用户表';
