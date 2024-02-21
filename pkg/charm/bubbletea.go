package charm

import (
	"fmt"
	"os"

	"github.com/charmbracelet/bubbles/list"
	tea "github.com/charmbracelet/bubbletea"
	"github.com/charmbracelet/lipgloss"
)

var docStyle = lipgloss.NewStyle().Margin(1, 2)

type item struct {
	title, desc string
}

func (i item) Title() string       { return i.title }
func (i item) Description() string { return i.desc }
func (i item) FilterValue() string { return i.title }

// AppState represents different states in the application.
type AppState int

const (
	MainMenu AppState = iota
	StockMarketData
	StockMarketNews
)

type model struct {
	list  list.Model
	state AppState
}

func (m model) Init() tea.Cmd {
	return nil
}

func (m model) Update(msg tea.Msg) (tea.Model, tea.Cmd) {
	switch msg := msg.(type) {
	case tea.KeyMsg:
		switch msg.String() {
		case "ctrl+c", "q":
			return m, tea.Quit

		// Check if "enter" is pressed.
		case "enter", " ":
			if selectedItem, ok := m.list.SelectedItem().(item); ok {
				switch selectedItem.title {
				case "Stock Market Data":
					m.state = StockMarketData
					fmt.Println("Navigating to Stock Market Data...")
					//TODO:  Replace with actual logic or navigation
					return m, nil
				case "Stock Market News":
					m.state = StockMarketNews
					fmt.Println("Navigating to Stock Market News...")
					//TODO:  Replace with actual logic or navigation
					return m, nil
				}
			}
		}
	case tea.WindowSizeMsg:
		h, v := docStyle.GetFrameSize()
		m.list.SetSize(msg.Width-h, msg.Height-v)
	}

	var cmd tea.Cmd
	m.list, cmd = m.list.Update(msg)
	return m, cmd
}

func (m model) View() string {
	switch m.state {
	case MainMenu:
		return docStyle.Render(m.list.View())
	case StockMarketData:
		return "Viewing Stock Market Data..."
	case StockMarketNews:
		return "Viewing Stock Market News..."
	}
	// Default view
	return ""
}

func Start() {
	items := []list.Item{
		item{title: "Stock Market Data", desc: "view stock prices and dividends."},
		item{title: "Stock Market News", desc: "view financial news."},
	}

	m := model{list: list.New(items, list.NewDefaultDelegate(), 0, 0)}
	m.list.Title = "Welcome to Mango Markets"

	p := tea.NewProgram(m, tea.WithAltScreen())

	if _, err := p.Run(); err != nil {
		fmt.Println("Error running program:", err)
		os.Exit(1)
	}
}
