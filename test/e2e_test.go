package test

import (
	"fmt"
	"net/http"
	"testing"
	"time"

	"github.com/jordanlanch/ionix-test/domain"
	"github.com/sirupsen/logrus"

	"github.com/jordanlanch/ionix-test/test/e2e"
)

const (
	statusOK                  = http.StatusOK
	statusUnprocessableEntity = http.StatusUnprocessableEntity
	errorInvalidPayload       = "error.invalid_payload"
	errorLoginEventBaseUserID = "login.event_base.user_id"
	errorDescription          = "[ERROR] invalid payload [" + errorLoginEventBaseUserID + ": String length must be greater than or equal to 1]"
)

func TestE2E(t *testing.T) {
	expect, teardown := e2e.Setup(t, "fixtures/login_test")
	defer teardown()

	log := logrus.New()
	log.SetFormatter(&logrus.TextFormatter{
		FullTimestamp: true,
		ForceColors:   true,
	})

	var (
		accessToken   string
		drugID        int
		vaccinationID int
	)

	t.Run("Signup error ", func(t *testing.T) {
		log.Info("Iniciando prueba:", t.Name())
		in := domain.SignupRequest{
			Email: "jordan.dev93@gmail.com",
		}
		obj := expect.POST("/signup").
			WithJSON(in).
			Expect().
			Status(http.StatusBadRequest).
			JSON().Object()
		obj.Value("message").String().IsEqual("[ERROR] invalid payload [(root): password is required]")
		log.Info("Finalizando prueba:", t.Name())
	})

	t.Run("Login user not found ", func(t *testing.T) {
		log.Info("Iniciando prueba:", t.Name())
		in := domain.LoginRequest{
			Email: "jordan.dev93@gmail.com",
		}
		obj := expect.POST("/login").
			WithJSON(in).
			Expect().
			Status(http.StatusNotFound).
			JSON().Object()
		obj.Value("message").String().IsEqual("User not found with the given email")
		log.Info("Finalizando prueba:", t.Name())
	})

	t.Run("Signup success", func(t *testing.T) {
		log.Info("Iniciando prueba:", t.Name())

		in := domain.SignupRequest{
			Email:    "jordan.dev93@gmail.com",
			Password: "pass1234*",
		}

		obj := expect.POST("/signup").
			WithJSON(in).
			Expect().
			Status(http.StatusOK).
			JSON().Object()
		obj.Value("accessToken").String().NotEqual("")
		log.Info("Finalizando prueba:", t.Name())
	})

	t.Run("Login success", func(t *testing.T) {
		log.Info("Iniciando prueba:", t.Name())

		in := domain.LoginRequest{
			Email:    "jordan.dev93@gmail.com",
			Password: "pass1234*",
		}

		obj := expect.POST("/login").
			WithJSON(in).
			Expect().
			Status(http.StatusOK).
			JSON().Object()
		accessToken = obj.Value("accessToken").String().Raw()
		obj.Value("accessToken").String().NotEqual("")
		log.Info("Finalizando prueba:", t.Name())
	})

	t.Run("Create and Use Drug", func(t *testing.T) {
		drugIn := domain.Drug{
			Name:        "Vacuna Prueba",
			Approved:    true,
			MinDose:     1,
			MaxDose:     2,
			AvailableAt: time.Now(),
		}

		// Creación de un nuevo medicamento
		drugObj := expect.POST("/drugs").
			WithHeader("Authorization", fmt.Sprintf("Bearer %s", accessToken)).
			WithJSON(drugIn).
			Expect().
			Status(http.StatusCreated).
			JSON().Object()

		// Almacenamiento del ID del medicamento para su uso en la creación de vacunaciones
		drugID = int(drugObj.Value("id").Number().Raw())
	})

	t.Run("Update Drug", func(t *testing.T) {
		updateDrugIn := domain.Drug{
			Name:        "Vacuna Prueba Actualizada",
			Approved:    true,
			MinDose:     1,
			MaxDose:     2,
			AvailableAt: time.Now(),
		}

		expect.PUT(fmt.Sprintf("/drugs/%d", drugID)).
			WithHeader("Authorization", fmt.Sprintf("Bearer %s", accessToken)).
			WithJSON(updateDrugIn).
			Expect().
			Status(http.StatusOK)
	})

	t.Run("Get Drug Details", func(t *testing.T) {
		expect.GET(fmt.Sprintf("/drugs/%d", drugID)).
			WithHeader("Authorization", fmt.Sprintf("Bearer %s", accessToken)).
			Expect().
			Status(http.StatusOK).
			JSON().Object().Value("id").IsEqual(drugID)
	})

	t.Run("Fetch All Drugs", func(t *testing.T) {
		expect.GET("/drugs").
			WithHeader("Authorization", fmt.Sprintf("Bearer %s", accessToken)).
			Expect().
			Status(http.StatusOK).
			JSON().Array().NotEmpty()
	})

	t.Run("Delete Drug", func(t *testing.T) {
		expect.DELETE(fmt.Sprintf("/drugs/%d", drugID)).
			WithHeader("Authorization", fmt.Sprintf("Bearer %s", accessToken)).
			Expect().
			Status(http.StatusOK)
	})

	t.Run("Create and Use Drug", func(t *testing.T) {
		drugIn := domain.Drug{
			Name:        "Vacuna Prueba",
			Approved:    true,
			MinDose:     1,
			MaxDose:     2,
			AvailableAt: time.Now(),
		}

		// Creación de un nuevo medicamento
		drugObj := expect.POST("/drugs").
			WithHeader("Authorization", fmt.Sprintf("Bearer %s", accessToken)).
			WithJSON(drugIn).
			Expect().
			Status(http.StatusCreated).
			JSON().Object()

		// Almacenamiento del ID del medicamento para su uso en la creación de vacunaciones
		drugID = int(drugObj.Value("id").Number().Raw())
	})

	t.Run("Create Vaccination with Drug", func(t *testing.T) {
		vaccinationIn := domain.Vaccination{
			Name:   "John Doe",
			DrugID: drugID,
			Dose:   1,
			Date:   time.Now(),
		}

		vaccinationObj := expect.POST("/vaccination").
			WithHeader("Authorization", fmt.Sprintf("Bearer %s", accessToken)).
			WithJSON(vaccinationIn).
			Expect().
			Status(http.StatusCreated).
			JSON().Object()

		// Almacenamiento del ID del medicamento para su uso en la creación de vacunaciones
		vaccinationID = int(vaccinationObj.Value("id").Number().Raw())
	})

	t.Run("Update Vaccination", func(t *testing.T) {
		updateVaccinationIn := domain.Vaccination{
			Name:   "Jane Doe Actualizada",
			DrugID: drugID,
			Dose:   2,
			Date:   time.Now(),
		}

		expect.PUT(fmt.Sprintf("/vaccination/%d", vaccinationID)).
			WithHeader("Authorization", fmt.Sprintf("Bearer %s", accessToken)).
			WithJSON(updateVaccinationIn).
			Expect().
			Status(http.StatusOK)
	})

	t.Run("Get Vaccination Details", func(t *testing.T) {
		expect.GET(fmt.Sprintf("/vaccination/%d", vaccinationID)). // Asumiendo que el ID de vacunación es 1.
										WithHeader("Authorization", fmt.Sprintf("Bearer %s", accessToken)).
										Expect().
										Status(http.StatusOK).
										JSON().Object().Value("id").IsEqual(1)
	})

	t.Run("Fetch All Vaccinations", func(t *testing.T) {
		expect.GET("/vaccination").
			WithHeader("Authorization", fmt.Sprintf("Bearer %s", accessToken)).
			Expect().
			Status(http.StatusOK).
			JSON().Array().NotEmpty()
	})

	t.Run("Delete Vaccination", func(t *testing.T) {
		expect.DELETE(fmt.Sprintf("/vaccination/%d", vaccinationID)). // Asumiendo que el ID de vacunación es 1.
										WithHeader("Authorization", fmt.Sprintf("Bearer %s", accessToken)).
										Expect().
										Status(http.StatusOK)
	})

}
