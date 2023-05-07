package graph

// This file will be automatically regenerated based on the schema, any resolver implementations
// will be copied through when generating and any unknown code will be moved to the end.
// Code generated by github.com/99designs/gqlgen version v0.17.30

import (
	"context"
	"fmt"

	"github.com/weeb-vip/anime-api/graph/generated"
	"github.com/weeb-vip/anime-api/graph/model"
)

// FindAnimeByID is the resolver for the findAnimeByID field.
func (r *entityResolver) FindAnimeByID(ctx context.Context, id string) (*model.Anime, error) {
	panic(fmt.Errorf("not implemented: FindAnimeByID - findAnimeByID"))
}

// Entity returns generated.EntityResolver implementation.
func (r *Resolver) Entity() generated.EntityResolver { return &entityResolver{r} }

type entityResolver struct{ *Resolver }
