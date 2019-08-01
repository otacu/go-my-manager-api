CREATE TABLE `my-manager`.`tb_user` (
  `id` int(10) unsigned NOT NULL AUTO_INCREMENT COMMENT '自增主键',
  `user_name` varchar(100) DEFAULT NULL COMMENT '登录名',
  `password` varchar(100) DEFAULT NULL COMMENT '登录密码',
  `name` varchar(100) DEFAULT NULL COMMENT '用户名',
  `avatar` varchar(255) DEFAULT NULL COMMENT '头像',
  `roles` varchar(255) DEFAULT NULL COMMENT '角色列表',
  `created_at` timestamp NULL DEFAULT NULL COMMENT '创建时间',
  `updated_at` timestamp NULL DEFAULT NULL COMMENT '更新时间',
  `deleted_at` timestamp NULL DEFAULT NULL COMMENT '删除时间',
  PRIMARY KEY (`id`)
) ENGINE=InnoDB DEFAULT CHARSET=utf8mb4;
INSERT INTO `tb_user` (`id`, `user_name`, `password`, `name`, `avatar`, `roles`, `created_at`, `updated_at`, `deleted_at`) VALUES (1, 'admin', 'qwe123', 'Super Admin', 'https://wpimg.wallstcn.com/f778738c-e4f8-4870-b634-56703b4acafe.gif', 'admin', '2019-8-1 14:39:22', '2019-8-1 14:39:24', NULL);
INSERT INTO `tb_user` (`id`, `user_name`, `password`, `name`, `avatar`, `roles`, `created_at`, `updated_at`, `deleted_at`) VALUES (2, 'editor', '123456', 'Normal Editor', 'https://wpimg.wallstcn.com/f778738c-e4f8-4870-b634-56703b4acafe.gif', 'editor', '2019-8-1 14:40:59', '2019-8-1 14:41:01', NULL);
