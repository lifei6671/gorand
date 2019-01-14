# gorand

Golang 随机字符串生成库

## 使用

```bash
go get github.com/lifei6671/gorand
```

## 实例

```go
//生成指定长度的字符串，使用 Apache 的算法
gorand.RandomAlphabetic(10)

//使用随机数算法生成字符串
gorand.KRand(20, KC_RAND_KIND_ALL)

//生成 UUID4 字符串
gorand.NewUUID4().String()
```
