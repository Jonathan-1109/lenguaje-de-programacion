package functions

import (
	"fmt"
	"taskManager/data"
	colors "taskManager/utils"

	"github.com/spf13/cobra"
)

var GetTask = &cobra.Command{
	Use:   "get [id]",
	Short: "Obtener una tarea",
	Args:  cobra.ExactArgs(1),
	Run: func(cmd *cobra.Command, args []string) {
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

		fmt.Println("ID:", task.ID,
			"\nNombre:", task.Name,
			"\nDescripción:", task.Description,
			"\nEstado:", task.State,
			"\nCreación:", task.Created,
			"\nActualzación reciente:", task.Update)
	},
}
