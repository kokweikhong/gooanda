package kw

type transactionFilter struct {
	ORDER                                 string // Order-related Transactions. These are the Transactions that create, cancel, fill or trigger Orders
	FUNDING                               string // Funding-related Transactions
	ADMIN                                 string // Administrative Transactions
	CREATE                                string // Account Create Transaction
	CLOSE                                 string // Account Close Transaction
	REOPEN                                string // Account Reopen Transaction
	CLIENT_CONFIGURE                      string // Client Configuration Transaction
	CLIENT_CONFIGURE_REJECT               string // Client Configuration Reject Transaction
	TRANSFER_FUNDS                        string // Transfer Funds Transaction
	TRANSFER_FUNDS_REJECT                 string // Transfer Funds Reject Transaction
	MARKET_ORDER                          string // Market Order Transaction
	MARKET_ORDER_REJECT                   string // Market Order Reject Transaction
	LIMIT_ORDER                           string // Limit Order Transaction
	LIMIT_ORDER_REJECT                    string // Limit Order Reject Transaction
	STOP_ORDER                            string // Stop Order Transaction
	STOP_ORDER_REJECT                     string // Stop Order Reject Transaction
	MARKET_IF_TOUCHED_ORDER               string // Market if Touched Order Transaction
	MARKET_IF_TOUCHED_ORDER_REJECT        string // Market if Touched Order Reject Transaction
	TAKE_PROFIT_ORDER                     string // Take Profit Order Transaction
	TAKE_PROFIT_ORDER_REJECT              string // Take Profit Order Reject Transaction
	STOP_LOSS_ORDER                       string // Stop Loss Order Transaction
	STOP_LOSS_ORDER_REJECT                string // Stop Loss Order Reject Transaction
	GUARANTEED_STOP_LOSS_ORDER            string // Guaranteed Stop Loss Order Transaction
	GUARANTEED_STOP_LOSS_ORDER_REJECT     string // Guaranteed Stop Loss Order Reject Transaction
	TRAILING_STOP_LOSS_ORDER              string // Trailing Stop Loss Order Transaction
	TRAILING_STOP_LOSS_ORDER_REJECT       string // Trailing Stop Loss Order Reject Transaction
	ONE_CANCELS_ALL_ORDER                 string // One Cancels All Order Transaction
	ONE_CANCELS_ALL_ORDER_REJECT          string // One Cancels All Order Reject Transaction
	ONE_CANCELS_ALL_ORDER_TRIGGERED       string // One Cancels All Order Trigger Transaction
	ORDER_FILL                            string // Order Fill Transaction
	ORDER_CANCEL                          string // Order Cancel Transaction
	ORDER_CANCEL_REJECT                   string // Order Cancel Reject Transaction
	ORDER_CLIENT_EXTENSIONS_MODIFY        string // Order Client Extensions Modify Transaction
	ORDER_CLIENT_EXTENSIONS_MODIFY_REJECT string // Order Client Extensions Modify Reject Transaction
	TRADE_CLIENT_EXTENSIONS_MODIFY        string // Trade Client Extensions Modify Transaction
	TRADE_CLIENT_EXTENSIONS_MODIFY_REJECT string // Trade Client Extensions Modify Reject Transaction
	MARGIN_CALL_ENTER                     string // Margin Call Enter Transaction
	MARGIN_CALL_EXTEND                    string // Margin Call Extend Transaction
	MARGIN_CALL_EXIT                      string // Margin Call Exit Transaction
	DELAYED_TRADE_CLOSURE                 string // Delayed Trade Closure Transaction
	DAILY_FINANCING                       string // Daily Financing Transaction
	RESET_RESETTABLE_PL                   string // Reset Resettable PL Transaction
}

var (
	TRANSACTIONFILTER = &transactionFilter{
		ORDER:                                 "ORDER",
		FUNDING:                               "FUNDING",
		ADMIN:                                 "ADMIN",
		CREATE:                                "CREATE",
		CLOSE:                                 "CLOSE",
		REOPEN:                                "REOPEN",
		CLIENT_CONFIGURE:                      "CLIENT_CONFIGURE",
		CLIENT_CONFIGURE_REJECT:               "CLIENT_CONFIGURE_REJECT",
		TRANSFER_FUNDS:                        "TRANSFER_FUNDS",
		TRANSFER_FUNDS_REJECT:                 "TRANSFER_FUNDS_REJECT",
		MARKET_ORDER:                          "MARKET_ORDER",
		MARKET_ORDER_REJECT:                   "MARKET_ORDER_REJECT",
		LIMIT_ORDER:                           "LIMIT_ORDER",
		LIMIT_ORDER_REJECT:                    "LIMIT_ORDER_REJECT",
		STOP_ORDER:                            "STOP_ORDER",
		STOP_ORDER_REJECT:                     "STOP_ORDER_REJECT",
		MARKET_IF_TOUCHED_ORDER:               "MARKET_IF_TOUCHED_ORDER",
		MARKET_IF_TOUCHED_ORDER_REJECT:        "MARKET_IF_TOUCHED_ORDER_REJECT",
		TAKE_PROFIT_ORDER:                     "TAKE_PROFIT_ORDER",
		TAKE_PROFIT_ORDER_REJECT:              "TAKE_PROFIT_ORDER_REJECT",
		STOP_LOSS_ORDER:                       "STOP_LOSS_ORDER",
		STOP_LOSS_ORDER_REJECT:                "STOP_LOSS_ORDER_REJECT",
		GUARANTEED_STOP_LOSS_ORDER:            "GUARANTEED_STOP_LOSS_ORDER",
		GUARANTEED_STOP_LOSS_ORDER_REJECT:     "GUARANTEED_STOP_LOSS_ORDER_REJECT",
		TRAILING_STOP_LOSS_ORDER:              "TRAILING_STOP_LOSS_ORDER",
		TRAILING_STOP_LOSS_ORDER_REJECT:       "TRAILING_STOP_LOSS_ORDER_REJECT",
		ONE_CANCELS_ALL_ORDER:                 "ONE_CANCELS_ALL_ORDER",
		ONE_CANCELS_ALL_ORDER_REJECT:          "ONE_CANCELS_ALL_ORDER_REJECT",
		ONE_CANCELS_ALL_ORDER_TRIGGERED:       "ONE_CANCELS_ALL_ORDER_TRIGGERED",
		ORDER_FILL:                            "ORDER_FILL",
		ORDER_CANCEL:                          "ORDER_CANCEL",
		ORDER_CANCEL_REJECT:                   "ORDER_CANCEL_REJECT",
		ORDER_CLIENT_EXTENSIONS_MODIFY:        "ORDER_CLIENT_EXTENSIONS_MODIFY",
		ORDER_CLIENT_EXTENSIONS_MODIFY_REJECT: "ORDER_CLIENT_EXTENSIONS_MODIFY_REJECT",
		TRADE_CLIENT_EXTENSIONS_MODIFY:        "TRADE_CLIENT_EXTENSIONS_MODIFY",
		TRADE_CLIENT_EXTENSIONS_MODIFY_REJECT: "TRADE_CLIENT_EXTENSIONS_MODIFY_REJECT",
		MARGIN_CALL_ENTER:                     "MARGIN_CALL_ENTER",
		MARGIN_CALL_EXTEND:                    "MARGIN_CALL_EXTEND",
		MARGIN_CALL_EXIT:                      "MARGIN_CALL_EXIT",
		DELAYED_TRADE_CLOSURE:                 "DELAYED_TRADE_CLOSURE",
		DAILY_FINANCING:                       "DAILY_FINANCING",
		RESET_RESETTABLE_PL:                   "RESET_RESETTABLE_PL",
	}
)
