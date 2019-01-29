package ggriot

import (
	"fmt"
	"log"
	"testing"

	"github.com/go-redis/redis"
	"github.com/tyhi/ggriot/cache"
)

func TestActiveGame(t *testing.T) {
	client := redis.NewClient(&redis.Options{
		Addr: "127.0.0.1:6379",
		Password: "",
		DB: 0,
	})

	SetAPIKey("")
	err := cache.UseCache(client)
	if err != nil {
		t.Error(err)
	}

}

func TestGetAc(t *testing.T) {
	e, err := SummonerByIGN(NA, "soowan")
	if err != nil {
		t.Error(err)
	}

	log.Println("sum id", e.ID)
	log.Println("acc id", e.AccountID)
}

func TestGetTotalMasteryLevel(t *testing.T) {
	e, err := TotalMasteryLevel(NA, "att1mlWZh48J3gVjokJ1NH9h2URkUq4HtGsV8RSEPNWzVv8")
	if err != nil {
		t.Error(err)
	}

	log.Println(e)
}

func TestGetChallengers(t *testing.T) {
	e, err := Challengers(NA, Ranked5s)
	if err != nil {
		t.Error(err)
	}

	fmt.Println(e)
}

func TestGetMasters(t *testing.T) {
	e, err := Masters(NA, Flex3s)
	if err != nil {
		t.Error(err)
	}

	fmt.Println(e)
}

func TestGetPlayerPosition(t *testing.T) {
	_, err := PlayerPosition(NA, "att1mlWZh48J3gVjokJ1NH9h2URkUq4HtGsV8RSEPNWzVv8")
	if err != nil {
		t.Error(err)
	}
}

func TestGetMatch(t *testing.T) {
	e, err := Match(NA, 2872782472)
	if err != nil {
		t.Error(err)
	}

	fmt.Println(e.GameVersion)
}

func TestGetMatchTimeline(t *testing.T) {
	e, err := MatchTimeline(NA, 2872782472)
	if err != nil {
		t.Error(err)
	}

	fmt.Println(e.Frames[10].ParticipantFrames.Num1.CurrentGold)
}

func TestSummonerByIGN(t *testing.T) {
	e, err := SummonerByIGN(NA, "Soowan")
	if err != nil {
		t.Error(err)
	}

	fmt.Println(e.ID)
}

func TestMatchHistory(t *testing.T) {
	s, err := SummonerByIGN(NA, "Soowan")
	if err != nil {
		t.Fatal(err)
	}
	mh, err := MatchHistory(NA, s.AccountID)
	if err != nil {
		t.Fatal(err)
	}

	fmt.Println(mh.Matches[0].GameID)
}

func TestClosing(t *testing.T) {
	err := cache.RedisConn.Close()
	if err != nil {
		t.Fatal(err)
	}
}
