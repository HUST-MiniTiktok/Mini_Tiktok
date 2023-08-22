package oss

import (
	"time"
)

const (
	RDCacheExpire = time.Hour * 24
)

func RDExistURLMaping(db_url string) bool {
	return RDClient.Exists(db_url)
}

func RDGetURLMaping(db_url string) string {
	return RDClient.Get(db_url)
}

func RDSetURLMaping(db_url string, real_url string) {
	RDClient.Set(db_url, real_url, RDCacheExpire)
}
