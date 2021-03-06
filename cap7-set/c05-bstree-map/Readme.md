## 1. 基于前面实现的二分搜索树实现集合
1. 因为我们可以采用不只一种方法实现集合，所以这里我们可以实现一个集合的接口，包含集合的基本操作方法
2. 其基本操作方法有：添加元素，删除元素，获取元素个数，非空判断，集合中是否包含某个元素
## 2. 基于链表实现集合
1. 自己实现removeElement(element)函数
## 3. 前面两种集合的性能对比
1. 通过写性能测试的方式来查看性能
2. 进行性能测试 `go test -test.bench=".*"`
3. 结果对比
    1. 二分搜索树    `30          47173083 ns/op`
    2. 链表         `2          871671950 ns/op`
4. ns/op:表示每次操作消耗的时间(纳秒)，通过相除可以得出，每次链表操作所花费的时间是二分搜索树的18.47倍
5. 为什么二分搜索树这么高效？可以简单分析一下时间复杂度
    1. 对于集合，一般我们的常规操作主要包括增、删、查
    
    1.1 链表实现集合时，分析如下
        
        1. 增：正常情况下链表添加元素时间复杂度是O(1),但由于集合元素不能重复，所以我们在添加元素之前要查找一下是否包含，也就是要遍历一遍，而这个时间复杂度是O(n)
        2. 删：删除节点的话，需要先找到我们待删除的节点(或者前面的节点)，所以，也会有一次遍历，所以它的时间复杂度也为O(n)
        3. 查：很明显这个时间复杂度是O(n)
        
    1.2 所以如果一个链表中有n个元素，那么其时间复杂度应该是O(n^2)    
        
        1. 增：增加元素的话，我们每次都会向下判断一下节点，所以，在满二叉树的情况下，复杂度应该是二叉树的调度O(h)
        2. 删：具体判断逻辑与上面类似
        3. 查：具体判断逻辑与上面类似
    1.3 这里需要搞清楚的是h和n的对比，我们假设链表和二分搜索树都有n个元素，判断一下它们的操作次数，其中二分搜索树考虑满节点的情况，具体如下
    
    | n的数量 | 链表复杂度(n) | 二分搜索树复杂度(h) | 对比    |
    | ------- | ---------- | ---------------- | ------- |
    | 1       | 1          | 1                | 1倍     |
    | 2       | 2          | 2                | 1倍     |
    | 4       | 4          | 3                | 1倍     |
    | 16      | 16         | 5                | 2倍     |
    | 128     | 128        | 8                | 16倍    |
    | 1024    | 1024       | 10               | 102倍   |
    | 1000000 | 1000000    | 20               | 50000倍 |
    上面列出了n和h的关系,可以看到，数量级越大，差距也越大，实际上，通过上面的计算，结论如下
    
    1. 如果是链表：平均复杂度应该是O(n)
    2. 如果是二分搜索树：平均复杂度应该是O(log2n)
    当然，如果是最坏情况的二分搜索树，也可能退化成链表 
## 4. 有序集合与无序集合
1. 有序集合中的元素具有顺序性：基于搜索树实现
2. 无序集合中的元素具有无序性：基于哈希表实现    
## 5. 基于链表实现映射
1. 这里与实现集合的方式类似，区别在于现在它有两个元素需要处理
        
## 6 基于二分搜索树实现映射
1. 实现方式与集合实现类似
## 7 有序映射和无序映射
1. 有序映射中的元素具有顺序性：基于搜索树实现
2. 无序映射中的元素具有无序性：基于哈希表实现  