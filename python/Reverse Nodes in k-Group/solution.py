# Definition for singly-linked list.
# class ListNode(object):
#     def __init__(self, val=0, next=None):
#         self.val = val
#         self.next = next
class Solution(object):
    def reverseKGroup(self, head, k):
        """
        :type head: ListNode
        :type k: int
        :rtype: ListNode
        """
        if k < 1: 
            return None
        
        if k == 1:
            return head
        else:
            cursor = head
            newHead = None
            lastTail = None
            firstGroup = True
            itr = 0
            while not cursor == None and not cursor.next == None:
                # Get the group to reorder
                subsetHead,subsetTail = self.getNthSizedGroup(cursor,k)
                if subsetHead == None:
                    print(cursor)
                    # size indivisable by k
                    return newHead
                
                # Record the state before doing in place reorder
                nextGroup = subsetTail.next
                # Isolate the current group
                subsetTail.next = None
                newSubsetHead, newSubsetTail = self.reverseSubset(subsetHead)
                if itr>0:
                    print(newSubsetHead)
                
                if firstGroup:
                    newHead = newSubsetHead
                    firstGroup = False
                else:
                    lastTail.next = newSubsetHead
                
                lastTail = newSubsetTail
                lastTail.next=nextGroup
                cursor = nextGroup
                itr = itr+1

            return newHead

    def printList(self, head):
        s = ""
        cursor = head
        while not cursor.next == None:
            s += str(cursor.val)
            cursor = cursor.next
        s += str(cursor.val)
        print(s)

    # returns head,tail of a subset
    def getNthSizedGroup(self, head, k):
        cursor = head
        for i in range(1,k):
            cursor, prev = self.advanceIndex(cursor)
            if cursor == None:
                return None,None
        return head, cursor

    def reverseSubset(self,head):
        
        if head.next == None:
            return head,None
        
        newHead = head
        cur = head
        
        while not cur.next == None:
            cur,prev = self.advanceIndex(cur)
            nex = cur.next
            # Make cur head
            cur.next = newHead
            newHead = cur
            prev.next = nex
            cur = prev
        return newHead,head


    # Returns index
    def advanceIndex(self, l): 
        if not l.next == None:
            return l.next,l
        else:
            return None,None
