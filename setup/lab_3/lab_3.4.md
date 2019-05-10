### 实验四 用逻辑做加法

[实验内容]
设计一个全加器，全加器有三个输入：A和B是两个加数，Ci是从下一位获得的进位，和两个输出：Sum，给上一位的进位C。逻辑可以表示为

```
C = A & B | A & Ci | B & Ci，
Sum = A & B & Ci | A & !B & !Ci | !A & B & !Ci | !A & !B Ci。
```

```
def FA(a, b, c):
    Carry = (a and b) or (b and c) or (a and c)
    Sum = (a and b and c) or (a and (not b) and (not c)) or ((not a) and b and (not c)) or ((not a) and (not b) and c)
    return Carry,Sum 
    
def add(x, y):
    while len(x) < len(y) : x = [False] + x
    while len(y) < len(x) : y = [False] + y
    L = [];Carry = False
    for i in range(len(x)-1,-1,-1):
        Carry,Sum = FA(x[i],y[i],Carry)
        L = [Sum] + L
    return (Carry, L)
​
print (add([True,True],[True,True,True]))
```
### [输出结果]

```
(True, [False, True, False])
```

那么，如何实现乘法？

```
def multiplier (x,y):
    S = []
    for i in range(len(y)-1, -1, -1):
        if y[i] == True:
            C,S = add(S,x)
            if C == True:
                S = [C] + S
        x = x + [False]
    return (S)
​
​
x = [True,True]
y = [True,False,True]
print (multiplier(x, y))
```

#### [输出结果]

```
[True, True, True, True]
```