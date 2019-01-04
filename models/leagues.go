package models

// LeagueRoster returns everyone in the  tier for the mode requested.
type LeagueRoster struct {
	Tier     string `json:"tier"`
	Queue    string `json:"queue"`
	LeagueID string `json:"leagueId"`
	Name     string `json:"name"`
	Entries  []struct {
		HotStreak        bool   `json:"hotStreak"`
		Wins             int    `json:"wins"`
		Veteran          bool   `json:"veteran"`
		Losses           int    `json:"losses"`
		Rank             string `json:"rank"`
		PlayerOrTeamName string `json:"playerOrTeamName"`
		Inactive         bool   `json:"inactive"`
		PlayerOrTeamID   string `json:"playerOrTeamId"`
		FreshBlood       bool   `json:"freshBlood"`
		LeaguePoints     int    `json:"leaguePoints"`
	} `json:"entries"`
}

// MGCTierList is the interface for Masters, Grandmaster and Challenger tiers.
type MGCTierList struct {
	Tier    string `json:"tier"`
	Entries []struct {
		Rank         string `json:"rank"`
		Wins         int    `json:"wins"`
		Losses       int    `json:"losses"`
		SummonerID   string `json:"summonerId"`
		LeaguePoints int    `json:"leaguePoints"`
		SummonerName string `json:"summonerName"`
	} `json:"entries"`
	LeagueID string `json:"leagueId"`
}

// LeaguePosition is what's returned when requesting a players league stats.
type LeaguePosition []struct {
	Rank         string `json:"rank"`
	Tier         string `json:"tier"`
	Wins         int    `json:"wins"`
	Losses       int    `json:"losses"`
	LeagueID     string `json:"leagueId"`
	QueueType    string `json:"queueType"`
	LeagueName   string `json:"leagueName"`
	SummonerID   string `json:"summonerId"`
	LeaguePoints int    `json:"leaguePoints"`
	SummonerName string `json:"summonerName"`
	MiniSeries   struct {
		Wins     int    `json:"wins"`
		Losses   int    `json:"losses"`
		Target   int    `json:"target"`
		Progress string `json:"progress"`
	} `json:"miniSeries,omitempty"`
}
