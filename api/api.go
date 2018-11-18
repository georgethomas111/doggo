package api

import (
	"encoding/json"
	"fmt"
	"net/http"
	"strconv"
	"strings"
	"time"
)

type API struct {
	db DB
	s  *http.Server
}

func (a *API) Percent(w http.ResponseWriter, r *http.Request) {
	sStr := strings.Split(r.URL.Path, "/")[2]
	eStr := strings.Split(r.URL.Path, "/")[3]
	fmt.Println("Percent handler hit", time.Now().UTC().Unix(), sStr, eStr)

	sTime, err := strconv.ParseInt(sStr, 10, 64)
	if err != nil {
		// respond with 404.
		fmt.Fprintf(w, err.Error())
	}

	eTime, err := strconv.ParseInt(eStr, 10, 64)
	if err != nil {
		// respond with 404
		fmt.Fprintf(w, err.Error())
	}

	res := a.db.Query(sTime, eTime)

	enc := json.NewEncoder(w)
	err = enc.Encode(res)
	if err != nil {
		// respond with 404
		fmt.Fprintf(w, err.Error())
	}
}

func (a *API) Close() {
	// stop listening to http
	a.s.Close()
}

func New(port string, database DB) *API {
	s := &http.Server{
		Addr:    port,
		Handler: http.DefaultServeMux,
	}

	a := &API{
		db: database,
		s:  s,
	}

	http.HandleFunc("/", a.Percent)
	fmt.Println("Listening for json requests on port", port)

	go func() {
		s.ListenAndServe()
	}()

	return a
}
