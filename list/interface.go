package list

type Iterator interface {
	HashNext() bool
	Next() interface{}
}
