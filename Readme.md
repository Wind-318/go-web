# 后端开发攻略(笔记)
***
## [中文文案排版指北](https://github.com/sparanoid/chinese-copywriting-guidelines/blob/master/README.zh-CN.md)
***
## [Markdown 基本语法](https://github.com/younghz/Markdown)
***
# Golang 攻略
- ## [基本数据类型](GO/基本类型.md)
- ## [条件语句和循环](GO/条件语句与循环.md)
- ## [反射与类型断言](GO/反射与类型断言.md)
- ## [文件操作](GO/文件操作.md)
- ## [并发](GO/并发.md)
  - #### 读写锁
  - #### 互斥锁
  - #### 原子操作
- ## [练习：](GO/练习.md)
  - ### [实现生产者消费者模型](GO/生产者消费者.go)
  - ### [实现一个线程安全的队列](GO/线程安全的队列.go)
  - ### [实现一个无锁队列](GO/无锁队列.go)
  - ### 实现一个线程池
- ## [GMP 调度模型](GO/GMP调度模型.md)
- ## [连接数据库](GO/database.md)
- ## Web 开发
  - ### [Go 网络编程](GO/Go%20Web开发/网络编程.md)
  - ### [正则表达式](GO/正则表达式.md)
  - ### [GraphQL](GO/Go%20Web开发/GraphQL.md)
  - ### Web 框架
    - #### Gin 框架
    - #### Beego 框架
- ### Echarts
- ### logrus
- ## [Gui 编程(walk)](GO/三方库/walk.md)
- ## [模拟键鼠操作](GO/三方库/robotgo.md)
***
# C++ 攻略
- ## STL
  - ### [vector](C++/STL/vector/vector创建和使用.md)
  - ### list
  - ### stack
  - ### queue
  - ### [set](C++/STL/Set/set.md)
  - ### map
  - ### algorithm
- ## 其它
  - ### [实现自定义排序规则](C++/其它/实现自定义排序规则.cpp)
  - ### [字符串与整数转换](C++/其它/字符串整数转换.md)
  - ### [常见问题](C++/其它/常见问题.md)
***
# Python 攻略
- ## [基本数据类型](Python/基本数据类型.md)
- ## [生成器](Python/生成器.md)
- ## [Turtle 库](Python/Turtle库.md)
- ## [文件操作](Python/文件操作.md)
- ## [并发](Python/并发.md)
- ## 三方库
  - ### [jieba](https://github.com/fxsjy/jieba)
  - ### [wordcloud](Python/第三方库/wordcloud.md)
- ## 连接数据库
- ## Web 开发
- ## 爬虫
  - ### [requests 库](Python/requests.md)
  - ### [正则表达式](Python/正则表达式.md)
- ## Web 框架
  - ### Django
- ## 爬虫框架
  - ### Scrapy
***
# Java 攻略
- TODO
***
# 数据结构与算法
- ## 基本数据结构
  - [链表](数据结构与算法/链表/链表.md)
  - [栈](数据结构与算法/栈/栈.md)
  - 队列
  - [优先队列](数据结构与算法/优先队列/优先队列.md)
  - [树](数据结构与算法/树/树.md)
  - 图
  - [哈希表](数据结构与算法/哈希表/哈希表.md)
  - [并查集](数据结构与算法/并查集/并查集.md)
  - [设计](数据结构与算法/设计/设计.md)
  - [Trie 树](数据结构与算法/Trie树/Trie树.md)
  - [Bitmap](数据结构与算法/Bitmap.md) 
- ## 算法
  - [前缀和](数据结构与算法/前缀和/前缀和.md)
  - [排序](数据结构与算法/排序/排序.md)
  - [动态规划](数据结构与算法/动态规划/动态规划.md)
  - [深度优先搜索(DFS)](数据结构与算法/深度优先搜索/深度优先搜索.md)
  - [广度优先搜索(BFS)](数据结构与算法/广度优先搜索/广度优先搜索.md)
  - [贪心算法](数据结构与算法/贪心/贪心.md)
  - [二分查找](数据结构与算法/二分查找/二分查找.md)
  - [双指针](数据结构与算法/双指针/双指针.md)
  - [回溯算法](数据结构与算法/回溯算法/回溯算法.md)
  - [位运算](数据结构与算法/位运算/位运算.md)
  - [滑动窗口](数据结构与算法/滑动窗口/滑动窗口.md)
  - [分治](数据结构与算法/分治/分治.md)
- ## [其它](数据结构与算法/其它.md)
- ## 参考
  - #### [Leetcode](https://leetcode-cn.com/problemset/all/)
  - #### [二分查找算法详解](https://mp.weixin.qq.com/s/uA2suoVykENmCQcKFMOSuQ)
***
# Linux 攻略
- TODO
***
# 设计模式
- ## [单例模式](设计模式/单例模式/单例模式.md)
- ## [简单工厂模式](设计模式/简单工厂/简单工厂.md)
- ## [工厂模式](设计模式/工厂模式/工厂模式.md)
- ## [观察者模式](设计模式/观察者模式/观察者模式.md)
- ## 参考
  - #### [设计模式 | 菜鸟教程 - RUNOOB.COM](https://www.runoob.com/design-pattern/design-pattern-tutorial.html)
***
# 数据库
- ## [数据库基本原理](数据库/数据库基本原理/数据库基本原理.md)
  - ### SQL基本语法
  - ### 事务和ACID
  - ### 并发
  - ### 封锁协议
  - ### 隔离级别
  - ### 多版本并发控制
  - ### SQL 注入
- ## [MySQL](数据库/MySQL/MySQL.md)
  - ### 数据类型
  - ### 索引
    - **B树和B+树**
    - **MySQL索引**
  - ### 查询优化
  - ### 存储引擎
    - **InnoDB**
    - **MyISAM**
  - ### 主从复制
  - ### 读写分离
  - ### 索引失效
  - ### 最左前缀原则
***
# 操作系统
- ## [概述](操作系统/概述.md)
  - ### 并发与并行
  - ### 异步
  - ### 用户态和内核态
  - ### 中断
- ## [进程与线程](操作系统/进程与线程.md)
  - ### 进程间通信
  - ### 调度
  - ### 经典 IPC 问题
    - #### 哲学家进餐问题
    - #### 读-写者问题
- ## [内存管理](操作系统/内存管理.md)
  - ### 物理内存和虚拟内存
  - ### 分页和分段机制
  - ### 页面置换算法
- ## [死锁](操作系统/死锁.md)
  - ### 死锁发生条件
  - ### 解决方法
- ## 参考
  - #### 现代操作系统（原书第 4 版）/（荷）Andrew S. Tanenbaum，（荷）Herbert Bos 著；陈向群等译. 一北京：机械工业出版社，2017.7
  - #### [计算机操作系统](https://github.com/CyC2018/CS-Notes/blob/master/notes/%E8%AE%A1%E7%AE%97%E6%9C%BA%E6%93%8D%E4%BD%9C%E7%B3%BB%E7%BB%9F%20-%20%E7%9B%AE%E5%BD%95.md)
## 
***
# 计算机网络
- ## [概述](计算机网络/概述.md)
- ## [链路层](计算机网络/链路层.md)
- ## [网络层](计算机网络/网络层.md)
- ## [传输层](计算机网络/传输层.md)
- ## [应用层](计算机网络/应用层.md)
- ## [HTTP 与 I/O](计算机网络/HTTP.md)
- ## 参考
  - #### JamesF.Kurose, KeithW.Ross, 库罗斯, 等. 计算机网络: 自顶向下方法 [M]. 机械工业出版社, 2014.
  - #### W.RichardStevens. TCP/IP 详解. 卷 1, 协议 [M]. 机械工业出版社, 2006.
  - #### [计算机网络](https://github.com/CyC2018/CS-Notes/blob/master/notes/%E8%AE%A1%E7%AE%97%E6%9C%BA%E7%BD%91%E7%BB%9C%20-%20%E7%9B%AE%E5%BD%95.md)
  - #### [面试题之计算机网络](https://www.nowcoder.com/discuss/468385?source_id=profile_create_nctrack&channel=-1)
***
# 编译原理
- TODO
***
# [分布式与集群](集群.md)
***
# 中间件
- TODO
***
# 缓存
- ## [Redis](缓存/Redis/Redis.md)
  - ### 数据类型
  - ### 语法
  - ### 持久化
  - ### 缓存问题
- ## Memcached
***
# 容器
- TODO