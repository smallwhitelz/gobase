package main

type ListV1[T any] interface {
	Add(index int, val T)
	Append(val T) error
	Delete(index int) error
}

type LinkedListV1[T any] struct {
	head *NodeV1[T]
}

func (l LinkedListV1[T]) Add(index int, val T) {
	//TODO implement me
	panic("implement me")
}

func (l LinkedListV1[T]) Append(val T) error {
	//TODO implement me
	panic("implement me")
}

func (l LinkedListV1[T]) Delete(index int) error {
	//TODO implement me
	panic("implement me")
}

type NodeV1[T any] struct {
	data T
}

func UserList() {
	l := &LinkedListV1[int]{}
	l.Add(1, 1)
	//l.Add(1,"2") //会报错，强制约束
}
