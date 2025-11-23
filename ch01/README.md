# Go Modules包管理
> 課程連結:[Golang學習記03，Go Modules包管理 (go mod項目管理)](https://www.youtube.com/watch?v=J5A6mbtezyU&list=PLBjZhzRvV2ChPTPNDx_apHdKa9Ha7LVpN&index=3)

## 創建 hello.go
- [hello.go](hello.go)
- 每個 package main 內都會有 main func

## 編譯 hello.go
```bash=
PS C:\Users\Tommy\OneDrive\文件\project\go_note\ch01> go build .\hello.go
```
編譯後依照環境包出執行檔案，windows => .exe 檔案

## 執行 hello(直接運行不編譯)
```bash=
PS C:\Users\Tommy\OneDrive\文件\project\go_note\ch01> go run .\hello.go 
# Hello World
```

## 使用 code runner 執行
在 hello.go 按下右鍵，執行 Run code

## 引入 MOD
告訴電腦這是什麼版本，哪個專案
```bash=
PS C:\Users\Tommy\OneDrive\文件\project\go_note\ch01> go mod init go_note
# go: creating new go.mod: module go_note
# go: to add module requirements and sums:
#         go mod tidy
```

# 在不同位置調用

## 同包不同文件
> main.go > hello.go
> 在 main.go 執行 go run .\main.go

Go 的執行路徑與檔案引用有關。
如果你只執行 main.go（例如 go run main.go），Go 只會編譯這個檔案，找不到 hello.go 中的 sayHello() 定義，導致錯誤。

而在 ch01 資料夾下執行 `go run .`，Go 會自動編譯同目錄下所有 .go 檔案（不含測試檔），所以可以找到 sayHello()。

## 不同包不同文件
> main.go > /neighbor/neighbor.go

透過引入 `mod/package` 以及方法字母大寫(視為 public) 可以使用到其他的方法