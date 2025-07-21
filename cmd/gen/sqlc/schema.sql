create table msg_logs (
    id UUID primary key default gen_random_uuid(),
    msg_id varchar(200),
    topic varchar(200),
    partition integer,
    trace_context varchar(200),
    "offset" varchar(200),
    consumed_at timestamp
);
