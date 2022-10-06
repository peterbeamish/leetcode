type IntHeap []int

func (h IntHeap) Len() int           { return len(h) }
func (h IntHeap) Less(i, j int) bool { return h[i] < h[j] }
func (h IntHeap) Swap(i, j int)      { h[i], h[j] = h[j], h[i] }

func (h *IntHeap) Push(x interface{}) {
	// Push and Pop use pointer receivers because they modify the slice's length,
	// not just its contents.
	*h = append(*h, x.(int))
}

func (h *IntHeap) Pop() interface{} {
	old := *h
	n := len(old)
	x := old[n-1]
	*h = old[0 : n-1]
	return x
}

func (h *IntHeap) Peek() interface{} {
    if len(*h) == 0 {
        return nil
    }
    return (*h)[0]
}


func minMeetingRooms(intervals [][]int) int {
    if len(intervals) == 0 {
        return 0
    }
    
    // Sort the collection by start time
    sort.Slice(intervals, func(i,j int) bool {
        return intervals[i][0]<intervals[j][0]
    })
        
    // allocate a minheap
    mh := &IntHeap{}
    heap.Init(mh)
    
    // push the first meeting's end time
    heap.Push(mh, intervals[0][1])
    
    // loop remaining meetings
    for i:=1; i<len(intervals); i++ {
        
        activeMeeting,_ := (mh.Peek()).(int)
        nextMeeting := intervals[i][0]
        
        if nextMeeting>= activeMeeting {
            heap.Pop(mh)
        }
        heap.Push(mh, intervals[i][1])
    }    
    return mh.Len()
}