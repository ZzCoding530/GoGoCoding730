package leetcode

/*
	LRU算法 需要实现
type LRUCache struct							结构体
func Constructor(capacity int) LRUCache			构造函数
func (this *LRUCache) Get(key int) int			Get方法
func (this *LRUCache) Put(key int, value int)	Put方法

*/

/*
	【1】	先建一个双向链表
*/
type DoubleListNode struct {
	key   int
	value int
	prev  *DoubleListNode
	next  *DoubleListNode
}

/*
	【2】	再来个缓存本身的类（结构体
*/
type LRUCache struct {
	size     int                     //目前的size
	capacity int                     //最大的容量，size<=capacity
	cache    map[int]*DoubleListNode //缓存本身的结构

	head *DoubleListNode //头节点
	tail *DoubleListNode //尾节点

}

/*
	【3】	双向链表的初始化函数 给key和value
*/
func constructorDoubleListNode(key, value int) *DoubleListNode {
	return &DoubleListNode{
		key:   key,
		value: value,
	}
}

/*
	【4】	构造函数 初始化一个容量为capacity的LRU缓存
*/
func Constructor(capacity int) LRUCache {
	lruCache := LRUCache{ //新建一个LRUCache
		capacity: capacity,                        //容量填进去
		cache:    map[int]*DoubleListNode{},       //新建一个缓存给进去
		head:     constructorDoubleListNode(0, 0), //初始化一个头节点进去
		tail:     constructorDoubleListNode(0, 0), //初始化一个尾节点进去
	}

	//刚才新建完了要把首尾指针连起来
	lruCache.head.next = lruCache.tail
	lruCache.tail.prev = lruCache.head

	return lruCache //把建好的返回出去
}

/*
	【6】	开始实现题目要求的两个方法
*/
func (this *LRUCache) Get(key int) int { //按照key找值
	//先判断这个key存在不存在啊，不存在返回-1
	if _, ok := this.cache[key]; !ok {
		return -1
	}

	//ok 如果存在呢，那就取出来呗
	node := this.cache[key]
	this.moveToHead(node) //记得取出来以后，把这个node放到最前面，因为要求把刚用过的放到最前面

	return node.value //别忘了要干啥，返回需要的值
}

func (this *LRUCache) Put(key int, value int) { //放进去一个node，这个有点复杂，因为涉及到容量
	//先判断要放进去的key是不是已经存在了
	if _, ok := this.cache[key]; !ok {
		//如果不存在那么
		node := constructorDoubleListNode(key, value) //先按照key，value初始化一个节点，这个节点就是要新放进去的
		this.cache[key] = node                        //放进去
		this.addToHead(node)                          //别忘了操作过的节点都需要放到最前面
		this.size++                                   //多了一个新的节点，size加1

		//放进去以后，关键来了，判断容量是不是超了，超了得删最久没用的（最尾部的
		if this.size > this.capacity {
			toBeRemoved := this.removeTail()
			delete(this.cache, toBeRemoved.key)
			//上面delete是golang中map自带的函数，因为把节点从链表中删除以后，map中存的也得删除
			//The delete built-in function deletes the element with the specified key
			//(m[key]) from the map. If m is nil or there is no such element, delete
			//is a no-op.

			this.size-- //删完size减1

		}
	} else { //这里else是 想要放进去的key已经存在，那就直接改里面的就行
		node := this.cache[key]
		node.value = value
		this.moveToHead(node) //必须要把node先取出来，因为要把这个刚操作过的node放到最前面

	}

}

/*
	【5】
		为了实现Get和Put函数，需要自己写几个API
		主要是addToHead（把节点添加到头部）	和	removeNode（移除某个节点）
		moveToHead	（把节点移动到头部）	和	removeTail（移除尾部节点）都可以通过调用前两个来实现
*/
func (this *LRUCache) addToHead(node *DoubleListNode) { //节点添加到首尾，head和tail是LRU自带的两个哨兵节点
	node.prev = this.head
	node.next = this.head.next
	this.head.next.prev = node
	this.head.next = node
}

func (this *LRUCache) removeNode(node *DoubleListNode) { //删除双向链表某个节点
	node.prev.next = node.next
	node.next.prev = node.prev
}

func (this *LRUCache) moveToHead(node *DoubleListNode) { //把某个节点移动至首尾
	this.removeNode(node) //实际上是先删除
	this.addToHead(node)  //再添加到首尾
}

func (this *LRUCache) removeTail() *DoubleListNode { //移除最后尾节点
	node := this.tail.prev //首先拿到尾部节点
	this.removeNode(node)  //然后移除
	return node            //把移除的节点返回出去
}
