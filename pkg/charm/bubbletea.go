package charm

import (
	"fmt"
	"os"
	"strings"

	"github.com/charmbracelet/bubbles/list"
	"github.com/charmbracelet/bubbles/textinput"
	tea "github.com/charmbracelet/bubbletea"
	"github.com/jp-mango/mangomarkets/internal/api"
	"github.com/jp-mango/mangomarkets/internal/config"
	"github.com/jp-mango/mangomarkets/util"
)

const listHeight = 14

type AppState int

const (
	MainMenu AppState = iota
	TickerInput
	TimeSeriesSelection
	DisplayData
)

type item string

func (i item) Title() string       { return string(i) }
func (i item) Description() string { return "" }
func (i item) FilterValue() string { return "" }

type dataMsg struct {
	data []string
	err  error
}

type model struct {
	list             list.Model
	textInput        textinput.Model
	state            AppState
	data             []string
	currentPage      int
	itemsPerPage     int
	fetchErrorMsg    string
	timeSeriesChoice string
}

func newModel() model {
	items := []list.Item{
		item("Stock Market"),
		item("Forex & Currencies"),
		item("Cryptocurrency"),
		item("Economic News"),
	}

	l := list.New(items, list.NewDefaultDelegate(), 50, listHeight)
	l.Title = "Welcome to Mango Markets"
	l.SetShowStatusBar(false)

	ti := textinput.New()
	ti.Placeholder = "Enter Ticker Symbol"
	ti.Focus()
	ti.Prompt = "> "

	return model{
		list:         l,
		textInput:    ti,
		state:        MainMenu,
		currentPage:  0,
		itemsPerPage: 5, // Adjust based on preference
	}
}

func (m model) Init() tea.Cmd {
	return nil
}

func (m model) Update(msg tea.Msg) (tea.Model, tea.Cmd) {
	switch m.state {
	case MainMenu:
		return m.updateMainMenu(msg)
	case TickerInput:
		return m.updateTickerInput(msg)
	case TimeSeriesSelection: // Add this case to handle the new state
		return m.updateTimeSeriesSelection(msg)
	case DisplayData:
		return m.updateDisplayData(msg)
	}
	return m, nil
}

func (m model) View() string {
	switch m.state {
	case MainMenu:
		return m.list.View()
	case TickerInput:
		return m.textInput.View() + "\nPress Enter to fetch data"
	case DisplayData:
		if m.fetchErrorMsg != "" {
			return "Error: " + m.fetchErrorMsg + "\nPress any key to return."
		}
		return m.renderPaginatedData() + "\nPress n for next, p for previous, any other key to return."
	}
	return ""
}

func (m *model) updateMainMenu(msg tea.Msg) (tea.Model, tea.Cmd) {
	var cmd tea.Cmd
	m.list, cmd = m.list.Update(msg)
	switch msg := msg.(type) {
	case tea.KeyMsg:
		switch msg.String() {
		case "q", "ctrl+c":
			return m, tea.Quit
		case "enter", " ":
			if m.list.SelectedItem().(item) == "Stock Market" {
				m.state = TickerInput
				m.textInput.Focus()
				return m, nil
			}
		}
	}
	return m, cmd
}

func (m *model) updateTickerInput(msg tea.Msg) (tea.Model, tea.Cmd) {
	var cmd tea.Cmd
	m.textInput, cmd = m.textInput.Update(msg)
	switch msg := msg.(type) {
	case tea.KeyMsg:
		if msg.Type == tea.KeyEnter && m.state == TickerInput {
			m.state = TimeSeriesSelection
			// You might want to reset or set up another textinput for this choice
			m.textInput.SetValue("")
			m.textInput.Placeholder = "Choose time-frame: Daily, Weekly, Monthly"
			m.textInput.Focus()
			return m, nil
		}
	case dataMsg:
		if msg.err != nil {
			m.fetchErrorMsg = msg.err.Error()
			m.state = DisplayData
		} else {
			m.data = msg.data
			m.state = DisplayData
		}
		return m, nil
	}
	return m, cmd
}

func (m *model) updateTimeSeriesSelection(msg tea.Msg) (tea.Model, tea.Cmd) {
	var cmd tea.Cmd
	m.textInput, cmd = m.textInput.Update(msg)
	switch msg := msg.(type) {
	case tea.KeyMsg:
		if msg.Type == tea.KeyEnter {
			m.timeSeriesChoice = m.textInput.Value()
			apiKey := config.LoadEnv()
			ticker := m.textInput.Value() // Assume you've stored the ticker somewhere
			// Dispatch command based on the choice
			return m, func() tea.Msg {
				switch m.timeSeriesChoice {
				case "daily":
					data, err := api.FetchTimeSeriesDaily(apiKey, ticker)
					return dataMsg{data: util.FormatTimeSeriesData(data), err: err}
				case "weekly":
					data, err := api.FetchTimeSeriesWeekly(apiKey, ticker)
					return dataMsg{data: util.FormatTimeSeriesData(data), err: err}
				case "monthly":
					data, err := api.FetchTimeSeriesMonthly(apiKey, ticker)
					return dataMsg{data: util.FormatTimeSeriesData(data), err: err}
				default:
					return dataMsg{err: fmt.Errorf("invalid choice")}
				}
			}
		}
	}
	return m, cmd
}

func (m *model) updateDisplayData(msg tea.Msg) (tea.Model, tea.Cmd) {
	switch msg := msg.(type) {
	case tea.KeyMsg:
		switch msg.String() {
		case "n":
			if (m.currentPage+1)*m.itemsPerPage < len(m.data) {
				m.currentPage++
			}
		case "p":
			if m.currentPage > 0 {
				m.currentPage--
			}
		default:
			m.state = MainMenu
			m.currentPage = 0 // Reset pagination
			m.fetchErrorMsg = ""
			m.data = nil
		}
	}
	return m, nil
}

func (m model) renderPaginatedData() string {
	start := m.currentPage * m.itemsPerPage
	end := start + m.itemsPerPage
	if end > len(m.data) {
		end = len(m.data)
	}

	var sb strings.Builder
	for i := start; i < end; i++ {
		sb.WriteString(m.data[i] + "\n")
	}
	return sb.String()
}

func Start() {
	m := newModel()

	p := tea.NewProgram(m)
	if _, err := p.Run(); err != nil {
		fmt.Println("Error running program:", err)
		os.Exit(1)
	}
}
