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
  - #### https是超文本传输安全协议，使用http通信，但是利用SSL加密数据包达到安全传输的目的。
  - #### 加密
    - 对称加密：加密和解密使用同一个秘钥
    - 非对称加密：拥有两个秘钥，公钥和私钥。如果使用公钥加密必须使用相应的私钥才能解密；如果使用私钥加密，必须使用相应的公钥才能解密。常见的非对称加密算法为**RSA算法**。
  - #### HTTPS工作流程
    - 客户端发起https请求
    - 客户端验证服务端证书
    - 服务端向客户端发送证书和公钥
    - 客户端发送用公钥加密后的数据
    - 服务端用私钥解密客户端的数据
    - 服务端发送私钥加密后的数据
    - 客户端用公钥解密服务端的数据
  - #### HTTP数据未加密，安全性差，HTTPS经过SSL加密，安全性好。
  - #### HTTP响应速度快，HTTP通过三次握手建立连接，需要交换三个数据包，HTTPS还有SSL握手的九个数据包，一共十二个。
  - #### HTTP使用80端口，HTTPS使用443端口。