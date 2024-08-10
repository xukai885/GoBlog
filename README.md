## 服务启动配置
#### 准备
```shell
## 版本
Go 1.20.4 
Node.js v20.9.0
Mysql 5.7
Redis 6.0.6
```
#### 初始化数据库
```shell
运行 dao/mysql/create_user.sql执行数据库初始化
```
#### 配置账户密钥
私钥生成指令(存放在服务运行目录)
```shell
openssl genrsa -out rsa_private_key.pem 1024
```
公钥: 根据私钥生成(存放在服务运行目录)
```shell
openssl rsa -in rsa_private_key.pem -pubout -out rsa_public_key.pem
```
#### 配置服务文件
```shell
config-dev.yaml #服务会根据APP_ENV指定变量读取相应的配置文件config-?.yaml;
```
#### 容器启动
```shell
docker run -itd -p9092:9092 --name=goblog -v /boke_web/uploads/:/uploads -v/boke_web/rsa_private_key.pem:/rsa_private_key.pem -v /boke_web/rsa_public_key.pem:/rsa_public_key.pem -v /boke_web/config-prod.yaml:/config-prod.yaml -eAPP_ENV=prod repository/image:tag
```
#### 前端
[GoBlog-Web](https://github.com/xukai885/GoBlog-Web)

#### 站点
[徐同学的博客](https://blog.xwnlearn.cn)