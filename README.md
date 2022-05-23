# Order Synchronizer

Synchronize decentralized orders for NitfyConnect Protocol

## Build

```shell
git clone https://github.com/NiftyConnect/Order-Synchronizer.git
cd Order-Synchronizer
make build
```

## Run

```shell
cp config-example.json build/config.json
cd build
 ./niftyconnect-order-synchronizer start --config-path config.json
```

## Query Orders

Query orders with these APIs:
```json
[
  "/api/v1/query_order/{blockchain}/{order_hash}",
  "/api/v1/query_orders/{blockchain}/{nft_address}",
  "/api/v1/query_orders/{blockchain}/{nft_address}/{token_id}"
]
```