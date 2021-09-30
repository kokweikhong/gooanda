
## TODO

#### OANDA endpoints

- Account
    - [x] [GET] Accounts
    - [x] [GET] AccountsByID
    - [x] [GET] AccountSummary
    - [x] [GET] AccountInstruments
    - [ ] [PATCH] AccountConfiguration
    - [ ] [GET] AccountChanges
- Instrument
    - [x] [GET] InstrumentCandles
    - [x] [GET] InstrumentOrderBook
    - [x] [GET] InstrumentPositionBook
- Order
    - [ ]Need to create data struct, now only json response.
    - [ ] [POST] OrderCreate
        - [X] MarketOrderRequest
        - [X] LimitOrderRequest
        - [X] StopOrderRequest
        - [X] MarketIfTouchedOrderRequest
        - [X] StopLossOrderRequest
        - [X] TakeProfitOrderRequest
        - [X] GuaranteedStopLossOrderRequest
        - [X] TrailingStopLossOrderRequest
    - [X] [GET] OrderList
    - [X] [GET] OrderPendingList
    - [X] [GET] OrderDetails
    - [X] [PUT] OrderReplace
    - [X] [PUT] OrderCancel
    - [ ] [PUT] OrderUpdateClientExt
- Trade
    - [ ]Need to create data struct, now only json response.
    - [X] [GET] TradeList
    - [X] [GET] TradesOpen
    - [X] [GET] TradeDetails
    - [X] [PUT] TradeClose
    - [ ] [PUT] TradeUpdateClientExt
    - [X] [PUT] TradeUpdateTPSL
- Position
    - [ ]Need to create data struct, now only json response.
    - [X] [GET] PositionList
    - [X] [GET] PositionOpenList
    - [X] [GET] PositionByAccountID
    - [X] [PUT] PositionCloseoutByInstrument
- Transaction
- Pricing
    - [x] [GET] CandlesLatest
    - [x] [GET] PricingInformation
    - [ ] [GET] PricingStream
        - Need to figure out timeout issue when retrieving data
    - [x] [GET] CandlestickInstrument

#### Features

- [ ] Function to convert struct to json format
- [ ] Add POST request function
