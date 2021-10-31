-- mysql ...
DROP TABLE IF EXISTS `user`;
CREATE TABLE `user`
(
  `id`         bigint(20) UNSIGNED NOT NULL AUTO_INCREMENT,
  `username`   varchar(255) CHARACTER SET utf8 COLLATE utf8_general_ci NOT NULL,
  `password`   varchar(255) CHARACTER SET utf8 COLLATE utf8_general_ci NOT NULL,
  `avatar`     varchar(255) NOT NULL,
  `created_at` datetime(3) NULL DEFAULT NULL,
  `updated_at` datetime(3) NULL DEFAULT NULL,
  `deleted_at` datetime(3) NULL DEFAULT NULL,
  PRIMARY KEY (`id`) USING BTREE
) ENGINE = InnoDB CHARACTER SET = utf8 COLLATE = utf8_general_ci ROW_FORMAT = Dynamic;
INSERT INTO `user`  VALUES (1, 'admin', '$2a$10$K3DRY94amW.0Shce.5FsmuDqjKPzQczll.s6fHA361yPJ0le1Igr.','https://encrypted-tbn0.gstatic.com/images?q=tbn:ANd9GcRmb_j0g0GkHwuF42L8gNpJtSrT10kRQyEMvQ&usqp=CAU','2021-10-22 17:05:14.764', NULL , NULL);
INSERT INTO `user`  VALUES (2, 'civy', '$2a$10$K3DRY94amW.0Shce.5FsmuDqjKPzQczll.s6fHA361yPJ0le1Igr.','https://encrypted-tbn0.gstatic.com/images?q=tbn:ANd9GcRmb_j0g0GkHwuF42L8gNpJtSrT10kRQyEMvQ&usqp=CAU','2021-10-24 17:05:14.764', NULL , NULL);

-- ----------------------------
-- Table structure for role
-- ----------------------------
DROP TABLE IF EXISTS `role`;
CREATE TABLE `role`
(
  `id`         bigint(20) UNSIGNED NOT NULL AUTO_INCREMENT,
  `name`       varchar(50) NOT NULL COMMENT '角色名字',
  `created_at` datetime(3) NULL DEFAULT NULL,
  `updated_at` datetime(3) NULL DEFAULT NULL,
  `deleted_at` datetime(3) NULL DEFAULT NULL,
  PRIMARY KEY (`id`) USING BTREE
) ENGINE = InnoDB CHARACTER SET = utf8 COLLATE = utf8_general_ci ROW_FORMAT = Dynamic;
-- ----------------------------
-- Records of role
-- ----------------------------
INSERT INTO `role` VALUES (1,'admin','2021-10-22 18:25:33.763', NULL , NULL);
INSERT INTO `role` VALUES (2,'user','2021-10-22 18:25:33.763', NULL , NULL);


-- ----------------------------
-- Table structure for user_role
-- ----------------------------
DROP TABLE IF EXISTS `user_role`;
CREATE TABLE `user_role`
(
  `id`         bigint(20) UNSIGNED NOT NULL AUTO_INCREMENT,
  `user_id`    bigint(20) UNSIGNED NOT NULL COMMENT '用户ID',
  `role_id`    bigint(20) UNSIGNED NOT NULL COMMENT '角色ID',
  `created_at` datetime(3) NULL DEFAULT NULL,
  `updated_at` datetime(3) NULL DEFAULT NULL,
  `deleted_at` datetime(3) NULL DEFAULT NULL,
  PRIMARY KEY (`id`) USING BTREE
) ENGINE = InnoDB CHARACTER SET = utf8 COLLATE = utf8_general_ci ROW_FORMAT = Dynamic;

-- ----------------------------
-- Records of user_role
-- ----------------------------
INSERT INTO `user_role`  VALUES (1, 1, 1,'2021-10-22 18:25:33.763', NULL , NULL);
INSERT INTO `user_role`  VALUES (1, 2, 2,'2021-10-22 18:25:33.763', NULL , NULL);


DROP TABLE IF EXISTS `menu`;
CREATE TABLE `menu`
(
  `id`         bigint(20) UNSIGNED NOT NULL AUTO_INCREMENT,
  `name`       varchar(50) NOT NULL COMMENT '资源名字',
  `path`       varchar(50) NOT NULL COMMENT '访问路径',
  `method`     varchar(50) NOT NULL COMMENT '资源请求方式',
  `created_at`  datetime(3) NULL DEFAULT NULL,
  `updated_at`  datetime(3) NULL DEFAULT NULL,
  `deleted_at`  datetime(3) NULL DEFAULT NULL,
  PRIMARY KEY (`id`) USING BTREE
);
INSERT INTO `menu` (id,name,path,method,created_at,updated_at,deleted_at) VALUES
(1,'获取用户信息','/v1/auth/userinfo','GET','2021-10-22 18:29:33.763', NULL , NULL),
(2, '获取用户列表', '/v1/user', 'GET', '2021-10-22 18:29:33.763', NULL , NULL),
(3, '获取单个用户', '/v1/user/:id', 'GET', '2021-10-22 18:29:33.763', NULL , NULL),
(4, '创建用户', '/v1/user/add', 'POST', '2021-10-22 18:29:33.763', NULL , NULL),
(5, '删除用户', '/v1/user/:id', 'DELETE', '2021-10-22 18:29:33.763', NULL , NULL),
(6, '修改用户', '/v1/user/:id', 'PUT', '2021-10-22 18:29:33.763', NULL , NULL),
(7, '查询所有菜单', '/v1/menu', 'GET', '2021-10-22 18:29:33.763', NULL , NULL),
(8, '查询单个菜单', '/v1/menu/:id', 'GET', '2021-10-22 18:29:33.763', NULL , NULL),
(9, '创建单个菜单', '/v1/menu/add', 'POST', '2021-10-22 18:29:33.763', NULL , NULL),
(10, '更新单个菜单', '/v1/menu/:id', 'PUT','2021-10-22 18:29:33.763', NULL , NULL),
(11, '删除单个菜单', '/v1/menu/:id', 'DELETE','2021-10-22 18:29:33.763', NULL , NULL),
(12, '查询所有角色', '/v1/role', 'GET','2021-10-22 18:29:33.763', NULL , NULL),
(13, '查询单个角色', '/v1/role/:id', 'GET','2021-10-22 18:29:33.763', NULL , NULL),
(14, '创建单个角色', '/v1/role/add', 'POST','2021-10-22 18:29:33.763', NULL , NULL),
(15, '更新单个角色', '/v1/roles/:id', 'PUT','2021-10-22 18:29:33.763', NULL , NULL),
(16, '删除单个角色', '/v1/roles/:id', 'DELETE','2021-10-22 18:29:33.763', NULL , NULL);


-- ----------------------------
-- Table structure for menu_role
-- ----------------------------
DROP TABLE IF EXISTS `menu_role`;
CREATE TABLE `menu_role`
(
  `id`         bigint(20) UNSIGNED NOT NULL AUTO_INCREMENT,
  `role_id`    bigint(20) UNSIGNED NOT NULL COMMENT '角色ID',
  `menu_id`    bigint(20) UNSIGNED NOT NULL COMMENT '菜单ID',
  `created_at` datetime(3) NULL DEFAULT NULL,
  `updated_at` datetime(3) NULL DEFAULT NULL,
  `deleted_at` datetime(3) NULL DEFAULT NULL,
  PRIMARY KEY (`id`) USING BTREE
) ENGINE = InnoDB CHARACTER SET = utf8 COLLATE = utf8_general_ci ROW_FORMAT = Dynamic;

-- ----------------------------
-- Records of role_menu
-- ----------------------------
INSERT INTO `menu_role` (id,role_id,menu_id,created_at,updated_at,deleted_at) VALUES
(1,1,1,'2021-10-22 18:29:33.763', NULL , NULL),
(2,1,2,'2021-10-22 18:29:33.763', NULL , NULL),
(3,1,3,'2021-10-22 18:29:33.763', NULL , NULL),


DROP TABLE IF EXISTS `article`;
CREATE TABLE `article`
(
  `id`          bigint(20) UNSIGNED NOT NULL AUTO_INCREMENT,
  `title`       varchar(255) NOT NULL,
  `content`     longtext CHARACTER SET utf8 COLLATE utf8_general_ci NULL,
  `category_id` bigint(20) UNSIGNED NOT NULL,
  `tag_id`      bigint(20) UNSIGNED NOT NULL,
  `user_id`     bigint(20) UNSIGNED NOT NULL,
  `created_at`  datetime(3) NULL DEFAULT NULL,
  `updated_at`  datetime(3) NULL DEFAULT NULL,
  `deleted_at`  datetime(3) NULL DEFAULT NULL,
  PRIMARY KEY (`id`) USING BTREE
) ENGINE = InnoDB CHARACTER SET = utf8 COLLATE = utf8_general_ci ROW_FORMAT = Dynamic;

INSERT INTO `article` VALUES (1,'Hello','Linux',1,1,1,'2021-10-22 18:29:33.763',NULL,NULL);
INSERT INTO `article` VALUES (2,' Be Patience','Stock for the long run.',1,1,1,'2021-10-22 18:29:33.763',NULL,NULL);
INSERT INTO `article` VALUES (3,' Keep Learning','The best way to keep learning.',1,1,1,'2021-10-22 18:29:33.763',NULL,NULL);

DROP TABLE IF EXISTS `category`;
CREATE TABLE `category`
(
  `id`            bigint(20) UNSIGNED NOT NULL AUTO_INCREMENT,
  `category_name` varchar(255) NOT NULL UNIQUE,
  `created_at`    datetime(3) NULL DEFAULT NULL,
  `updated_at`    datetime(3) NULL DEFAULT NULL,
  `deleted_at`    datetime(3) NULL DEFAULT NULL,
  PRIMARY KEY (`id`) USING BTREE
);
INSERT INTO `category` VALUES (1,'Technology','2021-10-22 18:29:33.763',NULL,NULL);
INSERT INTO `category` VALUES (2,'Humanities','2021-10-22 18:29:33.763',NULL,NULL);
INSERT INTO `category` VALUES (3,'Politics','2021-10-22 18:29:33.763',NULL,NULL);
INSERT INTO `category` VALUES (4,'Language','2021-10-22 18:29:33.763',NULL,NULL);

DROP TABLE IF EXISTS `tag`;
CREATE TABLE `tag`
(
  `id`         int(11) UNSIGNED NOT NULL AUTO_INCREMENT,
  `tag_name`   varchar(255) NOT NULL UNIQUE,
  `created_at` timestamp NULL,
  `updated_at` timestamp NULL,
  `deleted_at` timestamp NULL,
  PRIMARY KEY (`id`) USING BTREE
);
INSERT INTO `tag` VALUES  (1,'Learning Tip','2021-10-22 18:29:33.763',NULL,NULL);
INSERT INTO `tag` VALUES  (2,'Make','2021-10-22 18:29:33.763',NULL,NULL);
INSERT INTO `tag` VALUES  (3,'CryptoCurrency','2021-10-22 18:29:33.763',NULL,NULL);
INSERT INTO `tag` VALUES  (4,'DigitalCurrency','2021-10-22 18:29:33.763',NULL,NULL);


