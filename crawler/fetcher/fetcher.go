package fetcher

import (
	"bufio"
	"fmt"
	"golang.org/x/net/html/charset"
	"golang.org/x/text/encoding"
	"golang.org/x/text/encoding/unicode"
	"golang.org/x/text/transform"
	"io"
	"log"
	"net/http"
	"time"
)

// 该函数作用是从指定的 URL 下载内容，并将这些内容转换为 UTF-8 编码的文本
// 函数输入是一个字符串类型的 url 参数，输出是一个字节切片和一个错误值。字节切片用于存储从 URL 获取的内容
var rateLimiter = time.Tick(10 * time.Microsecond)

func Fetch(url string) ([]byte, error) {
	//<-rateLimiter // 防止http获取太快了
	resp, err := http.Get(url)
	if err != nil {
		return nil, err
	}
	// 使用 defer 关键字来确保在函数返回前关闭响应体（resp.Body）。这是处理 HTTP 响应时的常见做法，以避免资源泄露。
	defer resp.Body.Close()
	if resp.StatusCode != http.StatusOK {
		return nil,
			fmt.Errorf("wrong status code: %d",
				resp.StatusCode)
	}
	// 使用 bufio.NewReader 创建一个新的缓冲读取器，以便高效地读取响应体
	bodyReader := bufio.NewReader(resp.Body)
	// 调用 determineEncoding 函数来确定响应体的编码。
	e := determineEncoding(bodyReader)
	// 创建一个新的 transform.Reader，它将响应体从原始编码转换为 UTF-8 编码
	utf8Reader := transform.NewReader(resp.Body, e.NewDecoder())
	// 读取转换后的响应体，并将其作为字节切片返回。如果读取成功，返回字节切片和 nil 错误；如果失败，返回 nil 和错误对象。
	return io.ReadAll(utf8Reader)

}

// 这个函数尝试读取响应体的前 1024 个字节来猜测其编码
// 输入是一个 *bufio.Reader 参数，并返回一个 encoding.Encoding 对象，表示猜测到的编码。
func determineEncoding(
	r *bufio.Reader) encoding.Encoding {
	// 尝试读取缓冲读取器中的前 1024 个字节，而不会从缓冲区中移除这些字节。这用于猜测内容的编码。
	bytes, err := r.Peek(1024)
	// 如果尝试读取字节时发生错误，记录错误并假设编码为 UTF-8
	if err != nil {
		log.Printf("Fetcher error: %v", err)
		return unicode.UTF8
	}
	// 使用 charset.DetermineEncoding 函数猜测字节切片的编码
	e, _, _ := charset.DetermineEncoding(
		bytes, "")
	return e
}
