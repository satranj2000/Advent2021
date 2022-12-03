package utils

import "fmt"

type Node struct {
	data string
	next *Node
}

func SetupNode(val string) *Node {
	var n Node = Node{data: val, next: nil}
	return &n
}

func (head *Node) Add(val string) {
	if head == nil {
		head = SetupNode(val)
	} else {
		for head.next != nil {
			head = head.next
		}
		newnode := SetupNode(val)
		head.next = newnode
	}
}

func (curr *Node) AddAfter(val string) bool {
	if curr == nil {
		return false
	} else {
		newnode := SetupNode(val)
		newnode.next = curr.next
		curr.next = newnode
	}
	return true
}

func (head *Node) ListAll() {
	for head != nil {
		fmt.Println(head.data)
		head = head.next
	}
}

func (head *Node) ToArray() []string {
	str := make([]string, 0)
	for head != nil {
		val := head.data
		str = append(str, val)
		head = head.next
	}
	return str
}

func (head *Node) TwonodesFound(lside string, rside string) []*Node {

	foundNodes := make([]*Node, 0)
	head2 := head
	for head2 != nil {
		if head2.next != nil {
			if head2.data == lside && head2.next.data == rside {
				foundNodes = append(foundNodes, head2)
			}
		}
		head2 = head2.next
	}
	return foundNodes
}
