package teaser

const FREQUENTS_LIST_LENGTH = 3

type FrequentsList map[string]int

var FrequentsLists = make(map[string]FrequentsList)

var LoadFrequentsList func(username string) FrequentsList

func GetFrequentsList(username string) FrequentsList {
	if FrequentsLists[username] == nil {
		FrequentsLists[username] = LoadFrequentsList(username)
	}

	return FrequentsLists[username]
}
