package main

import (
	"fmt"
	"net"
	"strings"
	"sync"
	"time"
)

func main() {
	ipArr := make([]string, 0, 4)
	ipArr = append(ipArr, "localhost")

	interfaces, err := net.Interfaces()
	if err != nil {
		panic(err)
	}

	for _, interf := range interfaces {
		// Список адресов для каждого сетевого интерфейса
		addrs, err := interf.Addrs()
		if err != nil {
			panic(err)
		}

		if !strings.Contains(interf.Name, "enp") {
			continue
		}

		for i, add := range addrs {
			if i > 0 {
				break
			}

			if ip, ok := add.(*net.IPNet); ok {
				ipArr = append(ipArr, ip.IP.String())
			}
		}
	}

	wg := sync.WaitGroup{}

	for _, val := range ipArr {
		wg.Add(1)

		go func(val string) {
			defer wg.Done()
			w := sync.WaitGroup{}

			lstr, err := net.Listen("tcp", fmt.Sprintf("%s:1303", val))
			defer func() {
				lstr.Close()
			}()
			if err != nil {
				fmt.Println(err)
				return
			}

			w.Add(1)
			go func() {
				defer w.Done()
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

			w.Wait()

		}(val)

	}

	wg.Wait()
	fmt.Println("Сигнал завершения получен! Выключаем сервер.")
	return

}
