package daraja

var routes = map[string]string{
	"production":        "https://api.safaricom.co.ke",
	"sandbox":           "https://sandbox.safaricom.co.ke",
	"oauth":             "/oauth/v1/generate?grant_type=client_credentials",
	"b2c":               "/mpesa/b2c/v1/paymentrequest",
	"c2bregister":       "/mpesa/c2b/v1/registerurl",
	"c2bsimulate":       "/mpesa/c2b/v1/simulate",
	"c2bv2simulate":     "/mpesa/c2b/v2/simulate",
	"accountbalance":    "/mpesa/accountbalance/v1/query",
	"transactionstatus": "/mpesa/transactionstatus/v1/query",
	"reversal":          "/mpesa/reversal/v1/request",
	"stkpush":           "/mpesa/stkpush/v1/processrequest",
	"stkquery":          "/mpesa/stkpushquery/v1/query",
}
