# GoGoCoding730
使用Go开启新一轮的刷题咯！

一是复习做过的题，二是让Go的语法更熟悉　

三颗🌟表示需要再刷

 **🥳为了730li！**

---

### 一些很重要的基础

- 各种排序的Go版本联系
- Go的优先队列
- LRU多写几遍
- 左右边界的二分法查找
- 树的各种遍历迭代版熟练一些
- 单调栈要熟练

---

### 👨🏻‍💻链表

|   No.   |        Title         | Difficulty | Note                                         | Again？ |
| :-----: | :------------------: | :--------: | :------------------------------------------- | :-----: |
|   206   |       反转链表       |    Easy    | 没难度                                       |         |
|   25    |   K个一组反转链表    |    Hard    | 其实没难度                                   |         |
|   21    |   合并两个有序链表   |    Easy    | 也没啥难度                                   |         |
|   160   |       相交链表       |    Easy    | 脑筋急转弯                                   |         |
|   141   |       环形链表       |    Easy    | 没难度                                       |         |
|   142   |     环形链表 II      |   Medium   | 搞清楚第二次相遇的位置                       |         |
|   92    |     反转链表 II      |   Medium   | 值得多做几次                                 |    🌟    |
| Offer22 |    倒数第K个节点     |    Easy    | 没难度                                       |         |
|   19    |  删除倒数第k个节点   |   Medium   | 没啥难度，有点小细节                         |         |
|   234   |       回文链表       |    Easy    | 有思路就没难度                               |         |
|   143   |       重排链表       |   Medium   | 还是中点、反转各种组合起来                   |         |
|   83    |   删除链表重复元素   |    Easy    | 没难度                                       |         |
|   82    |  删除链表重复元素II  |   Medium   | 82的升级版，值得一做                         |    🌟    |
|   24    | 两两交换链表中的节点 |   Medium   | 每次写递归都像个傻子，重写                   |   🌟🌟    |
|   138   | 复制带随机指针的链表 |   Medium   | 看官方题解第三种解答，好题                   |   🌟🌟    |
|   328   |       奇偶链表       |   Medium   | 挺好玩，多做几遍                             |   🌟🌟    |
|   23    |   合并K个排序链表    |    Hard    | Java直接用优先队列很好做，<br />Go的我还没写 |  🌟🌟🌟🌟🌟  |

##### No24 两两反转链表，其实就相当于是25K个一组反转链表的特例情况！

---

### 👨🏻‍💻树

|   No.   |        Title         | Difficulty | Note                                                         | Again？ |
| :-----: | :------------------: | :--------: | ------------------------------------------------------------ | :-----: |
|   102   |    二叉树层序遍历    |   Medium   | Go没有各种集合真的不太方便= =                                |   🌟🌟🌟   |
|   103   |  二叉树Z字层序遍历   |   Medium   | ⚠️注意切片的“引用传递”其实是值传递！                          |   🌟🌟🌟   |
|   236   |  二叉树最近公共祖先  |   Medium   | 反正树基本上就是递归，<br />就看想不想得出思路               |   🌟🌟    |
|   94    |   二叉树的中序遍历   |    Easy    | 树的各种遍历，都要迭代，递归没意思                           |    🌟    |
|   199   |     二叉树右视图     |   Medium   | 就是层序遍历呗                                               |    🌟    |
|   144   |    二叉树前序遍历    |    Easy    | 其实切片用习惯了也还好                                       |    🌟    |
|   104   |    二叉树最大深度    |    Easy    | 没难度                                                       |         |
|   110   |    判断平衡二叉树    |    Easy    | 也算是没难度吧，都是递归                                     |         |
|   124   |    二叉树最长路径    |    Hard    | 提交了一个评论区大神的Go解法<br />一出手就知道是老程序员了   |   🌟🌟🌟   |
|   113   |      路径总和II      |   Medium   | 多做几遍，熟练Go的语法特性                                   |   🌟🌟🌟   |
|   105   | 前序和中序重建二叉树 |   Medium   | 树就是递归，每次的边界画图就好了                             |   🌟🌟🌟   |
|   543   |      二叉树直径      |    Easy    | 测试用例会改变全局变量初始值，好坑                           |         |
|   226   |      翻转二叉树      |    Easy    | 一定掌握迭代法，递归太没难度了                               |   🌟🌟🌟   |
|   101   |      对称二叉树      |    Easy    | 没意思                                                       |         |
|   112   |       路径总和       |    Easy    | = =                                                          |         |
|   129   | 根节点到叶子节点的和 |   Medium   | 第一遍居然没做出来，哭泣                                     |   🌟🌟🌟   |
|   98    |    判断搜索二叉树    |   Medium   | 递归还行，<br />层序遍历判断是否单调递增的<br />迭代法也需要掌握 |  🌟🌟🌟🌟   |
|   958   |    判断完全二叉树    |   Medium   | 层序遍历，当遍历到nil以后，<br />就不该再碰到一个不是nil的了 |   🌟🌟🌟   |
| offer54 |    BST第K大的节点    |    Easy    | 要想到BST中序遍历是有序的                                    |         |
|   145   |    二叉树后序遍历    |    Easy    | 遍历我从来都是迭代，递归有手就行                             |         |
|   662   |    二叉树最大宽度    |   Medium   | 带着序号进行层序遍历，值得再做                               |   🌟🌟🌟   |
|   230   |       BST第K小       |   Medium   | 就是中序遍历，没啥                                           |         |
|   297   | 二叉树序列化反序列化 |    Hard    | 思路不难，麻烦，要熟练Go的库                                 |   🌟🌟🌟   |
| offer27 |      二叉树镜像      |    Easy    | 就是翻转二叉树，起了个花里胡哨的名                           |         |
|   114   |     二叉树转链表     |   Medium   | 还算有点难度，再做吧                                         |   🌟🌟🌟   |
|   100   |       相同的树       |    Easy    | 递归就vans了                                                 |         |

**5月的最后一天结束了树的专题，剩下的不太好分类，一起做吧**

---

### 👨🏻‍💻各种题

|   No.   |                   Title                    | Difficulty | Note                                                         | Again？ |
| :-----: | :----------------------------------------: | :--------: | :----------------------------------------------------------- | :-----: |
|   912   |                  排序数组                  |   Medium   | 快速排序练习                                                 |         |
|   215   |           数组中的第K个最大元素            |   Medium   | 快排做，以后写个堆做的                                       |   🌟🌟🌟   |
|    3    |            无重复字符的最长子串            |   Medium   | 巨高频！录视频讲解吧                                         |  🌟🌟🌟🌟🌟  |
|   146   |                    LRU                     |   Medium   | 必会必会必会的思<br />不算难但是写起来需要很细心<br />录讲解 |  🌟🌟🌟🌟🌟  |
|    1    |                  两数之和                  |    Easy    | 梦开始的地方                                                 |         |
|   53    |                 最大子序和                 |    Easy    | 动态规划还挺好做的                                           |         |
|   15    |                  三数之和                  |   Medium   | 我觉得还挺有难度的，考验基本功                               |   🌟🌟🌟   |
|   121   |             买卖股票的最佳时机             |    Easy    | 经典                                                         |    🌟    |
|   415   |                 字符串相加                 |    Easy    | 没啥说的                                                     |         |
|   88    |              合并两个有序数组              |    Easy    | 从后往前确定                                                 |         |
|   20    |                  有效括号                  |    Easy    | 栈呗，搭配map                                                |         |
|   704   |                 二分法查找                 |    Easy    | =。=记得复习一下各种带重复值的<br />寻找左侧边界和右侧边界的二分查找 |   🌟🌟🌟   |
|   46    |                   全排列                   |   Medium   | 回溯法，很喜欢的一道题，多做几遍                             |   🌟🌟🌟   |
|   232   |                用栈实现队列                |    Easy    | 两个栈，很基础                                               |         |
|   54    |                  螺旋矩形                  |   Medium   | 理清思路，多做几遍，有意思的题                               |   🌟🌟🌟   |
|   200   |                  岛屿数量                  |   Medium   | 很好的系列题，必须多做                                       |  🌟🌟🌟🌟🌟  |
|    5    |                最长回文子串                |   Medium   | 全 文 背 诵                                                  |  🌟🌟🌟🌟🌟  |
|   33    |              搜索旋转排序数组              |   Medium   | 很考验二分法的基本功，<br />建议复习各种边界的二分法         |   🌟🌟🌟   |
|   42    |                   接雨水                   |    Hard    | 单调栈会写就秒杀                                             |   🌟🌟🌟   |
|   300   |               最长递增子序列               |   Medium   | 能想到dp就能做，想不到就麻烦了                               |   🌟🌟🌟   |
|   70    |                   爬楼梯                   |    Easy    | 显然是dp，这种求个数，长度，<br />单独数字的，基本上是dp     |         |
|    2    |                  两数相加                  |   Medium   | 咋说呢，不难也不简单，<br />需要思考又不会想不出来           |    🌟    |
|   151   |            翻转字符串里面的单词            |   Medium   | 我 是 一 个 无 情 的 API 调 用 选 手                         |         |
|   69    |                 x的平方根                  |    Easy    | 0到x二分吧，单独处理0，1以后100%<br />从0开始+1那种不行太慢了 |         |
|    8    |              atoi字符串转整数              |   Medium   | 双百   多帅哦，全 文 背 诵                                   |   🌟🌟🌟   |
|   155   |                   最小栈                   |    Easy    | 没啥，俩栈就行                                               |         |
|   169   |                  多数元素                  |    Easy    | 可以排序做，但了解一下摩尔投票法吧                           |         |
|   136   |              只出现一次的数字              |    Easy    | 位运算                                                       |         |
| offer10 |                斐波那契数列                |    Easy    | DP                                                           |         |
| offer21 |                调整数组奇偶                |    Easy    | 双指针 没啥                                                  |         |
|   283   |                   移动0                    |    Easy    | 双指针 跟上面一样的思路                                      |         |
|   14    |                最长公共前缀                |    Easy    | 没啥                                                         |         |
|  1143   |               最长公共子序列               |   Medium   | 一般求长度的都用DP，从头计算长度的<br />用一位数组，需要首尾划分计算长度的<br />用二维数组，注意Go二维数组的初始化 |   🌟🌟🌟   |
|   470   |          用Rand7( )实现Rand10( )           |   Medium   | 构造大的数然后取余啥的各种操作吧，<br />rand+rand的概率不变，rand*rand好像就变了，<br />构造的时候注意用常数乘然后再加 |  🌟🌟🌟🌟   |
|   31    |                 下一个排列                 |   Medium   | 难度适中我觉得                                               |   🌟🌟🌟   |
|   56    |                  合并区间                  |   Medium   | 值得一做，难度适中                                           |   🌟🌟🌟   |
|   148   |                  链表排序                  |   Medium   | 非常经典的一道题，递归和迭代当然选迭代                       |  🌟🌟🌟🌟   |
|   718   |               最长公共子数组               |   Medium   | 数组，字符串，求长度的都是dp                                 |   🌟🌟🌟   |
|   48    |                  旋转图像                  |   Medium   | 翻转两次即可，发现这点就没啥了                               |    🌟    |
|   39    |                  数组总和                  |   Medium   | 标准的回溯                                                   |   🌟🌟🌟   |
|   78    |                    子集                    |   Medium   | 回溯回溯回溯！都弄明白一次性                                 |   🌟🌟🌟   |
|   34    | 在排序数组中查找元素的第一个和最后一个位置 |   Medium   | 二分二分二分一次搞清楚好吧                                   |   🌟🌟🌟   |
|   108   |                  打家劫舍                  |   Medium   | dp 我自己做的奥力给                                          |   🌟🌟🌟   |
|   64    |                 最小路径和                 |    Hard    | 吓唬人的，Easy DP                                            |   🌟🌟🌟   |
|   239   |               滑动窗口最大值               |    Hard    | 蛮好的一道题，录一下吧                                       |   🌟🌟🌟   |
|   240   |               搜索二维矩阵II               |   Medium   | 类似二分法                                                   |    🌟    |
|   322   |                  零钱兑换                  |   Medium   | dpdpdpdp                                                     |    🌟    |
|   518   |                 零钱兑换II                 |   Medium   | dpdpdpdppdpdp                                                |    🌟    |
|   224   |                 基本计算器                 |    Hard    |                                                              |         |
|   59    |                 螺旋矩阵II                 |   Medium   | 细心点不是很难                                               |   🌟🌟🌟   |



---



