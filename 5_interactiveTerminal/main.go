package main

import (
	"grid"
	"interactiveTerminal/utils"

	"github.com/rivo/tview"
)

func main() {
	newApp := tview.NewApplication()
	grid.InitializeCinema()
	
	cinemaGrid := tview.NewGrid().SetRows(utils.ROWS).SetColumns(utils.COLUMNS)

	grid.PaintGridLayout(cinemaGrid);

	cinemaGrid.SetTitle("Seats").SetBorder(true)




	if err := newApp.SetRoot(cinemaGrid, true).Run(); err != nil {
		panic(err)
	}
}




// refreshTasks := func ()  {
// 	tasksViews.Clear()

// 	if(len(ListOfTasks) == 0){
// 		fmt.Fprintln(tasksViews, "No tasks added yet")
// 		return;
// 	}

// 	for index, task := range ListOfTasks {
// 		fmt.Fprintf(tasksViews, "%d. %s - Status:[%t] ",index + 1, task.Name, task.IsDone)
// 	}
// }

// nameInput := tview.NewInputField().SetLabel("Task Name: ")
// // deleteInput := tview.NewInputField().SetLabel("Write the task name to delete it: ")

// form := tview.NewForm().AddFormItem(nameInput).AddButton("Add task", func () {
// 	name := nameInput.GetText()

// 	if name == "" {
// 		return;
// 	}

// 	ListOfTasks = append(ListOfTasks, TaskShape{Name: name, IsDone: false})
// 	refreshTasks()
// 	nameInput.SetText("")
// }).AddButton("Delete all", func () {
// 	ListOfTasks = []TaskShape{}
// }).AddButton("Stop", func() {
// 	newApp.Stop()
// })


// refreshTasks()

// mainLayout := tview.NewFlex().AddItem(form,0,1,true).AddItem(tasksViews,0,1,false)

// if err := newApp.SetRoot(mainLayout, true).Run(); err != nil {
// 	panic(err)
// }
