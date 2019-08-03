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
	"fmt"
	"log"
	"net"

	pb "../manageuser"
	"google.golang.org/grpc"
)

const (
	port = ":50051"
)

var ListUsers []pb.UserInfo

// server is used to implement helloworld.GreeterServer.
type server struct{}

// SayHello implements helloworld.GreeterServer
func (s *server) SayHello(ctx context.Context, in *pb.HelloRequest) (*pb.HelloReply, error) {
	log.Printf("Received: %v", in.Name)
	return &pb.HelloReply{Message: "Hello " + in.Name}, nil
}
func (s *server) InsertUser(ctx context.Context, in *pb.UserInfo) (*pb.UserInfo, error) {
	log.Printf("Received: %v", in.Name)
	ListUsers = append(ListUsers, pb.UserInfo{UserID: in.UserID, Name: in.Name})
	fmt.Println(ListUsers)
	return &pb.UserInfo{UserID: in.UserID, Name: in.Name}, nil
}
func (s *server) DeleteUser(ctx context.Context, in *pb.DeleteUserRequest) (*pb.ReturnMessage, error) {
	log.Printf("Received:")
	return &pb.ReturnMessage{Msg: "Hello "}, nil
}
func (s *server) ListUser(ctx context.Context, in *pb.ListUserRequest) (*pb.ListUserResponse, error) {
	return &pb.ListUserResponse{UserInfos: []*pb.UserInfo{}}, nil
}

func main() {
	ListUsers = append(ListUsers, pb.UserInfo{UserID: 1, Name: "User1"})
	ListUsers = append(ListUsers, pb.UserInfo{UserID: 2, Name: "User2"})
	ListUsers = append(ListUsers, pb.UserInfo{UserID: 3, Name: "User3"})

	lis, err := net.Listen("tcp", port)
	if err != nil {
		log.Fatalf("failed to listen: %v", err)
	}
	s := grpc.NewServer()
	pb.RegisterUserServer(s, &server{})
	if err := s.Serve(lis); err != nil {
		log.Fatalf("failed to serve: %v", err)
	}
}