# 数据类型
- ### STRING
  - Key - Value 型数据
- HASH
  - 以键值对存储的无序列表
- LIST
  - 链表型结构，可以从两端插入数据
- SET
  - 无序集合，会自动去除重复数据
- ZSET
  - 有序集合，相比 SET 存储的数据是有序的
# 语法
- ## STRING
| 命令 | 作用 |
| :--: | :--: |
| SET key value | 设定指定 key 对应的 value |
| GET key | 获得指定 key 对应的值 |
| GETRANGE key start end | 获得字符串从 start 到 end 区间子字符串 |
| MGET key1 key2 | 获得多个 key 对应的值 |
| MSET key1 value1 key2 value2 | 设置多个 key 和 value |
| SETEX key seconds value | 设置 key 的值为 value，并将 key 的过期时间设为 seconds |
| SETNX key value | 只有在 key 不存在时设置 key 的值为 value |
| STRLEN key | 获得 key 对应值（字符串）的长度 |
| APPEND key value | 如果 key 已存在，将 value 追加到原来的 value 上 |
  - 示例：
  > set name 张三
  OK
  > get name
  张三
  > del name
  1
  > get name
  null
- ## LIST
| 命令 | 作用 |
| :--: | :--: |
| BLPOP key1 key2 timeout | 移除并获取列表的第一个元素，如果没有则等待超时或发现可移除元素 |
| BRPOP key1 key2 timeout | 移除并获取列表的最后一个元素，如果没有则等待超时或发现可移除元素 |
| LINDEX key index | 获得 index 下标处的元素 |
| LLEN key | 获得列表长度 |
| LPOP key | 从左侧移除元素 |
| RPOP key | 从右侧移除元素 |
| LPUSH key value1 value2 | 从左侧插入多个元素 |
| RPUSH key value1 value2 | 从右侧插入多个元素 |
| LRANGE key start end | 获得从 start 到 end 范围的元素，如果范围是 0 -1 则是全部元素 |
| LSET key index value | 通过索引设置元素的值 |
| LTRIM key start end | 只保留从 start 到 end 范围的值，其它部分移除 |
  - 示例
  > rpush TestKey name1
  1
  > rpush TestKey name2
  2
  > lrange TestKey 0 -1
  name1
  name2
  > lpush TestKey name3
  3
  > lrange TestKey 0 -1
  name3
  name1
  name2
  > rpop TestKey
  name2
  > lrange TestKey 0 -1
  name3
  name1
  > lpop TestKey
  name3
  > lrange TestKey 0 -1
  name1
  > llen TestKey
  1
  > Lindex TestKey 0
  name1
- ## HASH
| 命令 | 作用 |
| :--: | :--: |
|  |  |
|  |  |
|  |  |
|  |  |
|  |  |
|  |  |
|  |  |
|  |  |
|  |  |
|  |  |
|  |  |
|  |  |
  - 示例
  > hset test name1 张三
  1
  > hmset test name2 李四 name3 王五 name4 114514 name5 老八
  OK
  > hgetall test
  name1
  张三
  name2
  李四
  name3
  王五
  name4
  114514
  name5
  老八
  > hdel test name1
  1
  > hget test name1
  null
  > hget test name2
  李四
  > hkeys test
  name2
  name3
  name4
  name5
  > hlen test
  4
  > hvals test
  李四
  王五
  114514
  老八
- ## SET
| 命令 | 作用 |
| :--: | :--: |
|  |  |
|  |  |
|  |  |
|  |  |
|  |  |
|  |  |
|  |  |
|  |  |
|  |  |
|  |  |
|  |  |
|  |  |
  - 示例
- ## ZSET
| 命令 | 作用 |
| :--: | :--: |
|  |  |
|  |  |
|  |  |
|  |  |
|  |  |
|  |  |
|  |  |
|  |  |
|  |  |
|  |  |
|  |  |
|  |  |
  - 示例
# 持久化
- ## RDB
- ## AOF

# 缓存问题
- ## 缓存雪崩
  - 同一时间正好有大批数据同时失效，传来的请求直接打在数据库上，造成数据库承担过大压力。
  - **解决方法：**
    - 过期时间进行随机化调整，避免大量 Key 同时过期
    - 访问量过大时限流，避免同时请求过多
- ## 缓存穿透
  - 攻击者刻意大量去请求不存在的数据，造成缓存像是不存在一样，所有请求直接打在数据库上，造成数据库压力过大。
  - **解决方法：**
    - 存储无效 Key 设为NULL，下次查询时直接返回
    - 使用布隆过滤器
- ## 缓存击穿
  - 大量请求访问一个热点 Key 时，Key 突然失效，大量请求直接打在数据库上，造成数据库压力过大。
  - **解决方法：**
    - 延长热点 Key 过期时间或者不设置过期时间