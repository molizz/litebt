## Searcher

磁力服务组件之一的磁力搜索服务

### 主要功能

- 搜索接口
- 添加索引接口

#### 搜索接口

```
POST /search

key: hello
max: 20    (每次搜索最大100条)
page: 0    (翻页)
```

#### 添加索引

```
POST /add_index

title: 我是中国人
index: n+1
```