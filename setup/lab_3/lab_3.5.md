### 实验五 数据类型表示

```
数值类型的表示
1、产生10-20的随机浮点数
import random 
f = random.uniform(10,20)
print(f)

2、产生10-20的随机整数
import random
i = random.randint(10,20)
print(i)

3、布尔型例子
b = 100<101
print(b)

4、字符串
print("book's price")

print('book's price')

SyntaxError: invalid syntax
python无法判别book后面的单引号是字符串的结尾，还是字符串的符号，需要转义字符‘\’进行转义
print('book\'s price')  

5、列表
l = [1,1.3,'2','China',['I','am','another','list']]  
l[0]

列表元素可以包括整数型、浮点型、字符串、还可以是列表
l[0]#取出第一个元素
l[-1]

取出最后一个元素
l[4][0]

取出列表里列表的元素
l[1:]

取出第2个到最后一个元素
l1 = [1,1.3]
l2 = ['2','China',['I','am','another','list']]
l = l1+l2
print(l)

列表加法
6、字典
d = {'Michael':95,'Bob':89,'Tracy':92}
d['Bob']

```