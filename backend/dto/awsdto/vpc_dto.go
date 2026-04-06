package awsdto

type VPCDTO struct {
  ID        string `json:"id"`
  Name      string `json:"name"`
  CIDRBlock string `json:"cidrBlock"`
  State     string `json:"state"`
  IsDefault bool   `json:"isDefault"`
  Region    string `json:"region"`
}
