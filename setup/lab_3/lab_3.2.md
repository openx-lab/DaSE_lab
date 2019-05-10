### 实验二 任意进制转10进制

请写Python程序完成b-十的实数转换。
#### [参考代码]

```
def b2dec_int(x,b):
    d =0
    weight = b**(len(x)-1)
    for i in range(len(x)):
        if x[i] !='0':
            d = d+ weight*int(x[i])
        weight = weight // b
    #print (d)
    return d
​
​
def b2dec_float(y,b):
    d = 0
    for i in range(len(y)):
        if y[i]!='0':
            d += int(y[i])*(b**(-(i+1)))
    #print (d)
    return d
​
​
def b2dec_float2(y,b):
    d =0
    weight = b**(-len(y))
    for i in range(len(y)-1,-1,-1):
        if y[i]!='0':
            d = d+ weight*int(y[i])
        weight = weight*b
    return d
​
​
def b2dec(s,b):
    x,y = s.split('.')
    #print (x,y)
    x_=b2dec_int(x,b)
    y_=b2dec_float(y,b)
    print (x_+y_)
​
x = input()
43211.01234567
b = int(input())

b2dec(x,b)
```

#### [结果]

```
110.11
2
6.75
```
思考：可以参考b2dec_int函数的思路，改写b2dec_float函数么？
