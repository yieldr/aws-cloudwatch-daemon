package main

import (
	"github.com/aws/aws-sdk-go/aws"
	"github.com/aws/aws-sdk-go/aws/awsutil"
	"github.com/aws/aws-sdk-go/aws/session"
	"github.com/aws/aws-sdk-go/service/cloudwatch"
)

type CloudWatch struct {
	Client     *cloudwatch.CloudWatch
	Metrics    []*cloudwatch.MetricDatum
	Dimensions []*cloudwatch.Dimension
	Namespace  string
}

func (c *CloudWatch) AddDimention(name, value string) {
	c.Dimensions = append(c.Dimensions, &cloudwatch.Dimension{
		Name:  aws.String(name),
		Value: aws.String(value),
	})
}

func (c *CloudWatch) AddMetric(name, unit string, value float64) {
	c.Metrics = append(c.Metrics, &cloudwatch.MetricDatum{
		MetricName: aws.String(name),
		Unit:       aws.String(unit),
		Value:      aws.Float64(value),
		Dimensions: c.Dimensions,
	})

}

func (c *CloudWatch) ClearDimentions() {
	c.Dimensions = make([]*cloudwatch.Dimension, 0)
}

func (c *CloudWatch) ClearMetrics() {
	c.Metrics = make([]*cloudwatch.MetricDatum, 0)
}

func (c *CloudWatch) Send() (string, error) {
	in := &cloudwatch.PutMetricDataInput{
		MetricData: c.Metrics,
		Namespace:  aws.String(c.Namespace),
	}
	out, err := c.Client.PutMetricData(in)
	if err != nil {
		return "", err
	}
	c.ClearMetrics()
	return awsutil.StringValue(out), nil
}

func NewCloudWatch(namespace string) *CloudWatch {
	return &CloudWatch{
		Client:    cloudwatch.New(session.New()),
		Metrics:   make([]*cloudwatch.MetricDatum, 0),
		Namespace: namespace,
	}
}
