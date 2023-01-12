package ctrl

import (
	"net/http"
	"rst1/svc"
)

type RemoteUserController struct {
	service *svc.RemoteUserService
}

func NewRemoteUserController(service *svc.RemoteUserService) *RemoteUserController {
	return &RemoteUserController{
		service: service,
	}
}

func (c *RemoteUserController) HandleGetAddress(w http.ResponseWriter, r *http.Request) {
	addrResp := c.service.GetExpandedUser(1)
	if !addrResp.IsOk {
		SendResponse(w, map[string]string{"Error": addrResp.Error.Error()}, http.StatusInternalServerError)
		return
	}
	SendResponse(w, addrResp.Value, http.StatusOK)
}
