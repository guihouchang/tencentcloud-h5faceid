package faceid

import (
	"bytes"
	"context"
	"crypto/sha1"
	"encoding/json"
	"fmt"
	"github.com/gookit/goutil/structs"
	"github.com/gookit/goutil/strutil"
	"github.com/gookit/goutil/timex"
	"io/ioutil"
	"net/http"
	"net/url"
	"sort"
	"sync"
	"time"
)

const accessTokenURL = "https://miniprogram-kyc.tencentcloudapi.com/api/oauth2/access_token"
const apiTicketURL = "https://miniprogram-kyc.tencentcloudapi.com/api/oauth2/api_ticket"
const getH5FaceIDURL = "https://miniprogram-kyc.tencentcloudapi.com/api/server/h5/geth5faceid"

const queryH5FaceRecordURL = "https://miniprogram-kyc.tencentcloudapi.com/api/v2/base/queryfacerecord"

const defaultOptiamlURL = "https://miniprogram-kyc.tencentcloudapi.com"
const defaultClientCredential = "client_credential"
const defaultType = "SIGN"

type Face struct {
	accessTokenTTL time.Time
	accessToken    string
	apiTicket      string
	apiTicketTTL   time.Time
	c              *config
	sync.RWMutex
}

type config struct {
	appID         string
	secret        string
	version       string
	grantType     string
	optimalDomain string
}

type Options func(c *config)

func defaultConfig() *config {
	return &config{
		version:       "1.0.0",
		grantType:     defaultClientCredential,
		optimalDomain: defaultOptiamlURL,
	}
}

func WithAppID(appID string) Options {
	return func(c *config) {
		c.appID = appID
	}
}

func WithSecret(secret string) Options {
	return func(c *config) {
		c.secret = secret
	}
}

func WithOptimalDomain(optimalDomain string) Options {
	return func(c *config) {
		c.optimalDomain = optimalDomain
	}
}

func WithVersion(version string) Options {
	return func(c *config) {
		c.version = version
	}
}
func httpPostRequest(ctx context.Context, endpoint string, data map[string]any, headers map[string]string) ([]byte, error) {
	jsonData, err := json.Marshal(data)
	if err != nil {
		return nil, err
	}

	req, err := http.NewRequest("POST", endpoint, bytes.NewBuffer(jsonData))
	if err != nil {
		return nil, err
	}

	req = req.WithContext(ctx)

	// 设置请求头
	for key, value := range headers {
		req.Header.Set(key, value)
	}

	client := &http.Client{}
	resp, err := client.Do(req)
	if err != nil {
		return nil, err
	}
	defer resp.Body.Close()

	return ioutil.ReadAll(resp.Body)
}

func httpGetRequest(ctx context.Context, endpoint string, params map[string]any, headers map[string]any) ([]byte, error) {
	// 构建 URL 和查询参数
	baseURL, err := url.Parse(endpoint)
	if err != nil {
		return nil, fmt.Errorf("invalid URL: %v", err)
	}

	query := baseURL.Query()
	for key, value := range params {
		query.Set(key, strutil.MustString(value))
	}
	baseURL.RawQuery = query.Encode()

	// 创建 HTTP 请求
	req, err := http.NewRequest("GET", baseURL.String(), nil)
	if err != nil {
		return nil, fmt.Errorf("failed to create request: %v", err)
	}

	// 设置请求头
	for key, value := range headers {
		req.Header.Set(key, strutil.MustString(value))
	}

	// 创建 HTTP 客户端
	client := http.DefaultClient
	req = req.WithContext(ctx)
	// 发送请求
	resp, err := client.Do(req)
	if err != nil {
		return nil, fmt.Errorf("request failed: %v", err)
	}
	defer resp.Body.Close()

	// 读取响应内容
	body, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		return nil, fmt.Errorf("failed to read response: %v", err)
	}

	return body, nil
}

func (f *Face) GetAPITicket(ctx context.Context) (string, error) {
	accessToken, err := f.GetAccessToken(ctx)
	if err != nil {
		return "", err
	}

	// 提前十分钟应该刷新token
	now := time.Now().Add(time.Minute * -10)
	if !f.apiTicketTTL.IsZero() && now.After(f.apiTicketTTL) {
		return f.apiTicket, nil
	}

	f.Lock()
	defer f.Unlock()

	request := APITicketRequest{
		APPID:       f.c.appID,
		AccessToken: accessToken,
		Type:        defaultType,
		Version:     f.c.version,
	}

	maps := structs.ToMap(request)
	bytes, err := httpGetRequest(ctx, apiTicketURL, maps, nil)
	if err != nil {
		return "", err
	}

	response := &APITicketResponse{}
	err = json.Unmarshal(bytes, response)
	if err != nil {
		return "", err
	}

	if response.Code != "0" {
		return "", fmt.Errorf("failed to get api ticket: %s", response.Msg)
	}

	if len(response.Tickets) == 0 {
		return "", fmt.Errorf("failed to get api ticket: %s", "tickets is empty")
	}

	ticket := response.Tickets[0]

	t, err := timex.FromString(ticket.ExpireTime, "20060102150405")
	if err != nil {
		return "", err
	}
	f.apiTicketTTL = t.Time
	f.apiTicket = ticket.Value
	return f.apiTicket, nil
}

func (f *Face) GetAccessToken(ctx context.Context) (string, error) {
	// 提前十分钟应该刷新token
	now := time.Now().Add(time.Minute * -10)
	if !f.accessTokenTTL.IsZero() && now.After(f.accessTokenTTL) {
		return f.accessToken, nil
	}

	// 获取accessToken非并发安全
	f.Lock()
	defer f.Unlock()
	// post 请求
	request := AccessTokenRequest{
		AppID:     f.c.appID,
		GrantType: f.c.grantType,
		Secret:    f.c.secret,
		Version:   f.c.version,
	}

	maps := structs.ToMap(request)
	resp, err := httpGetRequest(ctx, accessTokenURL, maps, nil)
	if err != nil {
		return "", err
	}

	response := &AccessTokenResponse{}
	err = json.Unmarshal(resp, response)
	if err != nil {
		return "", err
	}

	if response.Code != "0" {
		return "", fmt.Errorf("failed to get access token: %s", response.Msg)
	}
	t, err := timex.FromString(response.ExpireTime, "20060102150405")
	if err != nil {
		return "", err
	}
	f.accessToken = response.AccessToken
	f.accessTokenTTL = t.Time
	return f.accessToken, nil
}

func (f *Face) sign(keys []string, data map[string]any) (string, error) {
	sortList := make([]string, 0, len(keys))

	for _, key := range keys {
		if value, ok := data[key]; ok {
			sortList = append(sortList, strutil.MustString(value))
		} else {
			return "", fmt.Errorf("%s is empty", key)
		}
	}

	sort.Strings(sortList)
	strBeSign := ""
	for _, value := range sortList {
		strBeSign += value
	}

	// sha1加密
	hashStr := sha1.Sum([]byte(strBeSign))
	return fmt.Sprintf("%x", hashStr), nil
}

func (f *Face) GetFaceID(ctx context.Context, request *H5FaceIDRequest) (*H5FaceIDResponse, error) {
	ticket, err := f.GetAPITicket(ctx)
	if err != nil {
		return nil, err
	}

	maps := structs.ToMap(request)
	maps["webankAppId"] = f.c.appID
	maps["version"] = f.c.version
	maps["ticket"] = ticket
	// 根据 webankAppId,orderNo,name,idNo,userId,version,ticket进行自定排序后生成签名
	keys := []string{"webankAppId", "orderNo", "name", "idNo", "userId", "version", "ticket"}
	sign, err := f.sign(keys, maps)
	if err != nil {
		return nil, err
	}

	maps["sign"] = sign
	endpoint := fmt.Sprintf(getH5FaceIDURL+"?orderNo=%s", request.OrderNO)
	resp, err := httpPostRequest(ctx, endpoint, maps, map[string]string{
		"Content-Type": "application/json",
	})

	if err != nil {
		return nil, err
	}

	response := &H5FaceIDResponse{}
	err = json.Unmarshal(resp, response)
	if err != nil {
		return nil, err
	}

	if response.Code != "0" {
		return nil, fmt.Errorf("failed to get face id: %s", response.Msg)
	}

	return response, nil
}

func (f *Face) GetFaceWebURL(ctx context.Context, request *H5FaceURLRequest) (string, error) {
	nonceTicket, err := f.GetNonceTicket(ctx, request.UserID)
	if err != nil {
		return "", err
	}

	maps := structs.ToMap(request)
	maps["webankAppId"] = f.c.appID
	maps["version"] = f.c.version
	maps["nonce"] = strutil.RandomCharsV2(32)
	maps["ticket"] = nonceTicket
	keys := []string{"webankAppId", "orderNo", "h5faceId", "userId", "nonce", "ticket", "version"}
	signStr, err := f.sign(keys, maps)
	if err != nil {
		return "", err
	}
	maps["sign"] = signStr
	endpoint := f.c.optimalDomain + "/api/web/login"
	baseURL, err := url.Parse(endpoint)
	if err != nil {
		return "", fmt.Errorf("invalid URL: %v", err)
	}

	query := baseURL.Query()
	for key, value := range maps {
		query.Set(key, strutil.MustString(value))
	}
	baseURL.RawQuery = query.Encode()
	return baseURL.String(), nil
}

func (f *Face) QueryFaceRecord(ctx context.Context, request *H5FaceRecordRequest) (*H5FaceRecordResponse, error) {
	apiTicket, err := f.GetAPITicket(ctx)
	if err != nil {
		return nil, err
	}

	maps := structs.ToMap(request)
	maps["ticket"] = apiTicket
	maps["version"] = f.c.version
	maps["appId"] = f.c.appID
	maps["nonce"] = strutil.RandomCharsV2(32)
	// 根据指定keys appId, orderNo, nonceStr, version, ticket进行自定排序后生成签名
	keys := []string{"appId", "orderNo", "nonce", "version", "ticket"}
	sign, err := f.sign(keys, maps)
	if err != nil {
		return nil, err
	}

	maps["sign"] = sign

	endpoint := fmt.Sprintf(queryH5FaceRecordURL+"?orderNo=%s", request.OrderNo)
	resp, err := httpPostRequest(ctx, endpoint, maps, map[string]string{
		"Content-Type": "application/json",
	})

	if err != nil {
		return nil, err
	}

	response := &H5FaceRecordResponse{}
	err = json.Unmarshal(resp, response)
	if err != nil {
		return nil, err
	}

	return response, nil
}

func (f *Face) GetNonceTicket(ctx context.Context, userID string) (string, error) {
	accessToken, err := f.GetAccessToken(ctx)
	if err != nil {
		return "", err
	}

	request := NonceAPITicketRequest{
		APPID:       f.c.appID,
		AccessToken: accessToken,
		Type:        defaultType,
		UserID:      userID,
		Version:     f.c.version,
	}

	maps := structs.ToMap(request)
	bytes, err := httpGetRequest(ctx, apiTicketURL, maps, nil)
	if err != nil {
		return "", err
	}

	resp := &APITicketResponse{}
	err = json.Unmarshal(bytes, resp)
	if err != nil {
		return "", err
	}

	if resp.Code != "0" {
		return "", fmt.Errorf("failed to get nonce ticket: %s", resp.Msg)
	}

	if len(resp.Tickets) == 0 {
		return "", fmt.Errorf("failed to get nonce ticket: %s", "tickets is empty")
	}

	return resp.Tickets[0].Value, nil
}

func NewFaceClient(opts ...Options) (*Face, error) {
	c := defaultConfig()
	for _, opt := range opts {
		opt(c)
	}

	return &Face{
		c: c,
	}, nil
}
