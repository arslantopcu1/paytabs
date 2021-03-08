package paytabs

import (
	"bytes"
	"encoding/json"
	//"fmt"
	"io/ioutil"
	"net/http"
	"net/url"
)

type payPageResponse struct {
	Result       string `json:"result,ommitempty"`
	ResponseCode string `json:"response_code,ommitempty"`
	PaymentURL   string `json:"payment_url,ommitempty"`
}

type generalResponse struct {
	Result       string `json:"result,ommitempty"`
	ResponseCode string `json:"response_code,ommitempty"`
}

type verifyResponse struct {
	Result        string `json:"result,ommitempty"`
	ResponseCode  string `json:"response_code,ommitempty"`
	Amount        string `json:"amount,ommitempty"`
	PTInvoiceID   string `json:"pt_invoice_id,ommitempty"`
	Currency      string `json:"currency,ommitempty"`
	ReferenceNo   string `json:"reference_no,ommitempty"`
	TransactionID string `json:"transaction_id,ommitempty"`
}

/**
Create PayPage API
**/
func CretaePayPage(data map[string]string) (payPageResponse, error) {
	location := "https://www.paytabs.com/apiv2/create_pay_page"
	resp, err := sendRequest(location, data)

	if err != nil {
		return payPageResponse{}, err //Error Happened
	}

	reader := bytes.NewReader(resp)
	_paypageResponse := payPageResponse{}
	json.NewDecoder(reader).Decode(&_paypageResponse)

	return _paypageResponse, nil
}

/**
Validate Secret key
**/
func ValidateSecretKey(data map[string]string) (generalResponse, error) {
	location := "https://www.paytabs.com/apiv2/validate_secret_key"
	resp, err := sendRequest(location, data)
	if err != nil {
		return generalResponse{}, err
	}

	reader := bytes.NewReader(resp)
	validateResponse := generalResponse{}
	json.NewDecoder(reader).Decode(&validateResponse)
	return validateResponse, nil
}

func VerifyPayment(data map[string]string) (verifyResponse, error) {
	location := "https://www.paytabs.com/apiv2/verify_payment"
	resp, err := sendRequest(location, data)
	if err != nil {
		return verifyResponse{}, err
	}

	reader := bytes.NewReader(resp)
	_verifyResponse := verifyResponse{}
	json.NewDecoder(reader).Decode(&_verifyResponse)
	return _verifyResponse, nil
}

/**
Post data to the server
**/
func sendRequest(link string, data map[string]string) ([]byte, error) {
	post_data := url.Values{}
	//Iterate the map
	for key, val := range data {
		post_data.Add(key, val)
	}

	resp, err := http.PostForm(
		link,
		post_data,
	)

	if err != nil {
		return nil, err
	}

	defer resp.Body.Close()
	body, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		return nil, err
	}

	return body, nil
}
