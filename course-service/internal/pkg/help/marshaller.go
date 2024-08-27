package help

import (
	"encoding/json"

	"google.golang.org/protobuf/encoding/protojson"
	"google.golang.org/protobuf/reflect/protoreflect"
)

func JsonToJson(data interface{}, js interface{}) error {
	body, err := json.Marshal(js)
	if err != nil {
		return err
	}

	err = json.Unmarshal(body, data)
	return err
}

func ProtoToStruct(data interface{}, m protoreflect.ProtoMessage) error {
	jsonMarshaller := protojson.MarshalOptions{
		AllowPartial:    true,
		EmitUnpopulated: true,
	}

	js, err := jsonMarshaller.Marshal(m)
	if err != nil {
		return err
	}

	err = json.Unmarshal(js, data)
	return err
}

func StructToProto(m protoreflect.ProtoMessage, data interface{}) error {
	jsonMarshaller := protojson.UnmarshalOptions{
		AllowPartial:   true,
		DiscardUnknown: true,
	}

	js, err := json.Marshal(data)
	if err != nil {
		return err
	}

	return jsonMarshaller.Unmarshal(js, m)
}

func StringToProto(m protoreflect.ProtoMessage, s string) error {
	jsonMarshaller := protojson.UnmarshalOptions{
		AllowPartial:   true,
		DiscardUnknown: true,
	}

	return jsonMarshaller.Unmarshal([]byte(s), m)
}

func ProtoToString(m protoreflect.ProtoMessage) (string, error) {
	jsonMarshaller := protojson.MarshalOptions{
		AllowPartial:    true,
		EmitUnpopulated: true,
	}

	js, err := jsonMarshaller.Marshal(m)
	if err != nil {
		return "", err
	}

	return string(js), nil
}
