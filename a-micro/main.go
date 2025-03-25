package main

import (
	"fmt"
	"io"
	"net/http"
)

func main() {
	http.HandleFunc("/", func(w http.ResponseWriter, r *http.Request) {
		if r.Method == http.MethodGet {
			res, err := http.Get("http://b-micro:8081")
			if err != nil {
				fmt.Fprintln(w, "A-Micro: Error to connect to B-Micro")
			}
			defer res.Body.Close()

			body, err := io.ReadAll(res.Body)
			if err != nil {
				fmt.Fprintln(w, "A-Micro: Error to react the Body")
			}
			str := fmt.Sprintf("A-Micro: response from B-Micro is — %v", string(body))
			fmt.Fprintln(w, str)
		} else {
			fmt.Fprintln(w, "A-Micro: Error")
		}
	})

	if err := http.ListenAndServe(":8080", nil); err != nil {
		fmt.Println("Server failed:", err)
	} else {
		fmt.Println("A-Micro is launched")
	}
}
