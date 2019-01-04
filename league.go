package ggriot

import (
	"time"

	"github.com/soowan/ggriot/models"

	"github.com/soowan/ggriot/cache"
)

var (
	// LeagueByIDExpire is how long this call should be considered "fresh".
	LeagueByIDExpire = time.Duration(15 * time.Minute)

	// ExpireMGC is how long Masters, Grandmasters, Challenger tier list calls should be considered "fresh"
	ExpireMGC = time.Duration(15 * time.Minute)

	// ExpireGetPlayerPosition is how long this call should be considered "fresh"
	ExpireGetPlayerPosition = time.Duration(15 * time.Minute)
)

// Challengers will return all the challengers in the requested queue.
// Valid queues are, Ranked5s(RANKED_SOLO_5x5), Flex3s(RANKED_FLEX_TT), and Flex5s(RANKED_FLEX_SR)
func Challengers(region string, mode string) (lr *models.MGCTierList, err error) {
	cp := cache.CachedParams{
		Cached:     true,
		Expire:     true,
		Expiration: ExpireMGC,
		CallType:   "league_challenger_by_queue_" + region,
		CallKey:    mode,
	}
	if err = apiRequest("https://"+region+"."+Base+BaseLeague+"/challengerleagues/by-queue/"+mode, &lr, cp); err != nil {
		return lr, nil
	}

	return lr, err
}

// Grandmasters will return all the challengers in the requested queue.
// Valid queues are, Ranked5s(RANKED_SOLO_5x5), Flex3s(RANKED_FLEX_TT), and Flex5s(RANKED_FLEX_SR)
func Grandmasters(region string, mode string) (lr *models.MGCTierList, err error) {
	cp := cache.CachedParams{
		Cached:     true,
		Expire:     true,
		Expiration: ExpireMGC,
		CallType:   "league_grandmaster_by_queue_" + region,
		CallKey:    mode,
	}
	if err = apiRequest("https://"+region+"."+Base+BaseLeague+"/grandmasterleagues/by-queue/"+mode, &lr, cp); err != nil {
		return lr, nil
	}

	return lr, err
}

// Masters returns the roster of all the players in the Masters tier for requested queue.
func Masters(region string, mode string) (lr *models.MGCTierList, err error) {
	cp := cache.CachedParams{
		Cached:     true,
		Expire:     true,
		Expiration: ExpireMGC,
		CallType:   "league_master_by_queue_" + region,
		CallKey:    mode,
	}
	if err = apiRequest("https://"+region+"."+Base+BaseLeague+"/grandmasterleagues/by-queue/"+mode, &lr, cp); err != nil {
		return lr, nil
	}

	return lr, err
}

// League will return the roster of the league requested, via UUID.
// You can and will get blocked from this call if you provide invalid UUIDs
func League(region string, leagueUUID string) (lr *models.LeagueRoster, err error) {
	cp := cache.CachedParams{
		Cached:     true,
		Expire:     true,
		Expiration: LeagueByIDExpire,
		CallType:   "league_by_id_" + region,
		CallKey:    leagueUUID,
	}
	if err = apiRequest("https://"+region+"."+Base+BaseLeague+"/leagues/"+leagueUUID, &lr, cp); err != nil {
		return lr, nil
	}

	return lr, err
}

// PlayerPosition will return the requested players league position in each of the three ranked queues.
func PlayerPosition(region string, summonerID string) (lp *models.LeaguePosition, err error) {
	cp := cache.CachedParams{
		Cached:     true,
		Expire:     true,
		Expiration: ExpireGetPlayerPosition,
		CallType:   "league_position_by_summoner_" + region,
		CallKey:    summonerID,
	}
	if err = apiRequest("https://"+region+"."+Base+BaseLeague+"/positions/by-summoner/"+summonerID, &lp, cp); err != nil {
		return nil, err
	}

	return lp, nil
}
