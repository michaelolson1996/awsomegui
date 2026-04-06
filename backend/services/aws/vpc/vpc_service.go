package vpc

import (
  "context"

  "github.com/aws/aws-sdk-go-v2/service/ec2"
  "awsomegui/backend/services/aws/client"
  awsdto "awsomegui/backend/dto/awsdto"
)

type VPCService struct {
  Factory *client.ClientFactory
}

func NewVPCService(
  factory *client.ClientFactory,
) *VPCService {
  return &VPCService{
    Factory: factory,
  }
}

func (s *VPCService) ListVPCs() (
  []awsdto.VPCDTO,
  error,
) {
  client, err := s.Factory.EC2Client()
  if err != nil {
    return nil, err
  }

  output, err := client.DescribeVpcs(
    context.TODO(),
    &ec2.DescribeVpcsInput{},
  )
  if err != nil {
    return nil, err
  }

  var result []awsdto.VPCDTO

  for _, vpc := range output.Vpcs {
    result = append(result, awsdto.VPCDTO{
      ID:        *vpc.VpcId,
      CIDRBlock: *vpc.CidrBlock,
      IsDefault: *vpc.IsDefault,
      State:     string(vpc.State),
    })
  }

  return result, nil
}
