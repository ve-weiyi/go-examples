# ankerwork2b-api


## 目录

目录约定
```shell

「
 |-- cache 一些数据的缓存目录可以删除
 |-- conf 配置文件目录
 |-- docs 存放文档
 |-- resource 资源文件
 |-- script 脚本文件
 |-- src 核心目录所有代码都存放在这里面
    「
     -- common 常用工具代码与业务关系不大
     -- dto http 或者 rpc 请求参数或返回参数定义
     -- infra 常用的基础业务代码，与业务可能相关
     -- model 模型实例，与数据库对应的类型
     -- repo 资源库接口，操作 model 类型存储数据的
     -- serve 系统服务，比如http,rpc,定时器服务
     -- service 业务服务逻辑层
 |-- tools 工具，如果想写一些工具请把代码放在这里面
 |-- main.go   ankerwork2b-api 的入口


```


## 启动命令

