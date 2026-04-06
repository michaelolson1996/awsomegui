package awsauth

import (
  "context"

  "github.com/aws/aws-sdk-go-v2/config"
  "github.com/aws/aws-sdk-go-v2/service/sts"
)

type CredentialService struct{}

func (s *CredentialService) ValidateProfile(
  profile string,
) error {
  cfg, err := config.LoadDefaultConfig(
    context.TODO(),
    config.WithSharedConfigProfile(profile),
  )
  if err != nil {
    return err
  }

  client := sts.NewFromConfig(cfg)

  _, err = client.GetCallerIdentity(
    context.TODO(),
    &sts.GetCallerIdentityInput{},
  )

  return err
}
