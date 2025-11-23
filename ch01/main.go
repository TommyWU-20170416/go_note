package main

import "ch01/neighbor"

func main() {
	// 與 main 同 package 但沒有定義或 import 的函式，不能 go run .\main.go，必須使用 go run .
	sayHello()
	// 因為有引入所以可以執行 go run .\main.go
	neighbor.HelloNeighbor()
}
