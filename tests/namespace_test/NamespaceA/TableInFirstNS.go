// Code generated by the FlatBuffers compiler. DO NOT EDIT.

package NamespaceA

import (
	flatbuffers "github.com/google/flatbuffers/go"
)

type TableInFirstNS struct {
	_tab flatbuffers.Table
}

func GetRootAsTableInFirstNS(buf []byte, offset flatbuffers.UOffsetT) *TableInFirstNS {
	n := flatbuffers.GetUOffsetT(buf[offset:])
	x := &TableInFirstNS{}
	x.Init(buf, n+offset)
	return x
}

func (rcv *TableInFirstNS) Init(buf []byte, i flatbuffers.UOffsetT) {
	rcv._tab.Bytes = buf
	rcv._tab.Pos = i
}

func (rcv *TableInFirstNS) Table() flatbuffers.Table {
	return rcv._tab
}

func (rcv *TableInFirstNS) FooTable(obj *TableInNestedNS) *TableInNestedNS {
	o := flatbuffers.UOffsetT(rcv._tab.Offset(4))
	if o != 0 {
		x := rcv._tab.Indirect(o + rcv._tab.Pos)
		if obj == nil {
			obj = new(TableInNestedNS)
		}
		obj.Init(rcv._tab.Bytes, x)
		return obj
	}
	return nil
}

func (rcv *TableInFirstNS) FooEnum() EnumInNestedNS {
	o := flatbuffers.UOffsetT(rcv._tab.Offset(6))
	if o != 0 {
		return rcv._tab.GetInt8(o + rcv._tab.Pos)
	}
	return 0
}

func (rcv *TableInFirstNS) MutateFooEnum(n EnumInNestedNS) bool {
	return rcv._tab.MutateInt8Slot(6, n)
}

func (rcv *TableInFirstNS) FooStruct(obj *StructInNestedNS) *StructInNestedNS {
	o := flatbuffers.UOffsetT(rcv._tab.Offset(8))
	if o != 0 {
		x := o + rcv._tab.Pos
		if obj == nil {
			obj = new(StructInNestedNS)
		}
		obj.Init(rcv._tab.Bytes, x)
		return obj
	}
	return nil
}

func TableInFirstNSStart(builder *flatbuffers.Builder) {
	builder.StartObject(3)
}
func TableInFirstNSAddFooTable(builder *flatbuffers.Builder, fooTable flatbuffers.UOffsetT) {
	builder.PrependUOffsetTSlot(0, flatbuffers.UOffsetT(fooTable), 0)
}
func TableInFirstNSAddFooEnum(builder *flatbuffers.Builder, fooEnum int8) {
	builder.PrependInt8Slot(1, fooEnum, 0)
}
func TableInFirstNSAddFooStruct(builder *flatbuffers.Builder, fooStruct flatbuffers.UOffsetT) {
	builder.PrependStructSlot(2, flatbuffers.UOffsetT(fooStruct), 0)
}
func TableInFirstNSEnd(builder *flatbuffers.Builder) flatbuffers.UOffsetT {
	return builder.EndObject()
}
