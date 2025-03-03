package samokat

import (
	"SamokatParser/pkg/utils"
	"bytes"
	"encoding/json"
	"fmt"
	"io"
	"net/http"
	"net/url"
	"strings"
)

type Client struct {
	httpClient  *http.Client
	bearerToken string
	userID      string
	deviceID    string
}

func NewClient() *Client {
	transport := &http.Transport{
		DialTLS: samokatDial,
	}

	return &Client{
		httpClient: &http.Client{
			Transport: transport,
		},
		userID:   utils.RandomNumber(10),
		deviceID: utils.GetMD5Hash(utils.RandomNumber(10)),
	}
}

func (c *Client) SetProxy(proxyURL string) error {
	parsedURL, err := url.Parse(proxyURL)
	if err != nil {
		return err
	}

	c.httpClient.Transport = &http.Transport{
		Proxy:   http.ProxyURL(parsedURL),
		DialTLS: samokatDial,
	}

	return nil
}

func (c *Client) setBasicHeaders(req *http.Request) {
	req.Header.Set("systemversion", "7.1.2")
	req.Header.Set("x-user-id", c.userID)
	req.Header.Set("x-user-type", "anonymous")
	req.Header.Set("x-application-platform", "android")
	req.Header.Set("x-application-version", "3.174.0")
	req.Header.Set("User-Agent", "smartspacestoreapp/3.174.0 (build: 29675; device: google Pixel 2; OS: Android 7.1.2)")
	req.Header.Set("deviceid", c.deviceID)
	req.Header.Set("Accept", "application/json, text/plain, */*")

	if c.bearerToken != "" {
		req.Header.Set("Authorization", "Bearer "+c.bearerToken)
	}
}

func (c *Client) DoRequest(req *http.Request) (*http.Response, error) {
	c.setBasicHeaders(req)

	return c.httpClient.Do(req)

}

func (c *Client) SetBearer(bearer string) {
	c.bearerToken = bearer
}

func (c *Client) GetOauthToken() (GetOAuthTokenResp, error) {
	URL := "https://api.samokat.ru/showcase/oauth/token"
	form := url.Values{}
	form.Set("grant_type", "urn:oauth:grant-type:anonymous")
	req, _ := http.NewRequest(
		http.MethodPost,
		URL,
		strings.NewReader(form.Encode()),
	)
	req.Header.Set("Content-Type", "application/x-www-form-urlencoded")

	resp, err := c.DoRequest(req)
	if err != nil {
		return GetOAuthTokenResp{}, err
	}

	defer resp.Body.Close()
	respBody, err := io.ReadAll(resp.Body)
	if err != nil {
		return GetOAuthTokenResp{}, err
	}
	fmt.Println(string(respBody))
	var getOAuthTokenRespnResp GetOAuthTokenResp
	err = json.Unmarshal(respBody, &getOAuthTokenRespnResp)
	if err != nil {
		return GetOAuthTokenResp{}, err
	}

	return getOAuthTokenRespnResp, nil
}

func (c *Client) AddAddress(address Address) (bool, error) {
	URL := "https://api.samokat.ru/showcase/v2/users/profile/addresses"
	jsonReqBody, err := json.Marshal(address)
	if err != nil {
		return false, err
	}
	req, _ := http.NewRequest(
		http.MethodPost,
		URL,
		bytes.NewReader(jsonReqBody),
	)

	req.Header.Set("Content-Type", "application/json")

	resp, err := c.DoRequest(req)
	if err != nil {
		return false, err
	}

	return resp.StatusCode == 200, nil
}

func (c *Client) GetShowcases(lat, lon float64) (Showcases, error) {
	URL := fmt.Sprintf(
		"https://api.samokat.ru/showcase/showcases/list?lat=%.6f&lon=%.6f",
		lat, lon)

	req, _ := http.NewRequest(
		http.MethodGet,
		URL,
		nil,
	)

	resp, err := c.DoRequest(req)
	if err != nil {
		return Showcases{}, err
	}

	defer resp.Body.Close()
	respBody, err := io.ReadAll(resp.Body)
	if err != nil {
		return Showcases{}, err
	}

	var showcases Showcases
	err = json.Unmarshal(respBody, &showcases)
	if err != nil {
		return Showcases{}, err
	}

	return showcases, nil
}

func (c *Client) GetShowcaseMain(showcaseUUID string) (ShowcaseMainRespBody, error) {
	URL := fmt.Sprintf(
		"https://api.samokat.ru/showcase/v2/showcases/%s/main",
		showcaseUUID)

	req, _ := http.NewRequest(
		http.MethodGet,
		URL,
		nil,
	)

	resp, err := c.DoRequest(req)
	if err != nil {
		return ShowcaseMainRespBody{}, err
	}

	defer resp.Body.Close()
	respBody, err := io.ReadAll(resp.Body)
	if err != nil {
		return ShowcaseMainRespBody{}, err
	}

	var showcaseMainRespBody ShowcaseMainRespBody
	err = json.Unmarshal(respBody, &showcaseMainRespBody)
	if err != nil {
		return ShowcaseMainRespBody{}, err
	}

	return showcaseMainRespBody, nil
}

func (c *Client) GetShowcaseCategoryGoods(showcaseUUID, categoryUUID string) (GetShowcaseCategoryResp, error) {
	URL := fmt.Sprintf(
		"https://api.samokat.ru/showcase/v2/showcases/%s/categories/%s?withPagination=false",
		showcaseUUID,
		categoryUUID)

	req, _ := http.NewRequest(
		http.MethodGet,
		URL,
		nil,
	)

	resp, err := c.DoRequest(req)
	if err != nil {
		return GetShowcaseCategoryResp{}, err
	}

	defer resp.Body.Close()
	respBody, err := io.ReadAll(resp.Body)
	if err != nil {
		return GetShowcaseCategoryResp{}, err
	}

	var getShowcaseCategoryResp GetShowcaseCategoryResp
	err = json.Unmarshal(respBody, &getShowcaseCategoryResp)
	if err != nil {
		return getShowcaseCategoryResp, nil
	}

	return getShowcaseCategoryResp, nil
}
