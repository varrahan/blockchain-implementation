# Blockchain Emulator

This project is a blockchain emulator that provides a set of HTTP endpoints for interacting with a simulated blockchain.

## Routes

The router for this application is created in the `CreateRouter` function, which sets up the following components:

1. A new blockchain instance
2. A handler context containing the blockchain
3. Various routes for interacting with the blockchain

Here's an overview of the routes:

### Available Routes

1. **POST /mineblock/{address}** - Mine a new block and assign the reward to the specified address
2. **POST /maketransaction** - Submit a new transaction to the blockchain
3. **GET /getbalance/{address}** - Retrieve the balance of a specific address
4. **GET /getblocks** - Retrieve all blocks in the blockchain

## Setup and Running

### Setup

1. Ensure you have Go installed on your system.
2. Clone this repository.
3. Navigate to the project directory.
4. Run `go mod tidy` to ensure all dependencies are installed.
5. Build the project with `go build`.
6. (Optional) To run as a docker container you have docker desktop or install docker engine inside WSL2 and expose it for Windows access

### Local running

To run service locally:
```bash
go run main.go
```

To run service on docker container:
```bash
docker build -t [image_name] .

docker compose up
```

## Dependencies

This project uses the following external dependency:

- `github.com/gorilla/mux` for HTTP request routing
- `github.com/joho/godotenv` to grab variables from dotenv file

## API Usage

Here's a brief overview of how to use the API:

1. **Mine a Block**
   ```
   POST /mineblock/{address}
   ```
   Replace `{address}` with the address that should receive the mining reward.

2. **Make a Transaction**
   ```
   POST /maketransaction
   ```
   Send a POST request with the transaction details in the request body. The request body should be a JSON object with the following structure:
   ```yaml
   {
     "From": "string",
     "To": "string",
     "Amount": float,
   }
   ```
   - `From`: The address of the sender
   - `To`: The address of the recipient
   - `Amount`: The amount to transfer (as a floating-point number)

3. **Get Balance**
   ```
   GET /getbalance/{address}
   ```
   Replace `{address}` with the address whose balance you want to check.

4. **Get All Blocks**
   ```
   GET /getblocks
   ```
   This will return all blocks in the blockchain.
