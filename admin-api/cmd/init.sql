insert into sys_user values(1,  2, 'admin', '若依', '00', 'ry@163.com', '15888888888', '0', '', '$2a$10$7JB720yubVSZvUI0rEqK/.VqGOZTH.ulu33dHOiBE8ByOhJIrdAu2', 1, '127.0.0.1', current_timestamp, '管理员', 1, 'admin', current_timestamp, '', null);
insert into sys_user values(2,  3, 'ry',    '若依', '00', 'ry@qq.com',  '15666666666', '0', '', '$2a$10$7JB720yubVSZvUI0rEqK/.VqGOZTH.ulu33dHOiBE8ByOhJIrdAu2', 1, '127.0.0.1', current_timestamp, '管理员', 1, 'admin', current_timestamp, '', null);


-- 初始化-部门表数据
-- ----------------------------
insert into sys_dept values(1,   0,   '0',        '若依科技',    0, '若依', '15888888888', 'ry@qq.com', 1, 1, 'admin', current_timestamp, '', null);
insert into sys_dept values(2,   1,   '0,1',      '深圳总公司',  1, '若依', '15888888888', 'ry@qq.com', 1, 1, 'admin', current_timestamp, '', null);
insert into sys_dept values(3,   1,   '0,1',      '长沙分公司',  2, '若依', '15888888888', 'ry@qq.com', 1, 1, 'admin', current_timestamp, '', null);
insert into sys_dept values(4,   2,   '0,1,2',    '研发部门',    1, '若依', '15888888888', 'ry@qq.com', 1, 1, 'admin', current_timestamp, '', null);
insert into sys_dept values(5,   2,   '0,1,2',    '市场部门',    2, '若依', '15888888888', 'ry@qq.com', 1, 1, 'admin', current_timestamp, '', null);
insert into sys_dept values(6,   2,   '0,1,2',    '测试部门',    3, '若依', '15888888888', 'ry@qq.com', 1, 1, 'admin', current_timestamp, '', null);
insert into sys_dept values(7,   2,   '0,1,2',    '财务部门',    4, '若依', '15888888888', 'ry@qq.com', 1, 1, 'admin', current_timestamp, '', null);
insert into sys_dept values(8,   2,   '0,1,2',    '运维部门',    5, '若依', '15888888888', 'ry@qq.com', 1, 1, 'admin', current_timestamp, '', null);
insert into sys_dept values(9,   3,   '0,1,3',    '市场部门',    1, '若依', '15888888888', 'ry@qq.com', 1, 1, 'admin', current_timestamp, '', null);
insert into sys_dept values(10,  3,   '0,1,3',    '财务部门',    2, '若依', '15888888888', 'ry@qq.com', 1, 1, 'admin', current_timestamp, '', null);

-- 初始化-岗位信息表数据
-- ----------------------------
insert into sys_post values(1, 'ceo',  '董事长',    1, 1, null, 'admin', current_timestamp, '', null);
insert into sys_post values(2, 'se',   '项目经理',  2, 1, null, 'admin', current_timestamp, '', null);
insert into sys_post values(3, 'hr',   '人力资源',  3, 1, null, 'admin', current_timestamp, '', null);
insert into sys_post values(4, 'user', '普通员工',  4, 1, null, 'admin', current_timestamp, '', null);


-- 初始化-角色信息表数据
-- ----------------------------
insert into sys_role values(1, '超级管理员',  'admin',  1, 1, true, true, 1, '超级管理员', 1, 'admin', current_timestamp, '', null);
insert into sys_role values(2, '普通角色',    'common', 2, 2, true, true, 1, '普通角色', 1, 'admin', current_timestamp, '', null);


-- 初始化-字典类型表数据
-- ----------------------------
insert into sys_dict_type values(1,  '用户性别', 'sys_user_sex',        1, '用户性别列表', 'admin', current_timestamp, '', null);
insert into sys_dict_type values(2,  '菜单状态', 'sys_show_hide',       1, '菜单状态列表', 'admin', current_timestamp, '', null);
insert into sys_dict_type values(3,  '系统开关', 'sys_normal_disable',  1, '系统开关列表', 'admin', current_timestamp, '', null);
insert into sys_dict_type values(4,  '任务状态', 'sys_job_status',      1, '任务状态列表', 'admin', current_timestamp, '', null);
insert into sys_dict_type values(5,  '任务分组', 'sys_job_group',       1, '任务分组列表', 'admin', current_timestamp, '', null);
insert into sys_dict_type values(6,  '系统是否', 'sys_yes_no',          1, '系统是否列表', 'admin', current_timestamp, '', null);
insert into sys_dict_type values(7,  '通知类型', 'sys_notice_type',     1, '通知类型列表', 'admin', current_timestamp, '', null);
insert into sys_dict_type values(8,  '通知状态', 'sys_notice_status',   1, '通知状态列表', 'admin', current_timestamp, '', null);
insert into sys_dict_type values(9,  '操作类型', 'sys_oper_type',       1, '操作类型列表', 'admin', current_timestamp, '', null);
insert into sys_dict_type values(10, '系统状态', 'sys_common_status',   1, '登录状态列表', 'admin', current_timestamp, '', null);


-- 初始化-字典数据表数据
-- ----------------------------
insert into sys_dict_data values(1,  1,  '男',       '0',       'sys_user_sex',        '',   '',        true,  1, '性别男',  'admin', current_timestamp, '', null);
insert into sys_dict_data values(2,  2,  '女',       '1',       'sys_user_sex',        '',   '',        false, 1, '性别女',  'admin', current_timestamp, '', null);
insert into sys_dict_data values(3,  3,  '未知',     '2',       'sys_user_sex',        '',   '',        false,  1,  '性别未知', 'admin', current_timestamp, '', null);
insert into sys_dict_data values(4,  1,  '显示',     '0',       'sys_show_hide',       '',   'primary', true,   1,  '显示菜单', 'admin', current_timestamp, '', null);
insert into sys_dict_data values(5,  2,  '隐藏',     '1',       'sys_show_hide',       '',   'danger',  false,  1,  '隐藏菜单', 'admin', current_timestamp, '', null);
insert into sys_dict_data values(6,  1,  '正常',     '0',       'sys_normal_disable',  '',   'primary', true,   1,  '正常状态', 'admin', current_timestamp, '', null);
insert into sys_dict_data values(7,  2,  '停用',     '1',       'sys_normal_disable',  '',   'danger',  false,  1,  '停用状态', 'admin', current_timestamp, '', null);
insert into sys_dict_data values(8,  1,  '正常',     '0',       'sys_job_status',      '',   'primary', true,   1,  '正常状态', 'admin', current_timestamp, '', null);
insert into sys_dict_data values(9,  2,  '暂停',     '1',       'sys_job_status',      '',   'danger',  false,  1,  '停用状态', 'admin', current_timestamp, '', null);
insert into sys_dict_data values(10, 1,  '默认',     'DEFAULT', 'sys_job_group',       '',   '',        true,   1,  '默认分组', 'admin', current_timestamp, '', null);
insert into sys_dict_data values(11, 2,  '系统',     'SYSTEM',  'sys_job_group',       '',   '',        false,  1,  '系统分组', 'admin', current_timestamp, '', null);
insert into sys_dict_data values(12, 1,  '是',       'Y',       'sys_yes_no',          '',   'primary', true,   1, '系统默认是', 'admin', current_timestamp, '', null);
insert into sys_dict_data values(13, 2,  '否',       'N',       'sys_yes_no',          '',   'danger',  false,  1, '系统默认否', 'admin', current_timestamp, '', null);
insert into sys_dict_data values(14, 1,  '通知',     '1',       'sys_notice_type',     '',   'warning', true,   1,  '通知', 'admin', current_timestamp, '', null);
insert into sys_dict_data values(15, 2,  '公告',     '2',       'sys_notice_type',     '',   'success', false,  1,  '公告', 'admin', current_timestamp, '', null);
insert into sys_dict_data values(16, 1,  '正常',     '0',       'sys_notice_status',   '',   'primary', true,   1, '正常状态', 'admin', current_timestamp, '', null);
insert into sys_dict_data values(17, 2,  '关闭',     '1',       'sys_notice_status',   '',   'danger',  false,  1, '关闭状态', 'admin', current_timestamp, '', null);
insert into sys_dict_data values(18, 99, '其他',     '0',       'sys_oper_type',       '',   'info',    false,  1, '其他操作', 'admin', current_timestamp, '', null);
insert into sys_dict_data values(19, 1,  '新增',     '1',       'sys_oper_type',       '',   'info',    false,  1, '新增操作', 'admin', current_timestamp, '', null);
insert into sys_dict_data values(20, 2,  '修改',     '2',       'sys_oper_type',       '',   'info',    false,  1, '修改操作', 'admin', current_timestamp, '', null);
insert into sys_dict_data values(21, 3,  '删除',     '3',       'sys_oper_type',       '',   'danger',  false,  1, '删除操作', 'admin', current_timestamp, '', null);
insert into sys_dict_data values(22, 4,  '授权',     '4',       'sys_oper_type',       '',   'primary', false,  1, '授权操作', 'admin', current_timestamp, '', null);
insert into sys_dict_data values(23, 5,  '导出',     '5',       'sys_oper_type',       '',   'warning', false,  1, '导出操作', 'admin', current_timestamp, '', null);
insert into sys_dict_data values(24, 6,  '导入',     '6',       'sys_oper_type',       '',   'warning', false,  1, '导入操作', 'admin', current_timestamp, '', null);
insert into sys_dict_data values(25, 7,  '强退',     '7',       'sys_oper_type',       '',   'danger',  false,  1, '强退操作',  'admin', current_timestamp, '', null);
insert into sys_dict_data values(26, 8,  '生成代码', '8',       'sys_oper_type',       '',   'warning', false,   1, '生成操作', 'admin', current_timestamp, '', null);
insert into sys_dict_data values(27, 9,  '清空数据', '9',       'sys_oper_type',       '',   'danger',  false,   1, '清空操作', 'admin', current_timestamp, '', null);
insert into sys_dict_data values(28, 1,  '成功',     '0',       'sys_common_status',   '',   'primary', false,  1, '正常状态', 'admin', current_timestamp, '', null);
insert into sys_dict_data values(29, 2,  '失败',     '1',       'sys_common_status',   '',   'danger',  false,  1, '停用状态', 'admin', current_timestamp, '', null);


-- 初始化-参数配置表数据
-- ----------------------------
insert into sys_setting values(1, '主框架页-默认皮肤样式名称',      'sys.index.skinName',            'skin-blue',     1, '蓝色 skin-blue、绿色 skin-green、紫色 skin-purple、红色 skin-red、黄色 skin-yellow', 'admin', current_timestamp, '', null);
insert into sys_setting values(2, '用户管理-账号初始密码',         'sys.user.initPassword',         '123456',        1, '初始化密码 123456', 'admin', current_timestamp, '', null);
insert into sys_setting values(3, '主框架页-侧边栏主题',           'sys.index.sideTheme',           'theme-dark',    1, '深色主题theme-dark，浅色主题theme-light', 'admin', current_timestamp, '', null);
insert into sys_setting values(4, '账号自助-验证码开关',           'sys.account.captchaEnabled',    'true',          1, '是否开启验证码功能（true开启，false关闭）', 'admin', current_timestamp, '', null);
insert into sys_setting values(5, '账号自助-是否开启用户注册功能',   'sys.account.registerUser',      'false',         1, '是否开启注册用户功能（true开启，false关闭）', 'admin', current_timestamp, '', null);
insert into sys_setting values(6, '用户登录-黑名单列表',           'sys.login.blackIPList',         '',              1, '设置登录IP黑名单限制，多个匹配项以;分隔，支持匹配（*通配、网段）', 'admin', current_timestamp, '', null);


-- 初始化-公告信息表数据
-- ----------------------------
insert into sys_notice values(1, '温馨提醒：2018-07-01 若依新版本发布啦', 2, '新版本内容', 1, '管理员', 'admin', current_timestamp, '', null);
insert into sys_notice values(2, '维护通知：2018-07-01 若依系统凌晨维护', 1, '维护内容',   1, '管理员', 'admin', current_timestamp, '', null);



-- 初始化-用户与岗位关联表数据
-- ----------------------------
insert into sys_user_post values (1, 1);
insert into sys_user_post values (2, 2);


-- 初始化 菜单
INSERT INTO sys_menu  VALUES (1, '系统首页', 0, 1, '/index', null, null, false, true, 'M', true, 1, null, 'lucky-shouye1', null, 'admin', '2023-07-28 07:59:34.297618 +00:00', '', null);
INSERT INTO sys_menu  VALUES (2, '系统管理', 0, 2, '/system', null, null, false, true, 'M', true, 1, null, 'lucky-shezhi1', null, 'admin', '2023-07-28 08:02:51.773475 +00:00', '', null);
INSERT INTO sys_menu  VALUES (3, '用户管理', 2, 1, '/system/user', null, null, false, true, 'C', true, 1, 'system:user', 'lucky-yonghu', null, 'admin', '2023-07-28 08:07:03.809457 +00:00', '', null);
INSERT INTO sys_menu  VALUES (4, '角色管理', 2, 2, '/system/role', null, null, false, true, 'C', true, 1, 'system:role', 'lucky-jiaose', null, 'admin', '2023-07-28 08:10:44.311937 +00:00', '', null);
INSERT INTO sys_menu  VALUES (5, '菜单管理', 2, 3, '/system/menu', null, null, false, true, 'C', true, 1, 'system:menu', 'lucky-caidan', null, 'admin', '2023-07-28 08:12:17.647120 +00:00', '', null);
INSERT INTO sys_menu  VALUES (6, '部门管理', 2, 4, '/system/dept', null, null, false, true, 'C', true, 1, 'system:dept', 'lucky-bumenguanli', null, 'admin', '2023-07-28 08:13:02.230920 +00:00', '', null);
INSERT INTO sys_menu  VALUES (7, '岗位管理', 2, 5, '/system/post', null, null, false, true, 'C', true, 1, 'system:post', 'lucky-gangwei', null, 'admin', '2023-07-28 08:13:58.618051 +00:00', '', null);
INSERT INTO sys_menu  VALUES (8, '字典管理', 2, 6, '/system/dict', null, null, false, true, 'C', true, 1, 'system:dict', 'lucky-zidianmuluguanli', null, 'admin', '2023-07-28 08:14:49.174098 +00:00', '', null);
INSERT INTO sys_menu  VALUES (9, '参数管理', 2, 7, '/system/param', null, null, false, true, 'C', true, 1, 'system:param', 'lucky-shujucanshu', null, 'admin', '2023-07-28 08:15:39.202023 +00:00', '', null);
INSERT INTO sys_menu  VALUES (10, '通知公告', 2, 8, '/system/inform', null, null, false, true, 'C', true, 1, 'system:inform', 'lucky-tongzhi', null, 'admin', '2023-07-28 08:16:24.368651 +00:00', '', null);
INSERT INTO sys_menu  VALUES (11, '日志管理', 0, 3, '/logger', null, null, false, true, 'M', true, 1, null, 'lucky-nav_icon_rzgl_spe', null, 'admin', '2023-07-28 08:17:18.044847 +00:00', '', null);
INSERT INTO sys_menu  VALUES (12, '登录日志', 11, 2, '/logger/login', null, null, false, true, 'C', true, 1, 'logger:login', 'lucky-denglurizhi', null, 'admin', '2023-07-28 08:24:07.345148 +00:00', '', null);
INSERT INTO sys_menu  VALUES (13, '系统监控', 0, 4, '/monitor', null, null, false, true, 'M', true, 1, null, 'lucky-jiankong', null, 'admin', '2023-07-28 08:25:03.676872 +00:00', '', null);
INSERT INTO sys_menu  VALUES (14, '在线用户', 13, 1, '/monitor/onlineUser', null, null, false, true, 'C', true, 1, 'monitor:onlineUser', 'lucky-zaixianyonghuguanli1', null, 'admin', '2023-07-28 08:25:36.849906 +00:00', '', null);
INSERT INTO sys_menu  VALUES (15, '定时任务', 13, 2, '/monitor/timeTask', null, null, false, true, 'C', true, 1, 'monitor:timeTask', 'lucky-dingshirenwuguanli', null, 'admin', '2023-07-28 08:26:14.120249 +00:00', '', null);
INSERT INTO sys_menu  VALUES (16, '服务器监控', 13, 3, '/monitor/server', null, null, false, true, 'C', true, 1, 'monitor:server', 'lucky-fuwuqijiankong', null, 'admin', '2023-07-28 08:27:20.483057 +00:00', '', null);
INSERT INTO sys_menu  VALUES (17, '缓冲监控', 13, 4, '/monitor/cache', null, null, false, true, 'C', true, 1, 'monitor:cache', 'lucky-huanchongfenxi', null, 'admin', '2023-07-28 08:27:55.814870 +00:00', '', null);
INSERT INTO sys_menu  VALUES (18, '缓冲列表', 13, 5, '/monitor/cacheList', null, null, false, true, 'C', true, 1, 'monitor:cacheList', 'lucky-cityworksjichugongnengtubiao-', null, 'admin', '2023-07-28 08:28:59.204073 +00:00', '', null);
INSERT INTO sys_menu  VALUES (19, '操作日志', 11, 4, '/logger/operate', '', '', false, true, 'C', true, 1, 'logger:operate', 'lucky-caozuorizhi', '', 'admin', '2023-07-28 08:23:19.839112 +00:00', 'admin', '2023-07-28 08:50:57.410301 +00:00');


-- 更新自增主键的下标
SELECT setval('sys_role_role_id_seq', (SELECT max(role_id) FROM sys_role));
SELECT setval('sys_user_user_id_seq', (SELECT max(user_id) FROM sys_user));
SELECT setval('sys_dept_dept_id_seq', (SELECT max(dept_id) FROM sys_dept));
SELECT setval('sys_menu_menu_id_seq', (SELECT max(menu_id) FROM sys_menu));
SELECT setval('sys_post_post_id_seq', (SELECT max(post_id) FROM sys_post));
SELECT setval('sys_dict_type_dict_id_seq', (SELECT max(dict_id) FROM sys_dict_type));
SELECT setval('sys_dict_data_data_id_seq', (SELECT max(data_id) FROM sys_dict_data));
SELECT setval('sys_setting_config_id_seq', (SELECT max(config_id) FROM sys_setting));
SELECT setval('sys_notice_notice_id_seq', (SELECT max(notice_id) FROM sys_notice));