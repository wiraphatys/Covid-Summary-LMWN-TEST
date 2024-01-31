# Covid-Summary-LMWN-TEST

## Overview
This project is a simple JSON API built with Go and the Gin web framework. It provides a summary of COVID-19 cases by province and age group, fetching data from a public API.

## Requirements
- Go (version 1.18 or later)
- Gin Web Framework
- Testify (for testing)

## Installation
1. **Clone the Repository**
   ```bash
   git clone https://github.com/wiraphatys/Covid-Summary-LMWN-TEST.git
   cd Covid-Summary-LMWN-TEST
   ```

2. **Install Dependencies**
   Navigate to the project directory and install the necessary Go modules:
   ```bash
   go mod tidy
   ```

## Running the Project
1. **Start the Server**
   In the project directory, start the server by running:
   ```bash
   go run main.go
   ```

2. **Accessing the API**
   With the server running, you can access the API at:
   `http://localhost:8080/covid/summary`

## Testing
1. **Running Tests**
   You can run the tests for this project using:
   ```bash
   go test
   ```

## API Endpoint
- **GET /covid/summary**
  - **Description**: Fetches a summary of COVID-19 cases categorized by province and age group.
  - **Response Format**:
    ```json
    {
      "Province": {
        "ProvinceName1": NumberOfCases,
        "ProvinceName2": NumberOfCases,
        ...
      },
      "AgeGroup": {
        "0-30": Count,
        "31-60": Count,
        "61+": Count,
        "N/A": Count
      }
    }
    ```

## Contributions
Contributions to this project are welcome. Please fork the repository and submit a pull request with your changes.
