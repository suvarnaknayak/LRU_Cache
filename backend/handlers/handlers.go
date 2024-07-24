package handlers

import (
    "encoding/json"
    "net/http"
    "time"
    "backend/cache"
)

var myCache = cache.NewLRUCache(100)

func GetCache(w http.ResponseWriter, r *http.Request) {
    key := r.URL.Query().Get("key")
    value, found := myCache.Get(key)
    if !found {
        http.Error(w, "Key not found", http.StatusNotFound)
        return
    }
    json.NewEncoder(w).Encode(value)
}

func SetCache(w http.ResponseWriter, r *http.Request) {
    var data struct {
        Key        string        `json:"key"`
        Value      string        `json:"value"`
        Expiration time.Duration `json:"expiration"`
    }
    if err := json.NewDecoder(r.Body).Decode(&data); err != nil {
        http.Error(w, err.Error(), http.StatusBadRequest)
        return
    }
    myCache.Set(data.Key, data.Value, data.Expiration)
    w.WriteHeader(http.StatusOK)
}

func DeleteCache(w http.ResponseWriter, r *http.Request) {
    key := r.URL.Query().Get("key")
    myCache.Delete(key)
    w.WriteHeader(http.StatusOK)
}
