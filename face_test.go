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
		IDNO:    "352225198703222010",
		Name:    "桂后昌",
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
		OrderNo: "17243090053082",
	}

	response, err := client.QueryFaceRecord(ctx, request)
	if err != nil {
		t.Fatal(err)
	}

	t.Log(response)
}

func TestGetFaceWebURL(t *testing.T) {
	client, err := NewFaceClient(
		WithAppID(""),
		WithSecret(""),
	)

	if err != nil {
		t.Fatal(err)
	}

	requestFaceID := &H5FaceIDRequest{
		IDNO:    "xxxxxxxxx",
		Name:    "xxx",
		OrderNO: "1234567891`",
		UserID:  "1",
	}

	faceIDResp, err := client.GetFaceID(ctx, requestFaceID)
	if err != nil {
		t.Fatal(err)
	}

	requestH5FaceURL := &H5FaceURLRequest{
		OrderNo:      faceIDResp.Result.OrderNo,
		H5FaceID:     faceIDResp.Result.H5FaceID,
		URL:          "http://www.baidu.com",
		ResultType:   "1",
		UserID:       "1",
		From:         "browser",
		RedirectType: "1",
	}

	url, err := client.GetFaceWebURL(ctx, requestH5FaceURL)
	if err != nil {
		t.Fatal(err)
	}

	t.Log(url)
}
