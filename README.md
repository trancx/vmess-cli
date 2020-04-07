# vmess-cli
Linux 下配置代理的一个小脚本，完成订阅URL等需求

## 准备工作
把项目`clone`下来然后`go build`
在项目中，设置`v2ray`目录链接到`v2ray-core`下，即这个目录下有可执行文件`v2ray`

## 设置订阅
运行`vmess-cli`敲以下指令 `subs 订阅URL`

## 常用指令说明
`update` -- 更新订阅  
`show` -- 展示所有可用的服务器  
`select idx` -- 选择代理服务器  
`start` -- 开启代理(默认全局)  
`stop` -- 终止代理  
`?`    -- 帮助  

## 注意
这个脚本只是方便修改配置文件，目前默认的监听接口为`10808`，程序没有设置修改的接口
如果有PAC的接口，需要用第三方的工具生成PAC，注意配置接口，这个很方便修改
然后将文件的路径配置在系统就可，或者在浏览器使用一些代理插件，方法很多。
