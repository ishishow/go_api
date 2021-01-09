```
mysql -u golang-test-user -h localhost -D golang-test-database -p


```

mysql> create table `users` (
`id` int NOT NULL AUTO_INCREMENT PRIMARY KEY COMMENT "ID",
`name` VARCHAR(100) NOT NULL COMMENT "user-name",
`token` VARCHAR(100) NOT NULL COMMENT "Token",
`created_at` datetime DEFAULT NULL COMMENT "created_at",
`updated_at` datetime DEFAULT NULL COMMENT "updatee_at"
) ENGINE=InnoDB DEFAULT CHARSET=utf8;

curl -X POST "http://localhost:8080/gacha/draw" -H "accept: application/json" -H "x-token: ishishow" -H "Content-Type: application/json" -d "{ \"times\": 2}"

2f45369a-0f11-46ce-a02a-1905f0d5fb99

05e2af0a-764f-46cf-80ca-a3aa017cb133
