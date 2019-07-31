// Code generated by protoc-gen-go. DO NOT EDIT.
// source: feedbackservice.proto

package feedbackservice

import (
	context "context"
	fmt "fmt"
	math "math"

	proto "github.com/golang/protobuf/proto"
	grpc "google.golang.org/grpc"
	codes "google.golang.org/grpc/codes"
	status "google.golang.org/grpc/status"
)

// Reference imports to suppress errors if they are not otherwise used.
var _ = proto.Marshal
var _ = fmt.Errorf
var _ = math.Inf

// This is a compile-time assertion to ensure that this generated file
// is compatible with the proto package it is being compiled against.
// A compilation error at this line likely means your copy of the
// proto package needs to be updated.
const _ = proto.ProtoPackageIsVersion3 // please upgrade the proto package

// The request message containing the user's name.
type HelloRequest struct {
	Name                 string   `protobuf:"bytes,1,opt,name=name,proto3" json:"name,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *HelloRequest) Reset()         { *m = HelloRequest{} }
func (m *HelloRequest) String() string { return proto.CompactTextString(m) }
func (*HelloRequest) ProtoMessage()    {}
func (*HelloRequest) Descriptor() ([]byte, []int) {
	return fileDescriptor_4b67b190e6a32652, []int{0}
}

func (m *HelloRequest) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_HelloRequest.Unmarshal(m, b)
}
func (m *HelloRequest) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_HelloRequest.Marshal(b, m, deterministic)
}
func (m *HelloRequest) XXX_Merge(src proto.Message) {
	xxx_messageInfo_HelloRequest.Merge(m, src)
}
func (m *HelloRequest) XXX_Size() int {
	return xxx_messageInfo_HelloRequest.Size(m)
}
func (m *HelloRequest) XXX_DiscardUnknown() {
	xxx_messageInfo_HelloRequest.DiscardUnknown(m)
}

var xxx_messageInfo_HelloRequest proto.InternalMessageInfo

func (m *HelloRequest) GetName() string {
	if m != nil {
		return m.Name
	}
	return ""
}

// The response message containing the greetings
type HelloReply struct {
	Message              string   `protobuf:"bytes,1,opt,name=message,proto3" json:"message,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *HelloReply) Reset()         { *m = HelloReply{} }
func (m *HelloReply) String() string { return proto.CompactTextString(m) }
func (*HelloReply) ProtoMessage()    {}
func (*HelloReply) Descriptor() ([]byte, []int) {
	return fileDescriptor_4b67b190e6a32652, []int{1}
}

func (m *HelloReply) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_HelloReply.Unmarshal(m, b)
}
func (m *HelloReply) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_HelloReply.Marshal(b, m, deterministic)
}
func (m *HelloReply) XXX_Merge(src proto.Message) {
	xxx_messageInfo_HelloReply.Merge(m, src)
}
func (m *HelloReply) XXX_Size() int {
	return xxx_messageInfo_HelloReply.Size(m)
}
func (m *HelloReply) XXX_DiscardUnknown() {
	xxx_messageInfo_HelloReply.DiscardUnknown(m)
}

var xxx_messageInfo_HelloReply proto.InternalMessageInfo

func (m *HelloReply) GetMessage() string {
	if m != nil {
		return m.Message
	}
	return ""
}

type PassengerFeedback struct {
	BookingCode          string   `protobuf:"bytes,1,opt,name=bookingCode,proto3" json:"bookingCode,omitempty"`
	PassengerID          int32    `protobuf:"varint,2,opt,name=passengerID,proto3" json:"passengerID,omitempty"`
	Feedback             string   `protobuf:"bytes,3,opt,name=feedback,proto3" json:"feedback,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *PassengerFeedback) Reset()         { *m = PassengerFeedback{} }
func (m *PassengerFeedback) String() string { return proto.CompactTextString(m) }
func (*PassengerFeedback) ProtoMessage()    {}
func (*PassengerFeedback) Descriptor() ([]byte, []int) {
	return fileDescriptor_4b67b190e6a32652, []int{2}
}

func (m *PassengerFeedback) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_PassengerFeedback.Unmarshal(m, b)
}
func (m *PassengerFeedback) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_PassengerFeedback.Marshal(b, m, deterministic)
}
func (m *PassengerFeedback) XXX_Merge(src proto.Message) {
	xxx_messageInfo_PassengerFeedback.Merge(m, src)
}
func (m *PassengerFeedback) XXX_Size() int {
	return xxx_messageInfo_PassengerFeedback.Size(m)
}
func (m *PassengerFeedback) XXX_DiscardUnknown() {
	xxx_messageInfo_PassengerFeedback.DiscardUnknown(m)
}

var xxx_messageInfo_PassengerFeedback proto.InternalMessageInfo

func (m *PassengerFeedback) GetBookingCode() string {
	if m != nil {
		return m.BookingCode
	}
	return ""
}

func (m *PassengerFeedback) GetPassengerID() int32 {
	if m != nil {
		return m.PassengerID
	}
	return 0
}

func (m *PassengerFeedback) GetFeedback() string {
	if m != nil {
		return m.Feedback
	}
	return ""
}

type PassengerFeedbackReply struct {
	Result               *PassengerFeedback `protobuf:"bytes,1,opt,name=result,proto3" json:"result,omitempty"`
	XXX_NoUnkeyedLiteral struct{}           `json:"-"`
	XXX_unrecognized     []byte             `json:"-"`
	XXX_sizecache        int32              `json:"-"`
}

func (m *PassengerFeedbackReply) Reset()         { *m = PassengerFeedbackReply{} }
func (m *PassengerFeedbackReply) String() string { return proto.CompactTextString(m) }
func (*PassengerFeedbackReply) ProtoMessage()    {}
func (*PassengerFeedbackReply) Descriptor() ([]byte, []int) {
	return fileDescriptor_4b67b190e6a32652, []int{3}
}

func (m *PassengerFeedbackReply) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_PassengerFeedbackReply.Unmarshal(m, b)
}
func (m *PassengerFeedbackReply) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_PassengerFeedbackReply.Marshal(b, m, deterministic)
}
func (m *PassengerFeedbackReply) XXX_Merge(src proto.Message) {
	xxx_messageInfo_PassengerFeedbackReply.Merge(m, src)
}
func (m *PassengerFeedbackReply) XXX_Size() int {
	return xxx_messageInfo_PassengerFeedbackReply.Size(m)
}
func (m *PassengerFeedbackReply) XXX_DiscardUnknown() {
	xxx_messageInfo_PassengerFeedbackReply.DiscardUnknown(m)
}

var xxx_messageInfo_PassengerFeedbackReply proto.InternalMessageInfo

func (m *PassengerFeedbackReply) GetResult() *PassengerFeedback {
	if m != nil {
		return m.Result
	}
	return nil
}

type GetFeedBackByPID struct {
	PassengerID          string   `protobuf:"bytes,1,opt,name=passengerID,proto3" json:"passengerID,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *GetFeedBackByPID) Reset()         { *m = GetFeedBackByPID{} }
func (m *GetFeedBackByPID) String() string { return proto.CompactTextString(m) }
func (*GetFeedBackByPID) ProtoMessage()    {}
func (*GetFeedBackByPID) Descriptor() ([]byte, []int) {
	return fileDescriptor_4b67b190e6a32652, []int{4}
}

func (m *GetFeedBackByPID) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_GetFeedBackByPID.Unmarshal(m, b)
}
func (m *GetFeedBackByPID) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_GetFeedBackByPID.Marshal(b, m, deterministic)
}
func (m *GetFeedBackByPID) XXX_Merge(src proto.Message) {
	xxx_messageInfo_GetFeedBackByPID.Merge(m, src)
}
func (m *GetFeedBackByPID) XXX_Size() int {
	return xxx_messageInfo_GetFeedBackByPID.Size(m)
}
func (m *GetFeedBackByPID) XXX_DiscardUnknown() {
	xxx_messageInfo_GetFeedBackByPID.DiscardUnknown(m)
}

var xxx_messageInfo_GetFeedBackByPID proto.InternalMessageInfo

func (m *GetFeedBackByPID) GetPassengerID() string {
	if m != nil {
		return m.PassengerID
	}
	return ""
}

type GetFeedBackByPIDReply struct {
	BookingCode          string   `protobuf:"bytes,1,opt,name=bookingCode,proto3" json:"bookingCode,omitempty"`
	Feedback             string   `protobuf:"bytes,2,opt,name=feedback,proto3" json:"feedback,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *GetFeedBackByPIDReply) Reset()         { *m = GetFeedBackByPIDReply{} }
func (m *GetFeedBackByPIDReply) String() string { return proto.CompactTextString(m) }
func (*GetFeedBackByPIDReply) ProtoMessage()    {}
func (*GetFeedBackByPIDReply) Descriptor() ([]byte, []int) {
	return fileDescriptor_4b67b190e6a32652, []int{5}
}

func (m *GetFeedBackByPIDReply) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_GetFeedBackByPIDReply.Unmarshal(m, b)
}
func (m *GetFeedBackByPIDReply) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_GetFeedBackByPIDReply.Marshal(b, m, deterministic)
}
func (m *GetFeedBackByPIDReply) XXX_Merge(src proto.Message) {
	xxx_messageInfo_GetFeedBackByPIDReply.Merge(m, src)
}
func (m *GetFeedBackByPIDReply) XXX_Size() int {
	return xxx_messageInfo_GetFeedBackByPIDReply.Size(m)
}
func (m *GetFeedBackByPIDReply) XXX_DiscardUnknown() {
	xxx_messageInfo_GetFeedBackByPIDReply.DiscardUnknown(m)
}

var xxx_messageInfo_GetFeedBackByPIDReply proto.InternalMessageInfo

func (m *GetFeedBackByPIDReply) GetBookingCode() string {
	if m != nil {
		return m.BookingCode
	}
	return ""
}

func (m *GetFeedBackByPIDReply) GetFeedback() string {
	if m != nil {
		return m.Feedback
	}
	return ""
}

type GetFeedBackByBookingID struct {
	BookingCode          string   `protobuf:"bytes,1,opt,name=bookingCode,proto3" json:"bookingCode,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *GetFeedBackByBookingID) Reset()         { *m = GetFeedBackByBookingID{} }
func (m *GetFeedBackByBookingID) String() string { return proto.CompactTextString(m) }
func (*GetFeedBackByBookingID) ProtoMessage()    {}
func (*GetFeedBackByBookingID) Descriptor() ([]byte, []int) {
	return fileDescriptor_4b67b190e6a32652, []int{6}
}

func (m *GetFeedBackByBookingID) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_GetFeedBackByBookingID.Unmarshal(m, b)
}
func (m *GetFeedBackByBookingID) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_GetFeedBackByBookingID.Marshal(b, m, deterministic)
}
func (m *GetFeedBackByBookingID) XXX_Merge(src proto.Message) {
	xxx_messageInfo_GetFeedBackByBookingID.Merge(m, src)
}
func (m *GetFeedBackByBookingID) XXX_Size() int {
	return xxx_messageInfo_GetFeedBackByBookingID.Size(m)
}
func (m *GetFeedBackByBookingID) XXX_DiscardUnknown() {
	xxx_messageInfo_GetFeedBackByBookingID.DiscardUnknown(m)
}

var xxx_messageInfo_GetFeedBackByBookingID proto.InternalMessageInfo

func (m *GetFeedBackByBookingID) GetBookingCode() string {
	if m != nil {
		return m.BookingCode
	}
	return ""
}

type GetFeedBackByBookingIDReply struct {
	PassengerID          string   `protobuf:"bytes,1,opt,name=passengerID,proto3" json:"passengerID,omitempty"`
	Feedback             string   `protobuf:"bytes,2,opt,name=feedback,proto3" json:"feedback,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *GetFeedBackByBookingIDReply) Reset()         { *m = GetFeedBackByBookingIDReply{} }
func (m *GetFeedBackByBookingIDReply) String() string { return proto.CompactTextString(m) }
func (*GetFeedBackByBookingIDReply) ProtoMessage()    {}
func (*GetFeedBackByBookingIDReply) Descriptor() ([]byte, []int) {
	return fileDescriptor_4b67b190e6a32652, []int{7}
}

func (m *GetFeedBackByBookingIDReply) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_GetFeedBackByBookingIDReply.Unmarshal(m, b)
}
func (m *GetFeedBackByBookingIDReply) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_GetFeedBackByBookingIDReply.Marshal(b, m, deterministic)
}
func (m *GetFeedBackByBookingIDReply) XXX_Merge(src proto.Message) {
	xxx_messageInfo_GetFeedBackByBookingIDReply.Merge(m, src)
}
func (m *GetFeedBackByBookingIDReply) XXX_Size() int {
	return xxx_messageInfo_GetFeedBackByBookingIDReply.Size(m)
}
func (m *GetFeedBackByBookingIDReply) XXX_DiscardUnknown() {
	xxx_messageInfo_GetFeedBackByBookingIDReply.DiscardUnknown(m)
}

var xxx_messageInfo_GetFeedBackByBookingIDReply proto.InternalMessageInfo

func (m *GetFeedBackByBookingIDReply) GetPassengerID() string {
	if m != nil {
		return m.PassengerID
	}
	return ""
}

func (m *GetFeedBackByBookingIDReply) GetFeedback() string {
	if m != nil {
		return m.Feedback
	}
	return ""
}

type DeleteFeedBackByPID struct {
	PassengerID          string   `protobuf:"bytes,1,opt,name=passengerID,proto3" json:"passengerID,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *DeleteFeedBackByPID) Reset()         { *m = DeleteFeedBackByPID{} }
func (m *DeleteFeedBackByPID) String() string { return proto.CompactTextString(m) }
func (*DeleteFeedBackByPID) ProtoMessage()    {}
func (*DeleteFeedBackByPID) Descriptor() ([]byte, []int) {
	return fileDescriptor_4b67b190e6a32652, []int{8}
}

func (m *DeleteFeedBackByPID) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_DeleteFeedBackByPID.Unmarshal(m, b)
}
func (m *DeleteFeedBackByPID) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_DeleteFeedBackByPID.Marshal(b, m, deterministic)
}
func (m *DeleteFeedBackByPID) XXX_Merge(src proto.Message) {
	xxx_messageInfo_DeleteFeedBackByPID.Merge(m, src)
}
func (m *DeleteFeedBackByPID) XXX_Size() int {
	return xxx_messageInfo_DeleteFeedBackByPID.Size(m)
}
func (m *DeleteFeedBackByPID) XXX_DiscardUnknown() {
	xxx_messageInfo_DeleteFeedBackByPID.DiscardUnknown(m)
}

var xxx_messageInfo_DeleteFeedBackByPID proto.InternalMessageInfo

func (m *DeleteFeedBackByPID) GetPassengerID() string {
	if m != nil {
		return m.PassengerID
	}
	return ""
}

type DeleteFeedBackByReply struct {
	Msg                  string   `protobuf:"bytes,1,opt,name=msg,proto3" json:"msg,omitempty"`
	Code                 string   `protobuf:"bytes,2,opt,name=code,proto3" json:"code,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *DeleteFeedBackByReply) Reset()         { *m = DeleteFeedBackByReply{} }
func (m *DeleteFeedBackByReply) String() string { return proto.CompactTextString(m) }
func (*DeleteFeedBackByReply) ProtoMessage()    {}
func (*DeleteFeedBackByReply) Descriptor() ([]byte, []int) {
	return fileDescriptor_4b67b190e6a32652, []int{9}
}

func (m *DeleteFeedBackByReply) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_DeleteFeedBackByReply.Unmarshal(m, b)
}
func (m *DeleteFeedBackByReply) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_DeleteFeedBackByReply.Marshal(b, m, deterministic)
}
func (m *DeleteFeedBackByReply) XXX_Merge(src proto.Message) {
	xxx_messageInfo_DeleteFeedBackByReply.Merge(m, src)
}
func (m *DeleteFeedBackByReply) XXX_Size() int {
	return xxx_messageInfo_DeleteFeedBackByReply.Size(m)
}
func (m *DeleteFeedBackByReply) XXX_DiscardUnknown() {
	xxx_messageInfo_DeleteFeedBackByReply.DiscardUnknown(m)
}

var xxx_messageInfo_DeleteFeedBackByReply proto.InternalMessageInfo

func (m *DeleteFeedBackByReply) GetMsg() string {
	if m != nil {
		return m.Msg
	}
	return ""
}

func (m *DeleteFeedBackByReply) GetCode() string {
	if m != nil {
		return m.Code
	}
	return ""
}

func init() {
	proto.RegisterType((*HelloRequest)(nil), "feedbackservice.HelloRequest")
	proto.RegisterType((*HelloReply)(nil), "feedbackservice.HelloReply")
	proto.RegisterType((*PassengerFeedback)(nil), "feedbackservice.PassengerFeedback")
	proto.RegisterType((*PassengerFeedbackReply)(nil), "feedbackservice.PassengerFeedbackReply")
	proto.RegisterType((*GetFeedBackByPID)(nil), "feedbackservice.GetFeedBackByPID")
	proto.RegisterType((*GetFeedBackByPIDReply)(nil), "feedbackservice.GetFeedBackByPIDReply")
	proto.RegisterType((*GetFeedBackByBookingID)(nil), "feedbackservice.GetFeedBackByBookingID")
	proto.RegisterType((*GetFeedBackByBookingIDReply)(nil), "feedbackservice.GetFeedBackByBookingIDReply")
	proto.RegisterType((*DeleteFeedBackByPID)(nil), "feedbackservice.DeleteFeedBackByPID")
	proto.RegisterType((*DeleteFeedBackByReply)(nil), "feedbackservice.DeleteFeedBackByReply")
}

func init() { proto.RegisterFile("feedbackservice.proto", fileDescriptor_4b67b190e6a32652) }

var fileDescriptor_4b67b190e6a32652 = []byte{
	// 434 bytes of a gzipped FileDescriptorProto
	0x1f, 0x8b, 0x08, 0x00, 0x00, 0x00, 0x00, 0x00, 0x02, 0xff, 0x94, 0x54, 0x4f, 0xef, 0xd2, 0x40,
	0x10, 0xa5, 0x80, 0x88, 0x83, 0x89, 0xb8, 0x02, 0x69, 0x4a, 0x4c, 0xea, 0xc6, 0x00, 0x07, 0xd3,
	0x03, 0x9a, 0x98, 0x90, 0x78, 0xb0, 0x12, 0x85, 0x5b, 0x83, 0x7a, 0xf2, 0x54, 0xda, 0xb1, 0x21,
	0xb4, 0xb4, 0x76, 0x8b, 0xb1, 0xdf, 0xc0, 0x8f, 0x6d, 0xfa, 0x67, 0xc9, 0xb2, 0xad, 0x85, 0xdf,
	0x6d, 0x77, 0x3a, 0x6f, 0xde, 0x7b, 0xd3, 0x97, 0x85, 0xf1, 0x4f, 0x44, 0x77, 0x6f, 0x3b, 0x47,
	0x86, 0xf1, 0xef, 0x83, 0x83, 0x46, 0x14, 0x87, 0x49, 0x48, 0x9e, 0x49, 0x65, 0x4a, 0xe1, 0xe9,
	0x06, 0x7d, 0x3f, 0xdc, 0xe1, 0xaf, 0x33, 0xb2, 0x84, 0x10, 0xe8, 0x9e, 0xec, 0x00, 0x55, 0x45,
	0x57, 0x16, 0x4f, 0x76, 0xf9, 0x99, 0xce, 0x00, 0xca, 0x9e, 0xc8, 0x4f, 0x89, 0x0a, 0x8f, 0x03,
	0x64, 0xcc, 0xf6, 0x78, 0x13, 0xbf, 0x52, 0x06, 0xcf, 0x2d, 0x9b, 0x31, 0x3c, 0x79, 0x18, 0x7f,
	0x2e, 0x79, 0x88, 0x0e, 0x83, 0x7d, 0x18, 0x1e, 0x0f, 0x27, 0xef, 0x53, 0xe8, 0x72, 0x88, 0x58,
	0xca, 0x3a, 0x22, 0x0e, 0xdb, 0xae, 0xd5, 0xb6, 0xae, 0x2c, 0x1e, 0xed, 0xc4, 0x12, 0xd1, 0xa0,
	0xcf, 0x75, 0xab, 0x9d, 0x7c, 0xc0, 0xe5, 0x4e, 0xbf, 0xc1, 0xa4, 0x42, 0x5a, 0x08, 0x5d, 0x41,
	0x2f, 0x46, 0x76, 0xf6, 0x93, 0x9c, 0x74, 0xb0, 0xa4, 0x86, 0xbc, 0x93, 0x2a, 0xb0, 0x44, 0xd0,
	0x77, 0x30, 0xfc, 0x82, 0x49, 0x56, 0x36, 0x6d, 0xe7, 0x68, 0xa6, 0xd6, 0x76, 0x2d, 0xeb, 0x2c,
	0x9d, 0x08, 0x25, 0xfa, 0x1d, 0xc6, 0x32, 0xaa, 0x90, 0x72, 0x7b, 0x09, 0xa2, 0xc5, 0xb6, 0x64,
	0x71, 0x05, 0x93, 0xab, 0xb1, 0x66, 0x81, 0x2b, 0x24, 0x35, 0xcf, 0xa5, 0x3f, 0x60, 0x5a, 0x8f,
	0xbd, 0x08, 0x6b, 0xf6, 0xd4, 0x28, 0xec, 0x3d, 0xbc, 0x58, 0xa3, 0x8f, 0x09, 0x3e, 0x74, 0x51,
	0x1f, 0x60, 0x2c, 0x03, 0x0b, 0x3d, 0x43, 0xe8, 0x04, 0xcc, 0x2b, 0x21, 0xd9, 0x31, 0x0b, 0xa4,
	0x93, 0x79, 0x2b, 0xb8, 0xf3, 0xf3, 0xf2, 0x6f, 0x17, 0xfa, 0x1c, 0x49, 0x36, 0xd0, 0xff, 0x6a,
	0xa7, 0x79, 0x40, 0xc9, 0xcb, 0xca, 0x2f, 0x16, 0xc3, 0xad, 0x4d, 0xff, 0xf7, 0x39, 0xf2, 0x53,
	0xda, 0x22, 0x08, 0xa3, 0x8f, 0xae, 0x7b, 0x15, 0x8a, 0x9c, 0xe1, 0x8e, 0xe0, 0x68, 0xf3, 0x3b,
	0xc2, 0x55, 0xd2, 0x78, 0xa0, 0x5e, 0xa7, 0x44, 0xd8, 0xf6, 0xab, 0xca, 0x18, 0x39, 0x50, 0xda,
	0xec, 0x66, 0x0b, 0x27, 0x62, 0x12, 0x91, 0x29, 0xe4, 0x6d, 0xde, 0x3c, 0xe5, 0x12, 0x13, 0xed,
	0xcd, 0x9d, 0x8d, 0x9c, 0x34, 0x80, 0x69, 0x25, 0x13, 0x82, 0xc1, 0xd7, 0x95, 0x71, 0x35, 0x09,
	0xaa, 0xf1, 0x58, 0x1b, 0x17, 0xda, 0x32, 0x57, 0xa0, 0x1f, 0x42, 0xc3, 0x8b, 0x23, 0xc7, 0xc0,
	0x3f, 0x76, 0x10, 0xf9, 0xc8, 0x64, 0xac, 0x39, 0x92, 0x0a, 0x56, 0xf6, 0x14, 0x5a, 0xca, 0xbe,
	0x97, 0xbf, 0x89, 0x6f, 0xff, 0x05, 0x00, 0x00, 0xff, 0xff, 0xdf, 0x16, 0x14, 0xa7, 0x2c, 0x05,
	0x00, 0x00,
}

// Reference imports to suppress errors if they are not otherwise used.
var _ context.Context
var _ grpc.ClientConn

// This is a compile-time assertion to ensure that this generated file
// is compatible with the grpc package it is being compiled against.
const _ = grpc.SupportPackageIsVersion4

// FeedBackClient is the client API for FeedBack service.
//
// For semantics around ctx use and closing/ending streaming RPCs, please refer to https://godoc.org/google.golang.org/grpc#ClientConn.NewStream.
type FeedBackClient interface {
	// Sends a greeting
	SayHello(ctx context.Context, in *HelloRequest, opts ...grpc.CallOption) (*HelloReply, error)
	AddPassengerFeedBack(ctx context.Context, in *PassengerFeedback, opts ...grpc.CallOption) (*PassengerFeedbackReply, error)
	GetFeedBackByPassengerID(ctx context.Context, in *GetFeedBackByPID, opts ...grpc.CallOption) (*GetFeedBackByPIDReply, error)
	GetFeedBackByBookingCode(ctx context.Context, in *GetFeedBackByBookingID, opts ...grpc.CallOption) (*GetFeedBackByBookingIDReply, error)
	DeleteFeedBackByPassengerID(ctx context.Context, in *DeleteFeedBackByPID, opts ...grpc.CallOption) (*DeleteFeedBackByReply, error)
}

type feedBackClient struct {
	cc *grpc.ClientConn
}

func NewFeedBackClient(cc *grpc.ClientConn) FeedBackClient {
	return &feedBackClient{cc}
}

func (c *feedBackClient) SayHello(ctx context.Context, in *HelloRequest, opts ...grpc.CallOption) (*HelloReply, error) {
	out := new(HelloReply)
	err := c.cc.Invoke(ctx, "/feedbackservice.FeedBack/SayHello", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *feedBackClient) AddPassengerFeedBack(ctx context.Context, in *PassengerFeedback, opts ...grpc.CallOption) (*PassengerFeedbackReply, error) {
	out := new(PassengerFeedbackReply)
	err := c.cc.Invoke(ctx, "/feedbackservice.FeedBack/AddPassengerFeedBack", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *feedBackClient) GetFeedBackByPassengerID(ctx context.Context, in *GetFeedBackByPID, opts ...grpc.CallOption) (*GetFeedBackByPIDReply, error) {
	out := new(GetFeedBackByPIDReply)
	err := c.cc.Invoke(ctx, "/feedbackservice.FeedBack/GetFeedBackByPassengerID", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *feedBackClient) GetFeedBackByBookingCode(ctx context.Context, in *GetFeedBackByBookingID, opts ...grpc.CallOption) (*GetFeedBackByBookingIDReply, error) {
	out := new(GetFeedBackByBookingIDReply)
	err := c.cc.Invoke(ctx, "/feedbackservice.FeedBack/GetFeedBackByBookingCode", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *feedBackClient) DeleteFeedBackByPassengerID(ctx context.Context, in *DeleteFeedBackByPID, opts ...grpc.CallOption) (*DeleteFeedBackByReply, error) {
	out := new(DeleteFeedBackByReply)
	err := c.cc.Invoke(ctx, "/feedbackservice.FeedBack/DeleteFeedBackByPassengerID", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

// FeedBackServer is the server API for FeedBack service.
type FeedBackServer interface {
	// Sends a greeting
	SayHello(context.Context, *HelloRequest) (*HelloReply, error)
	AddPassengerFeedBack(context.Context, *PassengerFeedback) (*PassengerFeedbackReply, error)
	GetFeedBackByPassengerID(context.Context, *GetFeedBackByPID) (*GetFeedBackByPIDReply, error)
	GetFeedBackByBookingCode(context.Context, *GetFeedBackByBookingID) (*GetFeedBackByBookingIDReply, error)
	DeleteFeedBackByPassengerID(context.Context, *DeleteFeedBackByPID) (*DeleteFeedBackByReply, error)
}

// UnimplementedFeedBackServer can be embedded to have forward compatible implementations.
type UnimplementedFeedBackServer struct {
}

func (*UnimplementedFeedBackServer) SayHello(ctx context.Context, req *HelloRequest) (*HelloReply, error) {
	return nil, status.Errorf(codes.Unimplemented, "method SayHello not implemented")
}
func (*UnimplementedFeedBackServer) AddPassengerFeedBack(ctx context.Context, req *PassengerFeedback) (*PassengerFeedbackReply, error) {
	return nil, status.Errorf(codes.Unimplemented, "method AddPassengerFeedBack not implemented")
}
func (*UnimplementedFeedBackServer) GetFeedBackByPassengerID(ctx context.Context, req *GetFeedBackByPID) (*GetFeedBackByPIDReply, error) {
	return nil, status.Errorf(codes.Unimplemented, "method GetFeedBackByPassengerID not implemented")
}
func (*UnimplementedFeedBackServer) GetFeedBackByBookingCode(ctx context.Context, req *GetFeedBackByBookingID) (*GetFeedBackByBookingIDReply, error) {
	return nil, status.Errorf(codes.Unimplemented, "method GetFeedBackByBookingCode not implemented")
}
func (*UnimplementedFeedBackServer) DeleteFeedBackByPassengerID(ctx context.Context, req *DeleteFeedBackByPID) (*DeleteFeedBackByReply, error) {
	return nil, status.Errorf(codes.Unimplemented, "method DeleteFeedBackByPassengerID not implemented")
}

func RegisterFeedBackServer(s *grpc.Server, srv FeedBackServer) {
	s.RegisterService(&_FeedBack_serviceDesc, srv)
}

func _FeedBack_SayHello_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(HelloRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(FeedBackServer).SayHello(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/feedbackservice.FeedBack/SayHello",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(FeedBackServer).SayHello(ctx, req.(*HelloRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _FeedBack_AddPassengerFeedBack_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(PassengerFeedback)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(FeedBackServer).AddPassengerFeedBack(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/feedbackservice.FeedBack/AddPassengerFeedBack",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(FeedBackServer).AddPassengerFeedBack(ctx, req.(*PassengerFeedback))
	}
	return interceptor(ctx, in, info, handler)
}

func _FeedBack_GetFeedBackByPassengerID_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(GetFeedBackByPID)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(FeedBackServer).GetFeedBackByPassengerID(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/feedbackservice.FeedBack/GetFeedBackByPassengerID",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(FeedBackServer).GetFeedBackByPassengerID(ctx, req.(*GetFeedBackByPID))
	}
	return interceptor(ctx, in, info, handler)
}

func _FeedBack_GetFeedBackByBookingCode_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(GetFeedBackByBookingID)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(FeedBackServer).GetFeedBackByBookingCode(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/feedbackservice.FeedBack/GetFeedBackByBookingCode",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(FeedBackServer).GetFeedBackByBookingCode(ctx, req.(*GetFeedBackByBookingID))
	}
	return interceptor(ctx, in, info, handler)
}

func _FeedBack_DeleteFeedBackByPassengerID_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(DeleteFeedBackByPID)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(FeedBackServer).DeleteFeedBackByPassengerID(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/feedbackservice.FeedBack/DeleteFeedBackByPassengerID",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(FeedBackServer).DeleteFeedBackByPassengerID(ctx, req.(*DeleteFeedBackByPID))
	}
	return interceptor(ctx, in, info, handler)
}

var _FeedBack_serviceDesc = grpc.ServiceDesc{
	ServiceName: "feedbackservice.FeedBack",
	HandlerType: (*FeedBackServer)(nil),
	Methods: []grpc.MethodDesc{
		{
			MethodName: "SayHello",
			Handler:    _FeedBack_SayHello_Handler,
		},
		{
			MethodName: "AddPassengerFeedBack",
			Handler:    _FeedBack_AddPassengerFeedBack_Handler,
		},
		{
			MethodName: "GetFeedBackByPassengerID",
			Handler:    _FeedBack_GetFeedBackByPassengerID_Handler,
		},
		{
			MethodName: "GetFeedBackByBookingCode",
			Handler:    _FeedBack_GetFeedBackByBookingCode_Handler,
		},
		{
			MethodName: "DeleteFeedBackByPassengerID",
			Handler:    _FeedBack_DeleteFeedBackByPassengerID_Handler,
		},
	},
	Streams:  []grpc.StreamDesc{},
	Metadata: "feedbackservice.proto",
}