// Code generated by protoc-gen-go-dubbo. DO NOT EDIT.

// Source: greet.proto
// Package: greet

package greet

import (
	dubbo_go_hessian2 "github.com/apache/dubbo-go-hessian2"
	java8_time "github.com/apache/dubbo-go-hessian2/java8_time"
	java_exception "github.com/apache/dubbo-go-hessian2/java_exception"
	java_sql_time "github.com/apache/dubbo-go-hessian2/java_sql_time"
	java_util "github.com/apache/dubbo-go-hessian2/java_util"
	java_math "github.com/dubbogo/gost/math/big"
)

type GreetEnum dubbo_go_hessian2.JavaEnum

const (
	GreetEnum_GREET_ENUM_0 GreetEnum = 0
	GreetEnum_GREET_ENUM_1 GreetEnum = 1
	GreetEnum_GREET_ENUM_2 GreetEnum = 2
	GreetEnum_GREET_ENUM_3 GreetEnum = 3
	GreetEnum_GREET_ENUM_4 GreetEnum = 4
)

// Enum value maps for GreetEnum
var (
	GreetEnum_name = map[int32]string{
		0: "GreetEnum_GREET_ENUM_0",
		1: "GreetEnum_GREET_ENUM_1",
		2: "GreetEnum_GREET_ENUM_2",
		3: "GreetEnum_GREET_ENUM_3",
		4: "GreetEnum_GREET_ENUM_4",
	}
	GreetEnum_value = map[string]int32{
		"GreetEnum_GREET_ENUM_0": 0,
		"GreetEnum_GREET_ENUM_1": 1,
		"GreetEnum_GREET_ENUM_2": 2,
		"GreetEnum_GREET_ENUM_3": 3,
		"GreetEnum_GREET_ENUM_4": 4,
	}
)

func (x GreetEnum) Enum() *GreetEnum {
	p := new(GreetEnum)
	*p = x
	return p
}

func (x GreetEnum) String() string {
	return GreetEnum_name[int32(x)]
}

func (x GreetEnum) Number() int32 {
	return GreetEnum_value[x.String()]
}

func (x GreetEnum) EnumValue(s string) dubbo_go_hessian2.JavaEnum {
	return dubbo_go_hessian2.JavaEnum(GreetEnum_value[x.String()])
}

func (x GreetEnum) JavaClassName() string {
	return "org.apache.greet.GreetEnum"
}

type GreetRequest struct {
	GreetRequest_Internal
	Name             string
	Way              GreetEnum
	Inner            *GreetRequest_Internal
	Time             *java_sql_time.Time
	Duration         *java8_time.Duration
	RuntimeException *java_exception.RuntimeException
	Uuid             *java_util.UUID
	SelfTime         *java_sql_time.Time
	BigInteger       *java_math.Integer
}

func (x *GreetRequest) JavaClassName() string {
	return "org.apache.greet.GreetRequest"
}

func (x *GreetRequest) String() string {
	e := dubbo_go_hessian2.NewEncoder()
	err := e.Encode(x)
	if err != nil {
		return ""
	}
	return string(e.Buffer())
}

func (x *GreetRequest) GetName() string {
	if x != nil {
		return x.Name
	}
	return ""
}

func (x *GreetRequest) GetWay() GreetEnum {
	if x != nil {
		return x.Way
	}
	return GreetEnum_GREET_ENUM_0
}

func (x *GreetRequest) GetInner() *GreetRequest_Internal {
	if x != nil {
		return x.Inner
	}
	return nil
}

func (x *GreetRequest) GetTime() *java_sql_time.Time {
	if x != nil {
		return x.Time
	}
	return nil
}

func (x *GreetRequest) GetDuration() *java8_time.Duration {
	if x != nil {
		return x.Duration
	}
	return nil
}

func (x *GreetRequest) GetRuntimeException() *java_exception.RuntimeException {
	if x != nil {
		return x.RuntimeException
	}
	return nil
}

func (x *GreetRequest) GetUuid() *java_util.UUID {
	if x != nil {
		return x.Uuid
	}
	return nil
}

func (x *GreetRequest) GetSelfTime() *java_sql_time.Time {
	if x != nil {
		return x.SelfTime
	}
	return nil
}

func (x *GreetRequest) GetBigInteger() *java_math.Integer {
	if x != nil {
		return x.BigInteger
	}
	return nil
}

type GreetRequest_Internal struct {
	Num int32
}

func (x *GreetRequest_Internal) JavaClassName() string {
	return "org.apache.greet.Inner"
}

func (x *GreetRequest_Internal) String() string {
	e := dubbo_go_hessian2.NewEncoder()
	err := e.Encode(x)
	if err != nil {
		return ""
	}
	return string(e.Buffer())
}

func (x *GreetRequest_Internal) GetNum() int32 {
	if x != nil {
		return x.Num
	}
	return 0
}

type GreetResponse struct {
	Greeting string
}

func (x *GreetResponse) JavaClassName() string {
	return "org.apache.greet.GreetResponse"
}

func (x *GreetResponse) String() string {
	e := dubbo_go_hessian2.NewEncoder()
	err := e.Encode(x)
	if err != nil {
		return ""
	}
	return string(e.Buffer())
}

func (x *GreetResponse) GetGreeting() string {
	if x != nil {
		return x.Greeting
	}
	return ""
}

func init() {
	dubbo_go_hessian2.RegisterPOJO(new(GreetRequest))
	dubbo_go_hessian2.RegisterPOJO(new(GreetRequest_Internal))
	dubbo_go_hessian2.RegisterPOJO(new(GreetResponse))

	for v := range GreetEnum_name {
		dubbo_go_hessian2.RegisterJavaEnum(GreetEnum(v))
		dubbo_go_hessian2.RegisterJavaEnum(GreetEnum(v))
		dubbo_go_hessian2.RegisterJavaEnum(GreetEnum(v))
		dubbo_go_hessian2.RegisterJavaEnum(GreetEnum(v))
		dubbo_go_hessian2.RegisterJavaEnum(GreetEnum(v))
	}
}