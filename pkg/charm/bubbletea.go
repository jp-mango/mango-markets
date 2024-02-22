package charm

import (
	"fmt"
	"os"
	"time"

	"github.com/charmbracelet/bubbles/list"
	"github.com/charmbracelet/bubbles/spinner"
	"github.com/charmbracelet/bubbles/textinput"
	tea "github.com/charmbracelet/bubbletea"
)

// Define application states
type state int

const (
	mainMenuState state = iota
	tickerInputState
	timeIntervalState
	fetchingDataState
	displayDataState
)

// Define menu items
type menuItem string

func (i menuItem) Title() string       { return string(i) }
func (i menuItem) Description() string { return "" }
func (i menuItem) FilterValue() string { return string(i) }

// Model represents the application state
type model struct {
	currentState  state
	list          list.Model
	tickerInput   textinput.Model
	spinner       spinner.Model
	fetchedData   string
	selectedStock string
	selectedTime  string
	apiKey        string
}

type dataMsg struct {
	data []string
}

func initialModel() model {
	// Initialize the main menu
	items := []list.Item{
		menuItem("Stock Market"),
		menuItem("Financial News"),
		menuItem("Forex Market"),
		menuItem("Cryptocurrency Market"),
		menuItem("Exit"),
	}
	l := list.New(items, list.NewDefaultDelegate(), 0, 0)
	l.Title = "Mango Markets 🥭"

	// Initialize the ticker input
	ti := textinput.New()
	ti.Placeholder = "Enter Ticker Symbol"
	ti.Focus()

	// Initialize the spinner
	s := spinner.New()
	s.Spinner = spinner.Dot

	return model{
		currentState: mainMenuState,
		list:         l,
		tickerInput:  ti,
		spinner:      s,
		apiKey:       "your_api_key_here", // Use your actual API key
	}
}

func (m model) Init() tea.Cmd {
	return nil
}

func (m model) Update(msg tea.Msg) (tea.Model, tea.Cmd) {
	switch m.currentState {
	case mainMenuState:
		return m.updateMainMenu(msg)
	case tickerInputState:
		return m.updateTickerInput(msg)
	case timeIntervalState:
		return m.updateTimeInterval(msg)
	case fetchingDataState:
		return m.updateFetchingData(msg)
	case displayDataState:
		return m.updateDisplayData(msg)
	}
	return m, nil
}

// Update function for the main menu state
func (m model) updateMainMenu(msg tea.Msg) (tea.Model, tea.Cmd) {
	switch msg := msg.(type) {
	case tea.KeyMsg:
		switch msg.String() {
		case "q", "ctrl+c", "esc":
			return m, tea.Quit
		case "enter":
			i, ok := m.list.SelectedItem().(menuItem)
			if !ok {
				return m, nil
			}
			switch i {
			case "Stock Market":
				m.currentState = tickerInputState
				m.tickerInput.Focus()
				return m, nil
			case "Exit":
				return m, tea.Quit
				// Handle other menu items similarly...
			}
		}
	}

	var cmd tea.Cmd
	m.list, cmd = m.list.Update(msg)
	return m, cmd
}

// Update function for the ticker input state
func (m model) updateTickerInput(msg tea.Msg) (tea.Model, tea.Cmd) {
	var cmd tea.Cmd
	m.tickerInput, cmd = m.tickerInput.Update(msg)

	switch msg := msg.(type) {
	case tea.KeyMsg:
		if msg.Type == tea.KeyEnter {
			m.selectedStock = m.tickerInput.Value()
			m.currentState = timeIntervalState
			// Reset ticker input for potential later use
			m.tickerInput.Reset()
			m.tickerInput.Blur()
			return m, nil
		}
	}
	return m, cmd
}

// Update function for the time interval selection state
func (m model) updateTimeInterval(msg tea.Msg) (tea.Model, tea.Cmd) {
	// This is a placeholder for handling time interval selection.
	// For simplicity, we're directly moving to the fetching data state.
	// You might want to implement a similar text input or list selection for the interval.
	m.currentState = fetchingDataState
	return m, m.fetchData()
}

// Example fetchData function that initiates data fetching and returns a command
func (m model) fetchData() tea.Cmd {
	return func() tea.Msg {
		// Simulate a fetch with a delay
		time.Sleep(2 * time.Second)
		// Simulate fetched data
		fetchedData := "Fetched data for " + m.selectedStock
		return dataMsg{data: []string{fetchedData}}
	}
}

// Update function for the fetching data state
func (m model) updateFetchingData(msg tea.Msg) (tea.Model, tea.Cmd) {
	switch msg := msg.(type) {
	case dataMsg:
		m.fetchedData = fmt.Sprintf("%v", msg.data)
		m.currentState = displayDataState
		return m, nil
	}

	var cmd tea.Cmd
	m.spinner, cmd = m.spinner.Update(msg)
	return m, cmd
}

// Update function for the display data state
func (m model) updateDisplayData(msg tea.Msg) (tea.Model, tea.Cmd) {
	switch msg := msg.(type) {
	case tea.KeyMsg:
		if msg.String() == "q" || msg.String() == "esc" {
			m.currentState = mainMenuState
			m.fetchedData = ""
			return m, nil
		}
	}
	return m, nil
}

func (m model) View() string {
	switch m.currentState {
	case mainMenuState:
		return m.list.View()
	case tickerInputState, timeIntervalState:
		return m.tickerInput.View()
	case fetchingDataState:
		return m.spinner.View() + "\nFetching data..."
	case displayDataState:
		return m.fetchedData
	}
	return "Loading..."
}

func Start() {
	p := tea.NewProgram(initialModel(), tea.WithAltScreen())
	if _, err := p.Run(); err != nil {
		fmt.Printf("Error running program: %v", err)
		os.Exit(1)
	}
}

// Implement other methods as needed, such as fetching data from the API,
// handling user input for ticker symbols and time intervals, and formatting
// the display of fetched data.
