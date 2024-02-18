package charm

import (
	"fmt"
	"os"

	"github.com/charmbracelet/bubbles/list"
	tea "github.com/charmbracelet/bubbletea"
	"github.com/charmbracelet/lipgloss"
	"github.com/jp-mango/mangomarkets/internal/api"
)

var (
	appStyle = lipgloss.NewStyle().Padding(1, 2)

	titleStyle = lipgloss.NewStyle().
			Foreground(lipgloss.Color("#FFFDF5")).
			Background(lipgloss.Color("#25A065")).
			Padding(0, 1)
)

type item struct {
	title       string
	description string
}

func (i item) Title() string       { return i.title }
func (i item) Description() string { return i.description }
func (i item) FilterValue() string { return i.title }

type model struct {
	list list.Model
}

func newModel(items []list.Item) model {
	l := list.New(items, list.NewDefaultDelegate(), 0, 0)
	l.Title = "Welcome Screen"
	l.SetShowStatusBar(false)
	l.SetShowFilter(false)
	l.Styles.Title = titleStyle

	return model{list: l}
}

func (m model) Init() tea.Cmd {
	return nil
}

func (m model) Update(msg tea.Msg) (tea.Model, tea.Cmd) {
	switch msg := msg.(type) {
	case tea.KeyMsg:
		switch msg.String() {
		case "q", "ctrl+c":
			return m, tea.Quit
		case "enter":
			i, ok := m.list.SelectedItem().(item)
			if !ok {
				return m, nil
			}
			switch i.title {
			case "View Stock Data":
				fmt.Println("Fetching stock data for IBM...")
				stockData, err := api.FetchTimeSeriesDaily("YOUR_API_KEY", "IBM")
				if err != nil {
					fmt.Printf("Error fetching stock data: %v\n", err)
					return m, tea.Quit
				}
				fmt.Printf("Latest stock data for IBM: %+v\n", stockData.MetaData)
				return m, tea.Quit
			}
		}
	}

	var cmd tea.Cmd
	m.list, cmd = m.list.Update(msg)
	return m, cmd
}

func (m model) View() string {
	return appStyle.Render(m.list.View())
}

func Start() {
	items := []list.Item{
		item{title: "View Stock Data", description: ""},
		item{title: "View Crypto Data", description: ""},
		item{title: "View Forex Data", description: ""},
		item{title: "View Stock News", description: ""},
	}

	p := tea.NewProgram(newModel(items))
	if err, _ := p.Run(); err != nil {
		fmt.Printf("Error running program: %v", err)
		os.Exit(1)
	}
}
