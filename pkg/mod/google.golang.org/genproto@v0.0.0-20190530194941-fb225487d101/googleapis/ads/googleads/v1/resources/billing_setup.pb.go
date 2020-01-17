// Code generated by protoc-gen-go. DO NOT EDIT.
// source: google/ads/googleads/v1/resources/billing_setup.proto

package resources

import (
	fmt "fmt"
	math "math"

	proto "github.com/golang/protobuf/proto"
	wrappers "github.com/golang/protobuf/ptypes/wrappers"
	enums "google.golang.org/genproto/googleapis/ads/googleads/v1/enums"
	_ "google.golang.org/genproto/googleapis/api/annotations"
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

// A billing setup across Ads and Payments systems; an association between a
// Payments account and an advertiser. A billing setup is specific to one
// advertiser.
type BillingSetup struct {
	// The resource name of the billing setup.
	// BillingSetup resource names have the form:
	//
	// `customers/{customer_id}/billingSetups/{billing_setup_id}`
	ResourceName string `protobuf:"bytes,1,opt,name=resource_name,json=resourceName,proto3" json:"resource_name,omitempty"`
	// The ID of the billing setup.
	Id *wrappers.Int64Value `protobuf:"bytes,2,opt,name=id,proto3" json:"id,omitempty"`
	// The status of the billing setup.
	Status enums.BillingSetupStatusEnum_BillingSetupStatus `protobuf:"varint,3,opt,name=status,proto3,enum=google.ads.googleads.v1.enums.BillingSetupStatusEnum_BillingSetupStatus" json:"status,omitempty"`
	// The resource name of the Payments account associated with this billing
	// setup. Payments resource names have the form:
	//
	// `customers/{customer_id}/paymentsAccounts/{payments_account_id}`
	// When setting up billing, this is used to signup with an existing Payments
	// account (and then payments_account_info should not be set).
	// When getting a billing setup, this and payments_account_info will be
	// populated.
	PaymentsAccount *wrappers.StringValue `protobuf:"bytes,11,opt,name=payments_account,json=paymentsAccount,proto3" json:"payments_account,omitempty"`
	// The Payments account information associated with this billing setup.
	// When setting up billing, this is used to signup with a new Payments account
	// (and then payments_account should not be set).
	// When getting a billing setup, this and payments_account will be
	// populated.
	PaymentsAccountInfo *BillingSetup_PaymentsAccountInfo `protobuf:"bytes,12,opt,name=payments_account_info,json=paymentsAccountInfo,proto3" json:"payments_account_info,omitempty"`
	// When creating a new billing setup, this is when the setup should take
	// effect. NOW is the only acceptable start time if the customer doesn't have
	// any approved setups.
	//
	// When fetching an existing billing setup, this is the requested start time.
	// However, if the setup was approved (see status) after the requested start
	// time, then this is the approval time.
	//
	// Types that are valid to be assigned to StartTime:
	//	*BillingSetup_StartDateTime
	//	*BillingSetup_StartTimeType
	StartTime isBillingSetup_StartTime `protobuf_oneof:"start_time"`
	// When the billing setup ends / ended. This is either FOREVER or the start
	// time of the next scheduled billing setup.
	//
	// Types that are valid to be assigned to EndTime:
	//	*BillingSetup_EndDateTime
	//	*BillingSetup_EndTimeType
	EndTime              isBillingSetup_EndTime `protobuf_oneof:"end_time"`
	XXX_NoUnkeyedLiteral struct{}               `json:"-"`
	XXX_unrecognized     []byte                 `json:"-"`
	XXX_sizecache        int32                  `json:"-"`
}

func (m *BillingSetup) Reset()         { *m = BillingSetup{} }
func (m *BillingSetup) String() string { return proto.CompactTextString(m) }
func (*BillingSetup) ProtoMessage()    {}
func (*BillingSetup) Descriptor() ([]byte, []int) {
	return fileDescriptor_4372b9605c7b101a, []int{0}
}

func (m *BillingSetup) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_BillingSetup.Unmarshal(m, b)
}
func (m *BillingSetup) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_BillingSetup.Marshal(b, m, deterministic)
}
func (m *BillingSetup) XXX_Merge(src proto.Message) {
	xxx_messageInfo_BillingSetup.Merge(m, src)
}
func (m *BillingSetup) XXX_Size() int {
	return xxx_messageInfo_BillingSetup.Size(m)
}
func (m *BillingSetup) XXX_DiscardUnknown() {
	xxx_messageInfo_BillingSetup.DiscardUnknown(m)
}

var xxx_messageInfo_BillingSetup proto.InternalMessageInfo

func (m *BillingSetup) GetResourceName() string {
	if m != nil {
		return m.ResourceName
	}
	return ""
}

func (m *BillingSetup) GetId() *wrappers.Int64Value {
	if m != nil {
		return m.Id
	}
	return nil
}

func (m *BillingSetup) GetStatus() enums.BillingSetupStatusEnum_BillingSetupStatus {
	if m != nil {
		return m.Status
	}
	return enums.BillingSetupStatusEnum_UNSPECIFIED
}

func (m *BillingSetup) GetPaymentsAccount() *wrappers.StringValue {
	if m != nil {
		return m.PaymentsAccount
	}
	return nil
}

func (m *BillingSetup) GetPaymentsAccountInfo() *BillingSetup_PaymentsAccountInfo {
	if m != nil {
		return m.PaymentsAccountInfo
	}
	return nil
}

type isBillingSetup_StartTime interface {
	isBillingSetup_StartTime()
}

type BillingSetup_StartDateTime struct {
	StartDateTime *wrappers.StringValue `protobuf:"bytes,9,opt,name=start_date_time,json=startDateTime,proto3,oneof"`
}

type BillingSetup_StartTimeType struct {
	StartTimeType enums.TimeTypeEnum_TimeType `protobuf:"varint,10,opt,name=start_time_type,json=startTimeType,proto3,enum=google.ads.googleads.v1.enums.TimeTypeEnum_TimeType,oneof"`
}

func (*BillingSetup_StartDateTime) isBillingSetup_StartTime() {}

func (*BillingSetup_StartTimeType) isBillingSetup_StartTime() {}

func (m *BillingSetup) GetStartTime() isBillingSetup_StartTime {
	if m != nil {
		return m.StartTime
	}
	return nil
}

func (m *BillingSetup) GetStartDateTime() *wrappers.StringValue {
	if x, ok := m.GetStartTime().(*BillingSetup_StartDateTime); ok {
		return x.StartDateTime
	}
	return nil
}

func (m *BillingSetup) GetStartTimeType() enums.TimeTypeEnum_TimeType {
	if x, ok := m.GetStartTime().(*BillingSetup_StartTimeType); ok {
		return x.StartTimeType
	}
	return enums.TimeTypeEnum_UNSPECIFIED
}

type isBillingSetup_EndTime interface {
	isBillingSetup_EndTime()
}

type BillingSetup_EndDateTime struct {
	EndDateTime *wrappers.StringValue `protobuf:"bytes,13,opt,name=end_date_time,json=endDateTime,proto3,oneof"`
}

type BillingSetup_EndTimeType struct {
	EndTimeType enums.TimeTypeEnum_TimeType `protobuf:"varint,14,opt,name=end_time_type,json=endTimeType,proto3,enum=google.ads.googleads.v1.enums.TimeTypeEnum_TimeType,oneof"`
}

func (*BillingSetup_EndDateTime) isBillingSetup_EndTime() {}

func (*BillingSetup_EndTimeType) isBillingSetup_EndTime() {}

func (m *BillingSetup) GetEndTime() isBillingSetup_EndTime {
	if m != nil {
		return m.EndTime
	}
	return nil
}

func (m *BillingSetup) GetEndDateTime() *wrappers.StringValue {
	if x, ok := m.GetEndTime().(*BillingSetup_EndDateTime); ok {
		return x.EndDateTime
	}
	return nil
}

func (m *BillingSetup) GetEndTimeType() enums.TimeTypeEnum_TimeType {
	if x, ok := m.GetEndTime().(*BillingSetup_EndTimeType); ok {
		return x.EndTimeType
	}
	return enums.TimeTypeEnum_UNSPECIFIED
}

// XXX_OneofWrappers is for the internal use of the proto package.
func (*BillingSetup) XXX_OneofWrappers() []interface{} {
	return []interface{}{
		(*BillingSetup_StartDateTime)(nil),
		(*BillingSetup_StartTimeType)(nil),
		(*BillingSetup_EndDateTime)(nil),
		(*BillingSetup_EndTimeType)(nil),
	}
}

// Container of Payments account information for this billing.
type BillingSetup_PaymentsAccountInfo struct {
	// A 16 digit id used to identify the Payments account associated with the
	// billing setup.
	//
	// This must be passed as a string with dashes, e.g. "1234-5678-9012-3456".
	PaymentsAccountId *wrappers.StringValue `protobuf:"bytes,1,opt,name=payments_account_id,json=paymentsAccountId,proto3" json:"payments_account_id,omitempty"`
	// The name of the Payments account associated with the billing setup.
	//
	// This enables the user to specify a meaningful name for a Payments account
	// to aid in reconciling monthly invoices.
	//
	// This name will be printed in the monthly invoices.
	PaymentsAccountName *wrappers.StringValue `protobuf:"bytes,2,opt,name=payments_account_name,json=paymentsAccountName,proto3" json:"payments_account_name,omitempty"`
	// A 12 digit id used to identify the Payments profile associated with the
	// billing setup.
	//
	// This must be passed in as a string with dashes, e.g. "1234-5678-9012".
	PaymentsProfileId *wrappers.StringValue `protobuf:"bytes,3,opt,name=payments_profile_id,json=paymentsProfileId,proto3" json:"payments_profile_id,omitempty"`
	// The name of the Payments profile associated with the billing setup.
	PaymentsProfileName *wrappers.StringValue `protobuf:"bytes,4,opt,name=payments_profile_name,json=paymentsProfileName,proto3" json:"payments_profile_name,omitempty"`
	// A secondary payments profile id present in uncommon situations, e.g.
	// when a sequential liability agreement has been arranged.
	SecondaryPaymentsProfileId *wrappers.StringValue `protobuf:"bytes,5,opt,name=secondary_payments_profile_id,json=secondaryPaymentsProfileId,proto3" json:"secondary_payments_profile_id,omitempty"`
	XXX_NoUnkeyedLiteral       struct{}              `json:"-"`
	XXX_unrecognized           []byte                `json:"-"`
	XXX_sizecache              int32                 `json:"-"`
}

func (m *BillingSetup_PaymentsAccountInfo) Reset()         { *m = BillingSetup_PaymentsAccountInfo{} }
func (m *BillingSetup_PaymentsAccountInfo) String() string { return proto.CompactTextString(m) }
func (*BillingSetup_PaymentsAccountInfo) ProtoMessage()    {}
func (*BillingSetup_PaymentsAccountInfo) Descriptor() ([]byte, []int) {
	return fileDescriptor_4372b9605c7b101a, []int{0, 0}
}

func (m *BillingSetup_PaymentsAccountInfo) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_BillingSetup_PaymentsAccountInfo.Unmarshal(m, b)
}
func (m *BillingSetup_PaymentsAccountInfo) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_BillingSetup_PaymentsAccountInfo.Marshal(b, m, deterministic)
}
func (m *BillingSetup_PaymentsAccountInfo) XXX_Merge(src proto.Message) {
	xxx_messageInfo_BillingSetup_PaymentsAccountInfo.Merge(m, src)
}
func (m *BillingSetup_PaymentsAccountInfo) XXX_Size() int {
	return xxx_messageInfo_BillingSetup_PaymentsAccountInfo.Size(m)
}
func (m *BillingSetup_PaymentsAccountInfo) XXX_DiscardUnknown() {
	xxx_messageInfo_BillingSetup_PaymentsAccountInfo.DiscardUnknown(m)
}

var xxx_messageInfo_BillingSetup_PaymentsAccountInfo proto.InternalMessageInfo

func (m *BillingSetup_PaymentsAccountInfo) GetPaymentsAccountId() *wrappers.StringValue {
	if m != nil {
		return m.PaymentsAccountId
	}
	return nil
}

func (m *BillingSetup_PaymentsAccountInfo) GetPaymentsAccountName() *wrappers.StringValue {
	if m != nil {
		return m.PaymentsAccountName
	}
	return nil
}

func (m *BillingSetup_PaymentsAccountInfo) GetPaymentsProfileId() *wrappers.StringValue {
	if m != nil {
		return m.PaymentsProfileId
	}
	return nil
}

func (m *BillingSetup_PaymentsAccountInfo) GetPaymentsProfileName() *wrappers.StringValue {
	if m != nil {
		return m.PaymentsProfileName
	}
	return nil
}

func (m *BillingSetup_PaymentsAccountInfo) GetSecondaryPaymentsProfileId() *wrappers.StringValue {
	if m != nil {
		return m.SecondaryPaymentsProfileId
	}
	return nil
}

func init() {
	proto.RegisterType((*BillingSetup)(nil), "google.ads.googleads.v1.resources.BillingSetup")
	proto.RegisterType((*BillingSetup_PaymentsAccountInfo)(nil), "google.ads.googleads.v1.resources.BillingSetup.PaymentsAccountInfo")
}

func init() {
	proto.RegisterFile("google/ads/googleads/v1/resources/billing_setup.proto", fileDescriptor_4372b9605c7b101a)
}

var fileDescriptor_4372b9605c7b101a = []byte{
	// 631 bytes of a gzipped FileDescriptorProto
	0x1f, 0x8b, 0x08, 0x00, 0x00, 0x00, 0x00, 0x00, 0x02, 0xff, 0xa4, 0x95, 0xdd, 0x6e, 0xd3, 0x30,
	0x14, 0x80, 0x49, 0x0a, 0x13, 0xf3, 0xda, 0x8d, 0x65, 0x42, 0x8a, 0xca, 0x40, 0x1b, 0x68, 0xd2,
	0x24, 0x84, 0xa3, 0x8e, 0x81, 0x50, 0xb8, 0x6a, 0xf9, 0xd9, 0x86, 0x10, 0xaa, 0xb2, 0xa9, 0x17,
	0x53, 0x45, 0xf0, 0x6a, 0x37, 0xb2, 0x94, 0xd8, 0x51, 0xec, 0x6c, 0xea, 0xd3, 0x20, 0x71, 0xc9,
	0x03, 0xf0, 0x10, 0x3c, 0x0a, 0x0f, 0x81, 0x50, 0xec, 0x38, 0xeb, 0xb2, 0x6e, 0xdd, 0xc4, 0x9d,
	0xed, 0x9c, 0xf3, 0xe9, 0xf3, 0x39, 0xc7, 0x2d, 0x78, 0x15, 0x71, 0x1e, 0xc5, 0xc4, 0x43, 0x58,
	0x78, 0x7a, 0x59, 0xac, 0x4e, 0x3b, 0x5e, 0x46, 0x04, 0xcf, 0xb3, 0x11, 0x11, 0xde, 0x09, 0x8d,
	0x63, 0xca, 0xa2, 0x50, 0x10, 0x99, 0xa7, 0x30, 0xcd, 0xb8, 0xe4, 0xce, 0xa6, 0x8e, 0x85, 0x08,
	0x0b, 0x58, 0xa5, 0xc1, 0xd3, 0x0e, 0xac, 0xd2, 0xda, 0x6f, 0xae, 0x22, 0x13, 0x96, 0x27, 0x35,
	0x6a, 0x28, 0x24, 0x92, 0xb9, 0xd0, 0xf0, 0xf6, 0x8b, 0xeb, 0x33, 0x25, 0x4d, 0x48, 0x28, 0x27,
	0x29, 0x29, 0xc3, 0x9f, 0x94, 0xe1, 0x6a, 0x77, 0x92, 0x8f, 0xbd, 0xb3, 0x0c, 0xa5, 0x29, 0xc9,
	0x0c, 0x6e, 0xdd, 0xe0, 0x52, 0xea, 0x21, 0xc6, 0xb8, 0x44, 0x92, 0x72, 0x56, 0x7e, 0x7d, 0xfa,
	0x7d, 0x11, 0x34, 0x7b, 0xda, 0xe5, 0xb0, 0x50, 0x71, 0x9e, 0x81, 0x96, 0xb9, 0x44, 0xc8, 0x50,
	0x42, 0x5c, 0x6b, 0xc3, 0xda, 0x5e, 0x0c, 0x9a, 0xe6, 0xf0, 0x0b, 0x4a, 0x88, 0xf3, 0x1c, 0xd8,
	0x14, 0xbb, 0xf6, 0x86, 0xb5, 0xbd, 0xb4, 0xf3, 0xa8, 0xac, 0x00, 0x34, 0x02, 0xf0, 0x80, 0xc9,
	0xd7, 0xbb, 0x03, 0x14, 0xe7, 0x24, 0xb0, 0x29, 0x76, 0xbe, 0x81, 0x05, 0x7d, 0x3f, 0xb7, 0xb1,
	0x61, 0x6d, 0x2f, 0xef, 0xec, 0xc3, 0xab, 0xaa, 0xa7, 0x2e, 0x08, 0xa7, 0x75, 0x0e, 0x55, 0xe2,
	0x07, 0x96, 0x27, 0x33, 0x8e, 0x83, 0x92, 0xeb, 0xec, 0x81, 0x07, 0x29, 0x9a, 0x24, 0x84, 0x49,
	0x11, 0xa2, 0xd1, 0x88, 0xe7, 0x4c, 0xba, 0x4b, 0x4a, 0x6e, 0xfd, 0x92, 0xdc, 0xa1, 0xcc, 0x28,
	0x8b, 0xb4, 0xdd, 0x8a, 0xc9, 0xea, 0xea, 0x24, 0xe7, 0x0c, 0x3c, 0xac, 0x83, 0x42, 0xca, 0xc6,
	0xdc, 0x6d, 0x2a, 0xda, 0x3b, 0x38, 0xb7, 0xef, 0x17, 0x34, 0x61, 0xff, 0x22, 0xff, 0x80, 0x8d,
	0x79, 0xb0, 0x96, 0x5e, 0x3e, 0x74, 0x3e, 0x82, 0x15, 0x21, 0x51, 0x26, 0x43, 0x8c, 0x24, 0x09,
	0x8b, 0x16, 0xbb, 0x8b, 0xf3, 0x2f, 0xb0, 0x7f, 0x27, 0x68, 0xa9, 0xb4, 0xf7, 0x48, 0x92, 0x23,
	0x9a, 0x10, 0xe7, 0xab, 0xe1, 0x54, 0x53, 0xe2, 0x02, 0x55, 0xf4, 0xdd, 0x39, 0x45, 0x2f, 0xb2,
	0x8f, 0x26, 0x29, 0x51, 0xa5, 0x36, 0x9b, 0x8a, 0x6f, 0x0e, 0x9c, 0x1e, 0x68, 0x11, 0x86, 0xa7,
	0x2c, 0x5b, 0x37, 0xb0, 0xb4, 0x82, 0x25, 0xc2, 0x70, 0xe5, 0x78, 0xac, 0x19, 0xe7, 0x86, 0xcb,
	0xff, 0x61, 0xa8, 0xd9, 0x66, 0xdb, 0xfe, 0xd5, 0x00, 0x6b, 0x33, 0x8a, 0xee, 0x7c, 0x06, 0x6b,
	0x97, 0x1b, 0x8b, 0xd5, 0x6c, 0xcf, 0x1b, 0x92, 0xd5, 0x7a, 0xbf, 0xb0, 0xd3, 0x9f, 0x31, 0x26,
	0xea, 0xad, 0xd8, 0x37, 0xe0, 0xd5, 0xfb, 0xaf, 0x1e, 0xd4, 0xb4, 0x5f, 0x9a, 0xf1, 0x31, 0x8d,
	0x49, 0xe1, 0xd7, 0xb8, 0x8d, 0x5f, 0x5f, 0xe7, 0xd5, 0xfc, 0x0c, 0x4d, 0xf9, 0xdd, 0xbd, 0x8d,
	0x5f, 0xc9, 0x53, 0x7e, 0x21, 0x78, 0x2c, 0xc8, 0x88, 0x33, 0x8c, 0xb2, 0x49, 0x38, 0xcb, 0xf4,
	0xde, 0x0d, 0xc8, 0xed, 0x0a, 0xd1, 0xaf, 0x2b, 0xf7, 0x9a, 0x00, 0x9c, 0x0f, 0x6e, 0x0f, 0x80,
	0xfb, 0x66, 0x44, 0x7a, 0x7f, 0x2d, 0xb0, 0x35, 0xe2, 0xc9, 0xfc, 0xa7, 0xd7, 0x5b, 0x9d, 0x7e,
	0x7b, 0xfd, 0xc2, 0xa0, 0x6f, 0x1d, 0x7f, 0x2a, 0xf3, 0x22, 0x1e, 0x23, 0x16, 0x41, 0x9e, 0x45,
	0x5e, 0x44, 0x98, 0xf2, 0x33, 0xbf, 0xae, 0x29, 0x15, 0xd7, 0xfc, 0x01, 0xbc, 0xad, 0x56, 0x3f,
	0xec, 0xc6, 0x5e, 0xb7, 0xfb, 0xd3, 0xde, 0xdc, 0xd3, 0xc8, 0x2e, 0x16, 0x50, 0x2f, 0x8b, 0xd5,
	0xa0, 0x03, 0x03, 0x13, 0xf9, 0xdb, 0xc4, 0x0c, 0xbb, 0x58, 0x0c, 0xab, 0x98, 0xe1, 0xa0, 0x33,
	0xac, 0x62, 0xfe, 0xd8, 0x5b, 0xfa, 0x83, 0xef, 0x77, 0xb1, 0xf0, 0xfd, 0x2a, 0xca, 0xf7, 0x07,
	0x1d, 0xdf, 0xaf, 0xe2, 0x4e, 0x16, 0x94, 0xec, 0xcb, 0x7f, 0x01, 0x00, 0x00, 0xff, 0xff, 0x50,
	0x93, 0xbd, 0xf6, 0xac, 0x06, 0x00, 0x00,
}
