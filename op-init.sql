/* used for op init */
/* init product */
/* init product end */
/* init app */
/* init app end */
/* init department end */
/* init user with password sha256 */
INSERT INTO user_type (id, name, description, created_by) VALUES(1,'internal','内部组件使用',1);
INSERT INTO user_type (id, name, description, created_by) VALUES(2,'api','API使用',1);
INSERT INTO user_type (id, name, description, created_by) VALUES(3,'common','普通用户',1);

/* admin/admin */
INSERT INTO user (id, d_id, t_id, username, nickname, password, gender, phone, mail,
    token, status, created_by, updated_by ) VALUES ( 1, 0, 1, 'admin', 'admin', '8c6976e5b5410415bde908bd4dee15dfb167a9c873fc4bb8a81f6f2ab448a918',
    'dummy', '', '', '21232f297a57a5a743894a0e4a801fc3', 1, 1, 1 );

/* wagent/wagent@ */
INSERT INTO user (id, d_id, t_id, username, nickname, password, gender, phone, mail,
    token, status, created_by, updated_by ) VALUES ( 2, 0, 1, 'wagent', 'wagent','0495583f3c1eba0aa3c57f644a63f96bbc16f3e920c765afa51e7c947b3581b3',
    'dummy', '', '', '692d3946c7292592cfc33f86e2705d8c', 1, 1, 1 );

/* guarder/guarder@ */
INSERT INTO user (id, d_id, t_id, username, nickname, password, gender, phone, mail,
    token, status, created_by, updated_by ) VALUES ( 3, 0, 1, 'guarder', 'guarder', '477b5b5d0503cf3f4fc1f0a237d045880bac6ba773ae9721430e71ab31b6cc1b',
    'dummy', '', '', '277f1c4abd678b6512518f5824bc5d58', 1, 1, 1 );

/* init user end */
/* init department */
INSERT INTO department (id, name, status, description, created_by)
VALUES (1, 'admin', 1, 1,1);
/* init hosttype */
INSERT INTO host_type (id, name, description, created_by) VALUES (1, "office", "办公网机房", 1);

/* init appenv */
INSERT INTO app_env (id, name, description, created_by) VALUES (1, "测试", "测试环境", 1);
INSERT INTO app_env (id, name, description, created_by) VALUES (2, "预发", "预发环境", 1);
INSERT INTO app_env (id, name, description, created_by) VALUES (3, "线上", "线上环境", 1);
/* init role */
/* init role end */
