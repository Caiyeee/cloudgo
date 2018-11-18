#### 简单 web 服务程序 cloudgo



1. ##### 监听端口：8080

2. ##### 模板介绍

   1. negroni
   2. mux
   3. render

3. ##### curl测试

   启动程序，`go run main.go`

   打开另一个终端窗口，输入命令：`curl -v http://localhost:8080/caiye`

   结果如下：

   ![](pics/curl.png)

   正确获取并返回输入的 name：caiye，测试成功

   

4. ##### 压力测试 ab

   输入命令：`ab -n 1000 -c 100 http://localhost:8080/caiye`

   结果如下：

   ![](pics/ab1.png)

   ![](pics/ab2.png)

   可以看到，0 失败，0 错误，总共用时 0.288 秒，平均每个请求耗时 28.824 ms.

   

   同时可以看到，运行监听程序的终端窗口，有如下信息（仅截出部分）：

   ![](pics/listen.png)

   

5. ##### 压力测试参数解释