package main

import (
	"github.com/gdamore/tcell/v2"
	"github.com/rivo/tview"
	"log"
	"os"
)

type item struct {
	is_dir bool
	name   string
}

func getItems() []item {
	result := []item{}
	entries, err := os.ReadDir("./")
	if err != nil {
		log.Fatal(err)
	}

	for _, e := range entries {
		i := item{
			is_dir: e.IsDir(),
			name:   e.Name(),
		}
		result = append(result, i)
	}
	return result
}

func main() {
	app := tview.NewApplication()

	list := tview.NewList().SetMainTextStyle(tcell.StyleDefault).SetSecondaryTextStyle(tcell.StyleDefault)
	for _, i := range getItems() {
		list.AddItem(i.name, "", 0, nil)
	}

	list.SetInputCapture(func(event *tcell.EventKey) *tcell.EventKey {
		switch event.Key() {
		case tcell.KeyRune:
			switch event.Rune() {
			case 'j': // Move down
				list.SetCurrentItem(list.GetCurrentItem() + 1)
			case 'k': // Move up
				list.SetCurrentItem(list.GetCurrentItem() - 1)
			case 'q': // Quit application
				app.Stop()
			}
		}
		return event
	})

	if err := app.SetRoot(list, true).Run(); err != nil {
		panic(err)
	}
}
