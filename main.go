package main

import (
	"bufio"
	"fmt"
	"io"
	"log"
	"math/rand"
	"net"
	"os"
	"strconv"
	"time"
)

func handleConnection(c net.Conn) {
	fmt.Printf("Serving %s\n", c.RemoteAddr().String())
	defer c.Close()
	buf := make([]byte, 0, 4096) // big buffer
	tmp := make([]byte, 256)     // using small tmo buffer for demonstrating
	for {
		n, err := c.Read(tmp)
		if err != nil {
			if err != io.EOF {
				//fmt.Println("read error:", err)
			} else {
				break
			}
		}
		//fmt.Println("got", n, "bytes.")
		buf = append(buf, tmp[:n]...)
	}

	c.Write([]byte(string(buf)))
	fmt.Println(string(buf))
}
func LoadPorts(file_name string) []int {
	var rtn_value []int
	inFile, err := os.Open(file_name)
	if err != nil {
		fmt.Println(err.Error() + `: ` + file_name)
		log.Fatal("no such port file.")
	}
	defer inFile.Close()

	scanner := bufio.NewScanner(inFile)

	for scanner.Scan() {
		log.Println(scanner.Text())
		var x, _ = strconv.Atoi(scanner.Text())
		rtn_value = append(rtn_value, x)
	}
	return rtn_value
}
func listen(port int) {
	log.Println("will listen on ", "0.0.0.0:"+string(port))
	l, err := net.Listen("tcp4", "0.0.0.0:"+strconv.Itoa(port))
	if err != nil {
		fmt.Println(err)
		return
	}
	rand.Seed(time.Now().Unix())

	for {
		c, err := l.Accept()
		if err != nil {
			fmt.Println(err)
			return
		}
		go handleConnection(c)
	}
}
func main() {
	arguments := os.Args
	if len(arguments) == 1 {
		fmt.Println("Please provide a input file!")
		return
	}
	ports := LoadPorts(arguments[1])
	for _, port := range ports {
		go func(i int) { listen(i) }(port)
	}
	for ; ; {
		time.Sleep(1 * time.Second)
	}
}
