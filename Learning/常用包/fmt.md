# fmt

## Sprintf

fmt.Sprintf(格式化样式, 参数列表…)

示例：

   sql := fmt.Sprintf(
			"insert into test_tab(id, date, name) values ('%d', '%s', '%s')",
			21,
			"2021-08-08",
			"test21",
		  )

    fmt.Println(sql)

执行结果：

    insert into test_tab(id, date, name) values ('21', '2021-08-08', 'test21')