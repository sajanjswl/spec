// Code generated by protoc-gen-go. DO NOT EDIT.
// source: github.com/container-object-storage-interface/spec/cosi.proto

package cosi

import (
	context "context"
	fmt "fmt"
	proto "github.com/golang/protobuf/proto"
	descriptor "github.com/golang/protobuf/protoc-gen-go/descriptor"
	_ "github.com/golang/protobuf/ptypes/timestamp"
	_ "github.com/golang/protobuf/ptypes/wrappers"
	grpc "google.golang.org/grpc"
	codes "google.golang.org/grpc/codes"
	status "google.golang.org/grpc/status"
	math "math"
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

type CreateBucketRequest struct {
	// The suggested name for the storage space. This field is REQUIRED.
	// It serves two purposes:
	// 1) Idempotency - This name is generated by the CO to achieve
	//    idempotency.  The Driver SHOULD ensure that multiple
	//    `CreateBucket` calls for the same name do not result in more
	//    than one piece of storage provisioned corresponding to that
	//    name. If a Driver is unable to enforce idempotency, the CO's
	//    error recovery logic could result in multiple (unused) volumes
	//    being provisioned.
	//    In the case of error, the CO MUST handle the gRPC error codes
	//    per the recovery behavior defined in the "CreateBucket Errors"
	//    section below.
	//    The CO is responsible for cleaning up buckets it provisioned
	//    that it no longer needs. If the CO is uncertain whether a bucket
	//    was provisioned or not when a `CreateBucket` call fails, the CO
	//    MAY call `CreateBucket` again, with the same name, to ensure the
	//    bucket exists and to retrieve the bucket's `bucket_id` (unless
	//    otherwise prohibited by "CreateBucket Errors").
	// 2) Suggested name - Some storage systems allow callers to specify
	//    an identifier by which to refer to the newly provisioned
	//    storage. If a storage system supports this, it can optionally
	//    use this name as the identifier for the new volume.
	// Any Unicode string that conforms to the length limit is allowed
	// except those containing the following banned characters:
	// U+0000-U+0008, U+000B, U+000C, U+000E-U+001F, U+007F-U+009F.
	// (These are control characters other than commonly used whitespace.)
	Name                 string   `protobuf:"bytes,1,opt,name=name,proto3" json:"name,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *CreateBucketRequest) Reset()         { *m = CreateBucketRequest{} }
func (m *CreateBucketRequest) String() string { return proto.CompactTextString(m) }
func (*CreateBucketRequest) ProtoMessage()    {}
func (*CreateBucketRequest) Descriptor() ([]byte, []int) {
	return fileDescriptor_a281b1e87a2def48, []int{0}
}

func (m *CreateBucketRequest) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_CreateBucketRequest.Unmarshal(m, b)
}
func (m *CreateBucketRequest) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_CreateBucketRequest.Marshal(b, m, deterministic)
}
func (m *CreateBucketRequest) XXX_Merge(src proto.Message) {
	xxx_messageInfo_CreateBucketRequest.Merge(m, src)
}
func (m *CreateBucketRequest) XXX_Size() int {
	return xxx_messageInfo_CreateBucketRequest.Size(m)
}
func (m *CreateBucketRequest) XXX_DiscardUnknown() {
	xxx_messageInfo_CreateBucketRequest.DiscardUnknown(m)
}

var xxx_messageInfo_CreateBucketRequest proto.InternalMessageInfo

func (m *CreateBucketRequest) GetName() string {
	if m != nil {
		return m.Name
	}
	return ""
}

type CreateBucketResponse struct {
	// Contains all attributes of the newly created bucket that are
	// relevant to the CO along with information required by the Driver
	// to uniquely identify the bucket. This field is REQUIRED.
	Bucket               *Bucket  `protobuf:"bytes,1,opt,name=bucket,proto3" json:"bucket,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *CreateBucketResponse) Reset()         { *m = CreateBucketResponse{} }
func (m *CreateBucketResponse) String() string { return proto.CompactTextString(m) }
func (*CreateBucketResponse) ProtoMessage()    {}
func (*CreateBucketResponse) Descriptor() ([]byte, []int) {
	return fileDescriptor_a281b1e87a2def48, []int{1}
}

func (m *CreateBucketResponse) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_CreateBucketResponse.Unmarshal(m, b)
}
func (m *CreateBucketResponse) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_CreateBucketResponse.Marshal(b, m, deterministic)
}
func (m *CreateBucketResponse) XXX_Merge(src proto.Message) {
	xxx_messageInfo_CreateBucketResponse.Merge(m, src)
}
func (m *CreateBucketResponse) XXX_Size() int {
	return xxx_messageInfo_CreateBucketResponse.Size(m)
}
func (m *CreateBucketResponse) XXX_DiscardUnknown() {
	xxx_messageInfo_CreateBucketResponse.DiscardUnknown(m)
}

var xxx_messageInfo_CreateBucketResponse proto.InternalMessageInfo

func (m *CreateBucketResponse) GetBucket() *Bucket {
	if m != nil {
		return m.Bucket
	}
	return nil
}

type Bucket struct {
	// The identifier for this bucket, generated by the driver.
	// This field is REQUIRED.
	// This field MUST contain enough information to uniquely identify
	// this specific bucket vs all other buckets supported by this driver.
	// This field SHALL be used by the CO in subsequent calls to refer to
	// this bucket.
	// The SP is NOT responsible for global uniqueness of bucket_id across
	// multiple SPs.
	BucketId             string   `protobuf:"bytes,1,opt,name=bucket_id,json=bucketId,proto3" json:"bucket_id,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *Bucket) Reset()         { *m = Bucket{} }
func (m *Bucket) String() string { return proto.CompactTextString(m) }
func (*Bucket) ProtoMessage()    {}
func (*Bucket) Descriptor() ([]byte, []int) {
	return fileDescriptor_a281b1e87a2def48, []int{2}
}

func (m *Bucket) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_Bucket.Unmarshal(m, b)
}
func (m *Bucket) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_Bucket.Marshal(b, m, deterministic)
}
func (m *Bucket) XXX_Merge(src proto.Message) {
	xxx_messageInfo_Bucket.Merge(m, src)
}
func (m *Bucket) XXX_Size() int {
	return xxx_messageInfo_Bucket.Size(m)
}
func (m *Bucket) XXX_DiscardUnknown() {
	xxx_messageInfo_Bucket.DiscardUnknown(m)
}

var xxx_messageInfo_Bucket proto.InternalMessageInfo

func (m *Bucket) GetBucketId() string {
	if m != nil {
		return m.BucketId
	}
	return ""
}

var E_AlphaEnum = &proto.ExtensionDesc{
	ExtendedType:  (*descriptor.EnumOptions)(nil),
	ExtensionType: (*bool)(nil),
	Field:         1060,
	Name:          "cosi.v1.alpha_enum",
	Tag:           "varint,1060,opt,name=alpha_enum",
	Filename:      "github.com/container-object-storage-interface/spec/cosi.proto",
}

var E_AlphaEnumValue = &proto.ExtensionDesc{
	ExtendedType:  (*descriptor.EnumValueOptions)(nil),
	ExtensionType: (*bool)(nil),
	Field:         1060,
	Name:          "cosi.v1.alpha_enum_value",
	Tag:           "varint,1060,opt,name=alpha_enum_value",
	Filename:      "github.com/container-object-storage-interface/spec/cosi.proto",
}

var E_CosiSecret = &proto.ExtensionDesc{
	ExtendedType:  (*descriptor.FieldOptions)(nil),
	ExtensionType: (*bool)(nil),
	Field:         1059,
	Name:          "cosi.v1.cosi_secret",
	Tag:           "varint,1059,opt,name=cosi_secret",
	Filename:      "github.com/container-object-storage-interface/spec/cosi.proto",
}

var E_AlphaField = &proto.ExtensionDesc{
	ExtendedType:  (*descriptor.FieldOptions)(nil),
	ExtensionType: (*bool)(nil),
	Field:         1060,
	Name:          "cosi.v1.alpha_field",
	Tag:           "varint,1060,opt,name=alpha_field",
	Filename:      "github.com/container-object-storage-interface/spec/cosi.proto",
}

var E_AlphaMessage = &proto.ExtensionDesc{
	ExtendedType:  (*descriptor.MessageOptions)(nil),
	ExtensionType: (*bool)(nil),
	Field:         1060,
	Name:          "cosi.v1.alpha_message",
	Tag:           "varint,1060,opt,name=alpha_message",
	Filename:      "github.com/container-object-storage-interface/spec/cosi.proto",
}

var E_AlphaMethod = &proto.ExtensionDesc{
	ExtendedType:  (*descriptor.MethodOptions)(nil),
	ExtensionType: (*bool)(nil),
	Field:         1060,
	Name:          "cosi.v1.alpha_method",
	Tag:           "varint,1060,opt,name=alpha_method",
	Filename:      "github.com/container-object-storage-interface/spec/cosi.proto",
}

var E_AlphaService = &proto.ExtensionDesc{
	ExtendedType:  (*descriptor.ServiceOptions)(nil),
	ExtensionType: (*bool)(nil),
	Field:         1060,
	Name:          "cosi.v1.alpha_service",
	Tag:           "varint,1060,opt,name=alpha_service",
	Filename:      "github.com/container-object-storage-interface/spec/cosi.proto",
}

func init() {
	proto.RegisterType((*CreateBucketRequest)(nil), "cosi.v1.CreateBucketRequest")
	proto.RegisterType((*CreateBucketResponse)(nil), "cosi.v1.CreateBucketResponse")
	proto.RegisterType((*Bucket)(nil), "cosi.v1.Bucket")
	proto.RegisterExtension(E_AlphaEnum)
	proto.RegisterExtension(E_AlphaEnumValue)
	proto.RegisterExtension(E_CosiSecret)
	proto.RegisterExtension(E_AlphaField)
	proto.RegisterExtension(E_AlphaMessage)
	proto.RegisterExtension(E_AlphaMethod)
	proto.RegisterExtension(E_AlphaService)
}

func init() {
	proto.RegisterFile("github.com/container-object-storage-interface/spec/cosi.proto", fileDescriptor_a281b1e87a2def48)
}

var fileDescriptor_a281b1e87a2def48 = []byte{
	// 439 bytes of a gzipped FileDescriptorProto
	0x1f, 0x8b, 0x08, 0x00, 0x00, 0x00, 0x00, 0x00, 0x02, 0xff, 0x84, 0x90, 0x5d, 0x6b, 0x13, 0x41,
	0x14, 0x86, 0x0d, 0x94, 0x98, 0x9c, 0xc4, 0x0f, 0x46, 0x2f, 0x4a, 0xec, 0x47, 0x0c, 0x88, 0xf5,
	0x22, 0xbb, 0x58, 0xef, 0x16, 0xa5, 0x90, 0x5a, 0xa1, 0x60, 0x29, 0x24, 0xe0, 0x85, 0x5e, 0x84,
	0xd9, 0xd9, 0x93, 0xcd, 0x68, 0x76, 0x66, 0x9c, 0x99, 0x8d, 0x3f, 0x48, 0x7f, 0xa8, 0xcc, 0x47,
	0x5a, 0xdd, 0xae, 0x78, 0xb7, 0x9c, 0xe7, 0x3d, 0xcf, 0xbc, 0x7b, 0xe0, 0x5d, 0xc9, 0xed, 0xba,
	0xce, 0x13, 0x26, 0xab, 0x94, 0x49, 0x61, 0x29, 0x17, 0xa8, 0xa7, 0x32, 0xff, 0x8a, 0xcc, 0x4e,
	0x8d, 0x95, 0x9a, 0x96, 0x38, 0xe5, 0xc2, 0xa2, 0x5e, 0x51, 0x86, 0xa9, 0x51, 0xc8, 0x52, 0x26,
	0x0d, 0x4f, 0x94, 0x96, 0x56, 0x92, 0xfb, 0xfe, 0x7b, 0xfb, 0x7a, 0x34, 0x2e, 0xa5, 0x2c, 0x37,
	0x98, 0xfa, 0x71, 0x5e, 0xaf, 0xd2, 0x02, 0x0d, 0xd3, 0x5c, 0x59, 0xa9, 0x43, 0x74, 0x74, 0xdc,
	0x4c, 0x58, 0x5e, 0xa1, 0xb1, 0xb4, 0x52, 0x31, 0x70, 0xd4, 0x0c, 0xfc, 0xd0, 0x54, 0x29, 0xd4,
	0x26, 0xf0, 0xc9, 0x2b, 0x78, 0x72, 0xae, 0x91, 0x5a, 0x9c, 0xd5, 0xec, 0x1b, 0xda, 0x39, 0x7e,
	0xaf, 0xd1, 0x58, 0x42, 0x60, 0x4f, 0xd0, 0x0a, 0xf7, 0x3b, 0xe3, 0xce, 0x49, 0x7f, 0xee, 0xbf,
	0x27, 0x67, 0xf0, 0xf4, 0xef, 0xa8, 0x51, 0x52, 0x18, 0x24, 0x2f, 0xa1, 0x9b, 0xfb, 0x89, 0x4f,
	0x0f, 0x4e, 0x1f, 0x25, 0xb1, 0x7f, 0x12, 0x83, 0x11, 0x4f, 0x5e, 0x40, 0x37, 0x4c, 0xc8, 0x33,
	0xe8, 0x87, 0xd9, 0x92, 0x17, 0xf1, 0x8d, 0x5e, 0x18, 0x5c, 0x16, 0xa7, 0x5f, 0x00, 0xce, 0xaf,
	0x17, 0x97, 0xef, 0x35, 0xdf, 0xa2, 0x26, 0x57, 0x30, 0xfc, 0xf3, 0x55, 0x72, 0x70, 0x63, 0x6f,
	0xe9, 0x3d, 0x3a, 0xfc, 0x07, 0x0d, 0x55, 0x27, 0xf7, 0xb2, 0xb7, 0x00, 0x74, 0xa3, 0xd6, 0x74,
	0x89, 0xa2, 0xae, 0xc8, 0x41, 0x12, 0xce, 0x93, 0xec, 0xce, 0x93, 0x5c, 0x88, 0xba, 0xba, 0x56,
	0x96, 0x4b, 0x61, 0xf6, 0x7f, 0xf5, 0xc6, 0x9d, 0x93, 0xde, 0xbc, 0xef, 0x17, 0x1c, 0xc8, 0x3e,
	0xc2, 0xe3, 0xdb, 0xed, 0xe5, 0x96, 0x6e, 0x6a, 0x24, 0xcf, 0x5b, 0x1d, 0x9f, 0x1c, 0x6b, 0x88,
	0x1e, 0xde, 0x88, 0x3c, 0xcd, 0xce, 0x60, 0xe0, 0xda, 0x2e, 0x0d, 0x32, 0x8d, 0x96, 0x1c, 0xde,
	0x11, 0x7d, 0xe0, 0xb8, 0x29, 0x76, 0x92, 0x9f, 0x41, 0x02, 0x6e, 0x65, 0xe1, 0x37, 0x9c, 0x20,
	0xd4, 0x59, 0xb9, 0xe0, 0xff, 0x04, 0xb1, 0x45, 0xf8, 0x7f, 0x4f, 0xb2, 0x0b, 0x78, 0x10, 0x04,
	0x15, 0x1a, 0x43, 0x4b, 0x24, 0xc7, 0x77, 0x14, 0x57, 0x81, 0x34, 0x24, 0x43, 0xbf, 0x16, 0x59,
	0x36, 0x83, 0xe1, 0x4e, 0x63, 0xd7, 0xb2, 0x20, 0x47, 0x2d, 0x16, 0x07, 0x1a, 0x92, 0x41, 0x94,
	0x38, 0x74, 0x5b, 0xc5, 0xa0, 0xde, 0x72, 0xd6, 0x56, 0x65, 0x11, 0x48, 0x6b, 0x95, 0xc8, 0x66,
	0xdd, 0xcf, 0x7b, 0xee, 0x40, 0x79, 0xd7, 0x6f, 0xbd, 0xf9, 0x1d, 0x00, 0x00, 0xff, 0xff, 0xf8,
	0xf7, 0x09, 0x5b, 0x8b, 0x03, 0x00, 0x00,
}

// Reference imports to suppress errors if they are not otherwise used.
var _ context.Context
var _ grpc.ClientConn

// This is a compile-time assertion to ensure that this generated file
// is compatible with the grpc package it is being compiled against.
const _ = grpc.SupportPackageIsVersion4

// COSIDriverClient is the client API for COSIDriver service.
//
// For semantics around ctx use and closing/ending streaming RPCs, please refer to https://godoc.org/google.golang.org/grpc#ClientConn.NewStream.
type COSIDriverClient interface {
	CreateBucket(ctx context.Context, in *CreateBucketRequest, opts ...grpc.CallOption) (*CreateBucketResponse, error)
}

type cOSIDriverClient struct {
	cc *grpc.ClientConn
}

func NewCOSIDriverClient(cc *grpc.ClientConn) COSIDriverClient {
	return &cOSIDriverClient{cc}
}

func (c *cOSIDriverClient) CreateBucket(ctx context.Context, in *CreateBucketRequest, opts ...grpc.CallOption) (*CreateBucketResponse, error) {
	out := new(CreateBucketResponse)
	err := c.cc.Invoke(ctx, "/cosi.v1.COSIDriver/CreateBucket", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

// COSIDriverServer is the server API for COSIDriver service.
type COSIDriverServer interface {
	CreateBucket(context.Context, *CreateBucketRequest) (*CreateBucketResponse, error)
}

// UnimplementedCOSIDriverServer can be embedded to have forward compatible implementations.
type UnimplementedCOSIDriverServer struct {
}

func (*UnimplementedCOSIDriverServer) CreateBucket(ctx context.Context, req *CreateBucketRequest) (*CreateBucketResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method CreateBucket not implemented")
}

func RegisterCOSIDriverServer(s *grpc.Server, srv COSIDriverServer) {
	s.RegisterService(&_COSIDriver_serviceDesc, srv)
}

func _COSIDriver_CreateBucket_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(CreateBucketRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(COSIDriverServer).CreateBucket(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/cosi.v1.COSIDriver/CreateBucket",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(COSIDriverServer).CreateBucket(ctx, req.(*CreateBucketRequest))
	}
	return interceptor(ctx, in, info, handler)
}

var _COSIDriver_serviceDesc = grpc.ServiceDesc{
	ServiceName: "cosi.v1.COSIDriver",
	HandlerType: (*COSIDriverServer)(nil),
	Methods: []grpc.MethodDesc{
		{
			MethodName: "CreateBucket",
			Handler:    _COSIDriver_CreateBucket_Handler,
		},
	},
	Streams:  []grpc.StreamDesc{},
	Metadata: "github.com/container-object-storage-interface/spec/cosi.proto",
}