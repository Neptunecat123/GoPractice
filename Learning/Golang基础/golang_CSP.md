# 并发机制

## CSP

Communicating squential processes

channel

## 多路选择机制

select

### 多渠道select

### 超时

```go
select {
    case ret:= <-Chan1:
        fmt.Println(ret)
    case <-time.After(time.Second * 30):
        fmt.Println("timeout")
}
```

 