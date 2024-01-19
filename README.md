

## example


## 前言

## 项目文档

（地址一）[doc/chinese](doc/chinese/) 即该项目的doc目录下

（地址二）go-zero官方微信公众号（微服务从代码到k8s部署应有尽有系列）

【注】go-zero的相关资料、知识点在这都能找到

![go-zero-public-qcode](./doc/chinese/images/1/go-zero-public-qcode.jpeg)



## 项目简介

整个项目使用了go-zero开发的微服务，基本包含了go-zero以及相关go-zero作者开发的一些中间件，所用到的技术栈基本是go-zero项目组的自研组件，基本是go-zero全家桶了

项目目录结构如下：


- app：所有业务代码包含api、rpc以及mq（消息队列、延迟队列、定时任务）
- common：通用组件 error、middleware、interceptor、tool、ctxdata等
- data：该项目包含该目录依赖所有中间件(mysql、es、redis、grafana等)产生的数据，此目录下的所有内容应该在git忽略文件中，不需要提交。
- deploy：

    - filebeat: docker部署filebeat配置
    - go-stash：go-stash配置
    - nginx: nginx网关配置
    - prometheus ： prometheus配置
    - script：
        - gencode：生成api、rpc，以及创建kafka语句，复制粘贴使用
        - mysql：生成model的sh工具
    - goctl: 该项目goctl的template，goctl生成自定义代码模版，template用法可参考go-zero文档，复制到家目录下.goctl即可
- doc : 该项目系列文档


## 开发模式

本项目使用的是微服务开发，api （http） + rpc（grpc） ， api充当聚合服务，复杂、涉及到其他业务调用的统一写在rpc中，如果一些不会被其他服务依赖使用的简单业务，可以直接写在api的logic中


## 感谢

go-zero 微服务: https://github.com/zeromicro/go-zero


