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
