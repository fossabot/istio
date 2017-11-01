// Code generated by protoc-gen-gogo.
// source: mixer/v1/config/client/mixer_filter_config.proto
// DO NOT EDIT!

package istio_mixer_v1_config_client

import proto "github.com/gogo/protobuf/proto"
import fmt "fmt"
import math "math"
import _ "github.com/gogo/protobuf/gogoproto"
import istio_mixer_v1 "istio.io/api/mixer/v1"

import strconv "strconv"

import strings "strings"
import reflect "reflect"
import github_com_gogo_protobuf_sortkeys "github.com/gogo/protobuf/sortkeys"

import io "io"

// Reference imports to suppress errors if they are not otherwise used.
var _ = proto.Marshal
var _ = fmt.Errorf
var _ = math.Inf

type MixerFilterConfig_NetworkFailPolicy int32

const (
	// If network fails, request is passed to the backend.
	FAIL_OPEN MixerFilterConfig_NetworkFailPolicy = 0
	// If network fails, request is rejected.
	FAIL_CLOSE MixerFilterConfig_NetworkFailPolicy = 1
)

var MixerFilterConfig_NetworkFailPolicy_name = map[int32]string{
	0: "FAIL_OPEN",
	1: "FAIL_CLOSE",
}
var MixerFilterConfig_NetworkFailPolicy_value = map[string]int32{
	"FAIL_OPEN":  0,
	"FAIL_CLOSE": 1,
}

func (MixerFilterConfig_NetworkFailPolicy) EnumDescriptor() ([]byte, []int) {
	return fileDescriptorMixerFilterConfig, []int{1, 0}
}

// Defines the per-service mixerclient control configuration.
type MixerControlConfig struct {
	// If true, call Mixer Check.
	EnableMixerCheck bool `protobuf:"varint,1,opt,name=enable_mixer_check,json=enableMixerCheck,proto3" json:"enable_mixer_check,omitempty"`
	// If true, call Mixer Report.
	EnableMixer_Report bool `protobuf:"varint,2,opt,name=enable_mixer_Report,json=enableMixerReport,proto3" json:"enable_mixer_Report,omitempty"`
	// Send these attributes to Mixer in both Check and Report. This
	// typically includes the "destination.service" attribute.
	MixerAttributes *istio_mixer_v1.Attributes `protobuf:"bytes,3,opt,name=mixer_attributes,json=mixerAttributes" json:"mixer_attributes,omitempty"`
	// HTTP API specifications to generate API attributes.
	HttpApiSpec []*HTTPAPISpec `protobuf:"bytes,4,rep,name=http_api_spec,json=httpApiSpec" json:"http_api_spec,omitempty"`
	// Quota specifications to generate quota requirements.
	QuotaSpec []*QuotaSpec `protobuf:"bytes,5,rep,name=quota_spec,json=quotaSpec" json:"quota_spec,omitempty"`
}

func (m *MixerControlConfig) Reset()      { *m = MixerControlConfig{} }
func (*MixerControlConfig) ProtoMessage() {}
func (*MixerControlConfig) Descriptor() ([]byte, []int) {
	return fileDescriptorMixerFilterConfig, []int{0}
}

// Defines the Mixerclient's envoy HTTP filter and TCP filter
// configuration.
type MixerFilterConfig struct {
	// The flag to disable check cache.
	DisableCheckCache bool `protobuf:"varint,1,opt,name=disable_check_cache,json=disableCheckCache,proto3" json:"disable_check_cache,omitempty"`
	// The flag to disable quota cache.
	DisableQuotaCache bool `protobuf:"varint,2,opt,name=disable_quota_cache,json=disableQuotaCache,proto3" json:"disable_quota_cache,omitempty"`
	// The flag to disable report batch.
	DisableReportBatch bool `protobuf:"varint,3,opt,name=disable_report_batch,json=disableReportBatch,proto3" json:"disable_report_batch,omitempty"`
	// Specifies the policy when failed to connect to Mixer server.
	NetworkFailPolicy MixerFilterConfig_NetworkFailPolicy `protobuf:"varint,4,opt,name=network_fail_policy,json=networkFailPolicy,proto3,enum=istio.mixer.v1.config.client.MixerFilterConfig_NetworkFailPolicy" json:"network_fail_policy,omitempty"`
	// Map of control configuration indexed by destination.service. This
	// is used to support per-service configuration for cases where a
	// mixerclient serves multiple services.
	ControlConfigs map[string]*MixerControlConfig `protobuf:"bytes,5,rep,name=control_configs,json=controlConfigs" json:"control_configs,omitempty" protobuf_key:"bytes,1,opt,name=key,proto3" protobuf_val:"bytes,2,opt,name=value"`
	// Default destination service name if none was specified in the
	// client request.
	DefaultDestinationService string `protobuf:"bytes,6,opt,name=default_destination_service,json=defaultDestinationService,proto3" json:"default_destination_service,omitempty"`
	// Default attributes to send to Mixer in both Check and
	// Report. This typically includes "destination.ip" and
	// "destination.uid" attributes.
	MixerAttributes *istio_mixer_v1.Attributes `protobuf:"bytes,7,opt,name=mixer_attributes,json=mixerAttributes" json:"mixer_attributes,omitempty"`
	// Default attributes to forward to upstream. This typically
	// includes the "source.ip" and "source.uid" attributes.
	ForwardAttributes *istio_mixer_v1.Attributes `protobuf:"bytes,8,opt,name=forward_attributes,json=forwardAttributes" json:"forward_attributes,omitempty"`
	// If set to true, disables mixer check calls for TCP connections
	DisableTCPCheckCalls bool `protobuf:"varint,9,opt,name=DisableTCPCheckCalls,proto3" json:"DisableTCPCheckCalls,omitempty"`
}

func (m *MixerFilterConfig) Reset()      { *m = MixerFilterConfig{} }
func (*MixerFilterConfig) ProtoMessage() {}
func (*MixerFilterConfig) Descriptor() ([]byte, []int) {
	return fileDescriptorMixerFilterConfig, []int{1}
}

func init() {
	proto.RegisterType((*MixerControlConfig)(nil), "istio.mixer.v1.config.client.MixerControlConfig")
	proto.RegisterType((*MixerFilterConfig)(nil), "istio.mixer.v1.config.client.MixerFilterConfig")
	proto.RegisterEnum("istio.mixer.v1.config.client.MixerFilterConfig_NetworkFailPolicy", MixerFilterConfig_NetworkFailPolicy_name, MixerFilterConfig_NetworkFailPolicy_value)
}
func (x MixerFilterConfig_NetworkFailPolicy) String() string {
	s, ok := MixerFilterConfig_NetworkFailPolicy_name[int32(x)]
	if ok {
		return s
	}
	return strconv.Itoa(int(x))
}
func (m *MixerControlConfig) Marshal() (dAtA []byte, err error) {
	size := m.Size()
	dAtA = make([]byte, size)
	n, err := m.MarshalTo(dAtA)
	if err != nil {
		return nil, err
	}
	return dAtA[:n], nil
}

func (m *MixerControlConfig) MarshalTo(dAtA []byte) (int, error) {
	var i int
	_ = i
	var l int
	_ = l
	if m.EnableMixerCheck {
		dAtA[i] = 0x8
		i++
		if m.EnableMixerCheck {
			dAtA[i] = 1
		} else {
			dAtA[i] = 0
		}
		i++
	}
	if m.EnableMixer_Report {
		dAtA[i] = 0x10
		i++
		if m.EnableMixer_Report {
			dAtA[i] = 1
		} else {
			dAtA[i] = 0
		}
		i++
	}
	if m.MixerAttributes != nil {
		dAtA[i] = 0x1a
		i++
		i = encodeVarintMixerFilterConfig(dAtA, i, uint64(m.MixerAttributes.Size()))
		n1, err := m.MixerAttributes.MarshalTo(dAtA[i:])
		if err != nil {
			return 0, err
		}
		i += n1
	}
	if len(m.HttpApiSpec) > 0 {
		for _, msg := range m.HttpApiSpec {
			dAtA[i] = 0x22
			i++
			i = encodeVarintMixerFilterConfig(dAtA, i, uint64(msg.Size()))
			n, err := msg.MarshalTo(dAtA[i:])
			if err != nil {
				return 0, err
			}
			i += n
		}
	}
	if len(m.QuotaSpec) > 0 {
		for _, msg := range m.QuotaSpec {
			dAtA[i] = 0x2a
			i++
			i = encodeVarintMixerFilterConfig(dAtA, i, uint64(msg.Size()))
			n, err := msg.MarshalTo(dAtA[i:])
			if err != nil {
				return 0, err
			}
			i += n
		}
	}
	return i, nil
}

func (m *MixerFilterConfig) Marshal() (dAtA []byte, err error) {
	size := m.Size()
	dAtA = make([]byte, size)
	n, err := m.MarshalTo(dAtA)
	if err != nil {
		return nil, err
	}
	return dAtA[:n], nil
}

func (m *MixerFilterConfig) MarshalTo(dAtA []byte) (int, error) {
	var i int
	_ = i
	var l int
	_ = l
	if m.DisableCheckCache {
		dAtA[i] = 0x8
		i++
		if m.DisableCheckCache {
			dAtA[i] = 1
		} else {
			dAtA[i] = 0
		}
		i++
	}
	if m.DisableQuotaCache {
		dAtA[i] = 0x10
		i++
		if m.DisableQuotaCache {
			dAtA[i] = 1
		} else {
			dAtA[i] = 0
		}
		i++
	}
	if m.DisableReportBatch {
		dAtA[i] = 0x18
		i++
		if m.DisableReportBatch {
			dAtA[i] = 1
		} else {
			dAtA[i] = 0
		}
		i++
	}
	if m.NetworkFailPolicy != 0 {
		dAtA[i] = 0x20
		i++
		i = encodeVarintMixerFilterConfig(dAtA, i, uint64(m.NetworkFailPolicy))
	}
	if len(m.ControlConfigs) > 0 {
		for k, _ := range m.ControlConfigs {
			dAtA[i] = 0x2a
			i++
			v := m.ControlConfigs[k]
			msgSize := 0
			if v != nil {
				msgSize = v.Size()
				msgSize += 1 + sovMixerFilterConfig(uint64(msgSize))
			}
			mapSize := 1 + len(k) + sovMixerFilterConfig(uint64(len(k))) + msgSize
			i = encodeVarintMixerFilterConfig(dAtA, i, uint64(mapSize))
			dAtA[i] = 0xa
			i++
			i = encodeVarintMixerFilterConfig(dAtA, i, uint64(len(k)))
			i += copy(dAtA[i:], k)
			if v != nil {
				dAtA[i] = 0x12
				i++
				i = encodeVarintMixerFilterConfig(dAtA, i, uint64(v.Size()))
				n2, err := v.MarshalTo(dAtA[i:])
				if err != nil {
					return 0, err
				}
				i += n2
			}
		}
	}
	if len(m.DefaultDestinationService) > 0 {
		dAtA[i] = 0x32
		i++
		i = encodeVarintMixerFilterConfig(dAtA, i, uint64(len(m.DefaultDestinationService)))
		i += copy(dAtA[i:], m.DefaultDestinationService)
	}
	if m.MixerAttributes != nil {
		dAtA[i] = 0x3a
		i++
		i = encodeVarintMixerFilterConfig(dAtA, i, uint64(m.MixerAttributes.Size()))
		n3, err := m.MixerAttributes.MarshalTo(dAtA[i:])
		if err != nil {
			return 0, err
		}
		i += n3
	}
	if m.ForwardAttributes != nil {
		dAtA[i] = 0x42
		i++
		i = encodeVarintMixerFilterConfig(dAtA, i, uint64(m.ForwardAttributes.Size()))
		n4, err := m.ForwardAttributes.MarshalTo(dAtA[i:])
		if err != nil {
			return 0, err
		}
		i += n4
	}
	if m.DisableTCPCheckCalls {
		dAtA[i] = 0x48
		i++
		if m.DisableTCPCheckCalls {
			dAtA[i] = 1
		} else {
			dAtA[i] = 0
		}
		i++
	}
	return i, nil
}

func encodeFixed64MixerFilterConfig(dAtA []byte, offset int, v uint64) int {
	dAtA[offset] = uint8(v)
	dAtA[offset+1] = uint8(v >> 8)
	dAtA[offset+2] = uint8(v >> 16)
	dAtA[offset+3] = uint8(v >> 24)
	dAtA[offset+4] = uint8(v >> 32)
	dAtA[offset+5] = uint8(v >> 40)
	dAtA[offset+6] = uint8(v >> 48)
	dAtA[offset+7] = uint8(v >> 56)
	return offset + 8
}
func encodeFixed32MixerFilterConfig(dAtA []byte, offset int, v uint32) int {
	dAtA[offset] = uint8(v)
	dAtA[offset+1] = uint8(v >> 8)
	dAtA[offset+2] = uint8(v >> 16)
	dAtA[offset+3] = uint8(v >> 24)
	return offset + 4
}
func encodeVarintMixerFilterConfig(dAtA []byte, offset int, v uint64) int {
	for v >= 1<<7 {
		dAtA[offset] = uint8(v&0x7f | 0x80)
		v >>= 7
		offset++
	}
	dAtA[offset] = uint8(v)
	return offset + 1
}
func (m *MixerControlConfig) Size() (n int) {
	var l int
	_ = l
	if m.EnableMixerCheck {
		n += 2
	}
	if m.EnableMixer_Report {
		n += 2
	}
	if m.MixerAttributes != nil {
		l = m.MixerAttributes.Size()
		n += 1 + l + sovMixerFilterConfig(uint64(l))
	}
	if len(m.HttpApiSpec) > 0 {
		for _, e := range m.HttpApiSpec {
			l = e.Size()
			n += 1 + l + sovMixerFilterConfig(uint64(l))
		}
	}
	if len(m.QuotaSpec) > 0 {
		for _, e := range m.QuotaSpec {
			l = e.Size()
			n += 1 + l + sovMixerFilterConfig(uint64(l))
		}
	}
	return n
}

func (m *MixerFilterConfig) Size() (n int) {
	var l int
	_ = l
	if m.DisableCheckCache {
		n += 2
	}
	if m.DisableQuotaCache {
		n += 2
	}
	if m.DisableReportBatch {
		n += 2
	}
	if m.NetworkFailPolicy != 0 {
		n += 1 + sovMixerFilterConfig(uint64(m.NetworkFailPolicy))
	}
	if len(m.ControlConfigs) > 0 {
		for k, v := range m.ControlConfigs {
			_ = k
			_ = v
			l = 0
			if v != nil {
				l = v.Size()
				l += 1 + sovMixerFilterConfig(uint64(l))
			}
			mapEntrySize := 1 + len(k) + sovMixerFilterConfig(uint64(len(k))) + l
			n += mapEntrySize + 1 + sovMixerFilterConfig(uint64(mapEntrySize))
		}
	}
	l = len(m.DefaultDestinationService)
	if l > 0 {
		n += 1 + l + sovMixerFilterConfig(uint64(l))
	}
	if m.MixerAttributes != nil {
		l = m.MixerAttributes.Size()
		n += 1 + l + sovMixerFilterConfig(uint64(l))
	}
	if m.ForwardAttributes != nil {
		l = m.ForwardAttributes.Size()
		n += 1 + l + sovMixerFilterConfig(uint64(l))
	}
	if m.DisableTCPCheckCalls {
		n += 2
	}
	return n
}

func sovMixerFilterConfig(x uint64) (n int) {
	for {
		n++
		x >>= 7
		if x == 0 {
			break
		}
	}
	return n
}
func sozMixerFilterConfig(x uint64) (n int) {
	return sovMixerFilterConfig(uint64((x << 1) ^ uint64((int64(x) >> 63))))
}
func (this *MixerControlConfig) String() string {
	if this == nil {
		return "nil"
	}
	s := strings.Join([]string{`&MixerControlConfig{`,
		`EnableMixerCheck:` + fmt.Sprintf("%v", this.EnableMixerCheck) + `,`,
		`EnableMixer_Report:` + fmt.Sprintf("%v", this.EnableMixer_Report) + `,`,
		`MixerAttributes:` + strings.Replace(fmt.Sprintf("%v", this.MixerAttributes), "Attributes", "istio_mixer_v1.Attributes", 1) + `,`,
		`HttpApiSpec:` + strings.Replace(fmt.Sprintf("%v", this.HttpApiSpec), "HTTPAPISpec", "HTTPAPISpec", 1) + `,`,
		`QuotaSpec:` + strings.Replace(fmt.Sprintf("%v", this.QuotaSpec), "QuotaSpec", "QuotaSpec", 1) + `,`,
		`}`,
	}, "")
	return s
}
func (this *MixerFilterConfig) String() string {
	if this == nil {
		return "nil"
	}
	keysForControlConfigs := make([]string, 0, len(this.ControlConfigs))
	for k, _ := range this.ControlConfigs {
		keysForControlConfigs = append(keysForControlConfigs, k)
	}
	github_com_gogo_protobuf_sortkeys.Strings(keysForControlConfigs)
	mapStringForControlConfigs := "map[string]*MixerControlConfig{"
	for _, k := range keysForControlConfigs {
		mapStringForControlConfigs += fmt.Sprintf("%v: %v,", k, this.ControlConfigs[k])
	}
	mapStringForControlConfigs += "}"
	s := strings.Join([]string{`&MixerFilterConfig{`,
		`DisableCheckCache:` + fmt.Sprintf("%v", this.DisableCheckCache) + `,`,
		`DisableQuotaCache:` + fmt.Sprintf("%v", this.DisableQuotaCache) + `,`,
		`DisableReportBatch:` + fmt.Sprintf("%v", this.DisableReportBatch) + `,`,
		`NetworkFailPolicy:` + fmt.Sprintf("%v", this.NetworkFailPolicy) + `,`,
		`ControlConfigs:` + mapStringForControlConfigs + `,`,
		`DefaultDestinationService:` + fmt.Sprintf("%v", this.DefaultDestinationService) + `,`,
		`MixerAttributes:` + strings.Replace(fmt.Sprintf("%v", this.MixerAttributes), "Attributes", "istio_mixer_v1.Attributes", 1) + `,`,
		`ForwardAttributes:` + strings.Replace(fmt.Sprintf("%v", this.ForwardAttributes), "Attributes", "istio_mixer_v1.Attributes", 1) + `,`,
		`DisableTCPCheckCalls:` + fmt.Sprintf("%v", this.DisableTCPCheckCalls) + `,`,
		`}`,
	}, "")
	return s
}
func valueToStringMixerFilterConfig(v interface{}) string {
	rv := reflect.ValueOf(v)
	if rv.IsNil() {
		return "nil"
	}
	pv := reflect.Indirect(rv).Interface()
	return fmt.Sprintf("*%v", pv)
}
func (m *MixerControlConfig) Unmarshal(dAtA []byte) error {
	l := len(dAtA)
	iNdEx := 0
	for iNdEx < l {
		preIndex := iNdEx
		var wire uint64
		for shift := uint(0); ; shift += 7 {
			if shift >= 64 {
				return ErrIntOverflowMixerFilterConfig
			}
			if iNdEx >= l {
				return io.ErrUnexpectedEOF
			}
			b := dAtA[iNdEx]
			iNdEx++
			wire |= (uint64(b) & 0x7F) << shift
			if b < 0x80 {
				break
			}
		}
		fieldNum := int32(wire >> 3)
		wireType := int(wire & 0x7)
		if wireType == 4 {
			return fmt.Errorf("proto: MixerControlConfig: wiretype end group for non-group")
		}
		if fieldNum <= 0 {
			return fmt.Errorf("proto: MixerControlConfig: illegal tag %d (wire type %d)", fieldNum, wire)
		}
		switch fieldNum {
		case 1:
			if wireType != 0 {
				return fmt.Errorf("proto: wrong wireType = %d for field EnableMixerCheck", wireType)
			}
			var v int
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowMixerFilterConfig
				}
				if iNdEx >= l {
					return io.ErrUnexpectedEOF
				}
				b := dAtA[iNdEx]
				iNdEx++
				v |= (int(b) & 0x7F) << shift
				if b < 0x80 {
					break
				}
			}
			m.EnableMixerCheck = bool(v != 0)
		case 2:
			if wireType != 0 {
				return fmt.Errorf("proto: wrong wireType = %d for field EnableMixer_Report", wireType)
			}
			var v int
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowMixerFilterConfig
				}
				if iNdEx >= l {
					return io.ErrUnexpectedEOF
				}
				b := dAtA[iNdEx]
				iNdEx++
				v |= (int(b) & 0x7F) << shift
				if b < 0x80 {
					break
				}
			}
			m.EnableMixer_Report = bool(v != 0)
		case 3:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field MixerAttributes", wireType)
			}
			var msglen int
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowMixerFilterConfig
				}
				if iNdEx >= l {
					return io.ErrUnexpectedEOF
				}
				b := dAtA[iNdEx]
				iNdEx++
				msglen |= (int(b) & 0x7F) << shift
				if b < 0x80 {
					break
				}
			}
			if msglen < 0 {
				return ErrInvalidLengthMixerFilterConfig
			}
			postIndex := iNdEx + msglen
			if postIndex > l {
				return io.ErrUnexpectedEOF
			}
			if m.MixerAttributes == nil {
				m.MixerAttributes = &istio_mixer_v1.Attributes{}
			}
			if err := m.MixerAttributes.Unmarshal(dAtA[iNdEx:postIndex]); err != nil {
				return err
			}
			iNdEx = postIndex
		case 4:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field HttpApiSpec", wireType)
			}
			var msglen int
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowMixerFilterConfig
				}
				if iNdEx >= l {
					return io.ErrUnexpectedEOF
				}
				b := dAtA[iNdEx]
				iNdEx++
				msglen |= (int(b) & 0x7F) << shift
				if b < 0x80 {
					break
				}
			}
			if msglen < 0 {
				return ErrInvalidLengthMixerFilterConfig
			}
			postIndex := iNdEx + msglen
			if postIndex > l {
				return io.ErrUnexpectedEOF
			}
			m.HttpApiSpec = append(m.HttpApiSpec, &HTTPAPISpec{})
			if err := m.HttpApiSpec[len(m.HttpApiSpec)-1].Unmarshal(dAtA[iNdEx:postIndex]); err != nil {
				return err
			}
			iNdEx = postIndex
		case 5:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field QuotaSpec", wireType)
			}
			var msglen int
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowMixerFilterConfig
				}
				if iNdEx >= l {
					return io.ErrUnexpectedEOF
				}
				b := dAtA[iNdEx]
				iNdEx++
				msglen |= (int(b) & 0x7F) << shift
				if b < 0x80 {
					break
				}
			}
			if msglen < 0 {
				return ErrInvalidLengthMixerFilterConfig
			}
			postIndex := iNdEx + msglen
			if postIndex > l {
				return io.ErrUnexpectedEOF
			}
			m.QuotaSpec = append(m.QuotaSpec, &QuotaSpec{})
			if err := m.QuotaSpec[len(m.QuotaSpec)-1].Unmarshal(dAtA[iNdEx:postIndex]); err != nil {
				return err
			}
			iNdEx = postIndex
		default:
			iNdEx = preIndex
			skippy, err := skipMixerFilterConfig(dAtA[iNdEx:])
			if err != nil {
				return err
			}
			if skippy < 0 {
				return ErrInvalidLengthMixerFilterConfig
			}
			if (iNdEx + skippy) > l {
				return io.ErrUnexpectedEOF
			}
			iNdEx += skippy
		}
	}

	if iNdEx > l {
		return io.ErrUnexpectedEOF
	}
	return nil
}
func (m *MixerFilterConfig) Unmarshal(dAtA []byte) error {
	l := len(dAtA)
	iNdEx := 0
	for iNdEx < l {
		preIndex := iNdEx
		var wire uint64
		for shift := uint(0); ; shift += 7 {
			if shift >= 64 {
				return ErrIntOverflowMixerFilterConfig
			}
			if iNdEx >= l {
				return io.ErrUnexpectedEOF
			}
			b := dAtA[iNdEx]
			iNdEx++
			wire |= (uint64(b) & 0x7F) << shift
			if b < 0x80 {
				break
			}
		}
		fieldNum := int32(wire >> 3)
		wireType := int(wire & 0x7)
		if wireType == 4 {
			return fmt.Errorf("proto: MixerFilterConfig: wiretype end group for non-group")
		}
		if fieldNum <= 0 {
			return fmt.Errorf("proto: MixerFilterConfig: illegal tag %d (wire type %d)", fieldNum, wire)
		}
		switch fieldNum {
		case 1:
			if wireType != 0 {
				return fmt.Errorf("proto: wrong wireType = %d for field DisableCheckCache", wireType)
			}
			var v int
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowMixerFilterConfig
				}
				if iNdEx >= l {
					return io.ErrUnexpectedEOF
				}
				b := dAtA[iNdEx]
				iNdEx++
				v |= (int(b) & 0x7F) << shift
				if b < 0x80 {
					break
				}
			}
			m.DisableCheckCache = bool(v != 0)
		case 2:
			if wireType != 0 {
				return fmt.Errorf("proto: wrong wireType = %d for field DisableQuotaCache", wireType)
			}
			var v int
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowMixerFilterConfig
				}
				if iNdEx >= l {
					return io.ErrUnexpectedEOF
				}
				b := dAtA[iNdEx]
				iNdEx++
				v |= (int(b) & 0x7F) << shift
				if b < 0x80 {
					break
				}
			}
			m.DisableQuotaCache = bool(v != 0)
		case 3:
			if wireType != 0 {
				return fmt.Errorf("proto: wrong wireType = %d for field DisableReportBatch", wireType)
			}
			var v int
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowMixerFilterConfig
				}
				if iNdEx >= l {
					return io.ErrUnexpectedEOF
				}
				b := dAtA[iNdEx]
				iNdEx++
				v |= (int(b) & 0x7F) << shift
				if b < 0x80 {
					break
				}
			}
			m.DisableReportBatch = bool(v != 0)
		case 4:
			if wireType != 0 {
				return fmt.Errorf("proto: wrong wireType = %d for field NetworkFailPolicy", wireType)
			}
			m.NetworkFailPolicy = 0
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowMixerFilterConfig
				}
				if iNdEx >= l {
					return io.ErrUnexpectedEOF
				}
				b := dAtA[iNdEx]
				iNdEx++
				m.NetworkFailPolicy |= (MixerFilterConfig_NetworkFailPolicy(b) & 0x7F) << shift
				if b < 0x80 {
					break
				}
			}
		case 5:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field ControlConfigs", wireType)
			}
			var msglen int
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowMixerFilterConfig
				}
				if iNdEx >= l {
					return io.ErrUnexpectedEOF
				}
				b := dAtA[iNdEx]
				iNdEx++
				msglen |= (int(b) & 0x7F) << shift
				if b < 0x80 {
					break
				}
			}
			if msglen < 0 {
				return ErrInvalidLengthMixerFilterConfig
			}
			postIndex := iNdEx + msglen
			if postIndex > l {
				return io.ErrUnexpectedEOF
			}
			var keykey uint64
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowMixerFilterConfig
				}
				if iNdEx >= l {
					return io.ErrUnexpectedEOF
				}
				b := dAtA[iNdEx]
				iNdEx++
				keykey |= (uint64(b) & 0x7F) << shift
				if b < 0x80 {
					break
				}
			}
			var stringLenmapkey uint64
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowMixerFilterConfig
				}
				if iNdEx >= l {
					return io.ErrUnexpectedEOF
				}
				b := dAtA[iNdEx]
				iNdEx++
				stringLenmapkey |= (uint64(b) & 0x7F) << shift
				if b < 0x80 {
					break
				}
			}
			intStringLenmapkey := int(stringLenmapkey)
			if intStringLenmapkey < 0 {
				return ErrInvalidLengthMixerFilterConfig
			}
			postStringIndexmapkey := iNdEx + intStringLenmapkey
			if postStringIndexmapkey > l {
				return io.ErrUnexpectedEOF
			}
			mapkey := string(dAtA[iNdEx:postStringIndexmapkey])
			iNdEx = postStringIndexmapkey
			if m.ControlConfigs == nil {
				m.ControlConfigs = make(map[string]*MixerControlConfig)
			}
			if iNdEx < postIndex {
				var valuekey uint64
				for shift := uint(0); ; shift += 7 {
					if shift >= 64 {
						return ErrIntOverflowMixerFilterConfig
					}
					if iNdEx >= l {
						return io.ErrUnexpectedEOF
					}
					b := dAtA[iNdEx]
					iNdEx++
					valuekey |= (uint64(b) & 0x7F) << shift
					if b < 0x80 {
						break
					}
				}
				var mapmsglen int
				for shift := uint(0); ; shift += 7 {
					if shift >= 64 {
						return ErrIntOverflowMixerFilterConfig
					}
					if iNdEx >= l {
						return io.ErrUnexpectedEOF
					}
					b := dAtA[iNdEx]
					iNdEx++
					mapmsglen |= (int(b) & 0x7F) << shift
					if b < 0x80 {
						break
					}
				}
				if mapmsglen < 0 {
					return ErrInvalidLengthMixerFilterConfig
				}
				postmsgIndex := iNdEx + mapmsglen
				if mapmsglen < 0 {
					return ErrInvalidLengthMixerFilterConfig
				}
				if postmsgIndex > l {
					return io.ErrUnexpectedEOF
				}
				mapvalue := &MixerControlConfig{}
				if err := mapvalue.Unmarshal(dAtA[iNdEx:postmsgIndex]); err != nil {
					return err
				}
				iNdEx = postmsgIndex
				m.ControlConfigs[mapkey] = mapvalue
			} else {
				var mapvalue *MixerControlConfig
				m.ControlConfigs[mapkey] = mapvalue
			}
			iNdEx = postIndex
		case 6:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field DefaultDestinationService", wireType)
			}
			var stringLen uint64
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowMixerFilterConfig
				}
				if iNdEx >= l {
					return io.ErrUnexpectedEOF
				}
				b := dAtA[iNdEx]
				iNdEx++
				stringLen |= (uint64(b) & 0x7F) << shift
				if b < 0x80 {
					break
				}
			}
			intStringLen := int(stringLen)
			if intStringLen < 0 {
				return ErrInvalidLengthMixerFilterConfig
			}
			postIndex := iNdEx + intStringLen
			if postIndex > l {
				return io.ErrUnexpectedEOF
			}
			m.DefaultDestinationService = string(dAtA[iNdEx:postIndex])
			iNdEx = postIndex
		case 7:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field MixerAttributes", wireType)
			}
			var msglen int
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowMixerFilterConfig
				}
				if iNdEx >= l {
					return io.ErrUnexpectedEOF
				}
				b := dAtA[iNdEx]
				iNdEx++
				msglen |= (int(b) & 0x7F) << shift
				if b < 0x80 {
					break
				}
			}
			if msglen < 0 {
				return ErrInvalidLengthMixerFilterConfig
			}
			postIndex := iNdEx + msglen
			if postIndex > l {
				return io.ErrUnexpectedEOF
			}
			if m.MixerAttributes == nil {
				m.MixerAttributes = &istio_mixer_v1.Attributes{}
			}
			if err := m.MixerAttributes.Unmarshal(dAtA[iNdEx:postIndex]); err != nil {
				return err
			}
			iNdEx = postIndex
		case 8:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field ForwardAttributes", wireType)
			}
			var msglen int
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowMixerFilterConfig
				}
				if iNdEx >= l {
					return io.ErrUnexpectedEOF
				}
				b := dAtA[iNdEx]
				iNdEx++
				msglen |= (int(b) & 0x7F) << shift
				if b < 0x80 {
					break
				}
			}
			if msglen < 0 {
				return ErrInvalidLengthMixerFilterConfig
			}
			postIndex := iNdEx + msglen
			if postIndex > l {
				return io.ErrUnexpectedEOF
			}
			if m.ForwardAttributes == nil {
				m.ForwardAttributes = &istio_mixer_v1.Attributes{}
			}
			if err := m.ForwardAttributes.Unmarshal(dAtA[iNdEx:postIndex]); err != nil {
				return err
			}
			iNdEx = postIndex
		case 9:
			if wireType != 0 {
				return fmt.Errorf("proto: wrong wireType = %d for field DisableTCPCheckCalls", wireType)
			}
			var v int
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowMixerFilterConfig
				}
				if iNdEx >= l {
					return io.ErrUnexpectedEOF
				}
				b := dAtA[iNdEx]
				iNdEx++
				v |= (int(b) & 0x7F) << shift
				if b < 0x80 {
					break
				}
			}
			m.DisableTCPCheckCalls = bool(v != 0)
		default:
			iNdEx = preIndex
			skippy, err := skipMixerFilterConfig(dAtA[iNdEx:])
			if err != nil {
				return err
			}
			if skippy < 0 {
				return ErrInvalidLengthMixerFilterConfig
			}
			if (iNdEx + skippy) > l {
				return io.ErrUnexpectedEOF
			}
			iNdEx += skippy
		}
	}

	if iNdEx > l {
		return io.ErrUnexpectedEOF
	}
	return nil
}
func skipMixerFilterConfig(dAtA []byte) (n int, err error) {
	l := len(dAtA)
	iNdEx := 0
	for iNdEx < l {
		var wire uint64
		for shift := uint(0); ; shift += 7 {
			if shift >= 64 {
				return 0, ErrIntOverflowMixerFilterConfig
			}
			if iNdEx >= l {
				return 0, io.ErrUnexpectedEOF
			}
			b := dAtA[iNdEx]
			iNdEx++
			wire |= (uint64(b) & 0x7F) << shift
			if b < 0x80 {
				break
			}
		}
		wireType := int(wire & 0x7)
		switch wireType {
		case 0:
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return 0, ErrIntOverflowMixerFilterConfig
				}
				if iNdEx >= l {
					return 0, io.ErrUnexpectedEOF
				}
				iNdEx++
				if dAtA[iNdEx-1] < 0x80 {
					break
				}
			}
			return iNdEx, nil
		case 1:
			iNdEx += 8
			return iNdEx, nil
		case 2:
			var length int
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return 0, ErrIntOverflowMixerFilterConfig
				}
				if iNdEx >= l {
					return 0, io.ErrUnexpectedEOF
				}
				b := dAtA[iNdEx]
				iNdEx++
				length |= (int(b) & 0x7F) << shift
				if b < 0x80 {
					break
				}
			}
			iNdEx += length
			if length < 0 {
				return 0, ErrInvalidLengthMixerFilterConfig
			}
			return iNdEx, nil
		case 3:
			for {
				var innerWire uint64
				var start int = iNdEx
				for shift := uint(0); ; shift += 7 {
					if shift >= 64 {
						return 0, ErrIntOverflowMixerFilterConfig
					}
					if iNdEx >= l {
						return 0, io.ErrUnexpectedEOF
					}
					b := dAtA[iNdEx]
					iNdEx++
					innerWire |= (uint64(b) & 0x7F) << shift
					if b < 0x80 {
						break
					}
				}
				innerWireType := int(innerWire & 0x7)
				if innerWireType == 4 {
					break
				}
				next, err := skipMixerFilterConfig(dAtA[start:])
				if err != nil {
					return 0, err
				}
				iNdEx = start + next
			}
			return iNdEx, nil
		case 4:
			return iNdEx, nil
		case 5:
			iNdEx += 4
			return iNdEx, nil
		default:
			return 0, fmt.Errorf("proto: illegal wireType %d", wireType)
		}
	}
	panic("unreachable")
}

var (
	ErrInvalidLengthMixerFilterConfig = fmt.Errorf("proto: negative length found during unmarshaling")
	ErrIntOverflowMixerFilterConfig   = fmt.Errorf("proto: integer overflow")
)

func init() {
	proto.RegisterFile("mixer/v1/config/client/mixer_filter_config.proto", fileDescriptorMixerFilterConfig)
}

var fileDescriptorMixerFilterConfig = []byte{
	// 655 bytes of a gzipped FileDescriptorProto
	0x1f, 0x8b, 0x08, 0x00, 0x00, 0x00, 0x00, 0x00, 0x02, 0xff, 0x94, 0x94, 0xcb, 0x6e, 0xda, 0x40,
	0x14, 0x86, 0x6d, 0x72, 0x69, 0x18, 0x94, 0x04, 0x86, 0x2c, 0x1c, 0x5a, 0x59, 0x08, 0xa9, 0x2a,
	0x95, 0x2a, 0x43, 0xe8, 0xa6, 0xea, 0xa2, 0x12, 0x21, 0xa0, 0x46, 0xca, 0x85, 0x3a, 0xd9, 0x5b,
	0xc3, 0x30, 0x84, 0x51, 0xa6, 0x1e, 0x67, 0x3c, 0x90, 0x66, 0xd7, 0x47, 0xe8, 0x63, 0xf4, 0x05,
	0xfa, 0x0e, 0x59, 0x66, 0xd9, 0x65, 0x71, 0x37, 0x5d, 0xe6, 0x11, 0x2a, 0xcf, 0x4c, 0x12, 0xc8,
	0xad, 0xcd, 0xce, 0x3e, 0xff, 0x77, 0xfe, 0xf1, 0xf9, 0xe7, 0x00, 0xa8, 0x7f, 0xa6, 0x5f, 0x88,
	0xa8, 0x8d, 0x37, 0x6a, 0x98, 0x87, 0x03, 0x7a, 0x54, 0xc3, 0x8c, 0x92, 0x50, 0xd6, 0x54, 0x39,
	0x18, 0x50, 0x26, 0x89, 0x08, 0xb4, 0xe4, 0x45, 0x82, 0x4b, 0x0e, 0x5f, 0xd0, 0x58, 0x52, 0xee,
	0x29, 0xc0, 0x1b, 0x6f, 0x78, 0x46, 0xd4, 0x7d, 0xa5, 0xb5, 0x23, 0x7e, 0xc4, 0x15, 0x58, 0x4b,
	0x9f, 0x74, 0x4f, 0x69, 0xfd, 0xfa, 0x14, 0x24, 0xa5, 0xa0, 0xbd, 0x91, 0x24, 0xb1, 0x91, 0x5e,
	0x3e, 0xf0, 0x01, 0x28, 0xa2, 0x41, 0x1c, 0x11, 0x6c, 0xb0, 0xca, 0x03, 0xd8, 0xc9, 0x88, 0x4b,
	0xa4, 0x99, 0xca, 0x45, 0x06, 0xc0, 0xdd, 0x14, 0x6b, 0xf1, 0x50, 0x0a, 0xce, 0x5a, 0x8a, 0x84,
	0x6f, 0x00, 0x24, 0x21, 0xea, 0x31, 0x12, 0xe8, 0xa1, 0xf0, 0x90, 0xe0, 0x63, 0xc7, 0x2e, 0xdb,
	0xd5, 0x25, 0x3f, 0xaf, 0x15, 0xdd, 0x95, 0xd6, 0xa1, 0x07, 0x8a, 0x33, 0xb4, 0x4f, 0x22, 0x2e,
	0xa4, 0x93, 0x51, 0x78, 0x61, 0x0a, 0xd7, 0x02, 0x6c, 0x83, 0xbc, 0x06, 0x6f, 0x26, 0x73, 0xe6,
	0xca, 0x76, 0x35, 0xd7, 0x28, 0x79, 0xb7, 0x92, 0x6a, 0x5e, 0x13, 0xfe, 0xaa, 0x2a, 0xde, 0x14,
	0xe0, 0x2e, 0x58, 0x1e, 0x4a, 0x19, 0x05, 0x57, 0x63, 0x3b, 0xf3, 0xe5, 0xb9, 0x6a, 0xae, 0xf1,
	0xda, 0x7b, 0x2c, 0x6d, 0xef, 0xe3, 0xe1, 0x61, 0xb7, 0xd9, 0xdd, 0x3e, 0x88, 0x08, 0xf6, 0x73,
	0x69, 0x7f, 0x33, 0xa2, 0xe9, 0x0b, 0xec, 0x00, 0xa0, 0x92, 0xd1, 0x5e, 0x0b, 0xca, 0xeb, 0xd5,
	0xe3, 0x5e, 0x9f, 0x52, 0x5e, 0x39, 0x65, 0x4f, 0xae, 0x1e, 0x2b, 0x3f, 0x16, 0x41, 0x41, 0x4d,
	0xdb, 0x51, 0x9b, 0x60, 0x12, 0xf5, 0x40, 0xb1, 0x4f, 0x63, 0x15, 0x92, 0x0a, 0x33, 0xc0, 0x08,
	0x0f, 0x89, 0x89, 0xb4, 0x60, 0x24, 0x15, 0x67, 0x2b, 0x15, 0xa6, 0x79, 0xfd, 0x55, 0x9a, 0xcf,
	0xcc, 0xf0, 0xea, 0x7c, 0xcd, 0xd7, 0xc1, 0xda, 0x15, 0x2f, 0x54, 0xca, 0x41, 0x0f, 0x49, 0x3c,
	0x54, 0xb9, 0x2e, 0xf9, 0xd0, 0x68, 0xfa, 0x02, 0x36, 0x53, 0x05, 0x9e, 0x80, 0x62, 0x48, 0xe4,
	0x29, 0x17, 0xc7, 0xc1, 0x00, 0x51, 0x16, 0x44, 0x9c, 0x51, 0x7c, 0xe6, 0xcc, 0x97, 0xed, 0xea,
	0x4a, 0xa3, 0xf9, 0xf8, 0xe0, 0x77, 0xe6, 0xf3, 0xf6, 0xb4, 0x55, 0x07, 0x51, 0xd6, 0x55, 0x46,
	0x7e, 0x21, 0xbc, 0x5d, 0x82, 0x0c, 0xac, 0x62, 0xbd, 0x67, 0xe6, 0xf7, 0x11, 0x9b, 0x9c, 0x5b,
	0x4f, 0x3d, 0x6e, 0x66, 0x5d, 0xe3, 0x76, 0x28, 0xc5, 0x99, 0xbf, 0x82, 0x67, 0x8a, 0xf0, 0x03,
	0x78, 0xde, 0x27, 0x03, 0x34, 0x62, 0x32, 0xe8, 0x93, 0x58, 0xd2, 0x10, 0x49, 0xca, 0xc3, 0x20,
	0x26, 0x62, 0x4c, 0x31, 0x71, 0x16, 0xcb, 0x76, 0x35, 0xeb, 0xaf, 0x1b, 0x64, 0xeb, 0x86, 0x38,
	0xd0, 0xc0, 0xbd, 0x6b, 0xfa, 0xec, 0xe9, 0x6b, 0xba, 0x0d, 0xe0, 0x80, 0x8b, 0x53, 0x24, 0xfa,
	0xd3, 0x46, 0x4b, 0xff, 0x34, 0x2a, 0x98, 0xae, 0x29, 0xab, 0x06, 0x58, 0xdb, 0xd2, 0x17, 0x79,
	0xd8, 0xea, 0x9a, 0x65, 0x61, 0x2c, 0x76, 0xb2, 0xea, 0x92, 0xef, 0xd5, 0x4a, 0x31, 0x28, 0xde,
	0x13, 0x16, 0xcc, 0x83, 0xb9, 0x63, 0x72, 0xa6, 0xf6, 0x2f, 0xeb, 0xa7, 0x8f, 0xb0, 0x03, 0x16,
	0xc6, 0x88, 0x8d, 0xf4, 0x8e, 0xe5, 0x1a, 0xf5, 0xff, 0xb8, 0x92, 0x19, 0x63, 0x5f, 0xb7, 0xbf,
	0xcf, 0xbc, 0xb3, 0x2b, 0x0d, 0x50, 0xb8, 0xb3, 0x10, 0x70, 0x19, 0x64, 0x3b, 0xcd, 0xed, 0x9d,
	0x60, 0xbf, 0xdb, 0xde, 0xcb, 0x5b, 0x70, 0x05, 0x00, 0xf5, 0xda, 0xda, 0xd9, 0x3f, 0x68, 0xe7,
	0xed, 0xcd, 0xfa, 0xf9, 0xc4, 0xb5, 0x2e, 0x26, 0xae, 0xf5, 0x73, 0xe2, 0x5a, 0x97, 0x13, 0xd7,
	0xfa, 0x9a, 0xb8, 0xf6, 0xf7, 0xc4, 0xb5, 0xce, 0x13, 0xd7, 0xbe, 0x48, 0x5c, 0xfb, 0x57, 0xe2,
	0xda, 0x7f, 0x12, 0xd7, 0xba, 0x4c, 0x5c, 0xfb, 0xdb, 0x6f, 0xd7, 0xea, 0x2d, 0xaa, 0xff, 0xb0,
	0xb7, 0x7f, 0x03, 0x00, 0x00, 0xff, 0xff, 0xc0, 0xc0, 0x2b, 0x6d, 0x91, 0x05, 0x00, 0x00,
}