package main

import "fmt"

const ArraySize = 7

// HashTable will hold an array
type HashTable struct {
	array [ArraySize]*bucket
}

// bucket is a linked list in each slot of the array
type bucket struct {
	head *bucketNode
}

// bucketNode structure
type bucketNode struct {
	key string
	next *bucketNode
}

func (h *HashTable) Insert(key string) {
	index := hash(key)
	h.array[index].insert(key)
}

func (h *HashTable) Search(key string) bool {
	index := hash(key)
	return h.array[index].search(key)
}

func (h *HashTable) Delete(key string) {
	index := hash(key)
	h.array[index].delete(key)
}


// insert
func (b *bucket) insert(k string) {
	if !b.search(k) {
		newNode := &bucketNode{ key: k }
		newNode.next = b.head
		b.head = newNode
	} else {
		fmt.Println(k, "already exists")
	}
	
}

// search
func (b *bucket) search(k string) bool {
	currentNode := b.head
	for currentNode != nil {
		if currentNode.key == k {
			return true
		}
		currentNode = currentNode.next
	}
	return false
}

// delete
func (b *bucket) delete(k string) {

	if b.head.key == k {
		b.head = b.head.next
		return
	}

	previousNode := b.head
	for previousNode.next != nil {
		if previousNode.next.key == k {
			previousNode.next = previousNode.next.next
		}
		previousNode = previousNode.next
	}
}

// hash function
func hash(key string) int {
	sum := 0
	for _, v := range key {
		sum += int(v)
	}

	return sum % ArraySize
}

// Init will create a bucket in each slot of the hash table
func Init() *HashTable{
	result := &HashTable{}
	for i := range result.array {
		result.array[i] = &bucket{}
	}
	return result
}


func main() {
	testHashTable := Init()
	testBucket := &bucket{}
	
	testBucket.insert("RANDY")
	testBucket.insert("RANDY") // "RANDY already exists"
	testBucket.delete("RANDY")
	fmt.Println(!testBucket.search("RANDY"))
	testBucket.insert("RANDY")
	fmt.Println(testBucket.search("RANDY"))
	fmt.Println(!testBucket.search("SALLY"))

	testHashTable.Insert("RANDY")
	fmt.Println(testHashTable.Search("RANDY"))
	testHashTable.Delete("RANDY")
	fmt.Println(!testHashTable.Search("RANDY"))
	fmt.Println(!testHashTable.Search("SALLY"))	
}