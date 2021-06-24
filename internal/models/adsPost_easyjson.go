// Code generated by easyjson for marshaling/unmarshaling. DO NOT EDIT.

package models

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

func easyjson6703527eDecodeTestTaskAdvertisingInternalModels(in *jlexer.Lexer, out *AdsPostRequest) {
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
			out.Id = uint64(in.Uint64())
		case "fields":
			if in.IsNull() {
				in.Skip()
				out.Fields = nil
			} else {
				in.Delim('[')
				if out.Fields == nil {
					if !in.IsDelim(']') {
						out.Fields = make([]string, 0, 4)
					} else {
						out.Fields = []string{}
					}
				} else {
					out.Fields = (out.Fields)[:0]
				}
				for !in.IsDelim(']') {
					var v1 string
					v1 = string(in.String())
					out.Fields = append(out.Fields, v1)
					in.WantComma()
				}
				in.Delim(']')
			}
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
func easyjson6703527eEncodeTestTaskAdvertisingInternalModels(out *jwriter.Writer, in AdsPostRequest) {
	out.RawByte('{')
	first := true
	_ = first
	{
		const prefix string = ",\"id\":"
		out.RawString(prefix[1:])
		out.Uint64(uint64(in.Id))
	}
	{
		const prefix string = ",\"fields\":"
		out.RawString(prefix)
		if in.Fields == nil && (out.Flags&jwriter.NilSliceAsEmpty) == 0 {
			out.RawString("null")
		} else {
			out.RawByte('[')
			for v2, v3 := range in.Fields {
				if v2 > 0 {
					out.RawByte(',')
				}
				out.String(string(v3))
			}
			out.RawByte(']')
		}
	}
	out.RawByte('}')
}

// MarshalJSON supports json.Marshaler interface
func (v AdsPostRequest) MarshalJSON() ([]byte, error) {
	w := jwriter.Writer{}
	easyjson6703527eEncodeTestTaskAdvertisingInternalModels(&w, v)
	return w.Buffer.BuildBytes(), w.Error
}

// MarshalEasyJSON supports easyjson.Marshaler interface
func (v AdsPostRequest) MarshalEasyJSON(w *jwriter.Writer) {
	easyjson6703527eEncodeTestTaskAdvertisingInternalModels(w, v)
}

// UnmarshalJSON supports json.Unmarshaler interface
func (v *AdsPostRequest) UnmarshalJSON(data []byte) error {
	r := jlexer.Lexer{Data: data}
	easyjson6703527eDecodeTestTaskAdvertisingInternalModels(&r, v)
	return r.Error()
}

// UnmarshalEasyJSON supports easyjson.Unmarshaler interface
func (v *AdsPostRequest) UnmarshalEasyJSON(l *jlexer.Lexer) {
	easyjson6703527eDecodeTestTaskAdvertisingInternalModels(l, v)
}
func easyjson6703527eDecodeTestTaskAdvertisingInternalModels1(in *jlexer.Lexer, out *AdsPostId) {
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
			out.Id = uint64(in.Uint64())
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
func easyjson6703527eEncodeTestTaskAdvertisingInternalModels1(out *jwriter.Writer, in AdsPostId) {
	out.RawByte('{')
	first := true
	_ = first
	{
		const prefix string = ",\"id\":"
		out.RawString(prefix[1:])
		out.Uint64(uint64(in.Id))
	}
	out.RawByte('}')
}

// MarshalJSON supports json.Marshaler interface
func (v AdsPostId) MarshalJSON() ([]byte, error) {
	w := jwriter.Writer{}
	easyjson6703527eEncodeTestTaskAdvertisingInternalModels1(&w, v)
	return w.Buffer.BuildBytes(), w.Error
}

// MarshalEasyJSON supports easyjson.Marshaler interface
func (v AdsPostId) MarshalEasyJSON(w *jwriter.Writer) {
	easyjson6703527eEncodeTestTaskAdvertisingInternalModels1(w, v)
}

// UnmarshalJSON supports json.Unmarshaler interface
func (v *AdsPostId) UnmarshalJSON(data []byte) error {
	r := jlexer.Lexer{Data: data}
	easyjson6703527eDecodeTestTaskAdvertisingInternalModels1(&r, v)
	return r.Error()
}

// UnmarshalEasyJSON supports easyjson.Unmarshaler interface
func (v *AdsPostId) UnmarshalEasyJSON(l *jlexer.Lexer) {
	easyjson6703527eDecodeTestTaskAdvertisingInternalModels1(l, v)
}
func easyjson6703527eDecodeTestTaskAdvertisingInternalModels2(in *jlexer.Lexer, out *AdsPostArrRequest) {
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
		case "start":
			out.Start = uint64(in.Uint64())
		case "count":
			out.Count = uint64(in.Uint64())
		case "sort":
			out.Sort = string(in.String())
		case "desc":
			out.Desc = bool(in.Bool())
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
func easyjson6703527eEncodeTestTaskAdvertisingInternalModels2(out *jwriter.Writer, in AdsPostArrRequest) {
	out.RawByte('{')
	first := true
	_ = first
	{
		const prefix string = ",\"start\":"
		out.RawString(prefix[1:])
		out.Uint64(uint64(in.Start))
	}
	{
		const prefix string = ",\"count\":"
		out.RawString(prefix)
		out.Uint64(uint64(in.Count))
	}
	{
		const prefix string = ",\"sort\":"
		out.RawString(prefix)
		out.String(string(in.Sort))
	}
	{
		const prefix string = ",\"desc\":"
		out.RawString(prefix)
		out.Bool(bool(in.Desc))
	}
	out.RawByte('}')
}

// MarshalJSON supports json.Marshaler interface
func (v AdsPostArrRequest) MarshalJSON() ([]byte, error) {
	w := jwriter.Writer{}
	easyjson6703527eEncodeTestTaskAdvertisingInternalModels2(&w, v)
	return w.Buffer.BuildBytes(), w.Error
}

// MarshalEasyJSON supports easyjson.Marshaler interface
func (v AdsPostArrRequest) MarshalEasyJSON(w *jwriter.Writer) {
	easyjson6703527eEncodeTestTaskAdvertisingInternalModels2(w, v)
}

// UnmarshalJSON supports json.Unmarshaler interface
func (v *AdsPostArrRequest) UnmarshalJSON(data []byte) error {
	r := jlexer.Lexer{Data: data}
	easyjson6703527eDecodeTestTaskAdvertisingInternalModels2(&r, v)
	return r.Error()
}

// UnmarshalEasyJSON supports easyjson.Unmarshaler interface
func (v *AdsPostArrRequest) UnmarshalEasyJSON(l *jlexer.Lexer) {
	easyjson6703527eDecodeTestTaskAdvertisingInternalModels2(l, v)
}
func easyjson6703527eDecodeTestTaskAdvertisingInternalModels3(in *jlexer.Lexer, out *AdsPostArrItem) {
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
			out.Id = uint64(in.Uint64())
		case "title":
			out.Title = string(in.String())
		case "photo":
			out.Photo = string(in.String())
		case "price":
			out.Price = uint64(in.Uint64())
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
func easyjson6703527eEncodeTestTaskAdvertisingInternalModels3(out *jwriter.Writer, in AdsPostArrItem) {
	out.RawByte('{')
	first := true
	_ = first
	{
		const prefix string = ",\"id\":"
		out.RawString(prefix[1:])
		out.Uint64(uint64(in.Id))
	}
	{
		const prefix string = ",\"title\":"
		out.RawString(prefix)
		out.String(string(in.Title))
	}
	{
		const prefix string = ",\"photo\":"
		out.RawString(prefix)
		out.String(string(in.Photo))
	}
	{
		const prefix string = ",\"price\":"
		out.RawString(prefix)
		out.Uint64(uint64(in.Price))
	}
	out.RawByte('}')
}

// MarshalJSON supports json.Marshaler interface
func (v AdsPostArrItem) MarshalJSON() ([]byte, error) {
	w := jwriter.Writer{}
	easyjson6703527eEncodeTestTaskAdvertisingInternalModels3(&w, v)
	return w.Buffer.BuildBytes(), w.Error
}

// MarshalEasyJSON supports easyjson.Marshaler interface
func (v AdsPostArrItem) MarshalEasyJSON(w *jwriter.Writer) {
	easyjson6703527eEncodeTestTaskAdvertisingInternalModels3(w, v)
}

// UnmarshalJSON supports json.Unmarshaler interface
func (v *AdsPostArrItem) UnmarshalJSON(data []byte) error {
	r := jlexer.Lexer{Data: data}
	easyjson6703527eDecodeTestTaskAdvertisingInternalModels3(&r, v)
	return r.Error()
}

// UnmarshalEasyJSON supports easyjson.Unmarshaler interface
func (v *AdsPostArrItem) UnmarshalEasyJSON(l *jlexer.Lexer) {
	easyjson6703527eDecodeTestTaskAdvertisingInternalModels3(l, v)
}
func easyjson6703527eDecodeTestTaskAdvertisingInternalModels4(in *jlexer.Lexer, out *AdsPost) {
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
			out.Id = uint64(in.Uint64())
		case "title":
			out.Title = string(in.String())
		case "description":
			out.Description = string(in.String())
		case "photos":
			if in.IsNull() {
				in.Skip()
				out.Photos = nil
			} else {
				in.Delim('[')
				if out.Photos == nil {
					if !in.IsDelim(']') {
						out.Photos = make([]string, 0, 4)
					} else {
						out.Photos = []string{}
					}
				} else {
					out.Photos = (out.Photos)[:0]
				}
				for !in.IsDelim(']') {
					var v4 string
					v4 = string(in.String())
					out.Photos = append(out.Photos, v4)
					in.WantComma()
				}
				in.Delim(']')
			}
		case "price":
			out.Price = uint64(in.Uint64())
		case "date":
			out.Date = string(in.String())
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
func easyjson6703527eEncodeTestTaskAdvertisingInternalModels4(out *jwriter.Writer, in AdsPost) {
	out.RawByte('{')
	first := true
	_ = first
	{
		const prefix string = ",\"id\":"
		out.RawString(prefix[1:])
		out.Uint64(uint64(in.Id))
	}
	{
		const prefix string = ",\"title\":"
		out.RawString(prefix)
		out.String(string(in.Title))
	}
	{
		const prefix string = ",\"description\":"
		out.RawString(prefix)
		out.String(string(in.Description))
	}
	{
		const prefix string = ",\"photos\":"
		out.RawString(prefix)
		if in.Photos == nil && (out.Flags&jwriter.NilSliceAsEmpty) == 0 {
			out.RawString("null")
		} else {
			out.RawByte('[')
			for v5, v6 := range in.Photos {
				if v5 > 0 {
					out.RawByte(',')
				}
				out.String(string(v6))
			}
			out.RawByte(']')
		}
	}
	{
		const prefix string = ",\"price\":"
		out.RawString(prefix)
		out.Uint64(uint64(in.Price))
	}
	{
		const prefix string = ",\"date\":"
		out.RawString(prefix)
		out.String(string(in.Date))
	}
	out.RawByte('}')
}

// MarshalJSON supports json.Marshaler interface
func (v AdsPost) MarshalJSON() ([]byte, error) {
	w := jwriter.Writer{}
	easyjson6703527eEncodeTestTaskAdvertisingInternalModels4(&w, v)
	return w.Buffer.BuildBytes(), w.Error
}

// MarshalEasyJSON supports easyjson.Marshaler interface
func (v AdsPost) MarshalEasyJSON(w *jwriter.Writer) {
	easyjson6703527eEncodeTestTaskAdvertisingInternalModels4(w, v)
}

// UnmarshalJSON supports json.Unmarshaler interface
func (v *AdsPost) UnmarshalJSON(data []byte) error {
	r := jlexer.Lexer{Data: data}
	easyjson6703527eDecodeTestTaskAdvertisingInternalModels4(&r, v)
	return r.Error()
}

// UnmarshalEasyJSON supports easyjson.Unmarshaler interface
func (v *AdsPost) UnmarshalEasyJSON(l *jlexer.Lexer) {
	easyjson6703527eDecodeTestTaskAdvertisingInternalModels4(l, v)
}
