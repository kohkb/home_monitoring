package gateway

import (
	"encoding/json"
	"io/ioutil"
	"net/http"
	"net/url"
	"os"
	"time"

	"github.com/joho/godotenv"
)

type NetatmoClient struct {
	BaseURL     string
	accessToken string
	HTTPClient  *http.Client
}

const (
	NetatmoBaseURL = "https://api.netatmo.com"
)

func NewNetatmoClient() *NetatmoClient {
	return &NetatmoClient{
		BaseURL:     NetatmoBaseURL,
		accessToken: newAccessToken(),
		HTTPClient: &http.Client{
			Timeout: time.Minute,
		},
	}
}

func newAccessToken() string {
	godotenv.Load(".env")

	values := url.Values{}
	values.Set("grant_type", "password")
	values.Set("client_id", os.Getenv("CLIENT_ID"))
	values.Set("scope", "read_station")
	values.Set("username", os.Getenv("USERNAME"))
	values.Set("password", os.Getenv("PASSWORD"))
	values.Set("client_secret", os.Getenv("CLIENT_SECRET"))

	resp, _ := http.PostForm(NetatmoBaseURL+"/oauth2/token", values)
	body, _ := ioutil.ReadAll(resp.Body)

	var resp_map map[string]interface{}
	err := json.Unmarshal([]byte(body), &resp_map)
	if err != nil {
		panic(err)
	}

	return resp_map["access_token"].(string)
}

func (c *NetatmoClient) GetCarbonDioxideConcentration() (int, error) {
	req, _ := http.NewRequest("GET", c.BaseURL+"/api/getstationsdata", nil)
	req.Header.Set("Authorization", "Bearer "+c.accessToken)

	resp, _ := c.HTTPClient.Do(req)
	body, _ := ioutil.ReadAll(resp.Body)

	var data map[string]interface{}
	err := json.Unmarshal([]byte(body), &data)
	if err != nil {
		return 0, err
	}

	return int(data["body"].(map[string]interface{})["devices"].([]interface{})[0].(map[string]interface{})["dashboard_data"].(map[string]interface{})["CO2"].(float64)), nil
}
