### 实验一 求解根号2的三种不同方法

#### 1.循环迭代，逐步逼近

#### 开平方1

```
def Square_root_1():
    c = 2
    i = 0
    g = 0
    for j in range(0, c+1):
        if (j * j > c and g == 0):
            g = j -1
    while(abs(g * g - c) > 0.0001):
        g += 0.00001
        i = i+1
        print ("%d:g = %.5f" % (i,g))
​
​
Square_root_1()
```

执行结果
​![pic1](http://kfcoding-static.oss-cn-hangzhou.aliyuncs.com/gitcourse-DaSE_lab/pic/2.1.png)

#### 2.二分查找法

#### 开平方2  二分法

```
def Square_root_2():
    i = 0
    c = 2
    m_max = c
    m_min = 0
    g = (m_min + m_max)/2
    while (abs(g * g - c)> 0.00000000001):
        if (g*g<c):
            m_min = g
        else :
            m_max = g
        g = (m_min + m_max)/2
        i = i + 1
        print ("%d:g = %.13f" % (i,g))
​
Square_root_2()
```

执行结果
![pic2](http://kfcoding-static.oss-cn-hangzhou.aliyuncs.com/gitcourse-DaSE_lab/pic/2.2.png)
3.牛顿法
参考文献：
https://zh.wikipedia.org/wiki/%E7%89%9B%E9%A1%BF%E6%B3%95
​​![pic2](http://kfcoding-static.oss-cn-hangzhou.aliyuncs.com/gitcourse-DaSE_lab/pic/2.3.png)

#### 牛顿法

```
def Square_root_3():
    c = 2
    g = c/2
    i = 0
    while (abs(g*g-c)>0.00000000001):
        g = (g + c/g)/2
        i = i+1
        print ("%d:%.13f"%(i,g))
​
Square_root_3()
```

执行结果
​![pic2](http://kfcoding-static.oss-cn-hangzhou.aliyuncs.com/gitcourse-DaSE_lab/pic/2.4.png)
