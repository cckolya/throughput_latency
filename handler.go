package main

import (
	"log"
	"net/http"

	"github.com/prometheus/client_golang/prometheus/promhttp"
)

//var (
//	hello = []byte("Hello world")
//)

type Handler struct {
	db    *Postgres
	cache *Redis
}

func NewHandler(db *Postgres, cache *Redis) *Handler {
	return &Handler{
		db:    db,
		cache: cache,
	}
}

func (h *Handler) Handle() error {
	http.HandleFunc("/", func(w http.ResponseWriter, r *http.Request) {
		w.WriteHeader(http.StatusOK)
		w.Write([]byte("Hello world"))
	})

	http.HandleFunc("/cpu", func(w http.ResponseWriter, r *http.Request) {
		time := r.URL.Query().Get("time")
		if time == "" {
			w.WriteHeader(http.StatusBadRequest)
			return
		}
		err := cpu(time)
		if err != nil {
			w.WriteHeader(http.StatusBadRequest)
			return
		}
		w.WriteHeader(http.StatusOK)
	})

	http.HandleFunc("/postgres", func(w http.ResponseWriter, r *http.Request) {
		err := h.db.SelectDual()
		if err != nil {
			log.Println(err)
			w.WriteHeader(http.StatusInternalServerError)
			return
		}
		w.WriteHeader(http.StatusOK)
	})

	http.HandleFunc("/redis", func(w http.ResponseWriter, r *http.Request) {
		//err := h.cache.Get()
		//if err != nil {
		//	log.Println(err)
		//	w.WriteHeader(http.StatusInternalServerError)
		//	return
		//}
		//w.WriteHeader(http.StatusOK)
	})

	http.Handle("/metrics", promhttp.Handler())

	return http.ListenAndServe(":8080", nil)
}
