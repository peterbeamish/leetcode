from heapq import heappop, heappush, heapify

class MaxHeap(object):
    def __init__(self):
        self.heap = []
        heapify(self.heap)
    
    def push(self, n):
        n = -1*n
        heappush(self.heap, n)
    def top(self):
        return self.heap[0]*-1
    def pop(self):
        return heappop(self.heap) *1
    def size(self):
        return len(self.heap)

class MinHeap(object):
    def __init__(self):
        self.heap = []
        heapify(self.heap)
    
    def push(self, n):
        heappush(self.heap, n)
    def top(self):
        return self.heap[0]
    def pop(self):
        return heappop(self.heap)
    def size(self):
        return len(self.heap)

class MedianFinder(object):
    def __init__(self):
        self.low = MaxHeap()
        self.hi = MinHeap()    

    def addNum(self, num):
        """
        :type num: int
        :rtype: None
        """
        self.low.push(num)
        self.hi.push(self.low.top())
        self.low.pop()
        
        if (self.low.size() < self.hi.size()):
            self.low.push(self.hi.top())
            self.hi.pop()

    def findMedian(self):
        if self.low.size() > self.hi.size():
            return self.low.top()
        else:
            topOfLow = self.low.top()
            topOfHi = self.hi.top()
            return (topOfLow + topOfHi)/2.0

# Your MedianFinder object will be instantiated and called as such:
# obj = MedianFinder()
# obj.addNum(num)
# param_2 = obj.findMedian()