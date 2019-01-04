package cache

import (
	"errors"
	"fmt"
	"github.com/jackc/pgx"
	"github.com/json-iterator/go"
	"time"
	"unsafe"
)

var (
	// ErrExpired is the error that is used when cached call expires
	ErrExpired = errors.New("call expired")

	// ErrNoCache is the error that is used when requested call doesn't get cached.
	ErrNoCache = errors.New("call will not be cached")

	// ErrNoData is the error that is used when data that doesn't expire isn't in database
	ErrNoData = errors.New("call is not cached")
)

// ReadCache will look to see if the request is in the cache, and if it has expired, and will load the data into the interface.
func ReadCache(s interface{}, cp *CachedParams) (err error) {
	if cp.Cached == true && Enabled == true {
		var Resp string
		var Updated time.Time
		if er := CDB.QueryRow(`SELECT updated_at, response FROM "`+cp.CallType+`" WHERE key=$1`, cp.CallKey).Scan(&Updated, &Resp); er != pgx.ErrNoRows {
			if cp.Expire == false || (time.Since(Updated) > cp.Expiration) == false {
				if er := jsoniter.UnmarshalFromString(Resp, &s); er != nil {
					return er
				}
				return
			}
			return ErrExpired

		}
		return ErrNoCache
	}
	return ErrNoCache
}

// StoreCache will store the call into the database if requested.
func StoreCache(cp *CachedParams, resp []byte) (err error) {
	if cp.Cached == true && Enabled == true {
		switch CDB.QueryRow("SELECT updated_at FROM "+cp.CallType+" WHERE key=$1", cp.CallKey).Scan() {
		case pgx.ErrNoRows:
			if _, er := CDB.Exec(`INSERT INTO "`+cp.CallType+`"(created_at, updated_at, key, response) VALUES($1, $1, $2, $3)`, time.Now(), cp.CallKey, bytesSlice2String(resp)); er != nil {
				fmt.Println(er)
				return er
			}
		default:
			if _, er := CDB.Exec(`UPDATE "`+cp.CallType+`" SET updated_at = $1, response = $2 WHERE key = $3`, time.Now(), bytesSlice2String(resp), cp.CallKey); err != nil {
				fmt.Println(er)
				return er
			}
		}
	}
	return
}

func bytesSlice2String(bs []byte) string {
	return *(*string)(unsafe.Pointer(&bs))
}
