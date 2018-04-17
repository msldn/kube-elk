package kube_elk

import (
	"encoding/json"
	"fmt"
	"io"
	"io/ioutil"
	"net/http"
	"strconv"
	"./elk"
	"github.com/gorilla/mux"
	"github.com/sirupsen/logrus"
)

func Index(w http.ResponseWriter, r *http.Request) {
	fmt.Fprint(w, "Welcome!\n")
}

func TodoIndex(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json; charset=UTF-8")
	w.WriteHeader(http.StatusOK)
	if err := json.NewEncoder(w).Encode(todos); err != nil {
		panic(err)
	}
}

func TodoShow(w http.ResponseWriter, r *http.Request) {
	vars := mux.Vars(r)
	var todoId int
	var err error
	if todoId, err = strconv.Atoi(vars["todoId"]); err != nil {
		panic(err)
	}
	todo := RepoFindTodo(todoId)
	if todo.Id > 0 {
		w.Header().Set("Content-Type", "application/json; charset=UTF-8")
		w.WriteHeader(http.StatusOK)
		if err := json.NewEncoder(w).Encode(todo); err != nil {
			panic(err)
		}
		return
	}

	// If we didn't find it, 404
	w.Header().Set("Content-Type", "application/json; charset=UTF-8")
	w.WriteHeader(http.StatusNotFound)
	if err := json.NewEncoder(w).Encode(jsonErr{Code: http.StatusNotFound, Text: "Not Found"}); err != nil {
		panic(err)
	}

}

/*
Test with this curl command:
curl -H "Content-Type: application/json" -d '{"name":"New Todo"}' http://localhost:8080/todos
*/
func TodoCreate(w http.ResponseWriter, r *http.Request) {
	var todo Todo
	body, err := ioutil.ReadAll(io.LimitReader(r.Body, 1048576))
	if err != nil {
		panic(err)
	}
	if err := r.Body.Close(); err != nil {
		panic(err)
	}
	if err := json.Unmarshal(body, &todo); err != nil {
		w.Header().Set("Content-Type", "application/json; charset=UTF-8")
		w.WriteHeader(422) // unprocessable entity
		if err := json.NewEncoder(w).Encode(err); err != nil {
			panic(err)
		}
	}

	t := RepoCreateTodo(todo)
	w.Header().Set("Content-Type", "application/json; charset=UTF-8")
	w.WriteHeader(http.StatusCreated)
	if err := json.NewEncoder(w).Encode(t); err != nil {
		panic(err)
	}
}



func TodoDelete(w http.ResponseWriter, r *http.Request) {
	vars := mux.Vars(r)
	var todoId int
	var err error
	if todoId, err = strconv.Atoi(vars["todoId"]); err != nil {
		panic(err)
	}

	todoIdx := RepoDeleteTodo(todoId)

	if todoIdx > 0 {
		w.Header().Set("Content-Type", "application/json; charset=UTF-8")
		w.WriteHeader(http.StatusOK)
		return
	}

	// If we didn't find it, 404
	w.Header().Set("Content-Type", "application/json; charset=UTF-8")
	w.WriteHeader(http.StatusNotFound)
	if err := json.NewEncoder(w).Encode(jsonErr{Code: http.StatusNotFound, Text: "Not Found"}); err != nil {
		panic(err)
	}
}


func ElkDelete(w http.ResponseWriter, r *http.Request) {
	vars := mux.Vars(r)
	var OrgId = vars["OrgId"]
	var err error

	var elkconfig = &elk.ElkConfig{
		Org: OrgId,
	}

	err = elk.ElkDelete("default", elkconfig)

	if err != nil  {
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
		Ls_p:31533,
		Kib_p:31535,
		Org: "org3",
	}
	body, err := ioutil.ReadAll(r.Body)
	if err != nil  {
		w.Header().Set("Content-Type", "application/json; charset=UTF-8")
		w.WriteHeader(http.StatusBadRequest)
		if err := json.NewEncoder(w).Encode(jsonErr{Code: http.StatusNotFound, Text: "Wrong Format"}); err != nil {
			panic(err)
		}
	}

	json.Unmarshal(body,elkconfig)
	err = elk.ElkCreate("default", elkconfig)

	if err != nil  {
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
		Ls_p:31533,
		Kib_p:31535,
		Org: vars["OrgId"],
	}
	body, err := ioutil.ReadAll(r.Body)
	if err != nil  {
		w.Header().Set("Content-Type", "application/json; charset=UTF-8")
		w.WriteHeader(http.StatusBadRequest)
		if err := json.NewEncoder(w).Encode(jsonErr{Code: http.StatusNotFound, Text: "Wrong Format"}); err != nil {
			panic(err)
		}
	}

	json.Unmarshal(body,elkconfig)
	err = elk.ElkGet("default", elkconfig)

	if err != nil  {
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


func ElkServiceList(w http.ResponseWriter, r *http.Request) {
	var err error
	vars := mux.Vars(r)

	var elkconfig = &elk.ElkConfig{
		Ls_p:31533,
		Kib_p:31535,
		Org: vars["OrgId"],
	}

	body, err := ioutil.ReadAll(r.Body)

	if err != nil  {
		w.Header().Set("Content-Type", "application/json; charset=UTF-8")
		w.WriteHeader(http.StatusBadRequest)
		if err := json.NewEncoder(w).Encode(jsonErr{Code: http.StatusNotFound, Text: "Wrong Format"}); err != nil {
			panic(err)
		}
	}

	json.Unmarshal(body,elkconfig)
	sl, err := elk.ElkServiceList("default", elkconfig)

	if err != nil  {
		w.Header().Set("Content-Type", "application/json; charset=UTF-8")
		w.WriteHeader(http.StatusNotFound)
		if err := json.NewEncoder(w).Encode(jsonErr{Code: http.StatusNotFound, Text: "Not Found"}); err != nil {
			panic(err)
		}
	}

	w.Header().Set("Content-Type", "application/json; charset=UTF-8")
	w.WriteHeader(http.StatusOK)
	_body,err := json.Marshal(sl)

	if err != nil {
		logrus.Errorf("Failed to ElkServiceList ",err)
	}
	w.Write(_body)

	return
}

func ElkDeployList(w http.ResponseWriter, r *http.Request) {
	var err error
	vars := mux.Vars(r)

	var elkconfig = &elk.ElkConfig{
		Ls_p:31533,
		Kib_p:31535,
		Org: vars["OrgId"],
	}

	body, err := ioutil.ReadAll(r.Body)

	if err != nil  {
		w.Header().Set("Content-Type", "application/json; charset=UTF-8")
		w.WriteHeader(http.StatusBadRequest)
		if err := json.NewEncoder(w).Encode(jsonErr{Code: http.StatusNotFound, Text: "Wrong Format"}); err != nil {
			panic(err)
		}
	}

	json.Unmarshal(body,elkconfig)
	items, err := elk.ElkDeployList("default", elkconfig)

	if err != nil  {
		w.Header().Set("Content-Type", "application/json; charset=UTF-8")
		w.WriteHeader(http.StatusNotFound)
		if err := json.NewEncoder(w).Encode(jsonErr{Code: http.StatusNotFound, Text: "Not Found"}); err != nil {
			panic(err)
		}
	}

	w.Header().Set("Content-Type", "application/json; charset=UTF-8")
	w.WriteHeader(http.StatusOK)
	_body,err := json.Marshal(items)

	if err != nil {
		logrus.Errorf("Failed to ElkDeployList ",err)
	}
	w.Write(_body)

	return
}