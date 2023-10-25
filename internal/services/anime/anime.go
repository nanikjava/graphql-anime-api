package anime

import (
	"context"
	"github.com/weeb-vip/anime-api/internal/db/repositories/anime"
	"gopkg.in/DataDog/dd-trace-go.v1/ddtrace/tracer"
)

type AnimeServiceImpl interface {
	AnimeByID(ctx context.Context, id string) (*anime.Anime, error)
	TopRatedAnime(ctx context.Context, limit int) ([]*anime.Anime, error)
	MostPopularAnime(ctx context.Context, limit int) ([]*anime.Anime, error)
	NewestAnime(ctx context.Context, limit int) ([]*anime.Anime, error)
}

type AnimeService struct {
	Repository anime.AnimeRepositoryImpl
}

func NewAnimeService(animeRepository anime.AnimeRepositoryImpl) AnimeServiceImpl {
	return &AnimeService{
		Repository: animeRepository,
	}
}

func (a *AnimeService) AnimeByID(ctx context.Context, id string) (*anime.Anime, error) {
	span := tracer.StartSpan("AnimeByID")
	span.SetTag("service", "anime")
	span.SetTag("type", "service")
	defer span.Finish()

	return a.Repository.FindById(id)
}

func (a *AnimeService) TopRatedAnime(ctx context.Context, limit int) ([]*anime.Anime, error) {
	span := tracer.StartSpan("TopRatedAnime")
	span.SetTag("service", "anime")
	span.SetTag("type", "service")
	defer span.Finish()

	return a.Repository.TopRatedAnime(limit)
}

func (a *AnimeService) MostPopularAnime(ctx context.Context, limit int) ([]*anime.Anime, error) {
	span := tracer.StartSpan("MostPopularAnime")
	span.SetTag("service", "anime")
	span.SetTag("type", "service")
	defer span.Finish()

	return a.Repository.MostPopularAnime(limit)
}

func (a *AnimeService) NewestAnime(ctx context.Context, limit int) ([]*anime.Anime, error) {
	span := tracer.StartSpan("NewestAnime")
	span.SetTag("service", "anime")
	span.SetTag("type", "service")
	defer span.Finish()

	return a.Repository.NewestAnime(limit)
}
