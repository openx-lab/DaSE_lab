### 实验三 10进制转2进制（包含小数部分）

完成十进制到二进制的包含小数的转换。输入时一个带小数点的十进制数，输出是带小数点的二进制数。

```
def dec2bin_int(x):
    if x <2 :
        return [x]
    r = x%2
    return (dec2bin_int(x//2) + [r])
​
def dec2bin_float(x):
    x -= int(x)
    bins = []
    while x:
        x *= 2
        bins.append(1 if x>=1. else 0)
        x -= int(x)
    return bins
    
def dec2bin(num):
    x = int (num)
    y = num - int(num)
    #print (x,y)
    x_=dec2bin_int(x)
    y_=dec2bin_float(y)
    s = ''
    for i in x_:
        s=s+str(i)
    s+='.'
    for i in y_:
        s+=str(i)
    print (s)
​
​
x = float(input())
dec2bin(x)

```

#### [结果]

```
10.75
1010.11
```