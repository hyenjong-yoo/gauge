// Copyright 2015 ThoughtWorks, Inc.

// This file is part of Gauge.

// Gauge is free software: you can redistribute it and/or modify
// it under the terms of the GNU General Public License as published by
// the Free Software Foundation, either version 3 of the License, or
// (at your option) any later version.

// Gauge is distributed in the hope that it will be useful,
// but WITHOUT ANY WARRANTY; without even the implied warranty of
// MERCHANTABILITY or FITNESS FOR A PARTICULAR PURPOSE.  See the
// GNU General Public License for more details.

// You should have received a copy of the GNU General Public License
// along with Gauge.  If not, see <http://www.gnu.org/licenses/>.

// Code generated by protoc-gen-go.
// source: messages.proto
// DO NOT EDIT!

/*
Package gauge_messages is a generated protocol buffer package.

It is generated from these files:
	messages.proto

It has these top-level messages:
	KillProcessRequest
	ExecutionStatusResponse
	ExecutionStartingRequest
	ExecutionEndingRequest
	SpecExecutionStartingRequest
	SpecExecutionEndingRequest
	ScenarioExecutionStartingRequest
	ScenarioExecutionEndingRequest
	StepExecutionStartingRequest
	StepExecutionEndingRequest
	ExecutionInfo
	SpecInfo
	ScenarioInfo
	StepInfo
	ExecuteStepRequest
	StepValidateRequest
	StepValidateResponse
	SuiteExecutionResult
	StepNamesRequest
	StepNamesResponse
	ScenarioDataStoreInitRequest
	SpecDataStoreInitRequest
	SuiteDataStoreInitRequest
	ParameterPosition
	RefactorRequest
	RefactorResponse
	StepNameRequest
	StepNameResponse
	Message
*/
package gauge_messages

import proto "github.com/golang/protobuf/proto"
import math "math"

// Reference imports to suppress errors if they are not otherwise used.
var _ = proto.Marshal
var _ = math.Inf

type Message_MessageType int32

const (
	Message_ExecutionStarting         Message_MessageType = 0
	Message_SpecExecutionStarting     Message_MessageType = 1
	Message_SpecExecutionEnding       Message_MessageType = 2
	Message_ScenarioExecutionStarting Message_MessageType = 3
	Message_ScenarioExecutionEnding   Message_MessageType = 4
	Message_StepExecutionStarting     Message_MessageType = 5
	Message_StepExecutionEnding       Message_MessageType = 6
	Message_ExecuteStep               Message_MessageType = 7
	Message_ExecutionEnding           Message_MessageType = 8
	Message_StepValidateRequest       Message_MessageType = 9
	Message_StepValidateResponse      Message_MessageType = 10
	Message_ExecutionStatusResponse   Message_MessageType = 11
	Message_StepNamesRequest          Message_MessageType = 12
	Message_StepNamesResponse         Message_MessageType = 13
	Message_KillProcessRequest        Message_MessageType = 14
	Message_SuiteExecutionResult      Message_MessageType = 15
	Message_ScenarioDataStoreInit     Message_MessageType = 16
	Message_SpecDataStoreInit         Message_MessageType = 17
	Message_SuiteDataStoreInit        Message_MessageType = 18
	Message_StepNameRequest           Message_MessageType = 19
	Message_StepNameResponse          Message_MessageType = 20
	Message_RefactorRequest           Message_MessageType = 21
	Message_RefactorResponse          Message_MessageType = 22
)

var Message_MessageType_name = map[int32]string{
	0:  "ExecutionStarting",
	1:  "SpecExecutionStarting",
	2:  "SpecExecutionEnding",
	3:  "ScenarioExecutionStarting",
	4:  "ScenarioExecutionEnding",
	5:  "StepExecutionStarting",
	6:  "StepExecutionEnding",
	7:  "ExecuteStep",
	8:  "ExecutionEnding",
	9:  "StepValidateRequest",
	10: "StepValidateResponse",
	11: "ExecutionStatusResponse",
	12: "StepNamesRequest",
	13: "StepNamesResponse",
	14: "KillProcessRequest",
	15: "SuiteExecutionResult",
	16: "ScenarioDataStoreInit",
	17: "SpecDataStoreInit",
	18: "SuiteDataStoreInit",
	19: "StepNameRequest",
	20: "StepNameResponse",
	21: "RefactorRequest",
	22: "RefactorResponse",
}
var Message_MessageType_value = map[string]int32{
	"ExecutionStarting":         0,
	"SpecExecutionStarting":     1,
	"SpecExecutionEnding":       2,
	"ScenarioExecutionStarting": 3,
	"ScenarioExecutionEnding":   4,
	"StepExecutionStarting":     5,
	"StepExecutionEnding":       6,
	"ExecuteStep":               7,
	"ExecutionEnding":           8,
	"StepValidateRequest":       9,
	"StepValidateResponse":      10,
	"ExecutionStatusResponse":   11,
	"StepNamesRequest":          12,
	"StepNamesResponse":         13,
	"KillProcessRequest":        14,
	"SuiteExecutionResult":      15,
	"ScenarioDataStoreInit":     16,
	"SpecDataStoreInit":         17,
	"SuiteDataStoreInit":        18,
	"StepNameRequest":           19,
	"StepNameResponse":          20,
	"RefactorRequest":           21,
	"RefactorResponse":          22,
}

func (x Message_MessageType) Enum() *Message_MessageType {
	p := new(Message_MessageType)
	*p = x
	return p
}
func (x Message_MessageType) String() string {
	return proto.EnumName(Message_MessageType_name, int32(x))
}
func (x *Message_MessageType) UnmarshalJSON(data []byte) error {
	value, err := proto.UnmarshalJSONEnum(Message_MessageType_value, data, "Message_MessageType")
	if err != nil {
		return err
	}
	*x = Message_MessageType(value)
	return nil
}

// / Default request. Tells the runner to shutdown.
type KillProcessRequest struct {
	XXX_unrecognized []byte `json:"-"`
}

func (m *KillProcessRequest) Reset()         { *m = KillProcessRequest{} }
func (m *KillProcessRequest) String() string { return proto.CompactTextString(m) }
func (*KillProcessRequest) ProtoMessage()    {}

// / Sends to any request which needs a execution status as response
// / usually step execution, hooks etc will return this
type ExecutionStatusResponse struct {
	ExecutionResult  *ProtoExecutionResult `protobuf:"bytes,1,req,name=executionResult" json:"executionResult,omitempty"`
	XXX_unrecognized []byte                `json:"-"`
}

func (m *ExecutionStatusResponse) Reset()         { *m = ExecutionStatusResponse{} }
func (m *ExecutionStatusResponse) String() string { return proto.CompactTextString(m) }
func (*ExecutionStatusResponse) ProtoMessage()    {}

func (m *ExecutionStatusResponse) GetExecutionResult() *ProtoExecutionResult {
	if m != nil {
		return m.ExecutionResult
	}
	return nil
}

// / Sent at start of Suite Execution. Tells the runner to execute `before_suite` hook.
type ExecutionStartingRequest struct {
	CurrentExecutionInfo *ExecutionInfo `protobuf:"bytes,1,opt,name=currentExecutionInfo" json:"currentExecutionInfo,omitempty"`
	XXX_unrecognized     []byte         `json:"-"`
}

func (m *ExecutionStartingRequest) Reset()         { *m = ExecutionStartingRequest{} }
func (m *ExecutionStartingRequest) String() string { return proto.CompactTextString(m) }
func (*ExecutionStartingRequest) ProtoMessage()    {}

func (m *ExecutionStartingRequest) GetCurrentExecutionInfo() *ExecutionInfo {
	if m != nil {
		return m.CurrentExecutionInfo
	}
	return nil
}

// / Sent at end of Suite Execution. Tells the runner to execute `after_suite` hook.
type ExecutionEndingRequest struct {
	CurrentExecutionInfo *ExecutionInfo `protobuf:"bytes,1,opt,name=currentExecutionInfo" json:"currentExecutionInfo,omitempty"`
	XXX_unrecognized     []byte         `json:"-"`
}

func (m *ExecutionEndingRequest) Reset()         { *m = ExecutionEndingRequest{} }
func (m *ExecutionEndingRequest) String() string { return proto.CompactTextString(m) }
func (*ExecutionEndingRequest) ProtoMessage()    {}

func (m *ExecutionEndingRequest) GetCurrentExecutionInfo() *ExecutionInfo {
	if m != nil {
		return m.CurrentExecutionInfo
	}
	return nil
}

// / Sent at start of Spec Execution. Tells the runner to execute `before_spec` hook.
type SpecExecutionStartingRequest struct {
	CurrentExecutionInfo *ExecutionInfo `protobuf:"bytes,1,opt,name=currentExecutionInfo" json:"currentExecutionInfo,omitempty"`
	XXX_unrecognized     []byte         `json:"-"`
}

func (m *SpecExecutionStartingRequest) Reset()         { *m = SpecExecutionStartingRequest{} }
func (m *SpecExecutionStartingRequest) String() string { return proto.CompactTextString(m) }
func (*SpecExecutionStartingRequest) ProtoMessage()    {}

func (m *SpecExecutionStartingRequest) GetCurrentExecutionInfo() *ExecutionInfo {
	if m != nil {
		return m.CurrentExecutionInfo
	}
	return nil
}

// / Sent at end of Spec Execution. Tells the runner to execute `after_spec` hook.
type SpecExecutionEndingRequest struct {
	CurrentExecutionInfo *ExecutionInfo `protobuf:"bytes,1,opt,name=currentExecutionInfo" json:"currentExecutionInfo,omitempty"`
	XXX_unrecognized     []byte         `json:"-"`
}

func (m *SpecExecutionEndingRequest) Reset()         { *m = SpecExecutionEndingRequest{} }
func (m *SpecExecutionEndingRequest) String() string { return proto.CompactTextString(m) }
func (*SpecExecutionEndingRequest) ProtoMessage()    {}

func (m *SpecExecutionEndingRequest) GetCurrentExecutionInfo() *ExecutionInfo {
	if m != nil {
		return m.CurrentExecutionInfo
	}
	return nil
}

// / Sent at start of Scenario Execution. Tells the runner to execute `before_scenario` hook.
type ScenarioExecutionStartingRequest struct {
	CurrentExecutionInfo *ExecutionInfo `protobuf:"bytes,1,opt,name=currentExecutionInfo" json:"currentExecutionInfo,omitempty"`
	XXX_unrecognized     []byte         `json:"-"`
}

func (m *ScenarioExecutionStartingRequest) Reset()         { *m = ScenarioExecutionStartingRequest{} }
func (m *ScenarioExecutionStartingRequest) String() string { return proto.CompactTextString(m) }
func (*ScenarioExecutionStartingRequest) ProtoMessage()    {}

func (m *ScenarioExecutionStartingRequest) GetCurrentExecutionInfo() *ExecutionInfo {
	if m != nil {
		return m.CurrentExecutionInfo
	}
	return nil
}

// / Sent at end of Scenario Execution. Tells the runner to execute `after_scenario` hook.
type ScenarioExecutionEndingRequest struct {
	CurrentExecutionInfo *ExecutionInfo `protobuf:"bytes,1,opt,name=currentExecutionInfo" json:"currentExecutionInfo,omitempty"`
	XXX_unrecognized     []byte         `json:"-"`
}

func (m *ScenarioExecutionEndingRequest) Reset()         { *m = ScenarioExecutionEndingRequest{} }
func (m *ScenarioExecutionEndingRequest) String() string { return proto.CompactTextString(m) }
func (*ScenarioExecutionEndingRequest) ProtoMessage()    {}

func (m *ScenarioExecutionEndingRequest) GetCurrentExecutionInfo() *ExecutionInfo {
	if m != nil {
		return m.CurrentExecutionInfo
	}
	return nil
}

// / Sent at start of Step Execution. Tells the runner to execute `before_step` hook.
type StepExecutionStartingRequest struct {
	CurrentExecutionInfo *ExecutionInfo `protobuf:"bytes,1,opt,name=currentExecutionInfo" json:"currentExecutionInfo,omitempty"`
	XXX_unrecognized     []byte         `json:"-"`
}

func (m *StepExecutionStartingRequest) Reset()         { *m = StepExecutionStartingRequest{} }
func (m *StepExecutionStartingRequest) String() string { return proto.CompactTextString(m) }
func (*StepExecutionStartingRequest) ProtoMessage()    {}

func (m *StepExecutionStartingRequest) GetCurrentExecutionInfo() *ExecutionInfo {
	if m != nil {
		return m.CurrentExecutionInfo
	}
	return nil
}

// / Sent at end of Step Execution. Tells the runner to execute `after_step` hook.
type StepExecutionEndingRequest struct {
	CurrentExecutionInfo *ExecutionInfo `protobuf:"bytes,1,opt,name=currentExecutionInfo" json:"currentExecutionInfo,omitempty"`
	XXX_unrecognized     []byte         `json:"-"`
}

func (m *StepExecutionEndingRequest) Reset()         { *m = StepExecutionEndingRequest{} }
func (m *StepExecutionEndingRequest) String() string { return proto.CompactTextString(m) }
func (*StepExecutionEndingRequest) ProtoMessage()    {}

func (m *StepExecutionEndingRequest) GetCurrentExecutionInfo() *ExecutionInfo {
	if m != nil {
		return m.CurrentExecutionInfo
	}
	return nil
}

// / Contains details of the execution.
// / Depending on the context (Step, Scenario, Spec or Suite), the respective fields are set.
type ExecutionInfo struct {
	// / Holds the information of the current Spec. Valid in context of Spec execution.
	CurrentSpec *SpecInfo `protobuf:"bytes,1,opt,name=currentSpec" json:"currentSpec,omitempty"`
	// / Holds the information of the current Scenario. Valid in context of Scenario execution.
	CurrentScenario *ScenarioInfo `protobuf:"bytes,2,opt,name=currentScenario" json:"currentScenario,omitempty"`
	// / Holds the information of the current Step. Valid in context of Step execution.
	CurrentStep *StepInfo `protobuf:"bytes,3,opt,name=currentStep" json:"currentStep,omitempty"`
	// / Stacktrace of the execution. Valid only if there is an error in execution.
	Stacktrace       *string `protobuf:"bytes,4,opt,name=stacktrace" json:"stacktrace,omitempty"`
	XXX_unrecognized []byte  `json:"-"`
}

func (m *ExecutionInfo) Reset()         { *m = ExecutionInfo{} }
func (m *ExecutionInfo) String() string { return proto.CompactTextString(m) }
func (*ExecutionInfo) ProtoMessage()    {}

func (m *ExecutionInfo) GetCurrentSpec() *SpecInfo {
	if m != nil {
		return m.CurrentSpec
	}
	return nil
}

func (m *ExecutionInfo) GetCurrentScenario() *ScenarioInfo {
	if m != nil {
		return m.CurrentScenario
	}
	return nil
}

func (m *ExecutionInfo) GetCurrentStep() *StepInfo {
	if m != nil {
		return m.CurrentStep
	}
	return nil
}

func (m *ExecutionInfo) GetStacktrace() string {
	if m != nil && m.Stacktrace != nil {
		return *m.Stacktrace
	}
	return ""
}

// / Contains details of the Spec execution.
type SpecInfo struct {
	// / Name of the current Spec being executed.
	Name *string `protobuf:"bytes,1,req,name=name" json:"name,omitempty"`
	// / Full File path containing the current Spec being executed.
	FileName *string `protobuf:"bytes,2,req,name=fileName" json:"fileName,omitempty"`
	// / Flag to indicate if the current Spec execution failed.
	IsFailed *bool `protobuf:"varint,3,req,name=isFailed" json:"isFailed,omitempty"`
	// / Tags relevant to the current Spec execution.
	Tags             []string `protobuf:"bytes,4,rep,name=tags" json:"tags,omitempty"`
	XXX_unrecognized []byte   `json:"-"`
}

func (m *SpecInfo) Reset()         { *m = SpecInfo{} }
func (m *SpecInfo) String() string { return proto.CompactTextString(m) }
func (*SpecInfo) ProtoMessage()    {}

func (m *SpecInfo) GetName() string {
	if m != nil && m.Name != nil {
		return *m.Name
	}
	return ""
}

func (m *SpecInfo) GetFileName() string {
	if m != nil && m.FileName != nil {
		return *m.FileName
	}
	return ""
}

func (m *SpecInfo) GetIsFailed() bool {
	if m != nil && m.IsFailed != nil {
		return *m.IsFailed
	}
	return false
}

func (m *SpecInfo) GetTags() []string {
	if m != nil {
		return m.Tags
	}
	return nil
}

// / Contains details of the Scenario execution.
type ScenarioInfo struct {
	// / Name of the current Scenario being executed.
	Name *string `protobuf:"bytes,1,req,name=name" json:"name,omitempty"`
	// / Flag to indicate if the current Scenario execution failed.
	IsFailed *bool `protobuf:"varint,2,req,name=isFailed" json:"isFailed,omitempty"`
	// / Tags relevant to the current Scenario execution.
	Tags             []string `protobuf:"bytes,3,rep,name=tags" json:"tags,omitempty"`
	XXX_unrecognized []byte   `json:"-"`
}

func (m *ScenarioInfo) Reset()         { *m = ScenarioInfo{} }
func (m *ScenarioInfo) String() string { return proto.CompactTextString(m) }
func (*ScenarioInfo) ProtoMessage()    {}

func (m *ScenarioInfo) GetName() string {
	if m != nil && m.Name != nil {
		return *m.Name
	}
	return ""
}

func (m *ScenarioInfo) GetIsFailed() bool {
	if m != nil && m.IsFailed != nil {
		return *m.IsFailed
	}
	return false
}

func (m *ScenarioInfo) GetTags() []string {
	if m != nil {
		return m.Tags
	}
	return nil
}

// / Contains details of the Step execution.
type StepInfo struct {
	// / The current request to execute Step
	Step *ExecuteStepRequest `protobuf:"bytes,1,req,name=step" json:"step,omitempty"`
	// / Flag to indicate if the current Step execution failed.
	IsFailed         *bool  `protobuf:"varint,2,req,name=isFailed" json:"isFailed,omitempty"`
	XXX_unrecognized []byte `json:"-"`
}

func (m *StepInfo) Reset()         { *m = StepInfo{} }
func (m *StepInfo) String() string { return proto.CompactTextString(m) }
func (*StepInfo) ProtoMessage()    {}

func (m *StepInfo) GetStep() *ExecuteStepRequest {
	if m != nil {
		return m.Step
	}
	return nil
}

func (m *StepInfo) GetIsFailed() bool {
	if m != nil && m.IsFailed != nil {
		return *m.IsFailed
	}
	return false
}

// / Request sent ot the runner to Execute a Step
type ExecuteStepRequest struct {
	// / Contains the actual text of the Step being executed.
	// / This contains the parameters as defined in the Spec.
	ActualStepText *string `protobuf:"bytes,1,req,name=actualStepText" json:"actualStepText,omitempty"`
	// / Contains the parsed text of the Step being executed.
	// / The paramters are replaced with placeholders.
	ParsedStepText *string `protobuf:"bytes,2,req,name=parsedStepText" json:"parsedStepText,omitempty"`
	// / Flag to indicate if the execution of the Scenario, containing the current Step, failed.
	ScenarioFailing *bool `protobuf:"varint,3,opt,name=scenarioFailing" json:"scenarioFailing,omitempty"`
	// / Collection of parameters applicable to the current Step.
	Parameters       []*Parameter `protobuf:"bytes,4,rep,name=parameters" json:"parameters,omitempty"`
	XXX_unrecognized []byte       `json:"-"`
}

func (m *ExecuteStepRequest) Reset()         { *m = ExecuteStepRequest{} }
func (m *ExecuteStepRequest) String() string { return proto.CompactTextString(m) }
func (*ExecuteStepRequest) ProtoMessage()    {}

func (m *ExecuteStepRequest) GetActualStepText() string {
	if m != nil && m.ActualStepText != nil {
		return *m.ActualStepText
	}
	return ""
}

func (m *ExecuteStepRequest) GetParsedStepText() string {
	if m != nil && m.ParsedStepText != nil {
		return *m.ParsedStepText
	}
	return ""
}

func (m *ExecuteStepRequest) GetScenarioFailing() bool {
	if m != nil && m.ScenarioFailing != nil {
		return *m.ScenarioFailing
	}
	return false
}

func (m *ExecuteStepRequest) GetParameters() []*Parameter {
	if m != nil {
		return m.Parameters
	}
	return nil
}

// / Request sent ot the runner to check if given Step is valid.
// / The runner should check if there is an implementation defined for the given Step Text.
type StepValidateRequest struct {
	// / The text is used to lookup Step implementation
	StepText *string `protobuf:"bytes,1,req,name=stepText" json:"stepText,omitempty"`
	// / The number of paramters in the Step
	NumberOfParameters *int32 `protobuf:"varint,2,req,name=numberOfParameters" json:"numberOfParameters,omitempty"`
	XXX_unrecognized   []byte `json:"-"`
}

func (m *StepValidateRequest) Reset()         { *m = StepValidateRequest{} }
func (m *StepValidateRequest) String() string { return proto.CompactTextString(m) }
func (*StepValidateRequest) ProtoMessage()    {}

func (m *StepValidateRequest) GetStepText() string {
	if m != nil && m.StepText != nil {
		return *m.StepText
	}
	return ""
}

func (m *StepValidateRequest) GetNumberOfParameters() int32 {
	if m != nil && m.NumberOfParameters != nil {
		return *m.NumberOfParameters
	}
	return 0
}

// / Response of StepValidateRequest.
// / The runner tells the caller if the Request was valid,
// / i.e. an implementation exists for given Step text.
// / Returns an error message if it is an error response.
type StepValidateResponse struct {
	IsValid          *bool   `protobuf:"varint,1,req,name=isValid" json:"isValid,omitempty"`
	ErrorMessage     *string `protobuf:"bytes,2,opt,name=errorMessage" json:"errorMessage,omitempty"`
	XXX_unrecognized []byte  `json:"-"`
}

func (m *StepValidateResponse) Reset()         { *m = StepValidateResponse{} }
func (m *StepValidateResponse) String() string { return proto.CompactTextString(m) }
func (*StepValidateResponse) ProtoMessage()    {}

func (m *StepValidateResponse) GetIsValid() bool {
	if m != nil && m.IsValid != nil {
		return *m.IsValid
	}
	return false
}

func (m *StepValidateResponse) GetErrorMessage() string {
	if m != nil && m.ErrorMessage != nil {
		return *m.ErrorMessage
	}
	return ""
}

// / Result of the Suite Execution.
type SuiteExecutionResult struct {
	SuiteResult      *ProtoSuiteResult `protobuf:"bytes,1,req,name=suiteResult" json:"suiteResult,omitempty"`
	XXX_unrecognized []byte            `json:"-"`
}

func (m *SuiteExecutionResult) Reset()         { *m = SuiteExecutionResult{} }
func (m *SuiteExecutionResult) String() string { return proto.CompactTextString(m) }
func (*SuiteExecutionResult) ProtoMessage()    {}

func (m *SuiteExecutionResult) GetSuiteResult() *ProtoSuiteResult {
	if m != nil {
		return m.SuiteResult
	}
	return nil
}

// / Requests Gauge to give all Step Names.
type StepNamesRequest struct {
	XXX_unrecognized []byte `json:"-"`
}

func (m *StepNamesRequest) Reset()         { *m = StepNamesRequest{} }
func (m *StepNamesRequest) String() string { return proto.CompactTextString(m) }
func (*StepNamesRequest) ProtoMessage()    {}

// / Response to StepNamesRequest
type StepNamesResponse struct {
	// / Collection of strings corresponding to Step texts.
	Steps            []string `protobuf:"bytes,1,rep,name=steps" json:"steps,omitempty"`
	XXX_unrecognized []byte   `json:"-"`
}

func (m *StepNamesResponse) Reset()         { *m = StepNamesResponse{} }
func (m *StepNamesResponse) String() string { return proto.CompactTextString(m) }
func (*StepNamesResponse) ProtoMessage()    {}

func (m *StepNamesResponse) GetSteps() []string {
	if m != nil {
		return m.Steps
	}
	return nil
}

// / Request runner to initialize Scenario DataStore
// / Scenario Datastore is reset after every Scenario execution.
type ScenarioDataStoreInitRequest struct {
	XXX_unrecognized []byte `json:"-"`
}

func (m *ScenarioDataStoreInitRequest) Reset()         { *m = ScenarioDataStoreInitRequest{} }
func (m *ScenarioDataStoreInitRequest) String() string { return proto.CompactTextString(m) }
func (*ScenarioDataStoreInitRequest) ProtoMessage()    {}

// / Request runner to initialize Spec DataStore
// / Spec Datastore is reset after every Spec execution.
type SpecDataStoreInitRequest struct {
	XXX_unrecognized []byte `json:"-"`
}

func (m *SpecDataStoreInitRequest) Reset()         { *m = SpecDataStoreInitRequest{} }
func (m *SpecDataStoreInitRequest) String() string { return proto.CompactTextString(m) }
func (*SpecDataStoreInitRequest) ProtoMessage()    {}

// / Request runner to initialize Suite DataStore
// / Suite Datastore is reset after every Suite execution.
type SuiteDataStoreInitRequest struct {
	XXX_unrecognized []byte `json:"-"`
}

func (m *SuiteDataStoreInitRequest) Reset()         { *m = SuiteDataStoreInitRequest{} }
func (m *SuiteDataStoreInitRequest) String() string { return proto.CompactTextString(m) }
func (*SuiteDataStoreInitRequest) ProtoMessage()    {}

// / Holds the new and old positions of a parameter.
// / Used when refactoring a Step.
type ParameterPosition struct {
	OldPosition      *int32 `protobuf:"varint,1,req,name=oldPosition" json:"oldPosition,omitempty"`
	NewPosition      *int32 `protobuf:"varint,2,req,name=newPosition" json:"newPosition,omitempty"`
	XXX_unrecognized []byte `json:"-"`
}

func (m *ParameterPosition) Reset()         { *m = ParameterPosition{} }
func (m *ParameterPosition) String() string { return proto.CompactTextString(m) }
func (*ParameterPosition) ProtoMessage()    {}

func (m *ParameterPosition) GetOldPosition() int32 {
	if m != nil && m.OldPosition != nil {
		return *m.OldPosition
	}
	return 0
}

func (m *ParameterPosition) GetNewPosition() int32 {
	if m != nil && m.NewPosition != nil {
		return *m.NewPosition
	}
	return 0
}

// / Tells the runner to refactor the specified Step.
type RefactorRequest struct {
	// / Old value, used to lookup Step to refactor
	OldStepValue *ProtoStepValue `protobuf:"bytes,1,req,name=oldStepValue" json:"oldStepValue,omitempty"`
	// / New value, the to-be value of Step being refactored.
	NewStepValue *ProtoStepValue `protobuf:"bytes,2,req,name=newStepValue" json:"newStepValue,omitempty"`
	// / Holds parameter positions of all parameters. Contains old and new parameter positions.
	ParamPositions   []*ParameterPosition `protobuf:"bytes,3,rep,name=paramPositions" json:"paramPositions,omitempty"`
	XXX_unrecognized []byte               `json:"-"`
}

func (m *RefactorRequest) Reset()         { *m = RefactorRequest{} }
func (m *RefactorRequest) String() string { return proto.CompactTextString(m) }
func (*RefactorRequest) ProtoMessage()    {}

func (m *RefactorRequest) GetOldStepValue() *ProtoStepValue {
	if m != nil {
		return m.OldStepValue
	}
	return nil
}

func (m *RefactorRequest) GetNewStepValue() *ProtoStepValue {
	if m != nil {
		return m.NewStepValue
	}
	return nil
}

func (m *RefactorRequest) GetParamPositions() []*ParameterPosition {
	if m != nil {
		return m.ParamPositions
	}
	return nil
}

// / Response of a RefactorRequest
type RefactorResponse struct {
	// / Flag indicating the success of Refactor operation.
	Success *bool `protobuf:"varint,1,req,name=success" json:"success,omitempty"`
	// / Error message, valid only if Refactor wasn't successful
	Error *string `protobuf:"bytes,2,opt,name=error" json:"error,omitempty"`
	// / List of files that were affected because of the refactoring.
	FilesChanged     []string `protobuf:"bytes,3,rep,name=filesChanged" json:"filesChanged,omitempty"`
	XXX_unrecognized []byte   `json:"-"`
}

func (m *RefactorResponse) Reset()         { *m = RefactorResponse{} }
func (m *RefactorResponse) String() string { return proto.CompactTextString(m) }
func (*RefactorResponse) ProtoMessage()    {}

func (m *RefactorResponse) GetSuccess() bool {
	if m != nil && m.Success != nil {
		return *m.Success
	}
	return false
}

func (m *RefactorResponse) GetError() string {
	if m != nil && m.Error != nil {
		return *m.Error
	}
	return ""
}

func (m *RefactorResponse) GetFilesChanged() []string {
	if m != nil {
		return m.FilesChanged
	}
	return nil
}

// / Request for details on a Single Step.
type StepNameRequest struct {
	// / Step text to lookup the Step.
	// / This is the parsed step value, i.e. with placeholders for parameters.
	StepValue        *string `protobuf:"bytes,1,req,name=stepValue" json:"stepValue,omitempty"`
	XXX_unrecognized []byte  `json:"-"`
}

func (m *StepNameRequest) Reset()         { *m = StepNameRequest{} }
func (m *StepNameRequest) String() string { return proto.CompactTextString(m) }
func (*StepNameRequest) ProtoMessage()    {}

func (m *StepNameRequest) GetStepValue() string {
	if m != nil && m.StepValue != nil {
		return *m.StepValue
	}
	return ""
}

// / Response to StepNameRequest.
type StepNameResponse struct {
	// / Flag indicating if there is a match for the given Step Text.
	IsStepPresent *bool `protobuf:"varint,1,req,name=isStepPresent" json:"isStepPresent,omitempty"`
	// / The Step name of the given step.
	StepName []string `protobuf:"bytes,2,rep,name=stepName" json:"stepName,omitempty"`
	// / Flag indicating if the given Step is an alias.
	HasAlias         *bool  `protobuf:"varint,3,req,name=hasAlias" json:"hasAlias,omitempty"`
	XXX_unrecognized []byte `json:"-"`
}

func (m *StepNameResponse) Reset()         { *m = StepNameResponse{} }
func (m *StepNameResponse) String() string { return proto.CompactTextString(m) }
func (*StepNameResponse) ProtoMessage()    {}

func (m *StepNameResponse) GetIsStepPresent() bool {
	if m != nil && m.IsStepPresent != nil {
		return *m.IsStepPresent
	}
	return false
}

func (m *StepNameResponse) GetStepName() []string {
	if m != nil {
		return m.StepName
	}
	return nil
}

func (m *StepNameResponse) GetHasAlias() bool {
	if m != nil && m.HasAlias != nil {
		return *m.HasAlias
	}
	return false
}

// / This is the message which gets transferred all the time
// / with proper message type set
// / One of the Request/Response fields will have value, depending on the MessageType set.
type Message struct {
	MessageType *Message_MessageType `protobuf:"varint,1,req,name=messageType,enum=gauge.messages.Message_MessageType" json:"messageType,omitempty"`
	// / A unique id to represent this message. A response to the message should copy over this value.
	// / This is used to synchronize messages & responses
	MessageId *int64 `protobuf:"varint,2,req,name=messageId" json:"messageId,omitempty"`
	// / [ExecutionStartingRequest](#gauge.messages.ExecutionStartingRequest)
	ExecutionStartingRequest *ExecutionStartingRequest `protobuf:"bytes,3,opt,name=executionStartingRequest" json:"executionStartingRequest,omitempty"`
	// / [SpecExecutionStartingRequest](#gauge.messages.SpecExecutionStartingRequest)
	SpecExecutionStartingRequest *SpecExecutionStartingRequest `protobuf:"bytes,4,opt,name=specExecutionStartingRequest" json:"specExecutionStartingRequest,omitempty"`
	// / [SpecExecutionEndingRequest](#gauge.messages.SpecExecutionEndingRequest)
	SpecExecutionEndingRequest *SpecExecutionEndingRequest `protobuf:"bytes,5,opt,name=specExecutionEndingRequest" json:"specExecutionEndingRequest,omitempty"`
	// / [ScenarioExecutionStartingRequest](#gauge.messages.ScenarioExecutionStartingRequest)
	ScenarioExecutionStartingRequest *ScenarioExecutionStartingRequest `protobuf:"bytes,6,opt,name=scenarioExecutionStartingRequest" json:"scenarioExecutionStartingRequest,omitempty"`
	// / [ScenarioExecutionEndingRequest](#gauge.messages.ScenarioExecutionEndingRequest)
	ScenarioExecutionEndingRequest *ScenarioExecutionEndingRequest `protobuf:"bytes,7,opt,name=scenarioExecutionEndingRequest" json:"scenarioExecutionEndingRequest,omitempty"`
	// / [StepExecutionStartingRequest](#gauge.messages.StepExecutionStartingRequest)
	StepExecutionStartingRequest *StepExecutionStartingRequest `protobuf:"bytes,8,opt,name=stepExecutionStartingRequest" json:"stepExecutionStartingRequest,omitempty"`
	// / [StepExecutionEndingRequest](#gauge.messages.StepExecutionEndingRequest)
	StepExecutionEndingRequest *StepExecutionEndingRequest `protobuf:"bytes,9,opt,name=stepExecutionEndingRequest" json:"stepExecutionEndingRequest,omitempty"`
	// / [ExecuteStepRequest](#gauge.messages.ExecuteStepRequest)
	ExecuteStepRequest *ExecuteStepRequest `protobuf:"bytes,10,opt,name=executeStepRequest" json:"executeStepRequest,omitempty"`
	// / [ExecutionEndingRequest](#gauge.messages.ExecutionEndingRequest)
	ExecutionEndingRequest *ExecutionEndingRequest `protobuf:"bytes,11,opt,name=executionEndingRequest" json:"executionEndingRequest,omitempty"`
	// / [StepValidateRequest](#gauge.messages.StepValidateRequest)
	StepValidateRequest *StepValidateRequest `protobuf:"bytes,12,opt,name=stepValidateRequest" json:"stepValidateRequest,omitempty"`
	// / [StepValidateResponse](#gauge.messages.StepValidateResponse)
	StepValidateResponse *StepValidateResponse `protobuf:"bytes,13,opt,name=stepValidateResponse" json:"stepValidateResponse,omitempty"`
	// / [ExecutionStatusResponse](#gauge.messages.ExecutionStatusResponse)
	ExecutionStatusResponse *ExecutionStatusResponse `protobuf:"bytes,14,opt,name=executionStatusResponse" json:"executionStatusResponse,omitempty"`
	// / [StepNamesRequest](#gauge.messages.StepNamesRequest)
	StepNamesRequest *StepNamesRequest `protobuf:"bytes,15,opt,name=stepNamesRequest" json:"stepNamesRequest,omitempty"`
	// / [StepNamesResponse](#gauge.messages.StepNamesResponse)
	StepNamesResponse *StepNamesResponse `protobuf:"bytes,16,opt,name=stepNamesResponse" json:"stepNamesResponse,omitempty"`
	// / [SuiteExecutionResult ](#gauge.messages.SuiteExecutionResult )
	SuiteExecutionResult *SuiteExecutionResult `protobuf:"bytes,17,opt,name=suiteExecutionResult" json:"suiteExecutionResult,omitempty"`
	// / [KillProcessRequest](#gauge.messages.KillProcessRequest)
	KillProcessRequest *KillProcessRequest `protobuf:"bytes,18,opt,name=killProcessRequest" json:"killProcessRequest,omitempty"`
	// / [ScenarioDataStoreInitRequest](#gauge.messages.ScenarioDataStoreInitRequest)
	ScenarioDataStoreInitRequest *ScenarioDataStoreInitRequest `protobuf:"bytes,19,opt,name=scenarioDataStoreInitRequest" json:"scenarioDataStoreInitRequest,omitempty"`
	// / [SpecDataStoreInitRequest](#gauge.messages.SpecDataStoreInitRequest)
	SpecDataStoreInitRequest *SpecDataStoreInitRequest `protobuf:"bytes,20,opt,name=specDataStoreInitRequest" json:"specDataStoreInitRequest,omitempty"`
	// / [SuiteDataStoreInitRequest](#gauge.messages.SuiteDataStoreInitRequest)
	SuiteDataStoreInitRequest *SuiteDataStoreInitRequest `protobuf:"bytes,21,opt,name=suiteDataStoreInitRequest" json:"suiteDataStoreInitRequest,omitempty"`
	// / [StepNameRequest](#gauge.messages.StepNameRequest)
	StepNameRequest *StepNameRequest `protobuf:"bytes,22,opt,name=stepNameRequest" json:"stepNameRequest,omitempty"`
	// / [StepNameResponse](#gauge.messages.StepNameResponse)
	StepNameResponse *StepNameResponse `protobuf:"bytes,23,opt,name=stepNameResponse" json:"stepNameResponse,omitempty"`
	// / [RefactorRequest](#gauge.messages.RefactorRequest)
	RefactorRequest *RefactorRequest `protobuf:"bytes,24,opt,name=refactorRequest" json:"refactorRequest,omitempty"`
	// / [RefactorResponse](#gauge.messages.RefactorResponse)
	RefactorResponse *RefactorResponse `protobuf:"bytes,25,opt,name=refactorResponse" json:"refactorResponse,omitempty"`
	XXX_unrecognized []byte            `json:"-"`
}

func (m *Message) Reset()         { *m = Message{} }
func (m *Message) String() string { return proto.CompactTextString(m) }
func (*Message) ProtoMessage()    {}

func (m *Message) GetMessageType() Message_MessageType {
	if m != nil && m.MessageType != nil {
		return *m.MessageType
	}
	return Message_ExecutionStarting
}

func (m *Message) GetMessageId() int64 {
	if m != nil && m.MessageId != nil {
		return *m.MessageId
	}
	return 0
}

func (m *Message) GetExecutionStartingRequest() *ExecutionStartingRequest {
	if m != nil {
		return m.ExecutionStartingRequest
	}
	return nil
}

func (m *Message) GetSpecExecutionStartingRequest() *SpecExecutionStartingRequest {
	if m != nil {
		return m.SpecExecutionStartingRequest
	}
	return nil
}

func (m *Message) GetSpecExecutionEndingRequest() *SpecExecutionEndingRequest {
	if m != nil {
		return m.SpecExecutionEndingRequest
	}
	return nil
}

func (m *Message) GetScenarioExecutionStartingRequest() *ScenarioExecutionStartingRequest {
	if m != nil {
		return m.ScenarioExecutionStartingRequest
	}
	return nil
}

func (m *Message) GetScenarioExecutionEndingRequest() *ScenarioExecutionEndingRequest {
	if m != nil {
		return m.ScenarioExecutionEndingRequest
	}
	return nil
}

func (m *Message) GetStepExecutionStartingRequest() *StepExecutionStartingRequest {
	if m != nil {
		return m.StepExecutionStartingRequest
	}
	return nil
}

func (m *Message) GetStepExecutionEndingRequest() *StepExecutionEndingRequest {
	if m != nil {
		return m.StepExecutionEndingRequest
	}
	return nil
}

func (m *Message) GetExecuteStepRequest() *ExecuteStepRequest {
	if m != nil {
		return m.ExecuteStepRequest
	}
	return nil
}

func (m *Message) GetExecutionEndingRequest() *ExecutionEndingRequest {
	if m != nil {
		return m.ExecutionEndingRequest
	}
	return nil
}

func (m *Message) GetStepValidateRequest() *StepValidateRequest {
	if m != nil {
		return m.StepValidateRequest
	}
	return nil
}

func (m *Message) GetStepValidateResponse() *StepValidateResponse {
	if m != nil {
		return m.StepValidateResponse
	}
	return nil
}

func (m *Message) GetExecutionStatusResponse() *ExecutionStatusResponse {
	if m != nil {
		return m.ExecutionStatusResponse
	}
	return nil
}

func (m *Message) GetStepNamesRequest() *StepNamesRequest {
	if m != nil {
		return m.StepNamesRequest
	}
	return nil
}

func (m *Message) GetStepNamesResponse() *StepNamesResponse {
	if m != nil {
		return m.StepNamesResponse
	}
	return nil
}

func (m *Message) GetSuiteExecutionResult() *SuiteExecutionResult {
	if m != nil {
		return m.SuiteExecutionResult
	}
	return nil
}

func (m *Message) GetKillProcessRequest() *KillProcessRequest {
	if m != nil {
		return m.KillProcessRequest
	}
	return nil
}

func (m *Message) GetScenarioDataStoreInitRequest() *ScenarioDataStoreInitRequest {
	if m != nil {
		return m.ScenarioDataStoreInitRequest
	}
	return nil
}

func (m *Message) GetSpecDataStoreInitRequest() *SpecDataStoreInitRequest {
	if m != nil {
		return m.SpecDataStoreInitRequest
	}
	return nil
}

func (m *Message) GetSuiteDataStoreInitRequest() *SuiteDataStoreInitRequest {
	if m != nil {
		return m.SuiteDataStoreInitRequest
	}
	return nil
}

func (m *Message) GetStepNameRequest() *StepNameRequest {
	if m != nil {
		return m.StepNameRequest
	}
	return nil
}

func (m *Message) GetStepNameResponse() *StepNameResponse {
	if m != nil {
		return m.StepNameResponse
	}
	return nil
}

func (m *Message) GetRefactorRequest() *RefactorRequest {
	if m != nil {
		return m.RefactorRequest
	}
	return nil
}

func (m *Message) GetRefactorResponse() *RefactorResponse {
	if m != nil {
		return m.RefactorResponse
	}
	return nil
}

func init() {
	proto.RegisterEnum("gauge.messages.Message_MessageType", Message_MessageType_name, Message_MessageType_value)
}
