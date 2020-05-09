# SerialTest
一种基于REST的串口自动化测试方案
## 目的
* 面向嵌入式设备串口调试或自动化测试流程
* 提供极其简单的RESTful交互接口
* 串口服务端与测试客户端分离 可通过局域网（有条件的可以公网但不建议）远程测试
* 避免XShell等工具对技术栈的强绑定 数据同步BUG WAIT阻塞BUG
## API
* GET: {SERVER}/serialtest/v1/register?url={LISTEN}
* GET: {CLIENT}/{LISTEN}?content=teststring&timestamp=1585710432855541900
* GET: {SERVER}/serialtest/v1/writer?content=teststring
## 编译
## 使用
1. ./scripts/build.sh 
2. ./bin/SerialTestServer -s com13 -b 38400 -p :8220 -l server.log
## Tips
* 服务端启动将生成一个access token 客户端HTTP请求需携带access_token头部