# 目录
## golang基本语法
***
- ### Hello world!
```go
package main

import "fmt"

func main() {
	fmt.Println("hello world")
}
```
- ### 基本数据类型
```go
package main

import "fmt"

func hello() {
	fmt.Println("hello world")
}

func main() {
	//arr是切片类型，可以装载空接口类型，由于go中所有类型默认都实现了空接口，所以所有类型都可以被装下
	arr := make([]interface{}, 0)
	//放入一个1
	arr = append(arr, 1)
	//放入一个bool类型
	arr = append(arr, true)
	//放入一个函数
	arr = append(arr, hello)
	//放入一个map
	arr = append(arr, make(map[int]string, 0))
	//放入一个channel
	arr = append(arr, make(chan int, 0))
	//打印arr
	fmt.Println(arr)

	/*
		输出：
		[1 true 0xdf57a0 map[] 0xc00004a060]
	*/
}

```
- ### 条件语句和循环
```go
package main

import "fmt"

func main() {
	//使用make创建一个切片
	arr := make([]int, 0)
	//循环10次，给切片中添加数字
	for i := 0; i < 10; i++ {
		arr = append(arr, i*100)
	}
	//range循环中，i为下标，j为值，如果不需要某一个，可以使用"_"来替代
	for i, j := range arr {
		fmt.Printf("第%d个数是%d\n", i, j)
	}
	/*
		输出为：
		    	第0个数是0
				第1个数是100
				第2个数是200
				第3个数是300
				第4个数是400
				第5个数是500
				第6个数是600
				第7个数是700
				第8个数是800
				第9个数是900
	*/

    //判断arr长度是否为0，输出相应结果
	if len(arr) != 0 {
		fmt.Println("arr不为空")
	} else {
		fmt.Println("arr为空")
	}
}
```
- ### 并发
```go
//开启协程和并发控制
package main

import (
	"fmt"
	"sync"
)

//声明一个等待组
var wg sync.WaitGroup

func hello(ch chan int) error {
	fmt.Println("开始执行")
	/*
		执行的操作
		有错误返回error
	*/
	//执行完任务后取出管道中的一个值，使得管道中有空位
	<-ch
	//等待组计数减一
	wg.Done()
	//没有错误，返回nil
	return nil
}

func main() {
	ch := make(chan int, 10)
	for i := 0; i < 100; i++ {
		//等待组计数加一
		wg.Add(1)
		//并发控制，管道中最多有10个数，剩下的协程将被阻塞，等待进入管道
		ch <- 1
		//使用go关键字开启一个协程，主线程继续向后执行
		go hello(ch)
	}
	//等待所有任务执行完毕
	wg.Wait()
}
```

```go
//互斥锁、读写锁、原子操作

```
### 练习：
- #### 实现生产者消费者模型
案例：
```go

```
- #### 实现一个无锁队列
案例：
```go

```
- #### leetcode算法练习
***
# 数据库
***
## 数据库基本原理
- ### SQL基本语法
增：
```sql
-- 创建一个teach数据库
CREATE DATABASE teach;
-- 使用teach数据库
USE teach;
--创建一个test表
CREATE TABLE test(
	-- ID作为主键，自动增长
    ID INT UNSIGNED PRIMARY KEY AUTO_INCREMENT ,
	-- varchar类型的name，长度最大为50
    NAME VARCHAR(50) NOT NULL
);
-- 插入一个路人甲和路人乙
INSERT INTO test VALUES(1, "路人甲");
INSERT INTO test VALUES(2, "路人乙");
```
删：
```sql
-- 删除路人乙
DELETE FROM test WHERE NAME = "路人乙";
-- 删除table表
DROP TABLE test;
```
改：
```sql
-- 将路人甲改名为龙套丙
UPDATE test SET NAME = "龙套丙" WHERE ID = 1;
-- 给表中添加一行年龄字段
ALTER TABLE test ADD COLUMN AGE INT UNSIGNED;
```
查：
```sql
-- 查询表中所有信息
SELECT * FROM test;
```
- ### 事务和ACID
  - 事务是一组不可分割的操作的组合，这组操作要么全部成功，要么全部失败，作为一个最小单元存在。同时，事务还满足ACID特性。
  - ACID指的是原子性（Atomicity）、一致性（Consistency）、隔离性（Isolation）、持久性（Durability）。
    - 原子性：事务中的一组操作是不可分割的最小单位，就像原子一样，这组操作作为一个整体，要么全部成功，要么全部失败。
    - 一致性：事务完成前后的数据才可被外界访问，事务进行时的中间状态数据不能被访问。即数据只能从一个事务完成前正确的状态转移到另一个事务完成后正确的状态，而正在进行中的事务中的数据是非法的。
    - 隔离性：一个事务正在进行操作时对其它操作来说是不可见的。
    - 持久性：当事务完成时修改会永久保留，不会因断电等突发情况导致丢失更改。
  - Mysql采用自动提交模式，执行一个操作后会自动提交该事务。如果想显示开启事务需要使用begin开启事务，使用commit提交事务。
- ### 并发时的问题：
  - 脏读：当隔离性没有被满足时，一个操作读到另一个正在进行的事务中的数据，由于事务没有完成，所以读到了非法数据。这个数据就是脏数据。
  - 不可重复读：一个事务多次读取一个数据，同时还有其它操作正在对此数据进行操作，此时事务两次读取的数据不一样。
  - 幻读：一个事务正在读取一些数据，另一个操作对这部分数据进行增加和删除，导致事务读取到的数据不一样。与不可重复读的区别：幻读是数据数目不同，而不可重复读是数据内容不同。
  - 丢失修改：一个事务的操作被另一个事务覆盖，比如A事务更改了数据，B事务随后又更改了数据，覆盖了A事务的操作，或者B事务删除了这个数据，导致A事务的操作无效。
- ### 隔离级别
  - 未提交读：允许读取未提交数据。
  - 已提交读：只允许读取已经提交了的数据。可解决脏读问题。
  - 可重复读：只允许读取已经提交了的数据，而且在两次读取一个数据期间不允许其它事务更新该数据。可解决脏读和不可重复读问题。
  - 可串行化：事务只能串行执行，可解决脏读、不可重复读和幻读问题。
## MySQL
- ### 索引
  - B树和B+树：
  - MySQL索引：
- ### 查询优化
- ### 存储引擎
  - InnoDB：
  - MyISAM：
- ### 复制
  - 主从复制：
  - 读写分离：
- ### go操作MySQL
    - 使用的第三方库：  
go get  
go get
## Redis
- ### 概述
- ### 数据类型
- ### 数据结构
- ### 使用场景
***
# 操作系统
***
待更新
***
# 计算机网络
***
待更新
***
# go web开发
***
- ## http
  - ### http教程
  - ### get请求
  - ### post请求
  - ### get请求和post请求的区别
- ## web框架
  - ### gin框架
  - ### beego框架
***
# 缓存、分布式、集群、消息队列、中间件
***
待更新
***