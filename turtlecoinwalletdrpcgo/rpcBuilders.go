package turtlecoinwalletdrpcgo

// rpcPayload is the struct with the right formatting for passing to the post request
type rpcPayload struct {
	JSONRPC  string                  `json:"jsonrpc"`
	Method   string                  `json:"method"`
	Params   *map[string]interface{} `json:"params,omitempty"`
	Password string                  `json:"password"`
	ID       int                     `json:"id"`
}

func buildRPC(
	method string,
	id int,
	rpcPassword string,
	params map[string]interface{}) rpcPayload {

	return rpcPayload{
		JSONRPC:  "2.0",
		Method:   method,
		Password: rpcPassword,
		ID:       id,
		Params:   &params}
}

func reset(
	id int,
	rpcPassword string,
	params map[string]interface{}) rpcPayload {

	return buildRPC(
		"reset",
		id,
		rpcPassword,
		params)
}

func save(
	id int,
	rpcPassword string,
	params map[string]interface{}) rpcPayload {

	return buildRPC(
		"save",
		id,
		rpcPassword,
		params)
}

func getViewKey(
	id int,
	rpcPassword string,
	params map[string]interface{}) rpcPayload {

	return buildRPC(
		"getViewKey",
		id,
		rpcPassword,
		params)
}

func getSpendKeys(
	id int,
	rpcPassword string,
	params map[string]interface{}) rpcPayload {

	return buildRPC(
		"getSpendKeys",
		id,
		rpcPassword,
		params)
}

func getStatus(
	id int,
	rpcPassword string,
	params map[string]interface{}) rpcPayload {

	return buildRPC(
		"getStatus",
		id,
		rpcPassword,
		params)
}

func getAddresses(
	id int,
	rpcPassword string,
	params map[string]interface{}) rpcPayload {

	return buildRPC(
		"getAddresses",
		id,
		rpcPassword,
		params)
}

func createAddress(
	id int,
	rpcPassword string,
	params map[string]interface{}) rpcPayload {

	return buildRPC(
		"createAddress",
		id,
		rpcPassword,
		params)
}

func deleteAddress(
	id int,
	rpcPassword string,
	params map[string]interface{}) rpcPayload {

	return buildRPC(
		"deleteAddress",
		id,
		rpcPassword,
		params)
}

func getBalance(
	id int,
	rpcPassword string,
	params map[string]interface{}) rpcPayload {

	return buildRPC(
		"getBalance",
		id,
		rpcPassword,
		params)
}

func getBlockHashes(
	id int,
	rpcPassword string,
	params map[string]interface{}) rpcPayload {

	return buildRPC(
		"getBlockHashes",
		id,
		rpcPassword,
		params)
}

func getTransactionHashes(
	id int,
	rpcPassword string,
	params map[string]interface{}) rpcPayload {

	return buildRPC(
		"getTransactionHashes",
		id,
		rpcPassword,
		params)
}

func getTransactions(
	id int,
	rpcPassword string,
	params map[string]interface{}) rpcPayload {

	return buildRPC(
		"getTransactionHashes",
		id,
		rpcPassword,
		params)
}

func getUnconfirmedTransactionHashes(
	id int,
	rpcPassword string,
	params map[string]interface{}) rpcPayload {

	return buildRPC(
		"getUnconfirmedTransactionHashes",
		id,
		rpcPassword,
		params)
}

func getTransaction(
	id int,
	rpcPassword string,
	params map[string]interface{}) rpcPayload {

	return buildRPC(
		"getTransaction",
		id,
		rpcPassword,
		params)
}

func sendTransaction(
	id int,
	rpcPassword string,
	params map[string]interface{}) rpcPayload {

	return buildRPC(
		"sendTransaction",
		id,
		rpcPassword,
		params)
}

func createDelayedTransaction(
	id int,
	rpcPassword string,
	params map[string]interface{}) rpcPayload {

	return buildRPC(
		"createDelayedTransaction",
		id,
		rpcPassword,
		params)
}

func getDelayedTransactionHashes(
	id int,
	rpcPassword string,
	params map[string]interface{}) rpcPayload {

	return buildRPC(
		"getDelayedTransactionHashes",
		id,
		rpcPassword,
		params)
}

func deleteDelayedTransaction(
	id int,
	rpcPassword string,
	params map[string]interface{}) rpcPayload {

	return buildRPC(
		"deleteDelayedTransaction",
		id,
		rpcPassword,
		params)
}

func sendDelayedTransaction(
	id int,
	rpcPassword string,
	params map[string]interface{}) rpcPayload {

	return buildRPC(
		"sendDelayedTransaction",
		id,
		rpcPassword,
		params)
}

func sendFusionTransaction(
	id int,
	rpcPassword string,
	params map[string]interface{}) rpcPayload {

	return buildRPC(
		"sendFusionTransaction",
		id,
		rpcPassword,
		params)
}

func estimateFusion(
	id int,
	rpcPassword string,
	params map[string]interface{}) rpcPayload {

	return buildRPC(
		"estimateFusion",
		id,
		rpcPassword,
		params)
}
