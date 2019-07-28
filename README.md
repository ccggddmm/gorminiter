# **Gorminiter**

### Introduction
Auto generating gorm init file. Helping developeer to build gorm structrure.

Usage: 

    ./gorminiter [-h] [options]


Parameters:

	-host	  database host, defaul: locolhost
	-port	  database port number, default: 3306
	-username database user name, default: root
	-password database password, default no password
	-bdname	  database name
	-table 	  table name

Example:
    
    ./gorminiter -port 3306 -host 127.0.0.1  -dbname mydb -password password -table mytable

    >>2019/07/26 20:37:23 DB connnect success

     type Table struct { 
            Id int64 `gorm:id` 
            UserName string `gorm:user_name` 
            Status string `gorm:status` 
            UpdateTime time.Time `gorm:update_time` 
    } 

Todo:

    Add primary key
    Add comment
    Generate dao file