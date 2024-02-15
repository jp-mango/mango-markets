package charm

import (
	"fmt"
	"log"
	"sort"
	"strings"

	tea "github.com/charmbracelet/bubbletea"
	"github.com/jp-mango/mangomarkets/internal/api"
	"github.com/jp-mango/mangomarkets/internal/config"
)

// Model represents the state of the TUI application.
type model struct {
	apiKey           string
	cursor           int
	fetching         bool
	errorMessage     string
	inMenu           bool
	inTickerInput    bool
	inputTicker      string
	choices          []string
	timeSeries       interface{}
	selectedInterval string
	displayingData   bool
}

// initialModel initializes the starting state of the TUI application.
func initialModel() model {
	return model{
		apiKey:  config.LoadEnv(),
		choices: []string{"Fetch Daily Time Series", "Fetch Weekly Time Series", "Fetch Monthly Time Series", "Quit"},
		inMenu:  true,
	}
}

// Init is called when the application starts.
func (m model) Init() tea.Cmd {
	return nil
}

// Update processes messages and updates the application state.
func (m model) Update(msg tea.Msg) (tea.Model, tea.Cmd) {
	switch msg := msg.(type) {
	case tea.KeyMsg:
		if m.inTickerInput {
			switch msg.Type {
			case tea.KeyRunes:
				m.inputTicker += string(msg.Runes)
			case tea.KeyBackspace:
				if len(m.inputTicker) > 0 {
					m.inputTicker = m.inputTicker[:len(m.inputTicker)-1]
				}
			case tea.KeyEnter:
				if m.inTickerInput {
					// Ensure we have valid input before proceeding
					if m.inputTicker == "" {
						// Optionally handle empty ticker input
						m.errorMessage = "Ticker symbol cannot be empty. Please enter a valid symbol."
						return m, nil
					}
					m.fetching = true
					m.inTickerInput = false
					parts := strings.Split(m.selectedInterval, " ")
					if len(parts) > 1 {
						// Correctly assuming "Fetch [Interval] Time Series" structure
						interval := strings.ToLower(parts[1]) // Now safely accessing the second element
						return m, fetchTimeSeries(m.apiKey, m.inputTicker, interval)
					} else {
						// Log error or handle unexpected selectedInterval format
						log.Println("Unexpected selectedInterval format:", m.selectedInterval)
						m.errorMessage = "An unexpected error occurred. Please try again."
						return m, nil
					}
				}
			case tea.KeyEsc:
				m.inTickerInput = false
				m.inMenu = true
			}
		} else if m.inMenu {
			switch msg.String() {
			case "up", "k":
				if m.cursor > 0 {
					m.cursor--
				}
			case "down", "j":
				if m.cursor < len(m.choices)-1 {
					m.cursor++
				}
			case "enter":
				selectedChoice := m.choices[m.cursor]
				if selectedChoice == "Quit" {
					return m, tea.Quit
				}
				m.selectedInterval = strings.ToLower(strings.Split(selectedChoice, " ")[1]) // Extracts "daily", "weekly", "monthly"
				m.inTickerInput = true
				m.inMenu = false
			}
		} else if m.displayingData {
			if msg.String() == "enter" || msg.String() == "q" {
				m.displayingData = false
				m.inMenu = true
				m.timeSeries = nil
			}
		}
	case timeSeriesDataMsg:
		m.fetching = false
		m.timeSeries = msg.data
		m.displayingData = true
		return m, nil
	case errMsg:
		m.fetching = false
		m.errorMessage = msg.err.Error()
		m.displayingData = true
		return m, nil
	}

	return m, nil
}

// View renders the application UI based on the current state.
func (m model) View() string {
	var b strings.Builder

	// Handling the fetching state
	if m.fetching {
		b.WriteString("Fetching data...\n")
	} else if m.displayingData {
		// Check if there is an error message to display
		if m.errorMessage != "" {
			b.WriteString(fmt.Sprintf("Error: %s\n", m.errorMessage))
		} else if m.timeSeries != nil {
			// Ensure the timeSeries data is of type TimeSeriesProvider
			if tsProvider, ok := m.timeSeries.(api.TimeSeriesProvider); ok {
				b.WriteString("\nFetched Data:\n")

				tsData := tsProvider.GetTimeSeriesData()
				var dates []string
				for date := range tsData {
					dates = append(dates, date)
				}
				// Example: Sorting dates in descending order; adjust as needed
				sort.Slice(dates, func(i, j int) bool {
					return dates[i] > dates[j]
				})

				for _, date := range dates {
					b.WriteString(fmt.Sprintf("Date: %s\n", date))
					metrics := tsData[date]
					// Assuming metrics keys have a consistent order you want to display; adjust as needed
					for key, value := range metrics {
						b.WriteString(fmt.Sprintf("%s: %s\n", key, value))
					}
					b.WriteString("----------------\n")
				}
			} else {
				b.WriteString("Unable to display data. Data format is unexpected.\n")
			}
		}

		b.WriteString("\nPress enter to return to menu.\n")
	} else if m.inTickerInput {
		b.WriteString("Enter Ticker Symbol: " + m.inputTicker + "\n")
	} else if m.inMenu {
		for i, choice := range m.choices {
			cursor := " "
			if i == m.cursor {
				cursor = ">"
			}
			b.WriteString(fmt.Sprintf("%s %s\n", cursor, choice))
		}
		b.WriteString("\nUse arrow keys to navigate, enter to select, and q to quit.\n")
	}

	return b.String()
}

// fetchTimeSeries constructs a command to fetch time series data based on the given interval.
func fetchTimeSeries(apiKey, ticker, interval string) tea.Cmd {
	return func() tea.Msg {
		var data api.TimeSeriesProvider
		var err error

		switch interval {
		case "daily":
			data, err = api.FetchTimeSeriesDaily(apiKey, ticker)
		case "weekly":
			data, err = api.FetchTimeSeriesWeekly(apiKey, ticker)
		case "monthly":
			data, err = api.FetchTimeSeriesMonthly(apiKey, ticker)
		default:
			err = fmt.Errorf("unknown interval type: %s", interval)
		}

		if err != nil {
			log.Printf("Error fetching time series: %v", err) // Log error
			return errMsg{err}
		}

		log.Println("Fetched time series data successfully.") // Log success
		return timeSeriesDataMsg{data}                        // Ensure this matches your struct definition
	}
}

// timeSeriesDataMsg is used to pass fetched time series data back to the Update function.
type timeSeriesDataMsg struct {
	data api.TimeSeriesProvider // Change to your specific time series data type as needed
}

// errMsg is used to pass error messages back to the Update function.
type errMsg struct {
	err error
}

// Start initializes and runs the TUI application.
func Start() {
	p := tea.NewProgram(initialModel())
	if _, err := p.Run(); err != nil { // Use Run() instead of Start()
		log.Fatalf("Error running TUI: %v", err)
	}
}
