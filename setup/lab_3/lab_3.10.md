### 实验十 排序

#### [实验内容]
完成merge（L1,L2）函数：输入参数是连个从小到大排好序的证书列表L1和L2，返回合成后的从小到大排好序的大列表X。例如merge（[1,4,5],[2,7]）会返回[1,2,4,5,7]，要求比较两列表元素大小的次数不能超过len(L1)+len(L2)。

```
def merge(L1,L2):
    i = 0
    j = 0
    L = []
    while (i<len(L1)) & (j<len(L2)):
        if L1[i]<L2[j]:
            L.append(L1[i])
            i+=1
            #print ("i",i)
        else :
            L.append(L2[j])
            j+=1
            #print ("j",j)
    if i==len(L1):
        for k in range(j,len(L2)):
            L.append(L2[k])
    if j==len(L2) :
        for k in range(i,len(L1)):
            L.append(L1[k])
    print (L)
​
​
L1 = [1,4,5]
L2 = [0,2,7,9,]
merge(L1,L2)

```

#### [结果]

```
[0, 1, 2, 4, 5, 7, 9]
```