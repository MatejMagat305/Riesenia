package main

import (
	"bufio"
	"fmt"
	"net"
	"sync"
)
type Worker struct {
	host   string
	port   string
	conn   net.Conn
	busy  bool
	reader *bufio.Reader
	writer *bufio.Writer
	mux sync.Mutex

}
func NewWorker(host string, port string) *Worker {
	conn, err := net.Dial("tcp", host+":"+(port))
	if err != nil {	fmt.Printf("worker %d zlyhal\n", port)	}
	return &Worker{host, port, conn, false,
		bufio.NewReader(conn), bufio.NewWriter(conn),sync.Mutex{} }
}



func (w *Worker) doit(t *Task) Result{
	w.writer.WriteString(t.toString())
	w.writer.Flush()
	fmt.Printf("worker %v pracuje\n", w.port)
	resbytes, _, _ := w.reader.ReadLine()
	res := buffer2Result(resbytes)
	w.busy = false
	//println(res.I1, "   ", res.I2,"   ",res.Skalar)
	return res
}