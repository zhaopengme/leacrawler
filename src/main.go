package main

import (
	"fmt"
	"github.com/zhaopengme/leacrawler"
	"time"
	// "os"
)

func main() {
	start := time.Now()
	fmt.Println("start...")

	lea := leacrawler.NewCrawler()

	url := "http://www.keenthemes.com/preview/metronic/theme/admin_1/"
	path := "/Users/zhaopeng/Desktop/adasd";
	lea.Fetch(url, path)

	fmt.Printf("time cost %v\n", time.Now().Sub(start))
}
