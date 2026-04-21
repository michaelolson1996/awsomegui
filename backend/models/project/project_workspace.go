package project

type CloudProviderConnection struct {
  Provider string
  Account  string
  Profile  string
  Regions  []string
}

type ProjectWorkspace struct {
  ProjectID  string
  Name       string
  Providers  []CloudProviderConnection
  OwnerID    string
  MemberIDs  []string
  IsPublic   bool
}
