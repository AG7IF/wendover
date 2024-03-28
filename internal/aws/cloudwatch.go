package aws

import (
	"context"
	"time"

	"github.com/aws/aws-sdk-go-v2/aws"
	"github.com/aws/aws-sdk-go-v2/service/cloudwatchlogs"
	"github.com/aws/aws-sdk-go-v2/service/cloudwatchlogs/types"
	"github.com/pkg/errors"
)

type cloudWatchLogEvent struct {
	Message   string
	Timestamp int64
}

type CloudwatchWriter struct {
	client          *cloudwatchlogs.Client
	logGroupName    string
	logStreamName   string
	buffer          []cloudWatchLogEvent
	bufferLength    int
	bufferByteCount int
}

func NewCloudWatchWriter(cfg aws.Config, logGroupName, logStreamName string, bufferLength int) *CloudwatchWriter {
	client := cloudwatchlogs.NewFromConfig(cfg)

	return &CloudwatchWriter{
		client:        client,
		logGroupName:  logGroupName,
		logStreamName: logStreamName,
		bufferLength:  bufferLength,
	}
}

func (cw *CloudwatchWriter) sendLogs() error {
	var events []types.InputLogEvent
	for _, v := range cw.buffer {
		event := types.InputLogEvent{
			Message:   &v.Message,
			Timestamp: &v.Timestamp,
		}

		events = append(events, event)
	}

	input := &cloudwatchlogs.PutLogEventsInput{
		LogEvents:     events,
		LogGroupName:  &cw.logGroupName,
		LogStreamName: &cw.logStreamName,
	}

	_, err := cw.client.PutLogEvents(context.TODO(), input)
	if err != nil {
		return errors.WithStack(err)
	}

	cw.buffer = make([]cloudWatchLogEvent, cw.bufferLength)
	cw.bufferByteCount = 0

	return nil
}

func (cw *CloudwatchWriter) Write(p []byte) (int, error) {
	event := cloudWatchLogEvent{
		Message:   string(p),
		Timestamp: time.Now().UnixMilli(),
	}

	cw.buffer = append(cw.buffer, event)
	cw.bufferByteCount = cw.bufferByteCount + len(p)

	if len(cw.buffer) >= cw.bufferLength {
		storedBytes := cw.bufferByteCount
		err := cw.sendLogs()
		if err != nil {
			return 0, errors.WithStack(err)
		}

		return storedBytes, nil
	}

	return 0, nil
}

func (cw *CloudwatchWriter) Close() error {
	return cw.sendLogs()
}
