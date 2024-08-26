# FaceID SDK 使用说明

## 介绍

`faceid` 是一个用于与腾讯云H5人脸核验 KYC 服务进行集成的 Go SDK，支持获取 API Ticket、Access Token，发起 H5 人脸识别，查询人脸识别记录等功能。

## 安装

```bash
go get -u github.com/guihouchang/tencentcloud-h5faceid
```
## 快速开始
以下是如何使用 faceid SDK 的示例代码
```golang
package main

import (
	"context"
	"fmt"
	"github.com/guihouchang/tencentcloud-h5faceid"
)

func main() {
	// 初始化FaceID客户端
	client, err := faceid.NewFaceClient(
		faceid.WithAppID("your-app-id"),
		faceid.WithSecret("your-secret"),
	)
	if err != nil {
		fmt.Println("Error creating FaceID client:", err)
		return
	}

	ctx := context.Background()

	// 获取API Ticket
	apiTicket, err := client.GetAPITicket(ctx)
	if err != nil {
		fmt.Println("Error getting API Ticket:", err)
		return
	}
	fmt.Println("API Ticket:", apiTicket)

	// 发起H5人脸识别请求
	faceIDRequest := &faceid.H5FaceIDRequest{
		OrderNO: "order123",
		Name:    "张三",
		IDNo:    "123456789012345678",
		UserID:  "user123",
	}
	faceIDResponse, err := client.GetFaceID(ctx, faceIDRequest)
	if err != nil {
		fmt.Println("Error getting Face ID:", err)
		return
	}
	fmt.Println("Face ID Response:", faceIDResponse)

	// 查询人脸识别记录
	faceRecordRequest := &faceid.H5FaceRecordRequest{
		OrderNo: "order123",
	}
	faceRecordResponse, err := client.QueryFaceRecord(ctx, faceRecordRequest)
	if err != nil {
		fmt.Println("Error querying face record:", err)
		return
	}
	fmt.Println("Face Record Response:", faceRecordResponse)
}
```
## 功能说明
1. #### 初始化 Face 客户端

在使用任何功能之前，需要先初始化 Face 客户端：
```golang
client, err := faceid.NewFaceClient(
    faceid.WithAppID("your-app-id"),
    faceid.WithSecret("your-secret"),
)
```
* `WithAppID`: 设置小程序的 AppID。
* `WithSecret`: 设置小程序的 Secret。
* `WithVersion`: 可选，设置接口版本，默认为 1.0.0。
2. #### 获取 Access Token

通过 GetAccessToken 方法获取 Access Token，Access Token 是请求 API 时的授权凭证：
```go
accessToken, err := client.GetAccessToken(ctx)
```
3. #### 获取 API Ticket
通过 GetAPITicket 方法获取 API Ticket，API Ticket 是请求 H5 人脸识别页面时使用的凭证：
```go
apiTicket, err := client.GetAPITicket(ctx)
```
4. #### 发起 H5 人脸识别请求
通过 GetFaceID 方法发起 H5 人脸识别请求，请求参数包括订单号、姓名、身份证号、用户 ID 等：
```go
faceIDRequest := &faceid.H5FaceIDRequest{
    OrderNO: "order123",
    Name:    "张三",
    IDNo:    "123456789012345678",
    UserID:  "user123",
    // 其他可选参数
}
faceIDResponse, err := client.GetFaceID(ctx, faceIDRequest)
```
5. #### 查询人脸识别记录
通过 QueryFaceRecord 方法查询人脸识别记录，请求参数包括订单号：
```go
faceRecordRequest := &faceid.H5FaceRecordRequest{
    OrderNo: "order123",
    // 其他可选参数
}
faceRecordResponse, err := client.QueryFaceRecord(ctx, faceRecordRequest)
```
6. #### 获取 H5 人脸识别页面 URL
通过 GetFaceIDURL 方法获取 H5 人脸识别页面 URL，请求参数包括订单号、姓名、身份证号、用户 ID 等：
```go
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
```

## 配置选项
在初始化 Face 客户端时，您可以传递以下配置选项：
1. `WithAppID`: 设置申请的 AppID。
2. `WithSecret`: 设置申请的 Secret。
3. `WithVersion`: 可选，设置接口版本，默认为 1.0.0。

## 依赖
`gookit/goutil`: 提供了一些便捷的工具函数，如字符串、结构体、时间处理等。

## 贡献
贡献代码请先 Fork 本仓库，然后创建一个新分支，提交您的更改，并在提交前运行测试，确保所有测试都通过。最后，在您的更改提交后，创建一个 Pull Request，我们将尽快进行合并。
