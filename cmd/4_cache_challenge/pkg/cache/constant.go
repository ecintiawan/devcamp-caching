package cache

import "time"

const (
	CacheFiveSeconds  = 5 * time.Second
	CacheOneMinute    = time.Minute
	CacheThreeMinutes = 3 * CacheOneMinute
	CacheFiveMinutes  = 5 * CacheOneMinute
)
