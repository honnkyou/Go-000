package main

import (
	"context"
	"net/http"
	"log"
	
	"golang.org/x/sync/errgroup"

	"os"
	"os/signal"

)

//基于 errgroup 实现一个 http server 的启动和关闭 ，
//以及 linux signal 信号的注册和处理，要保证能够一个退出，全部注销退出。
func serveApp(stop <-chan struct{}) error{
	mux := http.NewServeMux()
	mux.HandleFunc("/", func(resp http.ResponseWriter, req *http.Request){
		fmt.Fprintln(resp, "Hello, Qcon!")
	})
	srv := http.Server{Addr: "0.0.0.0:8080", handler: mux}
	go func(){
		sigint := make(chan os.Signal, 1)
		signal.Notify(sigint, os.Interrupt)
		select{
		case <-stop:
		case <-sigint:
		}
		srv.Shutdown(context.Background())
	}
	return srv.ListenAndServe()

}

func debugApp(stop <-chan struct{}) error{
	srv := http.Server{Addr: "127.0.0.1:8081"}
	go func(){
		sigint := make(chan os.Signal, 1)
		signal.Notify(sigint, os.Interrupt)
		select{
		case <-stop :
			log.Print("servers stopped by error.")

		case <-sigint :
			log.Print("servers shutdowned by signal.")
		}
		srv.Shutdown(context.Background())
	}
	return srv.ListenAndServe()
}


func main() {
	// make a stop signal for all servers stopping
	stop := make(chan struct{})

	// start http server service
	g := new(errgroup.Group)
	g.Go(serveApp(stop))
	g.Go(debugApp(stop))

	// shutdown all http servers
	if err := g.Wait(); err != nil {
		close(stop)
	}
	
}
//	ctx, cancel := context.withCancel(context.Background())

//	defer cancel()

// func registSig() {
// 	sigint := make(chan os.Signal, 1)
// 	signal.Notify(sigint, os.Interrupt)
// 	<-sigint

// 	// We received an interrupt signal, shut down.
// 	if err := srv.Shutdown(context.Background()); err != nil {
// 		// Error from closing listeners, or context timeout:
// 		log.Printf("HTTP server Shutdown: %v", err)
// 	}
// 	close(idleConnsClosed)
// }
