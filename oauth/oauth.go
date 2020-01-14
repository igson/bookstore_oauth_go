package oauth

import (
	"net/http"
	"strconv"
	"encoding/json"
	"fmt"
	"time"
	 "github.com/igson/bookstoreApi/utils/erros"
	"strings"
 	"github.com/mercadolibre/golang-restclient/rest"
)

const(
	headerXPublic = "X-Public"
	headerXClientId = "X-Client-Id"
	headerXCallerId = "X-Caller-Id"
	paramAccessToken string = "access_token"
)

var (
	oauthRestClient = rest.RequestBuilder{
		BaseURL: "http://localhost:8080/api",
		Timeout: 200 * time.Millisecond,
	}
)

type tokenAcesso struct {
	Id  string 	`json:"id"`
	UserId       int64  `json:"user_id"`
	ClienteId    int64  `json:"cliente_id"`
}

type oauthInterface interface {

}

func IsPublic(request *http.Request) bool {
	if request == nil {
		return true
	}
	return request.Header.Get(headerXPublic) == "true"
}


func Autenticacao(request *http.Request) *erros.MsgErroApi {
	if request == nil {
		return nil
	}

	cleanRequest(request)

	tokenID := strings.TrimSpace(request.URL.Query().Get(paramAccessToken))

	if tokenID == "" {
		return nil
	}

	token, erro := getTokenAcesso(tokenID)

	if erro != nil {
		if erro.Status == http.StatusNotFound {
			return nil
		}	
		return erro
	}

	request.Header.Add(headerXCallerId, fmt.Sprintf("%v", token.UserId))
	request.Header.Add(headerXClientId, fmt.Sprintf("%v", token.ClienteId))

	return nil
}


func GetCallerId(request *http.Request) int64 {
	if request == nil {
		return 0
	}
	
	callerID, erro := strconv.ParseInt(request.Header.Get(headerXCallerId), 10, 64)
	
	if erro != nil {
		return 0
	}
	return callerID
}


func GetClientId(request *http.Request) int64 {
	if request == nil {
		return 0
	}
	
	clientID, erro := strconv.ParseInt(request.Header.Get(headerXClientId), 10, 64)
	
	if erro != nil {
		return 0
	}
	return clientID
}


func cleanRequest(request *http.Request) {
	if request == nil {
		return 
	}
	request.Header.Del(headerXClientId)
	request.Header.Del(headerXCallerId)
	
}

func getTokenAcesso(tokenAcessoID string) (*tokenAcesso, *erros.MsgErroApi) {
	
	fmt.Println("Token ID: ", tokenAcessoID)

	resposta := oauthRestClient.Get(fmt.Sprintf("/oauth/tokens/%s", tokenAcessoID))
	
	if resposta == nil || resposta.Response == nil {
		return nil, erros.MsgInternalServerError("Resposta inválida ao tentar pegar token de acesso")
	}

	if resposta.StatusCode > 299 {
		var restErro erros.MsgErroApi
		if erro := json.Unmarshal(resposta.Bytes(), &restErro); erro != nil {
			return nil, erros.MsgInternalServerError("Erro na interface ao tentar efetuar token de acesso")
		}
		return nil, &restErro
	}

	var token tokenAcesso

	if erro := json.Unmarshal(resposta.Bytes(), &token); erro != nil {
		return nil, erros.MsgInternalServerError("erro ao tentar desserializar token ")
	}

	return &token, nil

}
