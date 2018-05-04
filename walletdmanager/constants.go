package walletdmanager

const (
	// DefaultTransferFee is the default fee. It is expressed in TRTL
	DefaultTransferFee float64 = 0.1
	// DefaultTransferMixin is the default mixin
	DefaultTransferMixin = 5

	logWalletdCurrentSessionFilename = "walletdCurrentSession.log"
	logWalletdAllSessionsFilename    = "walletd.log"
	walletdLogLevel                  = "3" // should be at least 3 as I use some logs messages to confirm creation of wallet
	walletdCommandName               = "walletd"
	turtlecoindCommandName           = "TurtleCoind"
)
