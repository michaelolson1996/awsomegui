package app

import (
  awsdto "awsomegui/backend/dto/awsdto"
//  vpcsvc "awsomegui/backend/services/aws/vpc"
)

func (a *App) GetAWSProfiles() (
  []awsdto.AWSProfileDTO,
  error,
) {
  return a.Profiles.ListProfiles()
}

func (a *App) SetAWSProfile(
  profile string,
) {
  a.Workspace.SetProfile("aws", profile)
}

func (a *App) GetAWSRegions() []string {
  return []string{
    "us-east-1",
    "us-east-2",
    "us-west-1",
    "us-west-2",
    "eu-west-1",
    "ap-southeast-1",
    "ap-southeast-2",
  }
}

func (a *App) SetAWSRegion(
  region string,
) {
  a.Workspace.SetRegion("aws", region)
}

func (a *App) GetAllVPCs() (
  []awsdto.VPCDTO,
  error,
) {
  return a.VPC.ListVPCs()
}
