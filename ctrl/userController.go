package ctrl

import (
	"encoding/json"
	"errors"
	"io"
	"net/http"
	"rst1/ctrl/util"
	"rst1/models"
	"rst1/svc"
	"strconv"

	"github.com/go-playground/validator"
	"github.com/gorilla/mux"
)

type Response struct {
	Message string `json:"message"`
}

type UserController struct {
	service svc.BasicCrudService[models.User]
}

func NewUserController(service svc.BasicCrudService[models.User]) *UserController {
	return &UserController{service: service}
}

func (c *UserController) HandleGetUsers(w http.ResponseWriter, r *http.Request) {
	usersResponse := c.service.FindAll()
	SendResponse(w, usersResponse.Value, http.StatusOK)
}

func getIntId(id string) (int, error) {
	intId, err := strconv.Atoi(id)
	if err != nil {
		return 0, err
	}
	return intId, nil
}

func getPathId(r *http.Request) (int, error) {
	vars := mux.Vars(r)
	id, ok := vars["id"]
	if !ok {
		println("id is missing in parameters")
		return 0, errors.New("bad request")
	}
	return getIntId(id)
}

func (c *UserController) HandleGetUser(w http.ResponseWriter, r *http.Request) {
	var statusCode int
	var response any
	intId, err := getPathId(r)
	if err != nil {
		statusCode = http.StatusBadRequest
		SendResponse(w, Response{Message: "Bad Request"}, statusCode)
	} else {
		usersResponse := c.service.FindById(intId)
		if usersResponse.Error != nil {
			statusCode = http.StatusNotFound
			response = Response{Message: "user not found"}
		} else {
			statusCode = http.StatusOK
			response = usersResponse.Value
		}
		SendResponse(w, response, statusCode)
	}
}

func parseBody(body io.ReadCloser) (models.User, error) {
	var user models.User
	err := json.NewDecoder(body).Decode(&user)
	if err != nil {
		println("error parsing body")
		return user, err
	}
	return user, nil
}

func validateBody(u *models.User) (int, error) {
	validate := validator.New()
	if err := validate.Struct(u); err != nil {
		println("validation error: ", err.Error())
		return 0, err
	}
	return 1, nil
}

func handleBody(body io.ReadCloser) (*models.User, error) {
	user, err := parseBody(body)
	if err != nil {
		return nil, err
	}
	_, verr := validateBody(&user)
	if verr != nil {
		return nil, verr
	}
	return &user, nil
}

func checkUserExists(user *models.User, s svc.BasicCrudService[models.User]) bool {
	resp := s.FindById(user.Id)
	if !resp.IsOk {
		return false
	}
	return true
}

func hashUserPassword(user *models.User) error {
	hashedPassword, herr := util.HashPassword(user.Password)
	if herr != nil {
		return herr
	}
	user.Password = hashedPassword
	return nil
}

func (c *UserController) HandleUpdateUser(w http.ResponseWriter, r *http.Request) {
	body := r.Body
	defer body.Close()
	user, err := handleBody(body)
	if err != nil {
		SendResponse(w, Response{Message: err.Error()}, http.StatusBadRequest)
		return
	}
	doesUserExist := checkUserExists(user, c.service)
	if !doesUserExist {
		SendResponse(w, Response{Message: "User does not exist"}, http.StatusBadRequest)
		return
	}
	herr := hashUserPassword(user)
	if herr != nil {
		SendResponse(w, Response{Message: err.Error()}, http.StatusInternalServerError)
		return
	}
	userResp := c.service.Update(user)
	if !userResp.IsOk {
		SendResponse(w, userResp.Error, http.StatusInternalServerError)
		return
	}
	SendResponse(w, userResp.Value, http.StatusAccepted)
}
