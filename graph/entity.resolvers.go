package graph

// This file will be automatically regenerated based on the schema, any resolver implementations
// will be copied through when generating and any unknown code will be moved to the end.

import (
	"context"

	"gitlab.slade360emr.com/go/base"
	"gitlab.slade360emr.com/go/base/graph/generated"
)

func (r *entityResolver) FindPageInfoByHasNextPage(ctx context.Context, hasNextPage bool) (*base.PageInfo, error) {
	return nil, nil
}

// Entity returns generated.EntityResolver implementation.
func (r *Resolver) Entity() generated.EntityResolver { return &entityResolver{r} }

type entityResolver struct{ *Resolver }
