package teaser

type FriendsList []string

var AddFriendToList func(username1 string, username2 string)
var LoadFriendsList func(username string) FriendsList

var FriendsLists = make(map[string]FriendsList)

func GetFriendsList(username string) FriendsList {
	if FriendsLists[username] == nil {
		FriendsLists[username] = LoadFriendsList(username)
	}

	return FriendsLists[username]
}
