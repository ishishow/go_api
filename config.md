


```
mysql -u golang-test-user -h localhost -D golang-test-database -p


```
mysql> create table `users` (
    -> `id` int NOT NULL AUTO_INCREMENT PRIMARY KEY COMMENT "ID",
    -> `name` VARCHAR(100) NOT NULL COMMENT "user-name",
    -> `token` VARCHAR(100) NOT NULL COMMENT "Token",
    -> `created_at` datetime DEFAULT NULL COMMENT "created_at",
    -> `updated_at` datetime DEFAULT NULL COMMENT "updatee_at"
    -> ) ENGINE=InnoDB DEFAULT CHARSET=utf8;
