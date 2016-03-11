package main

import (
	"errors"

	"github.com/aws/aws-sdk-go/aws/ec2metadata"
	"github.com/aws/aws-sdk-go/aws/session"
)

type EC2Metadata struct {
	InstanceID,
	ImageID,
	InstanceType string
}

func NewEc2Metadata() (*EC2Metadata, error) {
	service := ec2metadata.New(session.New())

	if !service.Available() {
		return nil, errors.New("EC2 Metadata service is unavailable. Make sure the application is running within an EC2 Instance.")
	}

	meta := &EC2Metadata{}
	meta.InstanceID, _ = service.GetMetadata("instance-id")
	meta.InstanceType, _ = service.GetMetadata("instance-type")
	meta.ImageID, _ = service.GetMetadata("ami-id")

	return meta, nil
}
