// Code generated by the FlatBuffers compiler. DO NOT EDIT.

package LMS

import (
	flatbuffers "github.com/google/flatbuffers/go"
)

type Phone struct {
	_tab flatbuffers.Table
}

func GetRootAsPhone(buf []byte, offset flatbuffers.UOffsetT) *Phone {
	n := flatbuffers.GetUOffsetT(buf[offset:])
	x := &Phone{}
	x.Init(buf, n+offset)
	return x
}

func (rcv *Phone) Init(buf []byte, i flatbuffers.UOffsetT) {
	rcv._tab.Bytes = buf
	rcv._tab.Pos = i
}

func (rcv *Phone) Table() flatbuffers.Table {
	return rcv._tab
}

func PhoneStart(builder *flatbuffers.Builder) {
	builder.StartObject(2)
}
func PhoneEnd(builder *flatbuffers.Builder) flatbuffers.UOffsetT {
	return builder.EndObject()
}
