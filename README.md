## 基于beego做后端 vue做前端的后台框架

## 部署指南

## 1.windows下的服务部署 ##

windows使用 server.bat 进行服务管理
## 2.linux下的服务部
linux下使用 宝塔 supervisor 进行服务管理 
软件商店安装 Supervisor管理器 1.1 并添加 程序到守护程序 以www身份运行
### supervisord开机自启动
参考 [Supervisor/initscripts](https://github.com/Supervisor/initscripts)

#### centos7+
#### 在安装宝塔的supervisor后，由于supervisor没有开机启动,我们设置开机启动
#### 删除已有文件
```
rm -d /lib/systemd/system/supervisord.service
vim /lib/systemd/system/supervisord.service
chmod +X /lib/systemd/system/supervisord.service
```
#填写以下内容
```
# supervisord service for systemd (CentOS 7.0+)
# by ET-CS (https://github.com/ET-CS)
[Unit]
Description=Supervisor daemon
After=rc-local.service

[Service]
Type=forking
ExecStart=/usr/bin/supervisord
ExecStop=/usr/bin/supervisorctl $OPTIONS shutdown
ExecReload=/usr/bin/supervisorctl $OPTIONS reload
KillMode=process
Restart=on-failure
RestartSec=42s

[Install]
WantedBy=multi-user.target
```
#关闭supervisor已运行服务
```shell script
supervisorctl shutdown
```

####开启supervisor服务
```shell script
systemctl  start  supervisord.service
systemctl enable supervisord.service
systemctl  restart  supervisord.service
systemctl start/restart/stop supervisord.service
```


## 3.搭配宝塔面板，实现客户管理的方便快捷 ##
搭配 nginx 更加专业，后续方便负载均衡的使用
## 关于日志 ##
开发模式(dev)下，日志实时写入文件，生产环境 prod 下会缓存 483 条数据,数据库日志单独文件
##打包命令
bee pack -exs=.go:.DS_Store:.tmp:tmp:vendor:.gitignore:mod:.sum:.log
