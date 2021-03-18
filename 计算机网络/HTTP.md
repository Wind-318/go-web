# HTTP
- ### HTTP(HyperText Transfer Protocol
  - 超文本传输协议的缩写，是一种应用广泛的网络传输协议。用户使用Http协议，通过浏览器作为客户端向web服务器发送请求。
  - HTTP是无连接的，指每次连接处理完一个请求后就会断开。
  - HTTP是无状态的，协议对于事务没有记忆，每次连接都是新连接，与上次无关。
- ### HTTP请求与响应
  - **请求报文**：发送一个http请求包含以下格式：请求行、请求头部、空行和请求数据。  
    **示例**：
    ```
    Accept-Language: zh-CN,zh;q=0.9,en;q=0.8,en-US;q=0.7
    Cookie: JSESSIONID=1F78DFDEDF9EBA5561AABDE63DD158E5; SERVERID=122; JSESSIONID=1AEB986F1EDF83C48AF8BE6BD2D60535
    Host: xxx
    Referer: xxx
    User-Agent: Mozilla/5.0 (Windows NT 10.0; Win64; x64) AppleWebKit/537.36 (KHTML, like Gecko) Chrome/89.0.4389.90 Safari/537.36 Edg/89.0.774.54
    ```
  - **响应报文**：响应报文也分为四个部分：状态行、消息报头、空行和响应正文。  
    **示例**：
    ```
    Cache-Control: no-cache
    Content-Encoding: gzip
    Content-Type: text/html;charset=utf-8
    Date: Thu, 18 Mar 2021 08:46:09 UTC
    Expires: Thu, 01 Jan 1970 00:00:00 GMT
    Pragma: No-cache
    Transfer-Encoding: chunked
    vary: accept-encoding
    x-frame-options: SAMEORIGIN
    ```
- ### URL：统一资源定位符，例如：https://cn.bing.com/
- ### HTML：一种标记语言
- ### 请求方法
  - #### Get：向服务器获取资源
  - #### Head：只返回响应头的get请求
  - #### Post：向服务器提交表单数据
  - #### Get请求和Post请求的区别
    get请求会将请求头与数据一起发送过去，post请求会先发送请求头，等待服务器响应后再发送数据。
- ### 常见HTTP状态码
  - **200**：请求成功
  - **301**：资源已经被永久转移到新位置
  - **404**：请求URL不存在
  - **500**：服务器错误
- ### HTTP与HTTPS的区别