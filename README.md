"# Learning-Go" 

This is a simple API built with Go, demonstrating the use of middleware, routing, and a mock database.

## Project Structure

```
.
├── cmd
│   └── api
│       └── main.go
├── internal
│   ├── handlers
│   │   ├── api.go
│   │   ├── get_coin_balance.go
│   │   └── ...
│   ├── middleware
│   │   └── authorization.go
│   └── tools
│       └── mockdb.go
└── go.mod
```
## Features
- **Middleware**: Custom middleware for authorization.
- **Routing**: Handlers for various API endpoints.
- **Mock Database**: Simulated database interactions for testing.

## Getting Started
1. Clone the repository:
   ```
   git clone https://github.com/sifunadonnex/Learning-Go.git
   ```
2. Navigate to the project directory:
   ```
   cd Learning-Go
   ```
3. Install the dependencies:
   ```
   go mod tidy
   ```
4. Start the API server:
   ```
   go run cmd/api/main.go
   ```
5. Test the API endpoints using a tool like Postman or curl.

## API Endpoints
- `GET /api/coins`: Retrieve all coins.
- `GET /api/coins/{id}`: Retrieve a specific coin by ID.
- `POST /api/coins`: Create a new coin.
- `PUT /api/coins/{id}`: Update a coin by ID.
- `DELETE /api/coins/{id}`: Delete a coin by ID.
- `GET /api/balance`: Retrieve the balance of a specific coin.
- `POST /api/balance`: Update the balance of a specific coin.

## Contributing
Contributions are welcome! Please fork the repository and submit a pull request with your changes.
## License
This project is licensed under the MIT License. See the [LICENSE](LICENSE) file for details.
## Contact
For any questions or feedback, feel free to reach out via email at [sifunadonnex@gmail.com](mailto:sifunadonnex@gmail.com).
## Acknowledgments
- Thanks to the Go community for their resources and support.
