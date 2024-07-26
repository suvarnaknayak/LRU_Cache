package handlers

import (
	"backend/cache"
	"encoding/json"
	"log"
	"net/http"
	"time"

	"github.com/gorilla/mux"
)

var myCache = cache.NewLRUCache(100)

type MessageResponse struct {
	Message interface{} `json:"message"`
}

func SetCache(w http.ResponseWriter, r *http.Request) {
	var data struct {
		Key        string `json:"key"`
		Value      string `json:"value"`
		Expiration string `json:"expiration"`
	}
	if err := json.NewDecoder(r.Body).Decode(&data); err != nil {
		http.Error(w, err.Error(), http.StatusBadRequest)
		return
	}

	expirationDuration, err := time.ParseDuration(data.Expiration)
	if err != nil {
		http.Error(w, "Invalid expiration duration format", http.StatusBadRequest)
		return
	}

	myCache.Set(data.Key, data.Value, expirationDuration)
	resp := MessageResponse{
		Message: "Successfully created",
	}

	jsonResp, err := json.Marshal(resp)
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusOK)
	w.Write(jsonResp)
}

func DeleteCache(w http.ResponseWriter, r *http.Request) {
	vars := mux.Vars(r)
	log.Println(vars)
	key := vars["key"]

	isDeleted := myCache.Delete(key)

	if !isDeleted {
		http.Error(w, "Key not found", http.StatusNotFound)
		return
	}

	resp := MessageResponse{
		Message: "Successfully deleted",
	}

	jsonResp, err := json.Marshal(resp)
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusOK)
	w.Write(jsonResp)
}

func GetCacheList(w http.ResponseWriter, r *http.Request) {

	value := myCache.GetList()
	log.Println((value))

	jsonResp, err := json.Marshal(value)
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusOK)
	w.Write(jsonResp)
}
