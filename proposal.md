## ACS-4210 - Utility Project - Finance Tracker

### Goal
The goal of this project is to create a finance tracker that allows users to track their expenses and income. The tracker will enable users to input their financial transactions and provide a summary of their financial status. Additionally, it will assist users in setting and tracking their financial goals.

### Features
- **User Authentication:** Users can authenticate using OAuth 2.0 with Google Sheets API to secure access to their spreadsheets.
- **Expense Tracking:** Users can input expenses, categorize them, and view historical expense data.
- **Income Tracking:** Users can record their income, categorize sources, and review income history.
- **Financial Summary and Goals:** Users can view a dashboard summarizing their financial status and track progress towards predefined financial goals.
- **Interactive CLI:** Utilize an interactive command line interface for a user-friendly experience in managing finances.

### Technologies
- **Go Language:** The backend of the application is written in Go, which offers excellent support for concurrent operations and integrates seamlessly with external APIs.
- **Google Sheets API:** To store and retrieve financial data, leveraging Google's established infrastructure for real-time data management and accessibility.
- **OAuth2 Package:** For secure authentication with Google APIs, ensuring that users’ data is protected.
- **Tview Package:** A Go package used to build rich, interactive command-line interfaces, enhancing user experience by providing visually appealing and navigable menus and displays.


### System Architecture
- **Command-Line Interface (CLI):** A robust CLI built with `tview` that interacts with the user to perform operations such as entering transactions, viewing reports, and setting goals.
- **Data Handling:** Utilizes the `sheets/v4` API to handle all CRUD operations on the Google Sheets document that serves as the database.
- **Authentication Service:** Implemented using `oauth2` and `google` packages to handle authentication and secure sessions between the user and their Google Sheets.

