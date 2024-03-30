package main

import (
	"github.com/gdamore/tcell/v2"
	"github.com/rivo/tview"
)

type Project struct {
	name        string
	description string
	date        string
}

var projects []Project
var app = tview.NewApplication()
var projectText = tview.NewTextView()
var text = tview.NewTextView().
	SetTextColor(tcell.ColorGreen).
	SetText("(a) to add a new project \n(q) to quit")
var form = tview.NewForm()
var pages = tview.NewPages()

var projectsList = tview.NewList().ShowSecondaryText(false)

var flex = tview.NewFlex()

func main() {

	projectsList.SetSelectedFunc(func(index int, name string, description string, shortcut rune) {
		setConcatText(&projects[index])
	})

	pages.AddPage("Menu", flex, true, true)
	pages.AddPage("Add Project", form, true, false)

	flex.SetDirection(tview.FlexRow).
		AddItem(tview.NewFlex().AddItem(projectsList, 0, 1, true).AddItem(projectText, 0, 4, false), 0, 6, false).
		AddItem(text, 0, 1, false)

	app.SetInputCapture(func(event *tcell.EventKey) *tcell.EventKey {
		if event.Rune() == 113 {
			app.Stop()
		} else if event.Rune() == 97 {
			form.Clear(true)
			addProjectForm()
			pages.SwitchToPage("Add Project")
		}
		return event
	})

	if err := app.SetRoot(pages, true).EnableMouse(true).Run(); err != nil {
		panic(err)
	}
}

func addProjectForm() {
	project := Project{}

	form.AddInputField("Name", "", 20, nil, func(name string) {
		project.name = name
	})

	form.AddInputField("Description", "", 20, nil, func(description string) {
		project.description = description
	})

	form.AddInputField("Date", "", 20, nil, func(date string) {
		project.date = date
	})

	form.AddButton("Save", func() {
		projects = append(projects, project)
		addProjectList()
		pages.SwitchToPage("Menu")
	})
}

func addProjectList() {
	projectsList.Clear()
	for index, project := range projects {
		projectsList.AddItem(project.name+" "+project.description, " ", rune(49+index), nil)
	}
}

func setConcatText(project *Project) {
	projectText.Clear()
	text := project.name + " " + project.description + "\n" + project.date
	projectText.SetText(text)
}
