package db

import (
	log "github.com/Sirupsen/logrus"
	"github.com/arrontaylor/teaser/teaser"
)

func LoadRank(sentcount int) teaser.Rank {
	row := Connection.QueryRow("SELECT max(count), name FROM ranks WHERE count<=?", sentcount)

	var rank string
	var count int
	err := row.Scan(&count, &rank)

	if err != nil {
		log.Warn(err.Error())
		return nil
	}

	return &rank
}
