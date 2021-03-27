// Code generated by the FlatBuffers compiler. DO NOT EDIT.

package models

import (
	flatbuffers "github.com/google/flatbuffers/go"
)

type StageT struct {
	Wait uint64
	Type uint16
	Verb byte
	Flag byte
	EnID string
	Text string
}

func (t *StageT) Pack(builder *flatbuffers.Builder) flatbuffers.UOffsetT {
	if t == nil { return 0 }
	EnIDOffset := builder.CreateString(t.EnID)
	TextOffset := builder.CreateString(t.Text)
	StageStart(builder)
	StageAddWait(builder, t.Wait)
	StageAddType(builder, t.Type)
	StageAddVerb(builder, t.Verb)
	StageAddFlag(builder, t.Flag)
	StageAddEnID(builder, EnIDOffset)
	StageAddText(builder, TextOffset)
	return StageEnd(builder)
}

func (rcv *Stage) UnPackTo(t *StageT) {
	t.Wait = rcv.Wait()
	t.Type = rcv.Type()
	t.Verb = rcv.Verb()
	t.Flag = rcv.Flag()
	t.EnID = string(rcv.EnID())
	t.Text = string(rcv.Text())
}

func (rcv *Stage) UnPack() *StageT {
	if rcv == nil { return nil }
	t := &StageT{}
	rcv.UnPackTo(t)
	return t
}

type Stage struct {
	_tab flatbuffers.Table
}

func GetRootAsStage(buf []byte, offset flatbuffers.UOffsetT) *Stage {
	n := flatbuffers.GetUOffsetT(buf[offset:])
	x := &Stage{}
	x.Init(buf, n+offset)
	return x
}

func (rcv *Stage) Init(buf []byte, i flatbuffers.UOffsetT) {
	rcv._tab.Bytes = buf
	rcv._tab.Pos = i
}

func (rcv *Stage) Table() flatbuffers.Table {
	return rcv._tab
}

func (rcv *Stage) Wait() uint64 {
	o := flatbuffers.UOffsetT(rcv._tab.Offset(4))
	if o != 0 {
		return rcv._tab.GetUint64(o + rcv._tab.Pos)
	}
	return 0
}

func (rcv *Stage) MutateWait(n uint64) bool {
	return rcv._tab.MutateUint64Slot(4, n)
}

func (rcv *Stage) Type() uint16 {
	o := flatbuffers.UOffsetT(rcv._tab.Offset(6))
	if o != 0 {
		return rcv._tab.GetUint16(o + rcv._tab.Pos)
	}
	return 0
}

func (rcv *Stage) MutateType(n uint16) bool {
	return rcv._tab.MutateUint16Slot(6, n)
}

func (rcv *Stage) Verb() byte {
	o := flatbuffers.UOffsetT(rcv._tab.Offset(8))
	if o != 0 {
		return rcv._tab.GetByte(o + rcv._tab.Pos)
	}
	return 0
}

func (rcv *Stage) MutateVerb(n byte) bool {
	return rcv._tab.MutateByteSlot(8, n)
}

func (rcv *Stage) Flag() byte {
	o := flatbuffers.UOffsetT(rcv._tab.Offset(10))
	if o != 0 {
		return rcv._tab.GetByte(o + rcv._tab.Pos)
	}
	return 0
}

func (rcv *Stage) MutateFlag(n byte) bool {
	return rcv._tab.MutateByteSlot(10, n)
}

func (rcv *Stage) EnID() []byte {
	o := flatbuffers.UOffsetT(rcv._tab.Offset(12))
	if o != 0 {
		return rcv._tab.ByteVector(o + rcv._tab.Pos)
	}
	return nil
}

func (rcv *Stage) Text() []byte {
	o := flatbuffers.UOffsetT(rcv._tab.Offset(14))
	if o != 0 {
		return rcv._tab.ByteVector(o + rcv._tab.Pos)
	}
	return nil
}

func StageStart(builder *flatbuffers.Builder) {
	builder.StartObject(6)
}
func StageAddWait(builder *flatbuffers.Builder, Wait uint64) {
	builder.PrependUint64Slot(0, Wait, 0)
}
func StageAddType(builder *flatbuffers.Builder, Type uint16) {
	builder.PrependUint16Slot(1, Type, 0)
}
func StageAddVerb(builder *flatbuffers.Builder, Verb byte) {
	builder.PrependByteSlot(2, Verb, 0)
}
func StageAddFlag(builder *flatbuffers.Builder, Flag byte) {
	builder.PrependByteSlot(3, Flag, 0)
}
func StageAddEnID(builder *flatbuffers.Builder, EnID flatbuffers.UOffsetT) {
	builder.PrependUOffsetTSlot(4, flatbuffers.UOffsetT(EnID), 0)
}
func StageAddText(builder *flatbuffers.Builder, Text flatbuffers.UOffsetT) {
	builder.PrependUOffsetTSlot(5, flatbuffers.UOffsetT(Text), 0)
}
func StageEnd(builder *flatbuffers.Builder) flatbuffers.UOffsetT {
	return builder.EndObject()
}
