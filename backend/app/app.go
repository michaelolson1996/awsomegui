package app

import (
  "context"

  awsauth "awsomegui/backend/services/auth/awsauth"
  vpcsvc "awsomegui/backend/services/aws/vpc"
  "awsomegui/backend/services/session"
)

type App struct {
  ctx       context.Context

  Workspace *session.WorkspaceManager
  Profiles  *awsauth.ProfileService
  VPC       *vpcsvc.VPCService
}

func NewApp(
  workspace *session.WorkspaceManager,
  profiles *awsauth.ProfileService,
  vpc *vpcsvc.VPCService,
) *App {
  return &App{
    Workspace: workspace,
    Profiles:  profiles,
    VPC:       vpc,
  }
}
