# Mango Markets - Financial Application Overview

Mango Markets is a comprehensive financial application aimed at providing users with a wide range of financial data and insights directly from their terminal. Leveraging the Alpha Vantage API, the application offers up-to-date information on various markets including stocks, forex, cryptocurrency, and financial news.

## Features

### 1. Stock Market

- **Top Gainers and Losers**: Retrieves the day's top gainers and losers in the stock market, providing a snapshot of market movements.
- **Ticker Search**: Allows users to input a stock ticker and choose from a range of data to display, including stock price, company overview, financial statements, and more.
- **Global Market Status**: Shows the open and close times of major global stock exchanges, along with their current status (open or closed).

### 2. Financial News

- Users can search for news related to a specific ticker or a general financial topic, offering insights and updates on market movements and corporate developments.

### 3. Forex Market

- This feature will allow users to view forex exchange rates and movements. (To be implemented)

### 4. Cryptocurrency Market

- Offers cryptocurrency market data including prices, market cap, and volume for popular cryptocurrencies. (To be implemented)

## Planned Enhancements

- **Forex and Cryptocurrency Data**: Expanding the application to include detailed forex and cryptocurrency market data.
- **Interactive Charts**: Implementing interactive charts for visual representation of financial data.
- **Technical Indicators**: Technical Analysis to make more informed dta driven solutions
- **Portfolio Tracking**: Allowing users to track their investment portfolio performance directly within the application.
- **Alerts and Notifications**: Setting up alerts for price movements, earnings announcements, or other significant events related to the user's interests.
- **Charmbracelet**: Making my app more appealing with bubbletea, bubbles, lipgloss, etc.

## Technical Overview

Mango Markets is built with Go, utilizing the standard library for API requests, data processing, and terminal interaction.

The application operates in a loop, prompting the user to navigate through the main menu to access different functionalities. Data fetched from the Alpha Vantage API is processed and displayed in a user-friendly format, offering valuable financial insights directly in the terminal.

## Getting Started

To run Mango Markets, ensure you have `Go`, `Docker`, and `Make` installed and set up on your machine. Clone the repository, navigate to the application directory, and run:

```bash
make all
```

Alternatively there is an exe located in bin
