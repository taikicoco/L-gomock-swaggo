package usecase_test

import (
	"context"
	"testing"

	"github.com/golang/mock/gomock"
	"github.com/stretchr/testify/require"

	"l-gomock/model"
	repository "l-gomock/repository/mock"
	"l-gomock/usecase"
)

func TestGetPlayerRankingInteractor_Execute(t *testing.T) {
	ctx := context.Background()
	resp := []*model.Player{
		{ID: "1", Name: "test-player-1", Ranking: 3},
		{ID: "2", Name: "test-player-2", Ranking: 2},
		{ID: "3", Name: "test-player-3", Ranking: 1},
	}

	// (1) モックを呼び出すための Controller を生成
    ctrl := gomock.NewController(t)
    defer ctrl.Finish()

    // (2) モックの生成
    pr := repository.NewMockPlayerAPIRepository(ctrl)
    // (3) テストに呼ばれるべきメソッドと引数・戻り値を指定
    pr.EXPECT().GetPlayerList(ctx).Return(resp, nil)

    // (4) テストの本体
    output, err := usecase.NewGetPlayerRankingUsecase(pr).Execute(ctx)
    require.NoError(t, err)
    require.Len(t, output.Ranking, 3)
    for i, p := range output.Ranking {
        require.Equal(t, i+1, p.Ranking)
    }
    require.Equal(t, 3, output.Count)
}
