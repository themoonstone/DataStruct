## 1. 制作属于自己的数组
1. 设置最多能存放的容量 capacity
2. 设置数组元素的大小 size
3. 初始化数组
4. 设置查询信息
    1. 获取数组元素个数
    2. 获取数组容量
    3. 检查数组是否为空
## 2. 向数组中添加元素
1. 分析：
    1. 添加元素分为三种，尾部追加，首部添加，中间插入
2. 思考
    1. 在添加之前需要判断当前容量是否还能继续添加
## 3. 数组中查询元素与修改元素
1. 实现数组格式化输出，元素之间以逗号分隔
2. 实现获取指定位置的元素
3. 实现更新指定位置的元素
## 4. 包含，搜索和删除
1. 实现包含功能
2. 实现搜索功能
    1. 查找指定元素存在的索引
3. 实现删除功能
    1. 删除指定位置元素
    2. 删除第一个元素
    3. 删除最后一个元素
4. 实现格式化输出
## 5. 使用泛型
1. 泛型:让数据结构可以放置"任何"数据类型
2. 在golang中，可以直接使用interface{}来替代
## 6. 动态数组实现
1. 弄清楚动态数组原理：
    1. 在追加的时候，如果当前的size已经达到了数组的容积，则重新为其分配内存，容积为原有的2倍
    2. 在删除元素的时候，如果当前的size已经小于了数组容量的1/2，则重新分配内存空间，大小为原来容积的1/2
## 7. 时间复杂度分析
1. 时间复杂度:简单来说，其实时间复杂度就是基本操作的执行次数
2. 数组时间复杂度分析
    1. 添加: O(n)
    2. 删除: O(n)
    3. 修改:已知索引是O(1),未知索引是O(n)
    4. 查找:已知索引是O(1),未知索引是O(n)
3. 均摊复杂度与复杂度震荡
    1. 复杂度震荡：就是复杂度不稳定的一种情况，以动态数组为例，在添加和删除元素的时候可能会触发resize操作，那么有可能出现以下情况：
        1. 元素已满，添加一个元素，触发resize，扩容为2倍
        2. 然后马上删除一个元素，其容积又变成扩容之后的1/2，所以执行缩容操作
        3. 不断重复上面两步
    2. 通过上面的步骤，我们发现每一次操作的时间复杂度都是O(n),而对于数组操作来说O(n)是其最坏的情况，这就是我们所说的复杂度震荡