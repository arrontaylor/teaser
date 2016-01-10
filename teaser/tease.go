package teaser

import "time"

type Tease struct {
	FromUsername string
	ToUsername   string
	Time         time.Time
	Read         bool
}

var CreateTease func(fromusername string, tousername string) *Tease
