package main

type SearchResult struct {
	Events []struct {
		Event struct {
			Accepted       int    `json:"accepted"`
			Address        string `json:"address"`
			Catch          string `json:"catch"`
			Description    string `json:"description"`
			EndedAt        string `json:"ended_at"`
			EventID        int    `json:"event_id"`
			EventURL       string `json:"event_url"`
			Lat            string `json:"lat"`
			Limit          int    `json:"limit"`
			Lon            string `json:"lon"`
			OwnerID        int    `json:"owner_id"`
			OwnerNickname  string `json:"owner_nickname"`
			OwnerTwitterID string `json:"owner_twitter_id"`
			Place          string `json:"place"`
			StartedAt      string `json:"started_at"`
			Title          string `json:"title"`
			UpdatedAt      string `json:"updated_at"`
			URL            string `json:"url"`
			Waiting        int    `json:"waiting"`
		} `json:"event"`
	} `json:"events"`
	ResultsReturned int `json:"results_returned"`
	ResultsStart    int `json:"results_start"`
}
