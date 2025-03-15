package Storage

import (
	"fmt"
	"sync"
	"time"
)

type Node struct {
	Value *TTL
	Next *Node
}

type LinkedList struct {
	mu sync.RWMutex
	Head *Node
	Tail *Node
}

func GetHead() *LinkedList {
	return &LinkedList{}
}

func (l *LinkedList) Add(data *TTL) {
	l.mu.Lock()
	defer l.mu.Unlock()
	if (l.Head == nil) {
		l.Head = &Node{Value: data}
		l.Tail = l.Head
	} else if(l.Head.Value.ttl.After(data.ttl)){
		l.Head = &Node{Value: data, Next: l.Head}
	}else if(l.Tail.Value.ttl.Before(data.ttl)){
		l.Tail.Next = &Node{Value: data}
		l.Tail = l.Tail.Next
	} else {
		temp := l.Head
		node := &Node{Value: data}
		for temp.Next != nil && !node.Value.ttl.Before(temp.Next.Value.ttl) {
			temp = temp.Next
		}
		node.Next = temp.Next
		temp.Next = node
	}
}

func (l *LinkedList) Delete(key string) {
	l.mu.Lock()
	defer l.mu.Unlock()
	temp := l.Head
	if l.Head == l.Tail {
		l.Head = nil
		l.Tail = nil
		return
	} else if (*temp.Value.key == key) {
		l.Head = temp.Next
		temp.Next = nil
		temp = nil
		return
	}
	for temp != nil && temp.Next != nil {
		fmt.Println("Checking if key: "+key+" is in linked list")
		if temp.Next.Value.key == &key {
			val := temp.Next
			temp.Next = temp.Next.Next
			val.Next = nil
			val = nil
			if temp.Next == nil {
				l.Tail = temp
			}
			break
		}
		temp = temp.Next
	}
	
}

func (db *Db) List() {
	db.link.mu.RLock()
	defer db.link.mu.RUnlock()
	temp := db.link.Head
	for temp != nil {
		fmt.Println("Key : "+(*temp.Value.key)+" Expires in : "+time.Until(temp.Value.ttl).String())
		temp = temp.Next
	}
}

func (db *Db) DropLink() {
	db.link.mu.Lock()
	defer db.link.mu.Unlock()
	db.link.Head = nil
	db.link.Tail = nil

}

func (l *LinkedList) Sweep(ttldb *TTLDB, db *Db) {
	for {
		l.AutoSweep(ttldb, db)
		time.Sleep(time.Millisecond* 500)
	}
}
func (l *LinkedList) AutoSweep(ttldb *TTLDB, db *Db) {
	temp := l.Head
	l.mu.Lock()
	defer l.mu.Unlock()
	for temp != nil && temp.Value.ttl.Before(time.Now()) {
		go db.Rmttl(*temp.Value.key)
		temp.Value.Data = nil
		temp = temp.Next
	}
	l.Head = temp
}