package main

import (
	"bytes"
	"fmt"
	"io"
	"os"
	"sync"
	"time"
)

var bufPool = sync.Pool{
	New: func() interface{} {
		fmt.Println("buffer created")
		return new(bytes.Buffer)
	},
}

func main() {

	log(os.Stdout, "debug-1")
	log(os.Stdout, "debug-2")

}

func log(w io.Writer, debug string) {
	b := bufPool.Get().(*bytes.Buffer)
	b.Reset()
	b.WriteString(time.Now().Format("15:05:55"))
	b.WriteString(":")
	b.WriteString(debug)
	b.WriteString("\n")

	w.Write(b.Bytes())
	bufPool.Put(b)
}
