package cron

import (
	"fmt"
	"github.com/adriein/tibia-mkt/pkg/constants"
	"github.com/adriein/tibia-mkt/pkg/service"
	"github.com/adriein/tibia-mkt/pkg/types"
	"github.com/google/uuid"
	"net/http"
	"strings"
	"time"
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

func (kc *KillStatisticsCron) Execute(cogs []types.Cog) ([]types.KillStatistic, error) {
	url := fmt.Sprintf(
		"%s/%s/%s/%s",
		constants.TibiaDataApiBaseUrl,
		constants.TibiaDataApiVersion,
		constants.TibiaDataApiKillStatisticsUrl,
		constants.WorldSecura,
	)

	request, requestCreationError := http.NewRequest(http.MethodGet, url, nil)

	if requestCreationError != nil {
		return nil, types.ApiError{
			Msg:      requestCreationError.Error(),
			Function: "Execute -> http.NewRequest()",
			File:     "kill-statistics.go",
		}
	}

	client := &http.Client{}
	response, requestError := client.Do(request)

	if requestError != nil {
		return nil, types.ApiError{
			Msg:      requestError.Error(),
			Function: "Execute -> client.Do()",
			File:     "kill-statistics.go",
		}
	}

	if response.StatusCode != http.StatusOK {
		return nil, types.ApiError{
			Msg:      fmt.Sprintf("TibiaApi responding with status code %s", response.Status),
			Function: "Execute -> client.Do()",
			File:     "kill-statistics.go",
			Values:   []string{url},
			Domain:   true,
		}
	}

	defer response.Body.Close()

	parsedResponse, decodeErr := service.Decode[TibiaApiKillStatisticsResponse](response.Body)

	if decodeErr != nil {
		return nil, types.ApiError{
			Msg:      decodeErr.Error(),
			Function: "Execute -> service.Decode()",
			File:     "kill-statistics.go",
		}
	}

	hashTable := make(map[string]int)

	var result []types.KillStatistic

	for _, statistic := range parsedResponse.KillStatistics.Entries {
		hashTable[statistic.Race] = statistic.LastDayKilled
	}

	for _, cog := range cogs {
		for _, creature := range cog.Creatures {
			id := uuid.New()

			name := kc.pluralize(creature.Name)

			killStatistic := hashTable[name]

			result = append(result, types.KillStatistic{
				Id:           id.String(),
				CreatureName: creature.Name,
				AmountKilled: killStatistic,
				DropRate:     creature.DropRate,
				ExecutedBy:   constants.TibiaMktCronUser,
				CreatedAt:    time.Now().UTC().Format(time.DateTime),
				UpdatedAt:    time.Now().UTC().Format(time.DateTime),
			})
		}
	}

	return result, nil
}

func (kc *KillStatisticsCron) pluralize(creatureName string) string {
	creatureName = strings.ToLower(creatureName)

	creatureName += "s"

	return creatureName
}
