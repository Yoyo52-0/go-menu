package main

import (
	"fmt"
	"os"
)

type LinkNode struct {
	next    *LinkNode
	command string
	exec_cb call_back
}

type LinkTable struct {
	head *LinkNode //头节点
	tail *LinkNode //尾节点
	len  int       //链表长度
}

type call_back interface {
	exec()
}

func initTable() *LinkTable {
	table := &LinkTable{
		head: nil,
		tail: nil,
		len:  0,
	}
	return table
}

func (this *LinkTable) claerTable() { //未使用
	this.head = nil
	this.tail = nil
	this.len = 0
	println("table destroyed")
	return
}

func createCommand(command string, call_back call_back) *LinkNode {
	newCommand := &LinkNode{
		next:    nil,
		command: command,
		exec_cb: call_back,
	}
	return newCommand
}

func (this *LinkTable) findNode(command string) *LinkNode {
	for temp := this.head; temp != nil; temp = temp.next {
		if temp.command == command {
			println("found command %s", command)
			return temp
		}
	}
	println("didn't find command %s", command)
	return nil
}

func (this *LinkTable) execCommand(command string) error {
	node := this.findNode(command)
	if node == nil {
		println("no such cmd!")
		return nil
	}
	node.exec_cb.exec()
	return nil
}

func (this *LinkTable) insertNode(node *LinkNode) {
	if this.head == nil { //首先判空
		this.head = node
		this.tail = node
		this.len = 1
		return
	}

	if this.findNode(node.command) != nil { //已存在的指令
		println("existing command!")
		return
	}

	this.tail.next = node
	this.tail = this.tail.next
	this.len++
	return
}

/*
func (this *LinkTable) insertNode(command string, exec_cb call_back) {
	if this.findNode(command) != nil { //已存在的指令
		println("existing command!")
		return
	}

	node := createCommand(command, exec_cb)

	if this.head == nil { //首先判空
		this.head = node
		this.tail = node
		this.len = 1
		return
	}

	this.tail.next = node
	this.tail = this.tail.next
	this.len++
	return
}
*/
func (this *LinkTable) deleteNode(index int) { //暂时不需要使用
	if index < 0 || index >= this.len {
		println("invalid index")
		return
	}
	node := this.head
	if index == 0 {
		this.head = node.next
	} else {
		node := this.head
		for i := 0; i < index-2; i++ {
			node = node.next
		}
		node.next = node.next.next
	}
	this.len--
}

type help struct {
}

func (this help) exec() {
	println("This is help command!")
}

type quit struct {
}

func (this quit) exec() {
	println("You have successfully exited")
	os.Exit(1)
}

func main() {
	table := initTable()
	cmd1 := createCommand("help", new(help))
	cmd2 := createCommand("quit", new(quit))
	table.insertNode(cmd1)
	table.insertNode(cmd2)
	//list.delectNode(cmd2)
	for {
		menu()
		var command string
		fmt.Scan(&command)
		table.execCommand(command)
	}
}
func menu() {
	fmt.Print("\tMenu with linktable\t\n>")
	fmt.Print("\tEnter your choice\t\n>")
	fmt.Print("\t1. help\t\t2. quit\t\n>")
}
