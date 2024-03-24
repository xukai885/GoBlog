## 服务启动配置
#### 数据库准备
```shell
mysql:5.7
redis:6.0.6
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


