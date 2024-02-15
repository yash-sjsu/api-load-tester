# API Load tester

This Go project, named `my-go-project`, demonstrates the capabilities of Go for making concurrent API calls, showcasing concurrency, HTTP request handling, and file operations. It focuses on sending multiple API requests in parallel, handling their responses, and saving the response data into files, making it particularly useful for scenarios like load testing or data fetching.

## Features

- **Concurrent API Calls:** Utilizes Go's goroutines to make multiple API calls concurrently.
- **API Authentication:** Includes sending an API key in the request header for authenticated API calls.
- **Dynamic Query Parameters:** Demonstrates how to add query parameters to API requests.
- **File Operations:** Saves each API response to a uniquely named file within a specified directory for analysis.

## Getting Started

### Prerequisites

Ensure you have Go installed on your system. This project requires Go 1.16 or later for module support.

### Installation

1. **Clone the Project**

   Clone the project to your local machine with the following command:

   ```bash
   git clone git@github.com:yash-sjsu/api-load-tester.git
   cd api-load-tester
   ```

2. **Initialize Go Module** (If not already done)

   ```bash
   go mod init my-go-project
   ```

3. **Install Dependencies**

   Running the project will automatically download the necessary dependencies, such as `github.com/google/uuid`.

### Configuration

- Update the `apiURL` and `apiKey` constants in `main.go` to your target API endpoint and authentication key.
- Adjust `queryParams`, `number_of_iterations`, and `number_of_api_calls` in the `main` function as needed.

### Running the Project

Execute the project with:

```bash
go run main.go
```

Or build an executable with:

```bash
go build
./my-go-project
```

### Output

The program performs concurrent API calls as per the configured parameters. Each response is saved in a uniquely named file within the `load-testing-results` directory.

## Customization

You can customize the project by adjusting the API URL, API key, query parameters, number of iterations, and number of API calls per iteration according to your requirements.
