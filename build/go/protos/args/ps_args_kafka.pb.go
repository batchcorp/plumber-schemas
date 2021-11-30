// Code generated by protoc-gen-go. DO NOT EDIT.
// source: ps_args_kafka.proto

package args

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

type SASLType int32

const (
	SASLType_NONE  SASLType = 0
	SASLType_PLAIN SASLType = 1
	SASLType_SCRAM SASLType = 2
)

var SASLType_name = map[int32]string{
	0: "NONE",
	1: "PLAIN",
	2: "SCRAM",
}

var SASLType_value = map[string]int32{
	"NONE":  0,
	"PLAIN": 1,
	"SCRAM": 2,
}

func (x SASLType) String() string {
	return proto.EnumName(SASLType_name, int32(x))
}

func (SASLType) EnumDescriptor() ([]byte, []int) {
	return fileDescriptor_5f98a21d9d819f0c, []int{0}
}

type KafkaConn struct {
	// @gotags: kong:"help='Kafka broker address (you may specify this flag multiple times',env=PLUMBER_RELAY_KAFKA_ADDRESS,default='localhost:9092',required"
	Address []string `protobuf:"bytes,1,rep,name=address,proto3" json:"address,omitempty" kong:"help='Kafka broker address (you may specify this flag multiple times',env=PLUMBER_RELAY_KAFKA_ADDRESS,default='localhost:9092',required"`
	// @gotags: kong:"help='Connect timeout',env=PLUMBER_RELAY_TIMEOUT_SECONDS,default=10"
	TimeoutSeconds int32 `protobuf:"varint,2,opt,name=timeout_seconds,json=timeoutSeconds,proto3" json:"timeout_seconds,omitempty" kong:"help='Connect timeout',env=PLUMBER_RELAY_TIMEOUT_SECONDS,default=10"`
	// @gotags: kong:"help='Enable TLS usage',env=PLUMBER_RELAY_USE_TLS"
	UseTls bool `protobuf:"varint,3,opt,name=use_tls,json=useTls,proto3" json:"use_tls,omitempty" kong:"help='Enable TLS usage',env=PLUMBER_RELAY_USE_TLS"`
	// @gotags: kong:"help='Allow insecure TLS usage',env=PLUMBER_RELAY_KAFKA_INSECURE_TLS"
	InsecureTls bool `protobuf:"varint,4,opt,name=insecure_tls,json=insecureTls,proto3" json:"insecure_tls,omitempty" kong:"help='Allow insecure TLS usage',env=PLUMBER_RELAY_KAFKA_INSECURE_TLS"`
	// @gotags: kong:"help='SASL authentication type (options: none, plain, scram)',env=PLUMBER_RELAY_KAFKA_SASL_TYPE,type=pbenum,pbenum_lowercase,default=none"
	SaslType SASLType `protobuf:"varint,5,opt,name=sasl_type,json=saslType,proto3,enum=protos.args.SASLType" json:"sasl_type,omitempty" kong:"help='SASL authentication type (options: none, plain, scram)',env=PLUMBER_RELAY_KAFKA_SASL_TYPE,type=pbenum,pbenum_lowercase,default=none"`
	// @gotags: kong:"help='SASL Username',env=PLUMBER_RELAY_KAFKA_USERNAME"
	SaslUsername string `protobuf:"bytes,6,opt,name=sasl_username,json=saslUsername,proto3" json:"sasl_username,omitempty" kong:"help='SASL Username',env=PLUMBER_RELAY_KAFKA_USERNAME"`
	// Required if sasl_type is not NONE
	// @gotags: kong:"help='SASL Password. If omitted, you will be prompted for the password',env=PLUMBER_RELAY_KAFKA_PASSWORD"
	SaslPassword         string   `protobuf:"bytes,7,opt,name=sasl_password,json=saslPassword,proto3" json:"sasl_password,omitempty" kong:"help='SASL Password. If omitted, you will be prompted for the password',env=PLUMBER_RELAY_KAFKA_PASSWORD"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *KafkaConn) Reset()         { *m = KafkaConn{} }
func (m *KafkaConn) String() string { return proto.CompactTextString(m) }
func (*KafkaConn) ProtoMessage()    {}
func (*KafkaConn) Descriptor() ([]byte, []int) {
	return fileDescriptor_5f98a21d9d819f0c, []int{0}
}

func (m *KafkaConn) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_KafkaConn.Unmarshal(m, b)
}
func (m *KafkaConn) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_KafkaConn.Marshal(b, m, deterministic)
}
func (m *KafkaConn) XXX_Merge(src proto.Message) {
	xxx_messageInfo_KafkaConn.Merge(m, src)
}
func (m *KafkaConn) XXX_Size() int {
	return xxx_messageInfo_KafkaConn.Size(m)
}
func (m *KafkaConn) XXX_DiscardUnknown() {
	xxx_messageInfo_KafkaConn.DiscardUnknown(m)
}

var xxx_messageInfo_KafkaConn proto.InternalMessageInfo

func (m *KafkaConn) GetAddress() []string {
	if m != nil {
		return m.Address
	}
	return nil
}

func (m *KafkaConn) GetTimeoutSeconds() int32 {
	if m != nil {
		return m.TimeoutSeconds
	}
	return 0
}

func (m *KafkaConn) GetUseTls() bool {
	if m != nil {
		return m.UseTls
	}
	return false
}

func (m *KafkaConn) GetInsecureTls() bool {
	if m != nil {
		return m.InsecureTls
	}
	return false
}

func (m *KafkaConn) GetSaslType() SASLType {
	if m != nil {
		return m.SaslType
	}
	return SASLType_NONE
}

func (m *KafkaConn) GetSaslUsername() string {
	if m != nil {
		return m.SaslUsername
	}
	return ""
}

func (m *KafkaConn) GetSaslPassword() string {
	if m != nil {
		return m.SaslPassword
	}
	return ""
}

type KafkaReadArgs struct {
	// @gotags: kong:"help='Topic(s) to read from',required"
	Topic []string `protobuf:"bytes,1,rep,name=topic,proto3" json:"topic,omitempty" kong:"help='Topic(s) to read from',required"`
	// @gotags: kong:"help='Specify what offset the consumer should read from (only works if --use-consumer-group is false)',default=0"
	ReadOffset int64 `protobuf:"varint,2,opt,name=read_offset,json=readOffset,proto3" json:"read_offset,omitempty" kong:"help='Specify what offset the consumer should read from (only works if --use-consumer-group is false)',default=0"`
	// @gotags: kong:"help='Whether plumber should use a consumer group',default=true"
	UseConsumerGroup bool `protobuf:"varint,3,opt,name=use_consumer_group,json=useConsumerGroup,proto3" json:"use_consumer_group,omitempty" kong:"help='Whether plumber should use a consumer group',default=true"`
	// @gotags: kong:"help='Specify a specific group-id to use when reading from kafka',default=plumber"
	ConsumerGroupName string `protobuf:"bytes,4,opt,name=consumer_group_name,json=consumerGroupName,proto3" json:"consumer_group_name,omitempty" kong:"help='Specify a specific group-id to use when reading from kafka',default=plumber"`
	// @gotags: kong:"help='How long to wait for new data when reading batches of messages',default=1"
	MaxWaitSeconds int32 `protobuf:"varint,5,opt,name=max_wait_seconds,json=maxWaitSeconds,proto3" json:"max_wait_seconds,omitempty" kong:"help='How long to wait for new data when reading batches of messages',default=1"`
	// @gotags: kong:"help='Minimum number of bytes to fetch in a single kafka request (throughput optimization)',default=1"
	MinBytes int32 `protobuf:"varint,6,opt,name=min_bytes,json=minBytes,proto3" json:"min_bytes,omitempty" kong:"help='Minimum number of bytes to fetch in a single kafka request (throughput optimization)',default=1"`
	// @gotags: kong:"help='Maximum number of bytes to fetch in a single kafka request (throughput optimization)',default=1"
	MaxBytes int32 `protobuf:"varint,7,opt,name=max_bytes,json=maxBytes,proto3" json:"max_bytes,omitempty" kong:"help='Maximum number of bytes to fetch in a single kafka request (throughput optimization)',default=1"`
	// @gotags: kong:"help='How often to commit offsets to broker (0 = synchronous)',default=5"
	CommitIntervalSeconds int32 `protobuf:"varint,8,opt,name=commit_interval_seconds,json=commitIntervalSeconds,proto3" json:"commit_interval_seconds,omitempty" kong:"help='How often to commit offsets to broker (0 = synchronous)',default=5"`
	// @gotags: kong:"help='How long a coordinator will wait for member joins as part of a rebalance',default=0"
	RebalanceTimeoutSeconds int32 `protobuf:"varint,9,opt,name=rebalance_timeout_seconds,json=rebalanceTimeoutSeconds,proto3" json:"rebalance_timeout_seconds,omitempty" kong:"help='How long a coordinator will wait for member joins as part of a rebalance',default=0"`
	// @gotags: kong:"help='Internal library queue capacity (throughput optimization)',default=1"
	QueueCapacity int32 `protobuf:"varint,10,opt,name=queue_capacity,json=queueCapacity,proto3" json:"queue_capacity,omitempty" kong:"help='Internal library queue capacity (throughput optimization)',default=1"`
	// @gotags: kong:"help='Display consumer offset stats during read'"
	IncludeOffsetInfo bool `protobuf:"varint,11,opt,name=include_offset_info,json=includeOffsetInfo,proto3" json:"include_offset_info,omitempty" kong:"help='Display consumer offset stats during read'"`
	// @gotags: kong:"help='Only lookup consumer group lag',group=lag"
	Lag bool `protobuf:"varint,12,opt,name=lag,proto3" json:"lag,omitempty" kong:"help='Only lookup consumer group lag',group=lag"`
	// @gotags: kong:"help='What consumer group to lookup lag for (required if --lag is set)',group=lag"
	LagConsumerGroup     string   `protobuf:"bytes,13,opt,name=lag_consumer_group,json=lagConsumerGroup,proto3" json:"lag_consumer_group,omitempty" kong:"help='What consumer group to lookup lag for (required if --lag is set)',group=lag"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *KafkaReadArgs) Reset()         { *m = KafkaReadArgs{} }
func (m *KafkaReadArgs) String() string { return proto.CompactTextString(m) }
func (*KafkaReadArgs) ProtoMessage()    {}
func (*KafkaReadArgs) Descriptor() ([]byte, []int) {
	return fileDescriptor_5f98a21d9d819f0c, []int{1}
}

func (m *KafkaReadArgs) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_KafkaReadArgs.Unmarshal(m, b)
}
func (m *KafkaReadArgs) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_KafkaReadArgs.Marshal(b, m, deterministic)
}
func (m *KafkaReadArgs) XXX_Merge(src proto.Message) {
	xxx_messageInfo_KafkaReadArgs.Merge(m, src)
}
func (m *KafkaReadArgs) XXX_Size() int {
	return xxx_messageInfo_KafkaReadArgs.Size(m)
}
func (m *KafkaReadArgs) XXX_DiscardUnknown() {
	xxx_messageInfo_KafkaReadArgs.DiscardUnknown(m)
}

var xxx_messageInfo_KafkaReadArgs proto.InternalMessageInfo

func (m *KafkaReadArgs) GetTopic() []string {
	if m != nil {
		return m.Topic
	}
	return nil
}

func (m *KafkaReadArgs) GetReadOffset() int64 {
	if m != nil {
		return m.ReadOffset
	}
	return 0
}

func (m *KafkaReadArgs) GetUseConsumerGroup() bool {
	if m != nil {
		return m.UseConsumerGroup
	}
	return false
}

func (m *KafkaReadArgs) GetConsumerGroupName() string {
	if m != nil {
		return m.ConsumerGroupName
	}
	return ""
}

func (m *KafkaReadArgs) GetMaxWaitSeconds() int32 {
	if m != nil {
		return m.MaxWaitSeconds
	}
	return 0
}

func (m *KafkaReadArgs) GetMinBytes() int32 {
	if m != nil {
		return m.MinBytes
	}
	return 0
}

func (m *KafkaReadArgs) GetMaxBytes() int32 {
	if m != nil {
		return m.MaxBytes
	}
	return 0
}

func (m *KafkaReadArgs) GetCommitIntervalSeconds() int32 {
	if m != nil {
		return m.CommitIntervalSeconds
	}
	return 0
}

func (m *KafkaReadArgs) GetRebalanceTimeoutSeconds() int32 {
	if m != nil {
		return m.RebalanceTimeoutSeconds
	}
	return 0
}

func (m *KafkaReadArgs) GetQueueCapacity() int32 {
	if m != nil {
		return m.QueueCapacity
	}
	return 0
}

func (m *KafkaReadArgs) GetIncludeOffsetInfo() bool {
	if m != nil {
		return m.IncludeOffsetInfo
	}
	return false
}

func (m *KafkaReadArgs) GetLag() bool {
	if m != nil {
		return m.Lag
	}
	return false
}

func (m *KafkaReadArgs) GetLagConsumerGroup() string {
	if m != nil {
		return m.LagConsumerGroup
	}
	return ""
}

type KafkaWriteArgs struct {
	// @gotags: kong:"help='Key to write to kafka (optional)'"
	Key string `protobuf:"bytes,1,opt,name=key,proto3" json:"key,omitempty" kong:"help='Key to write to kafka (optional)'"`
	// @gotags: kong:"help='Add one or more headers (optional; repeat flags to specify multiple)'"
	Headers map[string]string `protobuf:"bytes,2,rep,name=headers,proto3" json:"headers,omitempty" protobuf_key:"bytes,1,opt,name=key,proto3" protobuf_val:"bytes,2,opt,name=value,proto3" kong:"help='Add one or more headers (optional; repeat flags to specify multiple)'"`
	// @gotags: kong:"help='Topic(s) to write to',required"
	Topics               []string `protobuf:"bytes,3,rep,name=topics,proto3" json:"topics,omitempty" kong:"help='Topic(s) to write to',required"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *KafkaWriteArgs) Reset()         { *m = KafkaWriteArgs{} }
func (m *KafkaWriteArgs) String() string { return proto.CompactTextString(m) }
func (*KafkaWriteArgs) ProtoMessage()    {}
func (*KafkaWriteArgs) Descriptor() ([]byte, []int) {
	return fileDescriptor_5f98a21d9d819f0c, []int{2}
}

func (m *KafkaWriteArgs) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_KafkaWriteArgs.Unmarshal(m, b)
}
func (m *KafkaWriteArgs) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_KafkaWriteArgs.Marshal(b, m, deterministic)
}
func (m *KafkaWriteArgs) XXX_Merge(src proto.Message) {
	xxx_messageInfo_KafkaWriteArgs.Merge(m, src)
}
func (m *KafkaWriteArgs) XXX_Size() int {
	return xxx_messageInfo_KafkaWriteArgs.Size(m)
}
func (m *KafkaWriteArgs) XXX_DiscardUnknown() {
	xxx_messageInfo_KafkaWriteArgs.DiscardUnknown(m)
}

var xxx_messageInfo_KafkaWriteArgs proto.InternalMessageInfo

func (m *KafkaWriteArgs) GetKey() string {
	if m != nil {
		return m.Key
	}
	return ""
}

func (m *KafkaWriteArgs) GetHeaders() map[string]string {
	if m != nil {
		return m.Headers
	}
	return nil
}

func (m *KafkaWriteArgs) GetTopics() []string {
	if m != nil {
		return m.Topics
	}
	return nil
}

type KafkaRelayArgs struct {
	// @gotags: kong:"help='Topic(s) to read, write or get lag stats for',env=PLUMBER_RELAY_KAFKA_TOPIC,required"
	Topic []string `protobuf:"bytes,1,rep,name=topic,proto3" json:"topic,omitempty" kong:"help='Topic(s) to read, write or get lag stats for',env=PLUMBER_RELAY_KAFKA_TOPIC,required"`
	// @gotags: kong:"help='Specify what offset the consumer should read from (only works if --use-consumer-group is false)',env=PLUMBER_RELAY_KAFKA_READ_OFFSET,default=0"
	ReadOffset int64 `protobuf:"varint,2,opt,name=read_offset,json=readOffset,proto3" json:"read_offset,omitempty" kong:"help='Specify what offset the consumer should read from (only works if --use-consumer-group is false)',env=PLUMBER_RELAY_KAFKA_READ_OFFSET,default=0"`
	// @gotags: kong:"help='Whether plumber should use a consumer group',env=PLUMBER_RELAY_KAFKA_USE_CONSUMER_GROUP,default=true"
	UseConsumerGroup bool `protobuf:"varint,3,opt,name=use_consumer_group,json=useConsumerGroup,proto3" json:"use_consumer_group,omitempty" kong:"help='Whether plumber should use a consumer group',env=PLUMBER_RELAY_KAFKA_USE_CONSUMER_GROUP,default=true"`
	// @gotags: kong:"help='Specify a specific group-id to use when reading from kafka',env=PLUMBER_RELAY_KAFKA_GROUP_ID,default=plumber"
	ConsumerGroupName string `protobuf:"bytes,4,opt,name=consumer_group_name,json=consumerGroupName,proto3" json:"consumer_group_name,omitempty" kong:"help='Specify a specific group-id to use when reading from kafka',env=PLUMBER_RELAY_KAFKA_GROUP_ID,default=plumber"`
	// @gotags: kong:"help='How long to wait for new data when reading batches of messages',env=PLUMBER_RELAY_KAFKA_MAX_WAIT,default=5"
	MaxWaitSeconds int32 `protobuf:"varint,5,opt,name=max_wait_seconds,json=maxWaitSeconds,proto3" json:"max_wait_seconds,omitempty" kong:"help='How long to wait for new data when reading batches of messages',env=PLUMBER_RELAY_KAFKA_MAX_WAIT,default=5"`
	// @gotags: kong:"help='Minimum number of bytes to fetch in a single kafka request (throughput optimization)',env=PLUMBER_RELAY_KAFKA_MIN_BYTES,default=1048576"
	MinBytes int32 `protobuf:"varint,6,opt,name=min_bytes,json=minBytes,proto3" json:"min_bytes,omitempty" kong:"help='Minimum number of bytes to fetch in a single kafka request (throughput optimization)',env=PLUMBER_RELAY_KAFKA_MIN_BYTES,default=1048576"`
	// @gotags: kong:"help='Maximum number of bytes to fetch in a single kafka request (throughput optimization)',env=PLUMBER_RELAY_KAFKA_MAX_BYTES,default=1048576"
	MaxBytes int32 `protobuf:"varint,7,opt,name=max_bytes,json=maxBytes,proto3" json:"max_bytes,omitempty" kong:"help='Maximum number of bytes to fetch in a single kafka request (throughput optimization)',env=PLUMBER_RELAY_KAFKA_MAX_BYTES,default=1048576"`
	// @gotags: kong:"help='How often to commit offsets to broker (0 = synchronous)',env=PLUMBER_RELAY_KAFKA_COMMIT_INTERVAL,default=5"
	CommitIntervalSeconds int32 `protobuf:"varint,8,opt,name=commit_interval_seconds,json=commitIntervalSeconds,proto3" json:"commit_interval_seconds,omitempty" kong:"help='How often to commit offsets to broker (0 = synchronous)',env=PLUMBER_RELAY_KAFKA_COMMIT_INTERVAL,default=5"`
	// @gotags: kong:"help='How long a coordinator will wait for member joins as part of a rebalance',env=PLUMBER_RELAY_KAFKA_REBALANCE_TIMEOUT,default=5"
	RebalanceTimeoutSeconds int32 `protobuf:"varint,9,opt,name=rebalance_timeout_seconds,json=rebalanceTimeoutSeconds,proto3" json:"rebalance_timeout_seconds,omitempty" kong:"help='How long a coordinator will wait for member joins as part of a rebalance',env=PLUMBER_RELAY_KAFKA_REBALANCE_TIMEOUT,default=5"`
	// @gotags: kong:"help='Internal library queue capacity (throughput optimization)',env=PLUMBER_RELAY_KAFKA_QUEUE_CAPACITY,default=1000"
	QueueCapacity        int32    `protobuf:"varint,10,opt,name=queue_capacity,json=queueCapacity,proto3" json:"queue_capacity,omitempty" kong:"help='Internal library queue capacity (throughput optimization)',env=PLUMBER_RELAY_KAFKA_QUEUE_CAPACITY,default=1000"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *KafkaRelayArgs) Reset()         { *m = KafkaRelayArgs{} }
func (m *KafkaRelayArgs) String() string { return proto.CompactTextString(m) }
func (*KafkaRelayArgs) ProtoMessage()    {}
func (*KafkaRelayArgs) Descriptor() ([]byte, []int) {
	return fileDescriptor_5f98a21d9d819f0c, []int{3}
}

func (m *KafkaRelayArgs) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_KafkaRelayArgs.Unmarshal(m, b)
}
func (m *KafkaRelayArgs) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_KafkaRelayArgs.Marshal(b, m, deterministic)
}
func (m *KafkaRelayArgs) XXX_Merge(src proto.Message) {
	xxx_messageInfo_KafkaRelayArgs.Merge(m, src)
}
func (m *KafkaRelayArgs) XXX_Size() int {
	return xxx_messageInfo_KafkaRelayArgs.Size(m)
}
func (m *KafkaRelayArgs) XXX_DiscardUnknown() {
	xxx_messageInfo_KafkaRelayArgs.DiscardUnknown(m)
}

var xxx_messageInfo_KafkaRelayArgs proto.InternalMessageInfo

func (m *KafkaRelayArgs) GetTopic() []string {
	if m != nil {
		return m.Topic
	}
	return nil
}

func (m *KafkaRelayArgs) GetReadOffset() int64 {
	if m != nil {
		return m.ReadOffset
	}
	return 0
}

func (m *KafkaRelayArgs) GetUseConsumerGroup() bool {
	if m != nil {
		return m.UseConsumerGroup
	}
	return false
}

func (m *KafkaRelayArgs) GetConsumerGroupName() string {
	if m != nil {
		return m.ConsumerGroupName
	}
	return ""
}

func (m *KafkaRelayArgs) GetMaxWaitSeconds() int32 {
	if m != nil {
		return m.MaxWaitSeconds
	}
	return 0
}

func (m *KafkaRelayArgs) GetMinBytes() int32 {
	if m != nil {
		return m.MinBytes
	}
	return 0
}

func (m *KafkaRelayArgs) GetMaxBytes() int32 {
	if m != nil {
		return m.MaxBytes
	}
	return 0
}

func (m *KafkaRelayArgs) GetCommitIntervalSeconds() int32 {
	if m != nil {
		return m.CommitIntervalSeconds
	}
	return 0
}

func (m *KafkaRelayArgs) GetRebalanceTimeoutSeconds() int32 {
	if m != nil {
		return m.RebalanceTimeoutSeconds
	}
	return 0
}

func (m *KafkaRelayArgs) GetQueueCapacity() int32 {
	if m != nil {
		return m.QueueCapacity
	}
	return 0
}

func init() {
	proto.RegisterEnum("protos.args.SASLType", SASLType_name, SASLType_value)
	proto.RegisterType((*KafkaConn)(nil), "protos.args.KafkaConn")
	proto.RegisterType((*KafkaReadArgs)(nil), "protos.args.KafkaReadArgs")
	proto.RegisterType((*KafkaWriteArgs)(nil), "protos.args.KafkaWriteArgs")
	proto.RegisterMapType((map[string]string)(nil), "protos.args.KafkaWriteArgs.HeadersEntry")
	proto.RegisterType((*KafkaRelayArgs)(nil), "protos.args.KafkaRelayArgs")
}

func init() { proto.RegisterFile("ps_args_kafka.proto", fileDescriptor_5f98a21d9d819f0c) }

var fileDescriptor_5f98a21d9d819f0c = []byte{
	// 686 bytes of a gzipped FileDescriptorProto
	0x1f, 0x8b, 0x08, 0x00, 0x00, 0x00, 0x00, 0x00, 0x02, 0xff, 0xec, 0x54, 0xcd, 0x6e, 0xd3, 0x4c,
	0x14, 0xfd, 0x9c, 0x7f, 0xdf, 0xfc, 0x7c, 0xee, 0xb4, 0xa5, 0x06, 0x16, 0x84, 0x20, 0x84, 0x55,
	0x41, 0x22, 0x15, 0x09, 0x41, 0x59, 0xa5, 0x51, 0x05, 0x15, 0x25, 0xad, 0x9c, 0xa2, 0x4a, 0x6c,
	0xac, 0x89, 0x7d, 0xe3, 0x58, 0xb5, 0x3d, 0x66, 0xc6, 0x6e, 0x93, 0x37, 0x60, 0xcf, 0x2b, 0xf0,
	0x14, 0x3c, 0x1d, 0x9a, 0xb1, 0x93, 0x36, 0x15, 0x0f, 0xc0, 0x82, 0x55, 0xee, 0x3d, 0xe7, 0x5c,
	0xcd, 0xf8, 0xdc, 0x33, 0x81, 0xed, 0x44, 0x38, 0x94, 0xfb, 0xc2, 0xb9, 0xa2, 0xb3, 0x2b, 0xda,
	0x4f, 0x38, 0x4b, 0x19, 0x69, 0xaa, 0x1f, 0xd1, 0x97, 0x44, 0xef, 0x7b, 0x09, 0xf4, 0x4f, 0x92,
	0x1c, 0xb1, 0x38, 0x26, 0x26, 0xd4, 0xa9, 0xe7, 0x71, 0x14, 0xc2, 0xd4, 0xba, 0x65, 0x4b, 0xb7,
	0x57, 0x2d, 0x79, 0x01, 0xff, 0xa7, 0x41, 0x84, 0x2c, 0x4b, 0x1d, 0x81, 0x2e, 0x8b, 0x3d, 0x61,
	0x96, 0xba, 0x9a, 0x55, 0xb5, 0x3b, 0x05, 0x3c, 0xc9, 0x51, 0xb2, 0x07, 0xf5, 0x4c, 0xa0, 0x93,
	0x86, 0xc2, 0x2c, 0x77, 0x35, 0xab, 0x61, 0xd7, 0x32, 0x81, 0x17, 0xa1, 0x20, 0x4f, 0xa1, 0x15,
	0xc4, 0x02, 0xdd, 0x8c, 0xe7, 0x6c, 0x45, 0xb1, 0xcd, 0x15, 0x26, 0x25, 0x07, 0xa0, 0x0b, 0x2a,
	0x42, 0x27, 0x5d, 0x26, 0x68, 0x56, 0xbb, 0x9a, 0xd5, 0x39, 0xd8, 0xed, 0xdf, 0xb9, 0x6d, 0x7f,
	0x32, 0x9c, 0x9c, 0x5e, 0x2c, 0x13, 0xb4, 0x1b, 0x52, 0x27, 0x2b, 0xf2, 0x0c, 0xda, 0x6a, 0x26,
	0x13, 0xc8, 0x63, 0x1a, 0xa1, 0x59, 0xeb, 0x6a, 0x96, 0x6e, 0xb7, 0x24, 0xf8, 0xa5, 0xc0, 0xd6,
	0xa2, 0x84, 0x0a, 0x71, 0xc3, 0xb8, 0x67, 0xd6, 0x6f, 0x45, 0xe7, 0x05, 0xd6, 0xfb, 0x51, 0x81,
	0xb6, 0xb2, 0xc2, 0x46, 0xea, 0x0d, 0xb9, 0x2f, 0xc8, 0x0e, 0x54, 0x53, 0x96, 0x04, 0x6e, 0x61,
	0x46, 0xde, 0x90, 0x27, 0xd0, 0xe4, 0x48, 0x3d, 0x87, 0xcd, 0x66, 0x02, 0x53, 0x65, 0x43, 0xd9,
	0x06, 0x09, 0x9d, 0x29, 0x84, 0xbc, 0x04, 0x22, 0x2d, 0x70, 0x59, 0x2c, 0xb2, 0x08, 0xb9, 0xe3,
	0x73, 0x96, 0x25, 0x85, 0x1b, 0x46, 0x26, 0x70, 0x54, 0x10, 0x1f, 0x24, 0x4e, 0xfa, 0xb0, 0xbd,
	0xa9, 0x74, 0xd4, 0x67, 0x54, 0xd4, 0x0d, 0xb7, 0xdc, 0xbb, 0xda, 0xb1, 0xfc, 0x16, 0x0b, 0x8c,
	0x88, 0x2e, 0x9c, 0x1b, 0x1a, 0xdc, 0xae, 0xa2, 0x9a, 0xaf, 0x22, 0xa2, 0x8b, 0x4b, 0x1a, 0xac,
	0x57, 0xf1, 0x18, 0xf4, 0x28, 0x88, 0x9d, 0xe9, 0x32, 0x45, 0xa1, 0x6c, 0xa9, 0xda, 0x8d, 0x28,
	0x88, 0x8f, 0x64, 0xaf, 0x48, 0xba, 0x28, 0xc8, 0x7a, 0x41, 0xd2, 0x45, 0x4e, 0xbe, 0x81, 0x3d,
	0x97, 0x45, 0x51, 0x90, 0x3a, 0x41, 0x9c, 0x22, 0xbf, 0xa6, 0xe1, 0xfa, 0xa8, 0x86, 0x92, 0xee,
	0xe6, 0xf4, 0x49, 0xc1, 0xae, 0x4e, 0x3c, 0x84, 0x87, 0x1c, 0xa7, 0x34, 0xa4, 0xb1, 0x8b, 0xce,
	0xfd, 0xbc, 0xe8, 0x6a, 0x72, 0x6f, 0x2d, 0xb8, 0xd8, 0x0c, 0xce, 0x73, 0xe8, 0x7c, 0xcb, 0x30,
	0x43, 0xc7, 0xa5, 0x09, 0x75, 0x83, 0x74, 0x69, 0x82, 0x1a, 0x68, 0x2b, 0x74, 0x54, 0x80, 0xd2,
	0xae, 0x20, 0x76, 0xc3, 0xcc, 0xc3, 0x62, 0x01, 0x4e, 0x10, 0xcf, 0x98, 0xd9, 0x54, 0xee, 0x6e,
	0x15, 0x54, 0xbe, 0x88, 0x93, 0x78, 0xc6, 0x88, 0x01, 0xe5, 0x90, 0xfa, 0x66, 0x4b, 0xf1, 0xb2,
	0x94, 0xeb, 0x09, 0xa9, 0x7f, 0x7f, 0x3d, 0x6d, 0xe5, 0xb7, 0x11, 0x52, 0x7f, 0x63, 0x3d, 0xbd,
	0x5f, 0x1a, 0x74, 0x54, 0x2a, 0x2e, 0x79, 0x90, 0xa2, 0x8a, 0x85, 0x01, 0xe5, 0x2b, 0x5c, 0x9a,
	0x9a, 0x9a, 0x90, 0x25, 0x39, 0x82, 0xfa, 0x1c, 0xa9, 0x87, 0x5c, 0xbe, 0x8a, 0xb2, 0xd5, 0x3c,
	0xb0, 0x36, 0x62, 0xbb, 0x39, 0xdf, 0xff, 0x98, 0x4b, 0x8f, 0xe3, 0x94, 0x2f, 0xed, 0xd5, 0x20,
	0x79, 0x00, 0x35, 0x95, 0x2f, 0xf9, 0x6e, 0x64, 0xda, 0x8a, 0xee, 0xd1, 0x21, 0xb4, 0xee, 0x0e,
	0xfc, 0xe1, 0xf4, 0x1d, 0xa8, 0x5e, 0xd3, 0x30, 0x43, 0x15, 0x45, 0xdd, 0xce, 0x9b, 0xc3, 0xd2,
	0x5b, 0xad, 0xf7, 0xb3, 0x5c, 0x5c, 0xde, 0xc6, 0x90, 0x2e, 0xff, 0x65, 0xfa, 0xef, 0xcc, 0xf4,
	0xfe, 0x3e, 0x34, 0x56, 0xff, 0x6c, 0xa4, 0x01, 0x95, 0xf1, 0xd9, 0xf8, 0xd8, 0xf8, 0x8f, 0xe8,
	0x50, 0x3d, 0x3f, 0x1d, 0x9e, 0x8c, 0x0d, 0x4d, 0x96, 0x93, 0x91, 0x3d, 0xfc, 0x6c, 0x94, 0x8e,
	0xde, 0x7f, 0x7d, 0xe7, 0x07, 0xe9, 0x3c, 0x9b, 0xf6, 0x5d, 0x16, 0x0d, 0xa6, 0x34, 0x75, 0xe7,
	0x2e, 0xe3, 0xc9, 0x20, 0x09, 0xb3, 0x68, 0x8a, 0xfc, 0x95, 0x70, 0xe7, 0x18, 0x51, 0x31, 0x98,
	0x66, 0x41, 0xe8, 0x0d, 0x7c, 0x36, 0xc8, 0x83, 0x38, 0x90, 0x41, 0x9c, 0xd6, 0x54, 0xf3, 0xfa,
	0x77, 0x00, 0x00, 0x00, 0xff, 0xff, 0x32, 0x61, 0x60, 0x13, 0x18, 0x06, 0x00, 0x00,
}
