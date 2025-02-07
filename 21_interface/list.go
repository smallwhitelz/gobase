package main

type List interface {
	Add(index int, value any)
	Append(val any) error
	Delete(index int) error
}

type LinkedList struct {
	head node
}

func (l *LinkedList) Append(val any) error {
	//TODO implement me
	panic("implement me")
}

func (l *LinkedList) Delete(index int) error {
	//TODO implement me
	panic("implement me")
}

type node struct {
	next *node
}

func (l *LinkedList) Add(index int, value any) {
	// TODO 实现这个方法
}

func UserListV1() {
	l := &LinkedList{}
	l.Add(1, "1")
	l.Add(2, 2)
	l.Add(3, nil)

}
