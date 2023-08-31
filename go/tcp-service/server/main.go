/*
 * Copyright (c) 2023, WSO2 LLC. (https://www.wso2.com/) All Rights Reserved.
 *
 * WSO2 LLC. licenses this file to you under the Apache License,
 * Version 2.0 (the "License"); you may not use this file except
 * in compliance with the License.
 * You may obtain a copy of the License at
 *
 *     http://www.apache.org/licenses/LICENSE-2.0
 *
 * Unless required by applicable law or agreed to in writing,
 * software distributed under the License is distributed on an
 * "AS IS" BASIS, WITHOUT WARRANTIES OR CONDITIONS OF ANY
 * KIND, either express or implied. See the License for the
 * specific language governing permissions and limitations
 * under the License.
 */

package main

import (
	"context"
	"io"
	"log"
	"net"
	"os"
	"os/signal"
	"sync"
	"syscall"
)

// handles incoming client connection
func handleConnection(conn net.Conn) {
	defer conn.Close()
	log.Printf("Accepted connection from %v", conn.RemoteAddr())

	// Read data from the client
	for {
		buffer := make([]byte, 1024)
		n, err := conn.Read(buffer)
		if err != nil && err != io.EOF {
			log.Printf("Error reading data : %v", err)
			break
		}

		if n > 0 {
			log.Printf("Received data from client : %v", string(buffer[:n]))
			message := "Hello from TCP server"
			_, err = conn.Write([]byte(message))
			if err != nil {
				log.Printf("Error writing data : %v", err)
				break
			}

			log.Printf("Sent data to client : {%v}", message)
		} else {
			return
		}
	}
}

func main() {
	// Creating a TCP listener
	listener, err := net.Listen("tcp", ":8080")
	if err != nil {
		log.Fatalf("Error starting TCP server : %v", err)
	}

	defer listener.Close()
	log.Printf("TCP server listening on %s", listener.Addr())

	ctx, cancel := context.WithCancel(context.Background())
	wg := &sync.WaitGroup{}

	// Capture signals to handle graceful shutdown
	sigChan := make(chan os.Signal, 1)
	signal.Notify(sigChan, os.Interrupt, syscall.SIGTERM)

	// Handle graceful shutdown
	go func() {
		sig := <-sigChan
		log.Println("Received signal:", sig)
		cancel()

		close(sigChan)
		listener.Close()
		wg.Wait()
	}()

	// Accept incoming connections
	for {
		conn, err := listener.Accept()
		if err != nil {
			select {
			case <-ctx.Done():
				log.Println("Server shutting down gracefully...")
				wg.Wait()
				return
			default:
				log.Println("Error accepting connection:", err)
				continue
			}
		}

		wg.Add(1)
		// Handle incoming connection in a new goroutine
		go func(conn net.Conn) {
			handleConnection(conn)
			wg.Done()
		}(conn)
	}
}
