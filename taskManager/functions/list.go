package functions

import (
	"fmt"
	"slices"
	"strings"
	"taskManager/data"
	colors "taskManager/utils"

	"github.com/spf13/cobra"
)

var ListTask = &cobra.Command{
	Use:   "list [status]",
	Short: "Lista todas las tareas o aquellas en cierto estado: (All, Pending, InProgress, Completed)",
	Args:  cobra.MaximumNArgs(1),
	Run: func(cmd *cobra.Command, args []string) {
		var tasks []data.Task
		var status = ""
		if len(args) == 1 {
			status = strings.ToLower(args[0])
		} else {
			status = "all"
		}

		if !slices.Contains(data.ValidStatus, status) {
			fmt.Println(colors.Red, "Error: estado invalido", colors.Reset)
			cmd.Help()
			return
		}

		err := ReadJson(data.Route, &tasks)
		if err != nil {
			fmt.Println(colors.Red, "Error obteniendo las tareas", colors.Reset)
			return
		}
		if len(tasks) == 0 {
			fmt.Println(colors.Yellow, "No hay Tareas", colors.Reset)
			return
		}

		if status == "all" {
			fmt.Println("Todas las tareas: ")
		} else {
			fmt.Println("Tareas en:", status)
		}

		for _, task := range tasks {
			if status == "all" || task.State == data.State(slices.Index(data.ValidStatus, status)-1) {
				fmt.Printf("ID: %d, Nombre: %s, Estado: %s\n", task.ID, task.Name, task.State)
			}
		}
	},
}
