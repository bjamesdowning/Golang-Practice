package main

import (
	"bufio"
	"fmt"
	"log"
	"net"
	"strings"
)

func main() {
	li, err := net.Listen("tcp", "127.0.0.1:8080")
	if err != nil {
		log.Fatalln(err.Error())
	}
	defer li.Close()

	for {
		conn, err := li.Accept()
		if err != nil {
			log.Fatalln(err.Error())
		}
		go handle(conn)
	}
}

func handle(conn net.Conn) {
	defer conn.Close()
	request(conn)
}

func request(conn net.Conn) {
	i := 0
	scanner := bufio.NewScanner(conn)
	for scanner.Scan() {
		ln := scanner.Text()
		fmt.Println(ln)
		if i == 0 {
			mux(conn, ln)
		}
		if ln == "" {
			break //headers are gone
		}
		i++
	}
}

func mux(conn net.Conn, ln string) {
	m := strings.Fields(ln)[0] //method
	u := strings.Fields(ln)[1] // URL
	fmt.Println("***METHOD: ", m)
	fmt.Println("***URL: ", u)

	if m == "GET" && u == "/" {
		index(conn)
	}
	if m == "GET" && u == "/about" {
		about(conn)
	}
	if m == "GET" && u == "/apply" {
		apply(conn)
	}
	if m == "GET" && u == "/contact" {
		contact(conn)
	}
	if m == "POST" && u == "/apply" {
		applyProcess(conn)
	}

}

func index(conn net.Conn) {
	body := `<!DOCTYPE html><html lang="en"><head><meta
	charet="UTF-8"><title></title></head><body>
		<strong>INDEX</strong><br>
		<a href="/">index</a><br>
		<a href="/about">about</a><br>
		<a href="/apply">apply</a><br>
		<a href="/contact">contact</a><br>
		</body></html> `

	fmt.Fprint(conn, "HTTP/1.1 200 OK \r\n")
	fmt.Fprint(conn, "Content-Length: %d\r\n", len(body))
	fmt.Fprint(conn, "Content-Type: text/html\r\n")
	fmt.Fprint(conn, "\r\n")
	fmt.Fprint(conn, body)
}

func about(conn net.Conn) {
	body := `<!DOCTYPE html><html lang="en"><head><meta
	charet="UTF-8"><title></title></head><body>
		<strong>ABOUT</strong><br>
		<a href="/">index</a><br>
		<a href="/about">about</a><br>
		<a href="/apply">apply</a><br>
		<a href="/contact">contact</a><br>
		</body></html> `

	fmt.Fprint(conn, "HTTP/1.1 200 OK \r\n")
	fmt.Fprint(conn, "Content-Length: %d\r\n", len(body))
	fmt.Fprint(conn, "Content-Type: text/html\r\n")
	fmt.Fprint(conn, "\r\n")
	fmt.Fprint(conn, body)
}

func contact(conn net.Conn) {
	body := `<!DOCTYPE html><html lang="en"><head><meta
	charet="UTF-8"><title></title></head><body>
		<strong>CONTACT</strong><br>
		<a href="/">index</a><br>
		<a href="/about">about</a><br>
		<a href="/apply">apply</a><br>
		<a href="/contact">contact</a><br>
		</body></html> `

	fmt.Fprint(conn, "HTTP/1.1 200 OK \r\n")
	fmt.Fprint(conn, "Content-Length: %d\r\n", len(body))
	fmt.Fprint(conn, "Content-Type: text/html\r\n")
	fmt.Fprint(conn, "\r\n")
	fmt.Fprint(conn, body)
}

func apply(conn net.Conn) {
	body := `<!DOCTYPE html><html lang="en"><head><meta
	charet="UTF-8"><title></title></head><body>
		<strong>APPLY</strong><br>
		<a href="/">index</a><br>
		<a href="/about">about</a><br>
		<a href="/apply">apply</a><br>
		<a href="/contact">contact</a><br>
		<br>
		<br>
		<form action="/apply" method="post">
		<dt>Key:
		<dd><input type=text name=key>
		<dt>Value:
		<dd><input type=text name=value>
		<br>
		<input type="submit" value="test">
		</form>
		</body></html>`

	fmt.Fprint(conn, "HTTP/1.1 200 OK \r\n")
	fmt.Fprint(conn, "Content-Length: %d\r\n", len(body))
	fmt.Fprint(conn, "Content-Type: text/html\r\n")
	fmt.Fprint(conn, "\r\n")
	fmt.Fprint(conn, body)
}

func applyProcess(conn net.Conn) {
	body := `<!DOCTYPE html><html lang="en"><head><meta
	charet="UTF-8"><title></title></head><body>
		<strong>APPLY PROCESS</strong><br>
		<a href="/">index</a><br>
		<a href="/about">about</a><br>
		<a href="/apply">apply</a><br>
		<a href="/contact">contact</a><br>
		</body></html> `

	fmt.Fprint(conn, "HTTP/1.1 200 OK \r\n")
	fmt.Fprint(conn, "Content-Length: %d\r\n", len(body))
	fmt.Fprint(conn, "Content-Type: text/html\r\n")
	fmt.Fprint(conn, "\r\n")
	fmt.Fprint(conn, body)
}
