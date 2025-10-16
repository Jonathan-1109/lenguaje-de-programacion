package functions

import (
	"fmt"
	"taskManager/data"
	colors "taskManager/utils"

	"github.com/spf13/cobra"
)

var DeleteTask = &cobra.Command{
	Use:   "delete [id]",
	Short: "Borra una tarea",
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

		tasks = append(tasks[:index], tasks[index+1:]...)

		err = ChangeJson(data.Route, tasks)
		if err != nil {
			fmt.Println(colors.Red, err, colors.Reset)
			return
		}

		fmt.Println(colors.Green, "tarea borrada exitosamente", colors.Reset)
	},
}
