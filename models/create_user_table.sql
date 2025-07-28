CREATE TABLE IF NOT EXISTS users (
    id BIGINT(20) PRIMARY KEY AUTO_INCREMENT,
    user_id BIGINT(20) NOT NULL comment '用户ID',
    username VARCHAR(64) COLLATE utf8mb4_bin NOT NULL comment '用户名',
    password VARCHAR(64) COLLATE utf8mb4_bin NOT NULL comment '密码',
    email VARCHAR(64) COLLATE utf8mb4_bin NOT NULL comment '邮箱',
    gender TINYINT(4) NOT NULL DEFAULT 0 COMMENT '性别：0-未知，1-男，2-女',
    role TINYINT DEFAULT 0 COMMENT '0-普通用户，1-管理员',
    `create_time` timestamp NULL DEFAULT CURRENT_TIMESTAMP comment '创建时间',
    `update_time` timestamp NULL DEFAULT CURRENT_TIMESTAMP ON UPDATE CURRENT_TIMESTAMP comment '更新时间',
    `delete_time` timestamp NULL DEFAULT NULL comment '删除时间',
    UNIQUE KEY `idx_username` (`username`) USING BTREE,
    UNIQUE KEY `idx_user_id` (`user_id`) USING BTREE
) ENGINE=InnoDB DEFAULT CHARSET=utf8mb4 COLLATE=utf8mb4_bin COMMENT='用户表';

CREATE TABLE IF NOT EXISTS categories(
    `id` BIGINT PRIMARY KEY AUTO_INCREMENT COMMENT '分类id',
    `name` VARCHAR(100) NOT NULL COMMENT '分类名称',
    `create_time` timestamp NULL DEFAULT CURRENT_TIMESTAMP comment '创建时间',
    `update_time` timestamp NULL DEFAULT CURRENT_TIMESTAMP ON UPDATE CURRENT_TIMESTAMP comment '更新时间',
    `delete_time` timestamp NULL DEFAULT NULL comment '删除时间',
    UNIQUE KEY `idx_categories_name` (`name`)
) ENGINE=InnoDB DEFAULT CHARSET=utf8mb4 COMMENT='分类表';

CREATE Table IF NOT EXISTS `notes`(
    `id` BIGINT(20) PRIMARY KEY AUTO_INCREMENT COMMENT '笔记id',
    `title` VARCHAR(255) NOT NULL COMMENT '笔记标题',
    `content_md` TEXT NOT NULL COMMENT 'Markdown格式的笔记内容',
    `content_html` TEXT NOT NULL COMMENT 'HTML格式的笔记内容(由Markdown渲染)',
    `status` TINYINT NOT NULL DEFAULT 0 COMMENT '0:草稿，1:已发布',
    `author_id` BIGINT NOT NULL COMMENT '作者ID(关联users表)',
    `category_id` BIGINT NOT NULL COMMENT '分类ID(关联categories表)',
    `create_time` timestamp NULL DEFAULT CURRENT_TIMESTAMP comment '创建时间',
    `update_time` timestamp NULL DEFAULT CURRENT_TIMESTAMP ON UPDATE CURRENT_TIMESTAMP comment '更新时间',
    `delete_time` timestamp NULL DEFAULT NULL comment '删除时间',
    INDEX `idx_author_id` (`author_id` ASC),
    INDEX `idx_category_id` (`category_id` ASC)
)ENGINE=InnoDB DEFAULT CHARSET=utf8mb4 COMMENT='笔记表';