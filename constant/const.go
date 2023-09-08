package constant

const (
	HostIntegration              = "https://webpay3gint.transbank.cl"
	HostProduction               = "https://webpay3g.transbank.cl"
	MallCommerceCodeIntegration  = "597055555541"
	StoreCommerceCodeIntegration = 597055555542
	TimeOut                      = 120
	SecretKeyIntegration         = "579B532A7440BB0C9079DED94D31EA1615BACEB56610332264630D42D0A36B1C"
	Path                         = "rswebpaytransaction/api/webpay/v1.2/transactions"
	PathOneClick                 = "/rswebpaytransaction/api/oneclick/v1.2/"
)

var (
	Host             string
	MallCommerceCode string
	SecretKey        string
)
