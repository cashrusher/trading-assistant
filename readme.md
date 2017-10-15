# Golang version trading assistant for Bitfinex and Kraken.

Support auto trading(buy and sell) and trading history display.

# Front-end And Backed-end APIs

## Get kraken Supported currency

GET

    http://localhost:8080/assistant/currencies/kraken

Response:

```json
{
    "currencies":["XBT","BTH","BCH","ETC","LTC","ICN","GNO","MLN","REP"]
}

```

## Get bitfinex Supported currency

GET

    http://localhost:8080/assistant/currencies/bitfinex

Response:

```json
{
    "currencies":["XBT","BTH","BCH","ETC","LTC","ICN","GNO","MLN","REP"]
}

```

## Buy

POST

    http://localhost:8080/assistant/buy

Request:

```json

{
    "platform":"kraken",  //can be 'kraken' or 'bitfinex'
    "currency":"BTH", // it's the currency returned by the currencies api
    "amount":34.9982,
    "price":887.000293
}
```

Response:

success

```json
{
  "status":"success",
  "info":{
    "tradeID":"123455djsdkh33"
  }
}
```

Failed

```json
{
  "status":"failed",
  "message":"Net work error"
}
```

## Sell

POST 

    http://localhost:8080/assistant/sell

Request:

```json

{
    "platform":"kraken", //can be 'kraken' or 'bitfinex'
    "currency":"BTH", // it's the currency returned by the currencies api
    "amount":34.9982,
    "price":887.000293
}

```

Response:

success

```json
{
  "status":"success",
  "info":{
    "tradeID":"123455djsdkh33"
  }
}
```

Failed

```json
{
  "status":"failed",
  "message":"Net work error"
}
```
      
## Get All history

History will display in desc order by time, auto refresh by 5s

GET

    http://localhost:8080/assistant/trading-history
    
Response



```json
[
  {
    "time":"",
    "orderID":"",
    "platform":"",
    "orderType":"",
    "pair":"",
    "price":"",
    "volume":"",
    "amount":"",
    "fee":"",
    "status":""
  },
  {
   "time":"",
   "orderID":"",
   "platform":"",
   "orderType":"",
   "pair":"",
   "price":"",
   "volume":"",
   "amount":"",
   "fee":"",
   "status":""
  }
]
```
      