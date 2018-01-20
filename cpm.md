1、性能监控的目标:`定位那些地方耗cpu,那些地方耗内存`<br>
2、怎么做?(个人要把路由指向net/hhtp/pprof包下的函数)<br>
  1)、go tool pprof http://xxxx/debug/pprof/profile 针对cpu的<br>
      go tool pprof http://xxx/debug/pprof/heap 针对内存的
      
  2)、使用top[N]命令查看那些最消耗cpu,那些最消耗内存<br>
          接着可以使用peek命令看调用情况或者使用list命令查看源码<br>
          
  3)、如果观看的不直接，可以使用png/jpe这样的命令把情况生成图片，再进行定位<br>
      当然也可以使用web命令来查看情况<br>
3、下一步:希望在实际情况使用它<br>
4、其他(手动编写程序)<br>
cpu概要文件:
```go
func startCPUProfile() {
  if *cpuProfile != "" {
    f, err := os.Create(*cpuProfile)
    if err != nil {
      fmt.Fprintf(os.Stderr, "Can not create cpu profile output file: %s",err)
      return
    }
    if err := pprof.StartCPUProfile(f); err != nil {
      fmt.Fprintf(os.Stderr, "Can not start cpu profile: %s", err)
      f.Close()
      return
    }
  }
}
```
```go
func stopCPUProfile() {
  if *cpuProfile != "" {
    pprof.StopCPUProfile()
  }
}
```
内存概要文件:
```go
func startMemProfile() {
  if *memProfile != "" && *memProfileRate > 0 {
    runtime.MemProfileRate = *memProfileRate
  }
}
```
```go
func stopMemProfile() {
  if *memProfile != "" {
    f, err := os.Create(*memProfile)
    if err != nil {
      fmt.Fprintf(os.Stderr, "Can not create mem profile output file: %s", err)
      return
    }
    if err = pprof.WriteHeapProfile(f); err != nil {
      fmt.Fprintf(os.Stderr, "Can not write %s: %s", *memProfile, err)
    }
    f.Close()
  }
}
```
阻塞概要文件:
```go
func startBlockProfile() {
  if *blockProfile != "" && *blockProfileRate > 0 {
    runtime.SetBlockProfileRate(*blockProfileRate)
  }
}
```
```go
func stopBlockProfile() {
  if *blockProfile != "" && *blockProfileRate >= 0 {
    f, err := os.Create(*blockProfile)
    if err != nil {
      fmt.Fprintf(os.Stderr, "Can not create block profile output file: %s", err)
      return
    }
    if err = pprof.Lookup("block").WriteTo(f, 0); err != nil {
      fmt.Fprintf(os.Stderr, "Can not write %s: %s", *blockProfile, err)
    }
    f.Close()
  }
}
```
