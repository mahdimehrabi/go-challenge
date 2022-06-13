package interfaces

type MemoryDB interface {
	//save a sorted set
	SaveSortedSet(key string, member string, score string)

	//get a sorted set
	GetSortedSet(key string, member string, score string)

	//get list of sorted set by key in reverse order
	RevListStoredSet(key string)
}
