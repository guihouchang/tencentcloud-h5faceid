package faceid

import (
	"context"
	"testing"
)

var ctx = context.Background()

func TestGetAccessToken(t *testing.T) {
	client, err := NewFaceClient(
		WithAppID(""),
		WithSecret(""),
	)
	if err != nil {
		t.Fatal(err)
	}

	token, err := client.GetAccessToken(ctx)
	if err != nil {
		t.Fatal(err)
	}

	t.Log(token)
}

func TestGetAPITicket(t *testing.T) {
	client, err := NewFaceClient(
		WithAppID(""),
		WithSecret(""),
	)
	if err != nil {
		t.Fatal(err)
	}

	ticket, err := client.GetAPITicket(ctx)
	if err != nil {
		t.Fatal(err)
	}

	t.Log(ticket)
}

func TestGetNonceTicket(t *testing.T) {
	client, err := NewFaceClient(
		WithAppID(""),
		WithSecret(""),
	)

	if err != nil {
		t.Fatal(err)
	}

	ticket, err := client.GetNonceTicket(ctx, "1")
	if err != nil {
		t.Fatal(err)
	}
	t.Log(ticket)
}

func TestGetH5FaceID(t *testing.T) {
	client, err := NewFaceClient(
		WithAppID(""),
		WithSecret(""),
	)

	if err != nil {
		t.Fatal(err)
	}

	request := &H5FaceIDRequest{
		IDNO:    "",
		Name:    "",
		OrderNO: "123456789",
		UserID:  "1",
	}

	response, err := client.GetFaceID(ctx, request)
	if err != nil {
		t.Fatal(err)
	}

	t.Log(response)
}
func TestQueryFaceRecord(t *testing.T) {
	client, err := NewFaceClient(
		WithAppID(""),
		WithSecret(""),
	)

	if err != nil {
		t.Fatal(err)
	}

	request := &H5FaceRecordRequest{
		GetFile: "1",
		OrderNo: "123456789",
	}

	response, err := client.QueryFaceRecord(ctx, request)
	if err != nil {
		t.Fatal(err)
	}

	t.Log(response)
}
