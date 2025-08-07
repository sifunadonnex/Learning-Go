package api

//coin balance params
type CoinBalanceParams struct {
	Username string
}

// CoinBalanceResponse represents the response structure for coin balance
type CoinBalanceResponse struct {
	Code    int
	Balance int64
}

//Error represents the error structure
type Error struct {
	Code    int
	Message string
}
