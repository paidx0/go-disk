-- MySQL dump 10.13  Distrib 8.0.31, for macos12.6 (x86_64)
--
-- Host: 127.0.0.1    Database: cloud-disk
-- ------------------------------------------------------
-- Server version	8.0.31

/*!40101 SET @OLD_CHARACTER_SET_CLIENT=@@CHARACTER_SET_CLIENT */;
/*!40101 SET @OLD_CHARACTER_SET_RESULTS=@@CHARACTER_SET_RESULTS */;
/*!40101 SET @OLD_COLLATION_CONNECTION=@@COLLATION_CONNECTION */;
/*!50503 SET NAMES utf8mb4 */;
/*!40103 SET @OLD_TIME_ZONE=@@TIME_ZONE */;
/*!40103 SET TIME_ZONE='+00:00' */;
/*!40014 SET @OLD_UNIQUE_CHECKS=@@UNIQUE_CHECKS, UNIQUE_CHECKS=0 */;
/*!40014 SET @OLD_FOREIGN_KEY_CHECKS=@@FOREIGN_KEY_CHECKS, FOREIGN_KEY_CHECKS=0 */;
/*!40101 SET @OLD_SQL_MODE=@@SQL_MODE, SQL_MODE='NO_AUTO_VALUE_ON_ZERO' */;
/*!40111 SET @OLD_SQL_NOTES=@@SQL_NOTES, SQL_NOTES=0 */;

--
-- Table structure for table `repository_pool`
--

DROP TABLE IF EXISTS `repository_pool`;
/*!40101 SET @saved_cs_client     = @@character_set_client */;
/*!50503 SET character_set_client = utf8mb4 */;
CREATE TABLE `repository_pool` (
  `id` int unsigned NOT NULL AUTO_INCREMENT,
  `identity` varchar(36) CHARACTER SET utf8mb3 COLLATE utf8mb3_general_ci DEFAULT NULL COMMENT '随机编号',
  `hash` varchar(32) DEFAULT NULL COMMENT '文件的唯一标识',
  `name` varchar(255) CHARACTER SET utf8mb3 COLLATE utf8mb3_general_ci DEFAULT NULL COMMENT '文件名',
  `ext` varchar(30) CHARACTER SET utf8mb3 COLLATE utf8mb3_general_ci DEFAULT NULL COMMENT '文件后缀',
  `size` int DEFAULT NULL COMMENT '文件大小',
  `path` varchar(255) DEFAULT NULL COMMENT '文件路径',
  `created_at` datetime DEFAULT NULL,
  `updated_at` datetime DEFAULT NULL,
  `deleted_at` datetime DEFAULT NULL,
  PRIMARY KEY (`id`)
) ENGINE=InnoDB AUTO_INCREMENT=21 DEFAULT CHARSET=utf8mb3 COMMENT='存储仓库池';
/*!40101 SET character_set_client = @saved_cs_client */;

--
-- Dumping data for table `repository_pool`
--

LOCK TABLES `repository_pool` WRITE;
/*!40000 ALTER TABLE `repository_pool` DISABLE KEYS */;
INSERT INTO `repository_pool` VALUES (3,'4303eba0-984e-4d26-9ed4-631b5783d679','8ad7c63bedfa5e3e72bb307aab207f2c','a','.jpg',34101,'uploads/d54c9a8a-a160-476f-b8fb-344174aad53e.jpg','2022-11-28 12:35:34','2022-11-28 12:35:34',NULL),(4,'79858832-cc96-4d53-990d-e3c96fbed952','c12d06ff3f523b4c38ab996fa2abb6f4','c','.jpg',126213,'uploads/afa26ccc-1300-4f29-b87d-2826d2ead58c.jpg','2022-11-28 12:42:00','2022-11-28 12:42:00',NULL),(5,'2f2c5acf-cd01-46fb-a05a-412ab4699606','e97b40c130038228afa934c845fd2ecb','手工注入','.txt',4600,'uploads/72115212-843f-4bbb-8643-84cc2e92d210.txt','2022-11-28 12:43:22','2022-11-28 12:43:22',NULL),(6,'9b3d222e-920a-4caf-8d8a-a3cba30c5f31','d95f7988b047c2435e27098067e54634','dump','.rdb',88,'uploads/aed02b8b-2018-485e-a0be-d9233fe1c4af.rdb','2022-11-28 12:45:29','2022-11-28 12:45:29',NULL),(7,'48b7fd6c-1b9d-463f-976c-f78b3412b2dc','88fb232d3f595512ec75b6066ed23ec9','何逸辉 1520203795-实验一','.docx',987333,'uploads/02ffa5b5-928f-4cf6-8c36-21d4d9b1cfab.docx','2022-11-28 13:05:25','2022-11-28 13:05:25',NULL),(13,'c6827bbd-4eec-429c-b70f-653d4cb49450','6ec0c13157e3e30f7d6c6e4d12d7f076','本书配套的习题答案网址','.txt',37,'uploads/1a9e440d-10d4-4407-9710-1a0825ca4ea9.txt','2022-12-07 23:02:14','2022-12-07 23:02:14',NULL),(14,'6436daaf-1336-496b-b550-78394cc2d4dc','db8f83d778bff227dcc2526ca2b68887','16_','.png',603446,'uploads/abda1baa-e078-44a4-823f-c8173cc2e377.png','2022-12-08 02:37:10','2022-12-08 02:37:10',NULL),(15,'1695cc2f-c5d1-4815-a826-271ffeb91d29','1b44170962683c00a061b7121ed3afbb','密码学实验三','.docx',36998,'uploads/cccd56b3-6f27-48d3-b0c8-8fd1bc68ac4a.docx','2022-12-08 14:09:54','2022-12-08 14:09:54',NULL),(16,'9fe7ce16-0768-450c-942b-c6e0627e273e','a4e3e77bd4e0e08acce50f994a1a3196','16','.png',602180,'uploads/18786081-b94d-44f0-b309-0fce8705bbbd.png','2022-12-08 15:18:00','2022-12-08 15:18:00',NULL),(17,'1540055a-6e1a-4f9f-ae5b-11d96ae654a2','5c0282e6e7646c0329f9b93c85d8c673','17','.png',654148,'uploads/306e76eb-f0ab-4c87-a9f3-e1557d73fc49.png','2022-12-08 15:34:10','2022-12-08 15:34:10',NULL),(18,'060d1607-9f96-48ed-a794-f9b82b48815d','db046b011ed87b8fd6d8906cc6ef90cb','18','.png',617861,'uploads/a5a0f376-c9bf-4775-b6ad-f5c89df841e1.png','2022-12-08 16:15:38','2022-12-08 16:15:38',NULL),(19,'5307d8d1-77fb-4665-9b27-d3136a3df56e','4651a696373c8e7624e69cab1015878b','19','.png',540344,'uploads/574d5a5a-2454-469d-ae6a-2a6cab88f36a.png','2022-12-08 16:39:35','2022-12-08 16:39:35',NULL),(20,'bd6d1dbb-bda6-4a93-846b-1b6f9547ca28','6e3d593e9ddfe328ed3224e350f3e776','6','.png',623152,'uploads/1135b601-5972-4108-a8a0-bb736662f716.png','2022-12-08 17:23:47','2022-12-08 17:23:47',NULL);
/*!40000 ALTER TABLE `repository_pool` ENABLE KEYS */;
UNLOCK TABLES;

--
-- Table structure for table `share_basic`
--

DROP TABLE IF EXISTS `share_basic`;
/*!40101 SET @saved_cs_client     = @@character_set_client */;
/*!50503 SET character_set_client = utf8mb4 */;
CREATE TABLE `share_basic` (
  `id` int unsigned NOT NULL AUTO_INCREMENT,
  `identity` varchar(36) CHARACTER SET utf8mb3 COLLATE utf8mb3_general_ci DEFAULT NULL COMMENT '分享文件编号',
  `parent_id` int DEFAULT NULL COMMENT '父级ID',
  `user_identity` varchar(36) CHARACTER SET utf8mb3 COLLATE utf8mb3_general_ci DEFAULT NULL COMMENT '用户编号',
  `repository_identity` varchar(36) CHARACTER SET utf8mb3 COLLATE utf8mb3_general_ci DEFAULT NULL COMMENT '仓库池唯一标识',
  `expired_time` int DEFAULT NULL COMMENT '过期时间，0-永不失效',
  `click_num` int DEFAULT '0' COMMENT '点击次数',
  `created_at` datetime DEFAULT NULL,
  `updated_at` datetime DEFAULT NULL,
  `deleted_at` datetime DEFAULT NULL,
  PRIMARY KEY (`id`)
) ENGINE=InnoDB AUTO_INCREMENT=10 DEFAULT CHARSET=utf8mb3 COMMENT='文件分享池';
/*!40101 SET character_set_client = @saved_cs_client */;

--
-- Dumping data for table `share_basic`
--

LOCK TABLES `share_basic` WRITE;
/*!40000 ALTER TABLE `share_basic` DISABLE KEYS */;
INSERT INTO `share_basic` VALUES (4,'a2ae02e7-090e-4c5f-8a10-ba07e745c6ef',10,'b8433e0c-b318-4d1b-b511-2c2be29bc1c2','79858832-cc96-4d53-990d-e3c96fbed952',100,0,'2022-11-29 17:13:28','2022-11-29 17:13:28',NULL),(5,'85af2a8d-91c5-41c4-9c83-0c38c9b5be0c',10,'b8433e0c-b318-4d1b-b511-2c2be29bc1c2','48b7fd6c-1b9d-463f-976c-f78b3412b2dc',100,0,'2022-11-30 09:21:54','2022-11-30 09:21:54',NULL),(6,'3adc1f79-3e66-4217-a0c4-ac0b93d9bdb9',8,'e17423c1-fb9f-49f4-89b5-91614d614a35','2f2c5acf-cd01-46fb-a05a-412ab4699606',100,5,'2022-11-30 09:26:50','2022-11-30 09:26:50',NULL),(7,'295a1336-266a-426c-81ce-956eb4da6af4',8,'e17423c1-fb9f-49f4-89b5-91614d614a35','2f2c5acf-cd01-46fb-a05a-412ab4699606',100,0,'2022-11-30 09:27:08','2022-11-30 09:27:08',NULL),(8,'a2900acd-e58c-4998-8f62-55dc03ff52b9',8,'e17423c1-fb9f-49f4-89b5-91614d614a35','2f2c5acf-cd01-46fb-a05a-412ab4699606',100,0,'2022-12-07 23:25:05','2022-12-07 23:25:05',NULL),(9,'9cac0a5b-96b6-4fc5-a8d7-82f755dbcbc3',11,'e17423c1-fb9f-49f4-89b5-91614d614a35','c6827bbd-4eec-429c-b70f-653d4cb49450',100,4,'2022-12-07 23:27:22','2022-12-07 23:27:22',NULL);
/*!40000 ALTER TABLE `share_basic` ENABLE KEYS */;
UNLOCK TABLES;

--
-- Table structure for table `user_basic`
--

DROP TABLE IF EXISTS `user_basic`;
/*!40101 SET @saved_cs_client     = @@character_set_client */;
/*!50503 SET character_set_client = utf8mb4 */;
CREATE TABLE `user_basic` (
  `id` int unsigned NOT NULL AUTO_INCREMENT,
  `identity` varchar(36) CHARACTER SET utf8mb3 COLLATE utf8mb3_general_ci DEFAULT NULL COMMENT '随机编号',
  `name` varchar(60) DEFAULT NULL,
  `password` varchar(32) DEFAULT NULL,
  `email` varchar(100) DEFAULT NULL,
  `created_at` datetime DEFAULT NULL,
  `updated_at` datetime DEFAULT NULL,
  `deleted_at` datetime DEFAULT NULL,
  PRIMARY KEY (`id`)
) ENGINE=InnoDB AUTO_INCREMENT=13 DEFAULT CHARSET=utf8mb3 COMMENT='用户基本信息';
/*!40101 SET character_set_client = @saved_cs_client */;

--
-- Dumping data for table `user_basic`
--

LOCK TABLES `user_basic` WRITE;
/*!40000 ALTER TABLE `user_basic` DISABLE KEYS */;
INSERT INTO `user_basic` VALUES (1,'user_1','test','cc03e747a6afbbcbf8be7668acfebee5','123@qq.com',NULL,NULL,NULL),(2,'user_2','asd','bc575fe0557c046f433d556612b3e564','123@qq.com',NULL,NULL,NULL),(8,'0931c6db-cdef-4373-80b6-d0ab6d564fdd','paidxing0','827ccb0eea8a706c4c34a16891f84e7b','322798@qq.com','2022-11-27 21:48:16','2022-11-27 21:48:16',NULL),(10,'b8433e0c-b318-4d1b-b511-2c2be29bc1c2','test1','827ccb0eea8a706c4c34a16891f84e7b','11@qq.com','2022-11-29 03:13:36','2022-11-29 03:13:36',NULL),(11,'e17423c1-fb9f-49f4-89b5-91614d614a35','test2','827ccb0eea8a706c4c34a16891f84e7b','12@qq.com','2022-11-30 09:25:11','2022-11-30 09:25:11',NULL),(12,'64c19ea4-62e6-4f12-bdc3-38aac549fc40','test3','827ccb0eea8a706c4c34a16891f84e7b','3227055498@qq.com','2022-12-08 18:09:49','2022-12-08 18:09:49',NULL);
/*!40000 ALTER TABLE `user_basic` ENABLE KEYS */;
UNLOCK TABLES;

--
-- Table structure for table `user_repository`
--

DROP TABLE IF EXISTS `user_repository`;
/*!40101 SET @saved_cs_client     = @@character_set_client */;
/*!50503 SET character_set_client = utf8mb4 */;
CREATE TABLE `user_repository` (
  `id` int unsigned NOT NULL AUTO_INCREMENT,
  `identity` varchar(36) CHARACTER SET utf8mb3 COLLATE utf8mb3_general_ci DEFAULT NULL COMMENT '随机编号',
  `user_identity` varchar(36) CHARACTER SET utf8mb3 COLLATE utf8mb3_general_ci DEFAULT NULL COMMENT '用户编号',
  `parent_id` int DEFAULT NULL COMMENT '父级ID',
  `repository_identity` varchar(36) CHARACTER SET utf8mb3 COLLATE utf8mb3_general_ci DEFAULT NULL COMMENT '仓库文件编号',
  `ext` varchar(255) CHARACTER SET utf8mb3 COLLATE utf8mb3_general_ci DEFAULT NULL COMMENT '文件后缀',
  `name` varchar(255) CHARACTER SET utf8mb3 COLLATE utf8mb3_general_ci DEFAULT NULL COMMENT '文件名',
  `created_at` datetime DEFAULT NULL,
  `updated_at` datetime DEFAULT NULL,
  `deleted_at` datetime DEFAULT NULL,
  PRIMARY KEY (`id`)
) ENGINE=InnoDB AUTO_INCREMENT=28 DEFAULT CHARSET=utf8mb3 COMMENT='用户绑定仓库';
/*!40101 SET character_set_client = @saved_cs_client */;

--
-- Dumping data for table `user_repository`
--

LOCK TABLES `user_repository` WRITE;
/*!40000 ALTER TABLE `user_repository` DISABLE KEYS */;
INSERT INTO `user_repository` VALUES (11,'9980d7c1-8bdf-4bcb-b834-a53aa6c06eac','0931c6db-cdef-4373-80b6-d0ab6d564fdd',8,'4303eba0-984e-4d26-9ed4-631b5783d679','.jpg','a','2022-11-29 03:14:43','2022-11-29 03:14:43',NULL),(12,'f7e832f4-5454-467a-9f53-8762e8e45567','b8433e0c-b318-4d1b-b511-2c2be29bc1c2',10,'79858832-cc96-4d53-990d-e3c96fbed952','.jpg','a','2022-11-29 03:14:51','2022-11-29 03:14:51','2022-11-29 14:55:51'),(13,'c60cce75-ed6c-4ef8-865b-c47087887dc9','e17423c1-fb9f-49f4-89b5-91614d614a35',11,'2f2c5acf-cd01-46fb-a05a-412ab4699606','.jpg','testUpdate','2022-11-29 03:15:11','2022-11-29 15:45:48',NULL),(14,'bef4ea7d-de4b-4627-a6d6-18aa4aa45da1','e17423c1-fb9f-49f4-89b5-91614d614a35',11,'2f2c5acf-cd01-46fb-a05a-412ab4699606','.jpg','a','2022-11-29 03:25:26','2022-11-29 15:48:07',NULL),(15,'9767ff96-b314-4cc8-a11e-516dc1f2a8ae','e17423c1-fb9f-49f4-89b5-91614d614a35',11,'48b7fd6c-1b9d-463f-976c-f78b3412b2dc','.jpg','a','2022-11-29 03:26:09','2022-11-29 15:46:36',NULL),(16,'6bde3a97-8c66-4dc9-99cf-e4028ae68402','b8433e0c-b318-4d1b-b511-2c2be29bc1c2',10,'48b7fd6c-1b9d-463f-976c-f78b3412b2dc','.jpg','a','2022-11-29 03:34:54','2022-11-29 03:34:54',NULL),(17,'eb293b50-3cf4-4fa1-8ce9-732a13f99149','e17423c1-fb9f-49f4-89b5-91614d614a35',11,'2f2c5acf-cd01-46fb-a05a-412ab4699606','.txt','手工注入','2022-12-03 23:37:02','2022-12-03 23:37:02',NULL),(18,'7591682a-acee-42dd-8963-1c1e0f084d14','e17423c1-fb9f-49f4-89b5-91614d614a35',11,'2f2c5acf-cd01-46fb-a05a-412ab4699606','.txt','手工注入','2022-12-04 00:27:53','2022-12-04 00:27:53',NULL),(19,'e86cc444-7f54-4059-b286-beacb37aa66f','0931c6db-cdef-4373-80b6-d0ab6d564fdd',8,'1b657cb6-2b99-4fc8-abef-3c4084ef9084','.txt','本书配套的习题答案网址','2022-12-07 22:54:07','2022-12-07 23:18:04',NULL),(20,'5aaccdb4-de57-465d-b3a5-b73696a145b0','e17423c1-fb9f-49f4-89b5-91614d614a35',11,'c6827bbd-4eec-429c-b70f-653d4cb49450','.txt','本书配套的习题答案网址','2022-12-07 23:02:14','2022-12-07 23:02:14',NULL),(21,'9361e91d-dbab-403a-882a-fd4c72b70814','e17423c1-fb9f-49f4-89b5-91614d614a35',11,'6436daaf-1336-496b-b550-78394cc2d4dc','.png','16_','2022-12-08 02:37:10','2022-12-08 02:37:10',NULL),(22,'bb5d8a2b-a13d-4e1b-9f87-a8dbbb426424','e17423c1-fb9f-49f4-89b5-91614d614a35',11,'1695cc2f-c5d1-4815-a826-271ffeb91d29','.docx','密码学实验三','2022-12-08 14:09:54','2022-12-08 14:09:54',NULL),(23,'6df402df-ff6f-4105-ad2f-26eef3818324','e17423c1-fb9f-49f4-89b5-91614d614a35',11,'9fe7ce16-0768-450c-942b-c6e0627e273e','.png','16','2022-12-08 15:18:00','2022-12-08 15:18:00',NULL),(24,'8c238dbe-f49b-40c1-8f42-7df88b89a0d8','e17423c1-fb9f-49f4-89b5-91614d614a35',11,'1540055a-6e1a-4f9f-ae5b-11d96ae654a2','.png','17','2022-12-08 15:34:10','2022-12-08 15:34:10',NULL),(25,'e5747b31-8005-4c0a-81b1-28f93b6fc555','e17423c1-fb9f-49f4-89b5-91614d614a35',11,'060d1607-9f96-48ed-a794-f9b82b48815d','.png','18','2022-12-08 16:15:38','2022-12-08 16:15:38',NULL),(26,'2e6e6b7c-3419-4d36-a16f-3ac4a64ee5e3','e17423c1-fb9f-49f4-89b5-91614d614a35',11,'5307d8d1-77fb-4665-9b27-d3136a3df56e','.png','19','2022-12-08 16:39:35','2022-12-08 16:39:35',NULL),(27,'23e766cd-f790-45cf-bd57-c713bc813da1','e17423c1-fb9f-49f4-89b5-91614d614a35',11,'bd6d1dbb-bda6-4a93-846b-1b6f9547ca28','.png','6','2022-12-08 17:23:47','2022-12-08 17:23:47',NULL);
/*!40000 ALTER TABLE `user_repository` ENABLE KEYS */;
UNLOCK TABLES;

--
-- Dumping routines for database 'cloud-disk'
--
/*!40103 SET TIME_ZONE=@OLD_TIME_ZONE */;

/*!40101 SET SQL_MODE=@OLD_SQL_MODE */;
/*!40014 SET FOREIGN_KEY_CHECKS=@OLD_FOREIGN_KEY_CHECKS */;
/*!40014 SET UNIQUE_CHECKS=@OLD_UNIQUE_CHECKS */;
/*!40101 SET CHARACTER_SET_CLIENT=@OLD_CHARACTER_SET_CLIENT */;
/*!40101 SET CHARACTER_SET_RESULTS=@OLD_CHARACTER_SET_RESULTS */;
/*!40101 SET COLLATION_CONNECTION=@OLD_COLLATION_CONNECTION */;
/*!40111 SET SQL_NOTES=@OLD_SQL_NOTES */;

-- Dump completed on 2022-12-08 18:17:17
