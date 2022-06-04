package kraken

import (
	"cryptec/pkg/helpers"
	"crypto/hmac"
	"crypto/sha256"
	"crypto/sha512"
	"encoding/base64"
	"encoding/json"
	"errors"
	"fmt"
	"io/ioutil"
	"net/http"
	"net/url"
	"strings"
	"time"
)

const (
	apiURL     string = "https://api.kraken.com"
	userAgent  string = "cryptec-v0"
	apiVersion string = "0"
)

type Kraken struct {
	credentials helpers.Credentials
	client      *http.Client
}

// New kraken client
func NewKraken(key string, secret string) (kraken *Kraken) {
	return &Kraken{
		credentials: helpers.Credentials{
			SECRET:     secret,
			KEY:        key,
			PASSPHRASE: "",
		},
		client: &http.Client{},
	}
}

func (kraken *Kraken) getAccountBalance() (*AccountBalanceResponse, error) {
	values := url.Values{}

	response, err := kraken.callPrivate("Balance", values, &AccountBalanceResponse{})
	if err != nil {
		return nil, err
	}

	balanceResponse, _ := response.(AccountBalanceResponse)
	return &balanceResponse, nil

}

// Retrieve information about ledger entries. 50 results are returned at a time, the most recent by default.
// https://docs.kraken.com/rest/#operation/getLedgers
func (kraken *Kraken) getLedgersInfo(args getLedgersArgs) (*LedgersResponse, error) {
	values := url.Values{
		"aclass": []string{args.Aclass},
		"asset":  []string{args.Asset},
		"type":   []string{args.Type_},
		"start":  []string{fmt.Sprintf("%d", args.Start)},
		"end":    []string{fmt.Sprintf("%d", args.Start)},
		"ofs":    []string{fmt.Sprintf("%d", args.Offset)},
	}

	response, err := kraken.callPrivate("Ledgers", values, &LedgersResponse{})
	if err != nil {
		return nil, err
	}

	ledgersResponse, _ := response.(LedgersResponse)

	return &ledgersResponse, nil
}

func (kraken *Kraken) callPrivate(method string, values url.Values, responseType interface{}) (interface{}, error) {
	urlPath := fmt.Sprintf("/%s/private/%s", apiVersion, method)
	requestURL := apiURL + urlPath

	secret, err := base64.StdEncoding.DecodeString(kraken.credentials.SECRET)
	if err != nil {
		return nil, errors.New("failed to decode API secret")
	}

	values.Set("nonce", fmt.Sprintf("%d", time.Now().UnixNano()))

	signature := kraken.signature(urlPath, values, secret)

	header := http.Header{
		"API-Key":  []string{kraken.credentials.KEY},
		"API-Sign": []string{signature},
	}

	return kraken.executeRequest(requestURL, values, header, responseType)
}

func (kraken *Kraken) executeRequest(requestURL string, values url.Values, header http.Header, responseType interface{}) (interface{}, error) {
	request, err := http.NewRequest("POST", requestURL, strings.NewReader(values.Encode()))
	if err != nil {
		return nil, err
	}

	request.Header = header
	request.Header.Add("User-Agent", userAgent)

	response, err := kraken.client.Do(request)
	if err != nil {
		return nil, err
	}

	body, err := ioutil.ReadAll(response.Body)
	if err != nil {
		return nil, err
	}

	var jsonData Response
	if responseType != nil {
		jsonData.Result = responseType
	}

	err = json.Unmarshal(body, &jsonData)
	if err != nil {
		return nil, fmt.Errorf("could not Unmrshal response data field: %s", err)
	}

	if len(jsonData.Error) > 0 {
		return nil, fmt.Errorf("%s", jsonData.Error)
	}

	return jsonData.Result, nil
}

func (kraken Kraken) signature(urlPath string, values url.Values, secret []byte) (signature string) {
	sha := sha256.New()
	sha.Write([]byte(values.Get("nonce") + values.Encode()))
	shasum := sha.Sum(nil)

	mac := hmac.New(sha512.New, secret)
	mac.Write(append([]byte(urlPath), shasum...))
	macsum := mac.Sum(nil)
	return base64.StdEncoding.EncodeToString(macsum)
}
