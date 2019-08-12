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
	"flag"
	"fmt"
	"log"
	"net"
	"strconv"

	"google.golang.org/grpc/codes"

	pb "../feedbackservice"
	"google.golang.org/grpc"
	"google.golang.org/grpc/status"
)

var (
	echoEndpoint = flag.String("echo_endpoint", "localhost:50051", "endpoint of YourService")
)

const (
	port = ":50051"
)

var feedBackData []pb.PassengerFeedback

// server is used to implement helloworld.GreeterServer.
type server struct{}

// SayHello implements helloworld.GreeterServer
func (s *server) SayHello(ctx context.Context, in *pb.HelloRequest) (*pb.HelloReply, error) {
	log.Printf("Received: %v", in.Name)
	return &pb.HelloReply{Message: "Hello " + in.Name}, nil
}

func (s *server) AddPassengerFeedBack(ctx context.Context, in *pb.PassengerFeedback) (*pb.PassengerFeedbackReply, error) {
	fmt.Println("Received Add New Passenger FeedBack")
	//Check condition of assignment: 1 booking has only 1 feedback
	resultFeedBack := &pb.PassengerFeedback{BookingCode: in.BookingCode, PassengerID: in.PassengerID, Feedback: in.Feedback}
	for _, feedBack := range feedBackData {
		if feedBack.BookingCode == in.BookingCode {
			return &pb.PassengerFeedbackReply{Result: resultFeedBack}, status.Error(codes.AlreadyExists, "1 booking has only 1 feedback\nDuplicate feed for Booking Code:"+feedBack.BookingCode)
		}
	}
	// Insert to database
	PassegerFeedBack := pb.PassengerFeedback{BookingCode: in.BookingCode, PassengerID: in.PassengerID, Feedback: in.Feedback}
	feedBackData = append(feedBackData, PassegerFeedBack)
	fmt.Println("Currently FeedBack")
	fmt.Println(feedBackData)
	return &pb.PassengerFeedbackReply{Result: resultFeedBack}, nil
}

func (s *server) GetFeedBackByPassengerID(ctx context.Context, in *pb.GetFeedBackByPID) (*pb.GetFeedBackByPIDReply, error) {
	fmt.Println("Received GetFeedBackByPassengerID")
	iPID := in.PassengerID
	resultFeedBack := make([]*pb.PassengerFeedback, 0)
	for _, feed := range feedBackData {
		fmt.Println(feed.PassengerID)
		if iPID == feed.PassengerID {
			resultFeedBack = append(resultFeedBack, &pb.PassengerFeedback{BookingCode: feed.BookingCode, PassengerID: feed.PassengerID, Feedback: feed.Feedback})
		}
	}
	fmt.Println("Response GetFeedBackByPassengerID")
	fmt.Println(resultFeedBack)
	return &pb.GetFeedBackByPIDReply{Result: resultFeedBack}, nil
}

func (s *server) GetFeedBackByBookingCode(ctx context.Context, in *pb.GetFeedBackByBookingID) (*pb.GetFeedBackByBookingIDReply, error) {
	fmt.Println("Response GetFeedBackByBookingCode")
	sBookingCode := in.BookingCode
	resultFeedBack := make([]*pb.PassengerFeedback, 0)
	for _, feed := range feedBackData {
		fmt.Println(feed.PassengerID)
		if sBookingCode == feed.BookingCode {
			resultFeedBack = append(resultFeedBack, &pb.PassengerFeedback{BookingCode: feed.BookingCode, PassengerID: feed.PassengerID, Feedback: feed.Feedback})
		}
	}
	log.Printf("Response GetFeedBackByPassengerID")
	fmt.Println(resultFeedBack)
	return &pb.GetFeedBackByBookingIDReply{Result: resultFeedBack}, nil
}
func (s *server) DeleteFeedBackByPassengerID(ctx context.Context, in *pb.DeleteFeedBackByPID) (*pb.DeleteFeedBackByReply, error) {
	log.Printf("Received DeleteFeedBackByPassengerID")
	iPID := in.PassengerID

	countFeedBack := len(feedBackData)
	copiedArray := make([]pb.PassengerFeedback, 0)
	countFoundMatched := 0
	for index := 0; index < countFeedBack; index++ {
		if iPID != feedBackData[index].PassengerID {
			copiedArray = append(copiedArray, feedBackData[index])
		} else {
			countFoundMatched++
		}
	}
	feedBackData = copiedArray
	log.Printf("Response DeleteFeedBackByPassengerID")
	return &pb.DeleteFeedBackByReply{Msg: "Success Removed:" + strconv.FormatInt(int64(countFoundMatched), 10), Code: 200}, nil
}

func main() {
	// Create default database
	feedBackData = []pb.PassengerFeedback{
		{BookingCode: "1", PassengerID: 2, Feedback: "This is feedback1"},
		{BookingCode: "2", PassengerID: 2, Feedback: "This is feedback2"},
		{BookingCode: "3", PassengerID: 2, Feedback: "This is feedback3"},
	}

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
