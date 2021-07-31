package controller

import (
	"bytes"
	"encoding/json"
	"github.com/SemmiDev/fiber-go-clean-arch/model"
	"github.com/stretchr/testify/assert"
	"io/ioutil"
	"net/http/httptest"
	"testing"
)

func TestRegistrationController_Create(t *testing.T) {
	registrationRepository.DeleteAll()
	createRegistrationRequest := model.RegistrationRequest{
		Name:    "Sammi Aldhi Yanto",
		Email:   "sammidev@gmail.com",
		Phone:   "082387325971",
		Program: "S2",
	}

	requestBody, _ := json.Marshal(createRegistrationRequest)
	request := httptest.NewRequest("POST", "/api/v1/registration", bytes.NewBuffer(requestBody))
	request.Header.Set("Content-Type", "application/json")
	request.Header.Set("Accept", "application/json")

	response, _ := app.Test(request)
	assert.Equal(t, 201, response.StatusCode)
	responseBody, _ := ioutil.ReadAll(response.Body)

	webResponse := model.WebResponse{}
	json.Unmarshal(responseBody, &webResponse)
	assert.Equal(t, 201, webResponse.Code)
	assert.Equal(t, "Created", webResponse.Status)
	assert.Equal(t, false, webResponse.Error)
	assert.Equal(t, nil, webResponse.ErrorMessage)

	jsonData, _ := json.Marshal(webResponse.Data)
	createRegistrationResponse := model.RegistrationResponse{}
	json.Unmarshal(jsonData, &createRegistrationResponse)
	assert.NotNil(t, createRegistrationResponse.Username)
	assert.NotNil(t, createRegistrationResponse.Password)
	assert.Equal(t, model.S2Bill, createRegistrationResponse.Bill)
	assert.NotEqual(t, model.S1D3D4Bill, createRegistrationResponse.Bill)
}

func TestRegistrationController_CreateFailedEmailIsExists(t *testing.T) {
	registrationRepository.DeleteAll()
	createRegistrationRequest := model.RegistrationRequest{
		Name:  "Sammi Aldhi Yanto",
		Email: "sammidev@gmail.com",
		Phone: "082387325971",
	}
	registrationService.Create(&createRegistrationRequest, model.S2)

	createRegistrationRequest2 := model.RegistrationRequest{
		Name:    "Sammi Aldhi Yanto",
		Email:   "sammidev@gmail.com",
		Phone:   "082387325971",
		Program: "S2",
	}

	requestBody, _ := json.Marshal(createRegistrationRequest2)
	request := httptest.NewRequest("POST", "/api/v1/registration", bytes.NewBuffer(requestBody))
	request.Header.Set("Content-Type", "application/json")
	request.Header.Set("Accept", "application/json")

	response, _ := app.Test(request)
	assert.Equal(t, 400, response.StatusCode)
	responseBody, _ := ioutil.ReadAll(response.Body)

	webResponse := model.WebResponse{}
	json.Unmarshal(responseBody, &webResponse)
	assert.Equal(t, 400, webResponse.Code)
	assert.Equal(t, "Bad Request", webResponse.Status)
	assert.Equal(t, true, webResponse.Error)
	assert.Equal(t, "email has been recorded", webResponse.ErrorMessage)

	jsonData, _ := json.Marshal(webResponse.Data)
	createRegistrationResponse := model.RegistrationResponse{}
	json.Unmarshal(jsonData, &createRegistrationRequest2)
	assert.Empty(t, createRegistrationResponse.Username)
	assert.Empty(t, createRegistrationResponse.Password)
	assert.Empty(t, createRegistrationResponse.Bill)
	assert.Empty(t, createRegistrationResponse.VirtualAccount)
}

func TestRegistrationController_CreateFailedPhoneIsExists(t *testing.T) {
	registrationRepository.DeleteAll()
	createRegistrationRequest := model.RegistrationRequest{
		Name:  "Sammi Aldhi Yanto",
		Email: "sammidev@gmail.com",
		Phone: "082387325971",
	}
	registrationService.Create(&createRegistrationRequest, model.S2)

	createRegistrationRequest2 := model.RegistrationRequest{
		Name:    "Sammi Aldhi Yanto",
		Email:   "sammidev2@gmail.com",
		Phone:   "082387325971",
		Program: "S2",
	}

	requestBody, _ := json.Marshal(createRegistrationRequest2)
	request := httptest.NewRequest("POST", "/api/v1/registration", bytes.NewBuffer(requestBody))
	request.Header.Set("Content-Type", "application/json")
	request.Header.Set("Accept", "application/json")

	response, _ := app.Test(request)
	assert.Equal(t, 400, response.StatusCode)
	responseBody, _ := ioutil.ReadAll(response.Body)

	webResponse := model.WebResponse{}
	json.Unmarshal(responseBody, &webResponse)
	assert.Equal(t, 400, webResponse.Code)
	assert.Equal(t, "Bad Request", webResponse.Status)
	assert.Equal(t, true, webResponse.Error)
	assert.Equal(t, "phone has been recorded", webResponse.ErrorMessage)

	jsonData, _ := json.Marshal(webResponse.Data)
	createRegistrationResponse := model.RegistrationResponse{}
	json.Unmarshal(jsonData, &createRegistrationRequest2)
	assert.Empty(t, createRegistrationResponse.Username)
	assert.Empty(t, createRegistrationResponse.Password)
	assert.Empty(t, createRegistrationResponse.Bill)
	assert.Empty(t, createRegistrationResponse.VirtualAccount)
}

func TestRegistrationController_CreateFailedNameIsEmpty(t *testing.T) {
	registrationRepository.DeleteAll()
	createRegistrationRequest := model.RegistrationRequest{
		Name:    "",
		Email:   "sammidev@gmail.com",
		Phone:   "082387325971",
		Program: "S2",
	}

	requestBody, _ := json.Marshal(createRegistrationRequest)
	request := httptest.NewRequest("POST", "/api/v1/registration", bytes.NewBuffer(requestBody))
	request.Header.Set("Content-Type", "application/json")
	request.Header.Set("Accept", "application/json")

	response, _ := app.Test(request)
	assert.Equal(t, 400, response.StatusCode)
	responseBody, _ := ioutil.ReadAll(response.Body)

	webResponse := model.WebResponse{}
	json.Unmarshal(responseBody, &webResponse)
	assert.Equal(t, 400, webResponse.Code)
	assert.Equal(t, "Bad Request", webResponse.Status)
	assert.Equal(t, true, webResponse.Error)
	assert.Equal(t,
		map[string]interface{}(map[string]interface{}{"Required_Name": "Name Is Empty"}),
		webResponse.ErrorMessage)
	assert.Nil(t, webResponse.Data)

	jsonData, _ := json.Marshal(webResponse.Data)
	createRegistrationResponse := model.RegistrationResponse{}
	json.Unmarshal(jsonData, &createRegistrationResponse)
	assert.Empty(t, createRegistrationResponse.Username)
	assert.Empty(t, createRegistrationResponse.Password)
	assert.Empty(t, createRegistrationResponse.Bill)
	assert.Empty(t, createRegistrationResponse.VirtualAccount)
}

func TestRegistrationController_CreateFailedRequestsIsEmpty(t *testing.T) {
	registrationRepository.DeleteAll()
	createRegistrationRequest := model.RegistrationRequest{
		Name:    "",
		Email:   "",
		Phone:   "",
		Program: "S2",
	}

	requestBody, _ := json.Marshal(createRegistrationRequest)
	request := httptest.NewRequest("POST", "/api/v1/registration", bytes.NewBuffer(requestBody))
	request.Header.Set("Content-Type", "application/json")
	request.Header.Set("Accept", "application/json")

	response, _ := app.Test(request)
	assert.Equal(t, 400, response.StatusCode)
	responseBody, _ := ioutil.ReadAll(response.Body)

	webResponse := model.WebResponse{}
	json.Unmarshal(responseBody, &webResponse)
	assert.Equal(t, 400, webResponse.Code)
	assert.Equal(t, "Bad Request", webResponse.Status)
	assert.Equal(t, true, webResponse.Error)
	assert.Equal(t,
		map[string]interface{}{
			"Required_Email": "Email Is Empty",
			"Required_Name":  "Name Is Empty",
			"Required_Phone": "Phone Is Empty",
			"invalid_Phone":  "Phone Number Is Not Valid",
			"invalid_Email":  "Email Is Not Valid"},
		webResponse.ErrorMessage)
	assert.Nil(t, webResponse.Data)

	jsonData, _ := json.Marshal(webResponse.Data)
	createRegistrationResponse := model.RegistrationResponse{}
	json.Unmarshal(jsonData, &createRegistrationResponse)
	assert.Empty(t, createRegistrationResponse.Username)
	assert.Empty(t, createRegistrationResponse.Password)
	assert.Empty(t, createRegistrationResponse.Bill)
	assert.Empty(t, createRegistrationResponse.VirtualAccount)
}

func TestRegistrationController_CreateFailedInvalidPhone(t *testing.T) {
	registrationRepository.DeleteAll()
	createRegistrationRequest := model.RegistrationRequest{
		Name:    "sammi",
		Email:   "sammi@gmail.com",
		Phone:   "aoksoadal",
		Program: "S2",
	}

	requestBody, _ := json.Marshal(createRegistrationRequest)
	request := httptest.NewRequest("POST", "/api/v1/registration", bytes.NewBuffer(requestBody))
	request.Header.Set("Content-Type", "application/json")
	request.Header.Set("Accept", "application/json")

	response, _ := app.Test(request)
	assert.Equal(t, 400, response.StatusCode)
	responseBody, _ := ioutil.ReadAll(response.Body)

	webResponse := model.WebResponse{}
	json.Unmarshal(responseBody, &webResponse)
	assert.Equal(t, 400, webResponse.Code)
	assert.Equal(t, "Bad Request", webResponse.Status)
	assert.Equal(t, true, webResponse.Error)
	assert.Equal(t,
		map[string]interface{}{
			"invalid_Phone": "Phone Number Is Not Valid"},
		webResponse.ErrorMessage)
	assert.Nil(t, webResponse.Data)

	jsonData, _ := json.Marshal(webResponse.Data)
	createRegistrationResponse := model.RegistrationResponse{}
	json.Unmarshal(jsonData, &createRegistrationResponse)
	assert.Empty(t, createRegistrationResponse.Username)
	assert.Empty(t, createRegistrationResponse.Password)
	assert.Empty(t, createRegistrationResponse.Bill)
	assert.Empty(t, createRegistrationResponse.VirtualAccount)
}

func TestRegistrationController_CreateFailedInvalidPhoneAndEmail(t *testing.T) {
	registrationRepository.DeleteAll()
	createRegistrationRequest := model.RegistrationRequest{
		Name:    "sammi",
		Email:   "sammiasam",
		Phone:   "aoksoadal",
		Program: "S2",
	}

	requestBody, _ := json.Marshal(createRegistrationRequest)
	request := httptest.NewRequest("POST", "/api/v1/registration", bytes.NewBuffer(requestBody))
	request.Header.Set("Content-Type", "application/json")
	request.Header.Set("Accept", "application/json")

	response, _ := app.Test(request)
	assert.Equal(t, 400, response.StatusCode)
	responseBody, _ := ioutil.ReadAll(response.Body)

	webResponse := model.WebResponse{}
	json.Unmarshal(responseBody, &webResponse)
	assert.Equal(t, 400, webResponse.Code)
	assert.Equal(t, "Bad Request", webResponse.Status)
	assert.Equal(t, true, webResponse.Error)
	assert.Equal(t,
		map[string]interface{}{
			"invalid_Phone": "Phone Number Is Not Valid",
			"invalid_Email": "Email Is Not Valid",
		},
		webResponse.ErrorMessage)
	assert.Nil(t, webResponse.Data)

	jsonData, _ := json.Marshal(webResponse.Data)
	createRegistrationResponse := model.RegistrationResponse{}
	json.Unmarshal(jsonData, &createRegistrationResponse)
	assert.Empty(t, createRegistrationResponse.Username)
	assert.Empty(t, createRegistrationResponse.Password)
	assert.Empty(t, createRegistrationResponse.Bill)
	assert.Empty(t, createRegistrationResponse.VirtualAccount)
}

func TestRegistrationController_CreateFailedProgramNotRecognize(t *testing.T) {
	registrationRepository.DeleteAll()
	createRegistrationRequest := model.RegistrationRequest{
		Name:    "izzah",
		Email:   "izzah@gmail.com",
		Phone:   "08123912389123",
		Program: "xxxx",
	}

	requestBody, _ := json.Marshal(createRegistrationRequest)
	request := httptest.NewRequest("POST", "/api/v1/registration", bytes.NewBuffer(requestBody))
	request.Header.Set("Content-Type", "application/json")
	request.Header.Set("Accept", "application/json")

	response, _ := app.Test(request)
	assert.Equal(t, 400, response.StatusCode)
	responseBody, _ := ioutil.ReadAll(response.Body)

	webResponse := model.WebResponse{}
	json.Unmarshal(responseBody, &webResponse)
	assert.Equal(t, 400, webResponse.Code)
	assert.Equal(t, "Bad Request", webResponse.Status)
	assert.Equal(t, true, webResponse.Error)
	assert.Equal(t,
		map[string]interface{}{
			"Program_Not_Available": "Please Chose Between S1D3D4 or S2",
		},
		webResponse.ErrorMessage)
	assert.Nil(t, webResponse.Data)

	jsonData, _ := json.Marshal(webResponse.Data)
	createRegistrationResponse := model.RegistrationResponse{}
	json.Unmarshal(jsonData, &createRegistrationResponse)
	assert.Empty(t, createRegistrationResponse.Username)
	assert.Empty(t, createRegistrationResponse.Password)
	assert.Empty(t, createRegistrationResponse.Bill)
	assert.Empty(t, createRegistrationResponse.VirtualAccount)
}

func TestRegistrationController_UpdateSuccess(t *testing.T) {
	registrationRepository.DeleteAll()
	createRegistrationRequest := model.RegistrationRequest{
		Name:    "Sammi Aldhi Yanto",
		Email:   "sammidev@gmail.com",
		Phone:   "082387325971",
		Program: "S2",
	}
	temp, _ := registrationService.Create(&createRegistrationRequest, model.S2)

	updateStatusRequest := model.UpdateStatus{
		VirtualAccount: temp.VirtualAccount,
	}

	requestBody, _ := json.Marshal(updateStatusRequest)
	request := httptest.NewRequest("PUT", "/api/v1/registration/status", bytes.NewBuffer(requestBody))
	request.Header.Set("Content-Type", "application/json")
	request.Header.Set("Accept", "application/json")

	response, _ := app.Test(request)
	assert.Equal(t, 200, response.StatusCode)
	responseBody, _ := ioutil.ReadAll(response.Body)

	webResponse := model.WebResponse{}
	json.Unmarshal(responseBody, &webResponse)
	assert.Equal(t, 200, webResponse.Code)
	assert.Equal(t, "Ok", webResponse.Status)
	assert.Equal(t, false, webResponse.Error)
	assert.Equal(t, nil, webResponse.ErrorMessage)
	assert.Equal(t, map[string]interface{}(map[string]interface{}{"status": "updated"}), webResponse.Data)
}

func TestRegistrationController_UpdateFailedEmptyVA(t *testing.T) {
	registrationRepository.DeleteAll()
	createRegistrationRequest := model.RegistrationRequest{
		Name:    "Sammi Aldhi Yanto",
		Email:   "sammidev@gmail.com",
		Phone:   "082387325971",
		Program: "S2",
	}
	registrationService.Create(&createRegistrationRequest, model.S2)

	updateStatusRequest := model.UpdateStatus{
		VirtualAccount: "",
	}

	requestBody, _ := json.Marshal(updateStatusRequest)
	request := httptest.NewRequest("PUT", "/api/v1/registration/status", bytes.NewBuffer(requestBody))
	request.Header.Set("Content-Type", "application/json")
	request.Header.Set("Accept", "application/json")

	response, _ := app.Test(request)
	assert.Equal(t, 400, response.StatusCode)
	responseBody, _ := ioutil.ReadAll(response.Body)

	webResponse := model.WebResponse{}
	json.Unmarshal(responseBody, &webResponse)
	assert.Equal(t, 400, webResponse.Code)
	assert.Equal(t, "Bad Request", webResponse.Status)
	assert.Equal(t, true, webResponse.Error)
	assert.Equal(t, map[string]interface{}(map[string]interface{}{"Required_VA": "Virtual Account Is Empty"}), webResponse.ErrorMessage)
	assert.Nil(t, webResponse.Data)
}

func TestRegistrationController_UpdateFailedVaNotFound(t *testing.T) {
	registrationRepository.DeleteAll()
	updateStatusRequest := model.UpdateStatus{
		VirtualAccount: "1241231321231",
	}

	requestBody, _ := json.Marshal(updateStatusRequest)
	request := httptest.NewRequest("PUT", "/api/v1/registration/status", bytes.NewBuffer(requestBody))
	request.Header.Set("Content-Type", "application/json")
	request.Header.Set("Accept", "application/json")

	response, _ := app.Test(request)
	assert.Equal(t, 500, response.StatusCode)
	responseBody, _ := ioutil.ReadAll(response.Body)

	webResponse := model.WebResponse{}
	json.Unmarshal(responseBody, &webResponse)
	assert.Equal(t, 500, webResponse.Code)
	assert.Equal(t, "Internal Server Error", webResponse.Status)
	assert.Equal(t, true, webResponse.Error)
	assert.Equal(t, "va not found", webResponse.ErrorMessage)
	assert.Nil(t, webResponse.Data)
}
