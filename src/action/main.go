package action

import (
	. "mars_rover/src/domain"
)

type Action struct {
	repo Repository
}

func For(repo Repository) *Action {
	return &Action{
		repo: repo,
	}
}
