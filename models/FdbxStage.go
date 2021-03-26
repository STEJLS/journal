// Code generated by the FlatBuffers compiler. DO NOT EDIT.

package models

import (
	flatbuffers "github.com/google/flatbuffers/go"
)

type FdbxStageT struct {
	Dur int64
	Mtp int32
	Mid string
	Msg string
}

func (t *FdbxStageT) Pack(builder *flatbuffers.Builder) flatbuffers.UOffsetT {
	if t == nil { return 0 }
	midOffset := builder.CreateString(t.Mid)
	msgOffset := builder.CreateString(t.Msg)
	FdbxStageStart(builder)
	FdbxStageAddDur(builder, t.Dur)
	FdbxStageAddMtp(builder, t.Mtp)
	FdbxStageAddMid(builder, midOffset)
	FdbxStageAddMsg(builder, msgOffset)
	return FdbxStageEnd(builder)
}

func (rcv *FdbxStage) UnPackTo(t *FdbxStageT) {
	t.Dur = rcv.Dur()
	t.Mtp = rcv.Mtp()
	t.Mid = string(rcv.Mid())
	t.Msg = string(rcv.Msg())
}

func (rcv *FdbxStage) UnPack() *FdbxStageT {
	if rcv == nil { return nil }
	t := &FdbxStageT{}
	rcv.UnPackTo(t)
	return t
}

type FdbxStage struct {
	_tab flatbuffers.Table
}

func GetRootAsFdbxStage(buf []byte, offset flatbuffers.UOffsetT) *FdbxStage {
	n := flatbuffers.GetUOffsetT(buf[offset:])
	x := &FdbxStage{}
	x.Init(buf, n+offset)
	return x
}

func (rcv *FdbxStage) Init(buf []byte, i flatbuffers.UOffsetT) {
	rcv._tab.Bytes = buf
	rcv._tab.Pos = i
}

func (rcv *FdbxStage) Table() flatbuffers.Table {
	return rcv._tab
}

func (rcv *FdbxStage) Dur() int64 {
	o := flatbuffers.UOffsetT(rcv._tab.Offset(4))
	if o != 0 {
		return rcv._tab.GetInt64(o + rcv._tab.Pos)
	}
	return 0
}

func (rcv *FdbxStage) MutateDur(n int64) bool {
	return rcv._tab.MutateInt64Slot(4, n)
}

func (rcv *FdbxStage) Mtp() int32 {
	o := flatbuffers.UOffsetT(rcv._tab.Offset(6))
	if o != 0 {
		return rcv._tab.GetInt32(o + rcv._tab.Pos)
	}
	return 0
}

func (rcv *FdbxStage) MutateMtp(n int32) bool {
	return rcv._tab.MutateInt32Slot(6, n)
}

func (rcv *FdbxStage) Mid() []byte {
	o := flatbuffers.UOffsetT(rcv._tab.Offset(8))
	if o != 0 {
		return rcv._tab.ByteVector(o + rcv._tab.Pos)
	}
	return nil
}

func (rcv *FdbxStage) Msg() []byte {
	o := flatbuffers.UOffsetT(rcv._tab.Offset(10))
	if o != 0 {
		return rcv._tab.ByteVector(o + rcv._tab.Pos)
	}
	return nil
}

func FdbxStageStart(builder *flatbuffers.Builder) {
	builder.StartObject(4)
}
func FdbxStageAddDur(builder *flatbuffers.Builder, dur int64) {
	builder.PrependInt64Slot(0, dur, 0)
}
func FdbxStageAddMtp(builder *flatbuffers.Builder, mtp int32) {
	builder.PrependInt32Slot(1, mtp, 0)
}
func FdbxStageAddMid(builder *flatbuffers.Builder, mid flatbuffers.UOffsetT) {
	builder.PrependUOffsetTSlot(2, flatbuffers.UOffsetT(mid), 0)
}
func FdbxStageAddMsg(builder *flatbuffers.Builder, msg flatbuffers.UOffsetT) {
	builder.PrependUOffsetTSlot(3, flatbuffers.UOffsetT(msg), 0)
}
func FdbxStageEnd(builder *flatbuffers.Builder) flatbuffers.UOffsetT {
	return builder.EndObject()
}
