package grains

import (
	"fmt"
	"os"

	"github.com/asynkron/protoactor-go/cluster"
)

type OpportunityGrain struct {
}

func (v *OpportunityGrain) Init(ctx cluster.GrainContext) {
	fmt.Printf("ci-opportunity %v initialized\n", ctx.Identity())
}

func (v *OpportunityGrain) Started(ctx cluster.GrainContext) {
	fmt.Printf("opportunity %s started\n", ctx.Identity())
}

func (v *OpportunityGrain) GetOpportunity(opportunityRequest *GetOpportunityRequest, ctx cluster.GrainContext) (*GetOpportunityResponse, error) {
	machineName, _ := os.Hostname()
	opportunityId := ctx.Identity()

	fmt.Printf("opportunity-grain %s called for opportunity %s on %s\n", opportunityId, opportunityId, machineName)
	return &GetOpportunityResponse{
		OpportunityDetails: &OpportunityDetails{
			OpportunityId: opportunityId,
			Name:          "TestOpportunity",
		},
	}, nil
}

func (o *OpportunityGrain) Terminate(ctx cluster.GrainContext)      {}
func (o *OpportunityGrain) ReceiveDefault(ctx cluster.GrainContext) {}
