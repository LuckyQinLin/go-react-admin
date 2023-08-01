create table sys_menu
(
    menu_id     bigserial
        primary key,
    menu_code   varchar(50) not null,
    parent_id   bigint       default 0,
    order_num   bigint       default 0,
    path        varchar(200) default ''::character varying,
    component   varchar(255) default NULL::character varying,
    query       varchar(255) default NULL::character varying,
    is_frame    boolean      default false,
    is_cache    boolean      default true,
    menu_type   char         default ''::bpchar,
    visible     boolean      default true,
    status      bigint       default 1,
    perms       varchar(100) default NULL::character varying,
    icon        varchar(100) default '#'::character varying,
    remark      varchar(500) default NULL::character varying,
    create_by   varchar(64)  default ''::character varying,
    create_time timestamp with time zone,
    update_by   varchar(64)  default ''::character varying,
    update_time timestamp with time zone
);

comment on column sys_menu.menu_id is '菜单id';

comment on column sys_menu.menu_code is '菜单名称';

comment on column sys_menu.parent_id is '父菜单id';

comment on column sys_menu.order_num is '显示顺序';

comment on column sys_menu.path is '路由地址';

comment on column sys_menu.component is '组件路径';

comment on column sys_menu.query is '路由参数';

comment on column sys_menu.is_frame is '是否外链(true:是 false:不是)';

comment on column sys_menu.is_cache is '是否缓冲(true:是 false:不是)';

comment on column sys_menu.menu_type is '菜单类型(M目录 C菜单 F按钮)';

comment on column sys_menu.visible is '显隐状态(true显示 false隐藏)';

comment on column sys_menu.status is '菜单状态(1正常 0停用)';

comment on column sys_menu.perms is '权限标识';

comment on column sys_menu.icon is '菜单图标';

comment on column sys_menu.remark is '备注';

comment on column sys_menu.create_by is '创建者';

comment on column sys_menu.create_time is '创建时间';

comment on column sys_menu.update_by is '更新者';

comment on column sys_menu.update_time is '更新时间';

alter table sys_menu
    owner to admin;

INSERT INTO public.sys_menu (menu_id, menu_code, parent_id, order_num, path, component, query, is_frame, is_cache, menu_type, visible, status, perms, icon, remark, create_by, create_time, update_by, update_time) VALUES (6, '系统首页', 0, 1, '/index', null, null, false, true, 'M', true, 1, null, 'lucky-shouye1', null, 'admin', '2023-07-28 07:59:34.297618 +00:00', '', null);
INSERT INTO public.sys_menu (menu_id, menu_code, parent_id, order_num, path, component, query, is_frame, is_cache, menu_type, visible, status, perms, icon, remark, create_by, create_time, update_by, update_time) VALUES (7, '系统管理', 0, 2, '/system', null, null, false, true, 'M', true, 1, null, 'lucky-shezhi1', null, 'admin', '2023-07-28 08:02:51.773475 +00:00', '', null);
INSERT INTO public.sys_menu (menu_id, menu_code, parent_id, order_num, path, component, query, is_frame, is_cache, menu_type, visible, status, perms, icon, remark, create_by, create_time, update_by, update_time) VALUES (8, '用户管理', 7, 1, '/system/user', null, null, false, true, 'C', true, 1, 'system:user', 'lucky-yonghu', null, 'admin', '2023-07-28 08:07:03.809457 +00:00', '', null);
INSERT INTO public.sys_menu (menu_id, menu_code, parent_id, order_num, path, component, query, is_frame, is_cache, menu_type, visible, status, perms, icon, remark, create_by, create_time, update_by, update_time) VALUES (9, '角色管理', 7, 2, '/system/role', null, null, false, true, 'C', true, 1, 'system:role', 'lucky-jiaose', null, 'admin', '2023-07-28 08:10:44.311937 +00:00', '', null);
INSERT INTO public.sys_menu (menu_id, menu_code, parent_id, order_num, path, component, query, is_frame, is_cache, menu_type, visible, status, perms, icon, remark, create_by, create_time, update_by, update_time) VALUES (10, '菜单管理', 7, 3, '/system/menu', null, null, false, true, 'C', true, 1, 'system:menu', 'lucky-caidan', null, 'admin', '2023-07-28 08:12:17.647120 +00:00', '', null);
INSERT INTO public.sys_menu (menu_id, menu_code, parent_id, order_num, path, component, query, is_frame, is_cache, menu_type, visible, status, perms, icon, remark, create_by, create_time, update_by, update_time) VALUES (11, '部门管理', 7, 4, '/system/dept', null, null, false, true, 'C', true, 1, 'system:dept', 'lucky-bumenguanli', null, 'admin', '2023-07-28 08:13:02.230920 +00:00', '', null);
INSERT INTO public.sys_menu (menu_id, menu_code, parent_id, order_num, path, component, query, is_frame, is_cache, menu_type, visible, status, perms, icon, remark, create_by, create_time, update_by, update_time) VALUES (12, '岗位管理', 7, 5, '/system/post', null, null, false, true, 'C', true, 1, 'system:post', 'lucky-gangwei', null, 'admin', '2023-07-28 08:13:58.618051 +00:00', '', null);
INSERT INTO public.sys_menu (menu_id, menu_code, parent_id, order_num, path, component, query, is_frame, is_cache, menu_type, visible, status, perms, icon, remark, create_by, create_time, update_by, update_time) VALUES (13, '字典管理', 7, 6, '/system/dict', null, null, false, true, 'C', true, 1, 'system:dict', 'lucky-zidianmuluguanli', null, 'admin', '2023-07-28 08:14:49.174098 +00:00', '', null);
INSERT INTO public.sys_menu (menu_id, menu_code, parent_id, order_num, path, component, query, is_frame, is_cache, menu_type, visible, status, perms, icon, remark, create_by, create_time, update_by, update_time) VALUES (14, '参数管理', 7, 7, '/system/param', null, null, false, true, 'C', true, 1, 'system:param', 'lucky-shujucanshu', null, 'admin', '2023-07-28 08:15:39.202023 +00:00', '', null);
INSERT INTO public.sys_menu (menu_id, menu_code, parent_id, order_num, path, component, query, is_frame, is_cache, menu_type, visible, status, perms, icon, remark, create_by, create_time, update_by, update_time) VALUES (15, '通知公告', 7, 8, '/system/inform', null, null, false, true, 'C', true, 1, 'system:inform', 'lucky-tongzhi', null, 'admin', '2023-07-28 08:16:24.368651 +00:00', '', null);
INSERT INTO public.sys_menu (menu_id, menu_code, parent_id, order_num, path, component, query, is_frame, is_cache, menu_type, visible, status, perms, icon, remark, create_by, create_time, update_by, update_time) VALUES (16, '日志管理', 0, 3, '/logger', null, null, false, true, 'M', true, 1, null, 'lucky-nav_icon_rzgl_spe', null, 'admin', '2023-07-28 08:17:18.044847 +00:00', '', null);
INSERT INTO public.sys_menu (menu_id, menu_code, parent_id, order_num, path, component, query, is_frame, is_cache, menu_type, visible, status, perms, icon, remark, create_by, create_time, update_by, update_time) VALUES (18, '登录日志', 16, 2, '/logger/login', null, null, false, true, 'C', true, 1, 'logger:login', 'lucky-denglurizhi', null, 'admin', '2023-07-28 08:24:07.345148 +00:00', '', null);
INSERT INTO public.sys_menu (menu_id, menu_code, parent_id, order_num, path, component, query, is_frame, is_cache, menu_type, visible, status, perms, icon, remark, create_by, create_time, update_by, update_time) VALUES (19, '系统监控', 0, 4, '/monitor', null, null, false, true, 'M', true, 1, null, 'lucky-jiankong', null, 'admin', '2023-07-28 08:25:03.676872 +00:00', '', null);
INSERT INTO public.sys_menu (menu_id, menu_code, parent_id, order_num, path, component, query, is_frame, is_cache, menu_type, visible, status, perms, icon, remark, create_by, create_time, update_by, update_time) VALUES (20, '在线用户', 19, 1, '/monitor/onlineUser', null, null, false, true, 'C', true, 1, 'monitor:onlineUser', 'lucky-zaixianyonghuguanli1', null, 'admin', '2023-07-28 08:25:36.849906 +00:00', '', null);
INSERT INTO public.sys_menu (menu_id, menu_code, parent_id, order_num, path, component, query, is_frame, is_cache, menu_type, visible, status, perms, icon, remark, create_by, create_time, update_by, update_time) VALUES (21, '定时任务', 19, 2, '/monitor/timeTask', null, null, false, true, 'C', true, 1, 'monitor:timeTask', 'lucky-dingshirenwuguanli', null, 'admin', '2023-07-28 08:26:14.120249 +00:00', '', null);
INSERT INTO public.sys_menu (menu_id, menu_code, parent_id, order_num, path, component, query, is_frame, is_cache, menu_type, visible, status, perms, icon, remark, create_by, create_time, update_by, update_time) VALUES (22, '服务器监控', 19, 3, '/monitor/server', null, null, false, true, 'C', true, 1, 'monitor:server', 'lucky-fuwuqijiankong', null, 'admin', '2023-07-28 08:27:20.483057 +00:00', '', null);
INSERT INTO public.sys_menu (menu_id, menu_code, parent_id, order_num, path, component, query, is_frame, is_cache, menu_type, visible, status, perms, icon, remark, create_by, create_time, update_by, update_time) VALUES (23, '缓冲监控', 19, 4, '/monitor/cache', null, null, false, true, 'C', true, 1, 'monitor:cache', 'lucky-huanchongfenxi', null, 'admin', '2023-07-28 08:27:55.814870 +00:00', '', null);
INSERT INTO public.sys_menu (menu_id, menu_code, parent_id, order_num, path, component, query, is_frame, is_cache, menu_type, visible, status, perms, icon, remark, create_by, create_time, update_by, update_time) VALUES (24, '缓冲列表', 19, 5, '/monitor/cacheList', null, null, false, true, 'C', true, 1, 'monitor:cacheList', 'lucky-cityworksjichugongnengtubiao-', null, 'admin', '2023-07-28 08:28:59.204073 +00:00', '', null);
INSERT INTO public.sys_menu (menu_id, menu_code, parent_id, order_num, path, component, query, is_frame, is_cache, menu_type, visible, status, perms, icon, remark, create_by, create_time, update_by, update_time) VALUES (17, '操作日志', 16, 4, '/logger/operate', '', '', false, true, 'C', true, 1, 'logger:operate', 'lucky-caozuorizhi', '', 'admin', '2023-07-28 08:23:19.839112 +00:00', 'admin', '2023-07-28 08:50:57.410301 +00:00');
