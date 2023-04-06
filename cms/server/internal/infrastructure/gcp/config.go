package gcp

import "fmt"

type TaskConfig struct {
	GCPProject    string
	GCPRegion     string
	QueueName     string
	SubscriberURL string
	Topic         string
	GCSHost       string
	Timeout       int64 `default:"1800"` // second
}

func (c *TaskConfig) buildQueueUrl() (string, error) {
	if c.GCPProject == "" || c.GCPRegion == "" || c.QueueName == "" {
		return "", ErrMissignConfig
	}
	return fmt.Sprintf("projects/%s/locations/%s/queues/%s", c.GCPProject, c.GCPRegion, c.QueueName), nil
}
