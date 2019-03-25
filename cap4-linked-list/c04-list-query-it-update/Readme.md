## c01实现链表的基本结构
1. 链表的组成
    1. 实际数据
    2. 指向下一个节点的指针(引用)
2. 基本实现
    1. 实现链表的基本结构
    2. 实现链表的构造函数
    3. 实现链表的格式化输出    
## c02实现链表节点插入
1. 确定链表基本结构：一个链表之中，至少要有一个头节点head
2. 确定链表大小，用size表示
3. 功能实现
    1. 从链表头部插入数据：通过给定元素生成一个新的node，让该node的next指向head,然后让该node成为新的head
    2. 从链表指定位置插入数据
    3. 从链表尾部插入数据
## c03为链表添加一个虚拟头节点
1. 在前面的处理逻辑中，每次调用向指定位置插入元素的insert函数时，都需要判断一下是否是头节点，因为头节点不存在prev节点。
2. 为了让insert函数统一处理，向链表中添加一个虚拟头节点，不计入链表长度，不存入数据，但它的next指向头节点  
## c04链表的查询、遍历和修改  
1. 查询操作
    1. 实现查找指定位置的元素
    2. 实现查找头节点
    3. 实现查找尾节点
    4. 查找链表中是否包含了某个元素
2. 修改操作
    1. 实现修改指定位置的节点
3. 格式化输出    