package storage

import (
	"time"
)

type user struct {
	username string
	uid      int
	postlst  []post
}

type post struct {
	postDate time.Time
	body     []byte
	user     string
}
