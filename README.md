## snowflake-golang-client

安装 (mod):

```bash
go get github.com/yingzhuo/snowflake-golang-client
```

使用:

```golang
package main

import (
	"fmt"
	cli "github.com/yingzhuo/snowflake_golang_client"
)

func main() {

	client := cli.NewClient(&cli.Config{
		Host:         "localhost",
		Port:         18080,
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
