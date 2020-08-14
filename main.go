package main

import (
	"errors"
	"fmt"
	"golang.org/x/sync/errgroup"
	"strings"
	"sync"
	"time"
)

func main() {
	//chanTest()
	//rangeChan()
	// waitGroupTest()
	//errorGroupTest()
	sliceToString()
}

func sliceToString(){
	m:=[]int64{
		111,222,333,444,555,666,
	}
	// 将slice fid 转换为 1,2,3,4 这样的字符串
	fidStr := "'"+strings.Replace(strings.Trim(fmt.Sprint(m), "[]"), " ", "','", -1)+"'"
	fmt.Println(fidStr)
}

// 测试channel
func chanTest() {
	type Icc struct {
		Name string `json:"name"`
		Age  int    `json:"age"`
	}
	ch := make(chan *Icc, 5)
	defer close(ch)
	// 此处一定要注意 ！！！ defer 后边直接跟 close(ch) 会运行到 close(ch) 直接运行的！
	// defer close(ch)

	wg := &sync.WaitGroup{}
	for i := 0; i < 500; i++ {
		ch <- &Icc{
			Name: "小明",
			Age:  i,
		}
		wg.Add(1)
		go func(ch chan *Icc) {
			defer wg.Done()
			time.Sleep(500 * time.Microsecond)
			<-ch
		}(ch)
		fmt.Println("我在添加数据", i)
	}
	wg.Wait()
	//for {
	//	if icc, ok := <-ch; ok {
	//		fmt.Println("获取到了一个结果：", icc.Name, icc.Age)
	//	} else {
	//		break
	//	}
	//}
	fmt.Println("hello")
}

// 对于固定长度（缓冲区）的channel进行 range遍历
func rangeChan() {
	type Icc struct {
		Name string `json:"name"`
		Age  int    `json:"age"`
	}
	ch := make(chan *Icc, 5)
	wg := &sync.WaitGroup{}
	for i := 0; i < 5; i++ {
		wg.Add(1)
		go func(tempI int) {
			defer wg.Done()
			ch <- &Icc{
				Name: "小明",
				Age:  tempI,
			}
		}(i)
	}
	wg.Wait()
	// 可以直接close 允许读取一个关闭的channel 但是不允许往channel里放值
	close(ch)
	for v := range ch {
		fmt.Println("获取到了一个结果：", v.Name, v.Age)
	}
}

func waitGroupTest() {
	wg := &sync.WaitGroup{}
	wg.Add(1)
	var ret int32
	for i := 0; i <= 10; i++ {
		wg.Add(1)
		go func() {
			defer wg.Done()
			ret += 1
		}()
	}
	wg.Wait()
	fmt.Println(ret)
}

func errorGroupTest() {
	errGro := &errgroup.Group{}
	errGro.Go(func() error {
		return errors.New("1111")
	})
	errGro.Go(func() error {
		return nil
	})
	err := errGro.Wait()
	if err != nil {
		fmt.Println("批量执行，中间有一个出错了，", "err", err.Error())
	}
}
