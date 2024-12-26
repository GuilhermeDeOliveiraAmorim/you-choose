package repositories

import "github.com/GuilhermeDeOliveiraAmorim/you-choose/internal/entities"

type VoteRepository interface {
	CreateVote(movie entities.Vote) error
}
