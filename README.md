# Dev Quotes API

A public API for developer-related quotes. This API provides endpoints to retrieve quotes from famous developers, programmers, and tech leaders.

## Features

- Get random developer quotes
- Filter quotes by author or tag
- Get specific quotes by ID
- Rate limiting
- CORS enabled
- Swagger documentation
- Error handling
- Configuration management

## Getting Started

### Prerequisites

- Go 1.23 or higher
- Make (optional, for using Makefile commands)

### Installation

1. Clone the repository:

```bash
git clone https://github.com/shawnkost/dev-quotes-api.git
cd dev-quotes-api
```

2. Install dependencies:

```bash
go mod download
```

3. Copy the example environment file:

```bash
cp .env.example .env
```

4. Build the application:

```bash
make build
```

### Running the Application

Development mode:

```bash
make run
```

Production mode:

```bash
ENVIRONMENT=production ./bin/dev-quotes-api
```

### API Documentation

Once the server is running, you can access the Swagger documentation at:

```
http://localhost:8080/swagger/index.html
```

## API Endpoints

### Get Random Quote

```
GET /v1/quotes/random
```

### Get Quote by ID

```
GET /v1/quotes/:id
```

### Get Filtered Quotes

```
GET /v1/quotes?author=<author>&tag=<tag>&page=<page>&per_page=<per_page>
```

## API Usage Guidelines

### Rate Limiting

The API implements rate limiting to ensure fair usage:

- Default rate limit: 50 requests per minute per IP address

### Pagination

The filtered quotes endpoint supports pagination:

- `page`: Page number (default: 1)
- `per_page`: Items per page (default: 10, max: 100)
- Response includes pagination metadata:
  - `total`: Total number of quotes
  - `page`: Current page number
  - `per_page`: Items per page
  - `total_pages`: Total number of pages
  - `has_next`: Whether there is a next page
  - `has_previous`: Whether there is a previous page

### Example Responses

#### Successful Response (Paginated Quotes)

```json
{
  "quotes": [
    {
      "id": "1",
      "author": "Martin Fowler",
      "text": "Any fool can write code that a computer can understand. Good programmers write code that humans can understand.",
      "tags": ["programming", "clean-code"]
    }
  ],
  "total": 42,
  "page": 1,
  "per_page": 10,
  "total_pages": 5,
  "has_next": true,
  "has_previous": false
}
```

#### Error Response

```json
{
  "type": "VALIDATION",
  "message": "rate limit exceeded. please try again later",
  "code": 429,
  "details": {
    "retry_after": 60
  }
}
```

## Error Handling

The API uses standardized error responses:

| Error Type | HTTP Code | Description                |
| ---------- | --------- | -------------------------- |
| NOT_FOUND  | 404       | Resource not found         |
| VALIDATION | 400       | Invalid request parameters |
| INTERNAL   | 500       | Server error               |
| RATE_LIMIT | 429       | Rate limit exceeded        |

All error responses include:

- `type`: Error type identifier
- `message`: Human-readable error message
- `code`: HTTP status code
- `details`: Additional error information (when applicable)

## Configuration

The application can be configured using environment variables:

- `PORT`: Server port (default: 8080)
- `ENVIRONMENT`: Environment (development/production)
- `READ_TIMEOUT`: Server read timeout
- `WRITE_TIMEOUT`: Server write timeout
- `RATE_LIMIT`: API rate limit
- `RATE_LIMIT_TIME`: Rate limit time window

## Development

### Available Make Commands

- `make swag`: Generate Swagger documentation
- `make run`: Run the application
- `make build`: Build the binary
- `make tidy`: Tidy up go.mod/go.sum
- `make fmt`: Format code

### Project Structure

```
.
├── cmd/
│   └── server/        # Application entry point
├── internal/
│   ├── api/           # API handlers
│   ├── config/        # Configuration
│   ├── errors/        # Error handling
│   ├── repository/    # Data access
│   └── service/       # Business logic
├── configs/           # Configuration files
└── docs/              # Generated documentation
```

## Contributing

We welcome contributions! Please feel free to submit a Pull Request. For major changes, please open an issue first to discuss what you would like to change.

1. Fork the repository
2. Create your feature branch (`git checkout -b feature/amazing-feature`)
3. Commit your changes (`git commit -m 'Add some amazing feature'`)
4. Push to the branch (`git push origin feature/amazing-feature`)
5. Open a Pull Request

Please make sure to update tests as appropriate and follow the existing code style.

## License

This project is licensed under the MIT License - see the [LICENSE](LICENSE) file for details.

## Contact

Shawn Kost - [@shawnkost](https://github.com/shawnkost)

Project Link: [https://github.com/shawnkost/dev-quotes-api](https://github.com/shawnkost/dev-quotes-api)
