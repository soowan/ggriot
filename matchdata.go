package ggriot

import (
	"github.com/tyhi/ggriot/cache"
	"github.com/tyhi/ggriot/models"
	"strconv"
)

// Match will return the data for the match ID requested.
func Match(region string, matchID int64) (md *models.MatchData, err error) {
	mtID := strconv.FormatInt(matchID, 10)

	cp := cache.CachedParams{
		Cached:   true,
		Expire:   false,
		CallType: "league_match_by_id_" + region,
		CallKey:  mtID,
	}
	if err = apiRequest("https://"+region+"."+Base+BaseMatch+"/matches/"+mtID, &md, cp); err != nil {
		return nil, err
	}

	return md, nil
}

// MatchHistory will return an array of matches based on the parameters given.
// This doesn't actually return wins or loses, just basic information about the game.
// In order to get stats you have to request every game separately.
// TODO: Add ability to fully use the options when doing a matches call.
// TODO: Figure out if/how this can/should be cached.
func MatchHistory(region string, accountID string) (ms *models.MatchHistory, err error) {
	cp := cache.CachedParams{
		Cached: false,
	}
	if err = apiRequest("https://"+region+"."+Base+BaseMatch+"/matchlists/by-account/"+accountID, &ms, cp); err != nil {
		return nil, err
	}

	return ms, nil
}

// MatchTimeline will return the full timeline of the requested match ID.
// This is a pretty big struct that is returned so make sure you understand how to use this data first.
func MatchTimeline(region string, matchID int64) (mt *models.MatchTimeline, err error) {
	mtID := strconv.FormatInt(matchID, 10)
	cp := cache.CachedParams{
		Cached:   true,
		Expire:   false,
		CallType: "league_match_tl_by_id_" + region,
		CallKey:  mtID,
	}
	if err = apiRequest("https://"+region+"."+Base+BaseMatch+"/timelines/by-match/"+mtID, &mt, cp); err != nil {
		return nil, err
	}

	return mt, nil
}
