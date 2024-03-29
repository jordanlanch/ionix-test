package controller

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
	"net/http"

	"golang.org/x/crypto/bcrypt"

	"github.com/gin-gonic/gin"
	"github.com/invopop/jsonschema"
	"github.com/jordanlanch/ionix-test/bootstrap"
	"github.com/jordanlanch/ionix-test/domain"
	"github.com/stoewer/go-strcase"
	"github.com/xeipuuv/gojsonschema"
)

type LoginController struct {
	LoginUsecase domain.LoginUsecase
	Env          *bootstrap.Env
}

func (lc *LoginController) Login(c *gin.Context) {
	payload := domain.LoginRequest{}

	r := new(jsonschema.Reflector)
	r.KeyNamer = strcase.SnakeCase // from package github.com/stoewer/go-strcase
	// r.RequiredFromJSONSchemaTags = true
	payloadSchemaReflect := r.Reflect(domain.LoginRequest{})
	schemaLoader := gojsonschema.NewGoLoader(&payloadSchemaReflect)

	requestBody, err := ioutil.ReadAll(c.Request.Body)
	if err != nil {
		c.AbortWithStatusJSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	documentLoader := gojsonschema.NewBytesLoader(requestBody)

	result, err := gojsonschema.Validate(schemaLoader, documentLoader)
	if err != nil {
		c.JSON(http.StatusBadRequest, domain.ErrorResponse{Message: err.Error()})
		return
	}

	if !result.Valid() {
		err = fmt.Errorf("[ERROR] invalid payload %+v", result.Errors())
		c.JSON(http.StatusBadRequest, domain.ErrorResponse{Message: err.Error()})
		return
	}

	if err := json.Unmarshal(requestBody, &payload); err != nil {
		err = fmt.Errorf("[ERROR] error unmarshaling JSON %+v", err)
		c.JSON(http.StatusBadRequest, domain.ErrorResponse{Message: err.Error()})

		return
	}

	user, err := lc.LoginUsecase.GetUserByEmail(c, payload.Email)
	if err != nil {
		c.JSON(http.StatusNotFound, domain.ErrorResponse{Message: "User not found with the given email"})
		return
	}

	if bcrypt.CompareHashAndPassword([]byte(user.Password), []byte(payload.Password)) != nil {
		c.JSON(http.StatusUnauthorized, domain.ErrorResponse{Message: "Invalid credentials"})
		return
	}

	accessToken, err := lc.LoginUsecase.CreateAccessToken(user, lc.Env.AccessTokenSecret, lc.Env.AccessTokenExpiryHour)
	if err != nil {
		c.JSON(http.StatusInternalServerError, domain.ErrorResponse{Message: err.Error()})
		return
	}

	refreshToken, err := lc.LoginUsecase.CreateRefreshToken(user, lc.Env.RefreshTokenSecret, lc.Env.RefreshTokenExpiryHour)
	if err != nil {
		c.JSON(http.StatusInternalServerError, domain.ErrorResponse{Message: err.Error()})
		return
	}

	loginResponse := domain.LoginResponse{
		AccessToken:  accessToken,
		RefreshToken: refreshToken,
	}

	c.JSON(http.StatusOK, loginResponse)
}
