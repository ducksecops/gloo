// Code generated by protoc-gen-gogo. DO NOT EDIT.
// source: github.com/solo-io/gloo/projects/gloo/api/v1/upstream.proto

package v1

import (
	bytes "bytes"
	fmt "fmt"
	math "math"

	_ "github.com/gogo/protobuf/gogoproto"
	proto "github.com/gogo/protobuf/proto"
	cluster "github.com/solo-io/gloo/projects/gloo/pkg/api/external/envoy/api/v2/cluster"
	core1 "github.com/solo-io/gloo/projects/gloo/pkg/api/external/envoy/api/v2/core"
	aws "github.com/solo-io/gloo/projects/gloo/pkg/api/v1/options/aws"
	ec2 "github.com/solo-io/gloo/projects/gloo/pkg/api/v1/options/aws/ec2"
	azure "github.com/solo-io/gloo/projects/gloo/pkg/api/v1/options/azure"
	consul "github.com/solo-io/gloo/projects/gloo/pkg/api/v1/options/consul"
	kubernetes "github.com/solo-io/gloo/projects/gloo/pkg/api/v1/options/kubernetes"
	pipe "github.com/solo-io/gloo/projects/gloo/pkg/api/v1/options/pipe"
	static "github.com/solo-io/gloo/projects/gloo/pkg/api/v1/options/static"
	_ "github.com/solo-io/protoc-gen-ext/extproto"
	core "github.com/solo-io/solo-kit/pkg/api/v1/resources/core"
)

// Reference imports to suppress errors if they are not otherwise used.
var _ = proto.Marshal
var _ = fmt.Errorf
var _ = math.Inf

// This is a compile-time assertion to ensure that this generated file
// is compatible with the proto package it is being compiled against.
// A compilation error at this line likely means your copy of the
// proto package needs to be updated.
const _ = proto.GoGoProtoPackageIsVersion3 // please upgrade the proto package

//
// Upstreams represent destination for routing HTTP requests. Upstreams can be compared to
// [clusters](https://www.envoyproxy.io/docs/envoy/latest/api-v2/api/v2/cds.proto) in Envoy terminology.
// Each upstream in Gloo has a type. Supported types include `static`, `kubernetes`, `aws`, `consul`, and more.
// Each upstream type is handled by a corresponding Gloo plugin. (plugins currently need to be compiled into Gloo)
type Upstream struct {
	// Status indicates the validation status of the resource. Status is read-only by clients, and set by gloo during validation
	Status core.Status `protobuf:"bytes,1,opt,name=status,proto3" json:"status" testdiff:"ignore"`
	// Metadata contains the object metadata for this resource
	Metadata core.Metadata `protobuf:"bytes,2,opt,name=metadata,proto3" json:"metadata"`
	// Upstreams and their configuration can be automatically by Gloo Discovery
	// if this upstream is created or modified by Discovery, metadata about the operation will be placed here.
	DiscoveryMetadata *DiscoveryMetadata `protobuf:"bytes,3,opt,name=discovery_metadata,json=discoveryMetadata,proto3" json:"discovery_metadata,omitempty"`
	SslConfig         *UpstreamSslConfig `protobuf:"bytes,4,opt,name=ssl_config,json=sslConfig,proto3" json:"ssl_config,omitempty"`
	// Circuit breakers for this upstream. if not set, the defaults ones from the Gloo settings will be used.
	// if those are not set, [envoy's defaults](https://www.envoyproxy.io/docs/envoy/latest/api-v2/api/v2/cluster/circuit_breaker.proto#envoy-api-msg-cluster-circuitbreakers)
	// will be used.
	CircuitBreakers    *CircuitBreakerConfig     `protobuf:"bytes,5,opt,name=circuit_breakers,json=circuitBreakers,proto3" json:"circuit_breakers,omitempty"`
	LoadBalancerConfig *LoadBalancerConfig       `protobuf:"bytes,6,opt,name=load_balancer_config,json=loadBalancerConfig,proto3" json:"load_balancer_config,omitempty"`
	ConnectionConfig   *ConnectionConfig         `protobuf:"bytes,7,opt,name=connection_config,json=connectionConfig,proto3" json:"connection_config,omitempty"`
	HealthChecks       []*core1.HealthCheck      `protobuf:"bytes,8,rep,name=health_checks,json=healthChecks,proto3" json:"health_checks,omitempty"`
	OutlierDetection   *cluster.OutlierDetection `protobuf:"bytes,9,opt,name=outlier_detection,json=outlierDetection,proto3" json:"outlier_detection,omitempty"`
	// Use http2 when communicating with this upstream
	// this field is evaluated `true` for upstreams
	// with a grpc service spec. otherwise defaults to `false`
	UseHttp2 bool `protobuf:"varint,10,opt,name=use_http2,json=useHttp2,proto3" json:"use_http2,omitempty"`
	// Note to developers: new Upstream plugins must be added to this oneof field
	// to be usable by Gloo. (plugins currently need to be compiled into Gloo)
	//
	// Types that are valid to be assigned to UpstreamType:
	//	*Upstream_Kube
	//	*Upstream_Static
	//	*Upstream_Pipe
	//	*Upstream_Aws
	//	*Upstream_Azure
	//	*Upstream_Consul
	//	*Upstream_AwsEc2
	UpstreamType         isUpstream_UpstreamType `protobuf_oneof:"upstream_type"`
	XXX_NoUnkeyedLiteral struct{}                `json:"-"`
	XXX_unrecognized     []byte                  `json:"-"`
	XXX_sizecache        int32                   `json:"-"`
}

func (m *Upstream) Reset()         { *m = Upstream{} }
func (m *Upstream) String() string { return proto.CompactTextString(m) }
func (*Upstream) ProtoMessage()    {}
func (*Upstream) Descriptor() ([]byte, []int) {
	return fileDescriptor_b74df493149f644d, []int{0}
}
func (m *Upstream) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_Upstream.Unmarshal(m, b)
}
func (m *Upstream) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_Upstream.Marshal(b, m, deterministic)
}
func (m *Upstream) XXX_Merge(src proto.Message) {
	xxx_messageInfo_Upstream.Merge(m, src)
}
func (m *Upstream) XXX_Size() int {
	return xxx_messageInfo_Upstream.Size(m)
}
func (m *Upstream) XXX_DiscardUnknown() {
	xxx_messageInfo_Upstream.DiscardUnknown(m)
}

var xxx_messageInfo_Upstream proto.InternalMessageInfo

type isUpstream_UpstreamType interface {
	isUpstream_UpstreamType()
	Equal(interface{}) bool
}

type Upstream_Kube struct {
	Kube *kubernetes.UpstreamSpec `protobuf:"bytes,11,opt,name=kube,proto3,oneof" json:"kube,omitempty"`
}
type Upstream_Static struct {
	Static *static.UpstreamSpec `protobuf:"bytes,12,opt,name=static,proto3,oneof" json:"static,omitempty"`
}
type Upstream_Pipe struct {
	Pipe *pipe.UpstreamSpec `protobuf:"bytes,13,opt,name=pipe,proto3,oneof" json:"pipe,omitempty"`
}
type Upstream_Aws struct {
	Aws *aws.UpstreamSpec `protobuf:"bytes,14,opt,name=aws,proto3,oneof" json:"aws,omitempty"`
}
type Upstream_Azure struct {
	Azure *azure.UpstreamSpec `protobuf:"bytes,15,opt,name=azure,proto3,oneof" json:"azure,omitempty"`
}
type Upstream_Consul struct {
	Consul *consul.UpstreamSpec `protobuf:"bytes,16,opt,name=consul,proto3,oneof" json:"consul,omitempty"`
}
type Upstream_AwsEc2 struct {
	AwsEc2 *ec2.UpstreamSpec `protobuf:"bytes,17,opt,name=aws_ec2,json=awsEc2,proto3,oneof" json:"aws_ec2,omitempty"`
}

func (*Upstream_Kube) isUpstream_UpstreamType()   {}
func (*Upstream_Static) isUpstream_UpstreamType() {}
func (*Upstream_Pipe) isUpstream_UpstreamType()   {}
func (*Upstream_Aws) isUpstream_UpstreamType()    {}
func (*Upstream_Azure) isUpstream_UpstreamType()  {}
func (*Upstream_Consul) isUpstream_UpstreamType() {}
func (*Upstream_AwsEc2) isUpstream_UpstreamType() {}

func (m *Upstream) GetUpstreamType() isUpstream_UpstreamType {
	if m != nil {
		return m.UpstreamType
	}
	return nil
}

func (m *Upstream) GetStatus() core.Status {
	if m != nil {
		return m.Status
	}
	return core.Status{}
}

func (m *Upstream) GetMetadata() core.Metadata {
	if m != nil {
		return m.Metadata
	}
	return core.Metadata{}
}

func (m *Upstream) GetDiscoveryMetadata() *DiscoveryMetadata {
	if m != nil {
		return m.DiscoveryMetadata
	}
	return nil
}

func (m *Upstream) GetSslConfig() *UpstreamSslConfig {
	if m != nil {
		return m.SslConfig
	}
	return nil
}

func (m *Upstream) GetCircuitBreakers() *CircuitBreakerConfig {
	if m != nil {
		return m.CircuitBreakers
	}
	return nil
}

func (m *Upstream) GetLoadBalancerConfig() *LoadBalancerConfig {
	if m != nil {
		return m.LoadBalancerConfig
	}
	return nil
}

func (m *Upstream) GetConnectionConfig() *ConnectionConfig {
	if m != nil {
		return m.ConnectionConfig
	}
	return nil
}

func (m *Upstream) GetHealthChecks() []*core1.HealthCheck {
	if m != nil {
		return m.HealthChecks
	}
	return nil
}

func (m *Upstream) GetOutlierDetection() *cluster.OutlierDetection {
	if m != nil {
		return m.OutlierDetection
	}
	return nil
}

func (m *Upstream) GetUseHttp2() bool {
	if m != nil {
		return m.UseHttp2
	}
	return false
}

func (m *Upstream) GetKube() *kubernetes.UpstreamSpec {
	if x, ok := m.GetUpstreamType().(*Upstream_Kube); ok {
		return x.Kube
	}
	return nil
}

func (m *Upstream) GetStatic() *static.UpstreamSpec {
	if x, ok := m.GetUpstreamType().(*Upstream_Static); ok {
		return x.Static
	}
	return nil
}

func (m *Upstream) GetPipe() *pipe.UpstreamSpec {
	if x, ok := m.GetUpstreamType().(*Upstream_Pipe); ok {
		return x.Pipe
	}
	return nil
}

func (m *Upstream) GetAws() *aws.UpstreamSpec {
	if x, ok := m.GetUpstreamType().(*Upstream_Aws); ok {
		return x.Aws
	}
	return nil
}

func (m *Upstream) GetAzure() *azure.UpstreamSpec {
	if x, ok := m.GetUpstreamType().(*Upstream_Azure); ok {
		return x.Azure
	}
	return nil
}

func (m *Upstream) GetConsul() *consul.UpstreamSpec {
	if x, ok := m.GetUpstreamType().(*Upstream_Consul); ok {
		return x.Consul
	}
	return nil
}

func (m *Upstream) GetAwsEc2() *ec2.UpstreamSpec {
	if x, ok := m.GetUpstreamType().(*Upstream_AwsEc2); ok {
		return x.AwsEc2
	}
	return nil
}

// XXX_OneofWrappers is for the internal use of the proto package.
func (*Upstream) XXX_OneofWrappers() []interface{} {
	return []interface{}{
		(*Upstream_Kube)(nil),
		(*Upstream_Static)(nil),
		(*Upstream_Pipe)(nil),
		(*Upstream_Aws)(nil),
		(*Upstream_Azure)(nil),
		(*Upstream_Consul)(nil),
		(*Upstream_AwsEc2)(nil),
	}
}

// created by discovery services
type DiscoveryMetadata struct {
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *DiscoveryMetadata) Reset()         { *m = DiscoveryMetadata{} }
func (m *DiscoveryMetadata) String() string { return proto.CompactTextString(m) }
func (*DiscoveryMetadata) ProtoMessage()    {}
func (*DiscoveryMetadata) Descriptor() ([]byte, []int) {
	return fileDescriptor_b74df493149f644d, []int{1}
}
func (m *DiscoveryMetadata) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_DiscoveryMetadata.Unmarshal(m, b)
}
func (m *DiscoveryMetadata) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_DiscoveryMetadata.Marshal(b, m, deterministic)
}
func (m *DiscoveryMetadata) XXX_Merge(src proto.Message) {
	xxx_messageInfo_DiscoveryMetadata.Merge(m, src)
}
func (m *DiscoveryMetadata) XXX_Size() int {
	return xxx_messageInfo_DiscoveryMetadata.Size(m)
}
func (m *DiscoveryMetadata) XXX_DiscardUnknown() {
	xxx_messageInfo_DiscoveryMetadata.DiscardUnknown(m)
}

var xxx_messageInfo_DiscoveryMetadata proto.InternalMessageInfo

func init() {
	proto.RegisterType((*Upstream)(nil), "gloo.solo.io.Upstream")
	proto.RegisterType((*DiscoveryMetadata)(nil), "gloo.solo.io.DiscoveryMetadata")
}

func init() {
	proto.RegisterFile("github.com/solo-io/gloo/projects/gloo/api/v1/upstream.proto", fileDescriptor_b74df493149f644d)
}

var fileDescriptor_b74df493149f644d = []byte{
	// 827 bytes of a gzipped FileDescriptorProto
	0x1f, 0x8b, 0x08, 0x00, 0x00, 0x00, 0x00, 0x00, 0x02, 0xff, 0x94, 0x96, 0xdf, 0x4e, 0x2b, 0x45,
	0x1c, 0xc7, 0x4f, 0x0f, 0x7b, 0x38, 0xed, 0x00, 0xd2, 0x8e, 0x44, 0x37, 0x47, 0x85, 0xa6, 0x26,
	0x1e, 0x24, 0x61, 0x56, 0x4a, 0x8c, 0xa6, 0x06, 0x63, 0x5a, 0x48, 0x48, 0x00, 0x4d, 0x96, 0x78,
	0xe3, 0xcd, 0x66, 0x3a, 0x1d, 0xda, 0xb1, 0xcb, 0xce, 0x66, 0x67, 0xb6, 0x80, 0x97, 0x3c, 0x8d,
	0x8f, 0xe0, 0x0b, 0x98, 0xf8, 0x14, 0x5c, 0xf8, 0x06, 0x98, 0x78, 0x6f, 0xe6, 0x5f, 0xe9, 0x1f,
	0x4a, 0xd7, 0x0b, 0xba, 0x3b, 0x33, 0xdf, 0xef, 0x67, 0x7f, 0xfc, 0x76, 0xe6, 0x9b, 0x05, 0xdf,
	0xf5, 0x99, 0x1c, 0xe4, 0x5d, 0x44, 0xf8, 0x75, 0x20, 0x78, 0xcc, 0xf7, 0x19, 0x0f, 0xfa, 0x31,
	0xe7, 0x41, 0x9a, 0xf1, 0x5f, 0x29, 0x91, 0xc2, 0x8c, 0x70, 0xca, 0x82, 0xd1, 0x41, 0x90, 0xa7,
	0x42, 0x66, 0x14, 0x5f, 0xa3, 0x34, 0xe3, 0x92, 0xc3, 0x75, 0xb5, 0x86, 0x94, 0x0d, 0x31, 0xfe,
	0x6e, 0xab, 0xcf, 0xfb, 0x5c, 0x2f, 0x04, 0xea, 0xce, 0x68, 0xde, 0x41, 0x7a, 0x2b, 0xcd, 0x24,
	0xbd, 0x95, 0x76, 0x6e, 0x5b, 0x3f, 0x69, 0xc8, 0xa4, 0xe3, 0x5e, 0x53, 0x89, 0x7b, 0x58, 0x62,
	0xbb, 0xfe, 0xf9, 0xe2, 0x0a, 0x84, 0x88, 0xad, 0xe8, 0x85, 0x32, 0x09, 0xcb, 0x48, 0xce, 0x64,
	0xd4, 0xcd, 0x28, 0x1e, 0xd2, 0xcc, 0x1a, 0xf6, 0x17, 0x1b, 0x62, 0x8e, 0x7b, 0x51, 0x17, 0xc7,
	0x38, 0x21, 0x63, 0xf9, 0xde, 0x0b, 0x7c, 0x9e, 0x24, 0x94, 0x48, 0xc6, 0x13, 0xab, 0x3d, 0x5e,
	0xa0, 0xa5, 0xb7, 0x92, 0x66, 0x09, 0x8e, 0x03, 0x9a, 0x8c, 0xf8, 0x9d, 0xb1, 0x37, 0x03, 0xc2,
	0x33, 0x1a, 0x0c, 0x28, 0x8e, 0xe5, 0x20, 0x22, 0x03, 0x4a, 0x86, 0x96, 0xf2, 0xe9, 0x6c, 0x5b,
	0x84, 0xc4, 0x32, 0x17, 0x76, 0xf5, 0xfc, 0xff, 0x3d, 0x23, 0xce, 0x85, 0xa4, 0x59, 0xc0, 0x73,
	0x19, 0x33, 0x9a, 0x45, 0x3d, 0x2a, 0xa7, 0x2a, 0x9e, 0x7b, 0x05, 0x6e, 0x6c, 0xd7, 0xbf, 0x5e,
	0xfc, 0xdf, 0xf3, 0x54, 0x71, 0x84, 0xae, 0x8e, 0x11, 0x7b, 0xb1, 0xb6, 0x83, 0xe5, 0xb6, 0x94,
	0xa5, 0x54, 0xff, 0x58, 0xcb, 0xd1, 0x72, 0xcb, 0x30, 0xef, 0xd2, 0x2c, 0xa1, 0x92, 0x4e, 0xde,
	0x2e, 0xdf, 0x06, 0xce, 0x8e, 0x6f, 0xf4, 0x9f, 0x35, 0x1c, 0x16, 0x30, 0xfc, 0x96, 0x67, 0xd4,
	0xfc, 0x16, 0x6f, 0x07, 0xe1, 0x89, 0xc8, 0x63, 0x7b, 0xb1, 0xb6, 0x6f, 0x8a, 0x15, 0x47, 0x49,
	0x53, 0x5d, 0x23, 0x4a, 0x9a, 0xd6, 0xf8, 0x7e, 0xa9, 0xd1, 0x08, 0x1b, 0x7f, 0x56, 0x40, 0xf9,
	0x67, 0x7b, 0x2a, 0xe1, 0x19, 0x58, 0x35, 0x5b, 0xc6, 0x2f, 0xd5, 0x4b, 0xbb, 0x6b, 0xcd, 0x2d,
	0xa4, 0xb6, 0x9a, 0x3b, 0xa0, 0xe8, 0x52, 0xaf, 0xb5, 0x3f, 0xfb, 0xe3, 0x5f, 0xaf, 0xf4, 0xd7,
	0xc3, 0xce, 0xab, 0x7f, 0x1e, 0x76, 0x6a, 0x92, 0x0a, 0xd9, 0x63, 0x57, 0x57, 0xad, 0x06, 0xeb,
	0x27, 0x3c, 0xa3, 0x8d, 0xd0, 0x22, 0xe0, 0xb7, 0xa0, 0xec, 0x8e, 0xa5, 0xff, 0x5a, 0xe3, 0x3e,
	0x9a, 0xc6, 0x5d, 0xd8, 0xd5, 0xb6, 0xa7, 0x60, 0xe1, 0x58, 0x0d, 0x7f, 0x04, 0xb0, 0xc7, 0x04,
	0xe1, 0x23, 0x9a, 0xdd, 0x45, 0x63, 0xc6, 0x8a, 0x66, 0xec, 0xa0, 0xc9, 0xcc, 0x40, 0xc7, 0x4e,
	0xe7, 0x60, 0x61, 0xad, 0x37, 0x3b, 0x05, 0xbf, 0x07, 0x40, 0x88, 0x38, 0x22, 0x3c, 0xb9, 0x62,
	0x7d, 0xdf, 0x7b, 0x8e, 0xe3, 0x5a, 0x70, 0x29, 0xe2, 0x8e, 0x96, 0x85, 0x15, 0xe1, 0x6e, 0xe1,
	0x05, 0xa8, 0xce, 0x24, 0x82, 0xf0, 0xdf, 0x68, 0x4a, 0x63, 0x9a, 0xd2, 0x31, 0xaa, 0xb6, 0x11,
	0x59, 0xd0, 0x26, 0x99, 0x9a, 0x15, 0x30, 0x04, 0x5b, 0x53, 0x79, 0xe1, 0x0a, 0x5b, 0xd5, 0xc8,
	0xfa, 0x34, 0xf2, 0x9c, 0xe3, 0x5e, 0xdb, 0x0a, 0x2d, 0x10, 0xc6, 0x73, 0x73, 0xf0, 0x0c, 0xd4,
	0x9e, 0x42, 0xc5, 0x01, 0xdf, 0x6a, 0xe0, 0xf6, 0x4c, 0x8d, 0x63, 0x99, 0xc5, 0x55, 0xc9, 0xcc,
	0x0c, 0xec, 0x80, 0x8d, 0xc9, 0x74, 0x11, 0x7e, 0xb9, 0xbe, 0xa2, 0x41, 0x3a, 0x21, 0x10, 0x4e,
	0x19, 0x1a, 0x35, 0xcd, 0xbb, 0x3c, 0xd5, 0xba, 0x8e, 0x92, 0x85, 0xeb, 0x83, 0xa7, 0x81, 0x80,
	0x97, 0xa0, 0x36, 0x97, 0x1d, 0x7e, 0x45, 0x57, 0xf4, 0xc5, 0x0c, 0xc8, 0x44, 0x0d, 0xfa, 0xc9,
	0xc8, 0x8f, 0x9d, 0x3a, 0xac, 0xf2, 0x99, 0x19, 0xf8, 0x09, 0xa8, 0xe4, 0x82, 0x46, 0x03, 0x29,
	0xd3, 0xa6, 0x0f, 0xea, 0xa5, 0xdd, 0x72, 0x58, 0xce, 0x05, 0x3d, 0x55, 0x63, 0xd8, 0x01, 0x9e,
	0x3a, 0xdd, 0xfe, 0x9a, 0x7e, 0xc8, 0x3e, 0x9a, 0x38, 0xea, 0x6e, 0xcf, 0x3f, 0xff, 0xce, 0x53,
	0x4a, 0x4e, 0x5f, 0x85, 0xda, 0x0c, 0x3b, 0xe6, 0x08, 0x30, 0xe2, 0xaf, 0x6b, 0xcc, 0x97, 0xc8,
	0xe6, 0x53, 0x11, 0x84, 0xb5, 0xc2, 0x23, 0xe0, 0xa9, 0x80, 0xf2, 0x37, 0x34, 0xe2, 0x3d, 0xd2,
	0x69, 0x55, 0xa8, 0x06, 0xa5, 0x84, 0x2d, 0xb0, 0x82, 0x6f, 0x84, 0xff, 0x81, 0x6d, 0x96, 0x8a,
	0x9e, 0x22, 0x66, 0x65, 0x82, 0x3f, 0x80, 0x37, 0x3a, 0x77, 0xfc, 0x4d, 0xed, 0xde, 0x45, 0x26,
	0x85, 0x8a, 0xf8, 0x8d, 0x51, 0x75, 0xc0, 0x64, 0x90, 0x5f, 0xb5, 0x1d, 0xb0, 0x91, 0x54, 0xa8,
	0x03, 0x46, 0x0b, 0x4f, 0xc0, 0x5b, 0x1b, 0x48, 0x7e, 0x4d, 0x53, 0xf6, 0x90, 0x0b, 0xa8, 0x42,
	0x18, 0x7c, 0x23, 0x4e, 0x48, 0xb3, 0xf5, 0xf1, 0xfd, 0xa3, 0xe7, 0x81, 0xd7, 0xb9, 0xb8, 0x7f,
	0xf4, 0xd6, 0x60, 0xc5, 0x7d, 0x3e, 0x88, 0xf6, 0x26, 0xd8, 0x70, 0x83, 0x48, 0xde, 0xa5, 0xb4,
	0xf1, 0x21, 0xa8, 0xcd, 0x65, 0x41, 0xbb, 0xa5, 0x92, 0xea, 0xf7, 0xbf, 0xb7, 0x4b, 0xbf, 0x7c,
	0x55, 0xec, 0x33, 0x25, 0x1d, 0xf6, 0x6d, 0x4c, 0x76, 0x57, 0x75, 0x3e, 0x1e, 0xfe, 0x17, 0x00,
	0x00, 0xff, 0xff, 0x57, 0x55, 0x3e, 0x32, 0xe1, 0x08, 0x00, 0x00,
}

func (this *Upstream) Equal(that interface{}) bool {
	if that == nil {
		return this == nil
	}

	that1, ok := that.(*Upstream)
	if !ok {
		that2, ok := that.(Upstream)
		if ok {
			that1 = &that2
		} else {
			return false
		}
	}
	if that1 == nil {
		return this == nil
	} else if this == nil {
		return false
	}
	if !this.Status.Equal(&that1.Status) {
		return false
	}
	if !this.Metadata.Equal(&that1.Metadata) {
		return false
	}
	if !this.DiscoveryMetadata.Equal(that1.DiscoveryMetadata) {
		return false
	}
	if !this.SslConfig.Equal(that1.SslConfig) {
		return false
	}
	if !this.CircuitBreakers.Equal(that1.CircuitBreakers) {
		return false
	}
	if !this.LoadBalancerConfig.Equal(that1.LoadBalancerConfig) {
		return false
	}
	if !this.ConnectionConfig.Equal(that1.ConnectionConfig) {
		return false
	}
	if len(this.HealthChecks) != len(that1.HealthChecks) {
		return false
	}
	for i := range this.HealthChecks {
		if !this.HealthChecks[i].Equal(that1.HealthChecks[i]) {
			return false
		}
	}
	if !this.OutlierDetection.Equal(that1.OutlierDetection) {
		return false
	}
	if this.UseHttp2 != that1.UseHttp2 {
		return false
	}
	if that1.UpstreamType == nil {
		if this.UpstreamType != nil {
			return false
		}
	} else if this.UpstreamType == nil {
		return false
	} else if !this.UpstreamType.Equal(that1.UpstreamType) {
		return false
	}
	if !bytes.Equal(this.XXX_unrecognized, that1.XXX_unrecognized) {
		return false
	}
	return true
}
func (this *Upstream_Kube) Equal(that interface{}) bool {
	if that == nil {
		return this == nil
	}

	that1, ok := that.(*Upstream_Kube)
	if !ok {
		that2, ok := that.(Upstream_Kube)
		if ok {
			that1 = &that2
		} else {
			return false
		}
	}
	if that1 == nil {
		return this == nil
	} else if this == nil {
		return false
	}
	if !this.Kube.Equal(that1.Kube) {
		return false
	}
	return true
}
func (this *Upstream_Static) Equal(that interface{}) bool {
	if that == nil {
		return this == nil
	}

	that1, ok := that.(*Upstream_Static)
	if !ok {
		that2, ok := that.(Upstream_Static)
		if ok {
			that1 = &that2
		} else {
			return false
		}
	}
	if that1 == nil {
		return this == nil
	} else if this == nil {
		return false
	}
	if !this.Static.Equal(that1.Static) {
		return false
	}
	return true
}
func (this *Upstream_Pipe) Equal(that interface{}) bool {
	if that == nil {
		return this == nil
	}

	that1, ok := that.(*Upstream_Pipe)
	if !ok {
		that2, ok := that.(Upstream_Pipe)
		if ok {
			that1 = &that2
		} else {
			return false
		}
	}
	if that1 == nil {
		return this == nil
	} else if this == nil {
		return false
	}
	if !this.Pipe.Equal(that1.Pipe) {
		return false
	}
	return true
}
func (this *Upstream_Aws) Equal(that interface{}) bool {
	if that == nil {
		return this == nil
	}

	that1, ok := that.(*Upstream_Aws)
	if !ok {
		that2, ok := that.(Upstream_Aws)
		if ok {
			that1 = &that2
		} else {
			return false
		}
	}
	if that1 == nil {
		return this == nil
	} else if this == nil {
		return false
	}
	if !this.Aws.Equal(that1.Aws) {
		return false
	}
	return true
}
func (this *Upstream_Azure) Equal(that interface{}) bool {
	if that == nil {
		return this == nil
	}

	that1, ok := that.(*Upstream_Azure)
	if !ok {
		that2, ok := that.(Upstream_Azure)
		if ok {
			that1 = &that2
		} else {
			return false
		}
	}
	if that1 == nil {
		return this == nil
	} else if this == nil {
		return false
	}
	if !this.Azure.Equal(that1.Azure) {
		return false
	}
	return true
}
func (this *Upstream_Consul) Equal(that interface{}) bool {
	if that == nil {
		return this == nil
	}

	that1, ok := that.(*Upstream_Consul)
	if !ok {
		that2, ok := that.(Upstream_Consul)
		if ok {
			that1 = &that2
		} else {
			return false
		}
	}
	if that1 == nil {
		return this == nil
	} else if this == nil {
		return false
	}
	if !this.Consul.Equal(that1.Consul) {
		return false
	}
	return true
}
func (this *Upstream_AwsEc2) Equal(that interface{}) bool {
	if that == nil {
		return this == nil
	}

	that1, ok := that.(*Upstream_AwsEc2)
	if !ok {
		that2, ok := that.(Upstream_AwsEc2)
		if ok {
			that1 = &that2
		} else {
			return false
		}
	}
	if that1 == nil {
		return this == nil
	} else if this == nil {
		return false
	}
	if !this.AwsEc2.Equal(that1.AwsEc2) {
		return false
	}
	return true
}
func (this *DiscoveryMetadata) Equal(that interface{}) bool {
	if that == nil {
		return this == nil
	}

	that1, ok := that.(*DiscoveryMetadata)
	if !ok {
		that2, ok := that.(DiscoveryMetadata)
		if ok {
			that1 = &that2
		} else {
			return false
		}
	}
	if that1 == nil {
		return this == nil
	} else if this == nil {
		return false
	}
	if !bytes.Equal(this.XXX_unrecognized, that1.XXX_unrecognized) {
		return false
	}
	return true
}
