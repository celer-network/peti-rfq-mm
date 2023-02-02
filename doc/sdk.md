# SDK
[rfq](#rfq)

[rfqmm](#rfqmm)

## rfq
```
import "github.com/celer-network/peti-rfq-mm/sdk/service/rfq"
import rfqproto "github.com/celer-network/peti-rfq-mm/sdk/service/rfq/proto"
```
Package rfq provides the client of RFQ Server.

Package rfqproto provides the types for making API requests to RFQ Server.

### Using the client
To contact RFQ Server with the SDK use the NewClient function to create a new service client. With that client you can make API
requests to the server.

See [NewClient](#func-newclient) for more information on creating client for this service.

### Types
- [type Client](#type-client)
    - [func NewClient(server string, apiKey string, ops ...grpc.DialOption) *Client](#func-newclient)
    - [func (c *Client) PendingOrders(ctx context.Context, in *rfqproto.PendingOrdersRequest, opts ...grpc.CallOption) (*rfqproto.PendingOrdersResponse, error)](#func-client-pendingorders)
    - [func (c *Client) UpdateOrders(ctx context.Context, in *rfqproto.UpdateOrdersRequest, opts ...grpc.CallOption) (*rfqproto.UpdateOrdersResponse, error)](#func-client-updateorders)
    - [func (c *Client) UpdateConfigs(ctx context.Context, in *rfqproto.UpdateConfigsRequest, opts ...grpc.CallOption) (*rfqproto.UpdateConfigsResponse, error)](#func-client-updateconfigs)

### Protos
- [message PendingOrder](#message-pendingorder)
- [message OrderStatus](#message-orderstatus)
- [message PendingOrdersRequest](#message-pendingordersrequest)
- [message PendingOrdersResponse](#message-pendingordersresponse)
- [message PendingOrdersRequest](#message-updateordersrequest)
- [message PendingOrdersResponse](#message-updateordersresponse)
- [message PendingOrdersRequest](#message-updateconfigsrequest)
- [message PendingOrdersResponse](#message-updateconfigsresponse)

#### type Client
```go
type Client struct {
	rfqproto.MMApiClient
	// rfq server URL
	server string
	conn   *grpc.ClientConn
}
```
#### func NewClient
```go
func NewClient(server string, apiKey string, ops ...grpc.DialOption) *Client
```
NewClient creates a new instance of Client. Get the correct server url and request an API key before taking a try.

Example:
```go
rfqServerUrl := "<rfq-server-url>"
rfqServerApiKey := "<your-api-key>"
client := rfq.NewClient(rfqServerUrl, rfqServerApiKey, grpc.WithTransportCredentials(credentials.NewTLS(&tls.Config{})))
```
#### func (*Client) PendingOrders
```go
func (c *Client) PendingOrders(ctx context.Context, in *rfqproto.PendingOrdersRequest, opts ...grpc.CallOption) (*rfqproto.PendingOrdersResponse, error)
```
PendingOrders API returns all orders related to this Client with status in set
{`OrderStatus_STATUS_SRC_DEPOSITED`, `OrderStatus_STATUS_DST_TRANSFERRED`}.

Example:
```go
resp, err := client.PendingOrders(context.Background(), &rfqproto.PendingOrdersRequest{})
if err != nil {
    // handle error
}
```
#### func (*Client) UpdateOrders
```go
func (c *Client) UpdateOrders(ctx context.Context, in *rfqproto.UpdateOrdersRequest, opts ...grpc.CallOption) (*rfqproto.UpdateOrdersResponse, error)
```
UpdateOrders API updates orders maintained by RFQ Server.

Example:
```go
_, err := client.UpdateOrders(context.Background(),
	&rfqproto.UpdateOrdersRequest{
		Updates: []*rfqproto.OrderUpdates{{QuoteHash: quoteHash, OrderStatus: toStatus, ExecTxHash: txHash}},
	})
if err != nil {
    // handle error
}
```
#### func (*Client) UpdateConfigs
```go
func (c *Client) UpdateConfigs(ctx context.Context, in *rfqproto.UpdateConfigsRequest, opts ...grpc.CallOption) (*rfqproto.UpdateConfigsResponse, error)
```
UpdateConfigs API sends supported token list to RFQ Server.

Example:
```go
request := &rfqproto.UpdateConfigsRequest{Config: &rfqmmproto.Config{Tokens: tokens}}
_, err := client.UpdateConfigs(context.Background(), request)
if err != nil {
    // handle error
}
```

#### message PendingOrder
```protobuf
package service.rfq;
import "service/rfqmm/api.proto";
message PendingOrder {
  // quote defined in rfqmmproto
  service.rfqmm.Quote quote = 1;
  string src_deposit_tx_hash = 2;
  // indicates whether the user wants native token on the dst chain (only applicable if the dst token is a native wrap)
  bool dst_native = 3;
  // unix epoch seconds
  bytes exec_msg_call_data = 4;
  string quote_sig = 5;
  rfq.OrderStatus status = 6;
}
```

#### message OrderStatus
```protobuf
enum OrderStatus {
  // an order first enters the pending status when an MM returns a firm quote upon a user's quote request.
  STATUS_PENDING = 0;
  // reached after the user's call to the RFQ contract to deposit funds
  STATUS_SRC_DEPOSITED = 10;
  // reached only if an MM actively calls back to the RFQ server to mark the order as rejected. once marked as rejected, 
  // the same order will not appear in the PendingOrders() anymore.
  // note that MMs can choose to not implement this active call and hence this status will never be reached.
  STATUS_MM_REJECTED = 20;
  // reached only if an MM actively calls back to the RFQ server to mark the order as dst executed
  // when they finish submitting the tx on the dst chain to transfer fund to user.
  // note that MMs can choose to not implement this active call and hence this status will never be reached.
  STATUS_MM_DST_EXECUTED = 30;
  // this status marks the observation of the on-chain event DstTransferred
  // this also means that msg2 is on its way but not yet arrived on the src chain
  // note that to the user, when an order reaches this status, it can be considered completed
  STATUS_DST_TRANSFERRED = 40;
  // reached only if an MM actively calls back to the RFQ server to mark the order as src executed
  // when they finish submitting the tx on the src chain to release fund to MM.
  // note that MMs can choose to not implement this active call and hence this status will never be reached.
  STATUS_MM_SRC_EXECUTED = 50;
  // this status marks the observation of the on-chain event RefundInitiated upon msg1 execution
  STATUS_REFUND_INITIATED = 60;
  // this status marks the observation of the on-chain event SrcReleased upon msg2 execution
  STATUS_SRC_RELEASED = 70;
  // this status marks the observation of the on-chain event Refunded upon msg3 execution
  STATUS_REFUNDED = 80;
}
```

#### message PendingOrdersRequest
```protobuf
message PendingOrdersRequest {
}
```

#### message PendingOrdersResponse
```protobuf
message PendingOrdersResponse {
  repeated PendingOrder orders = 1;
}
```

#### message UpdateOrdersRequest
```protobuf
message OrderUpdates {
  string quote_hash = 1;
  OrderStatus order_status = 2;
  string exec_tx_hash = 3;
}

message UpdateOrdersRequest {
  repeated OrderUpdates updates = 1;
}
```

#### message UpdateOrdersResponse
```protobuf
message UpdateOrdersResponse {
}
```

#### message UpdateConfigsRequest
```protobuf
import "service/rfqmm/api.proto";
message UpdateConfigsRequest {
  service.rfqmm.Config config = 1;
}
```

#### message UpdateConfigsResponse
```protobuf
message UpdateConfigsResponse {
}
```

## rfqmm
```
import "github.com/celer-network/peti-rfq-mm/sdk/service/rfqmm"
import rfqmmproto "github.com/celer-network/peti-rfq-mm/sdk/service/rfqmm/proto"
```
Package rfqmm provides the client and MM application(server) and some default implementation of {ChainQuerier, LiquidityProvider,
AmountCalculator, RequestSigner}. Those four components are required to create a server of MM application.

Package rfqmmproto provides the types for making API requests to MM application.

### Using the client
To contact MM application with the SDK use the NewClient function to create a new service client. With that client you can make API
requests to the MM.

See [NewClient](#rfqmmfunc-newclient) for more information on creating client for this service.

### Using the server
To build an MM application with the SDK use the NewServer function to create a new server. With that client you can serve
requests and process pending orders from RFQ Server.

See [NewServer](#func-newserver) for more information on creating server.

### Types
- [type Client](#rfqmmtype-client)
  - [func NewClient(server string, ops ...grpc.DialOption) *Client](#rfqmmfunc-newclient)
  - [func (c *Client) Price(ctx context.Context, in *rfqmmproto.PriceRequest, opts ...grpc.CallOption) (*rfqmmproto.PriceResponse, error)](#func-client-price)
  - [func (c *Client) Quote(ctx context.Context, in *rfqmmproto.QuoteRequest, opts ...grpc.CallOption) (*rfqmmproto.QuoteResponse, error)](#func-client-quote)
  - [func (c *Client) Close()](#func-client-close)
- [interface ApiServer](#interface-apiserver)
- [type Server](#type-server)
  - [func NewServer(config *ServerConfig, client *rfq.Client, cm ChainQuerier, lp LiquidityProvider, ac AmountCalculator, rs RequestSigner) *Server](#func-newserver)
  - [func (s *Server) Serve(ops ...grpc.ServerOption)](#func-server-serve)
  - [func (s *Server) ReportConfigs()](#func-server-reportconfigs)
  - [func (s *Server) ValidateQuote(quote *proto.Quote, sig []byte) bool](#func-server-validatequote)
  - [func (s *Server) DefaultProcessOrder()](#func-server-defaultprocessorder)
  - [func (s *Server) StopProcessing(reason string)](#func-server-stopprocessing)
  - [func (s *Server) Price(ctx context.Context, request *rfqmmproto.PriceRequest) (response *rfqmmproto.PriceResponse, err error)](#func-server-price)
  - [func (s *Server) Quote(ctx context.Context, request *rfqmmproto.QuoteRequest) (response *rfqmmproto.QuoteResponse, err error)](#func-server-quote)
- [interface ChainQuerier](#interface-chainquerier)
- [type ChainManager](#type-chainmanager)
  - [func NewChainManager(configs []*RfqMmChainConfig) *ChainManager](#func-newchainmanager)
  - [func (cm *ChainManager) GetChain(chainId uint64) (*Chain, error)](#func-chainmanager-getchain)
  - [func (cm *ChainManager) GetRfqFee(srcChainId uint64, dstChainId uint64, amount *big.Int) (*big.Int, error)](#func-chainmanager-getrfqfee)
  - [func (cm *ChainManager) GetMsgFee(chainId uint64) (*big.Int, error)](#func-chainmanager-getmsgfee)
  - [func (cm *ChainManager) GetGasPrice(chainId uint64) (*big.Int, error)](#func-chainmanager-getgasprice)
  - [func (cm *ChainManager) GetNativeWrap(chainId uint64) (*common.Token, error)](#func-chainmanager-getnativetoken)
  - [func (cm *ChainManager) GetERC20Balance(chainId uint64, token eth.Addr, account eth.Addr) (*big.Int, error)](#func-chainmanager-geterc20balance)
  - [func (cm *ChainManager) GetNativeBalance(chainId uint64, account eth.Addr) (*big.Int, error)](#func-chainmanager-getnativebalance)
  - [func (cm *ChainManager) GetQuoteStatus(chainId uint64, quoteHash eth.Hash) (uint8, error)](#func-chainmanager-getquotestatus)
  - [func (cm *ChainManager) VerifyRfqEvent(chainId uint64, tx eth.Hash, evName string) (bool, error)](#func-chainmanager-verifyrfqevent)
- [interface LiquidityProvider](#interface-liquidityprovider)
- [type DefaultLiquidityProvider](#type-defaultliquidityprovider)
  - [func NewDefaultLiquidityProvider(cm *ChainManager, lm *LiqManager) *DefaultLiquidityProvider](#func-newdefaultliquidityprovider)
  - [func (lp DefaultLiquidityProvider) IsPaused() bool](#func-defaultliquidityprovider-ispaused)
  - [func (lp DefaultLiquidityProvider) GetTokens() []*common.Token](#func-defaultliquidityprovider-gettokens)
  - [func (lp DefaultLiquidityProvider) SetupTokenPairs(policies []string)](#func-defaultliquidityprovider-setuptokenpairs)
  - [func (lp DefaultLiquidityProvider) HasTokenPair(srcToken, dstToken *common.Token) bool](#func-defaultliquidityprovider-hastokenpair)
  - [func (lp DefaultLiquidityProvider) GetLiquidityProviderAddr(chainId uint64) (eth.Addr, error)](#func-defaultliquidityprovider-getliquidityprovideraddr)
  - [func (lp DefaultLiquidityProvider) AskForFreezing(chainId uint64, token eth.Addr, amount *big.Int, isNative bool) (int64, error)](#func-defaultliquidityprovider-askforfreezing)
  - [func (lp DefaultLiquidityProvider) FreezeLiquidity(chainId uint64, token eth.Addr, amount *big.Int, until int64, hash eth.Hash, isNative bool) error](#func-defaultliquidityprovider-freezeliquidity)
  - [func (lp DefaultLiquidityProvider) UnfreezeLiquidity(chainId uint64, hash eth.Hash) error](#func-defaultliquidityprovider-unfreezeliquidity)
  - [func (lp *DefaultLiquidityProvider) DstTransfer(transferNative bool, _quote rfq.RFQQuote, opts ...eth.TxOption) (eth.Hash, error)](#func-defaultliquidityprovider-dsttransfer)
  - [func (lp *DefaultLiquidityProvider) SrcRelease(_quote rfq.RFQQuote, _execMsgCallData []byte, opts ...eth.TxOption) (eth.Hash, error)](#func-defaultliquidityprovider-srcrelease)
- [type LiqManager](#type-liqmanager)
  - [func NewLiqManager(configs []*LPConfig) *LiqManager](#func-newliqmanager)
  - [func (lm *LiqManager) GetChains() []uint64](#func-liqmanager-getchains)
  - [func (lm *LiqManager) GetTokens() map[uint64][]*common.Token](#func-liqmanager-gettokens)
  - [func (lm *LiqManager) GetLiquidityProvider(chainId uint64) (eth.Addr, error)](#func-liqmanager-getliquidityprovider)
  - [func (lm *LiqManager) GetLiqNeedApprove(chainId uint64) ([]*common.Token, []*big.Int, error)](#func-liqmanager-getliqneedapprove)
  - [func (lm *LiqManager) AskForFreezing(chainId uint64, token eth.Addr, amount *big.Int) (int64, error)](#func-liqmanager-askforfreezing)
  - [func (lm *LiqManager) ReserveLiquidity(chainId uint64, token eth.Addr, amount *big.Int, until int64, hash eth.Hash) error](#func-liqmanager-reserveliquidity)
  - [func (lm *LiqManager) ConfirmLiquidity(chainId uint64, token eth.Addr, amount *big.Int, until int64, hash eth.Hash) error](#func-liqmanager-confirmliquidity)
  - [func (lm *LiqManager) UnfreezeLiquidity(chainId uint64, hash eth.Hash) error](#func-liqmanager-unfreezeliquidity)
  - [func (lm *LiqManager) TransferOutLiquidity(chainId uint64, token eth.Addr, amount *big.Int, hash eth.Hash) error](#func-liqmanager-transferoutliquidity)
  - [func (lm *LiqManager) ReleaseInLiquidity(chainId uint64, token eth.Addr, amount *big.Int) error](#func-liqmanager-releaseinliquidity)
  - [func (lm *LiqManager) ReleaseNative(chainId uint64) (bool, error)](#func-liqmanager-releasenative)
  - [func (lm *LiqManager) UpdateLiqAmt(querier ChainQuerier)](#func-liqmanager-updateliqamt)
  - [func (lm *LiqManager) GetLP(chainId uint64) (*LiqProvider, error)](#func-liqmanager-getlp)
  - [func (lm *LiqManager) GetSigner(chainId uint64) (eth.Addr, eth.Signer, error)](#func-liqmanager-getsigner)
- [interface AmountCalculator](#interface-amountcalculator)
- [type DefaultAmtCalculator](#type-defaultamtcalculator)
  - [func NewDefaultAmtCalculator(feeConfig *FeeConfig, querier ChainQuerier, priceProvider PriceProvider) *DefaultAmtCalculator](#func-newdefaultamtcalculator)
  - [func (ac *DefaultAmtCalculator) SetDstGasCost(gasCost uint64)](#func-defaultamtcalculator-setdstgascost)
  - [func (ac *DefaultAmtCalculator) SetSrcGasCost(gasCost uint64)](#func-defaultamtcalculator-setsrcgascost)
  - [func (ac *DefaultAmtCalculator) SetGlobalFeePerc(feePerc uint32) error](#func-defaultamtcalculator-setglobalfeeperc)
  - [func (ac *DefaultAmtCalculator) SetPerChainPairFeePercOverride(overrides []*ChainOverride) error](#func-defaultamtcalculator-setperchainpairfeepercoverride)
  - [func (ac *DefaultAmtCalculator) SetPerTokenPairFeePercOverride(overrides []*TokenOverride) error](#func-defaultamtcalculator-setpertokenpairfeepercoverride)
  - [func (ac *DefaultAmtCalculator) SetGasPrice(prices []*GasPrice)](#func-defaultamtcalculator-setgasprice)
  - [func (ac *DefaultAmtCalculator) CalRecvAmt(tokenIn *common.Token, tokenOut *common.Token, amountIn *big.Int) (amountOut *big.Int, releaseAmt *big.Int, fee *big.Int, err error)](#func-defaultamtcalculator-calrecvamt)
  - [func (ac *DefaultAmtCalculator) CalSendAmt(tokenIn *common.Token, tokenOut *common.Token, amountOut *big.Int) (*big.Int, *big.Int, *big.Int, error)](#func-defaultamtcalculator-calsendamt)
- [interface PriceProvider](#interface-priceprovider)
- [interface RequestSigner](#interface-requestsigner)




### Protos
- [message Price](#message-price)
- [message Quote](#message-quote)
- [message Config](#message-config)
- [message ErrCode](#message-errcode)
- [message PriceRequest](#message-pricerequest)
- [message PriceResponse](#message-priceresponse)
- [message QuoteRequest](#message-quoterequest)
- [message QuoteResponse](#message-quoteresponse)

#### (rfqmm)type Client
```go
type Client struct {
	proto.ApiClient
	// mm application url
	server string
	conn   *grpc.ClientConn
}
```
#### (rfqmm)func NewClient
```go
func NewClient(server string, ops ...grpc.DialOption) *Client
```
NewClient creates a new instance of Client.

Example:
```go
mmUrl := "<mm-application-url>"
client := rfqmm.NewClient(mmUrl, grpc.WithTransportCredentials(credentials.NewTLS(&tls.Config{})))
```

#### func (*Client) Price
```go
func (c *Client) Price(ctx context.Context, in *rfqmmproto.PriceRequest, opts ...grpc.CallOption) (*rfqmmproto.PriceResponse, error)
```
Price API is used to get price from MM for a swap within PriceRequest. In PriceRequest, at least one amount should be given. 
* If SrcAmount is given, MM application will return with how much DstToken will be received by User.
* If DstAmount is given, MM application will return with how much SrcToken User should deposit to receive such amount of DstToken.
    >NOTE. This has not yet been implemented by pre-build MM application.
* If both of SrcAmount and DstAmount is given, MM application will treat it as the first case.

Example:
```go
request := &rfqmmproto.PriceRequest{SrcToken: srcToken, DstToken: dstToken, SrcAmount: srcAmount, DstAmount: "", DstNative: false}
resp, err := client.Price(context.Background(), request)
if err != nil {
// handle error
}
```

#### func (*Client) Quote
```go
func (c *Client) Quote(ctx context.Context, in *rfqmmproto.QuoteRequest, opts ...grpc.CallOption) (*rfqmmproto.QuoteResponse, error)
```
Quote API is used to confirm a quotation from MM.

Example:
```go
// price is the result returned by MM application during Price request
// quote is built within RFQ Server based on price
request := &rfqmmproto.QuoteRequest{Price: price, Quote: quote, DstNative: false}
resp, err := client.Quote(context.Background(), request)
if err != nil {
// handle error
}
```

#### func (*Client) Close
```go
func (c *Client) Close()
```
Close Method closes the grpc connection used by this client.

#### interface ApiServer
```go
type ApiServer interface {
    Price(context.Context, *PriceRequest) (*PriceResponse, error)
    Quote(context.Context, *QuoteRequest) (*QuoteResponse, error)
}
```
ApiServer interface is required to be implemented by MM application

#### type Server
```go
type Server struct {
    Ctl chan bool
    RfqClient         *rfq.Client
    Config            *ServerConfig
    ChainCaller       ChainQuerier
    LiquidityProvider LiquidityProvider
    AmountCalculator  AmountCalculator
    RequestSigner     RequestSigner
}
```

#### func NewServer
```go
type ServerConfig struct {
    // the period for retrying report supported tokens to rfq server
    ReportRetryPeriod int64
    // the period for processing pending orders
    ProcessPeriod int64
    // indicates the period for a price to be valid
    PriceValidPeriod int64
    // minimum dst transfer period, in order to give mm enough time for dst transfer
    DstTransferPeriod int64
    // token pair policy list
    TPPolicyList []string
    // port num that mm would listen on
    PortListenOn int64
}
func NewServer(config *ServerConfig, client *rfq.Client, cm ChainQuerier, lp LiquidityProvider, ac AmountCalculator, rs RequestSigner) *Server
```
NewServer creates a new instance of MM application. Customize your own subcomponents or use the default implementation.

Example:
```go
// client := rfq.NewClient(...)
// chainQuerier := rfqmm.NewChainManager(...)
// liquidityManager := rfqmm.NewDefaultLiqManager(...)
// liquidityProvider := rfqmm.NewDefaultLiquidityProvider(...)
// amountCalculator := rfqmm.NewDefaultAmtCalculator(...)
// requestSigner, _ := liquidityManager.GetLP(...)
server := rfqmm.NewServer(serverConfig, client, chainQuerier, liquidityProvider, amountCalculator, requestSigner)
```

#### func (*Server) Serve
```go
func (s *Server) Serve(ops ...grpc.ServerOption)
```
Serve Method starts the Server. Then it will listen on specific port and serve requests. 

Example:
```go
// if you want to do anything else after Serve(), run it in a goroutine
server.Serve()
```

#### func (*Server) ReportConfigs
```go
func (s *Server) ReportConfigs()
```
ReportConfigs Method reports supported tokens of this MM application to RFQ Server via an internal rfq.client.

Example:
```go
server.ReportConfigs()
```

#### func (*Server) ValidateQuote
```go
func (s *Server) ValidateQuote(quote *proto.Quote, sig []byte) bool
```
ValidateQuote Method check basically a quote's validity.

Basic flow:
* verify quote signature
* check quote hash is matched

Example:
```go
ok := server.ValidateQuote(quote, sig)
if !ok {
	// handle invalid quote
}
```

#### func (*Server) DefaultProcessOrder
```go
func (s *Server) DefaultProcessOrder()
```
DefaultProcessOrder Method starts a goroutine for processing orders.

Basic flow:
* periodically get pending orders from RFQ Server
* process orders in parallel
  * if its status is `OrderStatus_STATUS_SRC_DEPOSITED`
    * check deadline for transfer has not yet passed
    * check tx which was sent from User on src chain contains expected event
    * check quote's status within RFQ contract on src chain
    * transfer token to User
    * update order
  * if its status is `OrderStatus_STATUS_DST_TRANSFERRED`
    * release token on src chain
    * update order

Example:
```go
server.DefaultProcessOrder()
```

#### func (*Server) StopProcessing
```go
func (s *Server) StopProcessing(reason string)
```
StopProcessing Method stops processing orders with a given reason.

Example:
```go
reason := "some reason"
server.StopProcessing(reason)
```

#### func (*Server) Price
```go
func (s *Server) Price(ctx context.Context, request *rfqmmproto.PriceRequest) (response *rfqmmproto.PriceResponse, err error)
```
Price API is a default implementation of responding a price request.

Basic flow:
* validate price request
* calculate price
* check if there is sufficient liquidity for requested token
* sign price

#### func (*Server) Quote
```go
func (s *Server) Quote(ctx context.Context, request *rfqmmproto.QuoteRequest) (response *rfqmmproto.QuoteResponse, err error)
```
Quote API is a default implementation of responding at a quote request.

Basic flow:
* validate quote request
* check price sig
* check release amount within request is correct
* check if there is sufficient liquidity for requested token
* sign quote

#### interface ChainQuerier
```go
type ChainQuerier interface {
	GetRfqFee(srcChainId, dstChainId uint64, amount *big.Int) (*big.Int, error)
	GetMsgFee(chainId uint64) (*big.Int, error)
	GetGasPrice(chainId uint64) (*big.Int, error)
	GetNativeWrap(chainId uint64) (*common.Token, error)
	GetERC20Balance(chainId uint64, token, account eth.Addr) (*big.Int, error)
	GetNativeBalance(chainId uint64, accoun eth.Addr) (*big.Int, error)
	GetQuoteStatus(chainId uint64, quoteHash eth.Hash) (uint8, error)
	VerifyRfqEvent(chainId uint64, tx eth.Hash, evName string) (bool, error)
}
```

#### type ChainManager
```go
type ChainManager struct {
	chains   map[uint64]*Chain
	// event id of RFQ contract event
	eventIDs map[string]eth.Hash
}
```
ChainManager is a default implementation of ChainQuerier.

#### func NewChainManager
```go
type RfqMmChainConfig struct {
    ChainId                                             uint64
    Name, Gateway                                       string
    BlkInterval, BlkDelay, MaxBlkDelta, ForwardBlkDelay uint64
    GasLimit                                            uint64
    AddGasEstimateRatio                                 float64
    // Legacy gas price flags
    AddGasGwei   uint64
    MinGasGwei   uint64
    MaxGasGwei   uint64
    ForceGasGwei string
    // EIP-1559 gas price flags
    MaxFeePerGasGwei         uint64
    MaxPriorityFeePerGasGwei uint64
    // if ProxyPort > 0, a proxy with this port will be created to support some special chain such as harmony, celo.
    // chainID will be used to determined which type proxy to create, so make sure the chainID is supported in the "endpoint-proxy"
    // create a proxy to the Gateway, and eth-client will be created to "127.0.0.1:ProxyPort"
    // more detail, https://github.com/celer-network/endpoint-proxy
    ProxyPort int

    Rfq    string
    Native *common.Token
}
func NewChainManager(configs []*RfqMmChainConfig) *ChainManager
```
NewChainManager creates a new instance of ChainManager.

Example:
```go
var chainConfig []*rfqmm.RfqMmChainConfig
err = viper.UnmarshalKey("multichain", &chainConfig)
if err != nil {
	// handle err
}
cm := rfqmm.NewChainManager(chainConfig)
```

#### func (*ChainManager) GetChain
```go
[func (cm *ChainManager) GetChain(chainId uint64) (*Chain, error)
```
GetChain Method returns the Chain with specific chainId.

Example:
```go
// get chain of Goerli
chain, err := cm.GetChain(5)
if err != nil {
	// handle err
}
```

#### func (*ChainManager) GetRfqFee
```go
func (cm *ChainManager) GetRfqFee(srcChainId uint64, dstChainId uint64, amount *big.Int) (*big.Int, error)
```
GetRfqFee Method get RFQ protocol fee amount by querying RFQ contract.

Example:
```go
// Goerli -> BSC Testnet. amount with decimals
rfqFee, err := cm.GetRfqFee(5, 97, 1000000000)
if err != nil {
	// handle err
}
```

#### func (*ChainManager) GetMsgFee
```go
func (cm *ChainManager) GetMsgFee(chainId uint64) (*big.Int, error)
```
GetMsgFee Method get required native token amount for sending a message with constant length 32.

Example:
```go
msgFee, err := cm.GetMsgFee(5)
if err != nil {
	// handle err
}
```

#### func (*ChainManager) GetGasPrice
```go
[func (cm *ChainManager) GetGasPrice(chainId uint64) (*big.Int, error)
```
GetGasPrice Method returns the suggested gas price on specific chain.

Example:
```go
// get chain of Goerli
gasPrice, err := cm.GetGasPrice(5)
if err != nil || gasPrice.Sign() == 0 {
	// handle err
}
```

#### func (*ChainManager) GetNativeWrap
```go
func (cm *ChainManager) GetNativeWrap(chainId uint64) (*common.Token, error)
```
GetNativeWrap Method get configured native token struct of specific chain.

Example:
```go
native, err := cm.GetNativeWrap(5)
if err != nil {
	// handle err
}
```

#### func (*ChainManager) GetERC20Balance
```go
func (cm *ChainManager) GetERC20Balance(chainId uint64, token eth.Addr, account eth.Addr) (*big.Int, error)
```
GetERC20Balance Method query and return requested ERC20 token balance.

Example:
```go
balance, err := cm.GetERC20Balance(5, token, account)
if err != nil {
	// handle err
}
```

#### func (*ChainManager) GetNativeBalance
```go
func (cm *ChainManager) GetNativeBalance(chainId uint64, account eth.Addr) (*big.Int, error)
```
GetNativeBalance Method query and return requested native token balance.

Example:
```go
balance, err := cm.GetNativeBalance(5, account)
if err != nil {
	// handle err
}
```

#### func (*ChainManager) GetQuoteStatus
```go
func (cm *ChainManager) GetQuoteStatus(chainId uint64, quoteHash eth.Hash) (uint8, error)
```
GetQuoteStatus Method query and return current status on chain of a specific quote.

Example:
```go
status, err := cm.GetQuoteStatus(5, quoteHash)
if err != nil {
    // handle err
}
```

#### func (*ChainManager) VerifyRfqEvent
```go
func (cm *ChainManager) VerifyRfqEvent(chainId uint64, tx eth.Hash, evName string) (bool, error)
```
VerifyRfqEvent Method try to find expected event within specific tx on specific chain.

Example:
```go
ok, err := cm.VerifyRfqEvent(5, txHash, "SrcDeposited")
if err != nil {
	// handle err
}
```

#### interface LiquidityProvider
```go
type LiquidityProvider interface {
	// IsPaused returns whether the LiquidityProvider is paused or not
	IsPaused() bool
	// GetTokens returns a list of all supported tokens
	GetTokens() []*common.Token
    // SetupTokenPairs sets up supported token pairs based on a given policy list.
    SetupTokenPairs(policies []string)
    // HasTokenPair check if a given token pair is supported
    HasTokenPair(srcToken, dstToken *common.Token) bool
	// GetLiquidityProviderAddr returns the address of liquidity provider on specified chain
	GetLiquidityProviderAddr(chainId uint64) (eth.Addr, error)
	// AskForFreezing checks if there is sufficient liquidity for specified token on specified chain and returns freeze time
	AskForFreezing(chainId uint64, token eth.Addr, amount *big.Int, isNative bool) (int64, error)
	// FreezeLiquidity will freeze a certain liquidity for specified amount until specified timestamp with an index of hash.
	FreezeLiquidity(chainId uint64, token eth.Addr, amount *big.Int, until int64, hash eth.Hash, isNative bool) error
	// UnfreezeLiquidity will try to unfreeze a certain liquidity with specified hash.
	UnfreezeLiquidity(chainId uint64, hash eth.Hash) error
	// DstTransfer should send tx on dstChain to transfer dstToken to user
	DstTransfer(transferNative bool, _quote rfq.RFQQuote, opts ...ethutils.TxOption) (eth.Hash, error)
	// SrcRelease should send tx on srcChain to release srcToken to mm
	SrcRelease(_quote rfq.RFQQuote, _execMsgCallData []byte, opts ...ethutils.TxOption) (eth.Hash, error)
}
```

#### type DefaultLiquidityProvider
```go
type DefaultLiquidityProvider struct {
	paused       bool
	txrs         map[uint64]*ethutils.Transactor
	chainManager *ChainManager
	liqManager   *LiqManager
}
```
DefaultLiquidityProvider is a default implementation of interface LiquidityProvider.

#### func NewDefaultLiquidityProvider
```go
func NewDefaultLiquidityProvider(cm *ChainManager, lm *LiqManager) *DefaultLiquidityProvider
```
NewDefaultLiquidityProvider creates a new instance of DefaultLiquidityProvider.

Example:
```go
cm := rfqmm.NewChainManager(chainConfig)
lm := rfqmm.NewLiqManager(lpConfig)
lp := rfqmm.NewDefaultLiquidityProvider(cm, lm)
```

#### func (*DefaultLiquidityProvider) IsPaused
```go
func (lp DefaultLiquidityProvider) IsPaused() bool
```
IsPaused Method returns whether the DefaultLiquidityProvider is paused or not.

#### func (*DefaultLiquidityProvider) GetTokens
```go
func (lp DefaultLiquidityProvider) GetTokens() []*common.Token
```
GetTokens Method returns a list of all supported tokens.

#### func (*DefaultLiquidityProvider) SetupTokenPairs
```go
func (lp DefaultLiquidityProvider) SetupTokenPairs(policies []string)
```
SetupTokenPairs Method sets up allowed token pairs according to policies.
Each policy string should be in one of the following formats:
>Note. Space is not allowed within any policy string.
1. `All`, means all supported tokens are grouped in pairs. If an MM supports 5 tokens on all chains, then this policy
will produce 20 pairs.
2. `Any2Of=<ChainId-TokenSymbol>,...`, means tokens described in policy are grouped in pairs.
    >Example: policy str = "Any2Of=5-USDC,5-USDT,97-USDC", token pairs = 5-USDC -> 5-USDT, 5-USDT -> 5-USDC, 5-USDC -> 97-USDC,
    97-USDC -> 5-USDC, 5-USDT -> 97-USDC, 5-USDT -> 97-USDC
3. `OneOf=<ChainId-TokenSymbol>,<ChainId-TokenSymbol>`, would produce only one token pair which is from the first token to
the second token.
    >Example: policy str = "OneOf=5-USDC,97-USDC", token pair = 5-USDC -> 97-USDC. Reverse direction is forbidden. Use Any2Of to
    support both directions.

#### func (*DefaultLiquidityProvider) HasTokenPair
```go
func (lp DefaultLiquidityProvider) HasTokenPair(srcToken, dstToken *common.Token) bool
```
HasTokenPair Method checks whether a token pair is allowed by this MM.

#### func (*DefaultLiquidityProvider) GetLiquidityProviderAddr
```go
func (lp DefaultLiquidityProvider) GetLiquidityProviderAddr(chainId uint64) (eth.Addr, error)
```
GetLiquidityProviderAddr Method returns the address of liquidity provider on specified chain.

Example:
```go
addr, err := lp.GetLiquidityProviderAddr(5)
if err != nil {
	// handle err
}
```

#### func (*DefaultLiquidityProvider) AskForFreezing
```go
func (lp DefaultLiquidityProvider) AskForFreezing(chainId uint64, token eth.Addr, amount *big.Int, isNative bool) (int64, error)
```
AskForFreezing Method checks if there is sufficient liquidity for specified token on specified chain and returns freeze time.
Freeze time indicates how long the requested liquidity will be frozen.

Example:
```go
freezeTime, err := lm.AskForFreezing(5, token, 1000000)
if err != nil {
	// handle err
}
```

#### func (*DefaultLiquidityProvider) FreezeLiquidity
```go
func (lp DefaultLiquidityProvider) FreezeLiquidity(chainId uint64, token eth.Addr, amount *big.Int, until int64, hash eth.Hash, isNative bool) error
```
FreezeLiquidity Method will freeze certain amount of specific liquidity with quoteHash until specific timestamp.
As native token and wrapped native token are managed differently, `isNative` is needed to indicate whether the frozen token
is native or not.

Example:
```go
// given a quote of type rfqmmproto.Quote
err := lp.FreezeLiquidity(quote.GetDstChainId(), quote.DstToken.GetAddr(), quote.GetDstAmt(), quote.SrcDeadline, quote.GetQuoteHash(), false)
if err != nil {
	// handle err
}
```

#### func (*DefaultLiquidityProvider) UnfreezeLiquidity
```go
func (lp DefaultLiquidityProvider) UnfreezeLiquidity(chainId uint64, hash eth.Hash) error
```
UnfreezeLiquidity Method will try to unfreeze a certain liquidity with specified hash.

Example:
```go
// given a quote of type rfqmmproto.Quote
err := lp.UnfreezeLiquidity(quote.GetDstChainId(), quote.GetQuoteHash())
if err != nil {
	// handle err
}
```

#### func (*DefaultLiquidityProvider) DstTransfer
```go
func (lp *DefaultLiquidityProvider) DstTransfer(transferNative bool, _quote binding.RFQQuote, opts ...eth.TxOption) (eth.Hash, error)
```
DstTransfer Method sends tx on dstChain to transfer dstToken to the User .

Example:
```go
// given a quote of type rfqmmproto.Quote
txHash, err := lp.DstTransfer(false, quote.ToQuoteOnChain())
if err != nil {
	// handle err
}
```

#### func (*DefaultLiquidityProvider) SrcRelease
```go
func (lp *DefaultLiquidityProvider) SrcRelease(_quote rfq.RFQQuote, _execMsgCallData []byte, opts ...eth.TxOption) (eth.Hash, error)
```
SrcRelease Method sends tx on srcChain to release srcToken to MM.

Example:
```go
// given a quote of type rfqmmproto.Quote
txHash, err := lp.SrcRelease(quote.ToQuoteOnChain(), _execMsgCallData)
if err != nil {
	// handle err
}
```

#### type LiqManager
```go
type LiqManager struct {
	LPs map[uint64]*LiqProvider
}
```
ChainManager is an example implementation of liquidity and provider account management. Each LiqProvider can be used as a [RequestSigner](#interface-requestsigner).

#### func NewLiqManager
```go
type LPConfig struct {
    ChainId       uint64
	Address       string
    Keystore      string
    Passphrase    string
    Liqs          []*LiquidityConfig
    ReleaseNative bool
}
type LiquidityConfig struct {
    Address    string
    Symbol     string
    Amount     string
    Approve    string
    Decimals   int32
    FreezeTime int64
}
func NewLiqManager(configs []*LPConfig) *LiqManager
```
NewLiqManager creates a new instance of LiqManager.

Example:
```go
var lpConfig []*rfqmm.LPConfig
err := viper.UnmarshalKey("lp", &lpConfig)
if err != nil {
	// handle err
}
lm := rfqmm.NewLiqManager(lpConfig)
```

#### func (*LiqManager) GetChains
```go
func (lm *LiqManager) GetChains() []uint64
```
GetChains Method returns a non-repeating list of chainId of all liquidity.

#### func (*LiqManager) GetTokens
```go
func (lm *LiqManager) GetTokens() map[uint64][]*common.Token
```
GetTokens Method returns a map from chainId to configured liquidity tokens.

#### func (*LiqManager) GetLiquidityProvider
```go
func (lm *LiqManager) GetLiquidityProvider(chainId uint64) (eth.Addr, error)
```
GetLiquidityProvider Method returns provider account's address of specific chain.

Example:
```go
addr, err := lm.GetLiquidityProvider(5)
if err != nil {
	// handle err
}
```

#### func (*LiqManager) GetLiqNeedApprove
```go
func (lm *LiqManager) GetLiqNeedApprove(chainId uint64) ([]*common.Token, []*big.Int, error)
```
GetLiqNeedApprove Method returns tokens with amount to be approved on specific chain.

Example:
```go
tokens, approveAmt, err := lm.GetLiqNeedApprove(5)
if err != nil {
	// handle err
}
for i, token := range tokens {
	// approve approveAmt[i] of token to RFQ contract
}
```

#### func (*LiqManager) AskForFreezing
```go
func (lm *LiqManager) AskForFreezing(chainId uint64, token eth.Addr, amount *big.Int) (int64, error)
```
AskForFreezing Method checks if there is sufficient liquidity for specified token on specified chain and returns freeze time.
Freeze time indicates how long the requested liquidity is reserved before the User deposit.

Example:
```go
freezeTime, err := lm.AskForFreezing(5, token, 1000000)
if err != nil {
	// handle err
}
```

#### func (*LiqManager) ReserveLiquidity
```go
func (lm *LiqManager) ReserveLiquidity(chainId uint64, token eth.Addr, amount *big.Int, until int64, hash eth.Hash) error
```
ReserveLiquidity Method used for reserving liquidity when the User confirms a quotation. Deadline of reservation and 
quoteHash should be supplied.

Example:
```go
// given a quote of type rfqmmproto.Quote
err := lm.ReserveLiquidity(quote.GetDstChainId(), quote.DstToken.GetAddr(), quote.GetDstAmt(), quote.SrcDeadline, quote.GetQuoteHash())
if err != nil {
	// handle err
}
```

#### func (*LiqManager) ConfirmLiquidity
```go
func (lm *LiqManager) ConfirmLiquidity(chainId uint64, token eth.Addr, amount *big.Int, until int64, hash eth.Hash) error
```
ConfirmLiquidity Method used for confirming liquidity when RFQ Server informs an MM that the User has deposited.
Deadline of confirmation and quoteHash should be supplied.

Example:
```go
// given a quote of type rfqmmproto.Quote
err := lm.ConfirmLiquidity(quote.GetDstChainId(), quote.DstToken.GetAddr(), quote.GetDstAmt(), quote.DstDeadline, quote.GetQuoteHash())
if err != nil {
    // handle err
}
```

#### func (*LiqManager) UnfreezeLiquidity
```go
func (lm *LiqManager) UnfreezeLiquidity(chainId uint64, hash eth.Hash) error
```
UnfreezeLiquidity Method used to unfreeze a certain liquidity by quoteHash. It applies to both of reservation and confirmation.

Example:
```go
// given a quote of type rfqmmproto.Quote
err := lm.UnfreezeLiquidity(quote.GetDstChainId(), quote.GetQuoteHash())
if err != nil {
	// handle err
}
```

#### func (*LiqManager) TransferOutLiquidity
```go
func (lm *LiqManager) TransferOutLiquidity(chainId uint64, token eth.Addr, amount *big.Int, hash eth.Hash) error
```
TransferOutLiquidity Method used to deduct liquidity amount after an MM has transferred token to the User.

Example:
```go
// given a quote of type rfqmmproto.Quote
err := lm.TransferOutLiquidity(quote.GetDstChainId(), quote.DstToken.GetAddr(), quote.GetDstAmt(), quote.GetQuoteHash())
if err != nil {
	// handle err
}
```

#### func (*LiqManager) ReleaseInLiquidity
```go
func (lm *LiqManager) ReleaseInLiquidity(chainId uint64, token eth.Addr, amount *big.Int) error
```
ReleaseInLiquidity Method used to augment liquidity amount after an MM has released token from src chain.

Example:
```go
// given a quote of type rfqmmproto.Quote
err := lm.ReleaseInLiquidity(quote.GetSrcChainId(), quote.SrcToken.GetAddr(), quote.GetSrcAmt())
if err != nil {
	// handle err
}
```

#### func (*LiqManager) ReleaseNative
```go
func (lm *LiqManager) ReleaseNative(chainId uint64) (bool, error)
```
ReleaseNative Method returns whether native token on specific chain is preferred during token releasing.

Example:
```go
// given a quote of type rfqmmproto.Quote
releaseNative, err := lm.ReleaseNative(quote.GetSrcChainId())
if err != nil {
	// handle err
}
```

#### func (*LiqManager) UpdateLiqAmt
```go
func (lm *LiqManager) UpdateLiqAmt(querier ChainQuerier)
```
UpdateLiqAmt Method updates local liquidity amount via a given ChainQuerier.

Example:
```go
// new a ChainManager that can be used as a ChainQuerier
cm := NewChainManager(chainConfig)
lm.UpdateLiqAmt(cm)
```

#### func (*LiqManager) GetLP
```go
func (lm *LiqManager) GetLP(chainId uint64) (*LiqProvider, error)
```
GetLP Method used to get a LiqProvider that can be used as a [RequestSigner](#interface-requestsigner).

Example:
```go
requestSigner, err := lm.GetLP(5)
if err != nil {
	// handle err
}
```

#### func (*LiqManager) GetSigner
```go
func (lm *LiqManager) GetSigner(chainId uint64) (eth.Addr, eth.Signer, error)
```
GetSigner Method returns the provider account address and an eth type signer which can be used to sign eth message or
construct a transactor.

Example:
```go
addr, signer, err := lm.GetSigner(5)
if err != nil {
	// handle err
}
```

#### interface AmountCalculator
```go
type AmountCalculator interface {
	CalRecvAmt(tokenIn, tokenOut *common.Token, amountIn *big.Int) (recvAmt, releaseAmt, fee *big.Int, err error)
	CalSendAmt(tokenIn, tokenOut *common.Token, amountOut *big.Int) (sendAmt, releaseAmt, fee *big.Int, err error)
}
```

#### type DefaultAmtCalculator
```go
type DefaultAmtCalculator struct {
    // fixed cost related fields
    DstGasCost uint64
    SrcGasCost uint64
    GasPrice   map[uint64]uint64

    // personalized fee related fieds
    // 100% = 1000000
    FeePercGlobal        uint32
    PerChainPairOverride map[uint64]map[uint64]uint32
    PerTokenPairOverride map[string]map[string]uint32

    // helper
    Querier       ChainQuerier
    PriceProvider PriceProvider
}
```
DefaultAmtCalculator is a default implementation of interface AmountCalculator

#### func NewDefaultAmtCalculator
```go
type FeeConfig struct {
    DstGasCost     uint64
    SrcGasCost     uint64
    PercGlobal     uint32
    ChainOverrides []*ChainOverride
    TokenOverrides []*TokenOverride
    GasPrices      []*GasPrice
}
type ChainOverride struct {
    SrcChainId, DstChainId uint64
    Perc                   uint32
}
type TokenOverride struct {
    SrcChainId, DstChainId uint64
    SrcToken, DstToken     string
    Perc                   uint32
}
type GasPrice struct {
    ChainId uint64
    Price   uint64
}
func NewDefaultAmtCalculator(feeConfig *FeeConfig, querier ChainQuerier, priceProvider PriceProvider) *DefaultAmtCalculator
```
NewDefaultAmtCalculator creates a new instance of DefaultAmtCalculator.

Example:
```go
feeConfig := new(rfqmm.FeeConfig)
err = viper.UnmarshalKey("fee", feeConfig)
if err != nil {
	// handle err
}
cm := rfqmm.NewChainManager(chainConfig)
// prepare your Price Provider
priceProvider := NewYourPriceProvider(...)
ac := rfqmm.NewDefaultAmtCalculator(feeConfig, cm, priceProvider)
```

#### func (*DefaultAmtCalculator) SetDstGasCost
```go
func (ac *DefaultAmtCalculator) SetDstGasCost(gasCost uint64)
```
SetDstGasCost Method sets gas cost charged by MM on dst chain .

#### func (*DefaultAmtCalculator) SetSrcGasCost
```go
func (ac *DefaultAmtCalculator) SetSrcGasCost(gasCost uint64)
```
SetSrcGasCost Mthtod sets gas cost charged by MM on src chain.


#### func (*DefaultAmtCalculator) SetGlobalFeePerc
```go
func (ac *DefaultAmtCalculator) SetGlobalFeePerc(feePerc uint32) error
```
SetGlobalFeePerc Method sets global fee percentage, of which maximum is 1000000(=100%).

Example:
```go
// set global fee percentage to 0.1%
err := ac.SetGlobalFeePerc(1000)
if err != nil {
	// handle err
}
```

#### func (*DefaultAmtCalculator) SetPerChainPairFeePercOverride
```go
func (ac *DefaultAmtCalculator) SetPerChainPairFeePercOverride(overrides []*ChainOverride) error
```
SetPerChainPairFeePercOverride Method override fee percentage per chain pair.

Example:
```go
overrides := []*ChainOverride{{5,97,2000}}
err := ac.SetPerChainPairFeePercOverride(overrides)
if err != nil {
	// handle err
}
```

#### func (*DefaultAmtCalculator) SetPerTokenPairFeePercOverride
```go
func (ac *DefaultAmtCalculator) SetPerTokenPairFeePercOverride(overrides []*TokenOverride) error
```
SetPerTokenPairFeePercOverride Method override fee percentage per token pair.

Example:
```go
overrides := []*TokenOverride{{5,97,"0x123...","0x234...",2000}}
err := ac.SetPerTokenPairFeePercOverride(overrides)
if err != nil {
	// handle err
}
```

#### func (*DefaultAmtCalculator) SetGasPrice
```go
func (ac *DefaultAmtCalculator) SetGasPrice(prices []*GasPrice)
```
SetGasPrice Method sets gas price charged for each gas in wei by MM.

Example:
```go
gasPrices := []*GasPrice{{5,10000000000/*10gwei*/}}
ac.SetGasPrice(gasPrices)
```

#### func (*DefaultAmtCalculator) CalRecvAmt
```go
func (ac *DefaultAmtCalculator) CalRecvAmt(tokenIn *common.Token, tokenOut *common.Token, amountIn *big.Int) (amountOut *big.Int, releaseAmt *big.Int, fee *big.Int, err error)
```
CalRecvAmt Method estimates how much `tokenOut` will be received by User, how much `tokenIn` will be received by MM and 
how much `tokenIn` is charged as fee in total.

Example:
```go
// given a request of type rfqmm.PriceRequest
receiveAmount, releaseAmount, fee, err = ac.CalRecvAmt(request.SrcToken, request.DstToken, request.GetSrcAmount())
if err != nil {
	// handle err
}
```

#### func (*DefaultAmtCalculator) CalSendAmt
```go
func (ac *DefaultAmtCalculator) CalSendAmt(tokenIn *common.Token, tokenOut *common.Token, amountOut *big.Int) (*big.Int, *big.Int, *big.Int, error)
```
CalSendAmt Method estimate how much `tokenIn` should be sent by User, how much `tokenIn` will be received by MM and how
much `tokenIn` is charged as fee in total.

Example:
```go
// given a request of type rfqmm.PriceRequest
receiveAmount, releaseAmount, fee, err = ac.CalSendAmt(request.SrcToken, request.DstToken, request.GetDstAmount())
if err != nil {
	// handle err
}
```

#### interface PriceProvider
```go
type PriceProvider interface {
	GetPrice(token *common.Token) (float64, error)
}
```
PriceProvider is an important part for creating a [DefaultAmtCalculator](#type-defaultamtcalculator)

#### interface RequestSigner
```go
type RequestSigner interface {
	Sign(data []byte) ([]byte, error)
	Verify(data, sig []byte) bool
}
```

#### message Price 
```protobuf
package common;
message Token {
  uint64 chain_id = 1;
  string symbol = 2;
  string address = 3;
  int32 decimals = 4;
  string name = 5;
  string logo_uri = 6;
}

package service.rfqmm;
import "common/token.proto";
message Price {
  common.Token src_token = 1;
  // src_amount reflects the total amount of src_token the user should deposit in the contract on 
  // the src chain it should include rfq protocol fee + msg fee + mm charged fee
  string src_amount = 2;
  common.Token dst_token = 3;
  string src_release_amount = 4;
  string dst_amount = 5;
  // fee = mm fee + msg fee + src tx gas cost + dst tx gas cost
  string fee_amount = 6;
  // unix epoch milliseconds. the time before which the price response is valid for Quote
  int64 valid_thru = 7;
  string mm_addr = 8;
  // sig(hash('rfq price', mm_addr, valid_thru, src_chain_id, token_in, amount_in, dst_chain_id, token_out, amount_out))
  // when calling Quote(), mm uses this signature to verify the price content is agreed by them previously
  // and is not beyond deadline.
  string sig = 9;
  // the maximum src deposit period that is expected by mm, will be started from the time when mm receives the quote request
  int64 src_deposit_period = 10;
  // the minimum dst transfer period that is expected by mm, will be started from the time when mm receives the quote request
  int64 dst_transfer_period = 11;
}
```
#### message Quote
```protobuf
import "common/token.proto";
message Quote {
  // quote hash
  string hash = 1;
  // the input token amount on the src chain
  common.Token src_token = 2;
  string src_amount = 3;
  // the token amount (same token as src_token) that the market maker will receive by filling this quote
  string src_release_amount = 4;
  // the token amount out on the dst chain to be received by the user
  common.Token dst_token = 5;
  string dst_amount = 6;
  // the deadline before which the user can submit on the src chain
  int64 src_deadline = 7;
  // the time after which the order is eligible for refund
  int64 dst_deadline = 8;
  // nonce that is determined by the server that is used to dedup quotes
  uint64 nonce = 9;
  // sender of the src tx (msg.sender). it's also the user who deposits the src fund
  string sender = 10; 
  // the receiver of the token on the dst chain
  string receiver = 11; 
  // the receiver of the refund (if any) on the src chain
  string refund_to = 12; 
  // the address of the liquidity provider who's going to transfer fund to the user on the dst chain
  string mm_addr = 13; 
}
```
#### message Config
```protobuf
import "common/token.proto";
message Config {
  repeated common.Token tokens = 2;
}
```
#### message ErrCode
```protobuf
enum ErrCode {
  ERROR_UNDEFINED = 0;
  ERROR_INVALID_ARGUMENTS = 1;
  ERROR_LIQUIDITY_PROVIDER = 2;
  ERROR_PRICE_PROVIDER = 3;
  ERROR_AMOUNT_CALCULATOR = 4;
  ERROR_REQUEST_SIGNER = 5;
  ERROR_LIQUIDITY_MANAGER = 6;
  ERROR_CHAIN_MANAGER = 7;
}
```
#### message PriceRequest
```protobuf
import "common/token.proto";
message PriceRequest {
  common.Token src_token = 1;
  common.Token dst_token = 2;
  string src_amount = 3;
  string dst_amount = 4;
  // indicates whether the user wants native token on the dst chain (only applicable if the dst token is a native wrap)
  bool dst_native = 5;
}
```
#### message PriceResponse
```protobuf
package common;
message Err {
  uint32 code = 1;
  string msg = 2;
}

package service.rfqmm;
import "common/error.proto";
message PriceResponse {
  common.Err err = 1;
  // if receiveAmount is specified in the request, it would be the price of receiveToken in sendToken and vice versa.
  Price price = 2;
}
```
#### message QuoteRequest
```protobuf
message QuoteRequest {
  Price price = 1;
  Quote quote = 2;
  // indicates whether the user wants native token on the dst chain (only applicable if the dst token is a native wrap)
  bool dst_native = 3;
}
```
#### message QuoteResponse
```protobuf
import "common/error.proto";
message QuoteResponse {
  common.Err err = 1;
  string quote_sig = 2;
}
```
