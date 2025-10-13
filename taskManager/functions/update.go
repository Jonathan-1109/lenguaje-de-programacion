package functions

import (
	"fmt"
	data "taskManager/data"
	colors "taskManager/utils"
	"time"

	"github.com/spf13/cobra"
)

var UpdateTask = &cobra.Command{
	Use:   "update [id]",
	Short: "Actualizar una tarea a su siguiente estado",
	Args:  cobra.ExactArgs(1),
	Run: func(cmd *cobra.Command, args []string) {
		flagChange, _ := cmd.Flags().GetBool("statusChange")
		flagName, _ := cmd.Flags().GetString("name")
		flagDescription, _ := cmd.Flags().GetString("description")
		band := false
		id := args[0]
		var tasks []data.Task

		err := ReadJson(data.Route, &tasks)
		if err != nil {
			fmt.Println(colors.Red, err, colors.Reset)
			return
		}

		index, err := GetIndex(id, tasks)
		if err != nil {
			fmt.Println(colors.Red, err, colors.Reset)
			return
		}

		task := &tasks[index]

		if flagChange {

			if task.State == data.Completed {
				fmt.Println("La tarea ya ha sido completada")
			} else {
				task.State++
				band = true
				fmt.Println("Tarea se actualizo a:", task.State)
			}
		}

		if flagName != "" {
			task.Name = flagName
			fmt.Println("Nombre de la tarea cambio a:", task.Name)
			band = true
		}

		if flagDescription != "" {
			task.Description = flagDescription
			fmt.Println("Descripción de la tarea cambio a:", task.Description)
			band = true
		}

		if band {
			task.Update = time.Now().Format("02/01/2006 03:04 PM")
		}

		err = ChangeJson(data.Route, tasks)
		if err != nil {
			fmt.Println(colors.Red, err, colors.Reset)
			return
		}
		fmt.Println(colors.Green, "Tarea actualizada", colors.Reset)
	},
}
