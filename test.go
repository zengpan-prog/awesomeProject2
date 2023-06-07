package main

import (
	"database/sql"
	"fmt"
)

/*
import (
	"flag"
	"fmt"
)
*/
/*//数据生产者
func producer(header string, channel chan<- string) {
	// 无限循环, 不停地生产数据
	for {
		// 将随机数和字符串格式化为字符串发送给通道
		channel <- fmt.Sprintf("%s: %v", header, rand.Int31())
		// 等待1秒
		time.Sleep(time.Second)
	}
}

// 数据消费者
func customer(channel <-chan string) {
	// 不停地获取数据
	for {
		// 从通道中取出数据, 此处会阻塞直到信道中返回数据
		message := <-channel
		// 打印数据
		fmt.Println(message)
	}
}*/
/*
var mode = flag.String("mode", "", "process mode")
func main() {
/*	// 创建一个字符串类型的通道
	channel := make(chan string)
	// 创建producer()函数的并发goroutine
	go producer("cat", channel)
	go producer("dog", channel)
	// 数据消费函数
	customer(channel)*/
/*
	var (
		a int
	)
	var aa int = 10
	fmt.Println(aa)
	fmt.Println(a)
	s, ss := 1, "string"
	fmt.Println(s)
	fmt.Println(ss)

	var attack = 40
	var defence = 20
	var damageRate float32 = 0.17
	var damage = float32(attack-defence) * damageRate
	fmt.Println(damage)

	var a1 int = 100
	var b1 int = 200
	b1, a1 = a1, b1
	fmt.Println(a1, b1)

	var cat int = 1
	var str string = "banana"
	t := &cat
	fmt.Printf("t value: %d \n",*t)
	fmt.Printf("t address: %p \n",t)
	fmt.Printf("cat address: %p \nstr addtess: %p", &cat, &str)

	// 解析命令行参数
	flag.Parse()
	// 输出命令行参数
	fmt.Println(*mode)

	str1 := new(string)
	*str1 = "Go语言教程"
	fmt.Printf("str1 type: %T \n",str)
	fmt.Printf("str1 address: %p \n",&str)
	fmt.Printf("str1 value: %s \n",*str1)
	fmt.Println("------------")

	var b [3]int             // 定义三个整数的数组
	fmt.Println(b[0])        // 打印第一个元素
	fmt.Println(b[len(b)-1]) // 打印最后一个元素
	// 打印索引和元素
	for i, v := range b {
		fmt.Printf("%d %d\n", i, v)
	}
	fmt.Println("------------")
	for _, v := range b {
		fmt.Printf("%d\n", v)
	}
	fmt.Println("------------")
	for v := range b {
		fmt.Printf("%d\n", v)
	}
	fmt.Println("------------")
	q := [...]int{1, 2, 3}
	fmt.Printf("%T\n", q) // "[3]int"

	fmt.Println("------------")
	var highRiseBuilding [30]int
	for i := 0; i < 30; i++ {
		highRiseBuilding[i] = i + 1
	}
	// 区间
	fmt.Println(highRiseBuilding[10:15])
	// 中间到尾部的所有元素
	fmt.Println(highRiseBuilding[20:])
	// 开头到中间指定位置的所有元素
	fmt.Println(highRiseBuilding[:2])
	// 所有元素输出
	fmt.Println(highRiseBuilding[:])

	fmt.Println("-----------------------")
	var i int
	for {
		if i > 10 {
			break
		}
		i++
	}
}
*/

/*
func main() {
	//创建一个新文件，写入内容 5句 hello,Go语言
	filePath := "E:/Go/goes.txt"
	file, err := os.OpenFile(filePath, os.O_WRONLY | os.O_APPEND, 0666)
	if err != nil{
		fmt.Println("open file err",err)
	}
	//及时关闭file句柄
	defer file.Close()

	//写入文件时，使用带缓存的 *Writer
	write := bufio.NewWriter(file)
	for i := 0; i < 20; i++ {
		_, _ = write.WriteString("hello,Go语言1 \n")
	}
	//Flush将缓存的文件真正写入到文件中
	_ = write.Flush()
}
*/

import (
	"database/sql"
	"fmt"
	_ "github.com/lib/pq"
)

func main() {
	// 连接数据库
	db, err := sql.Open("postgres",
		"host=192.168.0.201 port=5432 user=postgres password=zengpan8690341. dbname=postgres sslmode=disable")
	if err != nil {
		panic(err)
	}
	defer db.Close()

	// 执行查询
	rows, err := db.Query("SELECT * FROM userdeatail")
	if err != nil {
		panic(err)
	}
	defer rows.Close()

	// 遍历结果
	for rows.Next() {
		var id int
		var name string
		err = rows.Scan(&id, &name)
		if err != nil {
			panic(err)
		}
		fmt.Println(id, name)
	}

	// 检查查询错误
	err = rows.Err()
	if err != nil {
		panic(err)
	}
}
