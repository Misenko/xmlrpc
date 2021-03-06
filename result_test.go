package xmlrpc

import (
	"strings"
	"testing"
)

const endpointXML = "http://127.0.0.1:8000/file.xml"

const (
	wrongXMLFormat                  = "records/wrong_xml_format"
	wrongXMLResponse                = "records/wrong_xml_response"
	wrongValueTag                   = "records/wrong_value_tag"
	parseErrorOnechildtag           = "records/parse_error_onechildtag"
	parseErrorWrongtag              = "records/parse_error_wrongtag"
	parseErrorInt                   = "records/parse_error_int"
	parseErrorDouble                = "records/parse_error_double"
	parseErrorTime                  = "records/parse_error_time"
	parseErrorArray                 = "records/parse_error_array"
	parseErrorArrayElement          = "records/parse_error_array_element"
	parseErrorBoolean               = "records/parse_error_boolean"
	parseErrorStructNoname          = "records/parse_error_struct_noname"
	parseErrorStructNovalue         = "records/parse_error_struct_novalue"
	parseErrorStructNomember        = "records/parse_error_struct_nomember"
	parseErrorStructMultipleMembers = "records/parse_error_struct_multiple_members"
	parseErrorStructOnechildtag     = "records/parse_error_struct_onechildtag"
	parseErrorStructElement         = "records/parse_error_struct_element"
	parseErrorBase64                = "records/parse_error_base64"

	parseFaultError   = "records/parse_fault"
	parseFaultName    = "records/parse_fault_name"
	parseFaultMembers = "records/parse_fault_members"
)

func Test_wrongXMLFormat(t *testing.T) {
	res, err := MakeCallAndCreateRecord(t, wrongXMLFormat, endpointXML, "")
	if err == nil {
		t.Fatal("No error when parse wrong XML response.")
	}
	if !strings.Contains(err.Error(), "cannot parse XML RPC response") {
		t.Fatal("Unexpected error:", err)
	}
	if res != nil {
		t.Fatal("Method Call returns result when parse wrong XML response.")
	}
}

func Test_wrongXMLResponse(t *testing.T) {
	res, err := MakeCallAndCreateRecord(t, wrongXMLResponse, endpointXML, "")
	if err == nil {
		t.Fatal("No error when parse wrong XML response.")
	}
	if !strings.Contains(err.Error(), "cannot parse XML RPC response") {
		t.Fatal("Unexpected error:", err)
	}
	if res != nil {
		t.Fatal("Method Call returns result when parse wrong XML response.")
	}
}

func Test_wrongValueTag(t *testing.T) {
	res, err := MakeCallAndCreateRecord(t, wrongValueTag, endpointXML, "")
	if err == nil {
		t.Fatal("No error when parse wrong XML response.")
	}
	if !strings.Contains(err.Error(), "cannot parse XML RPC response") {
		t.Fatal("Unexpected error:", err)
	}
	if res != nil {
		t.Fatal("Method Call returns result when parse wrong XML response.")
	}
}

func Test_parseError_int(t *testing.T) {
	res, err := MakeCallAndCreateRecord(t, parseErrorInt, endpointXML, "")
	if err == nil {
		t.Fatal("No error when parse wrong XML response.")
	}
	if !strings.Contains(err.Error(), "cannot parse XML RPC response") {
		t.Fatal("Unexpected error:", err)
	}
	if res != nil {
		t.Fatal("Method Call returns result when parse wrong XML response.")
	}
}

func Test_parseError_double(t *testing.T) {
	res, err := MakeCallAndCreateRecord(t, parseErrorDouble, endpointXML, "")
	if err == nil {
		t.Fatal("No error when parse wrong XML response.")
	}
	if !strings.Contains(err.Error(), "cannot parse XML RPC response") {
		t.Fatal("Unexpected error:", err)
	}
	if res != nil {
		t.Fatal("Method Call returns result when parse wrong XML response.")
	}
}

func Test_parseError_time(t *testing.T) {
	res, err := MakeCallAndCreateRecord(t, parseErrorTime, endpointXML, "")
	if err == nil {
		t.Fatal("No error when parse wrong XML response.")
	}
	if !strings.Contains(err.Error(), "cannot parse XML RPC response") {
		t.Fatal("Unexpected error:", err)
	}
	if res != nil {
		t.Fatal("Method Call returns result when parse wrong XML response.")
	}
}

func Test_parseError_array(t *testing.T) {
	res, err := MakeCallAndCreateRecord(t, parseErrorArray, endpointXML, "")
	if err == nil {
		t.Fatal("No error when parse wrong XML response.")
	}
	if !strings.Contains(err.Error(), "cannot parse XML RPC response") {
		t.Fatal("Unexpected error:", err)
	}
	if res != nil {
		t.Fatal("Method Call returns result when parse wrong XML response.")
	}
}

func Test_parseError_array_element(t *testing.T) {
	res, err := MakeCallAndCreateRecord(t, parseErrorArrayElement, endpointXML, "")
	if err == nil {
		t.Fatal("No error when parse wrong XML response.")
	}
	if !strings.Contains(err.Error(), "cannot parse XML RPC response") {
		t.Fatal("Unexpected error:", err)
	}
	if res != nil {
		t.Fatal("Method Call returns result when parse wrong XML response.")
	}
}

func Test_parseError_base64(t *testing.T) {
	res, err := MakeCallAndCreateRecord(t, parseErrorBase64, endpointXML, "")
	if err == nil {
		t.Fatal("No error when parse wrong XML response.")
	}
	if !strings.Contains(err.Error(), "cannot parse XML RPC response") {
		t.Fatal("Unexpected error:", err)
	}
	if res != nil {
		t.Fatal("Method Call returns result when parse wrong XML response.")
	}
}

func Test_parseError_boolean(t *testing.T) {
	res, err := MakeCallAndCreateRecord(t, parseErrorBoolean, endpointXML, "")
	if err == nil {
		t.Fatal("No error when parse wrong XML response.")
	}
	if !strings.Contains(err.Error(), "cannot parse XML RPC response") {
		t.Fatal("Unexpected error:", err)
	}
	if res != nil {
		t.Fatal("Method Call returns result when parse wrong XML response.")
	}
}

func Test_parseError_oneChildTag(t *testing.T) {
	res, err := MakeCallAndCreateRecord(t, parseErrorOnechildtag, endpointXML, "")
	if err == nil {
		t.Fatal("No error when parse wrong XML response.")
	}
	if !strings.Contains(err.Error(), "cannot parse XML RPC response") {
		t.Fatal("Unexpected error:", err)
	}
	if res != nil {
		t.Fatal("Method Call returns result when parse wrong XML response.")
	}
}

func Test_parseError_struct_noName(t *testing.T) {
	res, err := MakeCallAndCreateRecord(t, parseErrorStructNoname, endpointXML, "")
	if err == nil {
		t.Fatal("No error when parse wrong XML response.")
	}
	if !strings.Contains(err.Error(), "cannot parse XML RPC response") {
		t.Fatal("Unexpected error:", err)
	}
	if res != nil {
		t.Fatal("Method Call returns result when parse wrong XML response.")
	}
}

func Test_parseError_struct_noValue(t *testing.T) {
	res, err := MakeCallAndCreateRecord(t, parseErrorStructNovalue, endpointXML, "")
	if err == nil {
		t.Fatal("No error when parse wrong XML response.")
	}
	if !strings.Contains(err.Error(), "cannot parse XML RPC response") {
		t.Fatal("Unexpected error:", err)
	}
	if res != nil {
		t.Fatal("Method Call returns result when parse wrong XML response.")
	}
}

func Test_parseError_struct_noMember(t *testing.T) {
	res, err := MakeCallAndCreateRecord(t, parseErrorStructNomember, endpointXML, "")
	if err == nil {
		t.Fatal("No error when parse wrong XML response.")
	}
	if !strings.Contains(err.Error(), "cannot parse XML RPC response") {
		t.Fatal("Unexpected error:", err)
	}
	if res != nil {
		t.Fatal("Method Call returns result when parse wrong XML response.")
	}
}

func Test_parseError_struct_element(t *testing.T) {
	res, err := MakeCallAndCreateRecord(t, parseErrorStructElement, endpointXML, "")
	if err == nil {
		t.Fatal("No error when parse wrong XML response.")
	}
	if !strings.Contains(err.Error(), "cannot parse XML RPC response") {
		t.Fatal("Unexpected error:", err)
	}
	if res != nil {
		t.Fatal("Method Call returns result when parse wrong XML response.")
	}
}

func Test_parseError_struct_multipleMembers(t *testing.T) {
	res, err := MakeCallAndCreateRecord(t, parseErrorStructMultipleMembers, endpointXML, "")
	if err == nil {
		t.Fatal("No error when parse wrong XML response.")
	}
	if !strings.Contains(err.Error(), "cannot parse XML RPC response") {
		t.Fatal("Unexpected error:", err)
	}
	if res != nil {
		t.Fatal("Method Call returns result when parse wrong XML response.")
	}
}

func Test_parseError_struct_oneChildTag(t *testing.T) {
	res, err := MakeCallAndCreateRecord(t, parseErrorStructOnechildtag, endpointXML, "")
	if err == nil {
		t.Fatal("No error when parse wrong XML response.")
	}
	if !strings.Contains(err.Error(), "cannot parse XML RPC response") {
		t.Fatal("Unexpected error:", err)
	}
	if res != nil {
		t.Fatal("Method Call returns result when parse wrong XML response.")
	}
}

func Test_parseError_wrongTag(t *testing.T) {
	res, err := MakeCallAndCreateRecord(t, parseErrorWrongtag, endpointXML, "")
	if err == nil {
		t.Fatal("No error when parse wrong XML response.")
	}
	if !strings.Contains(err.Error(), "cannot parse XML RPC response") {
		t.Fatal("Unexpected error:", err)
	}
	if res != nil {
		t.Fatal("Method Call returns result when parse wrong XML response.")
	}
}

func Test_parseFault(t *testing.T) {
	res, err := MakeCallAndCreateRecord(t, parseFaultError, endpointXML, "")
	if err == nil {
		t.Fatal("No error when parse wrong XML response.")
	}
	if !strings.Contains(err.Error(), "cannot parse XML RPC response") {
		t.Fatal("Unexpected error:", err)
	}
	if res != nil {
		t.Fatal("Method Call returns result when parse wrong XML response.")
	}
}

func Test_parseFault_nameNil(t *testing.T) {
	res, err := MakeCallAndCreateRecord(t, parseFaultName, endpointXML, "")
	if err == nil {
		t.Fatal("No error when parse wrong XML response.")
	}
	if !strings.Contains(err.Error(), "cannot parse XML RPC response") {
		t.Fatal("Unexpected error:", err)
	}
	if res != nil {
		t.Fatal("Method Call returns result when parse wrong XML response.")
	}
}

func Test_parseFault_members(t *testing.T) {
	res, err := MakeCallAndCreateRecord(t, parseFaultMembers, endpointXML, "")
	if err == nil {
		t.Fatal("No error when parse wrong XML response.")
	}
	if !strings.Contains(err.Error(), "cannot parse XML RPC response") {
		t.Fatal("Unexpected error:", err)
	}
	if res != nil {
		t.Fatal("Method Call returns result when parse wrong XML response.")
	}
}
