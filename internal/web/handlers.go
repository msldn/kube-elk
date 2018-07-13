package web

import (
	"encoding/json"
	"fmt"
	"github.com/gorilla/mux"
	"github.com/marek5050/kube-elk/internal/elk"
	"github.com/sirupsen/logrus"
	"io/ioutil"
	"net/http"
)

func Index(w http.ResponseWriter, r *http.Request) {
	fmt.Fprint(w, "Welcome!\n")
}

func ElkDelete(w http.ResponseWriter, r *http.Request) {
	vars := mux.Vars(r)
	var OrgId = vars["OrgId"]
	var err error

	var elkconfig = &elk.ElkConfig{
		Org: OrgId,
	}

	err = elk.ElkDelete(elkconfig)

	if err != nil {
		w.Header().Set("Content-Type", "application/json; charset=UTF-8")
		w.WriteHeader(http.StatusNotFound)
		if err := json.NewEncoder(w).Encode(jsonErr{Code: http.StatusNotFound, Text: "Not Found"}); err != nil {
			panic(err)
		}
	}

	w.Header().Set("Content-Type", "application/json; charset=UTF-8")
	w.WriteHeader(http.StatusOK)
	return
}

func ElkCreate(w http.ResponseWriter, r *http.Request) {
	var err error
	var elkconfig = &elk.ElkConfig{
		Ls_p:  31533,
		Kib_p: 31535,
		Org:   "org3",
	}
	body, err := ioutil.ReadAll(r.Body)
	if err != nil {
		w.Header().Set("Content-Type", "application/json; charset=UTF-8")
		w.WriteHeader(http.StatusBadRequest)
		if err := json.NewEncoder(w).Encode(jsonErr{Code: http.StatusNotFound, Text: "Wrong Format"}); err != nil {
			panic(err)
		}
	}

	json.Unmarshal(body, elkconfig)
	err = elk.ElkCreate(elkconfig)

	if err != nil {
		w.Header().Set("Content-Type", "application/json; charset=UTF-8")
		w.WriteHeader(http.StatusNotFound)
		if err := json.NewEncoder(w).Encode(jsonErr{Code: http.StatusNotFound, Text: "Not Found"}); err != nil {
			panic(err)
		}
	}

	w.Header().Set("Content-Type", "application/json; charset=UTF-8")
	w.WriteHeader(http.StatusOK)
	return
}

func ElkGet(w http.ResponseWriter, r *http.Request) {
	var err error
	vars := mux.Vars(r)

	var elkconfig = &elk.ElkConfig{
		Org: vars["OrgId"],
	}

	elk, err := elk.ElkGet( elkconfig)

	if err != nil {
		w.Header().Set("Content-Type", "application/json; charset=UTF-8")
		w.WriteHeader(http.StatusNotFound)
		if err := json.NewEncoder(w).Encode(jsonErr{Code: http.StatusNotFound, Text: "Not Found"}); err != nil {
			panic(err)
		}
	}

	w.Header().Set("Content-Type", "application/json; charset=UTF-8")
	w.WriteHeader(http.StatusOK)

	_body, err := json.Marshal(elk)

	if err != nil {
		logrus.Errorf("Failed to ElkGet ", err)
	}

	w.Write(_body)
	return
}

func ElkServiceList(w http.ResponseWriter, r *http.Request) {
	var err error
	vars := mux.Vars(r)

	var elkconfig = &elk.ElkConfig{
		Ls_p:  31533,
		Kib_p: 31535,
		Org:   vars["OrgId"],
	}

	body, err := ioutil.ReadAll(r.Body)

	if err != nil {
		w.Header().Set("Content-Type", "application/json; charset=UTF-8")
		w.WriteHeader(http.StatusBadRequest)
		if err := json.NewEncoder(w).Encode(jsonErr{Code: http.StatusNotFound, Text: "Wrong Format"}); err != nil {
			panic(err)
		}
	}

	json.Unmarshal(body, elkconfig)
	sl, err := elk.ElkServiceList(elkconfig)

	if err != nil {
		w.Header().Set("Content-Type", "application/json; charset=UTF-8")
		w.WriteHeader(http.StatusNotFound)
		if err := json.NewEncoder(w).Encode(jsonErr{Code: http.StatusNotFound, Text: "Not Found"}); err != nil {
			panic(err)
		}
	}

	w.Header().Set("Content-Type", "application/json; charset=UTF-8")
	w.WriteHeader(http.StatusOK)
	_body, err := json.Marshal(sl)

	if err != nil {
		logrus.Errorf("Failed to ElkServiceList ", err)
	}
	w.Write(_body)

	return
}

func ElkDeployList(w http.ResponseWriter, r *http.Request) {
	var err error
	vars := mux.Vars(r)

	var elkconfig = &elk.ElkConfig{
		Ls_p:  31533,
		Kib_p: 31535,
		Org:   vars["OrgId"],
	}

	body, err := ioutil.ReadAll(r.Body)

	if err != nil {
		w.Header().Set("Content-Type", "application/json; charset=UTF-8")
		w.WriteHeader(http.StatusBadRequest)
		if err := json.NewEncoder(w).Encode(jsonErr{Code: http.StatusNotFound, Text: "Wrong Format"}); err != nil {
			panic(err)
		}
	}

	json.Unmarshal(body, elkconfig)
	items, err := elk.ElkDeployList(elkconfig)

	if err != nil {
		w.Header().Set("Content-Type", "application/json; charset=UTF-8")
		w.WriteHeader(http.StatusNotFound)
		if err := json.NewEncoder(w).Encode(jsonErr{Code: http.StatusNotFound, Text: "Not Found"}); err != nil {
			panic(err)
		}
	}

	w.Header().Set("Content-Type", "application/json; charset=UTF-8")
	w.WriteHeader(http.StatusOK)
	_body, err := json.Marshal(items)

	if err != nil {
		logrus.Errorf("Failed to ElkDeployList ", err)
	}
	w.Write(_body)

	return
}
