package grid

import (
	"interactiveTerminal/utils"

	"github.com/rivo/tview"
)

type CinemaShape struct {
	SeatNumber int
	IsOccupied bool
}

var Cinema = [utils.ROWS][utils.COLUMNS]CinemaShape{}


func InitializeCinema () {
	var seatNumber int;
	for i := range Cinema {
		for j := range Cinema[i] {
			seatNumber++;
				Cinema[i][j] = CinemaShape{SeatNumber: seatNumber, IsOccupied: false}
		}
}
}

func PaintGridLayout(view *tview.Grid) {
	for i:= range utils.COLUMNS {
		for j := range utils.ROWS {
			view.AddItem(tview.NewBox().SetBorder(true), i, j, 1, 1, 10, 10, true)
			
		}
	}
}