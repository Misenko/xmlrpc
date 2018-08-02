package xmlrpc

import (
	"testing"

	"context"
	"net/http"
	"github.com/dnaeon/go-vcr/recorder"
	"strings"
)

const (
	endpointCorrect = "http://127.0.0.1:8000/RPC2"
	endpointWrong   = "http://127.0.0.1:8000/qwerty"
	endpointInvalid = "@#$%^&"
	endpointEmpty   = ""

	missingArgs     = "records/missing_args"
	wrongArgsType   = "records/wrong_args_type"
	wrongMethodName = "records/wrong_method_name"
	emptyMethodName = "records/empty_method_name"
	wrongEndpoint   = "records/wrong_endpoint"
)

func Test_NewClient(t *testing.T) {
	client := NewClient(endpointEmpty)
	if client == nil {
		t.Error("Wrong NewClient method, client is nil.")
	}
}

func MakeCallAndCreateRecord(ctx context.Context, t *testing.T, recorderName string, endpoint string, methodName string, args ...interface{}) (*Result, error){
	// Start our recorder
	r, err := recorder.New(recorderName)
	if err != nil {
		t.Fatal(err)
	}
	defer r.Stop() // Make sure recorder is stopped once done with it

	// Create an HTTP client and inject our transport
	cl := &http.Client{
		Transport: r, // Inject as transport!
	}

	// Create XML-RPC client and set HTTP client
	client := NewClient(endpoint)
	if client == nil {
		t.Fatal("Unable to create xml-rpc client.")
	}
	client.client = cl

	// Make call
	return client.Call(ctx, methodName, args...)
}

func Test_Call_preparePayload_nilArgs(t *testing.T) {
	// test expects fail before connection to the server, no record needed
	res, err := MakeCallAndCreateRecord(context.TODO(), t, "", endpointCorrect, "pow", nil)
	if err == nil {
		t.Fatal("No error when args contains nil.")
	}
	if !strings.Contains(err.Error(), "payload preparation failed") {
		t.Fatal("Unexpected error:", err)
	}
	if res != nil {
		t.Fatal("Method Call returns result when args contains nil.")
	}
}

func Test_Call_preparePayload_missingArgs(t *testing.T) {
	res, err := MakeCallAndCreateRecord(context.TODO(), t, missingArgs, endpointCorrect, "pow")
	if err == nil {
		t.Error("No error when args are missing.")
	}
	if !strings.Contains(err.Error(), "pow expected at least 2 arguments") {
		t.Fatal("Unexpected error:", err)
	}
	if res != nil {
		t.Error("Method Call returns result when args are missing.")
	}
}

func Test_Call_preparePayload_wrongArgsType(t *testing.T) {
	res, err := MakeCallAndCreateRecord(context.TODO(), t, wrongArgsType, endpointCorrect, "pow", "pizza", "lasagne")
	if err == nil {
		t.Error("No error when args are missing.")
	}
	if !strings.Contains(err.Error(), "unsupported operand type") {
		t.Fatal("Unexpected error:", err)
	}
	if res != nil {
		t.Error("Method Call returns result when args are missing.")
	}
}

func Test_Call_preparePayload_wrongMethodName(t *testing.T) {
	res, err := MakeCallAndCreateRecord(context.TODO(), t, wrongMethodName, endpointCorrect,  "pancake")
	if err == nil {
		t.Error("No error when method name is wrong.")
	}
	if !strings.Contains(err.Error(), "method \"pancake\" is not supported") {
		t.Fatal("Unexpected error:", err)
	}
	if res != nil {
		t.Error("Method Call returns result when method name is wrong.")
	}
}

func Test_Call_preparePayload_emptyMethodName(t *testing.T) {
	res, err := MakeCallAndCreateRecord(context.TODO(), t, emptyMethodName, endpointCorrect,  "")
	if err == nil {
		t.Error("No error when method name is empty.")
	}
	if !strings.Contains(err.Error(), "method \"\" is not supported") {
		t.Fatal("Unexpected error:", err)
	}
	if res != nil {
		t.Error("Method Call returns result when method name is empty.")
	}
}

func Test_Call_makeRequest_wrongEndpoint(t *testing.T) {
	res, err := MakeCallAndCreateRecord(context.TODO(), t, wrongEndpoint, endpointWrong, "pow", 2, 9)
	if err == nil {
		t.Fatal("No error when endpoint is wrong.")
	}
	if !strings.Contains(err.Error(), "code 404") {
		t.Fatal("Unexpected error:", err)
	}
	if res != nil {
		t.Fatal("Method Call returns result when endpoint is wrong.")
	}
}

func Test_Call_makeRequest_invalidEndpoint(t *testing.T) {
	// test expects fail before connection to the server, no record needed
	res, err := MakeCallAndCreateRecord(context.TODO(), t, "", endpointInvalid, "pow", 2, 9)
	if err == nil {
		t.Fatal("No error when endpoint is invalid.")
	}
	if !strings.Contains(err.Error(), "request preparation failed") {
		t.Fatal("Unexpected error:", err)
	}
	if res != nil {
		t.Fatal("Method Call returns result when endpoint is wrong.")
	}
}

func Test_Call_makeRequest_emptyEndpoint(t *testing.T) {
	// test expects fail before connection to the server, no record needed
	res, err := MakeCallAndCreateRecord(context.TODO(), t, "", endpointEmpty, "pow", 2, 9)
	if err == nil {
		t.Fatal("No error when endpoint is empty.")
	}
	if !strings.Contains(err.Error(), "unsupported protocol scheme") {
		t.Fatal("Unexpected error:", err)
	}
	if res != nil {
		t.Fatal("Method Call returns result when endpoint is empty.")
	}
}