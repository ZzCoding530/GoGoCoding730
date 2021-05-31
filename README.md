# GoGoCoding730
使用Go开启新一轮的刷题咯！

一是复习做过的题，二是让Go的语法更熟悉　

三颗🌟表示需要再刷

 **🥳为了730li！**

---

### 👨🏻‍💻链表

| No.     |        Title         | Difficulty | Note                                         | Again？ |
| :------ | :------------------: | :--------: | :------------------------------------------- | :-----: |
| 206     |       反转链表       |    Easy    | 没难度                                       |         |
| 25      |   K个一组反转链表    |    Hard    | 其实没难度                                   |         |
| 21      |   合并两个有序链表   |    Easy    | 也没啥难度                                   |         |
| 160     |       相交链表       |    Easy    | 脑筋急转弯                                   |         |
| 141     |       环形链表       |    Easy    | 没难度                                       |         |
| 142     |     环形链表 II      |   Medium   | 搞清楚第二次相遇的位置                       |         |
| 92      |     反转链表 II      |   Medium   | 值得多做几次                                 |    🌟    |
| Offer22 |    倒数第K个节点     |    Easy    | 没难度                                       |         |
| 19      |  删除倒数第k个节点   |   Medium   | 没啥难度，有点小细节                         |         |
| 234     |       回文链表       |    Easy    | 有思路就没难度                               |         |
| 143     |       重排链表       |   Medium   | 还是中点、反转各种组合起来                   |         |
| 83      |   删除链表重复元素   |    Easy    | 没难度                                       |         |
| 82      |  删除链表重复元素II  |   Medium   | 82的升级版，值得一做                         |    🌟    |
| 24      | 两两交换链表中的节点 |   Medium   | 每次写递归都像个傻子，重写                   |   🌟🌟    |
| 138     | 复制带随机指针的链表 |   Medium   | 看官方题解第三种解答，好题                   |   🌟🌟    |
| 328     |       奇偶链表       |   Medium   | 挺好玩，多做几遍                             |   🌟🌟    |
| 23      |   合并K个排序链表    |    Hard    | Java直接用优先队列很好做，<br />Go的我还没写 |  🌟🌟🌟🌟🌟  |

##### No24 两两反转链表，其实就相当于是25K个一组反转链表的特例情况！

---

### 👨🏻‍💻树

| No.     |        Title         | Difficulty | Note                                                         | Again？ |
| :------ | :------------------: | :--------: | ------------------------------------------------------------ | :-----: |
| 102     |    二叉树层序遍历    |   Medium   | Go没有各种集合真的不太方便= =                                |   🌟🌟🌟   |
| 103     |  二叉树Z字层序遍历   |   Medium   | ⚠️注意切片的“引用传递”其实是值传递！                          |   🌟🌟🌟   |
| 236     |  二叉树最近公共祖先  |   Medium   | 反正树基本上就是递归，<br />就看想不想得出思路               |   🌟🌟    |
| 94      |   二叉树的中序遍历   |    Easy    | 树的各种遍历，都要迭代，递归没意思                           |    🌟    |
| 199     |     二叉树右视图     |   Medium   | 就是层序遍历呗                                               |    🌟    |
| 144     |    二叉树前序遍历    |    Easy    | 其实切片用习惯了也还好                                       |    🌟    |
| 104     |    二叉树最大深度    |    Easy    | 没难度                                                       |         |
| 110     |    判断平衡二叉树    |    Easy    | 也算是没难度吧，都是递归                                     |         |
| 124     |    二叉树最长路径    |    Hard    | 提交了一个评论区大神的Go解法<br />一出手就知道是老程序员了   |   🌟🌟🌟   |
| 113     |      路径总和II      |   Medium   | 多做几遍，熟练Go的语法特性                                   |   🌟🌟🌟   |
| 105     | 前序和中序重建二叉树 |   Medium   | 树就是递归，每次的边界画图就好了                             |   🌟🌟🌟   |
| 543     |      二叉树直径      |    Easy    | 测试用例会改变全局变量初始值，好坑                           |         |
| 226     |      翻转二叉树      |    Easy    | 一定掌握迭代法，递归太没难度了                               |   🌟🌟🌟   |
| 101     |      对称二叉树      |    Easy    | 没意思                                                       |         |
| 112     |       路径总和       |    Easy    | = =                                                          |         |
| 129     | 根节点到叶子节点的和 |   Medium   | 第一遍居然没做出来，哭泣                                     |   🌟🌟🌟   |
| 98      |    判断搜索二叉树    |   Medium   | 递归还行，<br />层序遍历判断是否单调递增的<br />迭代法也需要掌握 |  🌟🌟🌟🌟   |
| 958     |    判断完全二叉树    |   Medium   | 层序遍历，当遍历到nil以后，<br />就不该再碰到一个不是nil的了 |   🌟🌟🌟   |
| offer54 |    BST第K大的节点    |    Easy    | 中序遍历，可以略优化                                         |         |
| 145     |    二叉树后序遍历    |    Easy    | 遍历我从来都是迭代，递归有手就行                             |         |
| 662     |    二叉树最大宽度    |   Medium   | 带着序号进行层序遍历，值得再做                               |   🌟🌟🌟   |
| 230     |       BST第K小       |   Medium   | 就是中序遍历，没啥                                           |         |
| 297     | 二叉树序列化反序列化 |    Hard    | 思路不难，麻烦，要熟练Go的库                                 |   🌟🌟🌟   |
| offer27 |      二叉树镜像      |    Easy    | 就是翻转二叉树，起了个花里胡哨的名                           |         |
| 114     |     二叉树转链表     |   Medium   | 还算有点难度，再做吧                                         |   🌟🌟🌟   |



5/27忙着做毕设 没写题，难过