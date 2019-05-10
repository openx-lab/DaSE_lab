### 实验十一 递归求幂

递归求13的n次方的Python程序，n作为输入。

```
def power(x,n):
    if n==0:
        return 1
    else:
        return x*power(x,n-1)
​
print('enter n:')
n = int(input())
4
power(13,n)
```

#### [结果]

```
enter n:
3
2197
```