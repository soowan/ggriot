package cache

import (
	"fmt"
	"github.com/jackc/pgx"
	"log"
	"strings"
	"sync"
	"time"
)

// CDB is the exported pointer to the opened cache server.
var CDB *pgx.ConnPool

// Enabled this is checked to see if ggriot should call the postgres db first or skip calling cache.
var Enabled = false

// UseCache will open a connection to a postgres server that will be used as a cache server.
func UseCache(gostring string) (err error) {
	pgxConfig, err := pgx.ParseConnectionString(gostring)
	if err != nil {
		return err
	}

	poolConfig := pgx.ConnPoolConfig{
		ConnConfig:     pgxConfig,
		MaxConnections: 80,
		AcquireTimeout: time.Duration(time.Second * 3),
	}

	CDB, err = pgx.NewConnPool(poolConfig)
	if err != nil {
		return err
	}

	Enabled = true

	// This will create all the tables needed, I'm not sure how to make this better yet, but I know it probably
	// needs to be changed.
	r := []string{
		"RU",
		"KR",
		"BR1",
		"OC1",
		"JP1",
		"NA1",
		"EUN1",
		"EUW1",
		"TR1",
		"LA1",
		"LA2",
	}

	c := []string{
		"mastery_by_summoner",
		"mastery_by_champion",
		"mastery_level",
		"champion_rotation",
		"league_by_queue",
		"league_by_id",
		"league_master_by_queue",
		"league_grandmaster_by_queue",
		"league_challenger_by_queue",
		"league_position_by_summoner",
		"league_match_by_id",
		"league_match_tl_by_id",
		"summoner_by_ign",
		"summoner_by_puuid",
	}

	var wg sync.WaitGroup
	for rr := range r {
		for cc := range c {
			go func(one string, two string) {
				wg.Add(1)
				fmt.Println("checking ", one+"_"+two)
				var d int
				if err := CDB.QueryRow("select 1 from information_schema.tables where table_name=$1", strings.ToLower(one+"_"+two)).Scan(&d); err == pgx.ErrNoRows {
					fmt.Println(err)
					_, err := CDB.Exec(`create table ` + strings.ToLower(one+"_"+two) + `(created_at timestamp with time zone, updated_at timestamp with time zone, key text, response jsonb)`)
					fmt.Println("creating ", strings.ToLower(one+"_"+two))
					if err != nil {
						log.Fatal(err)
					}
				}
				wg.Done()
			}(c[cc], r[rr])
		}
	}
	wg.Wait()
	return nil
}
