
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
    - [ ] [GET] OrderPendingList
    - [ ] [GET] OrderDetails
    - [ ] [PUT] OrderReplace
    - [ ] [PUT] OrderCancel
    - [ ] [PUT] OrderUpdateClientExt
- Trade
    - [ ] [GET] TradeList
    - [ ] [GET] TradesOpen
    - [ ] [GET] TradeDetails
    - [ ] [PUT] TradeClose
    - [ ] [PUT] TradeUpdateClientExt
    - [ ] [PUT] TradeUpdateTPSL
- Position
    - [ ] [GET] PositionList
    - [ ] [GET] PositionOpenList
    - [ ] [GET] PositionByAccountID
    - [ ] [PUT] PositionCloseoutByInstrument
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
