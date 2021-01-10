mysql> create table `users` (
`id` int NOT NULL AUTO_INCREMENT PRIMARY KEY COMMENT "ID",
`name` VARCHAR(100) NOT NULL COMMENT "user-name",
`token` VARCHAR(100) NOT NULL COMMENT "Token",
`created_at` datetime DEFAULT NULL COMMENT "created_at",
`updated_at` datetime DEFAULT NULL COMMENT "updatee_at"
) ENGINE=InnoDB DEFAULT CHARSET=utf8;

create table `user_characters` (
`id` int NOT NULL AUTO_INCREMENT PRIMARY KEY COMMENT "ID",
`character_id` int NOT NULL COMMENT "user-name",
`user_id` int NOT NULL COMMENT "Token",
`created_at` datetime DEFAULT NULL COMMENT "created_at",
`updated_at` datetime DEFAULT NULL COMMENT "updatee_at"
) ENGINE=InnoDB DEFAULT CHARSET=utf8;

create table `gachas` (
`id` int NOT NULL AUTO_INCREMENT PRIMARY KEY COMMENT "ID",
`name` VARCHAR(100) NOT NULL COMMENT "name",
`created_at` datetime DEFAULT NULL COMMENT "created_at",
`updated_at` datetime DEFAULT NULL COMMENT "updatee_at"
) ENGINE=InnoDB DEFAULT CHARSET=utf8;

create table `gacha_entries` (
`id` int NOT NULL AUTO_INCREMENT PRIMARY KEY COMMENT "ID",
`character_id` int NOT NULL COMMENT "user-name",
`user_id` int NOT NULL COMMENT "Token",
`weight` int NOT NULL COMMENT "user-name",
`created_at` datetime DEFAULT NULL COMMENT "created_at",
`updated_at` datetime DEFAULT NULL COMMENT "updatee_at"
) ENGINE=InnoDB DEFAULT CHARSET=utf8;

create table `characters` (
`id` int NOT NULL AUTO_INCREMENT PRIMARY KEY COMMENT "ID",
`name` VARCHAR(100) NOT NULL COMMENT "name",
`rarity` int NOT NULL COMMENT "rarity",
`created_at` datetime DEFAULT NULL COMMENT "created_at",
`updated_at` datetime DEFAULT NULL COMMENT "updatee_at"
) ENGINE=InnoDB DEFAULT CHARSET=utf8;
