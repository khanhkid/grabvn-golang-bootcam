/*
 *
 * Copyright 2015 gRPC authors.
 *
 * Licensed under the Apache License, Version 2.0 (the "License");
 * you may not use this file except in compliance with the License.
 * You may obtain a copy of the License at
 *
 *     http://www.apache.org/licenses/LICENSE-2.0
 *
 * Unless required by applicable law or agreed to in writing, software
 * distributed under the License is distributed on an "AS IS" BASIS,
 * WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND, either express or implied.
 * See the License for the specific language governing permissions and
 * limitations under the License.
 *
 */

// Package main implements a client for Greeter service.
package main

import (
	"bufio"
	"context"
	"fmt"
	"log"
	"os"
	"strings"
	"time"

	pb "../feedbackservice"
	"google.golang.org/grpc"
)

const (
	address     = "localhost:50051"
	defaultName = "world"
)

func main() {
	// Set up a connection to the server.
	conn, err := grpc.Dial(address, grpc.WithInsecure())
	if err != nil {
		log.Fatalf("did not connect: %v", err)
	}
	defer conn.Close()
	c := pb.NewFeedBackClient(conn)

	// Contact the server and print out its response.
	name := defaultName
	if len(os.Args) > 1 {
		name = os.Args[1]
	}
	ctx, cancel := context.WithTimeout(context.Background(), time.Second)
	defer cancel()

	reader := bufio.NewReader(os.Stdin)
	fmt.Println("Simple Shell")
	fmt.Println("---------------------")

	// fmt.Println("Please choose which part do you want to do ")
	// fmt.Println("1. Add passenger feedback")
	fmt.Print("-> ")
	text, _ := reader.ReadString('\n')
	// // convert CRLF to LF
	text = strings.Replace(text, "\n", "", -1)
	// switch text {
	// case "1":
	// 	fmt.Print("-> Input Feedback")
	// 	fmt.Print("-> BookingCode: ")
	// 	sBookingCode, _ := reader.ReadString('\n')
	// 	sBookingCode = strings.Replace(sBookingCode, "\n", "", -1)

	// 	fmt.Print("-> PID: ")
	// 	iPID, _ := reader.ReadString('\n')
	// 	iPID = strings.Replace(iPID, "\n", "", -1)

	// 	fmt.Print("-> Feedback: ")
	// 	sFeedBack, _ := reader.ReadString('\n')
	// 	sFeedBack = strings.Replace(sFeedBack, "\n", "", -1)

	// 	fmt.Printf("%v,%v,%v", sBookingCode, iPID, sFeedBack)

	// 	r2, err2 := c.AddPassengerFeedBack(ctx, &pb.PassengerFeedback{BookingCode: "1", PassengerID: 2, Feedback: "3"})
	// 	if err2 != nil {
	// 		log.Fatalf("could not greet: %v", err2)
	// 	}
	// 	log.Printf("Greeting: %v", r2)
	// }

	r, err := c.SayHello(ctx, &pb.HelloRequest{Name: name})
	if err != nil {
		log.Fatalf("could not greet: %v", err)
	}
	log.Printf("Greeting: %s", r.Message)

	r2, err2 := c.AddPassengerFeedBack(ctx, &pb.PassengerFeedback{BookingCode: "1", PassengerID: 2, Feedback: "3"})
	if err2 != nil {
		log.Fatalf("could not greet: %v", err2)
	}
	log.Printf("Greeting: %v", r2)
}
