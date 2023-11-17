package usecase

import (
	"context"
	"errors"
	"sort"

	"l-gomock/model"
	"l-gomock/repository"
)

type GetPlayerRankingInputPort interface {
	Execute(context.Context) (*GetPlayerRankingOutput, error)
}

type GetPlayerRankingOutput struct {
	Ranking []*model.Player
	Count   int
}

type GetPlayerRankingInteractor struct {
	PlayerAPIRepository repository.PlayerAPIRepository
}

func NewGetPlayerRankingUsecase(pr repository.PlayerAPIRepository) GetPlayerRankingInputPort {
	return &GetPlayerRankingInteractor{
		PlayerAPIRepository: pr,
	}
}

func (i *GetPlayerRankingInteractor) Execute(ctx context.Context) (*GetPlayerRankingOutput, error) {
	player, err := i.PlayerAPIRepository.GetPlayerList(ctx)
	if err != nil {
		return nil, err
	}

	if len(player) == 0 {
		return nil, errors.New("not found")
	}

	sort.Slice(player, func(i, j int) bool {
		return player[i].Ranking < player[j].Ranking
	})

	output := &GetPlayerRankingOutput{
		Ranking: player,
		Count:   len(player),
	}

	return output, nil
}
