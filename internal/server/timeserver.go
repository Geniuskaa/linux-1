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
	defer func() {
		lstr.Close()
	}()
	if err != nil {
		fmt.Println(err)
		return
	}

	go func() {
		for {
			con, err := lstr.Accept()
			if err != nil {
				fmt.Println(err)
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
	fmt.Println("Сигнал завершения получен! Выключаем сервер.")
	return
}
