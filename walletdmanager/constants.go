package walletdmanager

const (
	// DefaultTransferFee is the default fee. It is expressed in TRTL
	DefaultTransferFee float64 = 0.1
	// DefaultTransferMixin is the default mixin
	DefaultTransferMixin = 7

	logWalletdCurrentSessionFilename = "serviceCurrentSession.log"
	logWalletdAllSessionsFilename    = "service.log"
	walletdLogLevel                  = "3" // should be at least 3 as I use some logs messages to confirm creation of wallet
	walletdCommandName               = "service"
	turtlecoindCommandName           = "TurtleCoind"
)
