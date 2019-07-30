language: [English](#Introduction), [中文](#介绍)

# **Gorminiter**

### Introduction
Auto generating gorm init file. Helping developers to build gorm structrure.
:blush:

###Install

    git clone https://github.com/ccggddmm/gorminiter.git
    go install

For Mac User

    Just copy gorminiter, and use.
    
###Usage

    ./gorminiter [-hfc] [options]

Options:

	-host	  database host, defaul: locolhost
	-port	  database port number, default: 3306
	-username database user name, default: root
	-password database password, default no password
	-bd	  database name
	-table 	  table name

Pattern:

    -c : Use config file to set parameters
    -f : Gen file on ./your_table_name.go
    -h : Show usage

without -c and -f.
   
Example 1 :
    
    ./gorminiter -port 3306 -host 127.0.0.1  -db mydb -password password -table mytable

    >>DB connnect success

     type Mytable struct { 
            Id int64 `gorm:id` 
            UserName string `gorm:user_name` 
            Status string `gorm:status` 
            UpdateTime time.Time `gorm:update_time` 
    } 

With -c, set your config in setting.json, if there is no file on current dir, gorminiter will generate template file for you.

Example 2 :
    
    ./gorminiter -c

    >>DB connnect success

     type Mytable struct { 
            Id int64 `gorm:id` 
            UserName string `gorm:user_name` 
            Status string `gorm:status` 
            UpdateTime time.Time `gorm:update_time` 
    } 



Todo:

    Add comment
    Generate dao file
    

### 介绍

用golang做服务端，使用gorm操作数据库，发现要手动撸操作数据库的接口体？数据库有很多字段？
别崩溃，gorminiter帮你搞定！
gorminiter是一款可以通过配置mysql信息，自动登陆mysql获取数据表信息并自动生成对应gorm结构体的工具。
支持命令配置，文件配置。

###安装说明

    git clone https://github.com/ccggddmm/gorminiter.git
    go install

Mac用户：

    复制可执行文件直接使用

###使用说明

    ./gorminiter [-hfc] [options]


	-host	    数据库ip，默认值: locolhost
	-port	    数据库端口号, 默认值: 3306
	-username   用户名, 默认值: root
	-password   密码 , 默认密码为空
	-bd         数据库名
	-table 	    数据表名

模式：

    -c : 使用配置文件配置参数
    -f : 生成gorm是数据库操作模版文件 存放位置./your_table_name.go
    -h : 帮助

普通模式

Example 1 :
    
    ./gorminiter -port 3306 -host 127.0.0.1  -db mydb -password password -table mytable

    >>DB connnect success

     type Mytable struct { 
            Id int64 `gorm:id` 
            UserName string `gorm:user_name` 
            Status string `gorm:status` 
            UpdateTime time.Time `gorm:update_time` 
    } 

使用配置文件

配置文件是当前目录下setting.json, 执行./gorminiter -c如果当前目录没有配置文件，工具会帮你自动生成模版，修改模版配置文件即可。

Example 2 :
    
    ./gorminiter -c

    >>DB connnect success

     type Mytable struct { 
            Id int64 `gorm:id` 
            UserName string `gorm:user_name` 
            Status string `gorm:status` 
            UpdateTime time.Time `gorm:update_time` 
    } 

