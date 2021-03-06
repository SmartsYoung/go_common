package main

import (
	"errors"
	"fmt"
	"net/http"
)

const (
	DATETIME = "2019-04-04 12:12:12"
)

/*func init() {
	fmt.Println("init A")
}*/
var result string

/*func main() {\
	hello := "hello"
	var world string = "world"
	result = hello+world
	h.HandleFunc("/",IndexHandler)
	h.ListenAndServe(":8080",nil)
}*/

func IndexHandler(w http.ResponseWriter, r *http.Request) {
	fmt.Println(DATETIME)
	w.Write([]byte(result))
}
func QueryHelloWorld(str string) (string, error) {
	if len(result) == 0 {
		return result, errors.New("result is not null")
	}
	return result + str, nil
}

type I interface {
	f()
}

type S struct {
	a int64
}

func (this *S) f() {
	this.a++
}

//func main() {
//	c := make(chan int, 10)
//	close(c)
//	c <- 1
//	start := time.Now().UnixNano() / 1000 / 1000
//	var s S
//	var i I = &s
//	for j := 0; j < 1000000000; j++ {
//		i.f()
//	}
//	end := time.Now().UnixNano() / 1000 / 1000
//	println(end - start)
//}

func findByPage(pages []int, channel chan string) {
	for i := 0; i < len(pages); i++ {
		fmt.Printf("查询第 %d 页成功\n", pages[i])
	}
	channel <- "查询完毕"
	close(channel)
}

type student struct {
	Name string
	Age  int
}

func pase_student() map[string]*student {
	m := make(map[string]*student)
	stus := make([]student, 0)

	stus = append(stus, student{Name: "sss", Age: 25})

	for _, stu := range stus {
		stu.Name = "123"
	}
	fmt.Println(stus)
	return m
}
func main() {

	students := pase_student()
	for k, v := range students {
		fmt.Printf("key=%s,value=%v \n", k, v)
	}
}
