package model

type DecodeReq struct {
	Base64String string `json:inputString`
}

type DecodeResp struct {
	DecodeString string `json:outputString`
}
