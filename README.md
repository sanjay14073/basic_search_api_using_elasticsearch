# ElasticSearch Book Management API

This project demonstrates a simple API for managing books using Elasticsearch and Gin in Go. It includes functionality to add books from a JSON file to Elasticsearch and retrieve book data by ID.

## Features

- **Add Data**: Reads book data from a JSON file (`data.json`) and indexes it into Elasticsearch.
- **Get Data**: Retrieves book information from Elasticsearch using the book's ID.

---

## Prerequisites

Before running the application, ensure you have the following installed:

- **Go**: Version 1.19 or higher
- **Elasticsearch**: Elasticsearch cloud account with Cloud ID and API key
- **Gin**: Web framework for Go
- **dotenv**: For managing environment variables

---

## Installation

1. **Clone the repository**:
   ```bash
   git clone https://github.com/sanjay14073/basic_search_api_using_elasticsearch.git
   cd basic_search_api_using_elasticsearch
   ```

2. **Install dependencies**:
   ```bash
   go mod tidy
   ```

3. **Create a `.env` file** in the project root and add your Elasticsearch credentials:
   ```plaintext
   CLOUD_ID=<your-elasticsearch-cloud-id>
   API_KEY=<your-elasticsearch-api-key>
   ```

---

## Usage

### Running the Application

1. **Start the server**:
   ```bash
   go run server.go
   ```

2. **Access API Endpoints**:
   - **Add Data**: 
     - Endpoint: `POST /add`
     - This endpoint reads `data.json` and adds book entries to Elasticsearch.
   - **Get Data**:
     - Endpoint: `GET /get`
     - Query Parameter: `id` (default: `1159142`)
     - Example:
       ```
       http://localhost:3000/get?id=1
       ```

---

## API Endpoints

### Add Data

- **URL**: `/add`
- **Method**: `POST`
- **Description**: Reads `data.json` and indexes book data into Elasticsearch.
- **Response**:
  ```json
  {
    "message": "Working fine"
  }
  ```

---

### Get Data

- **URL**: `/get`
- **Method**: `GET`
- **Query Parameter**: 
  - `id`: ID of the book (default: `1159142`)
- **Response**:
  ```json
  {
    "data": {
      "_index": "my_index",
      "_id": "1",
      "_source": {
        "id": "1",
        "title": "Book Title",
        "author": "Author Name",
        ...
      }
    }
  }
  ```
---

## License

This project is licensed under the [MIT License](LICENSE).


## Acknowledgments

- [Elastic Go Client](https://github.com/elastic/go-elasticsearch)
- [Gin Framework](https://github.com/gin-gonic/gin)