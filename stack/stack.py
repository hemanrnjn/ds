from sys import maxsize

def createStack():
    stack = []
    return stack

def isEmpty(stack):
    if len(stack) > 0:
        return False
    return True

def push(stack, item):
    stack.append(item)
    return stack

def pop(stack):
    if isEmpty(stack):
        # print(maxsize)
        return str(-maxsize-1)
    return stack.pop()

def peek(stack):
    if isEmpty(stack):
        return str(-maxsize-1)
    return stack[len(stack) - 1]

stack = createStack() 
push(stack, str(10)) 
push(stack, str(20)) 
push(stack, str(30)) 
print(pop(stack) + " popped from stack")
print(stack)
