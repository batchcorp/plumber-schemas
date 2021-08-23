// Code generated by protoc-gen-go. DO NOT EDIT.
// source: records/azure-service-bus.proto

package records

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

type AzureServiceBus struct {
	ContentType          string                 `protobuf:"bytes,1,opt,name=content_type,json=contentType,proto3" json:"content_type,omitempty"`
	CorrelationId        string                 `protobuf:"bytes,2,opt,name=correlation_id,json=correlationId,proto3" json:"correlation_id,omitempty"`
	Value                []byte                 `protobuf:"bytes,3,opt,name=value,proto3" json:"value,omitempty"`
	DeliveryCount        uint32                 `protobuf:"varint,4,opt,name=delivery_count,json=deliveryCount,proto3" json:"delivery_count,omitempty"`
	SessionId            string                 `protobuf:"bytes,5,opt,name=session_id,json=sessionId,proto3" json:"session_id,omitempty"`
	GroupSequence        uint32                 `protobuf:"varint,6,opt,name=group_sequence,json=groupSequence,proto3" json:"group_sequence,omitempty"`
	Id                   string                 `protobuf:"bytes,7,opt,name=id,proto3" json:"id,omitempty"`
	Label                string                 `protobuf:"bytes,8,opt,name=label,proto3" json:"label,omitempty"`
	ReplyTo              string                 `protobuf:"bytes,9,opt,name=reply_to,json=replyTo,proto3" json:"reply_to,omitempty"`
	ReplyToGroupId       string                 `protobuf:"bytes,10,opt,name=reply_to_group_id,json=replyToGroupId,proto3" json:"reply_to_group_id,omitempty"`
	To                   string                 `protobuf:"bytes,11,opt,name=to,proto3" json:"to,omitempty"`
	Ttl                  int64                  `protobuf:"varint,12,opt,name=ttl,proto3" json:"ttl,omitempty"`
	LockToken            string                 `protobuf:"bytes,13,opt,name=lock_token,json=lockToken,proto3" json:"lock_token,omitempty"`
	SystemProperties     *AzureSystemProperties `protobuf:"bytes,14,opt,name=system_properties,json=systemProperties,proto3" json:"system_properties,omitempty"`
	UserProperties       map[string]string      `protobuf:"bytes,15,rep,name=user_properties,json=userProperties,proto3" json:"user_properties,omitempty" protobuf_key:"bytes,1,opt,name=key,proto3" protobuf_val:"bytes,2,opt,name=value,proto3"`
	Format               uint32                 `protobuf:"varint,16,opt,name=format,proto3" json:"format,omitempty"`
	XXX_NoUnkeyedLiteral struct{}               `json:"-"`
	XXX_unrecognized     []byte                 `json:"-"`
	XXX_sizecache        int32                  `json:"-"`
}

func (m *AzureServiceBus) Reset()         { *m = AzureServiceBus{} }
func (m *AzureServiceBus) String() string { return proto.CompactTextString(m) }
func (*AzureServiceBus) ProtoMessage()    {}
func (*AzureServiceBus) Descriptor() ([]byte, []int) {
	return fileDescriptor_fd29aeb29eb6a1a2, []int{0}
}

func (m *AzureServiceBus) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_AzureServiceBus.Unmarshal(m, b)
}
func (m *AzureServiceBus) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_AzureServiceBus.Marshal(b, m, deterministic)
}
func (m *AzureServiceBus) XXX_Merge(src proto.Message) {
	xxx_messageInfo_AzureServiceBus.Merge(m, src)
}
func (m *AzureServiceBus) XXX_Size() int {
	return xxx_messageInfo_AzureServiceBus.Size(m)
}
func (m *AzureServiceBus) XXX_DiscardUnknown() {
	xxx_messageInfo_AzureServiceBus.DiscardUnknown(m)
}

var xxx_messageInfo_AzureServiceBus proto.InternalMessageInfo

func (m *AzureServiceBus) GetContentType() string {
	if m != nil {
		return m.ContentType
	}
	return ""
}

func (m *AzureServiceBus) GetCorrelationId() string {
	if m != nil {
		return m.CorrelationId
	}
	return ""
}

func (m *AzureServiceBus) GetValue() []byte {
	if m != nil {
		return m.Value
	}
	return nil
}

func (m *AzureServiceBus) GetDeliveryCount() uint32 {
	if m != nil {
		return m.DeliveryCount
	}
	return 0
}

func (m *AzureServiceBus) GetSessionId() string {
	if m != nil {
		return m.SessionId
	}
	return ""
}

func (m *AzureServiceBus) GetGroupSequence() uint32 {
	if m != nil {
		return m.GroupSequence
	}
	return 0
}

func (m *AzureServiceBus) GetId() string {
	if m != nil {
		return m.Id
	}
	return ""
}

func (m *AzureServiceBus) GetLabel() string {
	if m != nil {
		return m.Label
	}
	return ""
}

func (m *AzureServiceBus) GetReplyTo() string {
	if m != nil {
		return m.ReplyTo
	}
	return ""
}

func (m *AzureServiceBus) GetReplyToGroupId() string {
	if m != nil {
		return m.ReplyToGroupId
	}
	return ""
}

func (m *AzureServiceBus) GetTo() string {
	if m != nil {
		return m.To
	}
	return ""
}

func (m *AzureServiceBus) GetTtl() int64 {
	if m != nil {
		return m.Ttl
	}
	return 0
}

func (m *AzureServiceBus) GetLockToken() string {
	if m != nil {
		return m.LockToken
	}
	return ""
}

func (m *AzureServiceBus) GetSystemProperties() *AzureSystemProperties {
	if m != nil {
		return m.SystemProperties
	}
	return nil
}

func (m *AzureServiceBus) GetUserProperties() map[string]string {
	if m != nil {
		return m.UserProperties
	}
	return nil
}

func (m *AzureServiceBus) GetFormat() uint32 {
	if m != nil {
		return m.Format
	}
	return 0
}

type AzureSystemProperties struct {
	LockedUntil            int64             `protobuf:"varint,1,opt,name=locked_until,json=lockedUntil,proto3" json:"locked_until,omitempty"`
	SequenceNumber         int64             `protobuf:"varint,2,opt,name=sequence_number,json=sequenceNumber,proto3" json:"sequence_number,omitempty"`
	PartitionId            int32             `protobuf:"varint,3,opt,name=partition_id,json=partitionId,proto3" json:"partition_id,omitempty"`
	PartitionKey           string            `protobuf:"bytes,4,opt,name=partition_key,json=partitionKey,proto3" json:"partition_key,omitempty"`
	EnqueuedTime           int64             `protobuf:"varint,5,opt,name=enqueued_time,json=enqueuedTime,proto3" json:"enqueued_time,omitempty"`
	DeadLetterSource       string            `protobuf:"bytes,6,opt,name=dead_letter_source,json=deadLetterSource,proto3" json:"dead_letter_source,omitempty"`
	ScheduledEnqueueTime   int64             `protobuf:"varint,7,opt,name=scheduled_enqueue_time,json=scheduledEnqueueTime,proto3" json:"scheduled_enqueue_time,omitempty"`
	EnqueuedSequenceNumber int64             `protobuf:"varint,8,opt,name=enqueued_sequence_number,json=enqueuedSequenceNumber,proto3" json:"enqueued_sequence_number,omitempty"`
	ViaPartitionKey        string            `protobuf:"bytes,9,opt,name=via_partition_key,json=viaPartitionKey,proto3" json:"via_partition_key,omitempty"`
	Annotations            map[string]string `protobuf:"bytes,10,rep,name=annotations,proto3" json:"annotations,omitempty" protobuf_key:"bytes,1,opt,name=key,proto3" protobuf_val:"bytes,2,opt,name=value,proto3"`
	XXX_NoUnkeyedLiteral   struct{}          `json:"-"`
	XXX_unrecognized       []byte            `json:"-"`
	XXX_sizecache          int32             `json:"-"`
}

func (m *AzureSystemProperties) Reset()         { *m = AzureSystemProperties{} }
func (m *AzureSystemProperties) String() string { return proto.CompactTextString(m) }
func (*AzureSystemProperties) ProtoMessage()    {}
func (*AzureSystemProperties) Descriptor() ([]byte, []int) {
	return fileDescriptor_fd29aeb29eb6a1a2, []int{1}
}

func (m *AzureSystemProperties) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_AzureSystemProperties.Unmarshal(m, b)
}
func (m *AzureSystemProperties) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_AzureSystemProperties.Marshal(b, m, deterministic)
}
func (m *AzureSystemProperties) XXX_Merge(src proto.Message) {
	xxx_messageInfo_AzureSystemProperties.Merge(m, src)
}
func (m *AzureSystemProperties) XXX_Size() int {
	return xxx_messageInfo_AzureSystemProperties.Size(m)
}
func (m *AzureSystemProperties) XXX_DiscardUnknown() {
	xxx_messageInfo_AzureSystemProperties.DiscardUnknown(m)
}

var xxx_messageInfo_AzureSystemProperties proto.InternalMessageInfo

func (m *AzureSystemProperties) GetLockedUntil() int64 {
	if m != nil {
		return m.LockedUntil
	}
	return 0
}

func (m *AzureSystemProperties) GetSequenceNumber() int64 {
	if m != nil {
		return m.SequenceNumber
	}
	return 0
}

func (m *AzureSystemProperties) GetPartitionId() int32 {
	if m != nil {
		return m.PartitionId
	}
	return 0
}

func (m *AzureSystemProperties) GetPartitionKey() string {
	if m != nil {
		return m.PartitionKey
	}
	return ""
}

func (m *AzureSystemProperties) GetEnqueuedTime() int64 {
	if m != nil {
		return m.EnqueuedTime
	}
	return 0
}

func (m *AzureSystemProperties) GetDeadLetterSource() string {
	if m != nil {
		return m.DeadLetterSource
	}
	return ""
}

func (m *AzureSystemProperties) GetScheduledEnqueueTime() int64 {
	if m != nil {
		return m.ScheduledEnqueueTime
	}
	return 0
}

func (m *AzureSystemProperties) GetEnqueuedSequenceNumber() int64 {
	if m != nil {
		return m.EnqueuedSequenceNumber
	}
	return 0
}

func (m *AzureSystemProperties) GetViaPartitionKey() string {
	if m != nil {
		return m.ViaPartitionKey
	}
	return ""
}

func (m *AzureSystemProperties) GetAnnotations() map[string]string {
	if m != nil {
		return m.Annotations
	}
	return nil
}

func init() {
	proto.RegisterType((*AzureServiceBus)(nil), "protos.records.AzureServiceBus")
	proto.RegisterMapType((map[string]string)(nil), "protos.records.AzureServiceBus.UserPropertiesEntry")
	proto.RegisterType((*AzureSystemProperties)(nil), "protos.records.AzureSystemProperties")
	proto.RegisterMapType((map[string]string)(nil), "protos.records.AzureSystemProperties.AnnotationsEntry")
}

func init() { proto.RegisterFile("records/azure-service-bus.proto", fileDescriptor_fd29aeb29eb6a1a2) }

var fileDescriptor_fd29aeb29eb6a1a2 = []byte{
	// 684 bytes of a gzipped FileDescriptorProto
	0x1f, 0x8b, 0x08, 0x00, 0x00, 0x00, 0x00, 0x00, 0x02, 0xff, 0x94, 0x53, 0x6b, 0x6b, 0xdb, 0x4a,
	0x10, 0xc5, 0x76, 0x9c, 0xc4, 0xeb, 0x67, 0xf6, 0xe6, 0x86, 0xbd, 0x81, 0x4b, 0xdd, 0x94, 0x50,
	0xb7, 0x34, 0x36, 0x24, 0xa5, 0x84, 0x52, 0x02, 0x49, 0x09, 0xc5, 0xb4, 0x94, 0x20, 0x3b, 0x50,
	0x4a, 0x61, 0x91, 0xb5, 0xd3, 0x64, 0x89, 0xa4, 0x55, 0xf6, 0x61, 0x50, 0x7f, 0x49, 0x7f, 0x48,
	0x7f, 0x60, 0xd9, 0x5d, 0xc9, 0x8f, 0x90, 0x0f, 0xed, 0x27, 0x69, 0xce, 0x9c, 0x39, 0xf3, 0xd0,
	0x11, 0x7a, 0x22, 0x21, 0x12, 0x92, 0xa9, 0x51, 0xf8, 0xc3, 0x48, 0x38, 0x52, 0x20, 0xe7, 0x3c,
	0x82, 0xa3, 0x99, 0x51, 0xc3, 0x4c, 0x0a, 0x2d, 0x70, 0xc7, 0x3d, 0xd4, 0xb0, 0xe0, 0x1d, 0xfc,
	0xac, 0xa3, 0xee, 0xb9, 0xe5, 0x4e, 0x3c, 0xf5, 0xc2, 0x28, 0xfc, 0x14, 0xb5, 0x22, 0x91, 0x6a,
	0x48, 0x35, 0xd5, 0x79, 0x06, 0xa4, 0xd2, 0xaf, 0x0c, 0x1a, 0x41, 0xb3, 0xc0, 0xa6, 0x79, 0x06,
	0xf8, 0x10, 0x75, 0x22, 0x21, 0x25, 0xc4, 0xa1, 0xe6, 0x22, 0xa5, 0x9c, 0x91, 0xaa, 0x23, 0xb5,
	0x57, 0xd0, 0x31, 0xc3, 0xbb, 0xa8, 0x3e, 0x0f, 0x63, 0x03, 0xa4, 0xd6, 0xaf, 0x0c, 0x5a, 0x81,
	0x0f, 0x6c, 0x31, 0x83, 0x98, 0xcf, 0x41, 0xe6, 0x34, 0x12, 0x26, 0xd5, 0x64, 0xa3, 0x5f, 0x19,
	0xb4, 0x83, 0x76, 0x89, 0xbe, 0xb7, 0x20, 0xfe, 0x1f, 0x21, 0x05, 0x4a, 0x15, 0xfa, 0x75, 0xa7,
	0xdf, 0x28, 0x90, 0x31, 0xb3, 0x2a, 0x37, 0x52, 0x98, 0x8c, 0x2a, 0xb8, 0x37, 0x90, 0x46, 0x40,
	0x36, 0xbd, 0x8a, 0x43, 0x27, 0x05, 0x88, 0x3b, 0xa8, 0xca, 0x19, 0xd9, 0x72, 0xd5, 0x55, 0xee,
	0x46, 0x8a, 0xc3, 0x19, 0xc4, 0x64, 0xdb, 0x41, 0x3e, 0xc0, 0xff, 0xa1, 0x6d, 0x09, 0x59, 0x9c,
	0x53, 0x2d, 0x48, 0xc3, 0x25, 0xb6, 0x5c, 0x3c, 0x15, 0xf8, 0x05, 0xda, 0x29, 0x53, 0xd4, 0x37,
	0xe4, 0x8c, 0x20, 0xc7, 0xe9, 0x14, 0x9c, 0x0f, 0x16, 0x1e, 0x33, 0xdb, 0x4b, 0x0b, 0xd2, 0xf4,
	0xbd, 0xb4, 0xc0, 0x3d, 0x54, 0xd3, 0x3a, 0x26, 0xad, 0x7e, 0x65, 0x50, 0x0b, 0xec, 0xab, 0xdd,
	0x29, 0x16, 0xd1, 0x1d, 0xd5, 0xe2, 0x0e, 0x52, 0xd2, 0xf6, 0x3b, 0x59, 0x64, 0x6a, 0x01, 0x1c,
	0xa0, 0x1d, 0x95, 0x2b, 0x0d, 0x09, 0xcd, 0xa4, 0xc8, 0x40, 0x6a, 0x0e, 0x8a, 0x74, 0xfa, 0x95,
	0x41, 0xf3, 0xf8, 0x70, 0xb8, 0xfe, 0xe5, 0x86, 0xfe, 0xab, 0x39, 0xf6, 0xd5, 0x82, 0x1c, 0xf4,
	0xd4, 0x03, 0x04, 0x7f, 0x43, 0x5d, 0xa3, 0x40, 0xae, 0x2a, 0x76, 0xfb, 0xb5, 0x41, 0xf3, 0xf8,
	0xe4, 0x71, 0xc5, 0x85, 0x0f, 0x86, 0xd7, 0x0a, 0xe4, 0x52, 0xe8, 0x32, 0xd5, 0x32, 0x0f, 0x3a,
	0x66, 0x0d, 0xc4, 0x7b, 0x68, 0xf3, 0xbb, 0x90, 0x49, 0xa8, 0x49, 0xcf, 0x5d, 0xbf, 0x88, 0xf6,
	0xcf, 0xd1, 0x3f, 0x8f, 0x94, 0xdb, 0x8b, 0xdc, 0x41, 0x5e, 0x38, 0xca, 0xbe, 0x2e, 0x2d, 0xe2,
	0x0d, 0xe4, 0x83, 0xb7, 0xd5, 0xd3, 0xca, 0xc1, 0xaf, 0x0d, 0xf4, 0xef, 0xa3, 0x4b, 0x5a, 0x83,
	0xda, 0x9b, 0x01, 0xa3, 0x26, 0xd5, 0x3c, 0x76, 0x72, 0xb5, 0xa0, 0xe9, 0xb1, 0x6b, 0x0b, 0xe1,
	0xe7, 0xa8, 0x5b, 0xfa, 0x82, 0xa6, 0x26, 0x99, 0x81, 0x74, 0x0d, 0x6a, 0x41, 0xa7, 0x84, 0x3f,
	0x3b, 0xd4, 0x6a, 0x65, 0xa1, 0xd4, 0xbc, 0xf4, 0xb1, 0x75, 0x6a, 0x3d, 0x68, 0x2e, 0xb0, 0x31,
	0xc3, 0xcf, 0x50, 0x7b, 0x49, 0xb1, 0xe3, 0x6f, 0xb8, 0x51, 0x97, 0x75, 0x1f, 0x21, 0xb7, 0x24,
	0x48, 0xef, 0x0d, 0x18, 0x60, 0x54, 0xf3, 0x04, 0x9c, 0x61, 0x6b, 0x41, 0xab, 0x04, 0xa7, 0x3c,
	0x01, 0xfc, 0x0a, 0x61, 0x06, 0x21, 0xa3, 0x31, 0x68, 0x0d, 0x92, 0x2a, 0x61, 0x64, 0xe1, 0xdb,
	0x46, 0xd0, 0xb3, 0x99, 0x4f, 0x2e, 0x31, 0x71, 0x38, 0x7e, 0x8d, 0xf6, 0x54, 0x74, 0x0b, 0xcc,
	0xc4, 0xc0, 0x68, 0xa1, 0xe3, 0xb5, 0xb7, 0x9c, 0xf6, 0xee, 0x22, 0x7b, 0xe9, 0x93, 0xae, 0xc7,
	0x29, 0x22, 0x8b, 0x41, 0x1e, 0x9e, 0x60, 0xdb, 0xd5, 0xed, 0x95, 0xf9, 0xc9, 0xfa, 0x29, 0x5e,
	0xa2, 0x9d, 0x39, 0x0f, 0xe9, 0xfa, 0xae, 0xfe, 0x6f, 0xe8, 0xce, 0x79, 0x78, 0xb5, 0xba, 0xee,
	0x17, 0xd4, 0x0c, 0xd3, 0x54, 0x68, 0xf7, 0xa7, 0x2b, 0x82, 0x9c, 0xa3, 0xde, 0xfc, 0x91, 0x47,
	0x87, 0xe7, 0xcb, 0x42, 0x6f, 0xaa, 0x55, 0xa9, 0xfd, 0x33, 0xd4, 0x7b, 0x48, 0xf8, 0x1b, 0xdb,
	0x5c, 0x9c, 0x7d, 0x7d, 0x77, 0xc3, 0xf5, 0xad, 0x99, 0x0d, 0x23, 0x91, 0x8c, 0x66, 0xa1, 0x8e,
	0x6e, 0x23, 0x21, 0xb3, 0x51, 0x16, 0xbb, 0x2d, 0x8f, 0xec, 0xd1, 0x92, 0x50, 0x8d, 0x66, 0x86,
	0xc7, 0x6c, 0x74, 0x23, 0x46, 0x7e, 0xe6, 0x51, 0x31, 0xf3, 0x6c, 0xd3, 0xc5, 0x27, 0xbf, 0x03,
	0x00, 0x00, 0xff, 0xff, 0xf9, 0xff, 0x1a, 0xa8, 0x4b, 0x05, 0x00, 0x00,
}
