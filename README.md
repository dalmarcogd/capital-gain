# capital-gain

This app calculates the tax on stock transactions.
If the transaction exceeds 20,000.00 then a tax of 20% will be applied to the profit.

## how to run using docker

```shell
docker build -t capgain:latest . && docker run -it capgain
```

## how to run using go
```shell
 go build -ldflags="-s -w" -o ./capgain ./cmd/capgain && ./capgain calculator
```

## test

After run the calculator paste on shell the above example:
```json
[{"operation":"buy","unit-cost":10.00,"quantity":10000},{"operation":"sell","unit-cost":2.00,"quantity":5000},{"operation":"sell","unit-cost":20.00,"quantity":2000},{"operation":"sell","unit-cost":20.00,"quantity":2000},{"operation":"sell","unit-cost":25.00,"quantity":1000},{"operation":"buy","unit-cost":20.00,"quantity":10000},{"operation":"sell","unit-cost":15.00,"quantity":5000},{"operation":"sell","unit-cost":30.00,"quantity":4350},{"operation":"sell","unit-cost":30.00,"quantity":650}]
```