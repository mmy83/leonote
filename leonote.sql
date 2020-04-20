create table leo_users (
    `id` int(11) not null auto_increment,
    `username` varchar(64) not null default "",
    `nickname` varchar(64) not null default "",
    `password` varchar(128) not null default "",
    `last_login_time` int(11) not null default 0,
    `isadmin` tinyint(1) not null default 0,
    `created_at` TIMESTAMP default CURRENT_TIMESTAMP,
    `updated_at` TIMESTAMP null ON UPDATE CURRENT_TIMESTAMP,
    `deleted_at` TIMESTAMP null,
    PRIMARY KEY (`id`),
    UNIQUE INDEX (`username`)
) ENGINE=InnoDB AUTO_INCREMENT=0 DEFAULT CHARSET=utf8 COMMENT='用户表';


create table leo_notebooks(
    `id` int(11) not null auto_increment,
    `parent_id` int(11) not null default 0,
    `user_id` int(11) not null default 0,
    `notebook_name` varchar(64) default null,
    `depth` tinyint(1) default 1,
    `note_num` int(11) default 0,
    `created_at` TIMESTAMP DEFAULT CURRENT_TIMESTAMP,
    `updated_at` TIMESTAMP null ON UPDATE CURRENT_TIMESTAMP,
    `deleted_at` TIMESTAMP null,
    PRIMARY KEY (`id`),
    INDEX (user_id)
) ENGINE=InnoDB AUTO_INCREMENT=0 DEFAULT CHARSET=utf8 COMMENT='笔记本表';


create table `leo_notes` (
    `id` int(11) not null auto_increment,
    `notebook_id` int(11) not null default 0,
    `user_id` int(11) not null default 0,
    `title` varchar(128) not null default "",
    `tags` varchar(128) not null default "",
    `content` text null ,
    `ismd` tinyint(1) not null default 0,
    `isshow` tinyint(1) not null default 0,
    `created_at` TIMESTAMP DEFAULT CURRENT_TIMESTAMP,
    `updated_at` TIMESTAMP null ON UPDATE CURRENT_TIMESTAMP,
    `deleted_at` TIMESTAMP null,
    PRIMARY KEY(`id`),
    INDEX (`user_id`)
) ENGINE=InnoDB AUTO_INCREMENT=0 DEFAULT CHARSET=utf8 COMMENT='笔记表';

create table `leo_tags` (
    `id` int(11) not null auto_increment,
    `user_id` int(11) not null default 0,
    `tag_name` varchar(64) not null default "",
    `count` int(10) not null default 0,
    `created_at` TIMESTAMP DEFAULT CURRENT_TIMESTAMP,
    `updated_at` TIMESTAMP null ON UPDATE CURRENT_TIMESTAMP,
    `deleted_at` TIMESTAMP null,
    PRIMARY KEY(`id`),
    INDEX(`user_id`)
) ENGINE=InnoDB AUTO_INCREMENT=0 DEFAULT CHARSET=utf8 COMMENT='标签表';

create table `leo_note_tags` (
    `id` int(11) not null auto_increment,
    `user_id` int(11) not null default 0,
    `note_id` int(11) not null default 0,
    `tag_id` int(11) not null default 0,
    `created_at` TIMESTAMP DEFAULT CURRENT_TIMESTAMP,
    `updated_at` TIMESTAMP null ON UPDATE CURRENT_TIMESTAMP,
    `deleted_at` TIMESTAMP null,
    PRIMARY KEY(`id`),
    INDEX(`user_id`)
) ENGINE=InnoDB AUTO_INCREMENT=0 DEFAULT CHARSET=utf8 COMMENT='笔记标签对应表';

