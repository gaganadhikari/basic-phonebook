package main

import(
	"fmt"
	"os"
	"path"
	"strconv"
	"math/rand"
	"time"
)

type Entry struct{
	Name string
	Surname string
	Tel string
}

var data = []Entry{}
var MIN = 0
var MAX = 26

func search(key string) *Entry{
	for i,v :=range data{
		if v.Tel == key{
			return &data[i]
		}
	}
	return nil
}

func list(){
	for _,v := range data{
		fmt.Println(v)
	}
}

func random(min,max int) int{
	return rand.Intn(max-min) + min
}
func getString(len int64) string {
	temp := ""
	startChar := "A"
	var i int64 = 1
	for {
		myRand := random(MIN, MAX)
		newChar := string(startChar[0] + byte(myRand))
		temp = temp + newChar
		if i == len {
			break
		}
		i++
	}
	return temp
}

func populate(n int, s []Entry) {
	for i := 0; i < n; i++ {
	name := getString(4)
	surname := getString(5)
	n := strconv.Itoa(random(100, 199))
	data = append(data, Entry{name, surname, n})
	}
}

func main(){
	arguments := os.Args
	if len(arguments) == 1{
		exe := path.Base(arguments[0])
		fmt.Printf("Usage: %s search|list <arguments>\n", exe)
		return
	}

	SEED := time.Now().Unix()
	rand.Seed(SEED)

	n := 100
	populate(n, data)
	fmt.Printf("Data has %d entries.\n", len(data))

	switch arguments[1] {
	case "search":
		if len(arguments) != 3 {
			fmt.Println("usage: search `surname`")
			return
		}
		result := search(arguments[2])
		if result == nil {
			fmt.Println(arguments[2],"surname not found")
			return
		}
		fmt.Println(*result)

	case "list":
		list()
	default:
		fmt.Println("Not a valid option")
	}
}