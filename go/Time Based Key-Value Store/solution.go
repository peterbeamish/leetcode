type PriorityQueue []*Item

type Item struct {
    value string
    priority int
    index int
}

func (pq PriorityQueue) Len() int {return len(pq)}
func (pq PriorityQueue) Less(i,j int) bool {
    return pq[i].priority > pq[j].priority
}
func (pq PriorityQueue) Swap(i,j int){
    pq[i],pq[j] = pq[j],pq[i]
    pq[i].index = i
    pq[j].index = j
}

func (pq *PriorityQueue) Push(x interface{}) {
	n := len(*pq)
	item := x.(*Item)
	item.index = n
	*pq = append(*pq, item)
}

func (pq *PriorityQueue) Pop() interface{} {
	old := *pq
	n := len(old)
	item := old[n-1]
	old[n-1] = nil  // avoid memory leak
	item.index = -1 // for safety
	*pq = old[0 : n-1]
	return item
}

// update modifies the priority and value of an Item in the queue.
func (pq *PriorityQueue) update(item *Item, value string, priority int) {
	item.value = value
	item.priority = priority
	heap.Fix(pq, item.index)
}


type TimeMap struct {
    data map[string]*PriorityQueue
}


func Constructor() TimeMap {
    return TimeMap{
        data: make(map[string]*PriorityQueue),
    }
}


func (this *TimeMap) Set(key string, value string, timestamp int)  {
    // Check if the key already exists
    // If the key exists, and the PriorityQueue is not nil, construct the item and insert
    if pq, ok := this.data[key]; ok{
        
        itm := &Item{
            value: value,
            priority: timestamp,
            index: len(*pq)-1,
        }
        if pq != nil{            
            heap.Push(pq,itm)       
        } else {
            // TODO handle error
        }
    } else {
        // If the key does not exist, allocate the PriorityQueue, push the item, insert into data
        itm := &Item{
            value: value,
            priority: timestamp,
            index: 0,
        }
        
        pq := make(PriorityQueue,1)
        pq[0] = itm
        
        // TODO check if needs error handling
        heap.Init(&pq)
        
        this.data[key] = &pq
    }
}


func (this *TimeMap) Get(key string, timestamp int) string {
    if pq, ok := this.data[key]; ok{
        if pq != nil {
            
            var target *Item = nil
            unusedItems := []*Item{}
            for len(*pq)!= 0 {
                itm, _ := (heap.Pop(pq)).(*Item)
                if itm != nil {
                    if itm.priority > timestamp {
                        unusedItems = append(unusedItems, itm)
                    } else {
                        // Item found!
                        target = itm
                        unusedItems = append(unusedItems, itm)
                        break
                    }
                } else {
                    // TODO handle nil item previously inserted
                    // LOG?
                }
            }
            
            for _,itm := range unusedItems {
                heap.Push(pq,itm)    
            }
            
            if target != nil {
                return target.value
            }
        }
            
            
        } else {
            // TODO handle error
        }
    
        return ""
    } 
    
 
/**
 * Your TimeMap object will be instantiated and called as such:
 * obj := Constructor();
 * obj.Set(key,value,timestamp);
 * param_2 := obj.Get(key,timestamp);
 */