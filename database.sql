create table `users` (
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
`gacha_id` int NOT NULL COMMENT "gacha_id",
`character_id` int NOT NULL COMMENT "character_id",
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

INSERT INTO gachas
  (id, name)
VALUES
  (1, "普通ガチャ"),
  (2, "最強ガチャ");

INSERT INTO gacha_entries
  (id, gacha_id, character_id, weight)
VALUES
  (1, 1, 1, 100),
  (2, 1, 2, 100),
  (3, 1, 3, 100),
  (4, 1, 4, 300),
  (5, 1, 5, 300),
  (6, 1, 6, 500),
  (7, 1, 7, 500),
  (8, 1, 8, 900),
  (9, 1, 9, 900),
  (10, 1, 10, 1000);

INSERT INTO characters
  (id, name, rarity)
VALUES
  (1, "FORTRAN", 5),
  (2, "Rust", 5),
  (3, "Fluter", 5),
  (4, "Go", 4),
  (5, "C++", 4),
  (6, "Python", 4),
  (7, "ruby", 4),
  (8, "css", 4),
  (9, "html", 4),
  (0, "nihongo", 4);



