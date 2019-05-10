### 实验二 使用蒙特卡洛法求Pi的值

参考文献：https://zh.wikipedia.org/wiki/%E8%92%99%E5%9C%B0%E5%8D%A1%E7%BE%85%E6%96%B9%E6%B3%95 

在一个边长为1的正方形内一均匀概率随机投点，该点落在此正方形的内切1/4圆中的概率即为内切圆与正方形的面积比值，所以Pi=落入圆的点数/所有点数*4。
​![pic2](http://kfcoding-static.oss-cn-hangzhou.aliyuncs.com/gitcourse-DaSE_lab/pic/2.5.png)
参考代码
​

#### 蒙特卡洛法求Pi

```
import random
def Pi(times):
    sum = 0
    for i in range(times):
        x = random.random()
        y = random.random()
        d2 = x*x + y*y
        if d2 <= 1 : 
            sum+=1
    return (sum/times*4)
​
​
times = 100000000
x = Pi(times)
print ("Pi = %.8f"%(x))
```