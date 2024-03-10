package api

import (
	"bytes"
	"encoding/json"
	"main/internal/adapters"
	"main/internal/app"
	"main/internal/domain"
	"main/internal/handlers"
	"main/pkg"
	"main/testsutils"
	"net/http"
	"net/http/httptest"
	"testing"

	"github.com/gin-gonic/gin"
	"github.com/stretchr/testify/suite"
)

type ServerTestSuite struct {
	suite.Suite
	router *gin.Engine
}

func (suite *ServerTestSuite) SetupTest() {
	walletRepo, _ := adapters.NewWalletPostgresRepo(&adapters.WalletPostgresRepoDeps{ConnPool: testsutils.GetDbPool()})
	walletService := app.NewWalletService(walletRepo, pkg.NewUUIDGenerator())
	walletFacade := handlers.NewWalletHandlersFacade(walletService)
	suite.router = getRouter(walletFacade)
}

func (s *ServerTestSuite) TestShouldReturnOkStatus() {
	req, _ := http.NewRequest("GET", "/api/v1/health", nil)
	w := httptest.NewRecorder()
	s.router.ServeHTTP(w, req)
	s.Equal(http.StatusOK, w.Code)
	body := w.Body.String()
	json := `{"status":"ok"}`
	s.Equal(json, body)
}

func (s *ServerTestSuite) TestShoudlReturn200WhenCreatingNewWallet() {
	req, _ := http.NewRequest("POST", "/api/v1/wallets", nil)
	w := httptest.NewRecorder()
	s.router.ServeHTTP(w, req)
	s.Equal(http.StatusCreated, w.Code)
}

func (s *ServerTestSuite) TestShoudlReturn200WhenDepositingToWallet() {
	req, _ := http.NewRequest("POST", "/api/v1/wallets", nil)
	w := httptest.NewRecorder()
	s.router.ServeHTTP(w, req)
	s.Equal(http.StatusCreated, w.Code)
	var wallet *domain.Wallet
	json.Unmarshal(w.Body.Bytes(), &wallet)

	depositBody := `{"amount":"100.00"}`
	body := bytes.NewBufferString(depositBody)
	depositRequest, _ := http.NewRequest("PUT", "/api/v1/wallets/"+wallet.ID+"/balance/deposit", body)
	depositRecorder := httptest.NewRecorder()
	s.router.ServeHTTP(depositRecorder, depositRequest)
	s.Equal(http.StatusOK, depositRecorder.Code)

	balanceRequest, _ := http.NewRequest("GET", "/api/v1/wallets/"+wallet.ID+"/balance", nil)
	balanceRecorder := httptest.NewRecorder()
	s.router.ServeHTTP(balanceRecorder, balanceRequest)
	s.Equal(http.StatusOK, balanceRecorder.Code)
	balance := `{"balance":"100.00"}`
	s.Equal(balance, balanceRecorder.Body.String())
}

func (s *ServerTestSuite) TestShoudlReturn200WhenGettingBalance() {
	req, _ := http.NewRequest("POST", "/api/v1/wallets", nil)
	w := httptest.NewRecorder()
	s.router.ServeHTTP(w, req)
	s.Equal(http.StatusCreated, w.Code)
	var wallet *domain.Wallet
	json.Unmarshal(w.Body.Bytes(), &wallet)

	balanceRequest, _ := http.NewRequest("GET", "/api/v1/wallets/"+wallet.ID+"/balance", nil)
	balanceRecorder := httptest.NewRecorder()
	s.router.ServeHTTP(balanceRecorder, balanceRequest)
	s.Equal(http.StatusOK, balanceRecorder.Code)
	balance := `{"balance":"0.00"}`
	s.Equal(balance, balanceRecorder.Body.String())
}

func (s *ServerTestSuite) TestShoudlReturn200WhenWithdrawingFromWallet() {
	req, _ := http.NewRequest("POST", "/api/v1/wallets", nil)
	w := httptest.NewRecorder()
	s.router.ServeHTTP(w, req)
	s.Equal(http.StatusCreated, w.Code)
	var wallet *domain.Wallet
	json.Unmarshal(w.Body.Bytes(), &wallet)

	depositBody := `{"amount":"100.00"}`
	body := bytes.NewBufferString(depositBody)
	depositRequest, _ := http.NewRequest("PUT", "/api/v1/wallets/"+wallet.ID+"/balance/deposit", body)
	depositRecorder := httptest.NewRecorder()
	s.router.ServeHTTP(depositRecorder, depositRequest)
	s.Equal(http.StatusOK, depositRecorder.Code)

	withdrawBody := `{"amount":"50.00"}`
	withdrawBodyBuffer := bytes.NewBufferString(withdrawBody)
	withdrawRequest, _ := http.NewRequest("PUT", "/api/v1/wallets/"+wallet.ID+"/balance/withdraw", withdrawBodyBuffer)
	withdrawRecorder := httptest.NewRecorder()
	s.router.ServeHTTP(withdrawRecorder, withdrawRequest)
	s.Equal(http.StatusOK, withdrawRecorder.Code)

	balanceRequest, _ := http.NewRequest("GET", "/api/v1/wallets/"+wallet.ID+"/balance", nil)
	balanceRecorder := httptest.NewRecorder()
	s.router.ServeHTTP(balanceRecorder, balanceRequest)
	s.Equal(http.StatusOK, balanceRecorder.Code)
	balance := `{"balance":"50.00"}`
	s.Equal(balance, balanceRecorder.Body.String())
}

func (s *ServerTestSuite) TestShoudlReturn400WhenWithdrawingMoreThanBalance() {
	req, _ := http.NewRequest("POST", "/api/v1/wallets", nil)
	w := httptest.NewRecorder()
	s.router.ServeHTTP(w, req)
	s.Equal(http.StatusCreated, w.Code)
	var wallet *domain.Wallet
	json.Unmarshal(w.Body.Bytes(), &wallet)

	depositBody := `{"amount":"100.00"}`
	body := bytes.NewBufferString(depositBody)
	depositRequest, _ := http.NewRequest("PUT", "/api/v1/wallets/"+wallet.ID+"/balance/deposit", body)
	depositRecorder := httptest.NewRecorder()
	s.router.ServeHTTP(depositRecorder, depositRequest)
	s.Equal(http.StatusOK, depositRecorder.Code)

	withdrawBody := `{"amount":"150.00"}`
	withdrawBodyBuffer := bytes.NewBufferString(withdrawBody)
	withdrawRequest, _ := http.NewRequest("PUT", "/api/v1/wallets/"+wallet.ID+"/balance/withdraw", withdrawBodyBuffer)
	withdrawRecorder := httptest.NewRecorder()
	s.router.ServeHTTP(withdrawRecorder, withdrawRequest)
	s.Equal(http.StatusBadRequest, withdrawRecorder.Code)
}

func (s *ServerTestSuite) TestShoudlReturn400WhenDepositingNegativeAmount() {
	req, _ := http.NewRequest("POST", "/api/v1/wallets", nil)
	w := httptest.NewRecorder()
	s.router.ServeHTTP(w, req)
	s.Equal(http.StatusCreated, w.Code)
	var wallet *domain.Wallet
	json.Unmarshal(w.Body.Bytes(), &wallet)

	depositBody := `{"amount":"-100.00"}`
	body := bytes.NewBufferString(depositBody)
	depositRequest, _ := http.NewRequest("PUT", "/api/v1/wallets/"+wallet.ID+"/balance/deposit", body)
	depositRecorder := httptest.NewRecorder()
	s.router.ServeHTTP(depositRecorder, depositRequest)
	s.Equal(http.StatusBadRequest, depositRecorder.Code)
}

func (s *ServerTestSuite) TestShoudlReturn400WhenDepositingZeroAmount() {
	req, _ := http.NewRequest("POST", "/api/v1/wallets", nil)
	w := httptest.NewRecorder()
	s.router.ServeHTTP(w, req)
	s.Equal(http.StatusCreated, w.Code)
	var wallet *domain.Wallet
	json.Unmarshal(w.Body.Bytes(), &wallet)

	depositBody := `{"amount":"0.00"}`
	body := bytes.NewBufferString(depositBody)
	depositRequest, _ := http.NewRequest("PUT", "/api/v1/wallets/"+wallet.ID+"/balance/deposit", body)
	depositRecorder := httptest.NewRecorder()
	s.router.ServeHTTP(depositRecorder, depositRequest)
	s.Equal(http.StatusBadRequest, depositRecorder.Code)
}

func (s *ServerTestSuite) TestShoudlReturn400WhenWithdrawingNegativeAmount() {
	req, _ := http.NewRequest("POST", "/api/v1/wallets", nil)
	w := httptest.NewRecorder()
	s.router.ServeHTTP(w, req)
	s.Equal(http.StatusCreated, w.Code)
	var wallet *domain.Wallet
	json.Unmarshal(w.Body.Bytes(), &wallet)

	withdrawBody := `{"amount":"-100.00"}`
	withdrawBodyBuffer := bytes.NewBufferString(withdrawBody)
	withdrawRequest, _ := http.NewRequest("PUT", "/api/v1/wallets/"+wallet.ID+"/balance/withdraw", withdrawBodyBuffer)
	withdrawRecorder := httptest.NewRecorder()
	s.router.ServeHTTP(withdrawRecorder, withdrawRequest)
	s.Equal(http.StatusBadRequest, withdrawRecorder.Code)
}

func (s *ServerTestSuite) TestShoudlReturn400WhenWithdrawingZeroAmount() {
	req, _ := http.NewRequest("POST", "/api/v1/wallets", nil)
	w := httptest.NewRecorder()
	s.router.ServeHTTP(w, req)
	s.Equal(http.StatusCreated, w.Code)
	var wallet *domain.Wallet
	json.Unmarshal(w.Body.Bytes(), &wallet)

	withdrawBody := `{"amount":"0.00"}`
	withdrawBodyBuffer := bytes.NewBufferString(withdrawBody)
	withdrawRequest, _ := http.NewRequest("PUT", "/api/v1/wallets/"+wallet.ID+"/balance/withdraw", withdrawBodyBuffer)
	withdrawRecorder := httptest.NewRecorder()
	s.router.ServeHTTP(withdrawRecorder, withdrawRequest)
	s.Equal(http.StatusBadRequest, withdrawRecorder.Code)
}

func (s *ServerTestSuite) TestShoudlReturn400WhenDepositingInvalidAmount() {
	req, _ := http.NewRequest("POST", "/api/v1/wallets", nil)
	w := httptest.NewRecorder()
	s.router.ServeHTTP(w, req)
	s.Equal(http.StatusCreated, w.Code)
	var wallet *domain.Wallet
	json.Unmarshal(w.Body.Bytes(), &wallet)

	depositBody := `{"amount":"invalid"}`
	body := bytes.NewBufferString(depositBody)
	depositRequest, _ := http.NewRequest("PUT", "/api/v1/wallets/"+wallet.ID+"/balance/deposit", body)
	depositRecorder := httptest.NewRecorder()
	s.router.ServeHTTP(depositRecorder, depositRequest)
	s.Equal(http.StatusBadRequest, depositRecorder.Code)
}

func TestRoutesTestSuite(t *testing.T) {
	suite.Run(t, new(ServerTestSuite))
}
