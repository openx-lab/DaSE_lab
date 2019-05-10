### 实验七 控制结构

```
1、if语句实现百分制转等级制
def if_test(score):
    if(score>=90):
        print('Excellent')
    elif(score>=80):
        print('Very good')
    elif(score>=70):
        print('Good')
    elif(score>=60):
        print('Pass')
    else:
        print('Fail')
​
if_test(45)

2、while循环判断b是否为质数
b = int(input())
9
a = b//2
while a>1:
    if b%a ==0:
        print('b is not prime')
        break
    a = a-1
else :
    print('b is prime')
​
​

3、for语句循环遍历
l = ['python','is','a','good','language']
for i in range(len(l)):
    print(i,l[i],end=' ')
```