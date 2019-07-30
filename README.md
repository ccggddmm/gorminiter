language: [English](#Introduction), [中文](#介绍)

# **Gorminiter**

### Introduction
Auto generating gorm init file. Helping developeer to build gorm structrure.
:blush:

Usage: 

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

     type Table struct { 
            Id int64 `gorm:id` 
            UserName string `gorm:user_name` 
            Status string `gorm:status` 
            UpdateTime time.Time `gorm:update_time` 
    } 

With -c, set your config in setting.json, if there is no file on current dir, gorminiter will generate template file for you.

Example 2 :
    
    ./gorminiter -c

    >>DB connnect success

     type Table struct { 
            Id int64 `gorm:id` 
            UserName string `gorm:user_name` 
            Status string `gorm:status` 
            UpdateTime time.Time `gorm:update_time` 
    } 



Todo:

    Add comment
    Generate dao file
    

### 介绍
