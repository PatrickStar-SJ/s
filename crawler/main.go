package main

import (
	"fmt"
	"io"
	"net/http"
)

func main() {
	resp, err := http.Get(
		"http://localhost:8080/mock/www.zhenai.com/zhenghun")
	if err != nil {
		panic(err)
	}

	//defer resp.Body.Close() 的作用就是无论函数如何返回，都确保 HTTP 响应体会被关闭
	//
	//defer 是 Go 语言提供的一种用于处理成对的操作的机制，如打开和关闭、连接和断开连接、加锁和释放锁。
	//通过 defer 语句，我们可以在函数执行完毕后（返回之前）执行一段代码，无论函数因为什么原因返回，都会执行这个 defer 语句。
	//
	//resp.Body.Close()  用来关闭 HTTP 响应体的，这是一个必要的步骤，否则可能会导致内存泄漏。
	//使用 http.Get 发送请求时，它会返回一个响应（resp）,在读取完数据后，你需要关闭这个流，以释放系统资源。

	defer resp.Body.Close()

	if resp.StatusCode != http.StatusOK {
		fmt.Println("Error: status code", resp.StatusCode)
		return
	}

	all, err := io.ReadAll(resp.Body)
	fmt.Printf("Received body: %s\n", all)
	if err != nil {
		panic(err)
	}

}
