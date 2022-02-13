/* Add ops platform mysql */
/* single - no dependency */
 CREATE DATABASE IF NOT EXISTS `op` DEFAULT CHARSET utf8mb4 DEFAULT COLLATE utf8mb4_unicode_ci;
 USE op;
/* add op user */
GRANT ALL PRIVILEGES ON  op.*  to 'op'@'127.0.0.1' IDENTIFIED BY 'liuliancao';
/* product table */
create TABLE IF NOT EXISTS `product` (
 `id` int(10) unsigned DEFAULT '0' COMMENT '产品ID',
 `name` varchar(100) DEFAULT '' COMMENT '产品名称',
 `description` varchar(100) DEFAULT '' COMMENT '产品描述',
 `status` int(3) unsigned DEFAULT '0' COMMENT '状态',
 `create_at` timestamp DEFAULT CURRENT_TIMESTAMP  COMMENT '创建时间',
 `update_at` timestamp ON UPDATE CURRENT_TIMESTAMP  COMMENT '修改时间',
  `delete_at` int(10) unsigned DEFAULT '0' COMMENT '删除时间',
 `create_by` varchar(50) DEFAULT '' COMMENT '创建人',
 `update_by` varchar(50) DEFAULT '' COMMENT '更新人'
) ENGINE=InnoDB DEFAULT CHARSET=utf8 COMMENT='产品表';

/* role table */
CREATE TABLE IF NOT EXISTS  `role` (
 `id` int(10) unsigned DEFAULT '0' COMMENT '角色ID',
 `name` varchar(100) DEFAULT '' COMMENT '角色名称',
 `permissions` varchar(100) DEFAULT '' COMMENT '角色权限',
 `status` int(3) unsigned DEFAULT '0' COMMENT '状态',
 `description` varchar(100) DEFAULT '' COMMENT '角色描述',
 `create_at` timestamp DEFAULT CURRENT_TIMESTAMP  COMMENT '创建时间',
 `update_at` timestamp ON UPDATE CURRENT_TIMESTAMP  COMMENT '修改时间',
  `delete_at` int(10) unsigned DEFAULT '0' COMMENT '删除时间',
 `create_by` varchar(50) DEFAULT '' COMMENT '创建人',
 `update_by` varchar(50) DEFAULT '' COMMENT '更新人'
) ENGINE=InnoDB DEFAULT CHARSET=utf8 COMMENT='角色表';

/* department table */
CREATE TABLE IF NOT EXISTS  `department` (
 `id` int(10) unsigned DEFAULT '0' COMMENT '部门ID',
 `name` varchar(100) DEFAULT '' COMMENT '部门名称',
 `status` int(3) unsigned DEFAULT '0' COMMENT '状态',
 `description` varchar(100) DEFAULT '' COMMENT '部门描述',
 `create_at` timestamp DEFAULT CURRENT_TIMESTAMP  COMMENT '创建时间',
 `update_at` timestamp ON UPDATE CURRENT_TIMESTAMP  COMMENT '修改时间',
  `delete_at` int(10) unsigned DEFAULT '0' COMMENT '删除时间',
 `create_by` varchar(50) DEFAULT '' COMMENT '创建人',
 `update_by` varchar(50) DEFAULT '' COMMENT '更新人'
) ENGINE=InnoDB DEFAULT CHARSET=utf8 COMMENT='部门表';

/* sshkey table */
CREATE TABLE IF NOT EXISTS  `sshkey` (
 `id` int(10) unsigned DEFAULT '0' COMMENT 'keyID',
 `u_id` int(10) unsigned DEFAULT '0' COMMENT '用户ID',
 `name` varchar(100) DEFAULT '' COMMENT 'key名称',
 `key` varchar(500) DEFAULT '' COMMENT 'ssh pub key',
 `create_at` timestamp DEFAULT CURRENT_TIMESTAMP  COMMENT '创建时间',
 `update_at` timestamp ON UPDATE CURRENT_TIMESTAMP  COMMENT '修改时间',
 `delete_at` int(10) unsigned DEFAULT '0' COMMENT '删除时间'
) ENGINE=InnoDB DEFAULT CHARSET=utf8 COMMENT='ssh key 表';

CREATE TABLE IF NOT EXISTS  `cluster` (
 `id` int(10) unsigned DEFAULT '0' COMMENT '集群ID',
 `name` varchar(100) DEFAULT '' COMMENT '集群名称',
 `status` int(3) unsigned DEFAULT '0' COMMENT '状态',
 `description` varchar(100) DEFAULT '' COMMENT '集群描述',
 `create_at` timestamp DEFAULT CURRENT_TIMESTAMP  COMMENT '创建时间',
 `update_at` timestamp ON UPDATE CURRENT_TIMESTAMP  COMMENT '修改时间',
  `delete_at` int(10) unsigned DEFAULT '0' COMMENT '删除时间',
  `create_by` varchar(50) DEFAULT '' COMMENT '创建人',
  `update_by` varchar(50) DEFAULT '' COMMENT '更新人'
) ENGINE=InnoDB DEFAULT CHARSET=utf8 COMMENT='集群表';
/* single - nodependency end */
/* onetomany */
CREATE TABLE IF NOT EXISTS  `user` (
 `id` int(10) unsigned DEFAULT '0' COMMENT '用户ID',
 `d_id` int(10) unsigned DEFAULT '0' COMMENT '部门ID',
 `username` varchar(100) DEFAULT '' COMMENT '用户名称',
 `password` varchar(100) DEFAULT '' COMMENT '密码',
 `gender` varchar(10) DEFAULT '' COMMENT '性别',
 `phone` varchar(30) DEFAULT '' COMMENT '电话',
 `mail` varchar(50) DEFAULT '' COMMENT '邮件',
 `token` varchar(100) DEFAULT '' COMMENT 'token',
 `status` int(3) unsigned DEFAULT 0 COMMENT '状态',
 `create_at` timestamp DEFAULT CURRENT_TIMESTAMP  COMMENT '创建时间',
 `update_at` timestamp ON UPDATE CURRENT_TIMESTAMP  COMMENT '修改时间',
  `delete_at` int(10) unsigned DEFAULT '0' COMMENT '删除时间',
 `create_by` varchar(50) DEFAULT '' COMMENT '创建人',
 `update_by` varchar(50) DEFAULT '' COMMENT '更新人'
) ENGINE=InnoDB DEFAULT CHARSET=utf8 COMMENT='用户表';

create TABLE IF NOT EXISTS `app` (
 `id` int(10) unsigned DEFAULT '0' COMMENT '应用ID',
 `p_id` int(10) unsigned DEFAULT '0' COMMENT '产品ID',
 `name` varchar(100) DEFAULT '' COMMENT '应用名称',
 `status` int(3) unsigned DEFAULT '0' COMMENT '状态',
 `description` varchar(100) DEFAULT '' COMMENT '应用描述',
 `create_at` timestamp DEFAULT CURRENT_TIMESTAMP  COMMENT '创建时间',
 `update_at` timestamp ON UPDATE CURRENT_TIMESTAMP  COMMENT '修改时间',
 `delete_at` int(10) unsigned DEFAULT '0' COMMENT '删除时间',
 `create_by` varchar(50) DEFAULT '' COMMENT '创建人',
 `update_by` varchar(50) DEFAULT '' COMMENT '更新人'
) ENGINE=InnoDB DEFAULT CHARSET=utf8 COMMENT='应用表';

create TABLE IF NOT EXISTS `tag` (
 `id` int(10) unsigned DEFAULT '0' COMMENT 'tag ID',
 `name` varchar(100) DEFAULT '' COMMENT '应用名称',
 `description` varchar(100) DEFAULT '' COMMENT '标签描述',
 `create_at` timestamp DEFAULT CURRENT_TIMESTAMP  COMMENT '创建时间',
 `update_at` timestamp ON UPDATE CURRENT_TIMESTAMP  COMMENT '修改时间',
 `delete_at` int(10) unsigned DEFAULT '0' COMMENT '删除时间',
 `create_by` varchar(50) DEFAULT '' COMMENT '创建人',
 `update_by` varchar(50) DEFAULT '' COMMENT '更新人'
) ENGINE=InnoDB DEFAULT CHARSET=utf8 COMMENT='标签表';


create TABLE IF NOT EXISTS `host` (
 `id` int(10) unsigned DEFAULT '0' COMMENT 'host ID',
 `hostname` varchar(100) DEFAULT '' COMMENT 'host主机名',
  `ip` varchar(16) DEFAULT '' COMMENT 'host ip',
  `os` varchar(20) DEFAULT '' COMMENT '操作系统',
  `htype` varchar(30) DEFAULT '' COMMENT 'host种类',
  `status` int(3) DEFAULT '0' COMMENT '主机状态',
  `uptime` varchar(30) DEFAULT '' COMMENT '启动时间',
  `extras` varchar(3000) DEFAULT '' COMMENT '扩展配置',
 `description` varchar(100) DEFAULT '' COMMENT '主机描述',
 `create_at` timestamp DEFAULT CURRENT_TIMESTAMP  COMMENT '创建时间',
 `update_at` timestamp ON UPDATE CURRENT_TIMESTAMP  COMMENT '修改时间',
 `delete_at` int(10) unsigned DEFAULT '0' COMMENT '删除时间',
 `create_by` varchar(50) DEFAULT '' COMMENT '创建人',
 `update_by` varchar(50) DEFAULT '' COMMENT '更新人'
) ENGINE=InnoDB DEFAULT CHARSET=utf8 COMMENT='主机表';

/* onetomany end */
/* manytomany */
CREATE TABLE IF NOT EXISTS  `app_user` (
 `id` int(10) unsigned DEFAULT '0' COMMENT '应用用户关联表id',
 `a_id` int(10) unsigned DEFAULT '0' COMMENT '应用id',
 `u_id` int(10) unsigned DEFAULT '0' COMMENT '用户id',
 `create_at` timestamp DEFAULT CURRENT_TIMESTAMP  COMMENT '创建时间',
 `update_at` timestamp ON UPDATE CURRENT_TIMESTAMP  COMMENT '修改时间',
  `delete_at` int(10) unsigned DEFAULT '0' COMMENT '删除时间',
 `create_by` varchar(50) DEFAULT '' COMMENT '创建人',
 `update_by` varchar(50) DEFAULT '' COMMENT '更新人'

) ENGINE=InnoDB DEFAULT CHARSET=utf8 COMMENT='应用用户关联表';

CREATE TABLE IF NOT EXISTS  `role_user` (
 `id` int(10) unsigned DEFAULT '0' COMMENT '角色用户关联表id',
 `r_id` int(10) unsigned DEFAULT '0' COMMENT '角色id',
 `u_id` int(10) unsigned DEFAULT '0' COMMENT '用户id',
 `create_at` timestamp DEFAULT CURRENT_TIMESTAMP  COMMENT '创建时间',
 `update_at` timestamp ON UPDATE CURRENT_TIMESTAMP  COMMENT '修改时间',
  `delete_at` int(10) unsigned DEFAULT '0' COMMENT '删除时间',
 `create_by` varchar(50) DEFAULT '' COMMENT '创建人',
 `update_by` varchar(50) DEFAULT '' COMMENT '更新人'

) ENGINE=InnoDB DEFAULT CHARSET=utf8 COMMENT='角色用户关联表';

CREATE TABLE IF NOT EXISTS  `app_cluster` (
 `id` int(10) unsigned DEFAULT '0' COMMENT '应用集群关联表id',
 `a_id` int(10) unsigned DEFAULT '0' COMMENT '应用id',
 `c_id` int(10) unsigned DEFAULT '0' COMMENT '用户id',
 `create_at` timestamp DEFAULT CURRENT_TIMESTAMP  COMMENT '创建时间',
 `update_at` timestamp ON UPDATE CURRENT_TIMESTAMP  COMMENT '修改时间',
 `delete_at` int(10) unsigned DEFAULT '0' COMMENT '删除时间',
 `create_by` varchar(50) DEFAULT '' COMMENT '创建人',
 `update_by` varchar(50) DEFAULT '' COMMENT '更新人'
) ENGINE=InnoDB DEFAULT CHARSET=utf8 COMMENT='应用集群关联表';

CREATE TABLE IF NOT EXISTS  `cluster_host` (
 `id` int(10) unsigned DEFAULT '0' COMMENT '集群主机关联表id',
 `c_id` int(10) unsigned DEFAULT '0' COMMENT '集群id',
 `h_id` int(10) unsigned DEFAULT '0' COMMENT '用户id',
 `create_at` timestamp DEFAULT CURRENT_TIMESTAMP  COMMENT '创建时间',
 `update_at` timestamp ON UPDATE CURRENT_TIMESTAMP  COMMENT '修改时间',
  `delete_at` int(10) unsigned DEFAULT '0' COMMENT '删除时间',
 `create_by` varchar(50) DEFAULT '' COMMENT '创建人',
 `update_by` varchar(50) DEFAULT '' COMMENT '更新人'
) ENGINE=InnoDB DEFAULT CHARSET=utf8 COMMENT='集群主机关联表';

CREATE TABLE IF NOT EXISTS  `tag_host` (
 `id` int(10) unsigned DEFAULT '0' COMMENT 'tag主机关联表id',
 `t_id` int(10) unsigned DEFAULT '0' COMMENT 'tag id',
 `h_id` int(10) unsigned DEFAULT '0' COMMENT '用户id',
 `create_at` timestamp DEFAULT CURRENT_TIMESTAMP  COMMENT '创建时间',
 `update_at` timestamp ON UPDATE CURRENT_TIMESTAMP  COMMENT '修改时间',
  `delete_at` int(10) unsigned DEFAULT '0' COMMENT '删除时间',
 `create_by` varchar(50) DEFAULT '' COMMENT '创建人',
 `update_by` varchar(50) DEFAULT '' COMMENT '更新人'
) ENGINE=InnoDB DEFAULT CHARSET=utf8 COMMENT='tag主机关联表';

/* manytomany end */
