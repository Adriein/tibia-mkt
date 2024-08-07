package cron

import (
	"fmt"
	"github.com/adriein/tibia-mkt/pkg/constants"
	"github.com/adriein/tibia-mkt/pkg/service"
	"github.com/adriein/tibia-mkt/pkg/types"
	"net/http"
)

type KillStatisticsCron struct{}

type Entries struct {
	Race                  string `json:"race"`
	LastDayPlayersKilled  int    `json:"last_day_players_killed"`
	LastDayKilled         int    `json:"last_day_killed"`
	LastWeekPlayersKilled int    `json:"last_week_players_killed"`
	LastWeekKilled        int    `json:"last_week_killed"`
}

type KillStatistics struct {
	Entries []Entries `json:"entries"`
}

type TibiaApiKillStatisticsResponse struct {
	World          string
	KillStatistics KillStatistics `json:"killStatistics"`
}

func NewKillStatisticsCron() *KillStatisticsCron {
	return &KillStatisticsCron{}
}

func (kc *KillStatisticsCron) Execute(cogs []types.Cog) error {
	request, requestCreationError := http.NewRequest(
		http.MethodGet,
		fmt.Sprintf(
			"%s/%s/%s/%s",
			constants.TibiaDataApiBaseUrl,
			constants.TibiaDataApiVersion,
			constants.TibiaDataApiKillStatisticsUrl,
			constants.WorldSecura,
		),
		nil,
	)

	if requestCreationError != nil {
		return types.ApiError{
			Msg:      requestCreationError.Error(),
			Function: "Execute -> http.NewRequest()",
			File:     "main.go",
		}
	}

	client := &http.Client{}
	response, requestError := client.Do(request)

	if requestError != nil {
		return types.ApiError{
			Msg:      requestError.Error(),
			Function: "Execute -> client.Do()",
			File:     "main.go",
		}
	}

	defer response.Body.Close()

	parsedResponse, decodeErr := service.Decode[TibiaApiKillStatisticsResponse](request)

	if decodeErr != nil {
		return types.ApiError{
			Msg:      decodeErr.Error(),
			Function: "Execute -> service.Decode",
			File:     "main.go",
		}
	}

	fmt.Println(parsedResponse)

	return nil
}
