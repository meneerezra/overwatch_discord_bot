package jsonModels

type Settings struct {
	DiscordToken string `json:"discord_token,omitempty"`
	OverfastURL  string `json:"overfast_url,omitempty"`
	BestOverwatchApiURL string `json:"best_overwatch_api_url,omitempty"`
}

func (s *Settings) DefaultValues() {
	s.DiscordToken = "yayap"
	s.OverfastURL = "https://overfast-api.tekrop.fr/"
	s.BestOverwatchApiURL = "https://best-overwatch-api.herokuapp.com/"
}
