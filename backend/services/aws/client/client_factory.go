package client

import (
  "context"

  "github.com/aws/aws-sdk-go-v2/config"
  "github.com/aws/aws-sdk-go-v2/aws"

  "awsomegui/backend/services/session"
)

type ClientFactory struct {
  Workspace *session.WorkspaceManager
}

func NewClientFactory(
  workspace *session.WorkspaceManager,
) *ClientFactory {
  return &ClientFactory{
    Workspace: workspace,
  }
}

func (f *ClientFactory) LoadAWSConfig() (
  aws.Config,
  error,
) {
  active := f.Workspace.GetSession("aws")

  return config.LoadDefaultConfig(
    context.TODO(),
    config.WithSharedConfigProfile(active.Profile),
    config.WithRegion(active.Region),
  )
}
