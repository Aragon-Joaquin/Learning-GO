package main

import (
	"fmt"

	"github.com/gdamore/tcell/v2"
	"github.com/rivo/tview"
)

type TaskShape struct {
	Name string
	IsDone bool
}

var ListOfTasks = []TaskShape{}

func main() {
	newApp := tview.NewApplication()

	tasksViews := tview.NewTextView().SetDynamicColors(true).SetWordWrap(true)
	tasksViews.SetBorder(true).SetTitle("Task List")

	refreshTasks := func ()  {
		tasksViews.Clear()

		if(len(ListOfTasks) == 0){
			fmt.Fprintln(tasksViews, "No tasks added yet")
			return;
		}

		for index, task := range ListOfTasks {
			fmt.Fprintf(tasksViews, "%d. %s - Status:[%t] ",index + 1, task.Name, task.IsDone)
		}
	}

	nameInput := tview.NewInputField().SetLabel("Task Name: ")
	// deleteInput := tview.NewInputField().SetLabel("Write the task name to delete it: ")

	form := tview.NewForm().AddFormItem(nameInput).AddButton("Add task", func () {
		name := nameInput.GetText()

		if name == "" {
			return;
		}

		ListOfTasks = append(ListOfTasks, TaskShape{Name: name, IsDone: false})
		refreshTasks()
		nameInput.SetText("")
	}).AddButton("Delete all", func () {
		ListOfTasks = []TaskShape{}
	}).AddButton("Stop", func() {
		newApp.Stop()
	})

	form.SetBorder(true).SetTitle("Dashboard?").SetTitleAlign(tview.AlignLeft).SetTitleColor(tcell.ColorDarkRed)

	refreshTasks()

	layout := tview.NewFlex().AddItem(form,0,1,true).AddItem(tasksViews,0,1,false)

	if err := newApp.SetRoot(layout, true).Run(); err != nil {
		panic(err)
	}
	
}


