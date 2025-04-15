package handlers

import (
	"github.com/GuilhermeDeOliveiraAmorim/you-choose/internal/factories"
	"github.com/GuilhermeDeOliveiraAmorim/you-choose/internal/util"
)

type HandlerFactory struct {
	MovieHandler *MovieHandler
	ListHandler  *ListHandler
	VoteHandler  *VoteHandler
	UserHandler  *UserHandler
	BrandHandler *BrandHandler
}

func NewHandlerFactory(inputFactory util.ImputFactory) *HandlerFactory {
	movieFactory := factories.NewMovieFactory(inputFactory)
	listFactory := factories.NewListFactory(inputFactory)
	voteFactory := factories.NewVoteFactory(inputFactory)
	userFactory := factories.NewUserFactory(inputFactory)
	brandFactory := factories.NewBrandFactory(inputFactory)

	return &HandlerFactory{
		MovieHandler: NewMovieHandler(movieFactory),
		ListHandler:  NewListHandler(listFactory),
		VoteHandler:  NewVoteHandler(voteFactory),
		UserHandler:  NewUserHandler(userFactory),
		BrandHandler: NewBrandHandler(brandFactory),
	}
}
