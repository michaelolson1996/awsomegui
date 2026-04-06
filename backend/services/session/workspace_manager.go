package session

type ProviderSession struct {
  Provider    string
  Profile     string
  Region      string
  IsConnected bool
  LastError   string
}

type WorkspaceManager struct {
  ActiveProvider string
  Sessions       map[string]*ProviderSession
}

func NewWorkspaceManager() *WorkspaceManager {
  return &WorkspaceManager{
    ActiveProvider: "",
    Sessions: make(map[string]*ProviderSession),
  }
}

func (w *WorkspaceManager) ConnectProvider(
  provider string,
  profile string,
  region string,
) {
  w.Sessions[provider] = &ProviderSession{
    Provider: provider,
    Profile: profile,
    Region: region,
    IsConnected: true,
  }

  w.ActiveProvider = provider
}

func (w *WorkspaceManager) DisconnectProvider(provider string) {
  delete(w.Sessions, provider)

  if w.ActiveProvider == provider {
    w.ActiveProvider = ""
  }
}

func (w *WorkspaceManager) GetActiveSession() *ProviderSession {
  return w.Sessions[w.ActiveProvider]
}

func (w *WorkspaceManager) SetActiveProvider(provider string) {
  if _, exists := w.Sessions[provider]; exists {
    w.ActiveProvider = provider
  }
}


func (w *WorkspaceManager) GetSession(
  provider string,
) *ProviderSession {
  return w.Sessions[provider]
}

func (w *WorkspaceManager) GetProfile(
  provider string,
) string {
  if session, exists := w.Sessions[provider]; exists {
    return session.Profile
  }
  return ""
}

func (w *WorkspaceManager) SetProfile(
  provider string,
  profile string,
) {
  if session, exists := w.Sessions[provider]; exists {
    session.Profile = profile
  }
}

func (w *WorkspaceManager) GetRegion(
  provider string,
) string {
  if session, exists := w.Sessions[provider]; exists {
    return session.Region
  }
  return ""
}

func (w *WorkspaceManager) SetRegion(
  provider string,
  region string,
) {
  if session, exists := w.Sessions[provider]; exists {
    session.Region = region
  }
}

