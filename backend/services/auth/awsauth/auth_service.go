package awsauth

import (
  "awsomegui/backend/services/session"
)

type AuthService struct {
  Workspace *session.WorkspaceManager
}

func NewAuthService(
  workspace *session.WorkspaceManager,
) *AuthService {
  return &AuthService{
    Workspace: workspace,
  }
}

func(a *AuthService) ConnectAWSProfile(
  profile string,
  region string,
) error {
  credService := &CredentialService{}

  err := credService.ValidateProfile(profile)
  if err != nil {
    return err
  }

  a.Workspace.ConnectProvider(
    "aws",
    profile,
    region,
  )

  return nil
}
