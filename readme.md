# Cryptocurrency rest api

Cryptocurrency Price REST API written in Golang. Data is scraped from CoinMarketCap.

*Please note that this is an unofficial API and is __not__ supported or controlled by CoinMarketCap itself.*

Live api base url: [https://go-crypto.herokuapp.com]('https://go-crypto.herokuapp.com')

## Usage
#### `GET /coins.json`

**Output:** JSON
Response:
```json
  [
     {
      "success": true,
      "timestamp": 1515959618,
      "amount_of_coins": 1433,
      "coins": [
          {
          "name": "bitcoin",
          "ticker": "BTC",
          "btc": "1.0",
          "price": "13615.6",
          "currency": "usd"
          },
          {
          "name": "ethereum",
          "ticker": "ETH",
          "btc": "0.0978496",
          "price": "1332.52",
          "currency": "usd"
          }
      ]
    }
  ]
  ...
```

## Run Locally
I use dep to manage my dependencies so this guide assumes you have that installed :)
```sh
$ dep ensure
$ go run server.go
```

## License
[WTFPL License](LICENSE)