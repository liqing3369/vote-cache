package constant

import "time"

const (
	ACTIVITY = "VOTE:ACTIVITY"
	PLAYER   = "VOTE:PLAYER"

	EXPIRE_TIME = time.Duration(60*24) * time.Hour
)
