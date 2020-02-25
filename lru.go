package go-lru

type lru interface{
	Get(key string)interface{}
	Set(key string,value interface{})error
}
type lruCache struct{
	cache map[string]interface{}
	length int
	size int 
	queue *Queue
}

var(
	ErrCacheFull = fmt.Errorf("Cache is full")
)


func newLruCache(s int)*lruCache{
	return &lruCache{
		cache: make(map[string]interface{}),
		length: 0,
		size: s,
		queue: nil,
	}
}

func (l *lruCache)Set(key string,value interface{})error{
	fmt.Println("Recieved key and value: ",key,value)
	if l.length == l.size {
		fmt.Println(ErrCacheFull)
	}
	// need to evict LRU node
	
}