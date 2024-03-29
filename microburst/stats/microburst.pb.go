// ----------------------------------------------------------------------------
// microburst.proto - microburst stats protobuf definitions
//
// March 2023
//
// Copyright (c) 2023 by Cisco Systems, Inc.
//
// Licensed under the Apache License, Version 2.0 (the "License");
// you may not use this file except in compliance with the License.
// You may obtain a copy of the License at
//
//     http://www.apache.org/licenses/LICENSE-2.0
//
// Unless required by applicable law or agreed to in writing, software
// distributed under the License is distributed on an "AS IS" BASIS,
// WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND, either express or implied.
// See the License for the specific language governing permissions and
// limitations under the License.
// ----------------------------------------------------------------------------

// Code generated by protoc-gen-go. DO NOT EDIT.
// versions:
// 	protoc-gen-go v1.31.0
// 	protoc        v3.15.8
// source: microburst.proto

package stats

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

type Microburst struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	// interface name
	InterfaceName string `protobuf:"bytes,1,opt,name=interfaceName,proto3" json:"interfaceName,omitempty"`
	// interface port queue-<idx>
	Queue string `protobuf:"bytes,2,opt,name=queue,proto3" json:"queue,omitempty"`
	// unicast (or) multicast
	QueueType string `protobuf:"bytes,3,opt,name=queueType,proto3" json:"queueType,omitempty"`
	// not used
	Threshold int32 `protobuf:"varint,4,opt,name=threshold,proto3" json:"threshold,omitempty"`
	// set to peak occupancy during burst
	Peak int32 `protobuf:"varint,5,opt,name=peak,proto3" json:"peak,omitempty"`
	// not used
	EndDepth int32 `protobuf:"varint,6,opt,name=endDepth,proto3" json:"endDepth,omitempty"`
	// set to burst duration
	Duration int64 `protobuf:"varint,7,opt,name=duration,proto3" json:"duration,omitempty"`
	// set to encoding timestamp
	Ts string `protobuf:"bytes,8,opt,name=ts,proto3" json:"ts,omitempty"`
	// set to up timestamp
	StartTs string `protobuf:"bytes,9,opt,name=startTs,proto3" json:"startTs,omitempty"`
	// set to down timestamp
	EndTs string `protobuf:"bytes,10,opt,name=endTs,proto3" json:"endTs,omitempty"`
	// set to burst peak timestamp
	PeakTs string `protobuf:"bytes,11,opt,name=peakTs,proto3" json:"peakTs,omitempty"`
	// set to node-<node_name>/microburst/interface-<port_name>/queue-[queue-<queue_idx>]
	SourceName string `protobuf:"bytes,12,opt,name=sourceName,proto3" json:"sourceName,omitempty"`
	// set to class-level<level>
	ClassLevel string `protobuf:"bytes,13,opt,name=classLevel,proto3" json:"classLevel,omitempty"`
}

func (x *Microburst) Reset() {
	*x = Microburst{}
	if protoimpl.UnsafeEnabled {
		mi := &file_microburst_proto_msgTypes[0]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *Microburst) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*Microburst) ProtoMessage() {}

func (x *Microburst) ProtoReflect() protoreflect.Message {
	mi := &file_microburst_proto_msgTypes[0]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use Microburst.ProtoReflect.Descriptor instead.
func (*Microburst) Descriptor() ([]byte, []int) {
	return file_microburst_proto_rawDescGZIP(), []int{0}
}

func (x *Microburst) GetInterfaceName() string {
	if x != nil {
		return x.InterfaceName
	}
	return ""
}

func (x *Microburst) GetQueue() string {
	if x != nil {
		return x.Queue
	}
	return ""
}

func (x *Microburst) GetQueueType() string {
	if x != nil {
		return x.QueueType
	}
	return ""
}

func (x *Microburst) GetThreshold() int32 {
	if x != nil {
		return x.Threshold
	}
	return 0
}

func (x *Microburst) GetPeak() int32 {
	if x != nil {
		return x.Peak
	}
	return 0
}

func (x *Microburst) GetEndDepth() int32 {
	if x != nil {
		return x.EndDepth
	}
	return 0
}

func (x *Microburst) GetDuration() int64 {
	if x != nil {
		return x.Duration
	}
	return 0
}

func (x *Microburst) GetTs() string {
	if x != nil {
		return x.Ts
	}
	return ""
}

func (x *Microburst) GetStartTs() string {
	if x != nil {
		return x.StartTs
	}
	return ""
}

func (x *Microburst) GetEndTs() string {
	if x != nil {
		return x.EndTs
	}
	return ""
}

func (x *Microburst) GetPeakTs() string {
	if x != nil {
		return x.PeakTs
	}
	return ""
}

func (x *Microburst) GetSourceName() string {
	if x != nil {
		return x.SourceName
	}
	return ""
}

func (x *Microburst) GetClassLevel() string {
	if x != nil {
		return x.ClassLevel
	}
	return ""
}

type Counter struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Name             string  `protobuf:"bytes,1,opt,name=name,proto3" json:"name,omitempty"`
	CurrentVal       int64   `protobuf:"varint,2,opt,name=currentVal,proto3" json:"currentVal,omitempty"`
	Min              int64   `protobuf:"varint,3,opt,name=min,proto3" json:"min,omitempty"`
	Max              int64   `protobuf:"varint,4,opt,name=max,proto3" json:"max,omitempty"`
	Avg              float64 `protobuf:"fixed64,5,opt,name=avg,proto3" json:"avg,omitempty"`
	Interval         int32   `protobuf:"varint,6,opt,name=interval,proto3" json:"interval,omitempty"`
	CumVal           int64   `protobuf:"varint,7,opt,name=cumVal,proto3" json:"cumVal,omitempty"`
	CurrentValDouble float64 `protobuf:"fixed64,8,opt,name=currentValDouble,proto3" json:"currentValDouble,omitempty"`
	BaseVal          int64   `protobuf:"varint,9,opt,name=baseVal,proto3" json:"baseVal,omitempty"`
	CurrentValU64    uint64  `protobuf:"varint,10,opt,name=currentValU64,proto3" json:"currentValU64,omitempty"`
}

func (x *Counter) Reset() {
	*x = Counter{}
	if protoimpl.UnsafeEnabled {
		mi := &file_microburst_proto_msgTypes[1]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *Counter) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*Counter) ProtoMessage() {}

func (x *Counter) ProtoReflect() protoreflect.Message {
	mi := &file_microburst_proto_msgTypes[1]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use Counter.ProtoReflect.Descriptor instead.
func (*Counter) Descriptor() ([]byte, []int) {
	return file_microburst_proto_rawDescGZIP(), []int{1}
}

func (x *Counter) GetName() string {
	if x != nil {
		return x.Name
	}
	return ""
}

func (x *Counter) GetCurrentVal() int64 {
	if x != nil {
		return x.CurrentVal
	}
	return 0
}

func (x *Counter) GetMin() int64 {
	if x != nil {
		return x.Min
	}
	return 0
}

func (x *Counter) GetMax() int64 {
	if x != nil {
		return x.Max
	}
	return 0
}

func (x *Counter) GetAvg() float64 {
	if x != nil {
		return x.Avg
	}
	return 0
}

func (x *Counter) GetInterval() int32 {
	if x != nil {
		return x.Interval
	}
	return 0
}

func (x *Counter) GetCumVal() int64 {
	if x != nil {
		return x.CumVal
	}
	return 0
}

func (x *Counter) GetCurrentValDouble() float64 {
	if x != nil {
		return x.CurrentValDouble
	}
	return 0
}

func (x *Counter) GetBaseVal() int64 {
	if x != nil {
		return x.BaseVal
	}
	return 0
}

func (x *Counter) GetCurrentValU64() uint64 {
	if x != nil {
		return x.CurrentValU64
	}
	return 0
}

type Stats struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	// not set
	Source string `protobuf:"bytes,1,opt,name=source,proto3" json:"source,omitempty"`
	// set to microburst
	StatName string `protobuf:"bytes,2,opt,name=statName,proto3" json:"statName,omitempty"`
	// not used
	Counters []*Counter `protobuf:"bytes,3,rep,name=counters,proto3" json:"counters,omitempty"`
	// microburst stats
	Microburst []*Microburst `protobuf:"bytes,4,rep,name=microburst,proto3" json:"microburst,omitempty"`
}

func (x *Stats) Reset() {
	*x = Stats{}
	if protoimpl.UnsafeEnabled {
		mi := &file_microburst_proto_msgTypes[2]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *Stats) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*Stats) ProtoMessage() {}

func (x *Stats) ProtoReflect() protoreflect.Message {
	mi := &file_microburst_proto_msgTypes[2]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use Stats.ProtoReflect.Descriptor instead.
func (*Stats) Descriptor() ([]byte, []int) {
	return file_microburst_proto_rawDescGZIP(), []int{2}
}

func (x *Stats) GetSource() string {
	if x != nil {
		return x.Source
	}
	return ""
}

func (x *Stats) GetStatName() string {
	if x != nil {
		return x.StatName
	}
	return ""
}

func (x *Stats) GetCounters() []*Counter {
	if x != nil {
		return x.Counters
	}
	return nil
}

func (x *Stats) GetMicroburst() []*Microburst {
	if x != nil {
		return x.Microburst
	}
	return nil
}

type StatsStream struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	// nodeName from switch telemetry configuration
	// # "use-nodeid <fabric_name/node_name>"
	NodeName string `protobuf:"bytes,1,opt,name=nodeName,proto3" json:"nodeName,omitempty"`
	// encode timestamp
	Timestamp string `protobuf:"bytes,2,opt,name=timestamp,proto3" json:"timestamp,omitempty"`
	// microburst stats
	StatObjects []*Stats `protobuf:"bytes,3,rep,name=statObjects,proto3" json:"statObjects,omitempty"`
	// fabricName is extracted from switch telemetry configuration
	// # "use-nodeid <fabric_name/node_name>"
	FabricName string `protobuf:"bytes,4,opt,name=fabricName,proto3" json:"fabricName,omitempty"`
	// set to CISCO_NX-OS
	Vendor string `protobuf:"bytes,5,opt,name=vendor,proto3" json:"vendor,omitempty"`
	// set to 2
	Version int32 `protobuf:"varint,6,opt,name=version,proto3" json:"version,omitempty"`
}

func (x *StatsStream) Reset() {
	*x = StatsStream{}
	if protoimpl.UnsafeEnabled {
		mi := &file_microburst_proto_msgTypes[3]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *StatsStream) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*StatsStream) ProtoMessage() {}

func (x *StatsStream) ProtoReflect() protoreflect.Message {
	mi := &file_microburst_proto_msgTypes[3]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use StatsStream.ProtoReflect.Descriptor instead.
func (*StatsStream) Descriptor() ([]byte, []int) {
	return file_microburst_proto_rawDescGZIP(), []int{3}
}

func (x *StatsStream) GetNodeName() string {
	if x != nil {
		return x.NodeName
	}
	return ""
}

func (x *StatsStream) GetTimestamp() string {
	if x != nil {
		return x.Timestamp
	}
	return ""
}

func (x *StatsStream) GetStatObjects() []*Stats {
	if x != nil {
		return x.StatObjects
	}
	return nil
}

func (x *StatsStream) GetFabricName() string {
	if x != nil {
		return x.FabricName
	}
	return ""
}

func (x *StatsStream) GetVendor() string {
	if x != nil {
		return x.Vendor
	}
	return ""
}

func (x *StatsStream) GetVersion() int32 {
	if x != nil {
		return x.Version
	}
	return 0
}

var File_microburst_proto protoreflect.FileDescriptor

var file_microburst_proto_rawDesc = []byte{
	0x0a, 0x10, 0x6d, 0x69, 0x63, 0x72, 0x6f, 0x62, 0x75, 0x72, 0x73, 0x74, 0x2e, 0x70, 0x72, 0x6f,
	0x74, 0x6f, 0x12, 0x05, 0x73, 0x74, 0x61, 0x74, 0x73, 0x22, 0xe8, 0x02, 0x0a, 0x0a, 0x4d, 0x69,
	0x63, 0x72, 0x6f, 0x62, 0x75, 0x72, 0x73, 0x74, 0x12, 0x24, 0x0a, 0x0d, 0x69, 0x6e, 0x74, 0x65,
	0x72, 0x66, 0x61, 0x63, 0x65, 0x4e, 0x61, 0x6d, 0x65, 0x18, 0x01, 0x20, 0x01, 0x28, 0x09, 0x52,
	0x0d, 0x69, 0x6e, 0x74, 0x65, 0x72, 0x66, 0x61, 0x63, 0x65, 0x4e, 0x61, 0x6d, 0x65, 0x12, 0x14,
	0x0a, 0x05, 0x71, 0x75, 0x65, 0x75, 0x65, 0x18, 0x02, 0x20, 0x01, 0x28, 0x09, 0x52, 0x05, 0x71,
	0x75, 0x65, 0x75, 0x65, 0x12, 0x1c, 0x0a, 0x09, 0x71, 0x75, 0x65, 0x75, 0x65, 0x54, 0x79, 0x70,
	0x65, 0x18, 0x03, 0x20, 0x01, 0x28, 0x09, 0x52, 0x09, 0x71, 0x75, 0x65, 0x75, 0x65, 0x54, 0x79,
	0x70, 0x65, 0x12, 0x1c, 0x0a, 0x09, 0x74, 0x68, 0x72, 0x65, 0x73, 0x68, 0x6f, 0x6c, 0x64, 0x18,
	0x04, 0x20, 0x01, 0x28, 0x05, 0x52, 0x09, 0x74, 0x68, 0x72, 0x65, 0x73, 0x68, 0x6f, 0x6c, 0x64,
	0x12, 0x12, 0x0a, 0x04, 0x70, 0x65, 0x61, 0x6b, 0x18, 0x05, 0x20, 0x01, 0x28, 0x05, 0x52, 0x04,
	0x70, 0x65, 0x61, 0x6b, 0x12, 0x1a, 0x0a, 0x08, 0x65, 0x6e, 0x64, 0x44, 0x65, 0x70, 0x74, 0x68,
	0x18, 0x06, 0x20, 0x01, 0x28, 0x05, 0x52, 0x08, 0x65, 0x6e, 0x64, 0x44, 0x65, 0x70, 0x74, 0x68,
	0x12, 0x1a, 0x0a, 0x08, 0x64, 0x75, 0x72, 0x61, 0x74, 0x69, 0x6f, 0x6e, 0x18, 0x07, 0x20, 0x01,
	0x28, 0x03, 0x52, 0x08, 0x64, 0x75, 0x72, 0x61, 0x74, 0x69, 0x6f, 0x6e, 0x12, 0x0e, 0x0a, 0x02,
	0x74, 0x73, 0x18, 0x08, 0x20, 0x01, 0x28, 0x09, 0x52, 0x02, 0x74, 0x73, 0x12, 0x18, 0x0a, 0x07,
	0x73, 0x74, 0x61, 0x72, 0x74, 0x54, 0x73, 0x18, 0x09, 0x20, 0x01, 0x28, 0x09, 0x52, 0x07, 0x73,
	0x74, 0x61, 0x72, 0x74, 0x54, 0x73, 0x12, 0x14, 0x0a, 0x05, 0x65, 0x6e, 0x64, 0x54, 0x73, 0x18,
	0x0a, 0x20, 0x01, 0x28, 0x09, 0x52, 0x05, 0x65, 0x6e, 0x64, 0x54, 0x73, 0x12, 0x16, 0x0a, 0x06,
	0x70, 0x65, 0x61, 0x6b, 0x54, 0x73, 0x18, 0x0b, 0x20, 0x01, 0x28, 0x09, 0x52, 0x06, 0x70, 0x65,
	0x61, 0x6b, 0x54, 0x73, 0x12, 0x1e, 0x0a, 0x0a, 0x73, 0x6f, 0x75, 0x72, 0x63, 0x65, 0x4e, 0x61,
	0x6d, 0x65, 0x18, 0x0c, 0x20, 0x01, 0x28, 0x09, 0x52, 0x0a, 0x73, 0x6f, 0x75, 0x72, 0x63, 0x65,
	0x4e, 0x61, 0x6d, 0x65, 0x12, 0x1e, 0x0a, 0x0a, 0x63, 0x6c, 0x61, 0x73, 0x73, 0x4c, 0x65, 0x76,
	0x65, 0x6c, 0x18, 0x0d, 0x20, 0x01, 0x28, 0x09, 0x52, 0x0a, 0x63, 0x6c, 0x61, 0x73, 0x73, 0x4c,
	0x65, 0x76, 0x65, 0x6c, 0x22, 0x93, 0x02, 0x0a, 0x07, 0x43, 0x6f, 0x75, 0x6e, 0x74, 0x65, 0x72,
	0x12, 0x12, 0x0a, 0x04, 0x6e, 0x61, 0x6d, 0x65, 0x18, 0x01, 0x20, 0x01, 0x28, 0x09, 0x52, 0x04,
	0x6e, 0x61, 0x6d, 0x65, 0x12, 0x1e, 0x0a, 0x0a, 0x63, 0x75, 0x72, 0x72, 0x65, 0x6e, 0x74, 0x56,
	0x61, 0x6c, 0x18, 0x02, 0x20, 0x01, 0x28, 0x03, 0x52, 0x0a, 0x63, 0x75, 0x72, 0x72, 0x65, 0x6e,
	0x74, 0x56, 0x61, 0x6c, 0x12, 0x10, 0x0a, 0x03, 0x6d, 0x69, 0x6e, 0x18, 0x03, 0x20, 0x01, 0x28,
	0x03, 0x52, 0x03, 0x6d, 0x69, 0x6e, 0x12, 0x10, 0x0a, 0x03, 0x6d, 0x61, 0x78, 0x18, 0x04, 0x20,
	0x01, 0x28, 0x03, 0x52, 0x03, 0x6d, 0x61, 0x78, 0x12, 0x10, 0x0a, 0x03, 0x61, 0x76, 0x67, 0x18,
	0x05, 0x20, 0x01, 0x28, 0x01, 0x52, 0x03, 0x61, 0x76, 0x67, 0x12, 0x1a, 0x0a, 0x08, 0x69, 0x6e,
	0x74, 0x65, 0x72, 0x76, 0x61, 0x6c, 0x18, 0x06, 0x20, 0x01, 0x28, 0x05, 0x52, 0x08, 0x69, 0x6e,
	0x74, 0x65, 0x72, 0x76, 0x61, 0x6c, 0x12, 0x16, 0x0a, 0x06, 0x63, 0x75, 0x6d, 0x56, 0x61, 0x6c,
	0x18, 0x07, 0x20, 0x01, 0x28, 0x03, 0x52, 0x06, 0x63, 0x75, 0x6d, 0x56, 0x61, 0x6c, 0x12, 0x2a,
	0x0a, 0x10, 0x63, 0x75, 0x72, 0x72, 0x65, 0x6e, 0x74, 0x56, 0x61, 0x6c, 0x44, 0x6f, 0x75, 0x62,
	0x6c, 0x65, 0x18, 0x08, 0x20, 0x01, 0x28, 0x01, 0x52, 0x10, 0x63, 0x75, 0x72, 0x72, 0x65, 0x6e,
	0x74, 0x56, 0x61, 0x6c, 0x44, 0x6f, 0x75, 0x62, 0x6c, 0x65, 0x12, 0x18, 0x0a, 0x07, 0x62, 0x61,
	0x73, 0x65, 0x56, 0x61, 0x6c, 0x18, 0x09, 0x20, 0x01, 0x28, 0x03, 0x52, 0x07, 0x62, 0x61, 0x73,
	0x65, 0x56, 0x61, 0x6c, 0x12, 0x24, 0x0a, 0x0d, 0x63, 0x75, 0x72, 0x72, 0x65, 0x6e, 0x74, 0x56,
	0x61, 0x6c, 0x55, 0x36, 0x34, 0x18, 0x0a, 0x20, 0x01, 0x28, 0x04, 0x52, 0x0d, 0x63, 0x75, 0x72,
	0x72, 0x65, 0x6e, 0x74, 0x56, 0x61, 0x6c, 0x55, 0x36, 0x34, 0x22, 0x9a, 0x01, 0x0a, 0x05, 0x53,
	0x74, 0x61, 0x74, 0x73, 0x12, 0x16, 0x0a, 0x06, 0x73, 0x6f, 0x75, 0x72, 0x63, 0x65, 0x18, 0x01,
	0x20, 0x01, 0x28, 0x09, 0x52, 0x06, 0x73, 0x6f, 0x75, 0x72, 0x63, 0x65, 0x12, 0x1a, 0x0a, 0x08,
	0x73, 0x74, 0x61, 0x74, 0x4e, 0x61, 0x6d, 0x65, 0x18, 0x02, 0x20, 0x01, 0x28, 0x09, 0x52, 0x08,
	0x73, 0x74, 0x61, 0x74, 0x4e, 0x61, 0x6d, 0x65, 0x12, 0x2a, 0x0a, 0x08, 0x63, 0x6f, 0x75, 0x6e,
	0x74, 0x65, 0x72, 0x73, 0x18, 0x03, 0x20, 0x03, 0x28, 0x0b, 0x32, 0x0e, 0x2e, 0x73, 0x74, 0x61,
	0x74, 0x73, 0x2e, 0x43, 0x6f, 0x75, 0x6e, 0x74, 0x65, 0x72, 0x52, 0x08, 0x63, 0x6f, 0x75, 0x6e,
	0x74, 0x65, 0x72, 0x73, 0x12, 0x31, 0x0a, 0x0a, 0x6d, 0x69, 0x63, 0x72, 0x6f, 0x62, 0x75, 0x72,
	0x73, 0x74, 0x18, 0x04, 0x20, 0x03, 0x28, 0x0b, 0x32, 0x11, 0x2e, 0x73, 0x74, 0x61, 0x74, 0x73,
	0x2e, 0x4d, 0x69, 0x63, 0x72, 0x6f, 0x62, 0x75, 0x72, 0x73, 0x74, 0x52, 0x0a, 0x6d, 0x69, 0x63,
	0x72, 0x6f, 0x62, 0x75, 0x72, 0x73, 0x74, 0x22, 0xc9, 0x01, 0x0a, 0x0b, 0x53, 0x74, 0x61, 0x74,
	0x73, 0x53, 0x74, 0x72, 0x65, 0x61, 0x6d, 0x12, 0x1a, 0x0a, 0x08, 0x6e, 0x6f, 0x64, 0x65, 0x4e,
	0x61, 0x6d, 0x65, 0x18, 0x01, 0x20, 0x01, 0x28, 0x09, 0x52, 0x08, 0x6e, 0x6f, 0x64, 0x65, 0x4e,
	0x61, 0x6d, 0x65, 0x12, 0x1c, 0x0a, 0x09, 0x74, 0x69, 0x6d, 0x65, 0x73, 0x74, 0x61, 0x6d, 0x70,
	0x18, 0x02, 0x20, 0x01, 0x28, 0x09, 0x52, 0x09, 0x74, 0x69, 0x6d, 0x65, 0x73, 0x74, 0x61, 0x6d,
	0x70, 0x12, 0x2e, 0x0a, 0x0b, 0x73, 0x74, 0x61, 0x74, 0x4f, 0x62, 0x6a, 0x65, 0x63, 0x74, 0x73,
	0x18, 0x03, 0x20, 0x03, 0x28, 0x0b, 0x32, 0x0c, 0x2e, 0x73, 0x74, 0x61, 0x74, 0x73, 0x2e, 0x53,
	0x74, 0x61, 0x74, 0x73, 0x52, 0x0b, 0x73, 0x74, 0x61, 0x74, 0x4f, 0x62, 0x6a, 0x65, 0x63, 0x74,
	0x73, 0x12, 0x1e, 0x0a, 0x0a, 0x66, 0x61, 0x62, 0x72, 0x69, 0x63, 0x4e, 0x61, 0x6d, 0x65, 0x18,
	0x04, 0x20, 0x01, 0x28, 0x09, 0x52, 0x0a, 0x66, 0x61, 0x62, 0x72, 0x69, 0x63, 0x4e, 0x61, 0x6d,
	0x65, 0x12, 0x16, 0x0a, 0x06, 0x76, 0x65, 0x6e, 0x64, 0x6f, 0x72, 0x18, 0x05, 0x20, 0x01, 0x28,
	0x09, 0x52, 0x06, 0x76, 0x65, 0x6e, 0x64, 0x6f, 0x72, 0x12, 0x18, 0x0a, 0x07, 0x76, 0x65, 0x72,
	0x73, 0x69, 0x6f, 0x6e, 0x18, 0x06, 0x20, 0x01, 0x28, 0x05, 0x52, 0x07, 0x76, 0x65, 0x72, 0x73,
	0x69, 0x6f, 0x6e, 0x42, 0x12, 0x5a, 0x10, 0x6d, 0x69, 0x63, 0x72, 0x6f, 0x62, 0x75, 0x72, 0x73,
	0x74, 0x2f, 0x73, 0x74, 0x61, 0x74, 0x73, 0x62, 0x06, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x33,
}

var (
	file_microburst_proto_rawDescOnce sync.Once
	file_microburst_proto_rawDescData = file_microburst_proto_rawDesc
)

func file_microburst_proto_rawDescGZIP() []byte {
	file_microburst_proto_rawDescOnce.Do(func() {
		file_microburst_proto_rawDescData = protoimpl.X.CompressGZIP(file_microburst_proto_rawDescData)
	})
	return file_microburst_proto_rawDescData
}

var file_microburst_proto_msgTypes = make([]protoimpl.MessageInfo, 4)
var file_microburst_proto_goTypes = []interface{}{
	(*Microburst)(nil),  // 0: stats.Microburst
	(*Counter)(nil),     // 1: stats.Counter
	(*Stats)(nil),       // 2: stats.Stats
	(*StatsStream)(nil), // 3: stats.StatsStream
}
var file_microburst_proto_depIdxs = []int32{
	1, // 0: stats.Stats.counters:type_name -> stats.Counter
	0, // 1: stats.Stats.microburst:type_name -> stats.Microburst
	2, // 2: stats.StatsStream.statObjects:type_name -> stats.Stats
	3, // [3:3] is the sub-list for method output_type
	3, // [3:3] is the sub-list for method input_type
	3, // [3:3] is the sub-list for extension type_name
	3, // [3:3] is the sub-list for extension extendee
	0, // [0:3] is the sub-list for field type_name
}

func init() { file_microburst_proto_init() }
func file_microburst_proto_init() {
	if File_microburst_proto != nil {
		return
	}
	if !protoimpl.UnsafeEnabled {
		file_microburst_proto_msgTypes[0].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*Microburst); i {
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
		file_microburst_proto_msgTypes[1].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*Counter); i {
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
		file_microburst_proto_msgTypes[2].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*Stats); i {
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
		file_microburst_proto_msgTypes[3].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*StatsStream); i {
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
			RawDescriptor: file_microburst_proto_rawDesc,
			NumEnums:      0,
			NumMessages:   4,
			NumExtensions: 0,
			NumServices:   0,
		},
		GoTypes:           file_microburst_proto_goTypes,
		DependencyIndexes: file_microburst_proto_depIdxs,
		MessageInfos:      file_microburst_proto_msgTypes,
	}.Build()
	File_microburst_proto = out.File
	file_microburst_proto_rawDesc = nil
	file_microburst_proto_goTypes = nil
	file_microburst_proto_depIdxs = nil
}
