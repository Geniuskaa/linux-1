package main

import (
	"context"
	"fmt"
	"net"
	"os"
	"os/signal"
	"time"
)

func main() {
	ctx, _ := signal.NotifyContext(context.Background(), os.Interrupt)

	lstr, err := net.Listen("tcp", "localhost:1303")
	if err != nil {
		fmt.Println(err)
		lstr.Close()
		return
	}

	go func() {
		for {
			con, err := lstr.Accept()
			if err != nil {
				fmt.Println(err)
				con.Close()
				return
			}

			go func(con net.Conn) {
				defer con.Close()
				_, err := con.Write([]byte(time.Now().Format("02.01.2006 15:04")))
				if err != nil {
					fmt.Println(err)
					return
				}
			}(con)
		}
	}()

	<-ctx.Done()
	lstr.Close()
	fmt.Println("Сигнал завершения получен! Выключаем сервер.")
	return
}
