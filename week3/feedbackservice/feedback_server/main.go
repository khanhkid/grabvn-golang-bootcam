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

//go:generate protoc -I ../helloworld --go_out=plugins=grpc:../helloworld ../helloworld/helloworld.proto

// Package main implements a server for Greeter service.
package main

import (
	"context"
	"log"
	"net"

	pb "../feedbackservice"
	"google.golang.org/grpc"
)

const (
	port = ":50051"
)

// server is used to implement helloworld.GreeterServer.
type server struct{}

// SayHello implements helloworld.GreeterServer
func (s *server) SayHello(ctx context.Context, in *pb.HelloRequest) (*pb.HelloReply, error) {
	log.Printf("Received: %v", in.Name)
	return &pb.HelloReply{Message: "Hello " + in.Name}, nil
}

func (s *server) AddPassengerFeedBack(ctx context.Context, in *pb.PassengerFeedback) (*pb.PassengerFeedbackReply, error) {
	log.Printf("Received: %v", in.BookingCode)
	resultFeedBack := &pb.PassengerFeedback{BookingCode: in.BookingCode, PassengerID: in.PassengerID, Feedback: in.Feedback}
	return &pb.PassengerFeedbackReply{Result: resultFeedBack}, nil
}

func (s *server) GetFeedBackByPassengerID(ctx context.Context, in *pb.GetFeedBackByPID) (*pb.GetFeedBackByPIDReply, error) {
	log.Printf("Received GetFeedBackByPassengerID")
	return &pb.GetFeedBackByPIDReply{BookingCode: "1", Feedback: "2"}, nil
}

func (s *server) GetFeedBackByBookingCode(ctx context.Context, in *pb.GetFeedBackByBookingID) (*pb.GetFeedBackByBookingIDReply, error) {
	log.Printf("Received GetFeedBackByPassengerID")
	return &pb.GetFeedBackByBookingIDReply{PassengerID: 1, Feedback: "2"}, nil
}
func (s *server) DeleteFeedBackByPassengerID(ctx context.Context, in *pb.DeleteFeedBackByPID) (*pb.DeleteFeedBackByReply, error) {
	log.Printf("Received GetFeedBackByPassengerID")
	return &pb.DeleteFeedBackByReply{Msg: "1", Code: 2}, nil
}

func main() {
	lis, err := net.Listen("tcp", port)
	if err != nil {
		log.Fatalf("failed to listen: %v", err)
	}
	s := grpc.NewServer()
	pb.RegisterFeedBackServer(s, &server{})
	if err := s.Serve(lis); err != nil {
		log.Fatalf("failed to serve: %v", err)
	}
}
