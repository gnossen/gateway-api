// Copyright 2024 The Kubernetes Authors.
//
// Licensed under the Apache License, Version 2.0 (the "License");
// you may not use this file except in compliance with the License.
// You may obtain a copy of the License at
//
//     http://www.apache.org/licenses/LICENSE-2.0
//
// Unless required by applicable law or agreed to in writing, software
// distributed under the License is distributed on an "AS IS" BASIS,
// WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND, either express or implied.
// See the License for the specific language governing permissions and
// limitations under the License.

// Code generated by protoc-gen-go. DO NOT EDIT.
// versions:
// 	protoc-gen-go v1.28.1
// 	protoc        v4.22.2
// source: grpcecho.proto

package grpcechoserver

import (
	protoreflect "google.golang.org/protobuf/reflect/protoreflect"
	protoimpl "google.golang.org/protobuf/runtime/protoimpl"
	reflect "reflect"
	sync "sync"
)

const (
	// Verify that this generated code is sufficiently up-to-date.
	_ = protoimpl.EnforceVersion(20 - protoimpl.MinVersion)
	// Verify that runtime/protoimpl is sufficiently up-to-date.
	_ = protoimpl.EnforceVersion(protoimpl.MaxVersion - 20)
)

type Header struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Key   string `protobuf:"bytes,1,opt,name=key,proto3" json:"key,omitempty"`
	Value string `protobuf:"bytes,2,opt,name=value,proto3" json:"value,omitempty"`
}

func (x *Header) Reset() {
	*x = Header{}
	if protoimpl.UnsafeEnabled {
		mi := &file_grpcecho_proto_msgTypes[0]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *Header) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*Header) ProtoMessage() {}

func (x *Header) ProtoReflect() protoreflect.Message {
	mi := &file_grpcecho_proto_msgTypes[0]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use Header.ProtoReflect.Descriptor instead.
func (*Header) Descriptor() ([]byte, []int) {
	return file_grpcecho_proto_rawDescGZIP(), []int{0}
}

func (x *Header) GetKey() string {
	if x != nil {
		return x.Key
	}
	return ""
}

func (x *Header) GetValue() string {
	if x != nil {
		return x.Value
	}
	return ""
}

type Context struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	// The Kubernetes namespace in which this server is running. Populated by the
	// NAMESPACE environment variable.
	Namespace string `protobuf:"bytes,1,opt,name=namespace,proto3" json:"namespace,omitempty"`
	// The name of the ingress controller under test. Populated by the INGRESS_NAME
	// environment variable.
	Ingress string `protobuf:"bytes,2,opt,name=ingress,proto3" json:"ingress,omitempty"`
	// The name service cannot be used here since it is a reserved word. Populated by the
	// SERVICE_NAME environment variable.
	ServiceName string `protobuf:"bytes,3,opt,name=service_name,json=serviceName,proto3" json:"service_name,omitempty"`
	// The name of the pod in which this server is running. Populated by the POD_NAME
	// environment variable.
	Pod string `protobuf:"bytes,4,opt,name=pod,proto3" json:"pod,omitempty"`
}

func (x *Context) Reset() {
	*x = Context{}
	if protoimpl.UnsafeEnabled {
		mi := &file_grpcecho_proto_msgTypes[1]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *Context) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*Context) ProtoMessage() {}

func (x *Context) ProtoReflect() protoreflect.Message {
	mi := &file_grpcecho_proto_msgTypes[1]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use Context.ProtoReflect.Descriptor instead.
func (*Context) Descriptor() ([]byte, []int) {
	return file_grpcecho_proto_rawDescGZIP(), []int{1}
}

func (x *Context) GetNamespace() string {
	if x != nil {
		return x.Namespace
	}
	return ""
}

func (x *Context) GetIngress() string {
	if x != nil {
		return x.Ingress
	}
	return ""
}

func (x *Context) GetServiceName() string {
	if x != nil {
		return x.ServiceName
	}
	return ""
}

func (x *Context) GetPod() string {
	if x != nil {
		return x.Pod
	}
	return ""
}

type TLSAssertions struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	// The TLS version used by the connection, e.g. "TLSv1.3"
	Version string `protobuf:"bytes,1,opt,name=version,proto3" json:"version,omitempty"`
	// The negotatiated protocol.
	NegotiatedProtocol string `protobuf:"bytes,2,opt,name=negotiated_protocol,json=negotiatedProtocol,proto3" json:"negotiated_protocol,omitempty"`
	// The server name indication extension sent by the client.
	ServerName string `protobuf:"bytes,3,opt,name=server_name,json=serverName,proto3" json:"server_name,omitempty"`
	// The cipher suite negotatiated for the connection, e.g. "TLS_EDCHE_ECDSA_WITH_AES_128_GCM_SHA256"
	CipherSuite string `protobuf:"bytes,4,opt,name=cipher_suite,json=cipherSuite,proto3" json:"cipher_suite,omitempty"`
	// The parsed certificates sent by the peer, in the order in which they were sent.
	PeerCertificates []string `protobuf:"bytes,5,rep,name=peer_certificates,json=peerCertificates,proto3" json:"peer_certificates,omitempty"`
}

func (x *TLSAssertions) Reset() {
	*x = TLSAssertions{}
	if protoimpl.UnsafeEnabled {
		mi := &file_grpcecho_proto_msgTypes[2]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *TLSAssertions) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*TLSAssertions) ProtoMessage() {}

func (x *TLSAssertions) ProtoReflect() protoreflect.Message {
	mi := &file_grpcecho_proto_msgTypes[2]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use TLSAssertions.ProtoReflect.Descriptor instead.
func (*TLSAssertions) Descriptor() ([]byte, []int) {
	return file_grpcecho_proto_rawDescGZIP(), []int{2}
}

func (x *TLSAssertions) GetVersion() string {
	if x != nil {
		return x.Version
	}
	return ""
}

func (x *TLSAssertions) GetNegotiatedProtocol() string {
	if x != nil {
		return x.NegotiatedProtocol
	}
	return ""
}

func (x *TLSAssertions) GetServerName() string {
	if x != nil {
		return x.ServerName
	}
	return ""
}

func (x *TLSAssertions) GetCipherSuite() string {
	if x != nil {
		return x.CipherSuite
	}
	return ""
}

func (x *TLSAssertions) GetPeerCertificates() []string {
	if x != nil {
		return x.PeerCertificates
	}
	return nil
}

type Assertions struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	// The fully qualified method of the current RPC, e.g.
	// "/gateway_api_conformance.grpc_echo.grpcecho.GrpcEcho/Echo"
	FullyQualifiedMethod string `protobuf:"bytes,1,opt,name=fully_qualified_method,json=fullyQualifiedMethod,proto3" json:"fully_qualified_method,omitempty"`
	// The headers present in the request.
	Headers []*Header `protobuf:"bytes,2,rep,name=headers,proto3" json:"headers,omitempty"`
	// The :authority pseudo-header of the request.
	Authority string `protobuf:"bytes,3,opt,name=authority,proto3" json:"authority,omitempty"`
	// Information associated with the conformance server deployment.
	Context *Context `protobuf:"bytes,4,opt,name=context,proto3" json:"context,omitempty"`
	// Information related to the TLS connection between the client and the server.
	TlsAssertions *TLSAssertions `protobuf:"bytes,5,opt,name=tls_assertions,json=tlsAssertions,proto3" json:"tls_assertions,omitempty"`
}

func (x *Assertions) Reset() {
	*x = Assertions{}
	if protoimpl.UnsafeEnabled {
		mi := &file_grpcecho_proto_msgTypes[3]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *Assertions) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*Assertions) ProtoMessage() {}

func (x *Assertions) ProtoReflect() protoreflect.Message {
	mi := &file_grpcecho_proto_msgTypes[3]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use Assertions.ProtoReflect.Descriptor instead.
func (*Assertions) Descriptor() ([]byte, []int) {
	return file_grpcecho_proto_rawDescGZIP(), []int{3}
}

func (x *Assertions) GetFullyQualifiedMethod() string {
	if x != nil {
		return x.FullyQualifiedMethod
	}
	return ""
}

func (x *Assertions) GetHeaders() []*Header {
	if x != nil {
		return x.Headers
	}
	return nil
}

func (x *Assertions) GetAuthority() string {
	if x != nil {
		return x.Authority
	}
	return ""
}

func (x *Assertions) GetContext() *Context {
	if x != nil {
		return x.Context
	}
	return nil
}

func (x *Assertions) GetTlsAssertions() *TLSAssertions {
	if x != nil {
		return x.TlsAssertions
	}
	return nil
}

type EchoRequest struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields
}

func (x *EchoRequest) Reset() {
	*x = EchoRequest{}
	if protoimpl.UnsafeEnabled {
		mi := &file_grpcecho_proto_msgTypes[4]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *EchoRequest) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*EchoRequest) ProtoMessage() {}

func (x *EchoRequest) ProtoReflect() protoreflect.Message {
	mi := &file_grpcecho_proto_msgTypes[4]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use EchoRequest.ProtoReflect.Descriptor instead.
func (*EchoRequest) Descriptor() ([]byte, []int) {
	return file_grpcecho_proto_rawDescGZIP(), []int{4}
}

type EchoResponse struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Assertions *Assertions  `protobuf:"bytes,1,opt,name=assertions,proto3" json:"assertions,omitempty"`
	Request    *EchoRequest `protobuf:"bytes,2,opt,name=request,proto3" json:"request,omitempty"`
}

func (x *EchoResponse) Reset() {
	*x = EchoResponse{}
	if protoimpl.UnsafeEnabled {
		mi := &file_grpcecho_proto_msgTypes[5]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *EchoResponse) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*EchoResponse) ProtoMessage() {}

func (x *EchoResponse) ProtoReflect() protoreflect.Message {
	mi := &file_grpcecho_proto_msgTypes[5]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use EchoResponse.ProtoReflect.Descriptor instead.
func (*EchoResponse) Descriptor() ([]byte, []int) {
	return file_grpcecho_proto_rawDescGZIP(), []int{5}
}

func (x *EchoResponse) GetAssertions() *Assertions {
	if x != nil {
		return x.Assertions
	}
	return nil
}

func (x *EchoResponse) GetRequest() *EchoRequest {
	if x != nil {
		return x.Request
	}
	return nil
}

var File_grpcecho_proto protoreflect.FileDescriptor

var file_grpcecho_proto_rawDesc = []byte{
	0x0a, 0x0e, 0x67, 0x72, 0x70, 0x63, 0x65, 0x63, 0x68, 0x6f, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f,
	0x12, 0x2a, 0x67, 0x61, 0x74, 0x65, 0x77, 0x61, 0x79, 0x5f, 0x61, 0x70, 0x69, 0x5f, 0x63, 0x6f,
	0x6e, 0x66, 0x6f, 0x72, 0x6d, 0x61, 0x6e, 0x63, 0x65, 0x2e, 0x67, 0x72, 0x70, 0x63, 0x5f, 0x65,
	0x63, 0x68, 0x6f, 0x2e, 0x67, 0x72, 0x70, 0x63, 0x65, 0x63, 0x68, 0x6f, 0x22, 0x30, 0x0a, 0x06,
	0x48, 0x65, 0x61, 0x64, 0x65, 0x72, 0x12, 0x10, 0x0a, 0x03, 0x6b, 0x65, 0x79, 0x18, 0x01, 0x20,
	0x01, 0x28, 0x09, 0x52, 0x03, 0x6b, 0x65, 0x79, 0x12, 0x14, 0x0a, 0x05, 0x76, 0x61, 0x6c, 0x75,
	0x65, 0x18, 0x02, 0x20, 0x01, 0x28, 0x09, 0x52, 0x05, 0x76, 0x61, 0x6c, 0x75, 0x65, 0x22, 0x76,
	0x0a, 0x07, 0x43, 0x6f, 0x6e, 0x74, 0x65, 0x78, 0x74, 0x12, 0x1c, 0x0a, 0x09, 0x6e, 0x61, 0x6d,
	0x65, 0x73, 0x70, 0x61, 0x63, 0x65, 0x18, 0x01, 0x20, 0x01, 0x28, 0x09, 0x52, 0x09, 0x6e, 0x61,
	0x6d, 0x65, 0x73, 0x70, 0x61, 0x63, 0x65, 0x12, 0x18, 0x0a, 0x07, 0x69, 0x6e, 0x67, 0x72, 0x65,
	0x73, 0x73, 0x18, 0x02, 0x20, 0x01, 0x28, 0x09, 0x52, 0x07, 0x69, 0x6e, 0x67, 0x72, 0x65, 0x73,
	0x73, 0x12, 0x21, 0x0a, 0x0c, 0x73, 0x65, 0x72, 0x76, 0x69, 0x63, 0x65, 0x5f, 0x6e, 0x61, 0x6d,
	0x65, 0x18, 0x03, 0x20, 0x01, 0x28, 0x09, 0x52, 0x0b, 0x73, 0x65, 0x72, 0x76, 0x69, 0x63, 0x65,
	0x4e, 0x61, 0x6d, 0x65, 0x12, 0x10, 0x0a, 0x03, 0x70, 0x6f, 0x64, 0x18, 0x04, 0x20, 0x01, 0x28,
	0x09, 0x52, 0x03, 0x70, 0x6f, 0x64, 0x22, 0xcb, 0x01, 0x0a, 0x0d, 0x54, 0x4c, 0x53, 0x41, 0x73,
	0x73, 0x65, 0x72, 0x74, 0x69, 0x6f, 0x6e, 0x73, 0x12, 0x18, 0x0a, 0x07, 0x76, 0x65, 0x72, 0x73,
	0x69, 0x6f, 0x6e, 0x18, 0x01, 0x20, 0x01, 0x28, 0x09, 0x52, 0x07, 0x76, 0x65, 0x72, 0x73, 0x69,
	0x6f, 0x6e, 0x12, 0x2f, 0x0a, 0x13, 0x6e, 0x65, 0x67, 0x6f, 0x74, 0x69, 0x61, 0x74, 0x65, 0x64,
	0x5f, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x63, 0x6f, 0x6c, 0x18, 0x02, 0x20, 0x01, 0x28, 0x09, 0x52,
	0x12, 0x6e, 0x65, 0x67, 0x6f, 0x74, 0x69, 0x61, 0x74, 0x65, 0x64, 0x50, 0x72, 0x6f, 0x74, 0x6f,
	0x63, 0x6f, 0x6c, 0x12, 0x1f, 0x0a, 0x0b, 0x73, 0x65, 0x72, 0x76, 0x65, 0x72, 0x5f, 0x6e, 0x61,
	0x6d, 0x65, 0x18, 0x03, 0x20, 0x01, 0x28, 0x09, 0x52, 0x0a, 0x73, 0x65, 0x72, 0x76, 0x65, 0x72,
	0x4e, 0x61, 0x6d, 0x65, 0x12, 0x21, 0x0a, 0x0c, 0x63, 0x69, 0x70, 0x68, 0x65, 0x72, 0x5f, 0x73,
	0x75, 0x69, 0x74, 0x65, 0x18, 0x04, 0x20, 0x01, 0x28, 0x09, 0x52, 0x0b, 0x63, 0x69, 0x70, 0x68,
	0x65, 0x72, 0x53, 0x75, 0x69, 0x74, 0x65, 0x12, 0x2b, 0x0a, 0x11, 0x70, 0x65, 0x65, 0x72, 0x5f,
	0x63, 0x65, 0x72, 0x74, 0x69, 0x66, 0x69, 0x63, 0x61, 0x74, 0x65, 0x73, 0x18, 0x05, 0x20, 0x03,
	0x28, 0x09, 0x52, 0x10, 0x70, 0x65, 0x65, 0x72, 0x43, 0x65, 0x72, 0x74, 0x69, 0x66, 0x69, 0x63,
	0x61, 0x74, 0x65, 0x73, 0x22, 0xdf, 0x02, 0x0a, 0x0a, 0x41, 0x73, 0x73, 0x65, 0x72, 0x74, 0x69,
	0x6f, 0x6e, 0x73, 0x12, 0x34, 0x0a, 0x16, 0x66, 0x75, 0x6c, 0x6c, 0x79, 0x5f, 0x71, 0x75, 0x61,
	0x6c, 0x69, 0x66, 0x69, 0x65, 0x64, 0x5f, 0x6d, 0x65, 0x74, 0x68, 0x6f, 0x64, 0x18, 0x01, 0x20,
	0x01, 0x28, 0x09, 0x52, 0x14, 0x66, 0x75, 0x6c, 0x6c, 0x79, 0x51, 0x75, 0x61, 0x6c, 0x69, 0x66,
	0x69, 0x65, 0x64, 0x4d, 0x65, 0x74, 0x68, 0x6f, 0x64, 0x12, 0x4c, 0x0a, 0x07, 0x68, 0x65, 0x61,
	0x64, 0x65, 0x72, 0x73, 0x18, 0x02, 0x20, 0x03, 0x28, 0x0b, 0x32, 0x32, 0x2e, 0x67, 0x61, 0x74,
	0x65, 0x77, 0x61, 0x79, 0x5f, 0x61, 0x70, 0x69, 0x5f, 0x63, 0x6f, 0x6e, 0x66, 0x6f, 0x72, 0x6d,
	0x61, 0x6e, 0x63, 0x65, 0x2e, 0x67, 0x72, 0x70, 0x63, 0x5f, 0x65, 0x63, 0x68, 0x6f, 0x2e, 0x67,
	0x72, 0x70, 0x63, 0x65, 0x63, 0x68, 0x6f, 0x2e, 0x48, 0x65, 0x61, 0x64, 0x65, 0x72, 0x52, 0x07,
	0x68, 0x65, 0x61, 0x64, 0x65, 0x72, 0x73, 0x12, 0x1c, 0x0a, 0x09, 0x61, 0x75, 0x74, 0x68, 0x6f,
	0x72, 0x69, 0x74, 0x79, 0x18, 0x03, 0x20, 0x01, 0x28, 0x09, 0x52, 0x09, 0x61, 0x75, 0x74, 0x68,
	0x6f, 0x72, 0x69, 0x74, 0x79, 0x12, 0x4d, 0x0a, 0x07, 0x63, 0x6f, 0x6e, 0x74, 0x65, 0x78, 0x74,
	0x18, 0x04, 0x20, 0x01, 0x28, 0x0b, 0x32, 0x33, 0x2e, 0x67, 0x61, 0x74, 0x65, 0x77, 0x61, 0x79,
	0x5f, 0x61, 0x70, 0x69, 0x5f, 0x63, 0x6f, 0x6e, 0x66, 0x6f, 0x72, 0x6d, 0x61, 0x6e, 0x63, 0x65,
	0x2e, 0x67, 0x72, 0x70, 0x63, 0x5f, 0x65, 0x63, 0x68, 0x6f, 0x2e, 0x67, 0x72, 0x70, 0x63, 0x65,
	0x63, 0x68, 0x6f, 0x2e, 0x43, 0x6f, 0x6e, 0x74, 0x65, 0x78, 0x74, 0x52, 0x07, 0x63, 0x6f, 0x6e,
	0x74, 0x65, 0x78, 0x74, 0x12, 0x60, 0x0a, 0x0e, 0x74, 0x6c, 0x73, 0x5f, 0x61, 0x73, 0x73, 0x65,
	0x72, 0x74, 0x69, 0x6f, 0x6e, 0x73, 0x18, 0x05, 0x20, 0x01, 0x28, 0x0b, 0x32, 0x39, 0x2e, 0x67,
	0x61, 0x74, 0x65, 0x77, 0x61, 0x79, 0x5f, 0x61, 0x70, 0x69, 0x5f, 0x63, 0x6f, 0x6e, 0x66, 0x6f,
	0x72, 0x6d, 0x61, 0x6e, 0x63, 0x65, 0x2e, 0x67, 0x72, 0x70, 0x63, 0x5f, 0x65, 0x63, 0x68, 0x6f,
	0x2e, 0x67, 0x72, 0x70, 0x63, 0x65, 0x63, 0x68, 0x6f, 0x2e, 0x54, 0x4c, 0x53, 0x41, 0x73, 0x73,
	0x65, 0x72, 0x74, 0x69, 0x6f, 0x6e, 0x73, 0x52, 0x0d, 0x74, 0x6c, 0x73, 0x41, 0x73, 0x73, 0x65,
	0x72, 0x74, 0x69, 0x6f, 0x6e, 0x73, 0x22, 0x0d, 0x0a, 0x0b, 0x45, 0x63, 0x68, 0x6f, 0x52, 0x65,
	0x71, 0x75, 0x65, 0x73, 0x74, 0x22, 0xb9, 0x01, 0x0a, 0x0c, 0x45, 0x63, 0x68, 0x6f, 0x52, 0x65,
	0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65, 0x12, 0x56, 0x0a, 0x0a, 0x61, 0x73, 0x73, 0x65, 0x72, 0x74,
	0x69, 0x6f, 0x6e, 0x73, 0x18, 0x01, 0x20, 0x01, 0x28, 0x0b, 0x32, 0x36, 0x2e, 0x67, 0x61, 0x74,
	0x65, 0x77, 0x61, 0x79, 0x5f, 0x61, 0x70, 0x69, 0x5f, 0x63, 0x6f, 0x6e, 0x66, 0x6f, 0x72, 0x6d,
	0x61, 0x6e, 0x63, 0x65, 0x2e, 0x67, 0x72, 0x70, 0x63, 0x5f, 0x65, 0x63, 0x68, 0x6f, 0x2e, 0x67,
	0x72, 0x70, 0x63, 0x65, 0x63, 0x68, 0x6f, 0x2e, 0x41, 0x73, 0x73, 0x65, 0x72, 0x74, 0x69, 0x6f,
	0x6e, 0x73, 0x52, 0x0a, 0x61, 0x73, 0x73, 0x65, 0x72, 0x74, 0x69, 0x6f, 0x6e, 0x73, 0x12, 0x51,
	0x0a, 0x07, 0x72, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x18, 0x02, 0x20, 0x01, 0x28, 0x0b, 0x32,
	0x37, 0x2e, 0x67, 0x61, 0x74, 0x65, 0x77, 0x61, 0x79, 0x5f, 0x61, 0x70, 0x69, 0x5f, 0x63, 0x6f,
	0x6e, 0x66, 0x6f, 0x72, 0x6d, 0x61, 0x6e, 0x63, 0x65, 0x2e, 0x67, 0x72, 0x70, 0x63, 0x5f, 0x65,
	0x63, 0x68, 0x6f, 0x2e, 0x67, 0x72, 0x70, 0x63, 0x65, 0x63, 0x68, 0x6f, 0x2e, 0x45, 0x63, 0x68,
	0x6f, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x52, 0x07, 0x72, 0x65, 0x71, 0x75, 0x65, 0x73,
	0x74, 0x32, 0x8a, 0x03, 0x0a, 0x08, 0x47, 0x72, 0x70, 0x63, 0x45, 0x63, 0x68, 0x6f, 0x12, 0x7b,
	0x0a, 0x04, 0x45, 0x63, 0x68, 0x6f, 0x12, 0x37, 0x2e, 0x67, 0x61, 0x74, 0x65, 0x77, 0x61, 0x79,
	0x5f, 0x61, 0x70, 0x69, 0x5f, 0x63, 0x6f, 0x6e, 0x66, 0x6f, 0x72, 0x6d, 0x61, 0x6e, 0x63, 0x65,
	0x2e, 0x67, 0x72, 0x70, 0x63, 0x5f, 0x65, 0x63, 0x68, 0x6f, 0x2e, 0x67, 0x72, 0x70, 0x63, 0x65,
	0x63, 0x68, 0x6f, 0x2e, 0x45, 0x63, 0x68, 0x6f, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x1a,
	0x38, 0x2e, 0x67, 0x61, 0x74, 0x65, 0x77, 0x61, 0x79, 0x5f, 0x61, 0x70, 0x69, 0x5f, 0x63, 0x6f,
	0x6e, 0x66, 0x6f, 0x72, 0x6d, 0x61, 0x6e, 0x63, 0x65, 0x2e, 0x67, 0x72, 0x70, 0x63, 0x5f, 0x65,
	0x63, 0x68, 0x6f, 0x2e, 0x67, 0x72, 0x70, 0x63, 0x65, 0x63, 0x68, 0x6f, 0x2e, 0x45, 0x63, 0x68,
	0x6f, 0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65, 0x22, 0x00, 0x12, 0x7e, 0x0a, 0x07, 0x45,
	0x63, 0x68, 0x6f, 0x54, 0x77, 0x6f, 0x12, 0x37, 0x2e, 0x67, 0x61, 0x74, 0x65, 0x77, 0x61, 0x79,
	0x5f, 0x61, 0x70, 0x69, 0x5f, 0x63, 0x6f, 0x6e, 0x66, 0x6f, 0x72, 0x6d, 0x61, 0x6e, 0x63, 0x65,
	0x2e, 0x67, 0x72, 0x70, 0x63, 0x5f, 0x65, 0x63, 0x68, 0x6f, 0x2e, 0x67, 0x72, 0x70, 0x63, 0x65,
	0x63, 0x68, 0x6f, 0x2e, 0x45, 0x63, 0x68, 0x6f, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x1a,
	0x38, 0x2e, 0x67, 0x61, 0x74, 0x65, 0x77, 0x61, 0x79, 0x5f, 0x61, 0x70, 0x69, 0x5f, 0x63, 0x6f,
	0x6e, 0x66, 0x6f, 0x72, 0x6d, 0x61, 0x6e, 0x63, 0x65, 0x2e, 0x67, 0x72, 0x70, 0x63, 0x5f, 0x65,
	0x63, 0x68, 0x6f, 0x2e, 0x67, 0x72, 0x70, 0x63, 0x65, 0x63, 0x68, 0x6f, 0x2e, 0x45, 0x63, 0x68,
	0x6f, 0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65, 0x22, 0x00, 0x12, 0x80, 0x01, 0x0a, 0x09,
	0x45, 0x63, 0x68, 0x6f, 0x54, 0x68, 0x72, 0x65, 0x65, 0x12, 0x37, 0x2e, 0x67, 0x61, 0x74, 0x65,
	0x77, 0x61, 0x79, 0x5f, 0x61, 0x70, 0x69, 0x5f, 0x63, 0x6f, 0x6e, 0x66, 0x6f, 0x72, 0x6d, 0x61,
	0x6e, 0x63, 0x65, 0x2e, 0x67, 0x72, 0x70, 0x63, 0x5f, 0x65, 0x63, 0x68, 0x6f, 0x2e, 0x67, 0x72,
	0x70, 0x63, 0x65, 0x63, 0x68, 0x6f, 0x2e, 0x45, 0x63, 0x68, 0x6f, 0x52, 0x65, 0x71, 0x75, 0x65,
	0x73, 0x74, 0x1a, 0x38, 0x2e, 0x67, 0x61, 0x74, 0x65, 0x77, 0x61, 0x79, 0x5f, 0x61, 0x70, 0x69,
	0x5f, 0x63, 0x6f, 0x6e, 0x66, 0x6f, 0x72, 0x6d, 0x61, 0x6e, 0x63, 0x65, 0x2e, 0x67, 0x72, 0x70,
	0x63, 0x5f, 0x65, 0x63, 0x68, 0x6f, 0x2e, 0x67, 0x72, 0x70, 0x63, 0x65, 0x63, 0x68, 0x6f, 0x2e,
	0x45, 0x63, 0x68, 0x6f, 0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65, 0x22, 0x00, 0x42, 0x3e,
	0x5a, 0x3c, 0x73, 0x69, 0x67, 0x73, 0x2e, 0x6b, 0x38, 0x73, 0x2e, 0x69, 0x6f, 0x2f, 0x67, 0x61,
	0x74, 0x65, 0x77, 0x61, 0x79, 0x2d, 0x61, 0x70, 0x69, 0x2f, 0x63, 0x6f, 0x6e, 0x66, 0x6f, 0x72,
	0x6d, 0x61, 0x6e, 0x63, 0x65, 0x2f, 0x67, 0x72, 0x70, 0x63, 0x2d, 0x65, 0x63, 0x68, 0x6f, 0x2f,
	0x67, 0x72, 0x70, 0x63, 0x65, 0x63, 0x68, 0x6f, 0x73, 0x65, 0x72, 0x76, 0x65, 0x72, 0x62, 0x06,
	0x70, 0x72, 0x6f, 0x74, 0x6f, 0x33,
}

var (
	file_grpcecho_proto_rawDescOnce sync.Once
	file_grpcecho_proto_rawDescData = file_grpcecho_proto_rawDesc
)

func file_grpcecho_proto_rawDescGZIP() []byte {
	file_grpcecho_proto_rawDescOnce.Do(func() {
		file_grpcecho_proto_rawDescData = protoimpl.X.CompressGZIP(file_grpcecho_proto_rawDescData)
	})
	return file_grpcecho_proto_rawDescData
}

var file_grpcecho_proto_msgTypes = make([]protoimpl.MessageInfo, 6)
var file_grpcecho_proto_goTypes = []interface{}{
	(*Header)(nil),        // 0: gateway_api_conformance.grpc_echo.grpcecho.Header
	(*Context)(nil),       // 1: gateway_api_conformance.grpc_echo.grpcecho.Context
	(*TLSAssertions)(nil), // 2: gateway_api_conformance.grpc_echo.grpcecho.TLSAssertions
	(*Assertions)(nil),    // 3: gateway_api_conformance.grpc_echo.grpcecho.Assertions
	(*EchoRequest)(nil),   // 4: gateway_api_conformance.grpc_echo.grpcecho.EchoRequest
	(*EchoResponse)(nil),  // 5: gateway_api_conformance.grpc_echo.grpcecho.EchoResponse
}
var file_grpcecho_proto_depIdxs = []int32{
	0, // 0: gateway_api_conformance.grpc_echo.grpcecho.Assertions.headers:type_name -> gateway_api_conformance.grpc_echo.grpcecho.Header
	1, // 1: gateway_api_conformance.grpc_echo.grpcecho.Assertions.context:type_name -> gateway_api_conformance.grpc_echo.grpcecho.Context
	2, // 2: gateway_api_conformance.grpc_echo.grpcecho.Assertions.tls_assertions:type_name -> gateway_api_conformance.grpc_echo.grpcecho.TLSAssertions
	3, // 3: gateway_api_conformance.grpc_echo.grpcecho.EchoResponse.assertions:type_name -> gateway_api_conformance.grpc_echo.grpcecho.Assertions
	4, // 4: gateway_api_conformance.grpc_echo.grpcecho.EchoResponse.request:type_name -> gateway_api_conformance.grpc_echo.grpcecho.EchoRequest
	4, // 5: gateway_api_conformance.grpc_echo.grpcecho.GrpcEcho.Echo:input_type -> gateway_api_conformance.grpc_echo.grpcecho.EchoRequest
	4, // 6: gateway_api_conformance.grpc_echo.grpcecho.GrpcEcho.EchoTwo:input_type -> gateway_api_conformance.grpc_echo.grpcecho.EchoRequest
	4, // 7: gateway_api_conformance.grpc_echo.grpcecho.GrpcEcho.EchoThree:input_type -> gateway_api_conformance.grpc_echo.grpcecho.EchoRequest
	5, // 8: gateway_api_conformance.grpc_echo.grpcecho.GrpcEcho.Echo:output_type -> gateway_api_conformance.grpc_echo.grpcecho.EchoResponse
	5, // 9: gateway_api_conformance.grpc_echo.grpcecho.GrpcEcho.EchoTwo:output_type -> gateway_api_conformance.grpc_echo.grpcecho.EchoResponse
	5, // 10: gateway_api_conformance.grpc_echo.grpcecho.GrpcEcho.EchoThree:output_type -> gateway_api_conformance.grpc_echo.grpcecho.EchoResponse
	8, // [8:11] is the sub-list for method output_type
	5, // [5:8] is the sub-list for method input_type
	5, // [5:5] is the sub-list for extension type_name
	5, // [5:5] is the sub-list for extension extendee
	0, // [0:5] is the sub-list for field type_name
}

func init() { file_grpcecho_proto_init() }
func file_grpcecho_proto_init() {
	if File_grpcecho_proto != nil {
		return
	}
	if !protoimpl.UnsafeEnabled {
		file_grpcecho_proto_msgTypes[0].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*Header); i {
			case 0:
				return &v.state
			case 1:
				return &v.sizeCache
			case 2:
				return &v.unknownFields
			default:
				return nil
			}
		}
		file_grpcecho_proto_msgTypes[1].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*Context); i {
			case 0:
				return &v.state
			case 1:
				return &v.sizeCache
			case 2:
				return &v.unknownFields
			default:
				return nil
			}
		}
		file_grpcecho_proto_msgTypes[2].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*TLSAssertions); i {
			case 0:
				return &v.state
			case 1:
				return &v.sizeCache
			case 2:
				return &v.unknownFields
			default:
				return nil
			}
		}
		file_grpcecho_proto_msgTypes[3].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*Assertions); i {
			case 0:
				return &v.state
			case 1:
				return &v.sizeCache
			case 2:
				return &v.unknownFields
			default:
				return nil
			}
		}
		file_grpcecho_proto_msgTypes[4].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*EchoRequest); i {
			case 0:
				return &v.state
			case 1:
				return &v.sizeCache
			case 2:
				return &v.unknownFields
			default:
				return nil
			}
		}
		file_grpcecho_proto_msgTypes[5].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*EchoResponse); i {
			case 0:
				return &v.state
			case 1:
				return &v.sizeCache
			case 2:
				return &v.unknownFields
			default:
				return nil
			}
		}
	}
	type x struct{}
	out := protoimpl.TypeBuilder{
		File: protoimpl.DescBuilder{
			GoPackagePath: reflect.TypeOf(x{}).PkgPath(),
			RawDescriptor: file_grpcecho_proto_rawDesc,
			NumEnums:      0,
			NumMessages:   6,
			NumExtensions: 0,
			NumServices:   1,
		},
		GoTypes:           file_grpcecho_proto_goTypes,
		DependencyIndexes: file_grpcecho_proto_depIdxs,
		MessageInfos:      file_grpcecho_proto_msgTypes,
	}.Build()
	File_grpcecho_proto = out.File
	file_grpcecho_proto_rawDesc = nil
	file_grpcecho_proto_goTypes = nil
	file_grpcecho_proto_depIdxs = nil
}
