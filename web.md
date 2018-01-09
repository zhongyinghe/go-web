go实现http的源码解读
==================
代码如下:
```
func main() {
	http.HandleFunc("/", sayhelloName) //设置访问的路由
	err := http.ListenAndServe(":9090", nil) //设置监听的端口
	if err != nil {
		log.Fatal("ListenAndServe: ", err)
	}
}
```
http包下的ListenAndServe()函数如下:
```
func ListenAndServe(addr string, handler Handler) error {
	server := &Server{Addr: addr, Handler: handler}
	return server.ListenAndServe()
}
```
根据代码，可以看出，它的实质是:创建Server对象，并调用Server的ListenAndServe()方法.<br>
而server.ListenAndServe()的代码如下:
```go
func (srv *Server) ListenAndServe() error {
	addr := srv.Addr
	if addr == "" {
		addr = ":http"
	}
	ln, err := net.Listen("tcp", addr)
	if err != nil {
		return err
	}
	return srv.Serve(tcpKeepAliveListener{ln.(*net.TCPListener)})
}
```
可以看出,它会调用Server的Serve()方法,而Serve()方法的代码如下:
```go
func (srv *Server) Serve(l net.Listener) error {
	defer l.Close()
	if fn := testHookServerServe; fn != nil {
		fn(srv, l)
	}
	var tempDelay time.Duration // how long to sleep on accept failure

	if err := srv.setupHTTP2_Serve(); err != nil {
		return err
	}

	srv.trackListener(l, true)
	defer srv.trackListener(l, false)

	baseCtx := context.Background() // base is always background, per Issue 16220
	ctx := context.WithValue(baseCtx, ServerContextKey, srv)
	ctx = context.WithValue(ctx, LocalAddrContextKey, l.Addr())
	for {
		rw, e := l.Accept()
		if e != nil {
			select {
			case <-srv.getDoneChan():
				return ErrServerClosed
			default:
			}
			if ne, ok := e.(net.Error); ok && ne.Temporary() {
				if tempDelay == 0 {
					tempDelay = 5 * time.Millisecond
				} else {
					tempDelay *= 2
				}
				if max := 1 * time.Second; tempDelay > max {
					tempDelay = max
				}
				srv.logf("http: Accept error: %v; retrying in %v", e, tempDelay)
				time.Sleep(tempDelay)
				continue
			}
			return e
		}
		tempDelay = 0
		c := srv.newConn(rw)
		c.setState(c.rwc, StateNew) // before Serve can return
		go c.serve(ctx)
	}
}
```
可以看出，它使用l.Accept()等待用户请求的到来，如果有请求到来则创建conn对象(c := srv.newConn(rw)),然后开启goroutine:`go c.serve(ctx)`<br>
可以看出，每个http请求都会开启一个goroutine,这就是它高并发的体现.<br>
而c.serve(ctx)方法会调用serverHandler的ServeHTTP()方法,而ServeHTTP()的代码如下:
```
func (sh serverHandler) ServeHTTP(rw ResponseWriter, req *Request) {
	handler := sh.srv.Handler
	if handler == nil {
		handler = DefaultServeMux
	}
	if req.RequestURI == "*" && req.Method == "OPTIONS" {
		handler = globalOptionsHandler{}
	}
	handler.ServeHTTP(rw, req)
}
```
这里是关键，如果在原先的http.ListenAndServe()这里，设置了第二个参数，则它就会按照设置的路由进行处理,个人路由器设计就是在这第二个参数入手。如果
没有设置第二个参数，则它会调用默认的处理handler:DefaultServeMux<br>
<br>
那么DefaultServeMux是一个什么样的对象?它是ServeMux对象，它的代码定义如下:
```go
type ServeMux struct {
	mu    sync.RWMutex
	m     map[string]muxEntry
	hosts bool // whether any patterns contain hostnames
}
```
而它的ServeHTTP()方法如下:
```
func (mux *ServeMux) ServeHTTP(w ResponseWriter, r *Request) {
	if r.RequestURI == "*" {
		if r.ProtoAtLeast(1, 1) {
			w.Header().Set("Connection", "close")
		}
		w.WriteHeader(StatusBadRequest)
		return
	}
	h, _ := mux.Handler(r)//对请求进行函数匹配
	h.ServeHTTP(w, r)
}
```
当它匹配到Handler对象，则调用该对象的ServeHTTP()方法,那么该对象是怎么样的呢?
Handler是接口,代码如下:
```
type Handler interface {
	ServeHTTP(ResponseWriter, *Request)
}
```
而HandlerFunc实现了该接口,它的代码如下:
```
// Handler that calls f.
type HandlerFunc func(ResponseWriter, *Request)

// ServeHTTP calls f(w, r).
func (f HandlerFunc) ServeHTTP(w ResponseWriter, r *Request) {
	f(w, r)
}
```
当我们在main()函数中调用http.HandleFunc("/", sayhelloName)时,该sayhellName()函数必须是HandlerFunc类型(定义函数类型，让相同签名的函数自动实现某个接口)。所以终极会调用如sayhelloName()这样的函数.那么sayhelloName这样的函数保存在哪里呢?<br>
它们保存在ServeMux结构体的m中,这个m是一个map[string]muxEntry类型，而muxEntry的代码如下:
```
type muxEntry struct {
	explicit bool
	h        Handler
	pattern  string
}
```
那么我们每次调用http.HandleFunc()时都会创建一个muxEntry对象，并保存进ServeMux的m中，以便在请求时进行匹配

