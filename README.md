# txd

An order matching engine written in go.

* JSON RPC interface.
* In-memory balances.
* Multi persistent adaptor(plain text, sql database, leveldb).
* Writen in go.

## Api

### SendTo

```JSON
  {
    "From": 100,
    "To": 101,
    "Amount": {
      "Currency": "BTC",
      "Value": "0.1"
    }
  }
```

### CreateOrder

### CancelOrder
