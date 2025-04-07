package overfast

import (
	"net/http"
	"net/url"
)

type PlayersResponse struct {
	Total   int      `json:"total"`
	Results []Player `json:"results"`
}

type Player struct {
	PlayerID      string `json:"player_id"`
	Name          string `json:"name"`
	Avatar        string `json:"avatar"`
	Namecard      string `json:"namecard"`
	Title         string `json:"title"`
	CareerURL     string `json:"career_url"`
	BlizzardID    string `json:"blizzard_id"`
	LastUpdatedAt int64  `json:"last_updated_at"`
}

func (c OverfastClient) GetPlayersByName(name string) (PlayersResponse, error) {
	var playersResponse PlayersResponse

	params := url.Values{}
	params.Add("name", name)

	req, err := c.NewRequest(http.MethodGet, "/players", nil, params)
	if err != nil {
	    return playersResponse, err
	}

	err = c.Do(req, &playersResponse)
	if err != nil {
		return playersResponse, err
	}


	return playersResponse, nil
}












