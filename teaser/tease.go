package teaser

import "time"

type Tease struct {
	TeaseId      int
	FromUsername string
	ToUsername   string
	Time         time.Time
	Read         bool
}

var CreateTease func(fromusername string, tousername string) *Tease
var ReadTease func(teaseid int, tousername string)
