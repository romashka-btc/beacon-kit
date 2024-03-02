// SPDX-License-Identifier: MIT
//
// Copyright (c) 2024 Berachain Foundation
//
// Permission is hereby granted, free of charge, to any person
// obtaining a copy of this software and associated documentation
// files (the "Software"), to deal in the Software without
// restriction, including without limitation the rights to use,
// copy, modify, merge, publish, distribute, sublicense, and/or sell
// copies of the Software, and to permit persons to whom the
// Software is furnished to do so, subject to the following
// conditions:
//
// The above copyright notice and this permission notice shall be
// included in all copies or substantial portions of the Software.
//
// THE SOFTWARE IS PROVIDED "AS IS", WITHOUT WARRANTY OF ANY KIND,
// EXPRESS OR IMPLIED, INCLUDING BUT NOT LIMITED TO THE WARRANTIES
// OF MERCHANTABILITY, FITNESS FOR A PARTICULAR PURPOSE AND
// NONINFRINGEMENT. IN NO EVENT SHALL THE AUTHORS OR COPYRIGHT
// HOLDERS BE LIABLE FOR ANY CLAIM, DAMAGES OR OTHER LIABILITY,
// WHETHER IN AN ACTION OF CONTRACT, TORT OR OTHERWISE, ARISING
// FROM, OUT OF OR IN CONNECTION WITH THE SOFTWARE OR THE USE OR
// OTHER DEALINGS IN THE SOFTWARE.

// Code generated by protoc-gen-go. DO NOT EDIT.
// versions:
// 	protoc-gen-go v1.32.0
// 	protoc        (unknown)
// source: beacon/core/types/v1/buoy.proto

package typesv1

import (
	v1 "github.com/prysmaticlabs/prysm/v5/proto/engine/v1"
	github_com_itsdevbear_bolaris_primitives "github.com/itsdevbear/bolaris/primitives"
	_ "github.com/prysmaticlabs/prysm/v5/proto/eth/ext"
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

// BeaconBuoyDeneb represents a beacon kit block during deneb network upgrade.
type BeaconBuoyDeneb struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	// Beacon chain slot that this block represents.
	Slot github_com_itsdevbear_bolaris_primitives.Slot `protobuf:"varint,1,opt,name=slot,proto3" json:"slot,omitempty" cast-type:"github.com/itsdevbear/bolaris/primitives.Slot"`
	// 32 byte root of the parent block.
	ParentRoot []byte `protobuf:"bytes,2,opt,name=parent_root,json=parentRoot,proto3" json:"parent_root,omitempty" ssz-size:"32"`
	// BeaconBlockBody contains the body of the beacon block.
	Body *BeaconBuoyBodyDeneb `protobuf:"bytes,3,opt,name=body,proto3" json:"body,omitempty"`
	// The payload value of the block.
	PayloadValue []byte `protobuf:"bytes,101,opt,name=payload_value,json=payloadValue,proto3" json:"payload_value,omitempty" ssz-size:"32"`
}

func (x *BeaconBuoyDeneb) Reset() {
	*x = BeaconBuoyDeneb{}
	if protoimpl.UnsafeEnabled {
		mi := &file_beacon_core_types_v1_buoy_proto_msgTypes[0]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *BeaconBuoyDeneb) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*BeaconBuoyDeneb) ProtoMessage() {}

func (x *BeaconBuoyDeneb) ProtoReflect() protoreflect.Message {
	mi := &file_beacon_core_types_v1_buoy_proto_msgTypes[0]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use BeaconBuoyDeneb.ProtoReflect.Descriptor instead.
func (*BeaconBuoyDeneb) Descriptor() ([]byte, []int) {
	return file_beacon_core_types_v1_buoy_proto_rawDescGZIP(), []int{0}
}

func (x *BeaconBuoyDeneb) GetSlot() github_com_itsdevbear_bolaris_primitives.Slot {
	if x != nil {
		return x.Slot
	}
	return github_com_itsdevbear_bolaris_primitives.Slot(0)
}

func (x *BeaconBuoyDeneb) GetParentRoot() []byte {
	if x != nil {
		return x.ParentRoot
	}
	return nil
}

func (x *BeaconBuoyDeneb) GetBody() *BeaconBuoyBodyDeneb {
	if x != nil {
		return x.Body
	}
	return nil
}

func (x *BeaconBuoyDeneb) GetPayloadValue() []byte {
	if x != nil {
		return x.PayloadValue
	}
	return nil
}

// BeaconBuoyBodyDeneb represents the body of a beacon block in Deneb.
type BeaconBuoyBodyDeneb struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	// The validators RANDAO reveal 96 byte value.
	RandaoReveal []byte `protobuf:"bytes,1,opt,name=randao_reveal,json=randaoReveal,proto3" json:"randao_reveal,omitempty" ssz-size:"96"`
	// 32 byte field of arbitrary data. This field may contain any data and
	// is not used for anything other than a fun message.
	Graffiti []byte `protobuf:"bytes,2,opt,name=graffiti,proto3" json:"graffiti,omitempty" ssz-size:"32"`
	// Deposits from the execution chain, at most MAX_DEPOSITS_PER_BLOCK.
	Deposits []*Deposit `protobuf:"bytes,3,rep,name=deposits,proto3" json:"deposits,omitempty" ssz-max:"16"`
	// Execution payload from the execution chain. New in Bellatrix network upgrade.
	ExecutionPayload *v1.ExecutionPayloadDeneb `protobuf:"bytes,4,opt,name=execution_payload,json=executionPayload,proto3" json:"execution_payload,omitempty"`
	// At most MAX_BLOBS_PER_BLOCK. New in Deneb network upgrade.
	BlobKzgCommitments [][]byte `protobuf:"bytes,5,rep,name=blob_kzg_commitments,json=blobKzgCommitments,proto3" json:"blob_kzg_commitments,omitempty" ssz-max:"16" ssz-size:"?,48"`
}

func (x *BeaconBuoyBodyDeneb) Reset() {
	*x = BeaconBuoyBodyDeneb{}
	if protoimpl.UnsafeEnabled {
		mi := &file_beacon_core_types_v1_buoy_proto_msgTypes[1]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *BeaconBuoyBodyDeneb) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*BeaconBuoyBodyDeneb) ProtoMessage() {}

func (x *BeaconBuoyBodyDeneb) ProtoReflect() protoreflect.Message {
	mi := &file_beacon_core_types_v1_buoy_proto_msgTypes[1]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use BeaconBuoyBodyDeneb.ProtoReflect.Descriptor instead.
func (*BeaconBuoyBodyDeneb) Descriptor() ([]byte, []int) {
	return file_beacon_core_types_v1_buoy_proto_rawDescGZIP(), []int{1}
}

func (x *BeaconBuoyBodyDeneb) GetRandaoReveal() []byte {
	if x != nil {
		return x.RandaoReveal
	}
	return nil
}

func (x *BeaconBuoyBodyDeneb) GetGraffiti() []byte {
	if x != nil {
		return x.Graffiti
	}
	return nil
}

func (x *BeaconBuoyBodyDeneb) GetDeposits() []*Deposit {
	if x != nil {
		return x.Deposits
	}
	return nil
}

func (x *BeaconBuoyBodyDeneb) GetExecutionPayload() *v1.ExecutionPayloadDeneb {
	if x != nil {
		return x.ExecutionPayload
	}
	return nil
}

func (x *BeaconBuoyBodyDeneb) GetBlobKzgCommitments() [][]byte {
	if x != nil {
		return x.BlobKzgCommitments
	}
	return nil
}

var File_beacon_core_types_v1_buoy_proto protoreflect.FileDescriptor

var file_beacon_core_types_v1_buoy_proto_rawDesc = []byte{
	0x0a, 0x1f, 0x62, 0x65, 0x61, 0x63, 0x6f, 0x6e, 0x2f, 0x63, 0x6f, 0x72, 0x65, 0x2f, 0x74, 0x79,
	0x70, 0x65, 0x73, 0x2f, 0x76, 0x31, 0x2f, 0x62, 0x75, 0x6f, 0x79, 0x2e, 0x70, 0x72, 0x6f, 0x74,
	0x6f, 0x12, 0x14, 0x62, 0x65, 0x61, 0x63, 0x6f, 0x6e, 0x2e, 0x63, 0x6f, 0x72, 0x65, 0x2e, 0x74,
	0x79, 0x70, 0x65, 0x73, 0x2e, 0x76, 0x31, 0x1a, 0x22, 0x62, 0x65, 0x61, 0x63, 0x6f, 0x6e, 0x2f,
	0x63, 0x6f, 0x72, 0x65, 0x2f, 0x74, 0x79, 0x70, 0x65, 0x73, 0x2f, 0x76, 0x31, 0x2f, 0x64, 0x65,
	0x70, 0x6f, 0x73, 0x69, 0x74, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x1a, 0x29, 0x65, 0x74, 0x68,
	0x65, 0x72, 0x65, 0x75, 0x6d, 0x2f, 0x65, 0x6e, 0x67, 0x69, 0x6e, 0x65, 0x2f, 0x76, 0x31, 0x2f,
	0x65, 0x78, 0x65, 0x63, 0x75, 0x74, 0x69, 0x6f, 0x6e, 0x5f, 0x65, 0x6e, 0x67, 0x69, 0x6e, 0x65,
	0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x1a, 0x1e, 0x65, 0x74, 0x68, 0x65, 0x72, 0x65, 0x75, 0x6d,
	0x2f, 0x65, 0x74, 0x68, 0x2f, 0x65, 0x78, 0x74, 0x2f, 0x6f, 0x70, 0x74, 0x69, 0x6f, 0x6e, 0x73,
	0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x22, 0xed, 0x01, 0x0a, 0x0f, 0x42, 0x65, 0x61, 0x63, 0x6f,
	0x6e, 0x42, 0x75, 0x6f, 0x79, 0x44, 0x65, 0x6e, 0x65, 0x62, 0x12, 0x45, 0x0a, 0x04, 0x73, 0x6c,
	0x6f, 0x74, 0x18, 0x01, 0x20, 0x01, 0x28, 0x04, 0x42, 0x31, 0x82, 0xb5, 0x18, 0x2d, 0x67, 0x69,
	0x74, 0x68, 0x75, 0x62, 0x2e, 0x63, 0x6f, 0x6d, 0x2f, 0x69, 0x74, 0x73, 0x64, 0x65, 0x76, 0x62,
	0x65, 0x61, 0x72, 0x2f, 0x62, 0x6f, 0x6c, 0x61, 0x72, 0x69, 0x73, 0x2f, 0x70, 0x72, 0x69, 0x6d,
	0x69, 0x74, 0x69, 0x76, 0x65, 0x73, 0x2e, 0x53, 0x6c, 0x6f, 0x74, 0x52, 0x04, 0x73, 0x6c, 0x6f,
	0x74, 0x12, 0x27, 0x0a, 0x0b, 0x70, 0x61, 0x72, 0x65, 0x6e, 0x74, 0x5f, 0x72, 0x6f, 0x6f, 0x74,
	0x18, 0x02, 0x20, 0x01, 0x28, 0x0c, 0x42, 0x06, 0x8a, 0xb5, 0x18, 0x02, 0x33, 0x32, 0x52, 0x0a,
	0x70, 0x61, 0x72, 0x65, 0x6e, 0x74, 0x52, 0x6f, 0x6f, 0x74, 0x12, 0x3d, 0x0a, 0x04, 0x62, 0x6f,
	0x64, 0x79, 0x18, 0x03, 0x20, 0x01, 0x28, 0x0b, 0x32, 0x29, 0x2e, 0x62, 0x65, 0x61, 0x63, 0x6f,
	0x6e, 0x2e, 0x63, 0x6f, 0x72, 0x65, 0x2e, 0x74, 0x79, 0x70, 0x65, 0x73, 0x2e, 0x76, 0x31, 0x2e,
	0x42, 0x65, 0x61, 0x63, 0x6f, 0x6e, 0x42, 0x75, 0x6f, 0x79, 0x42, 0x6f, 0x64, 0x79, 0x44, 0x65,
	0x6e, 0x65, 0x62, 0x52, 0x04, 0x62, 0x6f, 0x64, 0x79, 0x12, 0x2b, 0x0a, 0x0d, 0x70, 0x61, 0x79,
	0x6c, 0x6f, 0x61, 0x64, 0x5f, 0x76, 0x61, 0x6c, 0x75, 0x65, 0x18, 0x65, 0x20, 0x01, 0x28, 0x0c,
	0x42, 0x06, 0x8a, 0xb5, 0x18, 0x02, 0x33, 0x32, 0x52, 0x0c, 0x70, 0x61, 0x79, 0x6c, 0x6f, 0x61,
	0x64, 0x56, 0x61, 0x6c, 0x75, 0x65, 0x22, 0xc3, 0x02, 0x0a, 0x13, 0x42, 0x65, 0x61, 0x63, 0x6f,
	0x6e, 0x42, 0x75, 0x6f, 0x79, 0x42, 0x6f, 0x64, 0x79, 0x44, 0x65, 0x6e, 0x65, 0x62, 0x12, 0x2b,
	0x0a, 0x0d, 0x72, 0x61, 0x6e, 0x64, 0x61, 0x6f, 0x5f, 0x72, 0x65, 0x76, 0x65, 0x61, 0x6c, 0x18,
	0x01, 0x20, 0x01, 0x28, 0x0c, 0x42, 0x06, 0x8a, 0xb5, 0x18, 0x02, 0x39, 0x36, 0x52, 0x0c, 0x72,
	0x61, 0x6e, 0x64, 0x61, 0x6f, 0x52, 0x65, 0x76, 0x65, 0x61, 0x6c, 0x12, 0x22, 0x0a, 0x08, 0x67,
	0x72, 0x61, 0x66, 0x66, 0x69, 0x74, 0x69, 0x18, 0x02, 0x20, 0x01, 0x28, 0x0c, 0x42, 0x06, 0x8a,
	0xb5, 0x18, 0x02, 0x33, 0x32, 0x52, 0x08, 0x67, 0x72, 0x61, 0x66, 0x66, 0x69, 0x74, 0x69, 0x12,
	0x41, 0x0a, 0x08, 0x64, 0x65, 0x70, 0x6f, 0x73, 0x69, 0x74, 0x73, 0x18, 0x03, 0x20, 0x03, 0x28,
	0x0b, 0x32, 0x1d, 0x2e, 0x62, 0x65, 0x61, 0x63, 0x6f, 0x6e, 0x2e, 0x63, 0x6f, 0x72, 0x65, 0x2e,
	0x74, 0x79, 0x70, 0x65, 0x73, 0x2e, 0x76, 0x31, 0x2e, 0x44, 0x65, 0x70, 0x6f, 0x73, 0x69, 0x74,
	0x42, 0x06, 0x92, 0xb5, 0x18, 0x02, 0x31, 0x36, 0x52, 0x08, 0x64, 0x65, 0x70, 0x6f, 0x73, 0x69,
	0x74, 0x73, 0x12, 0x56, 0x0a, 0x11, 0x65, 0x78, 0x65, 0x63, 0x75, 0x74, 0x69, 0x6f, 0x6e, 0x5f,
	0x70, 0x61, 0x79, 0x6c, 0x6f, 0x61, 0x64, 0x18, 0x04, 0x20, 0x01, 0x28, 0x0b, 0x32, 0x29, 0x2e,
	0x65, 0x74, 0x68, 0x65, 0x72, 0x65, 0x75, 0x6d, 0x2e, 0x65, 0x6e, 0x67, 0x69, 0x6e, 0x65, 0x2e,
	0x76, 0x31, 0x2e, 0x45, 0x78, 0x65, 0x63, 0x75, 0x74, 0x69, 0x6f, 0x6e, 0x50, 0x61, 0x79, 0x6c,
	0x6f, 0x61, 0x64, 0x44, 0x65, 0x6e, 0x65, 0x62, 0x52, 0x10, 0x65, 0x78, 0x65, 0x63, 0x75, 0x74,
	0x69, 0x6f, 0x6e, 0x50, 0x61, 0x79, 0x6c, 0x6f, 0x61, 0x64, 0x12, 0x40, 0x0a, 0x14, 0x62, 0x6c,
	0x6f, 0x62, 0x5f, 0x6b, 0x7a, 0x67, 0x5f, 0x63, 0x6f, 0x6d, 0x6d, 0x69, 0x74, 0x6d, 0x65, 0x6e,
	0x74, 0x73, 0x18, 0x05, 0x20, 0x03, 0x28, 0x0c, 0x42, 0x0e, 0x8a, 0xb5, 0x18, 0x04, 0x3f, 0x2c,
	0x34, 0x38, 0x92, 0xb5, 0x18, 0x02, 0x31, 0x36, 0x52, 0x12, 0x62, 0x6c, 0x6f, 0x62, 0x4b, 0x7a,
	0x67, 0x43, 0x6f, 0x6d, 0x6d, 0x69, 0x74, 0x6d, 0x65, 0x6e, 0x74, 0x73, 0x42, 0x3c, 0x5a, 0x3a,
	0x67, 0x69, 0x74, 0x68, 0x75, 0x62, 0x2e, 0x63, 0x6f, 0x6d, 0x2f, 0x69, 0x74, 0x73, 0x64, 0x65,
	0x76, 0x62, 0x65, 0x61, 0x72, 0x2f, 0x62, 0x6f, 0x6c, 0x61, 0x72, 0x69, 0x73, 0x2f, 0x62, 0x65,
	0x61, 0x63, 0x6f, 0x6e, 0x2f, 0x63, 0x6f, 0x72, 0x65, 0x2f, 0x74, 0x79, 0x70, 0x65, 0x73, 0x2f,
	0x76, 0x31, 0x3b, 0x74, 0x79, 0x70, 0x65, 0x73, 0x76, 0x31, 0x62, 0x06, 0x70, 0x72, 0x6f, 0x74,
	0x6f, 0x33,
}

var (
	file_beacon_core_types_v1_buoy_proto_rawDescOnce sync.Once
	file_beacon_core_types_v1_buoy_proto_rawDescData = file_beacon_core_types_v1_buoy_proto_rawDesc
)

func file_beacon_core_types_v1_buoy_proto_rawDescGZIP() []byte {
	file_beacon_core_types_v1_buoy_proto_rawDescOnce.Do(func() {
		file_beacon_core_types_v1_buoy_proto_rawDescData = protoimpl.X.CompressGZIP(file_beacon_core_types_v1_buoy_proto_rawDescData)
	})
	return file_beacon_core_types_v1_buoy_proto_rawDescData
}

var file_beacon_core_types_v1_buoy_proto_msgTypes = make([]protoimpl.MessageInfo, 2)
var file_beacon_core_types_v1_buoy_proto_goTypes = []interface{}{
	(*BeaconBuoyDeneb)(nil),          // 0: beacon.core.types.v1.BeaconBuoyDeneb
	(*BeaconBuoyBodyDeneb)(nil),      // 1: beacon.core.types.v1.BeaconBuoyBodyDeneb
	(*Deposit)(nil),                  // 2: beacon.core.types.v1.Deposit
	(*v1.ExecutionPayloadDeneb)(nil), // 3: ethereum.engine.v1.ExecutionPayloadDeneb
}
var file_beacon_core_types_v1_buoy_proto_depIdxs = []int32{
	1, // 0: beacon.core.types.v1.BeaconBuoyDeneb.body:type_name -> beacon.core.types.v1.BeaconBuoyBodyDeneb
	2, // 1: beacon.core.types.v1.BeaconBuoyBodyDeneb.deposits:type_name -> beacon.core.types.v1.Deposit
	3, // 2: beacon.core.types.v1.BeaconBuoyBodyDeneb.execution_payload:type_name -> ethereum.engine.v1.ExecutionPayloadDeneb
	3, // [3:3] is the sub-list for method output_type
	3, // [3:3] is the sub-list for method input_type
	3, // [3:3] is the sub-list for extension type_name
	3, // [3:3] is the sub-list for extension extendee
	0, // [0:3] is the sub-list for field type_name
}

func init() { file_beacon_core_types_v1_buoy_proto_init() }
func file_beacon_core_types_v1_buoy_proto_init() {
	if File_beacon_core_types_v1_buoy_proto != nil {
		return
	}
	file_beacon_core_types_v1_deposit_proto_init()
	if !protoimpl.UnsafeEnabled {
		file_beacon_core_types_v1_buoy_proto_msgTypes[0].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*BeaconBuoyDeneb); i {
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
		file_beacon_core_types_v1_buoy_proto_msgTypes[1].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*BeaconBuoyBodyDeneb); i {
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
			RawDescriptor: file_beacon_core_types_v1_buoy_proto_rawDesc,
			NumEnums:      0,
			NumMessages:   2,
			NumExtensions: 0,
			NumServices:   0,
		},
		GoTypes:           file_beacon_core_types_v1_buoy_proto_goTypes,
		DependencyIndexes: file_beacon_core_types_v1_buoy_proto_depIdxs,
		MessageInfos:      file_beacon_core_types_v1_buoy_proto_msgTypes,
	}.Build()
	File_beacon_core_types_v1_buoy_proto = out.File
	file_beacon_core_types_v1_buoy_proto_rawDesc = nil
	file_beacon_core_types_v1_buoy_proto_goTypes = nil
	file_beacon_core_types_v1_buoy_proto_depIdxs = nil
}