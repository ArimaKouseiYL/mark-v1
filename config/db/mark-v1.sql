-- `mark-v1`.users definition

CREATE TABLE `users` (
                         `id` bigint DEFAULT NULL,
                         `user_name` varchar(100) CHARACTER SET utf8mb4 COLLATE utf8mb4_0900_ai_ci DEFAULT NULL,
                         `age` int DEFAULT NULL,
                         `address` varchar(100) DEFAULT NULL,
                         `password` varchar(100) DEFAULT NULL,
                         `phone` varchar(100) DEFAULT NULL
) ENGINE=InnoDB DEFAULT CHARSET=utf8mb4 COLLATE=utf8mb4_0900_ai_ci;

