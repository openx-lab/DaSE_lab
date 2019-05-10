### 实验十三 递归求表达式

根据实验七An的递归表达式，编程求An的精确解，输入为n，输出为An的精确解， 比如输入3，输出为3/5（小数表示亦可）

```
from fractions import Fraction
def fac(n):
    if n==0:
        return 1
    else:
        return 1/(1+fac(n-1))
 
​
print('enter n:')
n = int(input())
fac(n)
```