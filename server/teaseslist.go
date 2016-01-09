package server

type TeasesList []Tease

var LoadTeasesList func(username string) TeasesList

var TeasesLists = make(map[string]TeasesList)

func GetTeasesList(username string) TeasesList {
	if TeasesLists[username] == nil {
		TeasesLists[username] = LoadTeasesList(username)
	}

	return TeasesLists[username]
}
