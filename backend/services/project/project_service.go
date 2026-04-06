package project

import projectmodel "awsomegui/backend/models/project"

type ProjectService struct{}

func NewProjectService() *ProjectService {
  return &ProjectService
}

func (s *ProjectService) CreateProject(
  name string,
  ownerID string,
) *projectmodel.ProjectWorkspace {
  return &projectmodel.ProjectWorkspace{
    Name: name,
    OwnerID: ownerID,
    IsPublic: false,
  }
}

func (s *ProjectService) AddProviderConnection(
  project *projectmodel.ProjectWorkspace,
  provider *projectmodel.CloudProviderConnection,
) {
  project.Providers = append(
    project.Providers,
    provider,
  )
}
