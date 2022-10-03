from Queue import PriorityQueue

# Definition for singly-linked list.
# class ListNode(object):
#     def __init__(self, val=0, next=None):
#         self.val = val
#         self.next = next
class Solution(object):
    def mergeKLists(self, lists):
        """
        :type lists: List[ListNode]
        :rtype: ListNode
        """
        
        # PriorityQueue was selected for its sorting characteristics
        q = PriorityQueue()

        # iterate thru each nested list
        for ll in lists:
            # Put the first elem (head of each ll in the queue)

            if not ll == None:
                q.put((ll.val,ll))

        head = None
        last = None
        while not q.empty():
            val, node = q.get()
            if head == None:
                head = node
                last = head
            else:
                last.next = node
                last = node

            if not node.next == None:
                q.put((node.next.val,node.next))

        return head
