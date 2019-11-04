package main

import (
	"bufio"
	"fmt"
	"io"
	"net"
	"os"
	"time"
)

type Cislo int64
type Vektor []Cislo

func sucin(a Vektor, b Vektor) (s Cislo) {
	if len(a) != len(b) {
		return 0
	}
	s = 0
	for i := 0; i < len(a); i++ {
		s += a[i] * b[i]
	}
	return s
}

func handleConnection(conn net.Conn) {
	r := bufio.NewReader(conn)
	w := bufio.NewWriter(conn)
	for {
		inbytes, _, err := r.ReadLine()
		if err == io.EOF || len(inbytes) < 3 {
			break
		}
		task := buffer2Task(inbytes)
		fmt.Printf("prisla robota: %v\n", task)

		if task.I1 == -1 && task.I2 == -1 {
			fmt.Println("Client oznamil koniec brigady!")
			res := Result{-1, -1, Cislo(workTime.Milliseconds())}
			_, _ = w.WriteString(res.toString())
			_ = w.Flush()
			workTime = time.Duration(0)
		} else {
			start := time.Now()
			//time.Sleep(time.Second)
			res := Result{task.I1, task.I2, sucin(task.V1, task.V2)}
			fmt.Printf("vysledok: %v\n", res)
			_, _ = w.WriteString(res.toString())
			_ = w.Flush()
			workTime += time.Since(start)
		}
	}
}

var workTime = time.Duration(0)
// skalarnyServer.exe :port
func main() { 
	if len(os.Args) < 2 {
		fmt.Println("pouzitie:")
		fmt.Println("server.exe port\t- spusti skalarny server na porte")
	} else {
		port := os.Args[1]
		fmt.Println("pocuvam " + port)
		ln, err := net.Listen("tcp", ":" + port)
		if err != nil {
			fmt.Println("nepodarilo sa posadit Listenera na " + port)
		} else {
			for {
				conn, errc := ln.Accept()
				if errc != nil {
					fmt.Println("nepodarilo sa otvorit konekciu")
					return
				}
				go handleConnection(conn)
			}
		}
	}
}
