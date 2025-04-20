package steamclient

type GetFriendListFilter string

func (s GetFriendListFilter) String() string {
	return string(s)
}

var (
	GetFriendListFilterAll     GetFriendListFilter = "All"
	GetFriendListFilterFriends GetFriendListFilter = "Friends"
)
