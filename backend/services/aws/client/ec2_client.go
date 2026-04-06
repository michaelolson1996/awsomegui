package client

import (
  "github.com/aws/aws-sdk-go-v2/service/ec2"
)

func (f *ClientFactory) EC2Client() (
  *ec2.Client,
  error,
) {
  cfg, err := f.LoadAWSConfig()
  if err != nil {
    return nil, err
  }

  return ec2.NewFromConfig(cfg), nil
}
