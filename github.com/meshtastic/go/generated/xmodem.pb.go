// Code generated by protoc-gen-go. DO NOT EDIT.
// versions:
// 	protoc-gen-go v1.30.0
// 	protoc        v3.21.12
// source: meshtastic/xmodem.proto

package generated

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

type XModem_Control int32

const (
	XModem_NUL   XModem_Control = 0
	XModem_SOH   XModem_Control = 1
	XModem_STX   XModem_Control = 2
	XModem_EOT   XModem_Control = 4
	XModem_ACK   XModem_Control = 6
	XModem_NAK   XModem_Control = 21
	XModem_CAN   XModem_Control = 24
	XModem_CTRLZ XModem_Control = 26
)

// Enum value maps for XModem_Control.
var (
	XModem_Control_name = map[int32]string{
		0:  "NUL",
		1:  "SOH",
		2:  "STX",
		4:  "EOT",
		6:  "ACK",
		21: "NAK",
		24: "CAN",
		26: "CTRLZ",
	}
	XModem_Control_value = map[string]int32{
		"NUL":   0,
		"SOH":   1,
		"STX":   2,
		"EOT":   4,
		"ACK":   6,
		"NAK":   21,
		"CAN":   24,
		"CTRLZ": 26,
	}
)

func (x XModem_Control) Enum() *XModem_Control {
	p := new(XModem_Control)
	*p = x
	return p
}

func (x XModem_Control) String() string {
	return protoimpl.X.EnumStringOf(x.Descriptor(), protoreflect.EnumNumber(x))
}

func (XModem_Control) Descriptor() protoreflect.EnumDescriptor {
	return file_meshtastic_xmodem_proto_enumTypes[0].Descriptor()
}

func (XModem_Control) Type() protoreflect.EnumType {
	return &file_meshtastic_xmodem_proto_enumTypes[0]
}

func (x XModem_Control) Number() protoreflect.EnumNumber {
	return protoreflect.EnumNumber(x)
}

// Deprecated: Use XModem_Control.Descriptor instead.
func (XModem_Control) EnumDescriptor() ([]byte, []int) {
	return file_meshtastic_xmodem_proto_rawDescGZIP(), []int{0, 0}
}

type XModem struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Control XModem_Control `protobuf:"varint,1,opt,name=control,proto3,enum=meshtastic.XModem_Control" json:"control,omitempty"`
	Seq     uint32         `protobuf:"varint,2,opt,name=seq,proto3" json:"seq,omitempty"`
	Crc16   uint32         `protobuf:"varint,3,opt,name=crc16,proto3" json:"crc16,omitempty"`
	Buffer  []byte         `protobuf:"bytes,4,opt,name=buffer,proto3" json:"buffer,omitempty"`
}

func (x *XModem) Reset() {
	*x = XModem{}
	if protoimpl.UnsafeEnabled {
		mi := &file_meshtastic_xmodem_proto_msgTypes[0]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *XModem) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*XModem) ProtoMessage() {}

func (x *XModem) ProtoReflect() protoreflect.Message {
	mi := &file_meshtastic_xmodem_proto_msgTypes[0]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use XModem.ProtoReflect.Descriptor instead.
func (*XModem) Descriptor() ([]byte, []int) {
	return file_meshtastic_xmodem_proto_rawDescGZIP(), []int{0}
}

func (x *XModem) GetControl() XModem_Control {
	if x != nil {
		return x.Control
	}
	return XModem_NUL
}

func (x *XModem) GetSeq() uint32 {
	if x != nil {
		return x.Seq
	}
	return 0
}

func (x *XModem) GetCrc16() uint32 {
	if x != nil {
		return x.Crc16
	}
	return 0
}

func (x *XModem) GetBuffer() []byte {
	if x != nil {
		return x.Buffer
	}
	return nil
}

var File_meshtastic_xmodem_proto protoreflect.FileDescriptor

var file_meshtastic_xmodem_proto_rawDesc = []byte{
	0x0a, 0x17, 0x6d, 0x65, 0x73, 0x68, 0x74, 0x61, 0x73, 0x74, 0x69, 0x63, 0x2f, 0x78, 0x6d, 0x6f,
	0x64, 0x65, 0x6d, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x12, 0x0a, 0x6d, 0x65, 0x73, 0x68, 0x74,
	0x61, 0x73, 0x74, 0x69, 0x63, 0x22, 0xd3, 0x01, 0x0a, 0x06, 0x58, 0x4d, 0x6f, 0x64, 0x65, 0x6d,
	0x12, 0x34, 0x0a, 0x07, 0x63, 0x6f, 0x6e, 0x74, 0x72, 0x6f, 0x6c, 0x18, 0x01, 0x20, 0x01, 0x28,
	0x0e, 0x32, 0x1a, 0x2e, 0x6d, 0x65, 0x73, 0x68, 0x74, 0x61, 0x73, 0x74, 0x69, 0x63, 0x2e, 0x58,
	0x4d, 0x6f, 0x64, 0x65, 0x6d, 0x2e, 0x43, 0x6f, 0x6e, 0x74, 0x72, 0x6f, 0x6c, 0x52, 0x07, 0x63,
	0x6f, 0x6e, 0x74, 0x72, 0x6f, 0x6c, 0x12, 0x10, 0x0a, 0x03, 0x73, 0x65, 0x71, 0x18, 0x02, 0x20,
	0x01, 0x28, 0x0d, 0x52, 0x03, 0x73, 0x65, 0x71, 0x12, 0x14, 0x0a, 0x05, 0x63, 0x72, 0x63, 0x31,
	0x36, 0x18, 0x03, 0x20, 0x01, 0x28, 0x0d, 0x52, 0x05, 0x63, 0x72, 0x63, 0x31, 0x36, 0x12, 0x16,
	0x0a, 0x06, 0x62, 0x75, 0x66, 0x66, 0x65, 0x72, 0x18, 0x04, 0x20, 0x01, 0x28, 0x0c, 0x52, 0x06,
	0x62, 0x75, 0x66, 0x66, 0x65, 0x72, 0x22, 0x53, 0x0a, 0x07, 0x43, 0x6f, 0x6e, 0x74, 0x72, 0x6f,
	0x6c, 0x12, 0x07, 0x0a, 0x03, 0x4e, 0x55, 0x4c, 0x10, 0x00, 0x12, 0x07, 0x0a, 0x03, 0x53, 0x4f,
	0x48, 0x10, 0x01, 0x12, 0x07, 0x0a, 0x03, 0x53, 0x54, 0x58, 0x10, 0x02, 0x12, 0x07, 0x0a, 0x03,
	0x45, 0x4f, 0x54, 0x10, 0x04, 0x12, 0x07, 0x0a, 0x03, 0x41, 0x43, 0x4b, 0x10, 0x06, 0x12, 0x07,
	0x0a, 0x03, 0x4e, 0x41, 0x4b, 0x10, 0x15, 0x12, 0x07, 0x0a, 0x03, 0x43, 0x41, 0x4e, 0x10, 0x18,
	0x12, 0x09, 0x0a, 0x05, 0x43, 0x54, 0x52, 0x4c, 0x5a, 0x10, 0x1a, 0x42, 0x61, 0x0a, 0x13, 0x63,
	0x6f, 0x6d, 0x2e, 0x67, 0x65, 0x65, 0x6b, 0x73, 0x76, 0x69, 0x6c, 0x6c, 0x65, 0x2e, 0x6d, 0x65,
	0x73, 0x68, 0x42, 0x0c, 0x58, 0x6d, 0x6f, 0x64, 0x65, 0x6d, 0x50, 0x72, 0x6f, 0x74, 0x6f, 0x73,
	0x5a, 0x22, 0x67, 0x69, 0x74, 0x68, 0x75, 0x62, 0x2e, 0x63, 0x6f, 0x6d, 0x2f, 0x6d, 0x65, 0x73,
	0x68, 0x74, 0x61, 0x73, 0x74, 0x69, 0x63, 0x2f, 0x67, 0x6f, 0x2f, 0x67, 0x65, 0x6e, 0x65, 0x72,
	0x61, 0x74, 0x65, 0x64, 0xaa, 0x02, 0x14, 0x4d, 0x65, 0x73, 0x68, 0x74, 0x61, 0x73, 0x74, 0x69,
	0x63, 0x2e, 0x50, 0x72, 0x6f, 0x74, 0x6f, 0x62, 0x75, 0x66, 0x73, 0xba, 0x02, 0x00, 0x62, 0x06,
	0x70, 0x72, 0x6f, 0x74, 0x6f, 0x33,
}

var (
	file_meshtastic_xmodem_proto_rawDescOnce sync.Once
	file_meshtastic_xmodem_proto_rawDescData = file_meshtastic_xmodem_proto_rawDesc
)

func file_meshtastic_xmodem_proto_rawDescGZIP() []byte {
	file_meshtastic_xmodem_proto_rawDescOnce.Do(func() {
		file_meshtastic_xmodem_proto_rawDescData = protoimpl.X.CompressGZIP(file_meshtastic_xmodem_proto_rawDescData)
	})
	return file_meshtastic_xmodem_proto_rawDescData
}

var file_meshtastic_xmodem_proto_enumTypes = make([]protoimpl.EnumInfo, 1)
var file_meshtastic_xmodem_proto_msgTypes = make([]protoimpl.MessageInfo, 1)
var file_meshtastic_xmodem_proto_goTypes = []interface{}{
	(XModem_Control)(0), // 0: meshtastic.XModem.Control
	(*XModem)(nil),      // 1: meshtastic.XModem
}
var file_meshtastic_xmodem_proto_depIdxs = []int32{
	0, // 0: meshtastic.XModem.control:type_name -> meshtastic.XModem.Control
	1, // [1:1] is the sub-list for method output_type
	1, // [1:1] is the sub-list for method input_type
	1, // [1:1] is the sub-list for extension type_name
	1, // [1:1] is the sub-list for extension extendee
	0, // [0:1] is the sub-list for field type_name
}

func init() { file_meshtastic_xmodem_proto_init() }
func file_meshtastic_xmodem_proto_init() {
	if File_meshtastic_xmodem_proto != nil {
		return
	}
	if !protoimpl.UnsafeEnabled {
		file_meshtastic_xmodem_proto_msgTypes[0].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*XModem); i {
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
			RawDescriptor: file_meshtastic_xmodem_proto_rawDesc,
			NumEnums:      1,
			NumMessages:   1,
			NumExtensions: 0,
			NumServices:   0,
		},
		GoTypes:           file_meshtastic_xmodem_proto_goTypes,
		DependencyIndexes: file_meshtastic_xmodem_proto_depIdxs,
		EnumInfos:         file_meshtastic_xmodem_proto_enumTypes,
		MessageInfos:      file_meshtastic_xmodem_proto_msgTypes,
	}.Build()
	File_meshtastic_xmodem_proto = out.File
	file_meshtastic_xmodem_proto_rawDesc = nil
	file_meshtastic_xmodem_proto_goTypes = nil
	file_meshtastic_xmodem_proto_depIdxs = nil
}
