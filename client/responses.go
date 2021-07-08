package client

import (
	"net/http"
	"time"
)

type LoanOfferResponse struct {
	LoanOffers []struct {
		Type               string `json:"type"`
		Amount             int    `json:"amount"`
		Rate               int    `json:"rate"`
		TermInDays         int    `json:"termInDays"`
		CollateralRequired bool   `json:"collateralRequired"`
	} `json:"loans"`
}

type SpaceTradersClient struct {
	Token string
	client http.Client
}

type StatusResponse struct {
	Status string
}

type AccountResponse struct {
	User struct {
		Credits        int       `json:"credits"`
		JoinedAt       time.Time `json:"joinedAt"`
		ShipCount      int       `json:"shipCount"`
		StructureCount int       `json:"structureCount"`
		Username       string    `json:"username"`
	} `json:"user"`
}

type AcceptLoanResponse struct {
	Credits int `json:"credits"`
	Loan    struct {
		Due             time.Time `json:"due"`
		Id              string    `json:"id"`
		RepaymentAmount int       `json:"repaymentAmount"`
		Status          string    `json:"status"`
		Type            string    `json:"type"`
	} `json:"loan"`
}

type MyLoansResponse struct {
	Loans []struct {
		Type               string `json:"type"`
		Amount             int    `json:"amount"`
		Rate               int    `json:"rate"`
		TermInDays         int    `json:"termInDays"`
		CollateralRequired bool   `json:"collateralRequired"`
	} `json:"loans"`
}

type AvailableShipsResponse struct {
	ShipListings []struct {
		Type              string `json:"type"`
		Class             string `json:"class"`
		MaxCargo          int    `json:"maxCargo"`
		LoadingSpeed      int    `json:"loadingSpeed"`
		Speed             int    `json:"speed"`
		Manufacturer      string `json:"manufacturer"`
		Plating           int    `json:"plating"`
		Weapons           int    `json:"weapons"`
		PurchaseLocations []struct {
			System   string `json:"system"`
			Location string `json:"location"`
			Price    int    `json:"price"`
		} `json:"purchaseLocations"`
		RestrictedGoods []string `json:"restrictedGoods,omitempty"`
	} `json:"shipListings"`
}

type BuyShipResponse struct {
	Credits int `json:"credits"`
	Ship    struct {
		Cargo          []interface{} `json:"cargo"`
		Class          string        `json:"class"`
		Id             string        `json:"id"`
		LoadingSpeed   int           `json:"loadingSpeed"`
		Location       string        `json:"location"`
		Manufacturer   string        `json:"manufacturer"`
		MaxCargo       int           `json:"maxCargo"`
		Plating        int           `json:"plating"`
		SpaceAvailable int           `json:"spaceAvailable"`
		Speed          int           `json:"speed"`
		Type           string        `json:"type"`
		Weapons        int           `json:"weapons"`
		X              int           `json:"x"`
		Y              int           `json:"y"`
	} `json:"ship"`
}

type MyShipsResponse struct {
	Ships []struct {
		Cargo          []interface{} `json:"cargo"`
		Class          string        `json:"class"`
		Id             string        `json:"id"`
		LoadingSpeed   int           `json:"loadingSpeed"`
		Location       string        `json:"location"`
		Manufacturer   string        `json:"manufacturer"`
		MaxCargo       int           `json:"maxCargo"`
		Plating        int           `json:"plating"`
		SpaceAvailable int           `json:"spaceAvailable"`
		Speed          int           `json:"speed"`
		Type           string        `json:"type"`
		Weapons        int           `json:"weapons"`
		X              int           `json:"x"`
		Y              int           `json:"y"`
	} `json:"ships"`
}

type BuyGoodResponse struct {
	User struct {
		Credits int `json:"credits"`
	} `json:"user"`
	Order struct {
		Good         string `json:"good"`
		PricePerUnit int    `json:"pricePerUnit"`
		Quantity     int    `json:"quantity"`
		Total        int    `json:"total"`
	} `json:"order"`
	Ship struct {
		Cargo []struct {
			Good        string `json:"good"`
			Quantity    int    `json:"quantity"`
			TotalVolume int    `json:"totalVolume"`
		} `json:"cargo"`
		Class          string `json:"class"`
		Id             string `json:"id"`
		Location       string `json:"location"`
		Manufacturer   string `json:"manufacturer"`
		MaxCargo       int    `json:"maxCargo"`
		Plating        int    `json:"plating"`
		SpaceAvailable int    `json:"spaceAvailable"`
		Speed          int    `json:"speed"`
		Type           string `json:"type"`
		Weapons        int    `json:"weapons"`
		X              int    `json:"x"`
		Y              int    `json:"y"`
	} `json:"ship"`
}

type MarketplaceResponse struct {
	Marketplace []struct {
		Symbol               string `json:"symbol"`
		VolumePerUnit        int    `json:"volumePerUnit"`
		PricePerUnit         int    `json:"pricePerUnit"`
		Spread               int    `json:"spread"`
		PurchasePricePerUnit int    `json:"purchasePricePerUnit"`
		SellPricePerUnit     int    `json:"sellPricePerUnit"`
		QuantityAvailable    int    `json:"quantityAvailable"`
	} `json:"marketplace"`
}