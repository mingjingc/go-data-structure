package list

type Iterator interface {
	HashNext() bool
	Next() interface{}
}

type ISkipList interface {
	Find(key int) interface{}
	Put(key int, value interface{})
	Delete(key int)

	randomLevel()int
}
