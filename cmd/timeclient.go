package main

import (
	"bufio"
	"fmt"
	"net"
	"os"
	"strings"
)

func main() {
	in := bufio.NewReader(os.Stdin)

	fmt.Print("Введите IP адрес к которому вы хотите подключиться: ")
	addr := readStr(in)

	con, err := net.Dial("tcp", fmt.Sprintf("%s:1303", addr))
	defer con.Close()

	if err != nil {
		fmt.Println(err)
		return
	}

	reply := make([]byte, 1024)

	_, err = con.Read(reply)
	if err != nil {
		fmt.Println(err)
		return
	}

	fmt.Println(string(reply))
}

func readStr(in *bufio.Reader) string {
	nStr, _ := in.ReadString('\n')
	nStr = strings.ReplaceAll(nStr, "\r", "")
	nStr = strings.ReplaceAll(nStr, "\n", "")
	return nStr
}
