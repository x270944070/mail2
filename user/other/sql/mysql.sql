create table user_info
(
    id          bigint auto_increment primary key comment '主键',
    user_name   varchar(32)  default ''                                            not null comment '用户名',
    email       varchar(64)  default ''                                            not null comment '邮箱',
    user_mobile varchar(16)  default ''                                            not null comment '手机号',
    password    varchar(128) default ''                                            not null comment '密码',
    nick_name   varchar(32)  default ''                                            not null comment '用户昵称',
    create_time datetime     default CURRENT_TIMESTAMP                             not null comment '数据创建时间',
    update_time DATETIME     default CURRENT_TIMESTAMP on update CURRENT_TIMESTAMP not null comment '数据更新时间'
) comment '用户-账户信息表';