/*
分析需求_:
文章
用户
评论
种类
个人信息
*/

DROP TABLE IF EXISTS `users`;
CREATE TABLE `users`
(
  `id` bigint(20) UNSIGNED NOT NULL AUTO_INCREMENT,
  `username` varchar(20) CHARACTER SET utf8 COLLATE utf8_general_ci NOT NULL,
  `password` varchar(20) CHARACTER SET utf8 COLLATE utf8_general_ci NOT NULL,
  `role`
)