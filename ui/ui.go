package ui

import (
	"fmt"
	"github.com/austindoolittle/spacetraders/client"
	"github.com/jroimartin/gocui"
	"github.com/mgutz/ansi"
)

const sidebarView = "sidebar"

type Engine struct {
	gui *gocui.Gui
	client *client.SpaceTradersClient
	sidebarController SidebarController
}

type SidebarController struct {
	menuItems []string
	selectedIdx int
	indexChanged bool
	highlightAnsi func(string) string
}

func max(a, b int) int {
	if a > b {
		return a
	}
	return b
}

func min(a, b int) int {
	if a < b {
		return a
	}
	return b
}

func sidebarItems() []string {
	return []string{"Ship Marketplace", "Planetary Map", "Stats"}
}

func newSidebarController() SidebarController {
	highlightAnsi := ansi.ColorFunc("red+B:black")

	var editor = SidebarController{sidebarItems(), 0, true, highlightAnsi}
	return editor
}

func (s *SidebarController) draw(v *gocui.View) {
	if !s.indexChanged {
		return
	}

	defer func() {s.indexChanged = false}()

	v.Clear()
	for i, val := range s.menuItems {
		err := v.SetCursor(0, 0)
		if err != nil {
			return
		}

		if i == s.selectedIdx {
			val = s.highlightAnsi(val)
		}

		_, err = fmt.Fprintln(v, val)
		if err != nil {
			return
		}
	}

}

//goland:noinspection GoUnusedParameter
func (s *SidebarController) handleUp(g *gocui.Gui, v *gocui.View) error {
	s.selectedIdx = max(s.selectedIdx-1, 0)
	s.indexChanged = true
	return nil
}

func (s *SidebarController) handleDown(*gocui.Gui, *gocui.View) error {
	s.selectedIdx = min(s.selectedIdx+1, len(s.menuItems)-1)
	s.indexChanged = true
	return nil
}

//goland:noinspection GoNilness
func (ui *Engine) redraw(gui *gocui.Gui) error {
	maxX, maxY := gui.Size()
	if maxX == 0 || maxY == 0 {
		return nil
	}

	sidebarViewObj, err := gui.SetView(sidebarView, 0, 0, maxX/4, maxY-1)
	ui.sidebarController.draw(sidebarViewObj)

	if err != nil && err != gocui.ErrUnknownView {
		return fmt.Errorf("failed to initialize sidebar view: %w", err)
	}
	sidebarViewObj.Frame = false

	bodyView, err := gui.SetView("body", maxX/4, 0, maxX-1, maxY-1)
	if err != nil && err != gocui.ErrUnknownView {
		return fmt.Errorf("failed to initialize body view: %w", err)
	}

	bodyView.Editable = false

	_, err = gui.SetCurrentView(sidebarView)
	if err != nil {
		return fmt.Errorf("failed to set current view to sidebar: %w", err)
	}

	return nil
}

//goland:noinspection GoUnusedParameter
func quit(g *gocui.Gui, v *gocui.View) error {
	return gocui.ErrQuit
}

func NewUiEngine() (*Engine, error) {
	g, err := gocui.NewGui(gocui.OutputNormal)
	if err != nil {
		return &Engine{}, fmt.Errorf("failed to create new UI engine: %w", err)
	}

	spacetradersClient := client.NewSpaceTradersClient(client.SpacetradersToken)

	engine := &Engine{g, spacetradersClient, newSidebarController()}
	g.SetManagerFunc(engine.redraw)
	if err := g.SetKeybinding("", gocui.KeyCtrlC, gocui.ModNone, quit); err != nil {
		_ = engine.Close()
		return &Engine{}, err
	}

	if err := g.SetKeybinding(sidebarView, 'k', gocui.ModNone, engine.sidebarController.handleUp); err != nil {
		return &Engine{}, err
	}

	if err := g.SetKeybinding(sidebarView, 'j', gocui.ModNone, engine.sidebarController.handleDown); err != nil {
		return &Engine{}, err
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