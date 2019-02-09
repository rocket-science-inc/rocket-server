package graphql

import (
	"context"

	types "rocket-server/server/api/pkg/types"
	endpoint "rocket-server/server/api/pkg/endpoint"
)

type Resolver struct {
	Endpoints endpoint.Endpoints
}

func (r *Resolver) Query() QueryResolver {
	return &queryResolver{r}
}

func (r *Resolver) Mutation() MutationResolver {
	return &mutationResolver{r}
}

type queryResolver struct{ *Resolver }

func (r *queryResolver) GetEvents(ctx context.Context) ([]types.Event, error) {
	return r.Endpoints.GetEvents(ctx)
}

type mutationResolver struct{ *Resolver }

func (r *mutationResolver) AddEvent(ctx context.Context, event types.NewEvent) (types.Event, error) {
	return r.Endpoints.AddEvent(ctx, event)
}
