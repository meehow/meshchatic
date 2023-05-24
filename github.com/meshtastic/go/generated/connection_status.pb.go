// Code generated by protoc-gen-go. DO NOT EDIT.
// versions:
// 	protoc-gen-go v1.30.0
// 	protoc        v3.21.12
// source: meshtastic/connection_status.proto

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

type DeviceConnectionStatus struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	// WiFi Status
	Wifi *WifiConnectionStatus `protobuf:"bytes,1,opt,name=wifi,proto3,oneof" json:"wifi,omitempty"`
	// WiFi Status
	Ethernet *EthernetConnectionStatus `protobuf:"bytes,2,opt,name=ethernet,proto3,oneof" json:"ethernet,omitempty"`
	// Bluetooth Status
	Bluetooth *BluetoothConnectionStatus `protobuf:"bytes,3,opt,name=bluetooth,proto3,oneof" json:"bluetooth,omitempty"`
	// Serial Status
	Serial *SerialConnectionStatus `protobuf:"bytes,4,opt,name=serial,proto3,oneof" json:"serial,omitempty"`
}

func (x *DeviceConnectionStatus) Reset() {
	*x = DeviceConnectionStatus{}
	if protoimpl.UnsafeEnabled {
		mi := &file_meshtastic_connection_status_proto_msgTypes[0]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *DeviceConnectionStatus) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*DeviceConnectionStatus) ProtoMessage() {}

func (x *DeviceConnectionStatus) ProtoReflect() protoreflect.Message {
	mi := &file_meshtastic_connection_status_proto_msgTypes[0]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use DeviceConnectionStatus.ProtoReflect.Descriptor instead.
func (*DeviceConnectionStatus) Descriptor() ([]byte, []int) {
	return file_meshtastic_connection_status_proto_rawDescGZIP(), []int{0}
}

func (x *DeviceConnectionStatus) GetWifi() *WifiConnectionStatus {
	if x != nil {
		return x.Wifi
	}
	return nil
}

func (x *DeviceConnectionStatus) GetEthernet() *EthernetConnectionStatus {
	if x != nil {
		return x.Ethernet
	}
	return nil
}

func (x *DeviceConnectionStatus) GetBluetooth() *BluetoothConnectionStatus {
	if x != nil {
		return x.Bluetooth
	}
	return nil
}

func (x *DeviceConnectionStatus) GetSerial() *SerialConnectionStatus {
	if x != nil {
		return x.Serial
	}
	return nil
}

// WiFi connection status
type WifiConnectionStatus struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	// Connection status
	Status *NetworkConnectionStatus `protobuf:"bytes,1,opt,name=status,proto3" json:"status,omitempty"`
	// WiFi access point SSID
	Ssid string `protobuf:"bytes,2,opt,name=ssid,proto3" json:"ssid,omitempty"`
	// RSSI of wireless connection
	Rssi int32 `protobuf:"varint,3,opt,name=rssi,proto3" json:"rssi,omitempty"`
}

func (x *WifiConnectionStatus) Reset() {
	*x = WifiConnectionStatus{}
	if protoimpl.UnsafeEnabled {
		mi := &file_meshtastic_connection_status_proto_msgTypes[1]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *WifiConnectionStatus) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*WifiConnectionStatus) ProtoMessage() {}

func (x *WifiConnectionStatus) ProtoReflect() protoreflect.Message {
	mi := &file_meshtastic_connection_status_proto_msgTypes[1]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use WifiConnectionStatus.ProtoReflect.Descriptor instead.
func (*WifiConnectionStatus) Descriptor() ([]byte, []int) {
	return file_meshtastic_connection_status_proto_rawDescGZIP(), []int{1}
}

func (x *WifiConnectionStatus) GetStatus() *NetworkConnectionStatus {
	if x != nil {
		return x.Status
	}
	return nil
}

func (x *WifiConnectionStatus) GetSsid() string {
	if x != nil {
		return x.Ssid
	}
	return ""
}

func (x *WifiConnectionStatus) GetRssi() int32 {
	if x != nil {
		return x.Rssi
	}
	return 0
}

// Ethernet connection status
type EthernetConnectionStatus struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	// Connection status
	Status *NetworkConnectionStatus `protobuf:"bytes,1,opt,name=status,proto3" json:"status,omitempty"`
}

func (x *EthernetConnectionStatus) Reset() {
	*x = EthernetConnectionStatus{}
	if protoimpl.UnsafeEnabled {
		mi := &file_meshtastic_connection_status_proto_msgTypes[2]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *EthernetConnectionStatus) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*EthernetConnectionStatus) ProtoMessage() {}

func (x *EthernetConnectionStatus) ProtoReflect() protoreflect.Message {
	mi := &file_meshtastic_connection_status_proto_msgTypes[2]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use EthernetConnectionStatus.ProtoReflect.Descriptor instead.
func (*EthernetConnectionStatus) Descriptor() ([]byte, []int) {
	return file_meshtastic_connection_status_proto_rawDescGZIP(), []int{2}
}

func (x *EthernetConnectionStatus) GetStatus() *NetworkConnectionStatus {
	if x != nil {
		return x.Status
	}
	return nil
}

// Ethernet or WiFi connection status
type NetworkConnectionStatus struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	// IP address of device
	IpAddress uint32 `protobuf:"fixed32,1,opt,name=ip_address,json=ipAddress,proto3" json:"ip_address,omitempty"`
	// Whether the device has an active connection or not
	IsConnected bool `protobuf:"varint,2,opt,name=is_connected,json=isConnected,proto3" json:"is_connected,omitempty"`
	// Whether the device has an active connection to an MQTT broker or not
	IsMqttConnected bool `protobuf:"varint,3,opt,name=is_mqtt_connected,json=isMqttConnected,proto3" json:"is_mqtt_connected,omitempty"`
	// Whether the device is actively remote syslogging or not
	IsSyslogConnected bool `protobuf:"varint,4,opt,name=is_syslog_connected,json=isSyslogConnected,proto3" json:"is_syslog_connected,omitempty"`
}

func (x *NetworkConnectionStatus) Reset() {
	*x = NetworkConnectionStatus{}
	if protoimpl.UnsafeEnabled {
		mi := &file_meshtastic_connection_status_proto_msgTypes[3]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *NetworkConnectionStatus) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*NetworkConnectionStatus) ProtoMessage() {}

func (x *NetworkConnectionStatus) ProtoReflect() protoreflect.Message {
	mi := &file_meshtastic_connection_status_proto_msgTypes[3]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use NetworkConnectionStatus.ProtoReflect.Descriptor instead.
func (*NetworkConnectionStatus) Descriptor() ([]byte, []int) {
	return file_meshtastic_connection_status_proto_rawDescGZIP(), []int{3}
}

func (x *NetworkConnectionStatus) GetIpAddress() uint32 {
	if x != nil {
		return x.IpAddress
	}
	return 0
}

func (x *NetworkConnectionStatus) GetIsConnected() bool {
	if x != nil {
		return x.IsConnected
	}
	return false
}

func (x *NetworkConnectionStatus) GetIsMqttConnected() bool {
	if x != nil {
		return x.IsMqttConnected
	}
	return false
}

func (x *NetworkConnectionStatus) GetIsSyslogConnected() bool {
	if x != nil {
		return x.IsSyslogConnected
	}
	return false
}

// Bluetooth connection status
type BluetoothConnectionStatus struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	// The pairing PIN for bluetooth
	Pin uint32 `protobuf:"varint,1,opt,name=pin,proto3" json:"pin,omitempty"`
	// RSSI of bluetooth connection
	Rssi int32 `protobuf:"varint,2,opt,name=rssi,proto3" json:"rssi,omitempty"`
	// Whether the device has an active connection or not
	IsConnected bool `protobuf:"varint,3,opt,name=is_connected,json=isConnected,proto3" json:"is_connected,omitempty"`
}

func (x *BluetoothConnectionStatus) Reset() {
	*x = BluetoothConnectionStatus{}
	if protoimpl.UnsafeEnabled {
		mi := &file_meshtastic_connection_status_proto_msgTypes[4]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *BluetoothConnectionStatus) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*BluetoothConnectionStatus) ProtoMessage() {}

func (x *BluetoothConnectionStatus) ProtoReflect() protoreflect.Message {
	mi := &file_meshtastic_connection_status_proto_msgTypes[4]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use BluetoothConnectionStatus.ProtoReflect.Descriptor instead.
func (*BluetoothConnectionStatus) Descriptor() ([]byte, []int) {
	return file_meshtastic_connection_status_proto_rawDescGZIP(), []int{4}
}

func (x *BluetoothConnectionStatus) GetPin() uint32 {
	if x != nil {
		return x.Pin
	}
	return 0
}

func (x *BluetoothConnectionStatus) GetRssi() int32 {
	if x != nil {
		return x.Rssi
	}
	return 0
}

func (x *BluetoothConnectionStatus) GetIsConnected() bool {
	if x != nil {
		return x.IsConnected
	}
	return false
}

// Serial connection status
type SerialConnectionStatus struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	// Serial baud rate
	Baud uint32 `protobuf:"varint,1,opt,name=baud,proto3" json:"baud,omitempty"`
	// Whether the device has an active connection or not
	IsConnected bool `protobuf:"varint,2,opt,name=is_connected,json=isConnected,proto3" json:"is_connected,omitempty"`
}

func (x *SerialConnectionStatus) Reset() {
	*x = SerialConnectionStatus{}
	if protoimpl.UnsafeEnabled {
		mi := &file_meshtastic_connection_status_proto_msgTypes[5]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *SerialConnectionStatus) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*SerialConnectionStatus) ProtoMessage() {}

func (x *SerialConnectionStatus) ProtoReflect() protoreflect.Message {
	mi := &file_meshtastic_connection_status_proto_msgTypes[5]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use SerialConnectionStatus.ProtoReflect.Descriptor instead.
func (*SerialConnectionStatus) Descriptor() ([]byte, []int) {
	return file_meshtastic_connection_status_proto_rawDescGZIP(), []int{5}
}

func (x *SerialConnectionStatus) GetBaud() uint32 {
	if x != nil {
		return x.Baud
	}
	return 0
}

func (x *SerialConnectionStatus) GetIsConnected() bool {
	if x != nil {
		return x.IsConnected
	}
	return false
}

var File_meshtastic_connection_status_proto protoreflect.FileDescriptor

var file_meshtastic_connection_status_proto_rawDesc = []byte{
	0x0a, 0x22, 0x6d, 0x65, 0x73, 0x68, 0x74, 0x61, 0x73, 0x74, 0x69, 0x63, 0x2f, 0x63, 0x6f, 0x6e,
	0x6e, 0x65, 0x63, 0x74, 0x69, 0x6f, 0x6e, 0x5f, 0x73, 0x74, 0x61, 0x74, 0x75, 0x73, 0x2e, 0x70,
	0x72, 0x6f, 0x74, 0x6f, 0x12, 0x0a, 0x6d, 0x65, 0x73, 0x68, 0x74, 0x61, 0x73, 0x74, 0x69, 0x63,
	0x22, 0xd4, 0x02, 0x0a, 0x16, 0x44, 0x65, 0x76, 0x69, 0x63, 0x65, 0x43, 0x6f, 0x6e, 0x6e, 0x65,
	0x63, 0x74, 0x69, 0x6f, 0x6e, 0x53, 0x74, 0x61, 0x74, 0x75, 0x73, 0x12, 0x39, 0x0a, 0x04, 0x77,
	0x69, 0x66, 0x69, 0x18, 0x01, 0x20, 0x01, 0x28, 0x0b, 0x32, 0x20, 0x2e, 0x6d, 0x65, 0x73, 0x68,
	0x74, 0x61, 0x73, 0x74, 0x69, 0x63, 0x2e, 0x57, 0x69, 0x66, 0x69, 0x43, 0x6f, 0x6e, 0x6e, 0x65,
	0x63, 0x74, 0x69, 0x6f, 0x6e, 0x53, 0x74, 0x61, 0x74, 0x75, 0x73, 0x48, 0x00, 0x52, 0x04, 0x77,
	0x69, 0x66, 0x69, 0x88, 0x01, 0x01, 0x12, 0x45, 0x0a, 0x08, 0x65, 0x74, 0x68, 0x65, 0x72, 0x6e,
	0x65, 0x74, 0x18, 0x02, 0x20, 0x01, 0x28, 0x0b, 0x32, 0x24, 0x2e, 0x6d, 0x65, 0x73, 0x68, 0x74,
	0x61, 0x73, 0x74, 0x69, 0x63, 0x2e, 0x45, 0x74, 0x68, 0x65, 0x72, 0x6e, 0x65, 0x74, 0x43, 0x6f,
	0x6e, 0x6e, 0x65, 0x63, 0x74, 0x69, 0x6f, 0x6e, 0x53, 0x74, 0x61, 0x74, 0x75, 0x73, 0x48, 0x01,
	0x52, 0x08, 0x65, 0x74, 0x68, 0x65, 0x72, 0x6e, 0x65, 0x74, 0x88, 0x01, 0x01, 0x12, 0x48, 0x0a,
	0x09, 0x62, 0x6c, 0x75, 0x65, 0x74, 0x6f, 0x6f, 0x74, 0x68, 0x18, 0x03, 0x20, 0x01, 0x28, 0x0b,
	0x32, 0x25, 0x2e, 0x6d, 0x65, 0x73, 0x68, 0x74, 0x61, 0x73, 0x74, 0x69, 0x63, 0x2e, 0x42, 0x6c,
	0x75, 0x65, 0x74, 0x6f, 0x6f, 0x74, 0x68, 0x43, 0x6f, 0x6e, 0x6e, 0x65, 0x63, 0x74, 0x69, 0x6f,
	0x6e, 0x53, 0x74, 0x61, 0x74, 0x75, 0x73, 0x48, 0x02, 0x52, 0x09, 0x62, 0x6c, 0x75, 0x65, 0x74,
	0x6f, 0x6f, 0x74, 0x68, 0x88, 0x01, 0x01, 0x12, 0x3f, 0x0a, 0x06, 0x73, 0x65, 0x72, 0x69, 0x61,
	0x6c, 0x18, 0x04, 0x20, 0x01, 0x28, 0x0b, 0x32, 0x22, 0x2e, 0x6d, 0x65, 0x73, 0x68, 0x74, 0x61,
	0x73, 0x74, 0x69, 0x63, 0x2e, 0x53, 0x65, 0x72, 0x69, 0x61, 0x6c, 0x43, 0x6f, 0x6e, 0x6e, 0x65,
	0x63, 0x74, 0x69, 0x6f, 0x6e, 0x53, 0x74, 0x61, 0x74, 0x75, 0x73, 0x48, 0x03, 0x52, 0x06, 0x73,
	0x65, 0x72, 0x69, 0x61, 0x6c, 0x88, 0x01, 0x01, 0x42, 0x07, 0x0a, 0x05, 0x5f, 0x77, 0x69, 0x66,
	0x69, 0x42, 0x0b, 0x0a, 0x09, 0x5f, 0x65, 0x74, 0x68, 0x65, 0x72, 0x6e, 0x65, 0x74, 0x42, 0x0c,
	0x0a, 0x0a, 0x5f, 0x62, 0x6c, 0x75, 0x65, 0x74, 0x6f, 0x6f, 0x74, 0x68, 0x42, 0x09, 0x0a, 0x07,
	0x5f, 0x73, 0x65, 0x72, 0x69, 0x61, 0x6c, 0x22, 0x7b, 0x0a, 0x14, 0x57, 0x69, 0x66, 0x69, 0x43,
	0x6f, 0x6e, 0x6e, 0x65, 0x63, 0x74, 0x69, 0x6f, 0x6e, 0x53, 0x74, 0x61, 0x74, 0x75, 0x73, 0x12,
	0x3b, 0x0a, 0x06, 0x73, 0x74, 0x61, 0x74, 0x75, 0x73, 0x18, 0x01, 0x20, 0x01, 0x28, 0x0b, 0x32,
	0x23, 0x2e, 0x6d, 0x65, 0x73, 0x68, 0x74, 0x61, 0x73, 0x74, 0x69, 0x63, 0x2e, 0x4e, 0x65, 0x74,
	0x77, 0x6f, 0x72, 0x6b, 0x43, 0x6f, 0x6e, 0x6e, 0x65, 0x63, 0x74, 0x69, 0x6f, 0x6e, 0x53, 0x74,
	0x61, 0x74, 0x75, 0x73, 0x52, 0x06, 0x73, 0x74, 0x61, 0x74, 0x75, 0x73, 0x12, 0x12, 0x0a, 0x04,
	0x73, 0x73, 0x69, 0x64, 0x18, 0x02, 0x20, 0x01, 0x28, 0x09, 0x52, 0x04, 0x73, 0x73, 0x69, 0x64,
	0x12, 0x12, 0x0a, 0x04, 0x72, 0x73, 0x73, 0x69, 0x18, 0x03, 0x20, 0x01, 0x28, 0x05, 0x52, 0x04,
	0x72, 0x73, 0x73, 0x69, 0x22, 0x57, 0x0a, 0x18, 0x45, 0x74, 0x68, 0x65, 0x72, 0x6e, 0x65, 0x74,
	0x43, 0x6f, 0x6e, 0x6e, 0x65, 0x63, 0x74, 0x69, 0x6f, 0x6e, 0x53, 0x74, 0x61, 0x74, 0x75, 0x73,
	0x12, 0x3b, 0x0a, 0x06, 0x73, 0x74, 0x61, 0x74, 0x75, 0x73, 0x18, 0x01, 0x20, 0x01, 0x28, 0x0b,
	0x32, 0x23, 0x2e, 0x6d, 0x65, 0x73, 0x68, 0x74, 0x61, 0x73, 0x74, 0x69, 0x63, 0x2e, 0x4e, 0x65,
	0x74, 0x77, 0x6f, 0x72, 0x6b, 0x43, 0x6f, 0x6e, 0x6e, 0x65, 0x63, 0x74, 0x69, 0x6f, 0x6e, 0x53,
	0x74, 0x61, 0x74, 0x75, 0x73, 0x52, 0x06, 0x73, 0x74, 0x61, 0x74, 0x75, 0x73, 0x22, 0xb7, 0x01,
	0x0a, 0x17, 0x4e, 0x65, 0x74, 0x77, 0x6f, 0x72, 0x6b, 0x43, 0x6f, 0x6e, 0x6e, 0x65, 0x63, 0x74,
	0x69, 0x6f, 0x6e, 0x53, 0x74, 0x61, 0x74, 0x75, 0x73, 0x12, 0x1d, 0x0a, 0x0a, 0x69, 0x70, 0x5f,
	0x61, 0x64, 0x64, 0x72, 0x65, 0x73, 0x73, 0x18, 0x01, 0x20, 0x01, 0x28, 0x07, 0x52, 0x09, 0x69,
	0x70, 0x41, 0x64, 0x64, 0x72, 0x65, 0x73, 0x73, 0x12, 0x21, 0x0a, 0x0c, 0x69, 0x73, 0x5f, 0x63,
	0x6f, 0x6e, 0x6e, 0x65, 0x63, 0x74, 0x65, 0x64, 0x18, 0x02, 0x20, 0x01, 0x28, 0x08, 0x52, 0x0b,
	0x69, 0x73, 0x43, 0x6f, 0x6e, 0x6e, 0x65, 0x63, 0x74, 0x65, 0x64, 0x12, 0x2a, 0x0a, 0x11, 0x69,
	0x73, 0x5f, 0x6d, 0x71, 0x74, 0x74, 0x5f, 0x63, 0x6f, 0x6e, 0x6e, 0x65, 0x63, 0x74, 0x65, 0x64,
	0x18, 0x03, 0x20, 0x01, 0x28, 0x08, 0x52, 0x0f, 0x69, 0x73, 0x4d, 0x71, 0x74, 0x74, 0x43, 0x6f,
	0x6e, 0x6e, 0x65, 0x63, 0x74, 0x65, 0x64, 0x12, 0x2e, 0x0a, 0x13, 0x69, 0x73, 0x5f, 0x73, 0x79,
	0x73, 0x6c, 0x6f, 0x67, 0x5f, 0x63, 0x6f, 0x6e, 0x6e, 0x65, 0x63, 0x74, 0x65, 0x64, 0x18, 0x04,
	0x20, 0x01, 0x28, 0x08, 0x52, 0x11, 0x69, 0x73, 0x53, 0x79, 0x73, 0x6c, 0x6f, 0x67, 0x43, 0x6f,
	0x6e, 0x6e, 0x65, 0x63, 0x74, 0x65, 0x64, 0x22, 0x64, 0x0a, 0x19, 0x42, 0x6c, 0x75, 0x65, 0x74,
	0x6f, 0x6f, 0x74, 0x68, 0x43, 0x6f, 0x6e, 0x6e, 0x65, 0x63, 0x74, 0x69, 0x6f, 0x6e, 0x53, 0x74,
	0x61, 0x74, 0x75, 0x73, 0x12, 0x10, 0x0a, 0x03, 0x70, 0x69, 0x6e, 0x18, 0x01, 0x20, 0x01, 0x28,
	0x0d, 0x52, 0x03, 0x70, 0x69, 0x6e, 0x12, 0x12, 0x0a, 0x04, 0x72, 0x73, 0x73, 0x69, 0x18, 0x02,
	0x20, 0x01, 0x28, 0x05, 0x52, 0x04, 0x72, 0x73, 0x73, 0x69, 0x12, 0x21, 0x0a, 0x0c, 0x69, 0x73,
	0x5f, 0x63, 0x6f, 0x6e, 0x6e, 0x65, 0x63, 0x74, 0x65, 0x64, 0x18, 0x03, 0x20, 0x01, 0x28, 0x08,
	0x52, 0x0b, 0x69, 0x73, 0x43, 0x6f, 0x6e, 0x6e, 0x65, 0x63, 0x74, 0x65, 0x64, 0x22, 0x4f, 0x0a,
	0x16, 0x53, 0x65, 0x72, 0x69, 0x61, 0x6c, 0x43, 0x6f, 0x6e, 0x6e, 0x65, 0x63, 0x74, 0x69, 0x6f,
	0x6e, 0x53, 0x74, 0x61, 0x74, 0x75, 0x73, 0x12, 0x12, 0x0a, 0x04, 0x62, 0x61, 0x75, 0x64, 0x18,
	0x01, 0x20, 0x01, 0x28, 0x0d, 0x52, 0x04, 0x62, 0x61, 0x75, 0x64, 0x12, 0x21, 0x0a, 0x0c, 0x69,
	0x73, 0x5f, 0x63, 0x6f, 0x6e, 0x6e, 0x65, 0x63, 0x74, 0x65, 0x64, 0x18, 0x02, 0x20, 0x01, 0x28,
	0x08, 0x52, 0x0b, 0x69, 0x73, 0x43, 0x6f, 0x6e, 0x6e, 0x65, 0x63, 0x74, 0x65, 0x64, 0x42, 0x65,
	0x0a, 0x13, 0x63, 0x6f, 0x6d, 0x2e, 0x67, 0x65, 0x65, 0x6b, 0x73, 0x76, 0x69, 0x6c, 0x6c, 0x65,
	0x2e, 0x6d, 0x65, 0x73, 0x68, 0x42, 0x10, 0x43, 0x6f, 0x6e, 0x6e, 0x53, 0x74, 0x61, 0x74, 0x75,
	0x73, 0x50, 0x72, 0x6f, 0x74, 0x6f, 0x73, 0x5a, 0x22, 0x67, 0x69, 0x74, 0x68, 0x75, 0x62, 0x2e,
	0x63, 0x6f, 0x6d, 0x2f, 0x6d, 0x65, 0x73, 0x68, 0x74, 0x61, 0x73, 0x74, 0x69, 0x63, 0x2f, 0x67,
	0x6f, 0x2f, 0x67, 0x65, 0x6e, 0x65, 0x72, 0x61, 0x74, 0x65, 0x64, 0xaa, 0x02, 0x14, 0x4d, 0x65,
	0x73, 0x68, 0x74, 0x61, 0x73, 0x74, 0x69, 0x63, 0x2e, 0x50, 0x72, 0x6f, 0x74, 0x6f, 0x62, 0x75,
	0x66, 0x73, 0xba, 0x02, 0x00, 0x62, 0x06, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x33,
}

var (
	file_meshtastic_connection_status_proto_rawDescOnce sync.Once
	file_meshtastic_connection_status_proto_rawDescData = file_meshtastic_connection_status_proto_rawDesc
)

func file_meshtastic_connection_status_proto_rawDescGZIP() []byte {
	file_meshtastic_connection_status_proto_rawDescOnce.Do(func() {
		file_meshtastic_connection_status_proto_rawDescData = protoimpl.X.CompressGZIP(file_meshtastic_connection_status_proto_rawDescData)
	})
	return file_meshtastic_connection_status_proto_rawDescData
}

var file_meshtastic_connection_status_proto_msgTypes = make([]protoimpl.MessageInfo, 6)
var file_meshtastic_connection_status_proto_goTypes = []interface{}{
	(*DeviceConnectionStatus)(nil),    // 0: meshtastic.DeviceConnectionStatus
	(*WifiConnectionStatus)(nil),      // 1: meshtastic.WifiConnectionStatus
	(*EthernetConnectionStatus)(nil),  // 2: meshtastic.EthernetConnectionStatus
	(*NetworkConnectionStatus)(nil),   // 3: meshtastic.NetworkConnectionStatus
	(*BluetoothConnectionStatus)(nil), // 4: meshtastic.BluetoothConnectionStatus
	(*SerialConnectionStatus)(nil),    // 5: meshtastic.SerialConnectionStatus
}
var file_meshtastic_connection_status_proto_depIdxs = []int32{
	1, // 0: meshtastic.DeviceConnectionStatus.wifi:type_name -> meshtastic.WifiConnectionStatus
	2, // 1: meshtastic.DeviceConnectionStatus.ethernet:type_name -> meshtastic.EthernetConnectionStatus
	4, // 2: meshtastic.DeviceConnectionStatus.bluetooth:type_name -> meshtastic.BluetoothConnectionStatus
	5, // 3: meshtastic.DeviceConnectionStatus.serial:type_name -> meshtastic.SerialConnectionStatus
	3, // 4: meshtastic.WifiConnectionStatus.status:type_name -> meshtastic.NetworkConnectionStatus
	3, // 5: meshtastic.EthernetConnectionStatus.status:type_name -> meshtastic.NetworkConnectionStatus
	6, // [6:6] is the sub-list for method output_type
	6, // [6:6] is the sub-list for method input_type
	6, // [6:6] is the sub-list for extension type_name
	6, // [6:6] is the sub-list for extension extendee
	0, // [0:6] is the sub-list for field type_name
}

func init() { file_meshtastic_connection_status_proto_init() }
func file_meshtastic_connection_status_proto_init() {
	if File_meshtastic_connection_status_proto != nil {
		return
	}
	if !protoimpl.UnsafeEnabled {
		file_meshtastic_connection_status_proto_msgTypes[0].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*DeviceConnectionStatus); i {
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
		file_meshtastic_connection_status_proto_msgTypes[1].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*WifiConnectionStatus); i {
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
		file_meshtastic_connection_status_proto_msgTypes[2].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*EthernetConnectionStatus); i {
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
		file_meshtastic_connection_status_proto_msgTypes[3].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*NetworkConnectionStatus); i {
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
		file_meshtastic_connection_status_proto_msgTypes[4].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*BluetoothConnectionStatus); i {
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
		file_meshtastic_connection_status_proto_msgTypes[5].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*SerialConnectionStatus); i {
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
	file_meshtastic_connection_status_proto_msgTypes[0].OneofWrappers = []interface{}{}
	type x struct{}
	out := protoimpl.TypeBuilder{
		File: protoimpl.DescBuilder{
			GoPackagePath: reflect.TypeOf(x{}).PkgPath(),
			RawDescriptor: file_meshtastic_connection_status_proto_rawDesc,
			NumEnums:      0,
			NumMessages:   6,
			NumExtensions: 0,
			NumServices:   0,
		},
		GoTypes:           file_meshtastic_connection_status_proto_goTypes,
		DependencyIndexes: file_meshtastic_connection_status_proto_depIdxs,
		MessageInfos:      file_meshtastic_connection_status_proto_msgTypes,
	}.Build()
	File_meshtastic_connection_status_proto = out.File
	file_meshtastic_connection_status_proto_rawDesc = nil
	file_meshtastic_connection_status_proto_goTypes = nil
	file_meshtastic_connection_status_proto_depIdxs = nil
}
