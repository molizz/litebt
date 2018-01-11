### 下载器

进行下载的微服务

- 超时（当超过n个小时无下载数据，应该取消下载，释放资源）

### 功能

端口8882

```
下载文件

POST /download

参数:
url: 下载地址(支持磁力，http/s)
notify: 下载完成后通知(或者每10%通知一下？还要考虑下)

返回:
{
    "status":"ok",
    "task_id": "下载任务的id"
}
```

```
下载状态

POST /status

参数:
task_id: "下载任务的id"

返回:
{
    "status":"ok",
    "progress": 50,    //进度
    "size": 123123,    //原始文件总大小(byte)
    "downloaded_size": 123123, //已下载的文件总大小(byte)
}
```

```
取消下载

DELETE /cancel

参数:
task_id: "下载任务的id"

返回:
{
    "status":"ok"
}
```