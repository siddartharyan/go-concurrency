package main

import (
	"context"
	"io"
	"log"
	"net/http"
	"os"
	"time"
)

func main() {

	ctx, cancel := context.WithTimeout(context.Background(), 10000*time.Millisecond)

	defer cancel()

	req, err := http.NewRequest("GET", "https://andcloud.io", nil)

	req = req.WithContext(ctx)

	if err != nil {
		log.Fatal(err)
	}

	res, err := http.DefaultClient.Do(req)

	if err != nil {
		log.Fatal(err)
	}

	defer res.Body.Close()

	io.Copy(os.Stdout, res.Body)
}
