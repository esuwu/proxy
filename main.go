package main

import (
	"crypto/tls"
	"fmt"
	"github.com/valyala/fasthttp"
	"log"
)


func handleHTTP(ctx *fasthttp.RequestCtx) {
	fmt.Println(ctx.Request.String())
	req := fasthttp.AcquireRequest()
	req.SetRequestURI(ctx.URI().String())
	req.Header = ctx.Request.Header
	resp := fasthttp.AcquireResponse()
	client := &fasthttp.Client{}
	client.Do(req, resp)

	fmt.Println("kek")
	fmt.Print(resp)
}

func handleHTTPS(ctx *fasthttp.RequestCtx) {
	certificates := tls.Certificate{

	}
	conf := &tls.Config{

		//InsecureSkipVerify: true,
	}
	fmt.Println(string(ctx.RemoteAddr().String()))
	conn, err := tls.Dial("tcp", ctx.URI().String(), conf)
	if err != nil {
		log.Println(err)
		return
	}
	defer conn.Close()

	n, err := conn.Write([]byte("hello\n"))
	if err != nil {
		log.Println(n, err)
		return
	}

	buf := make([]byte, 100)
	n, err = conn.Read(buf)
	if err != nil {
		log.Println(n, err)
		return
	}

	println(string(buf[:n]))
}



func main() {
	server := &fasthttp.Server{
		Handler: fasthttp.RequestHandler(func(ctx *fasthttp.RequestCtx) {
			method := string(ctx.Method())
			if method == fasthttp.MethodConnect {
				handleHTTPS(ctx)
			} else {
				handleHTTP(ctx)
			}
		}),
	}

	err := fasthttp.ListenAndServe(":8081", server.Handler)
	if err != nil {
		log.Println(err)
		return
	}

	log.Println("Proxy started on 8081 port: ")
}
