distributed-task
================

> ` 用go语言编写的轻量级分布式任务调度框架. A light framework of task execution in distributed system Written with go `
> 

---

- 项目描述
-
> distributed-task 是一个开源的轻量级分布式任务调度框架，
> >
> > <p>当前1.0 目标是实现一个可以部署在本地和远端的
> > 运行框架，用户可以在框架中自定义执行的任务Task和任务数据集Data，并且用户可以自定义任务的分配方式和
> > 数据分配方式。用户可以在本地和远端分布式并行执行Task，并将执行完成后的结果集返回给本地，实现并发处理
> >一些任务的需求。（此处的本地可以看作中心节点）<p>
> ><br>
> ><p>2.0 实现代码的动态部署，即本地端可以任意重新定义Task，而不用重新部署远端（有点难，希望有大大可以加入）
> >.<p><br>
> > <p> 3.0 实现节点间的自动弹性调度，优化调度性能（主要从数据传输，网络性能，CPU等方面考虑）<p><br>
> > <P> 4.0 用一致性哈希实现计算过程的缓存和存储<P>
>  

	
- 设计思想
- 
> 借鉴MapReduce的设计思想，来分布式的并行计算Task，利用go routine 的特性实现多个机器的多核并行，
> 和高并发量；数据集目前是独立无关的任务数据

- 1.0模块和设计
-

  1. 网络模块

  	<p>负责网络传输，监听，心跳包<p>
  2. 命令模块
  
  	<p>负责根据收到的message 调用执行相关的模块<p>
  3. 调度器
  
  	<p>负责分发任务(节点级别)<p>
  4. 执行单元
  
  	<p>类似于Map Reduce 里面的 Map 将任务数据集并发在节点执行</p>
  5. 结果回收
  	
  	<P>类似于Map Reduce 里面的 Reduce 将结果集归于返回给本地端[MapReduce](http://static.googleusercontent.com/media/research.google.com/zh-CN//archive/mapreduce-osdi04.pdf)</p>
  6. 配置管理
  	
  	<p>配置节点ip，cpu核数，并发力度，负载度<p>
  	
  	![image](http://mytutu.qiniudn.com/go-task.jpg)
  
  
  ----
 join the team
 ===
 ----
 	给njutree dot at gmail dot com 发一封邮件加入我们的项目
 	Please send me a mail at njutree dot whitworth at gmail dot com.
  
  












