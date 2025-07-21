-- name: InsertLog :one
insert into msg_logs (msg_id, topic, partition, trace_context, "offset", consumed_at)
values ($1, $2, $3, $4, $5, $6) returning *;
