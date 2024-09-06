package pocketbase

import (
	"fmt"
	"io"
	"net/http"
)

type PocketBase struct {
	Addr string
}

func NewPocketBase(addr string) *PocketBase {
	return &PocketBase{
		Addr: addr,
	}
}

func request(method, query string, reader io.Reader) []byte {
	req, err := http.NewRequest(method, query, reader)
	if err != nil {
		panic(err)
	}
	req.Header.Set("Content-Type", "application/json")
	req.Header.Set("Authorization", fmt.Sprintf("Bearer %s", Bearer))

	res, err := http.DefaultClient.Do(req)
	if err != nil {
		panic(err)
	}
	defer func() {
		if err = res.Body.Close(); err != nil {
			panic(err)
		}
	}()
	if res.StatusCode == http.StatusOK {
		if data, err := io.ReadAll(res.Body); err != nil {
			panic(err)
		} else {
			return data
		}
	} else {
		panic(res.StatusCode)
	}
}