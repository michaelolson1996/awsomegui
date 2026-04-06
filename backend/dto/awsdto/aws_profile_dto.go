package awsdto

type AWSProfileDTO struct {
  Name      string `json:"name"`
  IsDefault bool   `json:"isDefault"`
  Region    string `json:"region"`
}
