-- +goose Up
-- +goose StatementBegin
SELECT 'up SQL query';
-- +goose StatementEnd
CREATE TABLE f_alarm(
                        id VARCHAR(32) NOT NULL,
                        dn VARCHAR(255) NOT NULL,
                        title VARCHAR(255) NOT NULL,
                        type int NOT NULL,
                        description text NOT NULL,
                        status int2 NOT NULL,
                        level int2 NOT NULL,
                        event_time timestamp NOT NULL,
                        create_at timestamp NOT NULL,
                        updated_at timestamp NOT NULL,
                        object_id VARCHAR(255) NOT NULL,
                        object_name VARCHAR(255) NOT NULL,
                        location VARCHAR(255) NOT NULL,
                        total_count int NOT NULL,
                        clear_time timestamp,
                        ack_flag int2,
                        archive_flag int2,
                        ack_time timestamp,
                        clear_user VARCHAR(255),
                        ack_user VARCHAR(255),
                        archive_time timestamp,
                        archive_user VARCHAR(255),
                        extra_data text,
                        PRIMARY KEY (id)
);

COMMENT ON TABLE f_alarm IS '告警表';
COMMENT ON COLUMN f_alarm.id IS 'id';
COMMENT ON COLUMN f_alarm.dn IS 'dn';
COMMENT ON COLUMN f_alarm.title IS '告警标题';
COMMENT ON COLUMN f_alarm.type IS '告警类型:1:系统,2:通讯';
COMMENT ON COLUMN f_alarm.description IS '告警描述';
COMMENT ON COLUMN f_alarm.status IS '0:已清除告警,1:未清除告警';
COMMENT ON COLUMN f_alarm.level IS '级别,1:严重,2:一般,3:警告,4:描述';
COMMENT ON COLUMN f_alarm.event_time IS '告警时间';
COMMENT ON COLUMN f_alarm.create_at IS '创建时间';
COMMENT ON COLUMN f_alarm.updated_at IS '更新时间';
COMMENT ON COLUMN f_alarm.object_id IS '告警对象id';
COMMENT ON COLUMN f_alarm.object_name IS '告警对象名称';
COMMENT ON COLUMN f_alarm.location IS '告警位置';
COMMENT ON COLUMN f_alarm.total_count IS '发生次数';
COMMENT ON COLUMN f_alarm.clear_time IS '清除时间';
COMMENT ON COLUMN f_alarm.ack_flag IS '是否确认';
COMMENT ON COLUMN f_alarm.ack_time IS '确认时间';
COMMENT ON COLUMN f_alarm.clear_user IS '清除人';
COMMENT ON COLUMN f_alarm.ack_user IS '确认人';
COMMENT ON COLUMN f_alarm.archive_flag IS '是否归档，归档后才转移到历史告警';
COMMENT ON COLUMN f_alarm.archive_time IS '归档时间';
COMMENT ON COLUMN f_alarm.archive_user IS '归档人id';


CREATE UNIQUE INDEX idx_f_alarm_dn ON f_alarm (dn);





CREATE TABLE f_history_alarm(
                                id VARCHAR(32) NOT NULL,
                                dn VARCHAR(255) NOT NULL,
                                title VARCHAR(255) NOT NULL,
                                type int NOT NULL,
                                description text NOT NULL,
                                status int2 NOT NULL,
                                level int2 NOT NULL,
                                event_time timestamp NOT NULL,
                                create_at timestamp NOT NULL,
                                updated_at timestamp NOT NULL,
                                object_id VARCHAR(255) NOT NULL,
                                object_name VARCHAR(255) NOT NULL,
                                location VARCHAR(255) NOT NULL,
                                total_count int NOT NULL,
                                clear_time timestamp,
                                ack_flag int2,
                                archive_flag int2,
                                ack_time timestamp,
                                clear_user VARCHAR(255),
                                ack_user VARCHAR(255),
                                archive_time timestamp,
                                archive_user VARCHAR(255),
                                extra_data text,
                                PRIMARY KEY (id)
);

COMMENT ON TABLE f_history_alarm IS '告警表';
COMMENT ON COLUMN f_history_alarm.id IS 'id';
COMMENT ON COLUMN f_history_alarm.dn IS 'dn';
COMMENT ON COLUMN f_history_alarm.title IS '告警标题';
COMMENT ON COLUMN f_history_alarm.type IS '告警类型';
COMMENT ON COLUMN f_history_alarm.description IS '告警描述';
COMMENT ON COLUMN f_history_alarm.status IS '0:已清除告警,1:未清除告警';
COMMENT ON COLUMN f_history_alarm.level IS '级别';
COMMENT ON COLUMN f_history_alarm.event_time IS '告警时间';
COMMENT ON COLUMN f_history_alarm.create_at IS '创建时间';
COMMENT ON COLUMN f_history_alarm.updated_at IS '更新时间';
COMMENT ON COLUMN f_history_alarm.object_id IS '告警对象id';
COMMENT ON COLUMN f_history_alarm.object_name IS '告警对象名称';
COMMENT ON COLUMN f_history_alarm.location IS '告警位置';
COMMENT ON COLUMN f_history_alarm.total_count IS '发生次数';
COMMENT ON COLUMN f_history_alarm.clear_time IS '清除时间';
COMMENT ON COLUMN f_history_alarm.ack_flag IS '是否确认';
COMMENT ON COLUMN f_history_alarm.ack_time IS '确认时间';
COMMENT ON COLUMN f_history_alarm.clear_user IS '清除人';
COMMENT ON COLUMN f_history_alarm.ack_user IS '确认人';
COMMENT ON COLUMN f_history_alarm.archive_flag IS '是否归档，归档后才转移到历史告警';
COMMENT ON COLUMN f_history_alarm.archive_time IS '归档时间';
COMMENT ON COLUMN f_history_alarm.archive_user IS '归档人id';

create table f_event(
    id VARCHAR(32) NOT NULL,
    dn VARCHAR(255) NOT NULL,
    event_title text not null,
    event_text  text not null,
    event_time timestamp not null,
    status int4 not null default 0,
    level int4 not null default 1,
    event_extra text null,
    primary key (id)
);
COMMENT ON TABLE f_event IS '事件表';
COMMENT ON COLUMN f_event.id IS 'id';
COMMENT ON COLUMN f_event.dn IS 'dn';

create table o_collector_cript(
    id varchar(255) not null,
    script text not null,
    primary key (id)
);
-- +goose Down
-- +goose StatementBegin
SELECT 'down SQL query';
DROP TABLE IF EXISTS f_alarm cascade ;
DROP TABLE IF EXISTS f_history_alarm cascade ;
DROP TABLE IF EXISTS f_event cascade ;
DROP TABLE IF EXISTS o_collector_parsescript cascade ;
-- +goose StatementEnd
