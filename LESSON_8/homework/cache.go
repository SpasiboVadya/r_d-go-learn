package lru

type LruCache interface {
	Put(key, value string)
	Get(key string) (string, bool)
}

type node struct {
	key   string
	value string
	prev  *node
	next  *node
}

type lruCache struct {
	capacity int
	cache    map[string]*node
	head     *node
	tail     *node
}

func NewLruCache(capacity int) LruCache {
	return &lruCache{
		capacity: capacity,
		cache:    make(map[string]*node),
	}
}

func (l *lruCache) Put(key, value string) {
	if node, exists := l.cache[key]; exists {
		// Update existing node
		node.value = value
		l.moveToFront(node)
		return
	}

	// Create new node
	newNode := &node{
		key:   key,
		value: value,
	}

	// Add to cache
	l.cache[key] = newNode
	l.addToFront(newNode)

	// Check capacity
	if len(l.cache) > l.capacity {
		l.removeLRU()
	}
}

func (l *lruCache) Get(key string) (string, bool) {
	if node, exists := l.cache[key]; exists {
		l.moveToFront(node)
		return node.value, true
	}
	return "", false
}

func (l *lruCache) addToFront(node *node) {
	if l.head == nil {
		l.head = node
		l.tail = node
		return
	}

	node.next = l.head
	l.head.prev = node
	l.head = node
}

func (l *lruCache) moveToFront(node *node) {
	if node == l.head {
		return
	}

	// Remove from current position
	if node.prev != nil {
		node.prev.next = node.next
	}
	if node.next != nil {
		node.next.prev = node.prev
	}
	if node == l.tail {
		l.tail = node.prev
	}

	// Add to front
	node.next = l.head
	node.prev = nil
	l.head.prev = node
	l.head = node
}

func (l *lruCache) removeLRU() {
	if l.tail == nil {
		return
	}

	// Remove from cache map
	delete(l.cache, l.tail.key)

	// Remove from linked list
	if l.tail.prev != nil {
		l.tail.prev.next = nil
		l.tail = l.tail.prev
	} else {
		l.head = nil
		l.tail = nil
	}
}
