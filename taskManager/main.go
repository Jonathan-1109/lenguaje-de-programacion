package main

import (
	command "taskManager/functions"

	"github.com/spf13/cobra"
)

var RootCmd = &cobra.Command{
	Use:   "task",
	Short: "Gestión de Tareas",
}

func main() {
	command.ExistDir()
	RootCmd.AddCommand(command.AddTask)
	RootCmd.AddCommand(command.UpdateTask)
	command.UpdateTask.Flags().BoolP("statusChange", "s", true, "Cambiar estado de la tarea")
	command.UpdateTask.Flags().StringP("name", "n", "", "Cambiar nombre de la tarea")
	command.UpdateTask.Flags().StringP("description", "d", "", "Cambiar descripción de la tarea")
	RootCmd.AddCommand(command.DeleteTask)
	RootCmd.AddCommand(command.ListTask)
	RootCmd.AddCommand(command.GetTask)
	RootCmd.Execute()
}
