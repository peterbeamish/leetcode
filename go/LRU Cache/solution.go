type LRUCache struct {
    dll *DLL
    // key to elem in the linked list
    lookup map[int]*Node
    
    capacity int
    
}

type Node struct {
    next *Node
    last *Node
    key int
    val int
}

type DLL struct {
    head *Node
    tail *Node
}

func (l *DLL) Add(node *Node){
    if l.head == nil {
        l.head = node
    } else {
        l.head.last = node
        node.next = l.head
        node.last = nil
        l.head = node
    }
    
    if l.tail == nil {
        l.tail = node
    }
    
}

func (l *DLL) MoveFront(node *Node){
    if node == l.head{
        return
    }
    
    last := node.last
    next := node.next
    
    if last != nil {
        last.next = next
    }
    if next != nil {
        next.last = last
    }
    oldHead := l.head
    
    oldHead.last = node
    node.next = oldHead
    node.last = nil
    l.head = node
    
    if l.tail == node && last != nil{
        l.tail = last
    }
}

func (l* DLL) DeleteTail() *Node{
    tTail := l.tail
    last := tTail.last
    
    l.tail = last
    tTail.last = nil
    
    
    return tTail
}


func Constructor(capacity int) LRUCache {
    return LRUCache{
        lookup: make(map[int]*Node),
        dll: &DLL{},
        capacity: capacity,
    }
}


func (this *LRUCache) Get(key int) int {
    // check if the key exists
    if nod, ok := this.lookup[key]; ok{
        
        this.dll.MoveFront(nod)
        
        return nod.val
    }
    return -1
}


func (this *LRUCache) Put(key int, value int)  {
    
    var node *Node
    var ok bool
    // Check if the key already exist. Get it to update the val.
    if node, ok = this.lookup[key]; !ok{
        // If it doesn't exist, allocate. Set val.
        node = &Node{
            key: key,
            val: value,
        }
        
        this.dll.Add(node)
        
    } else {
        node.val = value
        this.dll.MoveFront(node)
    }
    
    // Insert in dict
    this.lookup[key] = node
    
    // if len(lookup)>capacity, free the tail and adjust tail's last to be tail. Free the node from the dict as well.
    if len(this.lookup)>this.capacity{
        
        tTail := this.dll.DeleteTail()
        
        delete(this.lookup,tTail.key)
    }
}