// This file is automatically generated. Do not modify.

package scribe

import (
	"fmt"
	"strconv"
)

type ResultCode int32

var (
	ResultCodeOk       = ResultCode(0)
	ResultCodeTryLater = ResultCode(1)
	ResultCodeByName   = map[string]ResultCode{
		"ResultCode.OK":        ResultCodeOk,
		"ResultCode.TRY_LATER": ResultCodeTryLater,
	}
	ResultCodeByValue = map[ResultCode]string{
		ResultCodeOk:       "ResultCode.OK",
		ResultCodeTryLater: "ResultCode.TRY_LATER",
	}
)

func (e ResultCode) String() string {
	name := ResultCodeByValue[e]
	if name == "" {
		name = fmt.Sprintf("Unknown enum value ResultCode(%d)", e)
	}
	return name
}

func (e *ResultCode) UnmarshalJSON(b []byte) error {
	st := string(b)
	if st[0] == '"' {
		*e = ResultCode(ResultCodeByName[st[1:len(st)-1]])
		return nil
	}
	i, err := strconv.Atoi(st)
	*e = ResultCode(i)
	return err
}

type LogEntry struct {
	Category string `thrift:"1,required" json:"category"`
	Message  string `thrift:"2,required" json:"message"`
}

type RPCClient interface {
	Call(method string, request interface{}, response interface{}) error
}

type Scribe interface {
	Log(Messages []*LogEntry) (ResultCode, error)
}

type ScribeServer struct {
	Implementation Scribe
}

func (s *ScribeServer) Log(req *ScribeLogRequest, res *ScribeLogResponse) error {
	val, err := s.Implementation.Log(req.Messages)
	switch err.(type) {
	}
	res.Value = val
	return err
}

type ScribeLogRequest struct {
	Messages []*LogEntry `thrift:"1,required" json:"messages"`
}

type ScribeLogResponse struct {
	Value ResultCode `thrift:"0,required" json:"value"`
}

type ScribeClient struct {
	Client RPCClient
}

func (s *ScribeClient) Log(Messages []*LogEntry) (ResultCode, error) {
	req := &ScribeLogRequest{
		Messages: Messages,
	}
	res := &ScribeLogResponse{}
	err := s.Client.Call("Log", req, res)
	return res.Value, err
}
