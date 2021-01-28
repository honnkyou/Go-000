学习笔记
1. 用 Go 实现一个 tcp server ，用两个 goroutine 读写 conn，两个 goroutine 通过 chan 可以传递 message，能够正确退出
Week09 作业地址：
https://github.com/Go-000/Go-000/issues/82

refer to the example from ppt :

'''go
func main() {
	listen, err := net.Listen("tcp", "127.0.0.1:10000")
	if err != nil {
		log.Fatalf("listen error: %v\n", err)
	}
	for {
		conn, err := listen.Accept()
		if err != nil {
			log.Printf("accept error: %v\n", err)
			continue
		}
		// 开始goroutine监听连接
		go handleConn(conn)
	}
}

func handleConn(conn net.Conn) {
	defer conn.Close()
	// 读写缓冲区
	rd := bufio.NewReader(conn)
	wr := bufio.NewWriter(conn)
	for {
		line, _, err := rd.ReadLine()
		if err != nil {
			log.Printf("read error: %v\n", err)
			return
		}
		wr.WriteString("hello ")
		wr.Write(line)
		wr.Flush() // 一次性syscall
	}
}
'''