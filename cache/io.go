package cache

import (
	"errors"
	"fmt"
	"github.com/jackc/pgx"
	"github.com/json-iterator/go"
	"strings"
	"time"
	"unsafe"
)

var (
	ErrExpired = errors.New("call expired")
)

// ReadCache will look to see if the request is in the cache, and if it has expired, and will load the data into the interface.
func ReadCache(s interface{}, cp *CachedParams) (err error) {
	if cp.Cached == true && Enabled == true {
		var Resp string
		var Updated time.Time
		if er := CDB.QueryRow(`SELECT updated_at, response FROM `+strings.ToLower(cp.CallType)+` WHERE key=$1`, cp.CallKey).Scan(&Updated, &Resp); er != pgx.ErrNoRows {
			if cp.Expire == false || (time.Since(Updated) > cp.Expiration) == false {
				if er := jsoniter.UnmarshalFromString(Resp, &s); er != nil {
					return er
				}
				return
			}
			return ErrExpired

		}
		return
	}
	return
}

func StoreCache(cp *CachedParams, resp []byte) (err error) {
	fmt.Println(cp.CallKey)
	fmt.Println("store cache called")
	if cp.Cached == true && Enabled == true {
		fmt.Println("cache passed valid")
		err := CDB.QueryRow("SELECT updated_at FROM "+cp.CallType+" WHERE key=$1", cp.CallKey).Scan()
		fmt.Println(err)
		switch err {
		case pgx.ErrNoRows:
			if _, er := CDB.Exec(`INSERT INTO `+strings.ToLower(cp.CallType)+`(created_at, updated_at, key, response) VALUES($1, $1, $2, $3)`, time.Now(), cp.CallKey, ByteSlice2String(resp)); er != nil {
				fmt.Println(er)
				return er
			}
		default:
			if _, er := CDB.Exec(`UPDATE `+strings.ToLower(cp.CallType)+` SET updated_at = $1, response = $2 WHERE key = $3`, time.Now(), ByteSlice2String(resp), cp.CallKey); err != nil {
				fmt.Println(er)
				return er
			}
		}
	}
	return
}

func ByteSlice2String(bs []byte) string {
	return *(*string)(unsafe.Pointer(&bs))
}
