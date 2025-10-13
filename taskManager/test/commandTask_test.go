package test

import (
	fun "taskManager/functions"
	"testing"

	"github.com/spf13/cobra"
)

func TestAdd(t *testing.T) {
	var TestCmd = &cobra.Command{
		Use:   "test",
		Short: "Gestión de Tareas",
		Run:   fun.AddTask.Run,
	}
	TestCmd.SetArgs([]string{"nombre", "descripcion"})
	err := fun.ExistDir()
	if err != nil {
		t.Errorf("Error verificando directorio %v", err)
	}
	err = TestCmd.Execute()
	if err != nil {
		t.Errorf("Error al añadir: %v", err)
	}
}

func TestDelete(t *testing.T) {
	var TestCmd = &cobra.Command{
		Use:   "test",
		Short: "Gestión de Tareas",
		Run:   fun.DeleteTask.Run,
	}
	TestCmd.SetArgs([]string{"2"})
	err := fun.ExistDir()
	if err != nil {
		t.Errorf("Error verificando directorio %v", err)
	}
	err = TestCmd.Execute()
	if err != nil {
		t.Errorf("Error al borrar: %v", err)
	}
}

func TestUpdate(t *testing.T) {
	var TestCmd = &cobra.Command{
		Use:   "test",
		Short: "Gestión de Tareas",
		Run:   fun.UpdateTask.Run,
	}
	TestCmd.Flags().BoolP("statusChange", "s", true, "Cambiar estado de la tarea")
	TestCmd.Flags().StringP("name", "n", "", "Cambiar nombre de la tarea")
	TestCmd.Flags().StringP("description", "d", "", "Cambiar descripción de la tarea")
	err := fun.ExistDir()
	if err != nil {
		t.Errorf("Error verificando directorio %v", err)
	}
	TestCmd.SetArgs([]string{"1", "-n", "nuevo", "-d", "des"})
	err = TestCmd.Execute()
	if err != nil {
		t.Errorf("Error al actualizar: %v", err)
	}
}

func TestGet(t *testing.T) {
	var TestCmd = &cobra.Command{
		Use:   "test",
		Short: "Gestión de Tareas",
		Run:   fun.GetTask.Run,
	}
	err := fun.ExistDir()
	if err != nil {
		t.Errorf("Error verificando directorio %v", err)
	}
	TestCmd.SetArgs([]string{"2"})
	err = TestCmd.Execute()
	if err != nil {
		t.Errorf("Error al conseguir información: %v", err)
	}
}

func TestList(t *testing.T) {
	var TestCmd = &cobra.Command{
		Use:   "test",
		Short: "Gestión de Tareas",
		Run:   fun.ListTask.Run,
	}
	err := fun.ExistDir()
	if err != nil {
		t.Errorf("Error verificando directorio %v", err)
	}
	TestCmd.SetArgs([]string{"all"})
	err = TestCmd.Execute()
	if err != nil {
		t.Errorf("Error al conseguir información: %v", err)
	}
	TestCmd.SetArgs([]string{"inprogress"})
	err = TestCmd.Execute()
	if err != nil {
		t.Errorf("Error al conseguir información: %v", err)
	}
	TestCmd.SetArgs([]string{"completed"})
	err = TestCmd.Execute()
	if err != nil {
		t.Errorf("Error al conseguir información: %v", err)
	}
}
