// Code generated by protoc-gen-go.
// source: steammessages_deviceauth.steamclient.proto
// DO NOT EDIT!

package unified

import proto "github.com/golang/protobuf/proto"
import math "math"

// Reference imports to suppress errors if they are not otherwise used.
var _ = proto.Marshal
var _ = math.Inf

type CDeviceAuth_GetOwnAuthorizedDevices_Request struct {
	Steamid          *uint64 `protobuf:"fixed64,1,opt,name=steamid" json:"steamid,omitempty"`
	IncludeCanceled  *bool   `protobuf:"varint,2,opt,name=include_canceled" json:"include_canceled,omitempty"`
	XXX_unrecognized []byte  `json:"-"`
}

func (m *CDeviceAuth_GetOwnAuthorizedDevices_Request) Reset() {
	*m = CDeviceAuth_GetOwnAuthorizedDevices_Request{}
}
func (m *CDeviceAuth_GetOwnAuthorizedDevices_Request) String() string {
	return proto.CompactTextString(m)
}
func (*CDeviceAuth_GetOwnAuthorizedDevices_Request) ProtoMessage() {}

func (m *CDeviceAuth_GetOwnAuthorizedDevices_Request) GetSteamid() uint64 {
	if m != nil && m.Steamid != nil {
		return *m.Steamid
	}
	return 0
}

func (m *CDeviceAuth_GetOwnAuthorizedDevices_Request) GetIncludeCanceled() bool {
	if m != nil && m.IncludeCanceled != nil {
		return *m.IncludeCanceled
	}
	return false
}

type CDeviceAuth_GetOwnAuthorizedDevices_Response struct {
	Devices          []*CDeviceAuth_GetOwnAuthorizedDevices_Response_Device `protobuf:"bytes,1,rep,name=devices" json:"devices,omitempty"`
	XXX_unrecognized []byte                                                 `json:"-"`
}

func (m *CDeviceAuth_GetOwnAuthorizedDevices_Response) Reset() {
	*m = CDeviceAuth_GetOwnAuthorizedDevices_Response{}
}
func (m *CDeviceAuth_GetOwnAuthorizedDevices_Response) String() string {
	return proto.CompactTextString(m)
}
func (*CDeviceAuth_GetOwnAuthorizedDevices_Response) ProtoMessage() {}

func (m *CDeviceAuth_GetOwnAuthorizedDevices_Response) GetDevices() []*CDeviceAuth_GetOwnAuthorizedDevices_Response_Device {
	if m != nil {
		return m.Devices
	}
	return nil
}

type CDeviceAuth_GetOwnAuthorizedDevices_Response_Device struct {
	AuthDeviceToken  *uint64 `protobuf:"fixed64,1,opt,name=auth_device_token" json:"auth_device_token,omitempty"`
	DeviceName       *string `protobuf:"bytes,2,opt,name=device_name" json:"device_name,omitempty"`
	IsPending        *bool   `protobuf:"varint,3,opt,name=is_pending" json:"is_pending,omitempty"`
	IsCanceled       *bool   `protobuf:"varint,4,opt,name=is_canceled" json:"is_canceled,omitempty"`
	LastTimeUsed     *uint32 `protobuf:"varint,5,opt,name=last_time_used" json:"last_time_used,omitempty"`
	LastBorrowerId   *uint64 `protobuf:"fixed64,6,opt,name=last_borrower_id" json:"last_borrower_id,omitempty"`
	LastAppPlayed    *uint32 `protobuf:"varint,7,opt,name=last_app_played" json:"last_app_played,omitempty"`
	IsLimited        *bool   `protobuf:"varint,8,opt,name=is_limited" json:"is_limited,omitempty"`
	XXX_unrecognized []byte  `json:"-"`
}

func (m *CDeviceAuth_GetOwnAuthorizedDevices_Response_Device) Reset() {
	*m = CDeviceAuth_GetOwnAuthorizedDevices_Response_Device{}
}
func (m *CDeviceAuth_GetOwnAuthorizedDevices_Response_Device) String() string {
	return proto.CompactTextString(m)
}
func (*CDeviceAuth_GetOwnAuthorizedDevices_Response_Device) ProtoMessage() {}

func (m *CDeviceAuth_GetOwnAuthorizedDevices_Response_Device) GetAuthDeviceToken() uint64 {
	if m != nil && m.AuthDeviceToken != nil {
		return *m.AuthDeviceToken
	}
	return 0
}

func (m *CDeviceAuth_GetOwnAuthorizedDevices_Response_Device) GetDeviceName() string {
	if m != nil && m.DeviceName != nil {
		return *m.DeviceName
	}
	return ""
}

func (m *CDeviceAuth_GetOwnAuthorizedDevices_Response_Device) GetIsPending() bool {
	if m != nil && m.IsPending != nil {
		return *m.IsPending
	}
	return false
}

func (m *CDeviceAuth_GetOwnAuthorizedDevices_Response_Device) GetIsCanceled() bool {
	if m != nil && m.IsCanceled != nil {
		return *m.IsCanceled
	}
	return false
}

func (m *CDeviceAuth_GetOwnAuthorizedDevices_Response_Device) GetLastTimeUsed() uint32 {
	if m != nil && m.LastTimeUsed != nil {
		return *m.LastTimeUsed
	}
	return 0
}

func (m *CDeviceAuth_GetOwnAuthorizedDevices_Response_Device) GetLastBorrowerId() uint64 {
	if m != nil && m.LastBorrowerId != nil {
		return *m.LastBorrowerId
	}
	return 0
}

func (m *CDeviceAuth_GetOwnAuthorizedDevices_Response_Device) GetLastAppPlayed() uint32 {
	if m != nil && m.LastAppPlayed != nil {
		return *m.LastAppPlayed
	}
	return 0
}

func (m *CDeviceAuth_GetOwnAuthorizedDevices_Response_Device) GetIsLimited() bool {
	if m != nil && m.IsLimited != nil {
		return *m.IsLimited
	}
	return false
}

type CDeviceAuth_AcceptAuthorizationRequest_Request struct {
	Steamid          *uint64 `protobuf:"fixed64,1,opt,name=steamid" json:"steamid,omitempty"`
	AuthDeviceToken  *uint64 `protobuf:"fixed64,2,opt,name=auth_device_token" json:"auth_device_token,omitempty"`
	AuthCode         *uint64 `protobuf:"fixed64,3,opt,name=auth_code" json:"auth_code,omitempty"`
	FromSteamid      *uint64 `protobuf:"fixed64,4,opt,name=from_steamid" json:"from_steamid,omitempty"`
	XXX_unrecognized []byte  `json:"-"`
}

func (m *CDeviceAuth_AcceptAuthorizationRequest_Request) Reset() {
	*m = CDeviceAuth_AcceptAuthorizationRequest_Request{}
}
func (m *CDeviceAuth_AcceptAuthorizationRequest_Request) String() string {
	return proto.CompactTextString(m)
}
func (*CDeviceAuth_AcceptAuthorizationRequest_Request) ProtoMessage() {}

func (m *CDeviceAuth_AcceptAuthorizationRequest_Request) GetSteamid() uint64 {
	if m != nil && m.Steamid != nil {
		return *m.Steamid
	}
	return 0
}

func (m *CDeviceAuth_AcceptAuthorizationRequest_Request) GetAuthDeviceToken() uint64 {
	if m != nil && m.AuthDeviceToken != nil {
		return *m.AuthDeviceToken
	}
	return 0
}

func (m *CDeviceAuth_AcceptAuthorizationRequest_Request) GetAuthCode() uint64 {
	if m != nil && m.AuthCode != nil {
		return *m.AuthCode
	}
	return 0
}

func (m *CDeviceAuth_AcceptAuthorizationRequest_Request) GetFromSteamid() uint64 {
	if m != nil && m.FromSteamid != nil {
		return *m.FromSteamid
	}
	return 0
}

type CDeviceAuth_AcceptAuthorizationRequest_Response struct {
	XXX_unrecognized []byte `json:"-"`
}

func (m *CDeviceAuth_AcceptAuthorizationRequest_Response) Reset() {
	*m = CDeviceAuth_AcceptAuthorizationRequest_Response{}
}
func (m *CDeviceAuth_AcceptAuthorizationRequest_Response) String() string {
	return proto.CompactTextString(m)
}
func (*CDeviceAuth_AcceptAuthorizationRequest_Response) ProtoMessage() {}

type CDeviceAuth_AuthorizeRemoteDevice_Request struct {
	Steamid          *uint64 `protobuf:"fixed64,1,opt,name=steamid" json:"steamid,omitempty"`
	AuthDeviceToken  *uint64 `protobuf:"fixed64,2,opt,name=auth_device_token" json:"auth_device_token,omitempty"`
	XXX_unrecognized []byte  `json:"-"`
}

func (m *CDeviceAuth_AuthorizeRemoteDevice_Request) Reset() {
	*m = CDeviceAuth_AuthorizeRemoteDevice_Request{}
}
func (m *CDeviceAuth_AuthorizeRemoteDevice_Request) String() string { return proto.CompactTextString(m) }
func (*CDeviceAuth_AuthorizeRemoteDevice_Request) ProtoMessage()    {}

func (m *CDeviceAuth_AuthorizeRemoteDevice_Request) GetSteamid() uint64 {
	if m != nil && m.Steamid != nil {
		return *m.Steamid
	}
	return 0
}

func (m *CDeviceAuth_AuthorizeRemoteDevice_Request) GetAuthDeviceToken() uint64 {
	if m != nil && m.AuthDeviceToken != nil {
		return *m.AuthDeviceToken
	}
	return 0
}

type CDeviceAuth_AuthorizeRemoteDevice_Response struct {
	XXX_unrecognized []byte `json:"-"`
}

func (m *CDeviceAuth_AuthorizeRemoteDevice_Response) Reset() {
	*m = CDeviceAuth_AuthorizeRemoteDevice_Response{}
}
func (m *CDeviceAuth_AuthorizeRemoteDevice_Response) String() string {
	return proto.CompactTextString(m)
}
func (*CDeviceAuth_AuthorizeRemoteDevice_Response) ProtoMessage() {}

type CDeviceAuth_DeauthorizeRemoteDevice_Request struct {
	Steamid          *uint64 `protobuf:"fixed64,1,opt,name=steamid" json:"steamid,omitempty"`
	AuthDeviceToken  *uint64 `protobuf:"fixed64,2,opt,name=auth_device_token" json:"auth_device_token,omitempty"`
	XXX_unrecognized []byte  `json:"-"`
}

func (m *CDeviceAuth_DeauthorizeRemoteDevice_Request) Reset() {
	*m = CDeviceAuth_DeauthorizeRemoteDevice_Request{}
}
func (m *CDeviceAuth_DeauthorizeRemoteDevice_Request) String() string {
	return proto.CompactTextString(m)
}
func (*CDeviceAuth_DeauthorizeRemoteDevice_Request) ProtoMessage() {}

func (m *CDeviceAuth_DeauthorizeRemoteDevice_Request) GetSteamid() uint64 {
	if m != nil && m.Steamid != nil {
		return *m.Steamid
	}
	return 0
}

func (m *CDeviceAuth_DeauthorizeRemoteDevice_Request) GetAuthDeviceToken() uint64 {
	if m != nil && m.AuthDeviceToken != nil {
		return *m.AuthDeviceToken
	}
	return 0
}

type CDeviceAuth_DeauthorizeRemoteDevice_Response struct {
	XXX_unrecognized []byte `json:"-"`
}

func (m *CDeviceAuth_DeauthorizeRemoteDevice_Response) Reset() {
	*m = CDeviceAuth_DeauthorizeRemoteDevice_Response{}
}
func (m *CDeviceAuth_DeauthorizeRemoteDevice_Response) String() string {
	return proto.CompactTextString(m)
}
func (*CDeviceAuth_DeauthorizeRemoteDevice_Response) ProtoMessage() {}

type CDeviceAuth_GetUsedAuthorizedDevices_Request struct {
	Steamid          *uint64 `protobuf:"fixed64,1,opt,name=steamid" json:"steamid,omitempty"`
	XXX_unrecognized []byte  `json:"-"`
}

func (m *CDeviceAuth_GetUsedAuthorizedDevices_Request) Reset() {
	*m = CDeviceAuth_GetUsedAuthorizedDevices_Request{}
}
func (m *CDeviceAuth_GetUsedAuthorizedDevices_Request) String() string {
	return proto.CompactTextString(m)
}
func (*CDeviceAuth_GetUsedAuthorizedDevices_Request) ProtoMessage() {}

func (m *CDeviceAuth_GetUsedAuthorizedDevices_Request) GetSteamid() uint64 {
	if m != nil && m.Steamid != nil {
		return *m.Steamid
	}
	return 0
}

type CDeviceAuth_GetUsedAuthorizedDevices_Response struct {
	Devices          []*CDeviceAuth_GetUsedAuthorizedDevices_Response_Device `protobuf:"bytes,1,rep,name=devices" json:"devices,omitempty"`
	XXX_unrecognized []byte                                                  `json:"-"`
}

func (m *CDeviceAuth_GetUsedAuthorizedDevices_Response) Reset() {
	*m = CDeviceAuth_GetUsedAuthorizedDevices_Response{}
}
func (m *CDeviceAuth_GetUsedAuthorizedDevices_Response) String() string {
	return proto.CompactTextString(m)
}
func (*CDeviceAuth_GetUsedAuthorizedDevices_Response) ProtoMessage() {}

func (m *CDeviceAuth_GetUsedAuthorizedDevices_Response) GetDevices() []*CDeviceAuth_GetUsedAuthorizedDevices_Response_Device {
	if m != nil {
		return m.Devices
	}
	return nil
}

type CDeviceAuth_GetUsedAuthorizedDevices_Response_Device struct {
	AuthDeviceToken  *uint64 `protobuf:"fixed64,1,opt,name=auth_device_token" json:"auth_device_token,omitempty"`
	DeviceName       *string `protobuf:"bytes,2,opt,name=device_name" json:"device_name,omitempty"`
	OwnerSteamid     *uint64 `protobuf:"fixed64,3,opt,name=owner_steamid" json:"owner_steamid,omitempty"`
	LastTimeUsed     *uint32 `protobuf:"varint,4,opt,name=last_time_used" json:"last_time_used,omitempty"`
	LastAppPlayed    *uint32 `protobuf:"varint,5,opt,name=last_app_played" json:"last_app_played,omitempty"`
	XXX_unrecognized []byte  `json:"-"`
}

func (m *CDeviceAuth_GetUsedAuthorizedDevices_Response_Device) Reset() {
	*m = CDeviceAuth_GetUsedAuthorizedDevices_Response_Device{}
}
func (m *CDeviceAuth_GetUsedAuthorizedDevices_Response_Device) String() string {
	return proto.CompactTextString(m)
}
func (*CDeviceAuth_GetUsedAuthorizedDevices_Response_Device) ProtoMessage() {}

func (m *CDeviceAuth_GetUsedAuthorizedDevices_Response_Device) GetAuthDeviceToken() uint64 {
	if m != nil && m.AuthDeviceToken != nil {
		return *m.AuthDeviceToken
	}
	return 0
}

func (m *CDeviceAuth_GetUsedAuthorizedDevices_Response_Device) GetDeviceName() string {
	if m != nil && m.DeviceName != nil {
		return *m.DeviceName
	}
	return ""
}

func (m *CDeviceAuth_GetUsedAuthorizedDevices_Response_Device) GetOwnerSteamid() uint64 {
	if m != nil && m.OwnerSteamid != nil {
		return *m.OwnerSteamid
	}
	return 0
}

func (m *CDeviceAuth_GetUsedAuthorizedDevices_Response_Device) GetLastTimeUsed() uint32 {
	if m != nil && m.LastTimeUsed != nil {
		return *m.LastTimeUsed
	}
	return 0
}

func (m *CDeviceAuth_GetUsedAuthorizedDevices_Response_Device) GetLastAppPlayed() uint32 {
	if m != nil && m.LastAppPlayed != nil {
		return *m.LastAppPlayed
	}
	return 0
}

type CDeviceAuth_GetAuthorizedBorrowers_Request struct {
	Steamid          *uint64 `protobuf:"fixed64,1,opt,name=steamid" json:"steamid,omitempty"`
	IncludeCanceled  *bool   `protobuf:"varint,2,opt,name=include_canceled" json:"include_canceled,omitempty"`
	IncludePending   *bool   `protobuf:"varint,3,opt,name=include_pending" json:"include_pending,omitempty"`
	XXX_unrecognized []byte  `json:"-"`
}

func (m *CDeviceAuth_GetAuthorizedBorrowers_Request) Reset() {
	*m = CDeviceAuth_GetAuthorizedBorrowers_Request{}
}
func (m *CDeviceAuth_GetAuthorizedBorrowers_Request) String() string {
	return proto.CompactTextString(m)
}
func (*CDeviceAuth_GetAuthorizedBorrowers_Request) ProtoMessage() {}

func (m *CDeviceAuth_GetAuthorizedBorrowers_Request) GetSteamid() uint64 {
	if m != nil && m.Steamid != nil {
		return *m.Steamid
	}
	return 0
}

func (m *CDeviceAuth_GetAuthorizedBorrowers_Request) GetIncludeCanceled() bool {
	if m != nil && m.IncludeCanceled != nil {
		return *m.IncludeCanceled
	}
	return false
}

func (m *CDeviceAuth_GetAuthorizedBorrowers_Request) GetIncludePending() bool {
	if m != nil && m.IncludePending != nil {
		return *m.IncludePending
	}
	return false
}

type CDeviceAuth_GetAuthorizedBorrowers_Response struct {
	Borrowers        []*CDeviceAuth_GetAuthorizedBorrowers_Response_Borrower `protobuf:"bytes,1,rep,name=borrowers" json:"borrowers,omitempty"`
	XXX_unrecognized []byte                                                  `json:"-"`
}

func (m *CDeviceAuth_GetAuthorizedBorrowers_Response) Reset() {
	*m = CDeviceAuth_GetAuthorizedBorrowers_Response{}
}
func (m *CDeviceAuth_GetAuthorizedBorrowers_Response) String() string {
	return proto.CompactTextString(m)
}
func (*CDeviceAuth_GetAuthorizedBorrowers_Response) ProtoMessage() {}

func (m *CDeviceAuth_GetAuthorizedBorrowers_Response) GetBorrowers() []*CDeviceAuth_GetAuthorizedBorrowers_Response_Borrower {
	if m != nil {
		return m.Borrowers
	}
	return nil
}

type CDeviceAuth_GetAuthorizedBorrowers_Response_Borrower struct {
	Steamid          *uint64 `protobuf:"fixed64,1,opt,name=steamid" json:"steamid,omitempty"`
	IsPending        *bool   `protobuf:"varint,2,opt,name=is_pending" json:"is_pending,omitempty"`
	IsCanceled       *bool   `protobuf:"varint,3,opt,name=is_canceled" json:"is_canceled,omitempty"`
	TimeCreated      *uint32 `protobuf:"varint,4,opt,name=time_created" json:"time_created,omitempty"`
	XXX_unrecognized []byte  `json:"-"`
}

func (m *CDeviceAuth_GetAuthorizedBorrowers_Response_Borrower) Reset() {
	*m = CDeviceAuth_GetAuthorizedBorrowers_Response_Borrower{}
}
func (m *CDeviceAuth_GetAuthorizedBorrowers_Response_Borrower) String() string {
	return proto.CompactTextString(m)
}
func (*CDeviceAuth_GetAuthorizedBorrowers_Response_Borrower) ProtoMessage() {}

func (m *CDeviceAuth_GetAuthorizedBorrowers_Response_Borrower) GetSteamid() uint64 {
	if m != nil && m.Steamid != nil {
		return *m.Steamid
	}
	return 0
}

func (m *CDeviceAuth_GetAuthorizedBorrowers_Response_Borrower) GetIsPending() bool {
	if m != nil && m.IsPending != nil {
		return *m.IsPending
	}
	return false
}

func (m *CDeviceAuth_GetAuthorizedBorrowers_Response_Borrower) GetIsCanceled() bool {
	if m != nil && m.IsCanceled != nil {
		return *m.IsCanceled
	}
	return false
}

func (m *CDeviceAuth_GetAuthorizedBorrowers_Response_Borrower) GetTimeCreated() uint32 {
	if m != nil && m.TimeCreated != nil {
		return *m.TimeCreated
	}
	return 0
}

type CDeviceAuth_AddAuthorizedBorrowers_Request struct {
	Steamid          *uint64  `protobuf:"fixed64,1,opt,name=steamid" json:"steamid,omitempty"`
	SteamidBorrower  []uint64 `protobuf:"fixed64,2,rep,name=steamid_borrower" json:"steamid_borrower,omitempty"`
	XXX_unrecognized []byte   `json:"-"`
}

func (m *CDeviceAuth_AddAuthorizedBorrowers_Request) Reset() {
	*m = CDeviceAuth_AddAuthorizedBorrowers_Request{}
}
func (m *CDeviceAuth_AddAuthorizedBorrowers_Request) String() string {
	return proto.CompactTextString(m)
}
func (*CDeviceAuth_AddAuthorizedBorrowers_Request) ProtoMessage() {}

func (m *CDeviceAuth_AddAuthorizedBorrowers_Request) GetSteamid() uint64 {
	if m != nil && m.Steamid != nil {
		return *m.Steamid
	}
	return 0
}

func (m *CDeviceAuth_AddAuthorizedBorrowers_Request) GetSteamidBorrower() []uint64 {
	if m != nil {
		return m.SteamidBorrower
	}
	return nil
}

type CDeviceAuth_AddAuthorizedBorrowers_Response struct {
	SecondsToWait    *int32 `protobuf:"varint,1,opt,name=seconds_to_wait" json:"seconds_to_wait,omitempty"`
	XXX_unrecognized []byte `json:"-"`
}

func (m *CDeviceAuth_AddAuthorizedBorrowers_Response) Reset() {
	*m = CDeviceAuth_AddAuthorizedBorrowers_Response{}
}
func (m *CDeviceAuth_AddAuthorizedBorrowers_Response) String() string {
	return proto.CompactTextString(m)
}
func (*CDeviceAuth_AddAuthorizedBorrowers_Response) ProtoMessage() {}

func (m *CDeviceAuth_AddAuthorizedBorrowers_Response) GetSecondsToWait() int32 {
	if m != nil && m.SecondsToWait != nil {
		return *m.SecondsToWait
	}
	return 0
}

type CDeviceAuth_RemoveAuthorizedBorrowers_Request struct {
	Steamid          *uint64  `protobuf:"fixed64,1,opt,name=steamid" json:"steamid,omitempty"`
	SteamidBorrower  []uint64 `protobuf:"fixed64,2,rep,name=steamid_borrower" json:"steamid_borrower,omitempty"`
	XXX_unrecognized []byte   `json:"-"`
}

func (m *CDeviceAuth_RemoveAuthorizedBorrowers_Request) Reset() {
	*m = CDeviceAuth_RemoveAuthorizedBorrowers_Request{}
}
func (m *CDeviceAuth_RemoveAuthorizedBorrowers_Request) String() string {
	return proto.CompactTextString(m)
}
func (*CDeviceAuth_RemoveAuthorizedBorrowers_Request) ProtoMessage() {}

func (m *CDeviceAuth_RemoveAuthorizedBorrowers_Request) GetSteamid() uint64 {
	if m != nil && m.Steamid != nil {
		return *m.Steamid
	}
	return 0
}

func (m *CDeviceAuth_RemoveAuthorizedBorrowers_Request) GetSteamidBorrower() []uint64 {
	if m != nil {
		return m.SteamidBorrower
	}
	return nil
}

type CDeviceAuth_RemoveAuthorizedBorrowers_Response struct {
	XXX_unrecognized []byte `json:"-"`
}

func (m *CDeviceAuth_RemoveAuthorizedBorrowers_Response) Reset() {
	*m = CDeviceAuth_RemoveAuthorizedBorrowers_Response{}
}
func (m *CDeviceAuth_RemoveAuthorizedBorrowers_Response) String() string {
	return proto.CompactTextString(m)
}
func (*CDeviceAuth_RemoveAuthorizedBorrowers_Response) ProtoMessage() {}

func init() {
}
