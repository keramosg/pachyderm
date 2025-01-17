// Code generated by protoc-gen-zap (etc/proto/protoc-gen-zap). DO NOT EDIT.
//
// source: internal/storage/chunk/chunk.proto

package chunk

import (
	protoextensions "github.com/pachyderm/pachyderm/v2/src/protoextensions"
	zapcore "go.uber.org/zap/zapcore"
)

func (x *DataRef) MarshalLogObject(enc zapcore.ObjectEncoder) error {
	if x == nil {
		return nil
	}
	if obj, ok := interface{}(x.Ref).(zapcore.ObjectMarshaler); ok {
		enc.AddObject("ref", obj)
	} else {
		enc.AddReflected("ref", x.Ref)
	}
	protoextensions.AddBytes(enc, "hash", x.Hash)
	enc.AddInt64("offset_bytes", x.OffsetBytes)
	enc.AddInt64("size_bytes", x.SizeBytes)
	return nil
}

func (x *Ref) MarshalLogObject(enc zapcore.ObjectEncoder) error {
	if x == nil {
		return nil
	}
	protoextensions.AddBytes(enc, "id", x.Id)
	enc.AddInt64("size_bytes", x.SizeBytes)
	enc.AddBool("edge", x.Edge)
	protoextensions.AddBytes(enc, "dek", x.Dek)
	enc.AddString("encryption_algo", x.EncryptionAlgo.String())
	enc.AddString("compression_algo", x.CompressionAlgo.String())
	return nil
}
