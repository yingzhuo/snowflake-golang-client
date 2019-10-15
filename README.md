# snowflake-golang-client

### 安装:

```bash
go get github.com/yingzhuo/snowflake-golang-client
```

### 使用例:

```golang
package main

import (
	"fmt"
	cli "github.com/yingzhuo/snowflake-golang-client"
)

func main() {

	client := cli.NewClient(&cli.Config{
		Host:         "localhost",
		Port:         8080,
		ResponseType: cli.Protobuf,
	})

	// 生成多个ID
	for _, v := range client.NextIds(10) {
		fmt.Println(v)
	}

	// 生成单个ID
	fmt.Println(client.NextId())
}
```
