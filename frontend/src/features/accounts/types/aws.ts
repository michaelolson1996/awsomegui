export interface AWSProfileDTO {
  name: string;
  isDefault: boolean;
}

export interface VPCDTO {
  id: string;
  cidrBlock: string;
  isDefault: boolean;
  state: string;
}
