package client

import (
	"encoding/json"
	"errors"
	"fmt"
	"io"
	"net/http"
	"net/url"
	"strconv"
	"strings"
)

func createUrl(path string) string {
	val := url.URL{
		Scheme: "https",
		Host:   spacetradersBaseUrl,
		Path:   path,
	}

	return val.String()
}

func NewSpaceTradersClient(token string) *SpaceTradersClient {
	client := http.Client{}

	return &SpaceTradersClient{token, client}
}

func constructRequest(path string, method string, parameters map[string]string)  (*http.Request, error) {
	req, err := http.NewRequest(method, createUrl(path), nil)
	if err != nil {
		return req, fmt.Errorf("could not create request object for path %s: %w", path, err)
	}

	query := req.URL.Query()
	query.Add("token", SpacetradersToken)
	for k, v := range parameters {
		query.Add(k, v)
	}
	req.URL.RawQuery = query.Encode()

	return req, nil
}

func (c SpaceTradersClient) sendRequestWithParameters(path string, method string, obj interface{}, parameters map[string]string) error {
	req, err := constructRequest(path, method, parameters)
	if err != nil {
		return err
	}

	rsp, err := c.client.Do(req)
	if err != nil {
		return fmt.Errorf("failed to execute request to space traders API %s: %w", req.URL, err)
	}

	defer func(Body io.ReadCloser) {
		err := Body.Close()
		if err != nil {
			fmt.Println(fmt.Errorf("could not close response body buffer for url %s: %w", req.URL, err))
		}
	}(rsp.Body)

	if rsp.StatusCode < http.StatusOK || rsp.StatusCode >= http.StatusMultipleChoices {
		return fmt.Errorf("recieved not-OK status code %d when sending request to %s", rsp.StatusCode, req.URL)
	}

	responseBytes, err := io.ReadAll(rsp.Body)
	responseString := string(responseBytes)
	if err != nil {
		return fmt.Errorf("could not read from spacetraders api %s: %w", req.URL, err)
	}

	err = json.Unmarshal([]byte(responseString), &obj)
	if err != nil {
		return fmt.Errorf("could not decode response from spacetraders api %s: %w", req.URL, err)
	}

	return nil
}

func (c SpaceTradersClient) getWithParams(path string, obj interface{}, parameters map[string]string) error {
	return c.sendRequestWithParameters(path, "GET", &obj, parameters)
}

func (c SpaceTradersClient) postWithParams(path string, obj interface{}, parameters map[string]string) error {
	return c.sendRequestWithParameters(path, "POST", &obj, parameters)
}

func (c SpaceTradersClient) get(path string, obj interface{}) error {
	return c.getWithParams(path, &obj, map[string]string{})
}

func (c SpaceTradersClient) post(path string, obj interface{}) error {
	return c.postWithParams(path, &obj, map[string]string{})
}

func (c SpaceTradersClient) Connected() bool {
	status, err := c.Status()
	return err == nil && strings.Contains(status.Status, "currently online")
}

func (c SpaceTradersClient) Account() (AccountResponse, error) {
	var rspObj AccountResponse
	err := c.get("my/account", &rspObj)
	return rspObj, err
}

func (c SpaceTradersClient) Status() (StatusResponse, error) {
	var rspObj StatusResponse
	err := c.get("game/status", &rspObj)
	return rspObj, err
}

func (c SpaceTradersClient) AvailableLoans() (LoanOfferResponse, error) {
	var rspObj LoanOfferResponse
	err := c.get("types/loans", &rspObj)
	if err != nil {
		return rspObj, fmt.Errorf("failed to get loan result: %w", err)
	}

	return rspObj, nil
}

func (c SpaceTradersClient) AcceptLoan(loanType int) (MyLoansResponse, error) {
	var loanTypeStr string
	var rspObj MyLoansResponse

	switch loanType {
	case StartupLoan:
		loanTypeStr = "STARTUP"
	default:
		return rspObj, errors.New(fmt.Sprintf("Unrecognized loan type %d", loanType))
	}

	err := c.postWithParams("my/loans", &rspObj, map[string]string{"type": loanTypeStr})

	return rspObj, err
}

func (c SpaceTradersClient) MyLoans() (MyLoansResponse, error) {
	var rspObj MyLoansResponse
	err := c.get("my/loans", &rspObj)
	return rspObj, err
}

func (c SpaceTradersClient) AvailableShips(system string) (AvailableShipsResponse, error) {
	var rspObj AvailableShipsResponse
	err := c.get(fmt.Sprintf("systems/%s/ship-listings", system), &rspObj)
	return rspObj, err
}

func (c SpaceTradersClient) BuyShip(shipType string, shipLocation string) (BuyShipResponse, error) {
	var rspObj BuyShipResponse
	err := c.postWithParams("my/ships", &rspObj, map[string]string{"type": shipType, "location": shipLocation} )
	return rspObj, err
}

func (c SpaceTradersClient) MyShips() (MyShipsResponse, error) {
	var rspObj MyShipsResponse
	err := c.get("my/ships", &rspObj)
	return rspObj, err
}

func (c SpaceTradersClient) BuyGood(shipId string, good string, quantity int) (BuyGoodResponse, error) {
	var rspObj BuyGoodResponse
	err := c.postWithParams("my/purchase-orders", &rspObj,
		map[string]string{"shipId": shipId, "quantity": strconv.Itoa(quantity), "good": good} )
	return rspObj, err
}

func (c SpaceTradersClient) Marketplace(location string) (MarketplaceResponse, error) {
	var rspObj MarketplaceResponse
	err := c.get(fmt.Sprintf("locations/%s/marketplace", location), &rspObj)
	return rspObj, err
}

