package model

import "time"

type Recipe struct {
	ID           string    `json:"id" swag:"-"`
	Name         string    `json:"name" `
	Tags         []string  `json:"tags"`
	Ingredients  []string  `json:"ingredients"`
	Instructions []string  `json:"instructions"`
	PublishedAt  time.Time `json:"publishedAt" swag:"-"`
}
