package erros

import "net/http"

type MsgErroApi struct {
	Mensagem string `json:"mensagem,omitempty"`
	Status   int    `json:"status,omitempty"`
	Erro     string `json:"erro,omitempty"`
}

func MsgBadRequestErro(mensagem string) *MsgErroApi {
	return &MsgErroApi{
		Mensagem: mensagem, Status: http.StatusBadRequest, Erro: "BAD_REQUEST",
	}
}

func MsgNotFoundErro(mensagem string) *MsgErroApi {
	return &MsgErroApi{
		Mensagem: mensagem, Status: http.StatusNotFound, Erro: "NOT_FOUND",
	}
}

func MsgInternalServerError(mensagem string) *MsgErroApi {
	return &MsgErroApi{
		Mensagem: mensagem, Status: http.StatusInternalServerError, Erro: "INTERNAL_SERVER_ERROR",
	}
}
