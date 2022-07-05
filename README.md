# Pool
Coroutine Pool  
# Quick start
```
    p := NewPool(4)  
    go p.Run()  
    go http.ListenAndServe("0.0.0.0:6060", nil)  //pprof查看协程个数
    for {
        p.Put(Info{Param: "http://www.boyaa.com", Func: Proc})
        time.Sleep(time.Millisecond * 500)
    }
```

