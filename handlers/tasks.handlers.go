package handlers

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
	"net/http"
	"strconv"

	"github.com/gorilla/mux"
	"github.com/sergioriverafl/go-rest-api/data"
	"github.com/sergioriverafl/go-rest-api/models"
)

func CreateTask(w http.ResponseWriter, r *http.Request) {
	var newTask models.Task

	reqBody, err := ioutil.ReadAll(r.Body)

	if err != nil {
		fmt.Fprintf(w, "Insert a valid task")
	}

	json.Unmarshal(reqBody, &newTask)

	newTask.ID = len(data.TasksData) + 1
	data.TasksData = append(data.TasksData, newTask)

	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusCreated)
	json.NewEncoder(w).Encode(newTask)
}

func GetTasks(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode(data.TasksData)
}

func GetOneTask(w http.ResponseWriter, r *http.Request) {
	vars := mux.Vars(r)
	taskID, err := strconv.Atoi(vars["id"])

	if err != nil {
		fmt.Fprintf(w, "Invalid ID")
	}

	w.Header().Set("Content-Type", "application/json")

	for _, task := range data.TasksData {
		if task.ID == taskID {
			json.NewEncoder(w).Encode(task)
		}
	}
}

func DeleteTask(w http.ResponseWriter, r *http.Request) {
	vars := mux.Vars(r)
	taskID, err := strconv.Atoi(vars["id"])

	if err != nil {
		fmt.Fprintf(w, "Invalid ID")
	}

	for i, task := range data.TasksData {
		if task.ID == taskID {
			data.TasksData = append(data.TasksData[:i], data.TasksData[i+1:]...)
			w.WriteHeader(http.StatusAccepted)
			fmt.Fprintf(w, "The task witch ID %v has been remove successfully", taskID)
		}
	}
}

func UpdateTask(w http.ResponseWriter, r *http.Request) {

	vars := mux.Vars(r)
	taskID, err := strconv.Atoi(vars["id"])
	var updatedTask models.Task

	if err != nil {
		fmt.Fprintf(w, "Invalid ID")
	}

	reqBody, err := ioutil.ReadAll(r.Body)

	if err != nil {
		fmt.Fprintf(w, "Please enter valid data")
	}

	json.Unmarshal(reqBody, &updatedTask)

	for i, task := range data.TasksData {
		if task.ID == taskID {
			data.TasksData = append(data.TasksData[:i], data.TasksData[i+1:]...)
			updatedTask.ID = taskID
			data.TasksData = append(data.TasksData, updatedTask)

			fmt.Fprintf(w, "The task witch ID %v has been updated", taskID)
		}
	}
}
