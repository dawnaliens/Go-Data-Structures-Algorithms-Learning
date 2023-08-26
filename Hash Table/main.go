package main

import (
	"fmt"
	"hash/fnv"
)

type HashTable struct {
	buckets [][]KeyValuePair
	size    int
}

type KeyValuePair struct {
	key   string
	value interface{}
}

// NewHashTable creates a new hash table with a given size
func NewHashTable(size int) *HashTable {
	return &HashTable{
		buckets: make([][]KeyValuePair, size),
		size:    size,
	}
}

// Hash function to determine the bucket for a given key
func (ht *HashTable) hash(key string) int {
	h := fnv.New32a()
	h.Write([]byte(key))
	return int(h.Sum32()) % ht.size
}

func (ht *HashTable) Put(key string, value interface{}) {
	bucketIndex := ht.hash(key)
	bucket := ht.buckets[bucketIndex]

	// Check if key already exists in the bucket
	for i, kv := range bucket {
		if kv.key == key {
			bucket[i].value = value
			return
		}
	}

	// Key doesn't exist, add it to the bucket
	ht.buckets[bucketIndex] = append(bucket, KeyValuePair{key, value})
}

func (ht *HashTable) Get(key string) (interface{}, bool) {
	bucketIndex := ht.hash(key)
	bucket := ht.buckets[bucketIndex]

	for _, kv := range bucket {
		if kv.key == key {
			return kv.value, true
		}
	}
	return nil, false
}

func (ht *HashTable) Delete(key string) {
	bucketIndex := ht.hash(key)
	bucket := ht.buckets[bucketIndex]

	for i, kv := range bucket {
		if kv.key == key {
			// Remove the key-value pair from the bucket
			ht.buckets[bucketIndex] = append(bucket[:i], bucket[i+1:]...)
			return
		}
	}
}

func main() {
	ht := NewHashTable(10)
	ht.Put("name", "John")
	ht.Put("age", 30)

	value, exists := ht.Get("name")
	if exists {
		fmt.Println("name:", value)
	} else {
		fmt.Println("name not found")
	}

	value, exists = ht.Get("age")
	if exists {
		fmt.Println("age:", value)
	} else {
		fmt.Println("age not found")
	}

	ht.Delete("name")
	value, exists = ht.Get("name")
	if exists {
		fmt.Println("name after deletion:", value)
	} else {
		fmt.Println("name not found after deletion")
	}
}
