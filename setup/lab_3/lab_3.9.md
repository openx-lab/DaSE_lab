### 实验九 最大公约数

![](http://kfcoding-static.oss-cn-hangzhou.aliyuncs.com/gitcourse-DaSE_lab/pic/3.1.png)



```
def GCD(x,y):
    if x>y:
        a = x; b=y
    else :
        a = y; b = x
    if a%b ==0:
        return (b)
    return (GCD(a%b,b))
​
​
x = int(input())
625
y = int(input())
75
print ('x和y的最大公约数为：',GCD(x,y))
​
```