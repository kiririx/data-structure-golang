package linked_list

type LinkedList struct {
	nodes []node
}

func (l *LinkedList) add(v int) {
	//n := node{v: v}

}

func (l *LinkedList) del() {

}

type node struct {
	nextPtr *node
	lastPtr *node
	v       int
}
