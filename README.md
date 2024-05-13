# Finance Tracker

Finance Tracker is a command-line utility that allows users to track their income and expenses directly through a Google Sheet. The application features secure user authentication, the ability to categorize and log financial entries, as well as visualization of spending habits and income.

## Getting Started

These instructions will guide you on how to set up and run the Finance Tracker on your local machine.

### Prerequisites

Before you install the Finance Tracker, ensure you have the following:
- Go (at least version 1.15) installed on your machine. [Download Go](https://golang.org/dl/)
- `credentials.json` for OAuth2 authentication obtained from the gradescope.

### Installation

1. **Clone the repository**

   Clone this repository to your local machine using:

   ```bash
   git clone https://github.com/andrewphm/acs-4210-budget-tracker
   cd finance-tracker
   ```

2. **Set up Google Cloud credentials**

   NOTE: To allow easier testing, I've included the `credentials.json` and the `token.json`. These files will be used to authenticate with Google Sheets API while the project is being graded and will be removed shortly after. This is a temporary solution.

3. **Install dependencies**

   Run the following command in the root directory to install the required Go packages:

   ```bash
   go mod tidy
   ```

### Running the Program

To start the program, execute:

```bash
go run .
```

This command compiles and runs the application. Follow the on-screen prompts in the command line interface to interact with the finance tracker.

### Usage

1. **Authenticate**: On initial start, the application will request access to your Google Sheets via a browser. Log in and authenticate to proceed.

2. **Input Transactions**: Use the command-line menus to select whether to input an expense or income, fill out the necessary details, and categorize each transaction as directed.

3. **View Summary**: Access summarized views of your financial data directly through the CLI. Navigate using the provided keys to explore different categories and time frames.


## License

This project is licensed under the MIT License - see the LICENSE file for details.

## Acknowledgments

- Thanks to the developers of the dependencies used in this project, especially those who maintain `tview` and Google Sheets API packages.
