package test

import (
	"fmt"
	"os"
	data "taskManager/data"
	fun "taskManager/functions"
	"testing"
	"time"
)

func TestExistDir(t *testing.T) {
	os.RemoveAll("./tasks")
	err := fun.ExistDir()
	if err != nil {
		t.Fatalf("ExistDir falló: %v", err)
	}
}

func TestChangeJson(t *testing.T) {
	now := time.Now().Format("02/01/2006 03:04 PM")
	tasks := data.Task{ID: 1, Name: "Nombre", Description: "Descripcion", Created: now, State: data.Pending}
	err := fun.ChangeJson("./tasks/data.json", &tasks)
	if err != nil {
		t.Fatalf("ChangeJson falló: %v", err)
	}
	os.RemoveAll("./tasks")
}
func TestReadJson(t *testing.T) {
	var tasks []data.Task
	var config data.Config
	err := fun.ReadJson("./tasks/data.json", &tasks)
	if err != nil {
		t.Fatalf("ReadJson falló: %v", err)
	}
	fmt.Println(tasks)
	err = fun.ReadJson("./tasks/config.json", &config)
	if err != nil {
		t.Fatalf("ReadJson falló: %v", err)
	}
	fmt.Println(config.ActualID)
}

func TestGetID(t *testing.T) {
	id, err := fun.GetID()
	if err != nil {
		t.Fatalf("GetID falló: %v", err)
	}
	if id == 1 {
		fmt.Println("ID obtenido es 1")
	}
	fmt.Printf("\nID obtenido: %d", id)
}

func TestGetIndex(t *testing.T) {
	now := time.Now().Format("02/01/2006 03:04 PM")
	tasks := []data.Task{{ID: 5, Name: "Nombre", Description: "Descripcion", Created: now, State: data.Pending}}
	err := fun.ChangeJson("./tasks/data.json", tasks)
	if err != nil {
		t.Fatalf("ChangeJson falló: %v", err)
	}

	_, err = fun.GetIndex("5", tasks)
	if err != nil {
		t.Fatalf("GetIndex falló: %v", err)
	}
	os.RemoveAll("./tasks")
}
