package main

import "fmt"

type node struct {
	data int
	next *node
}

func createNode(val int) *node {
	return &node{
		data: val,
		next: nil,
	}
}

func insertAtFirst(val int, head **node) {
	var temp *node = createNode((val))
	temp.next = *head
	*head = temp
}

func insertAtLast(val int, root **node) {
	var temp *node = createNode((val))
	(*root).next = temp
	(*root) = temp
}
func deleteAtFirst(head **node) {
	var temp *node = (*head).next
	(*head).next = nil
	*head = temp

}
func deleteAtLast(head ,root **node) {
	var temp *node = *head
	for temp.next.next != nil {
		temp = temp.next
	}
	temp.next=nil
	*root=temp
}

func display(head *node) {
	for head != nil {
		fmt.Println(head.data)
		head = head.next
	}
}
func deleteElement(value int,head **node){
	var temp *node=*head
	for temp.next.data!=value {
		temp=temp.next	
	}
	temp.next=temp.next.next
	//fmt.Println(temp.data)

}
func main() {
	var head, root *node
	head = createNode(1)
	root = head
	insertAtFirst(0, &head)
	insertAtLast(2, &root)
	insertAtLast(3, &root)

	
	insertAtLast(4, &root)
	insertAtLast(5, &root)
	deleteAtFirst(&head)
	//deleteAtLast(&head,&root)
	deleteElement(3,&head)

	display(head)
	//fmt.Println(root.data)

}
