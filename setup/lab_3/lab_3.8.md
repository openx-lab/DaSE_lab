### 实验八 函数调用和参数传递

对于下面三个函数，你认为应该输出什么？

```
def func(L1):
    L1.append(1)
​
​
L=[2]
func(L)
print(L)

def func2(L1):
    x = L1 + [1]
    print (x,L1)
 
​
L = [2]
func2(L)
print (L)

def recursive(L):
    if L ==[]: return 
    L = L[0:len(L)-1]
    print ("L=",L)
    recursive(L)
    print ("L:",L)
    return 
​
​
recursive([1,2,3])
```