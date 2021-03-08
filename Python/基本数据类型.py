# 创建一个List, 并且放入一个字符串和一个整数
arr = ["asd", 123]
# 创建一个元组(tuple)，放入字符串
temp = ("test")
# 创建一个字典，姓名为路人甲，年龄为114514
dict = {"name":"先♂辈", "age":114514}
# 创建一个set，set中不会有重复的元素
set = {1,2,3,3}
# 给arr追加元素true
arr.append(True)
# 给arr追加元素temp、dict和set
arr.append(temp)
arr.append(dict)
arr.append(set)
# 打印arr中的元素和arr的长度
print(arr, len(arr))



# 表示单行注释

'''
表示
多行
注释
'''


list1 = [1,1]
# 复制
list2 = list1.copy()
print(list2)
# 追加
list1.append("asd")
print(list1)
# 统计数量
print(list1.count(1))
# 搜索指定区间，返回第一个出现值的下标
print(list1.index(1, 0, 3))
# 插入
list1.insert(2, "qwe")
# 删除
list1.pop(1)
# 删除
list1.remove("asd")
# 翻转
list1.reverse()
print(list1)
# 清空
list1.clear()
print(list1)
# 排序(list中的元素本身要能比较)
list2.sort()