package internal

import (
	"consumer/cmd/gen/sqlc/sqlc"
	"context"
	"github.com/confluentinc/confluent-kafka-go/kafka"
	"github.com/jackc/pgx/v5/pgtype"
	"time"
)

type Processor interface {
	Process(msg *kafka.Message) error
}

type KafkaProcessor struct {
	processor Processor
}

func NewKafkaProcessor(processor Processor) *KafkaProcessor {
	return &KafkaProcessor{processor: processor}
}

func (p *KafkaProcessor) Process(msg *kafka.Message) error {
	return p.processor.Process(msg)
}

type KafkaLogProcessor struct {
	processor     Processor
	logRepository *MessageLogRepository
}

func NewKafkaLogProcessor(p Processor, r *MessageLogRepository) *KafkaLogProcessor {
	return &KafkaLogProcessor{processor: p, logRepository: r}
}

func (p *KafkaLogProcessor) Process(msg *kafka.Message) error {
	tp := msg.TopicPartition
	h := msg.Headers

	ml := sqlc.InsertLogParams{
		Topic:        pgtype.Text{String: *tp.Topic, Valid: true},
		Partition:    pgtype.Int4{Int32: tp.Partition, Valid: true},
		TraceContext: extractHeader(h, "tc"),
		MsgID:        extractHeader(h, "id"),
		Offset:       pgtype.Text{String: tp.Offset.String(), Valid: true},
		ConsumedAt:   pgtype.Timestamp{Time: time.Now(), Valid: true},
	}
	ctx := context.Background()

	return p.logRepository.InsertLog(ctx, ml)
}

func extractHeader(headers []kafka.Header, key string) pgtype.Text {
	for _, h := range headers {
		if h.Key == key {
			return pgtype.Text{String: string(h.Value), Valid: true}
		}
	}
	return pgtype.Text{String: "", Valid: true}
}
