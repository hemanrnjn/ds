class Node:
    def __init__(self, data=None):
        self.data = data
        self.next = None

class LinkedList:
    def __init__(self):
        self.head = None

    def printList(self):
        x = self.head
        while(x is not None):
            print(x.data)
            x = x.next

    def insertAtBegin(self, val):
        x = Node(val)
        x.next = self.head
        self.head = x

    def insertAtEnd(self, val):
        x = Node(val)
        head = self.head
        while(head.next is not None):
            head = head.next
        head.next = x

    def findLength(self):
        head = self.head
        count = 0
        while(head is not None):
            head = head.next
            count += 1
        return count

    def insertAtIndex(self, val, index):
        count = 0
        head = self.head
        while(head is not None):
            if(index == count):
                if(index == 0):
                    self.insertAtBegin(val)
                    break
                else:
                    x = Node(val)
                    x.next = head.next
                    head.next = x
                    break
            count += 1
            head = head.next
        if(index >= count):
            print("Index doesn't exist")

    def deleteFromIndex(self, val, index):
        head = self.head
        


lList1 = LinkedList()
lList1.head = Node(1)

node2 = Node(2)
node3 = Node(3)

lList1.head.next = node2
node2.next = node3

# x = lList1.head
# while(x != None):
#     print(x.data)
#     x = x.next
# lList1.printList()
# lList1.insertAtBegin(0)
# lList1.printList()
# lList1.insertAtEnd(4)
# lList1.printList()
# print(lList1.findLength())
lList1.insertAtIndex(0, 3)
lList1.printList()