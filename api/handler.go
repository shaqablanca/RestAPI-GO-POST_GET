package api

import (
	"encoding/json"
	"fmt"
	"net/http"
	"orness/database"
	"orness/models"
	"sync/atomic"
)

var autoIncrement uint32

// GetNotes method is for retreiving data from the memory by using http request
func GetNotes(w http.ResponseWriter, r *http.Request) {

	w.Header().Add("Content-Type", "application/json")

	var notes []models.Note

	tag := r.URL.Query().Get("tag")

	if tag != "" {
		models := database.RDb[tag]

		w.Write(generateSuccessResponse(models))
		return
	}

	for _, v := range database.Db {
		notes = append(notes, v)
	}

	w.Write(generateSuccessResponse(notes))

}

// AddNotes method is for retreiving data from the memory by using http request
func AddNote(w http.ResponseWriter, r *http.Request) {
	w.Header().Add("Content-Type", "application/json")

	var req models.Note

	decoder := json.NewDecoder(r.Body)
	err := decoder.Decode(&req)

	if err != nil {
		fmt.Println(err)
		w.Write(generateErrorResponse("Bad Request:"+err.Error(), 400))
		return
	}

	if req.Message == "" {
		w.Write(generateErrorResponse("message cannot be empty", 500))
		return
	}
	atomic.AddUint32(&autoIncrement, 1)
	req.Id = autoIncrement
	tag := req.Tag
	database.Db[req.Id] = req
	if tag != "" {
		database.RDb[req.Tag] = append(database.RDb[req.Tag], req)
	}

	w.Write(generateSuccessResponse(req))

}

// Generating Error for corresponding statements in the method.
func generateErrorResponse(message string, code int) []byte {

	resp := struct {
		Success bool        `json:"success"`
		Code    int         `json:"code"`
		Message string      `json:"message"`
		Data    interface{} `json:"data"`
	}{
		Success: false,
		Code:    code,
		Message: message,
	}

	b, err := json.Marshal(resp)
	if err != nil {
		fmt.Println(err)
		return []byte("Internal server error")
	}

	return b
}

// Generating Success response for corresponding statements in the method
func generateSuccessResponse(data interface{}) []byte {

	resp := struct {
		Success bool        `json:"success"`
		Code    int         `json:"code"`
		Message string      `json:"message"`
		Data    interface{} `json:"data"`
	}{
		Success: true,
		Code:    200,
		Message: "",
		Data:    data,
	}

	b, err := json.Marshal(resp)
	if err != nil {
		fmt.Println(err)
	}

	return b
}
