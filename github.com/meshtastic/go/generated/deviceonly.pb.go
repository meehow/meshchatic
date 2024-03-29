// Code generated by protoc-gen-go. DO NOT EDIT.
// versions:
// 	protoc-gen-go v1.30.0
// 	protoc        v3.21.12
// source: meshtastic/deviceonly.proto

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

// TODO: REPLACE
type ScreenFonts int32

const (
	// TODO: REPLACE
	ScreenFonts_FONT_SMALL ScreenFonts = 0
	// TODO: REPLACE
	ScreenFonts_FONT_MEDIUM ScreenFonts = 1
	// TODO: REPLACE
	ScreenFonts_FONT_LARGE ScreenFonts = 2
)

// Enum value maps for ScreenFonts.
var (
	ScreenFonts_name = map[int32]string{
		0: "FONT_SMALL",
		1: "FONT_MEDIUM",
		2: "FONT_LARGE",
	}
	ScreenFonts_value = map[string]int32{
		"FONT_SMALL":  0,
		"FONT_MEDIUM": 1,
		"FONT_LARGE":  2,
	}
)

func (x ScreenFonts) Enum() *ScreenFonts {
	p := new(ScreenFonts)
	*p = x
	return p
}

func (x ScreenFonts) String() string {
	return protoimpl.X.EnumStringOf(x.Descriptor(), protoreflect.EnumNumber(x))
}

func (ScreenFonts) Descriptor() protoreflect.EnumDescriptor {
	return file_meshtastic_deviceonly_proto_enumTypes[0].Descriptor()
}

func (ScreenFonts) Type() protoreflect.EnumType {
	return &file_meshtastic_deviceonly_proto_enumTypes[0]
}

func (x ScreenFonts) Number() protoreflect.EnumNumber {
	return protoreflect.EnumNumber(x)
}

// Deprecated: Use ScreenFonts.Descriptor instead.
func (ScreenFonts) EnumDescriptor() ([]byte, []int) {
	return file_meshtastic_deviceonly_proto_rawDescGZIP(), []int{0}
}

// This message is never sent over the wire, but it is used for serializing DB
// state to flash in the device code
// FIXME, since we write this each time we enter deep sleep (and have infinite
// flash) it would be better to use some sort of append only data structure for
// the receive queue and use the preferences store for the other stuff
type DeviceState struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	// Read only settings/info about this node
	MyNode *MyNodeInfo `protobuf:"bytes,2,opt,name=my_node,json=myNode,proto3" json:"my_node,omitempty"`
	// My owner info
	Owner *User `protobuf:"bytes,3,opt,name=owner,proto3" json:"owner,omitempty"`
	// TODO: REPLACE
	NodeDb []*NodeInfo `protobuf:"bytes,4,rep,name=node_db,json=nodeDb,proto3" json:"node_db,omitempty"`
	// Received packets saved for delivery to the phone
	ReceiveQueue []*MeshPacket `protobuf:"bytes,5,rep,name=receive_queue,json=receiveQueue,proto3" json:"receive_queue,omitempty"`
	// A version integer used to invalidate old save files when we make
	// incompatible changes This integer is set at build time and is private to
	// NodeDB.cpp in the device code.
	Version uint32 `protobuf:"varint,8,opt,name=version,proto3" json:"version,omitempty"`
	// We keep the last received text message (only) stored in the device flash,
	// so we can show it on the screen.
	// Might be null
	RxTextMessage *MeshPacket `protobuf:"bytes,7,opt,name=rx_text_message,json=rxTextMessage,proto3" json:"rx_text_message,omitempty"`
	// Used only during development.
	// Indicates developer is testing and changes should never be saved to flash.
	NoSave bool `protobuf:"varint,9,opt,name=no_save,json=noSave,proto3" json:"no_save,omitempty"`
	// Some GPS receivers seem to have bogus settings from the factory, so we always do one factory reset.
	DidGpsReset bool `protobuf:"varint,11,opt,name=did_gps_reset,json=didGpsReset,proto3" json:"did_gps_reset,omitempty"`
	// We keep the last received waypoint stored in the device flash,
	// so we can show it on the screen.
	// Might be null
	RxWaypoint *MeshPacket `protobuf:"bytes,12,opt,name=rx_waypoint,json=rxWaypoint,proto3" json:"rx_waypoint,omitempty"`
	// The mesh's nodes with their available gpio pins for RemoteHardware module
	NodeRemoteHardwarePins []*NodeRemoteHardwarePin `protobuf:"bytes,13,rep,name=node_remote_hardware_pins,json=nodeRemoteHardwarePins,proto3" json:"node_remote_hardware_pins,omitempty"`
}

func (x *DeviceState) Reset() {
	*x = DeviceState{}
	if protoimpl.UnsafeEnabled {
		mi := &file_meshtastic_deviceonly_proto_msgTypes[0]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *DeviceState) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*DeviceState) ProtoMessage() {}

func (x *DeviceState) ProtoReflect() protoreflect.Message {
	mi := &file_meshtastic_deviceonly_proto_msgTypes[0]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use DeviceState.ProtoReflect.Descriptor instead.
func (*DeviceState) Descriptor() ([]byte, []int) {
	return file_meshtastic_deviceonly_proto_rawDescGZIP(), []int{0}
}

func (x *DeviceState) GetMyNode() *MyNodeInfo {
	if x != nil {
		return x.MyNode
	}
	return nil
}

func (x *DeviceState) GetOwner() *User {
	if x != nil {
		return x.Owner
	}
	return nil
}

func (x *DeviceState) GetNodeDb() []*NodeInfo {
	if x != nil {
		return x.NodeDb
	}
	return nil
}

func (x *DeviceState) GetReceiveQueue() []*MeshPacket {
	if x != nil {
		return x.ReceiveQueue
	}
	return nil
}

func (x *DeviceState) GetVersion() uint32 {
	if x != nil {
		return x.Version
	}
	return 0
}

func (x *DeviceState) GetRxTextMessage() *MeshPacket {
	if x != nil {
		return x.RxTextMessage
	}
	return nil
}

func (x *DeviceState) GetNoSave() bool {
	if x != nil {
		return x.NoSave
	}
	return false
}

func (x *DeviceState) GetDidGpsReset() bool {
	if x != nil {
		return x.DidGpsReset
	}
	return false
}

func (x *DeviceState) GetRxWaypoint() *MeshPacket {
	if x != nil {
		return x.RxWaypoint
	}
	return nil
}

func (x *DeviceState) GetNodeRemoteHardwarePins() []*NodeRemoteHardwarePin {
	if x != nil {
		return x.NodeRemoteHardwarePins
	}
	return nil
}

// The on-disk saved channels
type ChannelFile struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	// The channels our node knows about
	Channels []*Channel `protobuf:"bytes,1,rep,name=channels,proto3" json:"channels,omitempty"`
	// A version integer used to invalidate old save files when we make
	// incompatible changes This integer is set at build time and is private to
	// NodeDB.cpp in the device code.
	Version uint32 `protobuf:"varint,2,opt,name=version,proto3" json:"version,omitempty"`
}

func (x *ChannelFile) Reset() {
	*x = ChannelFile{}
	if protoimpl.UnsafeEnabled {
		mi := &file_meshtastic_deviceonly_proto_msgTypes[1]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *ChannelFile) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*ChannelFile) ProtoMessage() {}

func (x *ChannelFile) ProtoReflect() protoreflect.Message {
	mi := &file_meshtastic_deviceonly_proto_msgTypes[1]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use ChannelFile.ProtoReflect.Descriptor instead.
func (*ChannelFile) Descriptor() ([]byte, []int) {
	return file_meshtastic_deviceonly_proto_rawDescGZIP(), []int{1}
}

func (x *ChannelFile) GetChannels() []*Channel {
	if x != nil {
		return x.Channels
	}
	return nil
}

func (x *ChannelFile) GetVersion() uint32 {
	if x != nil {
		return x.Version
	}
	return 0
}

// This can be used for customizing the firmware distribution. If populated,
// show a secondary bootup screen with custom logo and text for 2.5 seconds.
type OEMStore struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	// The Logo width in Px
	OemIconWidth uint32 `protobuf:"varint,1,opt,name=oem_icon_width,json=oemIconWidth,proto3" json:"oem_icon_width,omitempty"`
	// The Logo height in Px
	OemIconHeight uint32 `protobuf:"varint,2,opt,name=oem_icon_height,json=oemIconHeight,proto3" json:"oem_icon_height,omitempty"`
	// The Logo in XBM bytechar format
	OemIconBits []byte `protobuf:"bytes,3,opt,name=oem_icon_bits,json=oemIconBits,proto3" json:"oem_icon_bits,omitempty"`
	// Use this font for the OEM text.
	OemFont ScreenFonts `protobuf:"varint,4,opt,name=oem_font,json=oemFont,proto3,enum=meshtastic.ScreenFonts" json:"oem_font,omitempty"`
	// Use this font for the OEM text.
	OemText string `protobuf:"bytes,5,opt,name=oem_text,json=oemText,proto3" json:"oem_text,omitempty"`
	// The default device encryption key, 16 or 32 byte
	OemAesKey []byte `protobuf:"bytes,6,opt,name=oem_aes_key,json=oemAesKey,proto3" json:"oem_aes_key,omitempty"`
	// A Preset LocalConfig to apply during factory reset
	OemLocalConfig *LocalConfig `protobuf:"bytes,7,opt,name=oem_local_config,json=oemLocalConfig,proto3" json:"oem_local_config,omitempty"`
	// A Preset LocalModuleConfig to apply during factory reset
	OemLocalModuleConfig *LocalModuleConfig `protobuf:"bytes,8,opt,name=oem_local_module_config,json=oemLocalModuleConfig,proto3" json:"oem_local_module_config,omitempty"`
}

func (x *OEMStore) Reset() {
	*x = OEMStore{}
	if protoimpl.UnsafeEnabled {
		mi := &file_meshtastic_deviceonly_proto_msgTypes[2]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *OEMStore) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*OEMStore) ProtoMessage() {}

func (x *OEMStore) ProtoReflect() protoreflect.Message {
	mi := &file_meshtastic_deviceonly_proto_msgTypes[2]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use OEMStore.ProtoReflect.Descriptor instead.
func (*OEMStore) Descriptor() ([]byte, []int) {
	return file_meshtastic_deviceonly_proto_rawDescGZIP(), []int{2}
}

func (x *OEMStore) GetOemIconWidth() uint32 {
	if x != nil {
		return x.OemIconWidth
	}
	return 0
}

func (x *OEMStore) GetOemIconHeight() uint32 {
	if x != nil {
		return x.OemIconHeight
	}
	return 0
}

func (x *OEMStore) GetOemIconBits() []byte {
	if x != nil {
		return x.OemIconBits
	}
	return nil
}

func (x *OEMStore) GetOemFont() ScreenFonts {
	if x != nil {
		return x.OemFont
	}
	return ScreenFonts_FONT_SMALL
}

func (x *OEMStore) GetOemText() string {
	if x != nil {
		return x.OemText
	}
	return ""
}

func (x *OEMStore) GetOemAesKey() []byte {
	if x != nil {
		return x.OemAesKey
	}
	return nil
}

func (x *OEMStore) GetOemLocalConfig() *LocalConfig {
	if x != nil {
		return x.OemLocalConfig
	}
	return nil
}

func (x *OEMStore) GetOemLocalModuleConfig() *LocalModuleConfig {
	if x != nil {
		return x.OemLocalModuleConfig
	}
	return nil
}

// RemoteHardwarePins associated with a node
type NodeRemoteHardwarePin struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	// The node_num exposing the available gpio pin
	NodeNum uint32 `protobuf:"varint,1,opt,name=node_num,json=nodeNum,proto3" json:"node_num,omitempty"`
	// The the available gpio pin for usage with RemoteHardware module
	Pin *RemoteHardwarePin `protobuf:"bytes,2,opt,name=pin,proto3" json:"pin,omitempty"`
}

func (x *NodeRemoteHardwarePin) Reset() {
	*x = NodeRemoteHardwarePin{}
	if protoimpl.UnsafeEnabled {
		mi := &file_meshtastic_deviceonly_proto_msgTypes[3]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *NodeRemoteHardwarePin) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*NodeRemoteHardwarePin) ProtoMessage() {}

func (x *NodeRemoteHardwarePin) ProtoReflect() protoreflect.Message {
	mi := &file_meshtastic_deviceonly_proto_msgTypes[3]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use NodeRemoteHardwarePin.ProtoReflect.Descriptor instead.
func (*NodeRemoteHardwarePin) Descriptor() ([]byte, []int) {
	return file_meshtastic_deviceonly_proto_rawDescGZIP(), []int{3}
}

func (x *NodeRemoteHardwarePin) GetNodeNum() uint32 {
	if x != nil {
		return x.NodeNum
	}
	return 0
}

func (x *NodeRemoteHardwarePin) GetPin() *RemoteHardwarePin {
	if x != nil {
		return x.Pin
	}
	return nil
}

var File_meshtastic_deviceonly_proto protoreflect.FileDescriptor

var file_meshtastic_deviceonly_proto_rawDesc = []byte{
	0x0a, 0x1b, 0x6d, 0x65, 0x73, 0x68, 0x74, 0x61, 0x73, 0x74, 0x69, 0x63, 0x2f, 0x64, 0x65, 0x76,
	0x69, 0x63, 0x65, 0x6f, 0x6e, 0x6c, 0x79, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x12, 0x0a, 0x6d,
	0x65, 0x73, 0x68, 0x74, 0x61, 0x73, 0x74, 0x69, 0x63, 0x1a, 0x18, 0x6d, 0x65, 0x73, 0x68, 0x74,
	0x61, 0x73, 0x74, 0x69, 0x63, 0x2f, 0x63, 0x68, 0x61, 0x6e, 0x6e, 0x65, 0x6c, 0x2e, 0x70, 0x72,
	0x6f, 0x74, 0x6f, 0x1a, 0x1a, 0x6d, 0x65, 0x73, 0x68, 0x74, 0x61, 0x73, 0x74, 0x69, 0x63, 0x2f,
	0x6c, 0x6f, 0x63, 0x61, 0x6c, 0x6f, 0x6e, 0x6c, 0x79, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x1a,
	0x15, 0x6d, 0x65, 0x73, 0x68, 0x74, 0x61, 0x73, 0x74, 0x69, 0x63, 0x2f, 0x6d, 0x65, 0x73, 0x68,
	0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x1a, 0x1e, 0x6d, 0x65, 0x73, 0x68, 0x74, 0x61, 0x73, 0x74,
	0x69, 0x63, 0x2f, 0x6d, 0x6f, 0x64, 0x75, 0x6c, 0x65, 0x5f, 0x63, 0x6f, 0x6e, 0x66, 0x69, 0x67,
	0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x22, 0x80, 0x04, 0x0a, 0x0b, 0x44, 0x65, 0x76, 0x69, 0x63,
	0x65, 0x53, 0x74, 0x61, 0x74, 0x65, 0x12, 0x2f, 0x0a, 0x07, 0x6d, 0x79, 0x5f, 0x6e, 0x6f, 0x64,
	0x65, 0x18, 0x02, 0x20, 0x01, 0x28, 0x0b, 0x32, 0x16, 0x2e, 0x6d, 0x65, 0x73, 0x68, 0x74, 0x61,
	0x73, 0x74, 0x69, 0x63, 0x2e, 0x4d, 0x79, 0x4e, 0x6f, 0x64, 0x65, 0x49, 0x6e, 0x66, 0x6f, 0x52,
	0x06, 0x6d, 0x79, 0x4e, 0x6f, 0x64, 0x65, 0x12, 0x26, 0x0a, 0x05, 0x6f, 0x77, 0x6e, 0x65, 0x72,
	0x18, 0x03, 0x20, 0x01, 0x28, 0x0b, 0x32, 0x10, 0x2e, 0x6d, 0x65, 0x73, 0x68, 0x74, 0x61, 0x73,
	0x74, 0x69, 0x63, 0x2e, 0x55, 0x73, 0x65, 0x72, 0x52, 0x05, 0x6f, 0x77, 0x6e, 0x65, 0x72, 0x12,
	0x2d, 0x0a, 0x07, 0x6e, 0x6f, 0x64, 0x65, 0x5f, 0x64, 0x62, 0x18, 0x04, 0x20, 0x03, 0x28, 0x0b,
	0x32, 0x14, 0x2e, 0x6d, 0x65, 0x73, 0x68, 0x74, 0x61, 0x73, 0x74, 0x69, 0x63, 0x2e, 0x4e, 0x6f,
	0x64, 0x65, 0x49, 0x6e, 0x66, 0x6f, 0x52, 0x06, 0x6e, 0x6f, 0x64, 0x65, 0x44, 0x62, 0x12, 0x3b,
	0x0a, 0x0d, 0x72, 0x65, 0x63, 0x65, 0x69, 0x76, 0x65, 0x5f, 0x71, 0x75, 0x65, 0x75, 0x65, 0x18,
	0x05, 0x20, 0x03, 0x28, 0x0b, 0x32, 0x16, 0x2e, 0x6d, 0x65, 0x73, 0x68, 0x74, 0x61, 0x73, 0x74,
	0x69, 0x63, 0x2e, 0x4d, 0x65, 0x73, 0x68, 0x50, 0x61, 0x63, 0x6b, 0x65, 0x74, 0x52, 0x0c, 0x72,
	0x65, 0x63, 0x65, 0x69, 0x76, 0x65, 0x51, 0x75, 0x65, 0x75, 0x65, 0x12, 0x18, 0x0a, 0x07, 0x76,
	0x65, 0x72, 0x73, 0x69, 0x6f, 0x6e, 0x18, 0x08, 0x20, 0x01, 0x28, 0x0d, 0x52, 0x07, 0x76, 0x65,
	0x72, 0x73, 0x69, 0x6f, 0x6e, 0x12, 0x3e, 0x0a, 0x0f, 0x72, 0x78, 0x5f, 0x74, 0x65, 0x78, 0x74,
	0x5f, 0x6d, 0x65, 0x73, 0x73, 0x61, 0x67, 0x65, 0x18, 0x07, 0x20, 0x01, 0x28, 0x0b, 0x32, 0x16,
	0x2e, 0x6d, 0x65, 0x73, 0x68, 0x74, 0x61, 0x73, 0x74, 0x69, 0x63, 0x2e, 0x4d, 0x65, 0x73, 0x68,
	0x50, 0x61, 0x63, 0x6b, 0x65, 0x74, 0x52, 0x0d, 0x72, 0x78, 0x54, 0x65, 0x78, 0x74, 0x4d, 0x65,
	0x73, 0x73, 0x61, 0x67, 0x65, 0x12, 0x17, 0x0a, 0x07, 0x6e, 0x6f, 0x5f, 0x73, 0x61, 0x76, 0x65,
	0x18, 0x09, 0x20, 0x01, 0x28, 0x08, 0x52, 0x06, 0x6e, 0x6f, 0x53, 0x61, 0x76, 0x65, 0x12, 0x22,
	0x0a, 0x0d, 0x64, 0x69, 0x64, 0x5f, 0x67, 0x70, 0x73, 0x5f, 0x72, 0x65, 0x73, 0x65, 0x74, 0x18,
	0x0b, 0x20, 0x01, 0x28, 0x08, 0x52, 0x0b, 0x64, 0x69, 0x64, 0x47, 0x70, 0x73, 0x52, 0x65, 0x73,
	0x65, 0x74, 0x12, 0x37, 0x0a, 0x0b, 0x72, 0x78, 0x5f, 0x77, 0x61, 0x79, 0x70, 0x6f, 0x69, 0x6e,
	0x74, 0x18, 0x0c, 0x20, 0x01, 0x28, 0x0b, 0x32, 0x16, 0x2e, 0x6d, 0x65, 0x73, 0x68, 0x74, 0x61,
	0x73, 0x74, 0x69, 0x63, 0x2e, 0x4d, 0x65, 0x73, 0x68, 0x50, 0x61, 0x63, 0x6b, 0x65, 0x74, 0x52,
	0x0a, 0x72, 0x78, 0x57, 0x61, 0x79, 0x70, 0x6f, 0x69, 0x6e, 0x74, 0x12, 0x5c, 0x0a, 0x19, 0x6e,
	0x6f, 0x64, 0x65, 0x5f, 0x72, 0x65, 0x6d, 0x6f, 0x74, 0x65, 0x5f, 0x68, 0x61, 0x72, 0x64, 0x77,
	0x61, 0x72, 0x65, 0x5f, 0x70, 0x69, 0x6e, 0x73, 0x18, 0x0d, 0x20, 0x03, 0x28, 0x0b, 0x32, 0x21,
	0x2e, 0x6d, 0x65, 0x73, 0x68, 0x74, 0x61, 0x73, 0x74, 0x69, 0x63, 0x2e, 0x4e, 0x6f, 0x64, 0x65,
	0x52, 0x65, 0x6d, 0x6f, 0x74, 0x65, 0x48, 0x61, 0x72, 0x64, 0x77, 0x61, 0x72, 0x65, 0x50, 0x69,
	0x6e, 0x52, 0x16, 0x6e, 0x6f, 0x64, 0x65, 0x52, 0x65, 0x6d, 0x6f, 0x74, 0x65, 0x48, 0x61, 0x72,
	0x64, 0x77, 0x61, 0x72, 0x65, 0x50, 0x69, 0x6e, 0x73, 0x22, 0x58, 0x0a, 0x0b, 0x43, 0x68, 0x61,
	0x6e, 0x6e, 0x65, 0x6c, 0x46, 0x69, 0x6c, 0x65, 0x12, 0x2f, 0x0a, 0x08, 0x63, 0x68, 0x61, 0x6e,
	0x6e, 0x65, 0x6c, 0x73, 0x18, 0x01, 0x20, 0x03, 0x28, 0x0b, 0x32, 0x13, 0x2e, 0x6d, 0x65, 0x73,
	0x68, 0x74, 0x61, 0x73, 0x74, 0x69, 0x63, 0x2e, 0x43, 0x68, 0x61, 0x6e, 0x6e, 0x65, 0x6c, 0x52,
	0x08, 0x63, 0x68, 0x61, 0x6e, 0x6e, 0x65, 0x6c, 0x73, 0x12, 0x18, 0x0a, 0x07, 0x76, 0x65, 0x72,
	0x73, 0x69, 0x6f, 0x6e, 0x18, 0x02, 0x20, 0x01, 0x28, 0x0d, 0x52, 0x07, 0x76, 0x65, 0x72, 0x73,
	0x69, 0x6f, 0x6e, 0x22, 0x84, 0x03, 0x0a, 0x08, 0x4f, 0x45, 0x4d, 0x53, 0x74, 0x6f, 0x72, 0x65,
	0x12, 0x24, 0x0a, 0x0e, 0x6f, 0x65, 0x6d, 0x5f, 0x69, 0x63, 0x6f, 0x6e, 0x5f, 0x77, 0x69, 0x64,
	0x74, 0x68, 0x18, 0x01, 0x20, 0x01, 0x28, 0x0d, 0x52, 0x0c, 0x6f, 0x65, 0x6d, 0x49, 0x63, 0x6f,
	0x6e, 0x57, 0x69, 0x64, 0x74, 0x68, 0x12, 0x26, 0x0a, 0x0f, 0x6f, 0x65, 0x6d, 0x5f, 0x69, 0x63,
	0x6f, 0x6e, 0x5f, 0x68, 0x65, 0x69, 0x67, 0x68, 0x74, 0x18, 0x02, 0x20, 0x01, 0x28, 0x0d, 0x52,
	0x0d, 0x6f, 0x65, 0x6d, 0x49, 0x63, 0x6f, 0x6e, 0x48, 0x65, 0x69, 0x67, 0x68, 0x74, 0x12, 0x22,
	0x0a, 0x0d, 0x6f, 0x65, 0x6d, 0x5f, 0x69, 0x63, 0x6f, 0x6e, 0x5f, 0x62, 0x69, 0x74, 0x73, 0x18,
	0x03, 0x20, 0x01, 0x28, 0x0c, 0x52, 0x0b, 0x6f, 0x65, 0x6d, 0x49, 0x63, 0x6f, 0x6e, 0x42, 0x69,
	0x74, 0x73, 0x12, 0x32, 0x0a, 0x08, 0x6f, 0x65, 0x6d, 0x5f, 0x66, 0x6f, 0x6e, 0x74, 0x18, 0x04,
	0x20, 0x01, 0x28, 0x0e, 0x32, 0x17, 0x2e, 0x6d, 0x65, 0x73, 0x68, 0x74, 0x61, 0x73, 0x74, 0x69,
	0x63, 0x2e, 0x53, 0x63, 0x72, 0x65, 0x65, 0x6e, 0x46, 0x6f, 0x6e, 0x74, 0x73, 0x52, 0x07, 0x6f,
	0x65, 0x6d, 0x46, 0x6f, 0x6e, 0x74, 0x12, 0x19, 0x0a, 0x08, 0x6f, 0x65, 0x6d, 0x5f, 0x74, 0x65,
	0x78, 0x74, 0x18, 0x05, 0x20, 0x01, 0x28, 0x09, 0x52, 0x07, 0x6f, 0x65, 0x6d, 0x54, 0x65, 0x78,
	0x74, 0x12, 0x1e, 0x0a, 0x0b, 0x6f, 0x65, 0x6d, 0x5f, 0x61, 0x65, 0x73, 0x5f, 0x6b, 0x65, 0x79,
	0x18, 0x06, 0x20, 0x01, 0x28, 0x0c, 0x52, 0x09, 0x6f, 0x65, 0x6d, 0x41, 0x65, 0x73, 0x4b, 0x65,
	0x79, 0x12, 0x41, 0x0a, 0x10, 0x6f, 0x65, 0x6d, 0x5f, 0x6c, 0x6f, 0x63, 0x61, 0x6c, 0x5f, 0x63,
	0x6f, 0x6e, 0x66, 0x69, 0x67, 0x18, 0x07, 0x20, 0x01, 0x28, 0x0b, 0x32, 0x17, 0x2e, 0x6d, 0x65,
	0x73, 0x68, 0x74, 0x61, 0x73, 0x74, 0x69, 0x63, 0x2e, 0x4c, 0x6f, 0x63, 0x61, 0x6c, 0x43, 0x6f,
	0x6e, 0x66, 0x69, 0x67, 0x52, 0x0e, 0x6f, 0x65, 0x6d, 0x4c, 0x6f, 0x63, 0x61, 0x6c, 0x43, 0x6f,
	0x6e, 0x66, 0x69, 0x67, 0x12, 0x54, 0x0a, 0x17, 0x6f, 0x65, 0x6d, 0x5f, 0x6c, 0x6f, 0x63, 0x61,
	0x6c, 0x5f, 0x6d, 0x6f, 0x64, 0x75, 0x6c, 0x65, 0x5f, 0x63, 0x6f, 0x6e, 0x66, 0x69, 0x67, 0x18,
	0x08, 0x20, 0x01, 0x28, 0x0b, 0x32, 0x1d, 0x2e, 0x6d, 0x65, 0x73, 0x68, 0x74, 0x61, 0x73, 0x74,
	0x69, 0x63, 0x2e, 0x4c, 0x6f, 0x63, 0x61, 0x6c, 0x4d, 0x6f, 0x64, 0x75, 0x6c, 0x65, 0x43, 0x6f,
	0x6e, 0x66, 0x69, 0x67, 0x52, 0x14, 0x6f, 0x65, 0x6d, 0x4c, 0x6f, 0x63, 0x61, 0x6c, 0x4d, 0x6f,
	0x64, 0x75, 0x6c, 0x65, 0x43, 0x6f, 0x6e, 0x66, 0x69, 0x67, 0x22, 0x63, 0x0a, 0x15, 0x4e, 0x6f,
	0x64, 0x65, 0x52, 0x65, 0x6d, 0x6f, 0x74, 0x65, 0x48, 0x61, 0x72, 0x64, 0x77, 0x61, 0x72, 0x65,
	0x50, 0x69, 0x6e, 0x12, 0x19, 0x0a, 0x08, 0x6e, 0x6f, 0x64, 0x65, 0x5f, 0x6e, 0x75, 0x6d, 0x18,
	0x01, 0x20, 0x01, 0x28, 0x0d, 0x52, 0x07, 0x6e, 0x6f, 0x64, 0x65, 0x4e, 0x75, 0x6d, 0x12, 0x2f,
	0x0a, 0x03, 0x70, 0x69, 0x6e, 0x18, 0x02, 0x20, 0x01, 0x28, 0x0b, 0x32, 0x1d, 0x2e, 0x6d, 0x65,
	0x73, 0x68, 0x74, 0x61, 0x73, 0x74, 0x69, 0x63, 0x2e, 0x52, 0x65, 0x6d, 0x6f, 0x74, 0x65, 0x48,
	0x61, 0x72, 0x64, 0x77, 0x61, 0x72, 0x65, 0x50, 0x69, 0x6e, 0x52, 0x03, 0x70, 0x69, 0x6e, 0x2a,
	0x3e, 0x0a, 0x0b, 0x53, 0x63, 0x72, 0x65, 0x65, 0x6e, 0x46, 0x6f, 0x6e, 0x74, 0x73, 0x12, 0x0e,
	0x0a, 0x0a, 0x46, 0x4f, 0x4e, 0x54, 0x5f, 0x53, 0x4d, 0x41, 0x4c, 0x4c, 0x10, 0x00, 0x12, 0x0f,
	0x0a, 0x0b, 0x46, 0x4f, 0x4e, 0x54, 0x5f, 0x4d, 0x45, 0x44, 0x49, 0x55, 0x4d, 0x10, 0x01, 0x12,
	0x0e, 0x0a, 0x0a, 0x46, 0x4f, 0x4e, 0x54, 0x5f, 0x4c, 0x41, 0x52, 0x47, 0x45, 0x10, 0x02, 0x42,
	0x5f, 0x0a, 0x13, 0x63, 0x6f, 0x6d, 0x2e, 0x67, 0x65, 0x65, 0x6b, 0x73, 0x76, 0x69, 0x6c, 0x6c,
	0x65, 0x2e, 0x6d, 0x65, 0x73, 0x68, 0x42, 0x0a, 0x44, 0x65, 0x76, 0x69, 0x63, 0x65, 0x4f, 0x6e,
	0x6c, 0x79, 0x5a, 0x22, 0x67, 0x69, 0x74, 0x68, 0x75, 0x62, 0x2e, 0x63, 0x6f, 0x6d, 0x2f, 0x6d,
	0x65, 0x73, 0x68, 0x74, 0x61, 0x73, 0x74, 0x69, 0x63, 0x2f, 0x67, 0x6f, 0x2f, 0x67, 0x65, 0x6e,
	0x65, 0x72, 0x61, 0x74, 0x65, 0x64, 0xaa, 0x02, 0x14, 0x4d, 0x65, 0x73, 0x68, 0x74, 0x61, 0x73,
	0x74, 0x69, 0x63, 0x2e, 0x50, 0x72, 0x6f, 0x74, 0x6f, 0x62, 0x75, 0x66, 0x73, 0xba, 0x02, 0x00,
	0x62, 0x06, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x33,
}

var (
	file_meshtastic_deviceonly_proto_rawDescOnce sync.Once
	file_meshtastic_deviceonly_proto_rawDescData = file_meshtastic_deviceonly_proto_rawDesc
)

func file_meshtastic_deviceonly_proto_rawDescGZIP() []byte {
	file_meshtastic_deviceonly_proto_rawDescOnce.Do(func() {
		file_meshtastic_deviceonly_proto_rawDescData = protoimpl.X.CompressGZIP(file_meshtastic_deviceonly_proto_rawDescData)
	})
	return file_meshtastic_deviceonly_proto_rawDescData
}

var file_meshtastic_deviceonly_proto_enumTypes = make([]protoimpl.EnumInfo, 1)
var file_meshtastic_deviceonly_proto_msgTypes = make([]protoimpl.MessageInfo, 4)
var file_meshtastic_deviceonly_proto_goTypes = []interface{}{
	(ScreenFonts)(0),              // 0: meshtastic.ScreenFonts
	(*DeviceState)(nil),           // 1: meshtastic.DeviceState
	(*ChannelFile)(nil),           // 2: meshtastic.ChannelFile
	(*OEMStore)(nil),              // 3: meshtastic.OEMStore
	(*NodeRemoteHardwarePin)(nil), // 4: meshtastic.NodeRemoteHardwarePin
	(*MyNodeInfo)(nil),            // 5: meshtastic.MyNodeInfo
	(*User)(nil),                  // 6: meshtastic.User
	(*NodeInfo)(nil),              // 7: meshtastic.NodeInfo
	(*MeshPacket)(nil),            // 8: meshtastic.MeshPacket
	(*Channel)(nil),               // 9: meshtastic.Channel
	(*LocalConfig)(nil),           // 10: meshtastic.LocalConfig
	(*LocalModuleConfig)(nil),     // 11: meshtastic.LocalModuleConfig
	(*RemoteHardwarePin)(nil),     // 12: meshtastic.RemoteHardwarePin
}
var file_meshtastic_deviceonly_proto_depIdxs = []int32{
	5,  // 0: meshtastic.DeviceState.my_node:type_name -> meshtastic.MyNodeInfo
	6,  // 1: meshtastic.DeviceState.owner:type_name -> meshtastic.User
	7,  // 2: meshtastic.DeviceState.node_db:type_name -> meshtastic.NodeInfo
	8,  // 3: meshtastic.DeviceState.receive_queue:type_name -> meshtastic.MeshPacket
	8,  // 4: meshtastic.DeviceState.rx_text_message:type_name -> meshtastic.MeshPacket
	8,  // 5: meshtastic.DeviceState.rx_waypoint:type_name -> meshtastic.MeshPacket
	4,  // 6: meshtastic.DeviceState.node_remote_hardware_pins:type_name -> meshtastic.NodeRemoteHardwarePin
	9,  // 7: meshtastic.ChannelFile.channels:type_name -> meshtastic.Channel
	0,  // 8: meshtastic.OEMStore.oem_font:type_name -> meshtastic.ScreenFonts
	10, // 9: meshtastic.OEMStore.oem_local_config:type_name -> meshtastic.LocalConfig
	11, // 10: meshtastic.OEMStore.oem_local_module_config:type_name -> meshtastic.LocalModuleConfig
	12, // 11: meshtastic.NodeRemoteHardwarePin.pin:type_name -> meshtastic.RemoteHardwarePin
	12, // [12:12] is the sub-list for method output_type
	12, // [12:12] is the sub-list for method input_type
	12, // [12:12] is the sub-list for extension type_name
	12, // [12:12] is the sub-list for extension extendee
	0,  // [0:12] is the sub-list for field type_name
}

func init() { file_meshtastic_deviceonly_proto_init() }
func file_meshtastic_deviceonly_proto_init() {
	if File_meshtastic_deviceonly_proto != nil {
		return
	}
	file_meshtastic_channel_proto_init()
	file_meshtastic_localonly_proto_init()
	file_meshtastic_mesh_proto_init()
	file_meshtastic_module_config_proto_init()
	if !protoimpl.UnsafeEnabled {
		file_meshtastic_deviceonly_proto_msgTypes[0].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*DeviceState); i {
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
		file_meshtastic_deviceonly_proto_msgTypes[1].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*ChannelFile); i {
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
		file_meshtastic_deviceonly_proto_msgTypes[2].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*OEMStore); i {
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
		file_meshtastic_deviceonly_proto_msgTypes[3].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*NodeRemoteHardwarePin); i {
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
			RawDescriptor: file_meshtastic_deviceonly_proto_rawDesc,
			NumEnums:      1,
			NumMessages:   4,
			NumExtensions: 0,
			NumServices:   0,
		},
		GoTypes:           file_meshtastic_deviceonly_proto_goTypes,
		DependencyIndexes: file_meshtastic_deviceonly_proto_depIdxs,
		EnumInfos:         file_meshtastic_deviceonly_proto_enumTypes,
		MessageInfos:      file_meshtastic_deviceonly_proto_msgTypes,
	}.Build()
	File_meshtastic_deviceonly_proto = out.File
	file_meshtastic_deviceonly_proto_rawDesc = nil
	file_meshtastic_deviceonly_proto_goTypes = nil
	file_meshtastic_deviceonly_proto_depIdxs = nil
}
