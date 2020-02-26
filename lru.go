package goLRU

import (
	"fmt"

	queue "github.com/souvikhaldar/go-ADT/queue"
)

type lruCache struct {
	cache  map[string]interface{}
	length int
	size   int
	queue  *queue.Queue
}

var (
	ErrCacheFull   = fmt.Errorf("Cache is full")
	ErrKeyNotFound = fmt.Errorf("The given key is not found in the cache")
)

func NewLRUCache(s int) *lruCache {
	return &lruCache{
		cache:  make(map[string]interface{}),
		length: 0,
		size:   s,
		queue:  queue.NewQueue(s),
	}
}

// Set receives (key,value) as arguments and puts it in the cache
func (l *lruCache) Set(key string, value interface{}) error {
	fmt.Println("Recieved key and value: ", key, value)
	if l.length == l.size {
		fmt.Println(ErrCacheFull)
		valToBeDeletedFromCache := l.queue.PeekRear()
		// need to evict LRU node from queue and map
		if err := l.queue.Enqueue(); err != nil {
			return err
		}
		delete(l.cache, valToBeDeletedFromCache)
		l.length--
	}

	// add it to cache map
	l.cache[key] = value
	// add the entry to queue
	if err := l.queue.Enqueue(key); err != nil {
		return err
	}
	return nil
}

// Get receives key as the argument and returns the value
// associated with the key
func (l *lruCache) Get(key string) (interface{}, error) {
	value, ok := l.cache[key]
	if !ok {
		return nil, ErrKeyNotFound
	}

	if err := l.queue.MoveNodeToEnd(key); err != nil {
		fmt.Println(err)
	}

	return value, nil
}
