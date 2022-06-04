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
