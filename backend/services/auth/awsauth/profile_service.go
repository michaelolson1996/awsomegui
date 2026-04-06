package awsauth

import (
  "bufio"
  "strings"
  "os"

  "github.com/aws/aws-sdk-go-v2/config"

  awsdto "awsomegui/backend/dto/awsdto"
)

type ProfileService struct{}

func NewProfileService() *ProfileService {
  return &ProfileService{}
}

func (s *ProfileService) ListProfiles() (
  []awsdto.AWSProfileDTO,
  error,
) {
  path := config.DefaultSharedConfigFilename()

  file, err := os.Open(path)
  if err != nil {
    return nil, err
  }
  defer file.Close()

  var profiles []awsdto.AWSProfileDTO
  scanner := bufio.NewScanner(file)

  for scanner.Scan() {
    line := strings.TrimSpace(scanner.Text())

    if strings.HasPrefix(line, "[") &&
        strings.HasSuffix(line, "]") {

      name := strings.Trim(line, "[]")
      name = strings.TrimPrefix(name, "profile ")

      profiles = append(profiles, awsdto.AWSProfileDTO{
        Name:      name,
        IsDefault: name == "default",
      })
    }
  }

  return profiles, scanner.Err()
}
