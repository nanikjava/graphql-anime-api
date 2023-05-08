// Code generated by github.com/99designs/gqlgen, DO NOT EDIT.

package model

import (
	"time"
)

type Anime struct {
	ID            string     `json:"id"`
	Anidbid       *string    `json:"anidbid,omitempty"`
	TitleEn       *string    `json:"titleEn,omitempty"`
	TitleJp       *string    `json:"titleJp,omitempty"`
	TitleRomaji   *string    `json:"titleRomaji,omitempty"`
	TitleKanji    *string    `json:"titleKanji,omitempty"`
	TitleSynonyms []string   `json:"titleSynonyms,omitempty"`
	Description   *string    `json:"description,omitempty"`
	ImageURL      *string    `json:"imageUrl,omitempty"`
	Tags          []string   `json:"tags,omitempty"`
	Studios       []string   `json:"studios,omitempty"`
	AnimeStatus   *string    `json:"animeStatus,omitempty"`
	Episodes      *int       `json:"episodes,omitempty"`
	Duration      *string    `json:"duration,omitempty"`
	Rating        *string    `json:"rating,omitempty"`
	CreatedAt     string     `json:"createdAt"`
	UpdatedAt     string     `json:"updatedAt"`
	StartDate     *time.Time `json:"startDate,omitempty"`
	EndDate       *time.Time `json:"endDate,omitempty"`
	Broadcast     *string    `json:"broadcast,omitempty"`
	Source        *string    `json:"source,omitempty"`
	Licensors     []string   `json:"licensors,omitempty"`
	Ranking       *int       `json:"ranking,omitempty"`
}

func (Anime) IsEntity() {}

type AnimeAPI struct {
	// Version of event publish service
	Version string `json:"version"`
}

type AnimeSearchInput struct {
	// Search query
	Query string `json:"query"`
	// Page number
	Page int `json:"page"`
	// Items per page
	PerPage int `json:"perPage"`
	// Sort by
	SortBy *string `json:"sortBy,omitempty"`
	// Sort direction
	SortDirection *string `json:"sortDirection,omitempty"`
	// Tags
	Tags []string `json:"tags,omitempty"`
	// Studios
	Studios []string `json:"studios,omitempty"`
	// Anime statuses
	AnimeStatuses []string `json:"animeStatuses,omitempty"`
}

type APIInfo struct {
	AnimeAPI *AnimeAPI `json:"animeApi"`
	Name     string    `json:"name"`
}
