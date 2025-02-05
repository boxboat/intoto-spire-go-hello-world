package main

import (
	"flag"
	"fmt"
	"log"
	"net/http"
)

const responseTemplate = `
<html>
    <link rel="icon" href="data:;base64,=">
    <body>
    <p>You've been here %v times</p>
    </body>
</html>
`

func main() {
	port := flag.Int("p", 8080, "port to listen on")
	flag.Parse()

	visitsByIp := make(map[string]uint)
	http.HandleFunc("/", func(w http.ResponseWriter, r *http.Request) {
		ip := getRequestIp(r)
		visits := visitsByIp[ip]
		visits += 1
		visitsByIp[ip] = visits
		response := fmt.Sprintf(responseTemplate, visits)
		w.Header().Set("Content-Type", "text/html; charset=utf-8")
		w.Write([]byte(response))
		log.Println(fmt.Sprintf("Visited by %v", ip))
	})

	addr := fmt.Sprintf(":%v", *port)
	log.Println(fmt.Sprintf("Listening on %v", addr))
	log.Fatal(http.ListenAndServe(addr, nil))
}

func getRequestIp(r *http.Request) string {
	ip := r.Header.Get("X-FORWARDED-FOR")
	if len(ip) == 0 {
		ip = r.RemoteAddr
	}

	return ip
}
