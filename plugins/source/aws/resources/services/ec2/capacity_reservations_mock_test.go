package ec2

import (
	"testing"

	"github.com/aws/aws-sdk-go-v2/service/ec2"
	"github.com/aws/aws-sdk-go-v2/service/ec2/types"
	"github.com/cloudquery/cloudquery/plugins/source/aws/client"
	"github.com/cloudquery/cloudquery/plugins/source/aws/client/mocks"
	"github.com/cloudquery/plugin-sdk/v3/faker"
	"github.com/golang/mock/gomock"
)

func buildEc2CapacityReservations(t *testing.T, ctrl *gomock.Controller) client.Services {
	m := mocks.NewMockEc2Client(ctrl)
	l := types.CapacityReservation{}
	err := faker.FakeObject(&l)
	if err != nil {
		t.Fatal(err)
	}
	m.EXPECT().DescribeCapacityReservations(gomock.Any(), gomock.Any(), gomock.Any()).Return(
		&ec2.DescribeCapacityReservationsOutput{
			CapacityReservations: []types.CapacityReservation{l},
		}, nil)
	return client.Services{
		Ec2: m,
	}
}

func TestEc2CapacityReservations(t *testing.T) {
	client.AwsMockTestHelper(t, CapacityReservations(), buildEc2CapacityReservations, client.TestOptions{})
}
