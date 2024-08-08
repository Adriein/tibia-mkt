package cron

import (
	"fmt"
	"github.com/adriein/tibia-mkt/pkg/constants"
	"github.com/adriein/tibia-mkt/pkg/service"
	"github.com/adriein/tibia-mkt/pkg/types"
	"net/http"
	"strings"
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
			File:     "kill-statistics.go",
		}
	}

	client := &http.Client{}
	response, requestError := client.Do(request)

	if requestError != nil {
		return types.ApiError{
			Msg:      requestError.Error(),
			Function: "Execute -> client.Do()",
			File:     "kill-statistics.go",
		}
	}

	defer response.Body.Close()

	parsedResponse, decodeErr := service.Decode[TibiaApiKillStatisticsResponse](response.Body)

	if decodeErr != nil {
		return types.ApiError{
			Msg:      decodeErr.Error(),
			Function: "Execute -> service.Decode()",
			File:     "kill-statistics.go",
		}
	}

	hashTable := make(map[string]int)

	for _, statistic := range parsedResponse.KillStatistics.Entries {
		hashTable[statistic.Race] = statistic.LastDayKilled
	}

	for _, cog := range cogs {
		for _, creature := range cog.Creatures {
			name := kc.pluralize(creature.Name)

			killStatistic := hashTable[name]
		}
	}

	return nil
}

func (kc *KillStatisticsCron) pluralize(creatureName string) string {
	creatureName = strings.ToLower(creatureName)

	creatureName += "s"

	return creatureName
}
