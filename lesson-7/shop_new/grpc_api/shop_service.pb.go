// Code generated by protoc-gen-go. DO NOT EDIT.
// source: shop_service.proto

package grpc_api

import (
	fmt "fmt"
	proto "github.com/golang/protobuf/proto"
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

// represents methods which not returns any data
type Empty struct {
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *Empty) Reset()         { *m = Empty{} }
func (m *Empty) String() string { return proto.CompactTextString(m) }
func (*Empty) ProtoMessage()    {}
func (*Empty) Descriptor() ([]byte, []int) {
	return fileDescriptor_acb80bbe5ec7a30f, []int{0}
}

func (m *Empty) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_Empty.Unmarshal(m, b)
}
func (m *Empty) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_Empty.Marshal(b, m, deterministic)
}
func (m *Empty) XXX_Merge(src proto.Message) {
	xxx_messageInfo_Empty.Merge(m, src)
}
func (m *Empty) XXX_Size() int {
	return xxx_messageInfo_Empty.Size(m)
}
func (m *Empty) XXX_DiscardUnknown() {
	xxx_messageInfo_Empty.DiscardUnknown(m)
}

var xxx_messageInfo_Empty proto.InternalMessageInfo

type Timestamp struct {
	// UTC seconds from unix epoch
	Seconds              int64    `protobuf:"varint,1,opt,name=seconds,proto3" json:"seconds,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *Timestamp) Reset()         { *m = Timestamp{} }
func (m *Timestamp) String() string { return proto.CompactTextString(m) }
func (*Timestamp) ProtoMessage()    {}
func (*Timestamp) Descriptor() ([]byte, []int) {
	return fileDescriptor_acb80bbe5ec7a30f, []int{1}
}

func (m *Timestamp) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_Timestamp.Unmarshal(m, b)
}
func (m *Timestamp) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_Timestamp.Marshal(b, m, deterministic)
}
func (m *Timestamp) XXX_Merge(src proto.Message) {
	xxx_messageInfo_Timestamp.Merge(m, src)
}
func (m *Timestamp) XXX_Size() int {
	return xxx_messageInfo_Timestamp.Size(m)
}
func (m *Timestamp) XXX_DiscardUnknown() {
	xxx_messageInfo_Timestamp.DiscardUnknown(m)
}

var xxx_messageInfo_Timestamp proto.InternalMessageInfo

func (m *Timestamp) GetSeconds() int64 {
	if m != nil {
		return m.Seconds
	}
	return 0
}

type Item struct {
	Id                   int32      `protobuf:"varint,1,opt,name=id,proto3" json:"id,omitempty"`
	Name                 string     `protobuf:"bytes,2,opt,name=name,proto3" json:"name,omitempty"`
	Price                int64      `protobuf:"varint,3,opt,name=price,proto3" json:"price,omitempty"`
	Createdat            *Timestamp `protobuf:"bytes,4,opt,name=createdat,proto3" json:"createdat,omitempty"`
	Updatedat            *Timestamp `protobuf:"bytes,5,opt,name=updatedat,proto3" json:"updatedat,omitempty"`
	XXX_NoUnkeyedLiteral struct{}   `json:"-"`
	XXX_unrecognized     []byte     `json:"-"`
	XXX_sizecache        int32      `json:"-"`
}

func (m *Item) Reset()         { *m = Item{} }
func (m *Item) String() string { return proto.CompactTextString(m) }
func (*Item) ProtoMessage()    {}
func (*Item) Descriptor() ([]byte, []int) {
	return fileDescriptor_acb80bbe5ec7a30f, []int{2}
}

func (m *Item) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_Item.Unmarshal(m, b)
}
func (m *Item) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_Item.Marshal(b, m, deterministic)
}
func (m *Item) XXX_Merge(src proto.Message) {
	xxx_messageInfo_Item.Merge(m, src)
}
func (m *Item) XXX_Size() int {
	return xxx_messageInfo_Item.Size(m)
}
func (m *Item) XXX_DiscardUnknown() {
	xxx_messageInfo_Item.DiscardUnknown(m)
}

var xxx_messageInfo_Item proto.InternalMessageInfo

func (m *Item) GetId() int32 {
	if m != nil {
		return m.Id
	}
	return 0
}

func (m *Item) GetName() string {
	if m != nil {
		return m.Name
	}
	return ""
}

func (m *Item) GetPrice() int64 {
	if m != nil {
		return m.Price
	}
	return 0
}

func (m *Item) GetCreatedat() *Timestamp {
	if m != nil {
		return m.Createdat
	}
	return nil
}

func (m *Item) GetUpdatedat() *Timestamp {
	if m != nil {
		return m.Updatedat
	}
	return nil
}

type ItemFilter struct {
	Limit                int32    `protobuf:"varint,1,opt,name=limit,proto3" json:"limit,omitempty"`
	Offset               int32    `protobuf:"varint,2,opt,name=offset,proto3" json:"offset,omitempty"`
	Priceleft            int64    `protobuf:"varint,3,opt,name=priceleft,proto3" json:"priceleft,omitempty"`
	Priceright           int64    `protobuf:"varint,4,opt,name=priceright,proto3" json:"priceright,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *ItemFilter) Reset()         { *m = ItemFilter{} }
func (m *ItemFilter) String() string { return proto.CompactTextString(m) }
func (*ItemFilter) ProtoMessage()    {}
func (*ItemFilter) Descriptor() ([]byte, []int) {
	return fileDescriptor_acb80bbe5ec7a30f, []int{3}
}

func (m *ItemFilter) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_ItemFilter.Unmarshal(m, b)
}
func (m *ItemFilter) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_ItemFilter.Marshal(b, m, deterministic)
}
func (m *ItemFilter) XXX_Merge(src proto.Message) {
	xxx_messageInfo_ItemFilter.Merge(m, src)
}
func (m *ItemFilter) XXX_Size() int {
	return xxx_messageInfo_ItemFilter.Size(m)
}
func (m *ItemFilter) XXX_DiscardUnknown() {
	xxx_messageInfo_ItemFilter.DiscardUnknown(m)
}

var xxx_messageInfo_ItemFilter proto.InternalMessageInfo

func (m *ItemFilter) GetLimit() int32 {
	if m != nil {
		return m.Limit
	}
	return 0
}

func (m *ItemFilter) GetOffset() int32 {
	if m != nil {
		return m.Offset
	}
	return 0
}

func (m *ItemFilter) GetPriceleft() int64 {
	if m != nil {
		return m.Priceleft
	}
	return 0
}

func (m *ItemFilter) GetPriceright() int64 {
	if m != nil {
		return m.Priceright
	}
	return 0
}

type ItemList struct {
	Items                []*Item  `protobuf:"bytes,1,rep,name=items,proto3" json:"items,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *ItemList) Reset()         { *m = ItemList{} }
func (m *ItemList) String() string { return proto.CompactTextString(m) }
func (*ItemList) ProtoMessage()    {}
func (*ItemList) Descriptor() ([]byte, []int) {
	return fileDescriptor_acb80bbe5ec7a30f, []int{4}
}

func (m *ItemList) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_ItemList.Unmarshal(m, b)
}
func (m *ItemList) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_ItemList.Marshal(b, m, deterministic)
}
func (m *ItemList) XXX_Merge(src proto.Message) {
	xxx_messageInfo_ItemList.Merge(m, src)
}
func (m *ItemList) XXX_Size() int {
	return xxx_messageInfo_ItemList.Size(m)
}
func (m *ItemList) XXX_DiscardUnknown() {
	xxx_messageInfo_ItemList.DiscardUnknown(m)
}

var xxx_messageInfo_ItemList proto.InternalMessageInfo

func (m *ItemList) GetItems() []*Item {
	if m != nil {
		return m.Items
	}
	return nil
}

type ItemID struct {
	Id                   int32    `protobuf:"varint,1,opt,name=id,proto3" json:"id,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *ItemID) Reset()         { *m = ItemID{} }
func (m *ItemID) String() string { return proto.CompactTextString(m) }
func (*ItemID) ProtoMessage()    {}
func (*ItemID) Descriptor() ([]byte, []int) {
	return fileDescriptor_acb80bbe5ec7a30f, []int{5}
}

func (m *ItemID) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_ItemID.Unmarshal(m, b)
}
func (m *ItemID) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_ItemID.Marshal(b, m, deterministic)
}
func (m *ItemID) XXX_Merge(src proto.Message) {
	xxx_messageInfo_ItemID.Merge(m, src)
}
func (m *ItemID) XXX_Size() int {
	return xxx_messageInfo_ItemID.Size(m)
}
func (m *ItemID) XXX_DiscardUnknown() {
	xxx_messageInfo_ItemID.DiscardUnknown(m)
}

var xxx_messageInfo_ItemID proto.InternalMessageInfo

func (m *ItemID) GetId() int32 {
	if m != nil {
		return m.Id
	}
	return 0
}

type Order struct {
	Id                   int32      `protobuf:"varint,1,opt,name=id,proto3" json:"id,omitempty"`
	Customername         string     `protobuf:"bytes,2,opt,name=customername,proto3" json:"customername,omitempty"`
	Customerphone        string     `protobuf:"bytes,3,opt,name=customerphone,proto3" json:"customerphone,omitempty"`
	Customeremail        string     `protobuf:"bytes,4,opt,name=customeremail,proto3" json:"customeremail,omitempty"`
	Itemids              []*ItemID  `protobuf:"bytes,5,rep,name=itemids,proto3" json:"itemids,omitempty"`
	Createdat            *Timestamp `protobuf:"bytes,6,opt,name=createdat,proto3" json:"createdat,omitempty"`
	Updatedat            *Timestamp `protobuf:"bytes,7,opt,name=updatedat,proto3" json:"updatedat,omitempty"`
	XXX_NoUnkeyedLiteral struct{}   `json:"-"`
	XXX_unrecognized     []byte     `json:"-"`
	XXX_sizecache        int32      `json:"-"`
}

func (m *Order) Reset()         { *m = Order{} }
func (m *Order) String() string { return proto.CompactTextString(m) }
func (*Order) ProtoMessage()    {}
func (*Order) Descriptor() ([]byte, []int) {
	return fileDescriptor_acb80bbe5ec7a30f, []int{6}
}

func (m *Order) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_Order.Unmarshal(m, b)
}
func (m *Order) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_Order.Marshal(b, m, deterministic)
}
func (m *Order) XXX_Merge(src proto.Message) {
	xxx_messageInfo_Order.Merge(m, src)
}
func (m *Order) XXX_Size() int {
	return xxx_messageInfo_Order.Size(m)
}
func (m *Order) XXX_DiscardUnknown() {
	xxx_messageInfo_Order.DiscardUnknown(m)
}

var xxx_messageInfo_Order proto.InternalMessageInfo

func (m *Order) GetId() int32 {
	if m != nil {
		return m.Id
	}
	return 0
}

func (m *Order) GetCustomername() string {
	if m != nil {
		return m.Customername
	}
	return ""
}

func (m *Order) GetCustomerphone() string {
	if m != nil {
		return m.Customerphone
	}
	return ""
}

func (m *Order) GetCustomeremail() string {
	if m != nil {
		return m.Customeremail
	}
	return ""
}

func (m *Order) GetItemids() []*ItemID {
	if m != nil {
		return m.Itemids
	}
	return nil
}

func (m *Order) GetCreatedat() *Timestamp {
	if m != nil {
		return m.Createdat
	}
	return nil
}

func (m *Order) GetUpdatedat() *Timestamp {
	if m != nil {
		return m.Updatedat
	}
	return nil
}

type OrderFilter struct {
	Limit                int32    `protobuf:"varint,1,opt,name=limit,proto3" json:"limit,omitempty"`
	Offset               int32    `protobuf:"varint,2,opt,name=offset,proto3" json:"offset,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *OrderFilter) Reset()         { *m = OrderFilter{} }
func (m *OrderFilter) String() string { return proto.CompactTextString(m) }
func (*OrderFilter) ProtoMessage()    {}
func (*OrderFilter) Descriptor() ([]byte, []int) {
	return fileDescriptor_acb80bbe5ec7a30f, []int{7}
}

func (m *OrderFilter) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_OrderFilter.Unmarshal(m, b)
}
func (m *OrderFilter) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_OrderFilter.Marshal(b, m, deterministic)
}
func (m *OrderFilter) XXX_Merge(src proto.Message) {
	xxx_messageInfo_OrderFilter.Merge(m, src)
}
func (m *OrderFilter) XXX_Size() int {
	return xxx_messageInfo_OrderFilter.Size(m)
}
func (m *OrderFilter) XXX_DiscardUnknown() {
	xxx_messageInfo_OrderFilter.DiscardUnknown(m)
}

var xxx_messageInfo_OrderFilter proto.InternalMessageInfo

func (m *OrderFilter) GetLimit() int32 {
	if m != nil {
		return m.Limit
	}
	return 0
}

func (m *OrderFilter) GetOffset() int32 {
	if m != nil {
		return m.Offset
	}
	return 0
}

type OrderList struct {
	Orders               []*Order `protobuf:"bytes,1,rep,name=orders,proto3" json:"orders,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *OrderList) Reset()         { *m = OrderList{} }
func (m *OrderList) String() string { return proto.CompactTextString(m) }
func (*OrderList) ProtoMessage()    {}
func (*OrderList) Descriptor() ([]byte, []int) {
	return fileDescriptor_acb80bbe5ec7a30f, []int{8}
}

func (m *OrderList) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_OrderList.Unmarshal(m, b)
}
func (m *OrderList) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_OrderList.Marshal(b, m, deterministic)
}
func (m *OrderList) XXX_Merge(src proto.Message) {
	xxx_messageInfo_OrderList.Merge(m, src)
}
func (m *OrderList) XXX_Size() int {
	return xxx_messageInfo_OrderList.Size(m)
}
func (m *OrderList) XXX_DiscardUnknown() {
	xxx_messageInfo_OrderList.DiscardUnknown(m)
}

var xxx_messageInfo_OrderList proto.InternalMessageInfo

func (m *OrderList) GetOrders() []*Order {
	if m != nil {
		return m.Orders
	}
	return nil
}

func init() {
	proto.RegisterType((*Empty)(nil), "grpc_api.Empty")
	proto.RegisterType((*Timestamp)(nil), "grpc_api.Timestamp")
	proto.RegisterType((*Item)(nil), "grpc_api.Item")
	proto.RegisterType((*ItemFilter)(nil), "grpc_api.ItemFilter")
	proto.RegisterType((*ItemList)(nil), "grpc_api.ItemList")
	proto.RegisterType((*ItemID)(nil), "grpc_api.ItemID")
	proto.RegisterType((*Order)(nil), "grpc_api.Order")
	proto.RegisterType((*OrderFilter)(nil), "grpc_api.OrderFilter")
	proto.RegisterType((*OrderList)(nil), "grpc_api.OrderList")
}

func init() { proto.RegisterFile("shop_service.proto", fileDescriptor_acb80bbe5ec7a30f) }

var fileDescriptor_acb80bbe5ec7a30f = []byte{
	// 527 bytes of a gzipped FileDescriptorProto
	0x1f, 0x8b, 0x08, 0x00, 0x00, 0x00, 0x00, 0x00, 0x02, 0xff, 0x9c, 0x54, 0xc1, 0x6e, 0xd3, 0x40,
	0x14, 0x8c, 0x93, 0x3a, 0xa9, 0x5f, 0x4a, 0x8a, 0x5e, 0x0b, 0xb2, 0x22, 0x84, 0xa2, 0x55, 0x11,
	0x11, 0x87, 0x40, 0x02, 0x5c, 0xe8, 0x0d, 0x02, 0x28, 0x12, 0x12, 0x92, 0x05, 0x17, 0x2e, 0x95,
	0xb1, 0x5f, 0x9a, 0x95, 0xe2, 0xec, 0x6a, 0x77, 0x8b, 0xe8, 0x8d, 0x3b, 0x5f, 0xc0, 0x0f, 0xf0,
	0x9d, 0x68, 0x77, 0xed, 0x26, 0x36, 0x05, 0x15, 0x6e, 0x99, 0x79, 0xb3, 0xde, 0x79, 0x33, 0xb1,
	0x01, 0xf5, 0x4a, 0xc8, 0x33, 0x4d, 0xea, 0x0b, 0xcf, 0x68, 0x22, 0x95, 0x30, 0x02, 0xf7, 0xcf,
	0x95, 0xcc, 0xce, 0x52, 0xc9, 0x59, 0x0f, 0xc2, 0xd7, 0x85, 0x34, 0x97, 0xec, 0x01, 0x44, 0x1f,
	0x78, 0x41, 0xda, 0xa4, 0x85, 0xc4, 0x18, 0x7a, 0x9a, 0x32, 0xb1, 0xc9, 0x75, 0x1c, 0x8c, 0x82,
	0x71, 0x27, 0xa9, 0x20, 0xfb, 0x19, 0xc0, 0xde, 0xc2, 0x50, 0x81, 0x03, 0x68, 0xf3, 0xdc, 0x4d,
	0xc3, 0xa4, 0xcd, 0x73, 0x44, 0xd8, 0xdb, 0xa4, 0x05, 0xc5, 0xed, 0x51, 0x30, 0x8e, 0x12, 0xf7,
	0x1b, 0x8f, 0x21, 0x94, 0x8a, 0x67, 0x14, 0x77, 0xdc, 0x43, 0x3c, 0xc0, 0x29, 0x44, 0x99, 0xa2,
	0xd4, 0x50, 0x9e, 0x9a, 0x78, 0x6f, 0x14, 0x8c, 0xfb, 0xb3, 0xa3, 0x49, 0x65, 0x68, 0x72, 0x65,
	0x22, 0xd9, 0xaa, 0xec, 0x91, 0x0b, 0x99, 0x97, 0x47, 0xc2, 0xbf, 0x1c, 0xb9, 0x52, 0xb1, 0xaf,
	0x00, 0xd6, 0xe7, 0x1b, 0xbe, 0x36, 0xa4, 0xac, 0x93, 0x35, 0x2f, 0xb8, 0x29, 0x0d, 0x7b, 0x80,
	0x77, 0xa1, 0x2b, 0x96, 0x4b, 0x4d, 0xc6, 0xb9, 0x0e, 0x93, 0x12, 0xe1, 0x3d, 0x88, 0x9c, 0xd5,
	0x35, 0x2d, 0x4d, 0xe9, 0x7d, 0x4b, 0xe0, 0x7d, 0x00, 0x07, 0x14, 0x3f, 0x5f, 0xf9, 0x05, 0x3a,
	0xc9, 0x0e, 0xc3, 0x9e, 0xc0, 0xbe, 0xbd, 0xf9, 0x1d, 0xd7, 0x06, 0x4f, 0x20, 0xe4, 0x86, 0x0a,
	0x1b, 0x63, 0x67, 0xdc, 0x9f, 0x0d, 0xb6, 0xa6, 0xad, 0x24, 0xf1, 0x43, 0x16, 0x43, 0xd7, 0xc2,
	0xc5, 0xbc, 0x99, 0x2a, 0xfb, 0xd1, 0x86, 0xf0, 0xbd, 0xca, 0x49, 0xfd, 0x96, 0x37, 0x83, 0x83,
	0xec, 0x42, 0x1b, 0x51, 0x90, 0xda, 0xc9, 0xbd, 0xc6, 0xe1, 0x09, 0xdc, 0xaa, 0xb0, 0x5c, 0x89,
	0x8d, 0xef, 0x21, 0x4a, 0xea, 0xe4, 0xae, 0x8a, 0x8a, 0x94, 0xaf, 0xdd, 0x4a, 0x3b, 0x2a, 0x47,
	0xe2, 0x23, 0xe8, 0x59, 0xb3, 0x3c, 0xd7, 0x71, 0xe8, 0x76, 0xb9, 0x5d, 0xdf, 0x65, 0x31, 0x4f,
	0x2a, 0x41, 0xbd, 0xe1, 0xee, 0xbf, 0x37, 0xdc, 0xbb, 0x51, 0xc3, 0xa7, 0xd0, 0x77, 0xd1, 0xfc,
	0x4f, 0xc5, 0xec, 0x19, 0x44, 0xee, 0xb0, 0x6b, 0xe9, 0x21, 0x74, 0x85, 0x05, 0x55, 0x4d, 0x87,
	0xdb, 0x9b, 0x9d, 0x28, 0x29, 0xc7, 0xb3, 0xef, 0x6d, 0x18, 0xb8, 0xe2, 0x48, 0x0a, 0xcd, 0x8d,
	0x50, 0x97, 0x38, 0x01, 0x78, 0xe5, 0xb6, 0xf0, 0x6f, 0x45, 0x3d, 0x94, 0x61, 0x03, 0xb3, 0x96,
	0xd5, 0x7f, 0x74, 0x2b, 0xdc, 0x50, 0xff, 0x18, 0x60, 0x4e, 0x6b, 0xfa, 0x83, 0x7e, 0xc7, 0xa9,
	0x7f, 0x8d, 0x5b, 0x38, 0x85, 0xde, 0x5b, 0x32, 0x4e, 0x7d, 0x5c, 0x57, 0xfb, 0xa0, 0xae, 0xb9,
	0xe3, 0x39, 0x44, 0x36, 0x07, 0x8b, 0x34, 0x62, 0x7d, 0x6c, 0x07, 0xc3, 0x6b, 0x38, 0xd6, 0x9a,
	0x7d, 0x0b, 0xe0, 0xd0, 0xe7, 0xb3, 0x8d, 0x63, 0x0a, 0x7d, 0x1f, 0x87, 0xff, 0xd7, 0x36, 0x93,
	0x1c, 0x36, 0x09, 0xd6, 0xc2, 0x17, 0x00, 0xf6, 0x81, 0x0e, 0x6a, 0xbc, 0xd3, 0x10, 0x94, 0xa6,
	0x8f, 0x1a, 0xb4, 0xb7, 0xf0, 0xf2, 0xe0, 0x13, 0x4c, 0x4e, 0xab, 0xc9, 0xe7, 0xae, 0xfb, 0xba,
	0x3d, 0xfd, 0x15, 0x00, 0x00, 0xff, 0xff, 0x17, 0x45, 0xf0, 0x2d, 0xf3, 0x04, 0x00, 0x00,
}
