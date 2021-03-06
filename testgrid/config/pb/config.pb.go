/*
Copyright 2016 The Kubernetes Authors.

Licensed under the Apache License, Version 2.0 (the "License");
you may not use this file except in compliance with the License.
You may obtain a copy of the License at

    http://www.apache.org/licenses/LICENSE-2.0

Unless required by applicable law or agreed to in writing, software
distributed under the License is distributed on an "AS IS" BASIS,
WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND, either express or implied.
See the License for the specific language governing permissions and
limitations under the License.
*/

// Code generated by protoc-gen-go.
// source: config.proto
// DO NOT EDIT!

/*
Package config is a generated protocol buffer package.

It is generated from these files:
	config.proto

It has these top-level messages:
	TestGroup
	Dashboard
	LinkTemplate
	LinkOptionsTemplate
	DashboardTab
	Configuration
	DefaultConfiguration
*/
package config

import proto "github.com/golang/protobuf/proto"
import fmt "fmt"
import math "math"

// Reference imports to suppress errors if they are not otherwise used.
var _ = proto.Marshal
var _ = fmt.Errorf
var _ = math.Inf

// This is a compile-time assertion to ensure that this generated file
// is compatible with the proto package it is being compiled against.
// A compilation error at this line likely means your copy of the
// proto package needs to be updated.
const _ = proto.ProtoPackageIsVersion2 // please upgrade the proto package

type TestGroup_TestsName int32

const (
	TestGroup_TESTS_NAME_MIN     TestGroup_TestsName = 0
	TestGroup_TESTS_NAME_IGNORE  TestGroup_TestsName = 1
	TestGroup_TESTS_NAME_REPLACE TestGroup_TestsName = 2
	TestGroup_TESTS_NAME_APPEND  TestGroup_TestsName = 3
)

var TestGroup_TestsName_name = map[int32]string{
	0: "TESTS_NAME_MIN",
	1: "TESTS_NAME_IGNORE",
	2: "TESTS_NAME_REPLACE",
	3: "TESTS_NAME_APPEND",
}
var TestGroup_TestsName_value = map[string]int32{
	"TESTS_NAME_MIN":     0,
	"TESTS_NAME_IGNORE":  1,
	"TESTS_NAME_REPLACE": 2,
	"TESTS_NAME_APPEND":  3,
}

func (x TestGroup_TestsName) String() string {
	return proto.EnumName(TestGroup_TestsName_name, int32(x))
}
func (TestGroup_TestsName) EnumDescriptor() ([]byte, []int) { return fileDescriptor0, []int{0, 0} }

// Specifies a group of tests to gather.
type TestGroup struct {
	// Name of this TestGroup, for mapping dashboard tabs to tests.
	Name string `protobuf:"bytes,1,opt,name=name" yaml:"name,omitempty"`
	// Path to the test result stored in gcs
	GcsPrefix string `protobuf:"bytes,2,opt,name=gcs_prefix,json=gcsPrefix" yaml:"gcs_prefix,omitempty"`
	// Number of days of test results to gather and serve.
	DaysOfResults int32 `protobuf:"varint,3,opt,name=days_of_results,json=daysOfResults" yaml:"days_of_results,omitempty"`
	// What to do with the 'Tests name' configuration value. It can replace the
	// name of the test, be appended to the name of the test, or ignored. If it is
	// ignored, then the name of the tests will be the build target.
	TestsNamePolicy TestGroup_TestsName       `protobuf:"varint,6,opt,name=tests_name_policy,json=testsNamePolicy,enum=TestGroup_TestsName" yaml:"tests_name_policy,omitempty"`
	ColumnHeader    []*TestGroup_ColumnHeader `protobuf:"bytes,9,rep,name=column_header,json=columnHeader" yaml:"column_header,omitempty"`
	// deprecated - always set to true
	UseKubernetesClient bool `protobuf:"varint,24,opt,name=use_kubernetes_client,json=useKubernetesClient" yaml:"use_kubernetes_client,omitempty"`
	// deprecated - always set to true
	IsExternal bool `protobuf:"varint,25,opt,name=is_external,json=isExternal" yaml:"is_external,omitempty"`
}

func (m *TestGroup) Reset()                    { *m = TestGroup{} }
func (m *TestGroup) String() string            { return proto.CompactTextString(m) }
func (*TestGroup) ProtoMessage()               {}
func (*TestGroup) Descriptor() ([]byte, []int) { return fileDescriptor0, []int{0} }

func (m *TestGroup) GetColumnHeader() []*TestGroup_ColumnHeader {
	if m != nil {
		return m.ColumnHeader
	}
	return nil
}

// Custom column headers for defining extra column-heading rows from values in
// the test result.
type TestGroup_ColumnHeader struct {
	ConfigurationValue string `protobuf:"bytes,3,opt,name=configuration_value,json=configurationValue" yaml:"configuration_value,omitempty"`
}

func (m *TestGroup_ColumnHeader) Reset()                    { *m = TestGroup_ColumnHeader{} }
func (m *TestGroup_ColumnHeader) String() string            { return proto.CompactTextString(m) }
func (*TestGroup_ColumnHeader) ProtoMessage()               {}
func (*TestGroup_ColumnHeader) Descriptor() ([]byte, []int) { return fileDescriptor0, []int{0, 0} }

// Specifies a dashboard.
type Dashboard struct {
	// A list of the tabs on the dashboard.
	DashboardTab []*DashboardTab `protobuf:"bytes,1,rep,name=dashboard_tab,json=dashboardTab" yaml:"dashboard_tab,omitempty"`
	// A name for the Dashboard.
	Name string `protobuf:"bytes,2,opt,name=name" yaml:"name,omitempty"`
}

func (m *Dashboard) Reset()                    { *m = Dashboard{} }
func (m *Dashboard) String() string            { return proto.CompactTextString(m) }
func (*Dashboard) ProtoMessage()               {}
func (*Dashboard) Descriptor() ([]byte, []int) { return fileDescriptor0, []int{1} }

func (m *Dashboard) GetDashboardTab() []*DashboardTab {
	if m != nil {
		return m.DashboardTab
	}
	return nil
}

type LinkTemplate struct {
	// The URL template.
	Url string `protobuf:"bytes,1,opt,name=url" yaml:"url,omitempty"`
	// The options templates.
	Options []*LinkOptionsTemplate `protobuf:"bytes,2,rep,name=options" yaml:"options,omitempty"`
}

func (m *LinkTemplate) Reset()                    { *m = LinkTemplate{} }
func (m *LinkTemplate) String() string            { return proto.CompactTextString(m) }
func (*LinkTemplate) ProtoMessage()               {}
func (*LinkTemplate) Descriptor() ([]byte, []int) { return fileDescriptor0, []int{2} }

func (m *LinkTemplate) GetOptions() []*LinkOptionsTemplate {
	if m != nil {
		return m.Options
	}
	return nil
}

// A simple key/value pair for link options.
type LinkOptionsTemplate struct {
	// The key for the option. This is not expanded.
	Key string `protobuf:"bytes,1,opt,name=key" yaml:"key,omitempty"`
	// The value for the option. This is expanded the same as the LinkTemplate.
	Value string `protobuf:"bytes,2,opt,name=value" yaml:"value,omitempty"`
}

func (m *LinkOptionsTemplate) Reset()                    { *m = LinkOptionsTemplate{} }
func (m *LinkOptionsTemplate) String() string            { return proto.CompactTextString(m) }
func (*LinkOptionsTemplate) ProtoMessage()               {}
func (*LinkOptionsTemplate) Descriptor() ([]byte, []int) { return fileDescriptor0, []int{3} }

// A single tab on a dashboard.
type DashboardTab struct {
	// The name of the dashboard tab to display in the client.
	Name string `protobuf:"bytes,1,opt,name=name" yaml:"name,omitempty"`
	// The name of the TestGroup specifying the test results for this tab.
	TestGroupName string `protobuf:"bytes,2,opt,name=test_group_name,json=testGroupName" yaml:"test_group_name,omitempty"`
	// Default bug component for manually filing bugs from the dashboard
	BugComponent int32 `protobuf:"varint,3,opt,name=bug_component,json=bugComponent" yaml:"bug_component,omitempty"`
	// Default code search path for changelist search links
	CodeSearchPath string `protobuf:"bytes,4,opt,name=code_search_path,json=codeSearchPath" yaml:"code_search_path,omitempty"`
	// The URL template to visit after clicking on a cell.
	OpenTestTemplate *LinkTemplate `protobuf:"bytes,7,opt,name=open_test_template,json=openTestTemplate" yaml:"open_test_template,omitempty"`
	// The URL template to visit when filing a bug.
	FileBugTemplate *LinkTemplate `protobuf:"bytes,8,opt,name=file_bug_template,json=fileBugTemplate" yaml:"file_bug_template,omitempty"`
	// The URL template to visit when attaching a bug
	AttachBugTemplate *LinkTemplate `protobuf:"bytes,9,opt,name=attach_bug_template,json=attachBugTemplate" yaml:"attach_bug_template,omitempty"`
	// Text to show in the about menu as a link to another view of the results.
	ResultsText string `protobuf:"bytes,10,opt,name=results_text,json=resultsText" yaml:"results_text,omitempty"`
	// The URL template to visit after clicking.
	ResultsUrlTemplate *LinkTemplate `protobuf:"bytes,11,opt,name=results_url_template,json=resultsUrlTemplate" yaml:"results_url_template,omitempty"`
	// The URL template to visit when searching for changelists.
	CodeSearchUrlTemplate *LinkTemplate `protobuf:"bytes,12,opt,name=code_search_url_template,json=codeSearchUrlTemplate" yaml:"code_search_url_template,omitempty"`
}

func (m *DashboardTab) Reset()                    { *m = DashboardTab{} }
func (m *DashboardTab) String() string            { return proto.CompactTextString(m) }
func (*DashboardTab) ProtoMessage()               {}
func (*DashboardTab) Descriptor() ([]byte, []int) { return fileDescriptor0, []int{4} }

func (m *DashboardTab) GetOpenTestTemplate() *LinkTemplate {
	if m != nil {
		return m.OpenTestTemplate
	}
	return nil
}

func (m *DashboardTab) GetFileBugTemplate() *LinkTemplate {
	if m != nil {
		return m.FileBugTemplate
	}
	return nil
}

func (m *DashboardTab) GetAttachBugTemplate() *LinkTemplate {
	if m != nil {
		return m.AttachBugTemplate
	}
	return nil
}

func (m *DashboardTab) GetResultsUrlTemplate() *LinkTemplate {
	if m != nil {
		return m.ResultsUrlTemplate
	}
	return nil
}

func (m *DashboardTab) GetCodeSearchUrlTemplate() *LinkTemplate {
	if m != nil {
		return m.CodeSearchUrlTemplate
	}
	return nil
}

// A service configuration consisting of multiple test groups and dashboards.
type Configuration struct {
	// A list of groups of tests to gather.
	TestGroups []*TestGroup `protobuf:"bytes,1,rep,name=test_groups,json=testGroups" yaml:"test_groups,omitempty"`
	// A list of all of the dashboards for a server.
	Dashboards []*Dashboard `protobuf:"bytes,2,rep,name=dashboards" yaml:"dashboards,omitempty"`
}

func (m *Configuration) Reset()                    { *m = Configuration{} }
func (m *Configuration) String() string            { return proto.CompactTextString(m) }
func (*Configuration) ProtoMessage()               {}
func (*Configuration) Descriptor() ([]byte, []int) { return fileDescriptor0, []int{5} }

func (m *Configuration) GetTestGroups() []*TestGroup {
	if m != nil {
		return m.TestGroups
	}
	return nil
}

func (m *Configuration) GetDashboards() []*Dashboard {
	if m != nil {
		return m.Dashboards
	}
	return nil
}

type DefaultConfiguration struct {
	// A default testgroup with default initialization data
	DefaultTestGroup *TestGroup `protobuf:"bytes,1,opt,name=default_test_group,json=defaultTestGroup" yaml:"default_test_group,omitempty"`
	// A default dashboard with default initialization data
	DefaultDashboardTab *DashboardTab `protobuf:"bytes,2,opt,name=default_dashboard_tab,json=defaultDashboardTab" yaml:"default_dashboard_tab,omitempty"`
}

func (m *DefaultConfiguration) Reset()                    { *m = DefaultConfiguration{} }
func (m *DefaultConfiguration) String() string            { return proto.CompactTextString(m) }
func (*DefaultConfiguration) ProtoMessage()               {}
func (*DefaultConfiguration) Descriptor() ([]byte, []int) { return fileDescriptor0, []int{6} }

func (m *DefaultConfiguration) GetDefaultTestGroup() *TestGroup {
	if m != nil {
		return m.DefaultTestGroup
	}
	return nil
}

func (m *DefaultConfiguration) GetDefaultDashboardTab() *DashboardTab {
	if m != nil {
		return m.DefaultDashboardTab
	}
	return nil
}

func init() {
	proto.RegisterType((*TestGroup)(nil), "TestGroup")
	proto.RegisterType((*TestGroup_ColumnHeader)(nil), "TestGroup.ColumnHeader")
	proto.RegisterType((*Dashboard)(nil), "Dashboard")
	proto.RegisterType((*LinkTemplate)(nil), "LinkTemplate")
	proto.RegisterType((*LinkOptionsTemplate)(nil), "LinkOptionsTemplate")
	proto.RegisterType((*DashboardTab)(nil), "DashboardTab")
	proto.RegisterType((*Configuration)(nil), "Configuration")
	proto.RegisterType((*DefaultConfiguration)(nil), "DefaultConfiguration")
	proto.RegisterEnum("TestGroup_TestsName", TestGroup_TestsName_name, TestGroup_TestsName_value)
}

func init() { proto.RegisterFile("config.proto", fileDescriptor0) }

var fileDescriptor0 = []byte{
	// 780 bytes of a gzipped FileDescriptorProto
	0x1f, 0x8b, 0x08, 0x00, 0x00, 0x09, 0x6e, 0x88, 0x02, 0xff, 0x6c, 0x55, 0xdf, 0xaf, 0xe2, 0x44,
	0x14, 0xb6, 0xd0, 0x7b, 0x6f, 0x7b, 0x28, 0xf7, 0x0e, 0x03, 0x68, 0x5d, 0x63, 0xc4, 0x9a, 0x6c,
	0x88, 0x26, 0x98, 0xe0, 0x8b, 0x46, 0x37, 0x8a, 0x80, 0xeb, 0x76, 0x77, 0xb9, 0xa4, 0xa0, 0xaf,
	0x93, 0xa1, 0x0c, 0xd0, 0x50, 0xda, 0xa6, 0x33, 0x35, 0xf0, 0x1f, 0xf8, 0x7e, 0xff, 0x61, 0xd3,
	0xe9, 0x4f, 0x0c, 0x6f, 0x33, 0xdf, 0xf9, 0xce, 0x39, 0x33, 0xdf, 0x7c, 0xa7, 0x05, 0xc3, 0x0d,
	0x83, 0x9d, 0xb7, 0x1f, 0x45, 0x71, 0x28, 0x42, 0xeb, 0x45, 0x05, 0x7d, 0xcd, 0xb8, 0x78, 0x1b,
	0x87, 0x49, 0x84, 0x31, 0xa8, 0x01, 0x3d, 0x31, 0x53, 0x19, 0x28, 0x43, 0xdd, 0x91, 0x6b, 0xfc,
	0x25, 0xc0, 0xde, 0xe5, 0x24, 0x8a, 0xd9, 0xce, 0x3b, 0x9b, 0x0d, 0x19, 0xd1, 0xf7, 0x2e, 0x5f,
	0x4a, 0x00, 0xbf, 0x86, 0xa7, 0x2d, 0xbd, 0x70, 0x12, 0xee, 0x48, 0xcc, 0x78, 0xe2, 0x0b, 0x6e,
	0x36, 0x07, 0xca, 0xf0, 0xce, 0x69, 0xa7, 0xf0, 0xf3, 0xce, 0xc9, 0x40, 0xfc, 0x1b, 0x74, 0x04,
	0xe3, 0x82, 0x93, 0xb4, 0x28, 0x89, 0x42, 0xdf, 0x73, 0x2f, 0xe6, 0xfd, 0x40, 0x19, 0x3e, 0x8e,
	0x7b, 0xa3, 0xf2, 0x04, 0x72, 0xc5, 0x17, 0xf4, 0xc4, 0x9c, 0x27, 0x51, 0x2c, 0x97, 0x92, 0x8c,
	0x7f, 0x81, 0xb6, 0x1b, 0xfa, 0xc9, 0x29, 0x20, 0x07, 0x46, 0xb7, 0x2c, 0x36, 0xf5, 0x41, 0x73,
	0xd8, 0x1a, 0x7f, 0x56, 0xcb, 0x9e, 0xca, 0xf8, 0x9f, 0x32, 0xec, 0x18, 0x6e, 0x6d, 0x87, 0xc7,
	0xd0, 0x4f, 0x38, 0x23, 0xc7, 0x64, 0xc3, 0xe2, 0x80, 0x09, 0xc6, 0x89, 0xeb, 0x7b, 0x2c, 0x10,
	0xa6, 0x39, 0x50, 0x86, 0x9a, 0xd3, 0x4d, 0x38, 0x7b, 0x5f, 0xc6, 0xa6, 0x32, 0x84, 0xbf, 0x82,
	0x96, 0xc7, 0x09, 0x3b, 0x0b, 0x16, 0x07, 0xd4, 0x37, 0x3f, 0x97, 0x4c, 0xf0, 0xf8, 0x3c, 0x47,
	0x5e, 0xbd, 0x07, 0xa3, 0xde, 0x12, 0x7f, 0x0f, 0xdd, 0x4c, 0xdd, 0x24, 0xa6, 0xc2, 0x0b, 0x03,
	0xf2, 0x0f, 0xf5, 0x13, 0x26, 0x05, 0xd1, 0x1d, 0x7c, 0x15, 0xfa, 0x3b, 0x8d, 0xd8, 0xaa, 0xa6,
	0xa0, 0x86, 0xad, 0x6a, 0x0d, 0xd4, 0xb4, 0x58, 0xf6, 0x12, 0xf2, 0xca, 0x18, 0xc3, 0xe3, 0x7a,
	0xbe, 0x5a, 0xaf, 0xc8, 0x62, 0xf2, 0x71, 0x4e, 0x3e, 0xbe, 0x5b, 0xa0, 0x4f, 0x70, 0x1f, 0x3a,
	0x35, 0xec, 0xdd, 0xdb, 0xc5, 0xb3, 0x33, 0x47, 0x0a, 0xfe, 0x14, 0x70, 0x0d, 0x76, 0xe6, 0xcb,
	0x0f, 0x93, 0xe9, 0x1c, 0x35, 0xfe, 0x47, 0x9f, 0x2c, 0x97, 0xf3, 0xc5, 0x0c, 0x35, 0x6d, 0x55,
	0x53, 0xd1, 0x9d, 0xad, 0x6a, 0x77, 0xe8, 0xde, 0x56, 0xb5, 0x07, 0xa4, 0xd9, 0xaa, 0xa6, 0x21,
	0xdd, 0x56, 0x35, 0x40, 0xa6, 0xad, 0x6a, 0xaf, 0xd0, 0x17, 0xd6, 0x0a, 0xf4, 0x19, 0xe5, 0x87,
	0x4d, 0x48, 0xe3, 0x2d, 0x1e, 0x43, 0x7b, 0x5b, 0x6c, 0x88, 0xa0, 0x1b, 0x53, 0x91, 0xba, 0xb7,
	0x47, 0x25, 0x65, 0x4d, 0x37, 0x8e, 0xb1, 0xad, 0xed, 0x4a, 0x23, 0x35, 0x2a, 0x23, 0x59, 0x4b,
	0x30, 0x3e, 0x78, 0xc1, 0x71, 0xcd, 0x4e, 0x91, 0x4f, 0x05, 0xc3, 0x08, 0x9a, 0x49, 0xec, 0xe7,
	0x5e, 0x4b, 0x97, 0x78, 0x04, 0x0f, 0x61, 0x94, 0x8a, 0xc3, 0xcd, 0x86, 0xec, 0xd1, 0x1b, 0xa5,
	0x19, 0xcf, 0x19, 0x56, 0x24, 0x3a, 0x05, 0xc9, 0x7a, 0x03, 0xdd, 0x1b, 0xf1, 0xb4, 0xf0, 0x91,
	0x5d, 0x8a, 0xc2, 0x47, 0x76, 0xc1, 0x3d, 0xb8, 0xcb, 0x5e, 0x22, 0x3b, 0x4f, 0xb6, 0xb1, 0xfe,
	0x55, 0xc1, 0x98, 0xdd, 0x3a, 0x75, 0xdd, 0xfe, 0xaf, 0x41, 0x1a, 0x91, 0xec, 0x53, 0x83, 0x91,
	0xda, 0xa5, 0xda, 0xa2, 0xb0, 0x9d, 0x7c, 0xb0, 0x6f, 0xa0, 0xbd, 0x49, 0xf6, 0xc4, 0x0d, 0x4f,
	0x51, 0x18, 0xa4, 0xbe, 0xca, 0xa6, 0xc0, 0xd8, 0x24, 0xfb, 0x69, 0x81, 0xe1, 0x21, 0x20, 0x37,
	0xdc, 0x32, 0xc2, 0x19, 0x8d, 0xdd, 0x03, 0x89, 0xa8, 0x38, 0x98, 0xaa, 0xac, 0xf6, 0x98, 0xe2,
	0x2b, 0x09, 0x2f, 0xa9, 0x38, 0xe0, 0x9f, 0x01, 0x87, 0x11, 0x0b, 0x88, 0xec, 0x2d, 0xf2, 0x9b,
	0x99, 0x0f, 0x03, 0x45, 0x2a, 0x5f, 0xd7, 0xd1, 0x41, 0x29, 0x31, 0x75, 0x4e, 0x29, 0xc0, 0x4f,
	0xd0, 0xd9, 0x79, 0x3e, 0x23, 0xe9, 0x81, 0xca, 0x5c, 0xed, 0x56, 0xee, 0x53, 0xca, 0xfb, 0x3d,
	0xd9, 0x97, 0xa9, 0x6f, 0xa0, 0x4b, 0x85, 0xa0, 0xee, 0xe1, 0x3a, 0x59, 0xbf, 0x95, 0xdc, 0xc9,
	0x98, 0xf5, 0xf4, 0xaf, 0xc1, 0xc8, 0xbf, 0x02, 0x44, 0xb0, 0xb3, 0x30, 0x41, 0x5e, 0xae, 0x95,
	0x63, 0x6b, 0x76, 0x16, 0xf8, 0x57, 0xe8, 0x15, 0x94, 0x24, 0xf6, 0xab, 0x16, 0xad, 0x5b, 0x2d,
	0x70, 0x4e, 0xfd, 0x2b, 0xf6, 0xcb, 0x1e, 0x7f, 0x80, 0x59, 0x17, 0xf1, 0xaa, 0x88, 0x71, 0xab,
	0x48, 0xbf, 0xd2, 0xb6, 0x56, 0xa7, 0x1c, 0x81, 0x7b, 0xf4, 0x60, 0x1d, 0xa0, 0x3d, 0xad, 0x4f,
	0x27, 0xfe, 0x0e, 0x5a, 0xd5, 0xb3, 0xf3, 0xdc, 0xf2, 0x50, 0x7d, 0x6a, 0x1c, 0x28, 0x9f, 0x9f,
	0xe3, 0x6f, 0x01, 0x4a, 0xf7, 0x17, 0xd6, 0x85, 0x6a, 0x3c, 0x9c, 0x5a, 0xd4, 0x7a, 0x51, 0xa0,
	0x37, 0x63, 0x3b, 0x9a, 0xf8, 0xe2, 0xba, 0xe3, 0x8f, 0x80, 0xb7, 0x19, 0x4e, 0xaa, 0xce, 0xd2,
	0x8a, 0xd7, 0x8d, 0x51, 0xce, 0xaa, 0xbe, 0xda, 0x13, 0xe8, 0x17, 0x99, 0xd7, 0x83, 0xda, 0xc8,
	0xd5, 0xb8, 0x1a, 0xd4, 0x6e, 0xce, 0xad, 0x83, 0x9b, 0x7b, 0xf9, 0x37, 0xf8, 0xe1, 0xbf, 0x00,
	0x00, 0x00, 0xff, 0xff, 0xf5, 0x02, 0x5c, 0xb5, 0x1d, 0x06, 0x00, 0x00,
}
