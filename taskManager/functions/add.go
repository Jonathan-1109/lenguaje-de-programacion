package functions

import (
	"fmt"
	"taskManager/data"
	colors "taskManager/utils"
	"time"

	"github.com/spf13/cobra"
)

var AddTask = &cobra.Command{
	Use:   "add [name] [description]",
	Short: "Añadir una tarea con nombre y descripción",
	Args:  cobra.MatchAll(cobra.MinimumNArgs(1), cobra.MaximumNArgs(2)),
	Run: func(cmd *cobra.Command, args []string) {

		err := ExistDir()
		if err != nil {
			fmt.Println(colors.Red, err, colors.Reset)
			return
		}
		name := args[0]
		var description = ""
		if len(args) > 1 {
			description = args[1]
		}

		t := time.Now().Format("02/01/2006 03:04 PM")
		id, err := GetID()
		var newTasks []data.Task

		if err != nil {
			fmt.Println(colors.Red, err, colors.Reset)
			return
		}

		task := data.Task{ID: id, Name: name, Created: t, State: data.Pending, Description: description}
		err = ReadJson(data.Route, &newTasks)
		if err != nil {
			fmt.Println(err)
			return
		}

		newTasks = append(newTasks, task)
		err = ChangeJson(data.Route, newTasks)
		if err != nil {
			fmt.Println(colors.Red, err, colors.Reset)
			return
		}

		fmt.Println(colors.Green, "Tarea creada ID:", id, colors.Reset)
	},
}
