package main

import (
	"fmt"
	"net"
	"os"
)

func filterConnection(conn net.Conn, rules []string) {
	remoteAddr := conn.RemoteAddr().String()
	allowed := false

	for _, rule := range rules {
		if rule == remoteAddr {
			allowed = true
			break
		}
	}

	if allowed {
		fmt.Printf("Allowed connection from %s\n", remoteAddr)
	} else {
		fmt.Printf("Blocked connection from %s\n", remoteAddr)
		conn.Close()
	}
}

func getenv(key string, def string) string {
	val, ok := os.LookupEnv(key)
	if !ok {
		val = def
	}
	return val
}

func main() {
	port := getenv("PORT", "8080")
	rules := []string{"192.168.0.1", "192.168.0.2"}

	listener, err := net.Listen("tcp", ":"+port)
	if err != nil {
		fmt.Println("Error starting server:", err)
		os.Exit(1)
	}
	fmt.Printf("Server listening on port %s\n", port)

	for {
		conn, err := listener.Accept()
		if err != nil {
			fmt.Println("Error accepting connection:", err)
			continue
		}
		fmt.Printf("Server listening ro rules %v", rules)
		go filterConnection(conn, rules)
	}
}
