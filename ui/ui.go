package ui

import (
	"fmt"
	"github.com/austindoolittle/spacetraders/client"
	"github.com/jroimartin/gocui"
)

type Engine struct {
	gui *gocui.Gui
	client *client.SpaceTradersClient
}

//goland:noinspection GoNilness
func (ui Engine) redraw(gui *gocui.Gui) error {
	maxX, maxY := gui.Size()

	sidebarView, err := gui.SetView("sidebar", 0, 0, maxX/3, maxY-1)

	if err != nil && err != gocui.ErrUnknownView {
		return fmt.Errorf("failed to initialize sidebar view: %w", err)
	}

	//goland:noinspection GoNilness
	sidebarView.Clear()

	_, err = fmt.Fprintln(sidebarView, "sidebar!")
	if err != nil {
		return fmt.Errorf("failed to write to sidebar view: %w", err)
	}

	bodyView, err := gui.SetView("body", maxX/3, 0, maxX-1, maxY-1)
	if err != nil && err != gocui.ErrUnknownView {
		return fmt.Errorf("failed to initialize body view: %w", err)
	}

	bodyView.Clear()
	_, err = fmt.Fprintln(bodyView, "body!")
	if err != nil {
		return fmt.Errorf("failed to write to body view: %w", err)
	}

	return nil
}

//goland:noinspection GoUnusedParameter
func quit(g *gocui.Gui, v *gocui.View) error {
	fmt.Println("CTRL+C, We're done here")
	return gocui.ErrQuit
}

func NewUiEngine() (Engine, error) {
	var engine Engine

	g, err := gocui.NewGui(gocui.OutputNormal)
	if err != nil {
		return engine, fmt.Errorf("failed to create new UI engine: %w", err)
	}

	spacetradersClient := client.NewSpaceTradersClient(client.SpacetradersToken)

	engine = Engine{g, spacetradersClient}
	g.SetManagerFunc(engine.redraw)
	if err := g.SetKeybinding("", gocui.KeyCtrlC, gocui.ModNone, quit); err != nil {
		_ = engine.Close()
		return engine, err
	}
	g.Mouse = true

	return engine, nil
}

func (ui Engine) Run() error {
	if err := ui.gui.MainLoop(); err != nil && err != gocui.ErrQuit {
		return err
	}

	return nil
}

func (ui Engine) Close() error {
	return ui.Close()
}