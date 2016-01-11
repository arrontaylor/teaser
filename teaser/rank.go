package teaser

type Rank *string

var LoadRank func(sentcount int) Rank

var Ranks = make(map[int]Rank)

func GetRank(sentcount int) Rank {
	if Ranks[sentcount] == nil {
		Ranks[sentcount] = LoadRank(sentcount)
	}

	return Ranks[sentcount]
}
