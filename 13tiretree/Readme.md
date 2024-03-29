#### Trie树（字典树）
- 介绍
```
trie树 是一种多叉树，专门处理字符串匹配的数据结构，用来解决在一组字符串集合中快速查找某个字符串的问题；
存储方式：子节点用一个数组保存，key=字符的哈希值（可以使用ASCII），value=这个节点的地址
type TireTree struct {
    Data Byte
    child [26]*TireTree
}

假设我们的字符串中只有从 a 到 z 这 26 个小写字母，我们在数组中下标为 0 的位置，存储指向子节点 a 的指针，
下标为 1 的位置存储指向子节点 b 的指针，以此类推，下标为 25 的位置，存储的是指向的子节点 z 的指针。
如果某个字符的子节点不存在，我们就在对应的下标的位置存储 null。
```

- 常规算法
```
1.插入（只考虑英文字符）
    对一个要插入tire树的字符串进行扫描，对字符进行 ASCII - 'a'确定数组index，
    然后将所对应的value设置[26]*TireTree,插入结束后对最后一个字符的isEnd 设置true
2.查找
    
3.删除
```

#### AC自动机（多匹配）

[🔗](https://www.cnblogs.com/sclbgw7/p/9260756.html)

- 介绍
```
案例：对用户输入进行敏感字检查，可以将敏感字提前生成一个tire树，然后对用户的输入进行匹配，
    但是这种方案，每次都只能对一个敏感词匹配；
AC 自动机就是为了解决这种问题，当一次匹配失败后，AC自动机会从当前字符开始再次匹配，可以一次性与多个敏感词做比较
```
- 实现原理
```
主串：用户输入
模式串：敏感字组成的tiretree,

当主串在模式串上匹配的时候，如果遇到匹配不到的字符，就从tiretree上拿到已经匹配到的字符串{m}，
然后在这个字符串上找到可以在tiretree上匹配到的m最长前缀，
然后拿到这个最长前缀在tiretree可以匹配到的字符串的末尾开始匹配；
```
![匹配失败后的实效](https://github.com/lqhandsome/Alg/blob/master/13tiretree/failnext.png)

- 失效函数
```
失效函数，就是当一个节点匹配不到后，下次应该匹配的位置；
类似于KMP的next函数，

当构建的模式串在匹配的时候，如果一个节点（p）的孩子节点为空，
或者孩子节点无法匹配的情况，就从root->p 的这串字符中找到一个最长的后缀，
在tiretree上匹配，如果有的话，那p的failnext 就指向这个匹配到的串的结尾；

--- 如何构建失败指针（就这个节点而言，孩子无法匹配，）

 假设p 的失败指针是q pc是p的子节点，qc是q的子节点
1.root 的失败指针为 nil
2.如果 p的子节点 pc无法匹配，就查找p的失效指针，看是否能与pc重合，
3.如果能，则pc的失效指针就是qc ;
4.如果不能令p=p.fail,重复2,3操作
```
