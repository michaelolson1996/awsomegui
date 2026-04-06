package project

type CloudProviderConnection struct {
  Provider string
  Account  string
  Profile  string
  Region   string
}

type ProjectWorkspace struct {
  ProjectID  string
  Name       string
  Providers  []CloudProviderConnection
  IsPublic   bool
  OwnerID    string
  TeamIDs    []string
}
