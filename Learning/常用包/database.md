# golang操作mysql

Go本身不提供具体数据库驱动，只提供驱动接口和管理。各个数据库驱动需要第三方实现，并且注册到Go中的驱动管理中。

Mysql库`https://github.com/go-sql-driver/mysql`

代码中需要注册mysql数据库驱动，通过引入空白导入mysql包来完成，只执行mysql包的初始化代码（代码位于`%GOPATH%/github.com/go-sql-driver/mysql/driver.go`）

    import (
        "database/sql"
        _ "github.com/go-sql-driver/mysql"
    )

## Query

查询多行

## Query

## Scan

把数据库取出的字段值赋值给指定的数据结构

    var version string

    err2 := db.QueryRow("SELECT VERSION()").Scan(&version)