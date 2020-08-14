# Top Coin

Simple Application to fetch the top performing coins almost realtime.

## Stack

-   Go Lang and Mux router for HTTP communication
-   Rabbit MQ for inter service communication

## Architecture

This applciation has 3 microservices

-   Pricing service - This service gets all the available coins and the corresponding prices
-   Rating service - This service sorts the coins according to the price
-   Client service - This service handles all client interactions and returns the top coins based on the parameters.

### Pricing Service

Pricing service fetches the latest price of all coins almost real time(5 seconds buffer).[CoinmarketCap API](https://coinmarketcap.com/api/) is used to get the current USD prices and [Cryptocompare API](https://www.cryptocompare.com/api) to get all the active coins. For every 5 second, all the active stocks and prices are fetched.

Some of the assets in `cryptocompare` are not present in `coinmarketcap`. We ignore those coins and fetch the prices of other coins.

Price details are passed to Rating Service using `RabbitMQ`.

### Rating Service

Rating Service receives the price details and sorts the price from high to low based on the currency. The data is stored in memory but we can make it persistent using a Database. If Pricing service goes down and there is already data available, it sends the cached data.

HTTP service is exposed on port `8000`

To fetch the prices

```
GET localhost:8000/rank?limit=<limit>
```

### Client Service

Client Service handles all the client interactions. HTTP service is exposed on port `8001`. Client hits the following URL

```
GET localhost:8001/topcoin?limit=<limit>
```

Client service inturn triggers the Rating service and fetches the latest information.

#### Errors

-   Invalid Limit - Limit variable is not sent or limit value is 0
-   Server is down. Please check later - Either Price Service or Rating Service is down
-   Internal Server Error - Unexpected Error

## Steps to Install

-   Clone the directory

```
https://github.com/VaishakVk/topcoin.git
```

-   <ins>Pricing service</ins>

*   Navigate to `clientService`
*   Create `.env` file. Reference keys are avaialble in `.env.example`
*   From the terminal, run

```
go run main.go
```

-   <ins>Rating service</ins>

*   Navigate to `ratingService`
*   Create `.env` file. Reference keys are avaialble in `.env.example`
*   From the terminal, run

```
go run main.go
```

-   <ins>Client service</ins>

*   Navigate to `clientService`
*   Create `.env` file. Reference keys are avaialble in `.env.example`
*   From the terminal, run

```
go run main.go
```

When all the services are running, hit the following URL from browser or Postman

```
GET localhost:8001/topcoin?limit=<limit:int>
```
