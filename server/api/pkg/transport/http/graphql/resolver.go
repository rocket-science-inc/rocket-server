package graphql

import (
	"context"
)

type Resolver struct{}

func (r *Resolver) Mutation() MutationResolver {
	return &mutationResolver{r}
}
func (r *Resolver) Query() QueryResolver {
	return &queryResolver{r}
}

type mutationResolver struct{ *Resolver }

func (r *mutationResolver) AddEvent(ctx context.Context, title string, info string) (*Event, error) {
	panic("not implemented")
}

type queryResolver struct{ *Resolver }

func (r *queryResolver) Events(ctx context.Context) ([]*Event, error) {
	panic("not implemented")
}
func (r *queryResolver) Event(ctx context.Context, id string) (*Event, error) {
	panic("not implemented")
}
