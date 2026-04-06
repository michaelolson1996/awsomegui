package projectdto

type CloudProviderConnectionDTO struct {
  Provider string `json:"provider"`
  Account  string `json:"account"`
  Profile  string `json:"profile"`
  Region   string `json:"region"`
}

type ProjectWorkspaceDTO struct {
  Name      string                       `json:"name"`
  IsPublic  bool                         `json:"isPublic"`
  Providers []CloudProviderConnectionDTO `json:"providers"`
}
