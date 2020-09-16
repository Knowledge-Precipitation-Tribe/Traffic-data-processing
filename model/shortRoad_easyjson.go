// Code generated by easyjson for marshaling/unmarshaling. DO NOT EDIT.

package model

import (
	json "encoding/json"
	easyjson "github.com/mailru/easyjson"
	jlexer "github.com/mailru/easyjson/jlexer"
	jwriter "github.com/mailru/easyjson/jwriter"
)

// suppress unused package warning
var (
	_ *json.RawMessage
	_ *jlexer.Lexer
	_ *jwriter.Writer
	_ easyjson.Marshaler
)

func easyjsonC332dce6DecodeTrafficDataProcessingModel(in *jlexer.Lexer, out *ShortRoad) {
	isTopLevel := in.IsStart()
	if in.IsNull() {
		if isTopLevel {
			in.Consumed()
		}
		in.Skip()
		return
	}
	in.Delim('{')
	for !in.IsDelim('}') {
		key := in.UnsafeFieldName(false)
		in.WantColon()
		if in.IsNull() {
			in.Skip()
			in.WantComma()
			continue
		}
		switch key {
		case "id":
			out.Id = int(in.Int())
		case "tti_id":
			out.TTiId = int(in.Int())
		case "tti_name":
			out.TTIName = string(in.String())
		case "start_longitude":
			out.StartLongitude = string(in.String())
		case "start_latitude":
			out.StartLatitude = string(in.String())
		case "stop_longitude":
			out.StopLongitude = string(in.String())
		case "stop_latitude":
			out.StopLatitude = string(in.String())
		default:
			in.SkipRecursive()
		}
		in.WantComma()
	}
	in.Delim('}')
	if isTopLevel {
		in.Consumed()
	}
}
func easyjsonC332dce6EncodeTrafficDataProcessingModel(out *jwriter.Writer, in ShortRoad) {
	out.RawByte('{')
	first := true
	_ = first
	{
		const prefix string = ",\"id\":"
		out.RawString(prefix[1:])
		out.Int(int(in.Id))
	}
	{
		const prefix string = ",\"tti_id\":"
		out.RawString(prefix)
		out.Int(int(in.TTiId))
	}
	{
		const prefix string = ",\"tti_name\":"
		out.RawString(prefix)
		out.String(string(in.TTIName))
	}
	{
		const prefix string = ",\"start_longitude\":"
		out.RawString(prefix)
		out.String(string(in.StartLongitude))
	}
	{
		const prefix string = ",\"start_latitude\":"
		out.RawString(prefix)
		out.String(string(in.StartLatitude))
	}
	{
		const prefix string = ",\"stop_longitude\":"
		out.RawString(prefix)
		out.String(string(in.StopLongitude))
	}
	{
		const prefix string = ",\"stop_latitude\":"
		out.RawString(prefix)
		out.String(string(in.StopLatitude))
	}
	out.RawByte('}')
}

// MarshalJSON supports json.Marshaler interface
func (v ShortRoad) MarshalJSON() ([]byte, error) {
	w := jwriter.Writer{}
	easyjsonC332dce6EncodeTrafficDataProcessingModel(&w, v)
	return w.Buffer.BuildBytes(), w.Error
}

// MarshalEasyJSON supports easyjson.Marshaler interface
func (v ShortRoad) MarshalEasyJSON(w *jwriter.Writer) {
	easyjsonC332dce6EncodeTrafficDataProcessingModel(w, v)
}

// UnmarshalJSON supports json.Unmarshaler interface
func (v *ShortRoad) UnmarshalJSON(data []byte) error {
	r := jlexer.Lexer{Data: data}
	easyjsonC332dce6DecodeTrafficDataProcessingModel(&r, v)
	return r.Error()
}

// UnmarshalEasyJSON supports easyjson.Unmarshaler interface
func (v *ShortRoad) UnmarshalEasyJSON(l *jlexer.Lexer) {
	easyjsonC332dce6DecodeTrafficDataProcessingModel(l, v)
}