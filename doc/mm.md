# RFQ Market Maker (MM) docs

## Overview 

Request For Quote (RFQ) system is running on top of Celer Inter-Chain Message Framework ([Celer IM](https://im-docs.celer.network/developer/celer-im-overview)) to enable secure and efficient intra- or inter-chain token swaps.

This document describes the functions and operations of the market maker (MM), which is responsible for quoting and fulfilling orders for RFQ transactions.

## Outline

- [RFQ Basics](#rfq-basics)
  - [Reach an agreement](#reach-an-agreement)
  - [Swap on chain](#swap-on-chain)
    - [SrcDeposit](#srcdeposit)
    - [DstTransfer](#dsttransfer)
    - [SrcRelease](#srcrelease)
  - [Relayer](#relayer)
- [Become an MM](#become-an-mm)
  - [Request for MM qualification](#request-for-mm-qualification)
  - [Run MM application](#run-mm-application)
- [Default MM application](#default-mm-application)
  - [Installation](#installation)
  - [Configuration](#configuration)
  - [Running](#running)
- [Light MM application](#light-mm-application)
  - [Light MM installation](#light-mm-installation)
  - [Light MM Configuration](#light-mm-configuration)
  - [Light MM Running](#light-mm-running)
- [Customize your own MM application](#customize-your-own-mm-application)
  - [Customize subcomponents](#customize-subcomponents)
  - [Customize order processing](#customize-order-processing)
  - [Customize request serving](#customize-request-serving)
- [Support](#support)
  - [Information](#information)
  - [Debugging](#debugging)



## RFQ Basics

A successfull RFQ transaction consists two main processes:

1. User and MM reach a quote agreement through off-chain communications via an RFQ Server.
2. User and MM execute the quote by swapping tokens through the RFQ contracts and Celer IM.

### Reach an agreement

```
 ╔══════╗                            ╔════════════╗                           ╔═════╗
 ║      ║                            ║            ║                           ║     ║
 ║      ║                            ║      R     ║ < = Supported Tokens < =  ║     ║
 ║      ║                            ║      F     ║                           ║     ║
 ║   U  ║ = > Request Quotation = >  ║      Q     ║ = > Price Request = > = > ║     ║
 ║   S  ║                            ║            ║                           ║  M  ║
 ║   E  ║ < = < Quotation = < = < =  ║      S     ║ < = Price Response < = <  ║  M  ║
 ║   R  ║                            ║      E     ║                           ║  s  ║
 ║      ║ = > Confirm quotation = >  ║      V     ║ = > Quote Request = > = > ║     ║
 ║      ║                            ║      E     ║                           ║     ║
 ║      ║                            ║      R     ║ < = Quote Response < = <  ║     ║
 ║      ║                            ║            ║                           ║     ║
 ╚══════╝                            ╚════════════╝                           ╚═════╝
```

>Prerequisite: All MMs should report their supported tokens in a list to RFQ Server via [UpdateConfigs API](./sdk.md#func-client-updateconfigs)
once after MM is ready.

1. User requests quotation from RFQ Server for a possible swap: token X on chain A -> token Y on chain B
2. RFQ Server got the request from User, and takes a look at all MMs' token config in order to determine who is available to 
fulfill this swap. Then RFQ Server will send [Price request](./sdk.md#message-pricerequest) to all available MMs.
3. MM got the price request from RFQ Server, and calculate how much token Y on chain B he would like to pay for exchanging
token X on chain A, according to his fee strategy. Then MM will return his price response to RFQ Server along with his
signature of this price response and a period for this price to be valid.
4. RFQ Server collects price responses from available MMs and chooses only one MM for taking this swap order. **The MM of 
which price response has the highest amount of token Y on chain B will be chosen by RFQ Server.** Then the best price response
will be returned to User as the requested quotation of his possible swap.
5. If User accepts this quotation, he needs to confirm it through RFQ Server.
6. When User confirms the quotation, RFQ Server will send [Quote request](./sdk.md#message-quoterequest) to the chosen MM.
Quote request includes the signature produced by the MM during step 3, a suggested SrcDeadline for User by which User should finish
locking his token X on chain A, a suggested DstDeadline for MM by which MM should finish transferring token Y to User on chain B.
7. MM got the quote request from RFQ Server and verify his signature of price. As long as the 
   1. signature is valid,
   2. the period for this price to be valid has not yet passed,
   3. suggested SrcDeadline and DstDeadline are both acceptable,
   4. (optional) has sufficient token Y on chain B and freezes it,

    MM can sign this quotation and return this signature to RFQ Server for later verification. 

Once a good quote response is returned from an MM to RFQ Server, an agreement between certain User and MM is reached.

### Swap on chain
#### SrcDeposit
```
 ──────────────────────────────────────────────┬───────────────────────────────────────────────────
 CHAIN A                                       │                                            CHAIN B
 ┏━━━━┓                     ┏━━━━━━━━━━━━━┓    │    ┏━━━━━━━━━━━━━┓                    ┏━━━━┓            
 ┃USER┃ > = srcDeposit >  > ┃     RFQ     ┃    │    ┃     RFQ     ┃                    ┃ MM ┃
 ┗━━━━┛                     ┗━━━━━━━━━━━━━┛    │    ┗━━━━━━━━━━━━━┛                    ┗━━━━┛
                                   ∨ send      │                                          ∧
                                   v message1  │                                          ∧
                            ┏━━━━━━━━━━━━━┓    │    ┏━━━━━━━━━━━━━┓                       ∧ inform MM:
                            ┃ Message Bus ┃    │    ┃ Message Bus ┃                       ∧ User has
                            ┗━━━━━━━━━━━━━┛    │    ┗━━━━━━━━━━━━━┛                       ∧ deposited
                                   v           │                                          ∧
 ──────────────────────────────────────────────┴───────────────────────────────────────────────────
                                   v listened by sgn                                      ∧
                             ╔═══════════════════════════════════╗                   ╔════════════╗
                             ║   SGN (State Guardian Network)    ║ > query message > ║ RFQ Server ║
                             ╚═══════════════════════════════════╝                   ╚════════════╝                                              
                                                                                      
```
After User confirms a quotation, User is required to deposit token X to RFQ contract on chain A, by calling [srcDeposit](https://github.com/celer-network/sgn-v2-contracts/blob/c66326d458b9d34058ed960f610af69d8514716c/contracts/message/apps/RFQ.sol#L80).
During `srcDeposit`, a message would be sent via Message Bus, the core contract of Celer IM. As a consequence, the message
would be catched by SGN via event listener, and co-signed by SGN's validators. Once the message has sufficient voting power,
RFQ Server can fetch it from SGN and mark the corresponding swap to status `OrderStatus.STATUS_SRC_DEPOSITED`. Then the chosen MM will 
be informed(in a polling way, see [PendingOrders](./sdk.md#func-client-pendingorders)) that User has deposited on chain A.
Along with it, the signature generated by the MM during above step 7 will be returned to MM for verification.

#### DstTransfer
```
 ──────────────────────────────────────────────┬───────────────────────────────────────────────────
 CHAIN A                                       │                                            CHAIN B
                                               │                                       ┏━━━━┓
                                               │          > = > transfer token > = > > ┃USER┃
                                               │          ∧                            ┗━━━━┛
 ┏━━━━┓                     ┏━━━━━━━━━━━━━┓    │    ┏━━━━━━━━━━━━━┓                    ┏━━━━┓          
 ┃ MM ┃                     ┃     RFQ     ┃    │    ┃     RFQ     ┃ <  dstTransfer < < ┃ MM ┃
 ┗━━━━┛                     ┗━━━━━━━━━━━━━┛    │    ┗━━━━━━━━━━━━━┛                    ┗━━━━┛
   ∧                                           │          ∨ send                          
   ∧ inform mm                                 │          v message2                       
   ∧ and give him           ┏━━━━━━━━━━━━━┓    │    ┏━━━━━━━━━━━━━┓                       
   ∧ a proof                ┃ Message Bus ┃    │    ┃ Message Bus ┃                        
   ∧                        ┗━━━━━━━━━━━━━┛    │    ┗━━━━━━━━━━━━━┛                       
   ∧                                           │          v                               
 ──────────────────────────────────────────────┴───────────────────────────────────────────────────
   ∧                                    listened by sgn   v                               
 ╔════════════╗              ╔═══════════════════════════════════╗                   
 ║ RFQ Server ║  < message < ║   SGN (State Guardian Network)    ║ 
 ╚════════════╝              ╚═══════════════════════════════════╝                                                                 
                                                                                      
```
When MM is informed that User has deposited, MM should 
* verify the signature of quotation to make sure that MM did have made an agreement on it before
* double-check the validity of information from RFQ Server in case of a hacked or malicious server, that User did have
deposited token X on chain A. 

If everything goes well, then MM can call [dstTransfer](https://github.com/celer-network/sgn-v2-contracts/blob/c66326d458b9d34058ed960f610af69d8514716c/contracts/message/apps/RFQ.sol#L136) to transfer token Y to User on chain B.
During `dstTransfer`, 
* a message would be sent via Message Bus contract, catched by SGN, and co-signed by SGN's validators.
* certain amount of token Y would be transferred from MM, and transferred to User after the message is successfully sent out.

>NOTE: Light MM can choose to let a central service called relayer to submit tx for him on dst chain.

Once the message has sufficient voting power, RFQ Server can fetch it from SGN and mark the corresponding swap to status
`OrderStatus.STATUS_DST_TRANSFERRED`. Then the chosen MM will be informed(in a polling way, see [PendingOrders](./sdk.md#func-client-pendingorders))
that `dstTransfer` is successful, and `srcRelease`is available on chain A to release token. 
As a proof of order fulfillment generated by SGN is required to release token, RFQ Server will also help deliver it to MM.
 
#### SrcRelease
```
 ──────────────────────────────────────────────┬───────────────────────────────────────────────────
 CHAIN A                                       │                                            CHAIN B
    < < < < < transfer token < < < <           │
    v                              ∧           │
 ┏━━━━┓                     ┏━━━━━━━━━━━━━┓    │    ┏━━━━━━━━━━━━━┓                                
 ┃ MM ┃ > = srcRelease >  > ┃     RFQ     ┃    │    ┃     RFQ     ┃                    
 ┗━━━━┛                     ┗━━━━━━━━━━━━━┛    │    ┗━━━━━━━━━━━━━┛                    
                                   ∨ verify    │                                          
                                   v proof     │                                          
                            ┏━━━━━━━━━━━━━┓    │    ┏━━━━━━━━━━━━━┓                        
                            ┃ Message Bus ┃    │    ┃ Message Bus ┃                     
                            ┗━━━━━━━━━━━━━┛    │    ┗━━━━━━━━━━━━━┛                       
                                               │                                          
 ──────────────────────────────────────────────┴───────────────────────────────────────────────────
                                                                        
```
When MM got the proof of order fulfillment, MM can call [srcRelease](https://github.com/celer-network/sgn-v2-contracts/blob/c66326d458b9d34058ed960f610af69d8514716c/contracts/message/apps/RFQ.sol#L188) 
to release token X on chain A. During `srcRelease`, the proof is verified via MessageBus contract. If all checks are passed,
the locked token X which was deposited by User would be transferred to MM after deducting RFQ protocol fee. 

>NOTE: Light MM can choose to let a central service called relayer to submit tx for him on src chain.

Then the swap on chain between User and MM is completed.

### Relayer
Relayer is a central service hosted by Peti protocol to help those MMs who:
* doesn't want to request rfq server for reporting tokens and getting pending orders
* doesn't want to send any tx on both of dst and src chain.

As a consequence of that the relayer sends tx for MMs:
* MM should expose more api in order to finish the whole swap, including an api for signing specific data.
* Base fees(tx gas cost and message fee) would be charged by Peti protocol instead of MM, and be accumulated in RFQ contract.

Relayer's working mechanism is very simple and is based on MM's signature verification. Instead of directly sending tokens
to User, MM should sign the quote's hash to express that he allows a third party to transfer token from his address to user.
Then relayer can take MM's sig and call RFQ contract at `dstTransferWithSig`, where MM's signature is verified and tokens
are transferred from MM to User. 

Relayer would help release token on src chain for MM as well. This would require nothing more, because `srcRelease` is 
callable for anyone by design. And tokens will only be released to MM's address no matter who call `srcRelease`.

## Become an MM

### Request for MM qualification
An API key is needed for an MM to use RFQ Server's services. Contact us for requesting an API key.

### Run MM application
For default MM application, see the guide at [Default MM Application](#default-mm-application).

For light MM application, see the guide at [Light MM Application](#light-mm-application).

For customized MM application, run it as you preferred.

## Default MM application

### Installation

1. Download `peti-rfq-mm`
```
git clone https://github.com/celer-network/peti-rfq-mm.git
cd peti-rfq-mm
```
2. Build `peti-rfq-mm`
```
make install
```

### Configuration
Make a new folder to store your configuration file and ETH keystore file.
```
mkdir .peti-rfq-mm
cd .peti-rfq-mm
mkdir config eth-ks
touch config/chain.toml config/lp.toml config/fee.toml config/mm.toml
// move all used address's keystore file into .peti-rfq-mm/eth-ks/
mv <path-to-your-eth-keystore-file> eth-ks/<give-a-name>.json
```
The `.peti-rfq-mm` folder's structure will looks like:
```
.peti-rfq-mm/
  - config/
      - chain.toml
      - lp.toml
      - fee.toml
      - mm.toml
  - eth-ks/
      - <give-a-name>.json
      - <give-b-name>.json
```

1. Chain configuration

Each chain is configured by a `multichain`.
Take Goerli as an example. Before using, don't forget to update `chainId`, `name`, and fill up `gateway` and `rfq`.
RFQ contract address could be found at [Information](#information).

```
[[multichain]]
chainID = 5
name = "Goerli"
gateway = "<your-goerli-rpc>" # fill in your Goerli rpc provider url
rfq = "<copy-addr-from-'Support->Contract address'>"
blkdelay = 5 # how many blocks confirmations are required
blkinterval = 15 # polling interval for querying tx's status
# belows are optional transaction options
# maxfeepergasgwei = 10 # acceptable max fee price
# maxPriorityFeePerGasGwei = 2 # acceptable max priority fee price
# gaslimit = 200000 # fix gas limit and skip gas estimation, often used for debuging
# addgasestimateratio = 0.3 # adjust result from gas estimation, actual gasLimit = (1+addgasestimateratio)*estimation 
[multichain.native]
symbol = "ETH"
# if any liquidity of native token or wrapped native token on this chain is configured in lp.toml, this address should 
# be set, and set to wrapped native token address.
address = "<wrapped-native-token-address>"
decimals = 18
```

Transaction options are used in condition. Normally, if you got some error about "out of gas", we recommend using
`addgasestimateratio` with value `0.3` at first, and gradually increase its value if the error still occurs.

A dubug tip: if you got any error not about gas during pre-running, and you don't figure out the reason, try give a `gaslimit`
that is big enough. After the transaction is sent, debug it in [Tenderly](https://dashboard.tenderly.co/).

2. Liquidity configuration

Liquidity is configured per chain and per token. An example full configuration of liquidity on Goerli is:
```
[[lp]]
chainid = 5
address = "<lp address>"
# if this lp is a contract, should keep keystore unset or empty string
keystore = "./eth-ks/<give-a-name>.json"
passphrase = "<password-of-your-keystore>"
# release native token or wrapped native token on this chain, used when the token deposited by User is native token or wrapped native token
releasenative = false
[[lp.liqs]]
address = "0xf4B2cbc3bA04c478F0dC824f4806aC39982Dce73"
symbol = "USDT"
# token's available amount. if not set, would query the current token balance during initialization
amount = "5000000000"
# the amount of token to be approved to RFQ contract during initialization
approve = "1000000000000"
decimals = 6
# how long you prefer to freeze this token. The unit is second.
freezetime = 300
[[lp.liqs]]
# address of full `f` represents native token
address = "0xffffffffffffffffffffffffffffffffffffffff"
symbol = "ETH"
#amount = "0"
decimals = 18
freezetime = 200
```
You can use different account for each chain or just use one account for all chains. `keystore` should set to path of your
keystore file relative to `.peti-rfq-mm` floder.
If you're going to use a contract as a liquidity provider, let `keystore` be empty.

For each token, `address`, `symbol`, `decimals` and `freezetime` are required, while `amount` and `approve` are optional.

- `freezetime`: How long the MM prefer to freeze a token. Take 300 as example. It means, counting from the User confirm a 
quotation, he should finish depositing token within 300 second.
- `amount`: How much token the MM could supply. MM can set it to any value regardless of current token balance. If it is
not set, current token balance would be used instead.
- `approve`: How much token will be approved to RFQ contract. If it is set, transaction would be sent during initialization
for approving. *Once MM has approved sufficient amount, remove this field to prevent re-approve.*
- `address`: Token address. In particular, `0xffffffffffffffffffffffffffffffffffffffff` is used to represent native token.
*If native token is configured for one chain, relatively `multichain.native.address` must be set and set to wrapped native token address*.

3. Fee configuration

Fee strategy is configured globally with overrides per chain pair and per token pair. An example full configuration of 
fee is:
```
[fee]
# how much gas of dst chain you wanna charge, should be higher than actual consumption
dstgascost = 100000
# how much gas of src chain you wanna charge, should be higher than actual consumption
srcgascost = 150000
# global percentage fee, 100% = 1000000
percglobal = 1000

[[fee.gasprices]]
chainid = 5
# how much wei you wanna charge for each gas consumed when failed to call eth_gasPrice 
price = 5000000000

[[fee.gasprices]]
chainid = 97
# how much wei you wanna charge for each gas consumed when failed to call eth_gasPrice
price = 7000000000

[[fee.chainoverrides]]
# override percentage fee from srcchainid to dstchainid
srcchainid = 5
dstchainid = 97
perc = 2000

[[fee.tokenoverrides]]
# override percentage fee from srctoken on srcchainid to dsttoken on dstchainid
srcchainid = 5
srctoken = "0xf4B2cbc3bA04c478F0dC824f4806aC39982Dce73"
dstchainid = 97
dsttoken = "0x7d43AABC515C356145049227CeE54B608342c0ad"
perc = 3000
```

Except of `fee.chainoverrides` and `fee.tokenoverrides`, all the other fields are required for fee configuration. Generally,
MM needs to separately send one tx on dst and src chain, in order to complete a swap order. That's  why we need
to configure `fee.dstgascost`, `fee.srcgascost` and `fee.gasprice`. The actual charged fee value to cover gas consumption
on two chains will be
`fee.dstgascost * <current-gasprice-on-dst> * <current-native-token-price-in-wei> + fee.srcgascost * <current-gasprice-on-src> * <current-native-token-price-in-wei>`. 
At last, the fee value will be converted to token amount which is deducted from the amount of token transferred to User.

4. MM configuration

This configuration contains several important parameters related to MM application's operation.
```
[priceprovider]
# url required by default price provider implementation
url = "https://cbridge-stat.s3.us-west-2.amazonaws.com/prod2/cbridge-price.json"

[rfqserver]
url = "<url-of-rfq-server>"
apikey = "<your-api-key>"

[requestsigner]
# indicates which chain's signer will be used as request signer
chainid = 5
# Optional. if keystore(file path) is not empty, then the address denoted by the keystore will be used as request signer.
keystore = ""
passphrase = ""

[mm]
# token pair policy list indicates from which token to which token the mm is interested in
tpPolicyList = ["All"]
# port that mm listens on
portListenOn = 6666
# all periods' unit is second
# indicates the period during which a price response from this mm is valid
priceValidPeriod = 300
# indicates the minimum period for this mm to complete transferring on dst chain, couting from the user confirms the quotation
dstTransferPeriod = 600
# if faled to report token configs to rfq server, mm will be stucked and retry every <reportperiod> seconds until success.
reportRetryPeriod = 5
# time interval for getting and processing pending orders from rfq server
processPeriod = 5
# indicates whether this mm is light versioned
lightMM = false
# change host to "0.0.0.0" in need
host="localhost"
```
Token pair policy format can be found in [SDK doc](./sdk.md#func-defaultliquidityprovider-setuptokenpairs).

Do not modify `priceprovider.url`. A large json format data of token prices stored under `priceprovider.url`, and is updated
periodically by other external process. At present, it's the only implementation of price service within default MM application.
If you're not comfortable with this implementation, you can either try to customize your own MM application or waiting for
later updates of default MM application.

Get `rfqserver.url` at [Information](#information) and fill up your API key.

As mentioned before, MM should be able to sign any data and verify own signatures. Within default MM application, one 
of the accounts configured in `lp.toml` is used as request signer to sign price and quote response. The chosen account is
identified by `requestsigner.chainid` of which value matches with `lp.chainid`.

### Running

Create a peti-rfq-mm system service
```
touch /etc/systemd/system/peti-rfq-mm.service
# create the log directory for executor
mkdir -p /var/log/peti-rfq-mm
```

>IMPORTANT: check if the user, user group and paths defined in your systemd file are correct.

```
# peti-rfq-mm.service

[Unit]
Description=Default MM application
After=network-online.target

[Service]
Environment=HOME=/home/ubuntu
ExecStart=/home/ubuntu/go/bin/peti-rfq-mm start \
  --home /home/ubuntu/.peti-rfq-mm/ \
  --logdir /var/log/peti-rfq-mm/app --loglevel debug
StandardError=append:/var/log/peti-rfq-mm/error.log
Restart=always
RestartSec=10
User=ubuntu
Group=ubuntu
LimitNOFILE=4096

[Install]
WantedBy=multi-user.target
```
Enable and start executor service
```
sudo systemctl enable peti-rfq-mm
sudo systemctl start peti-rfq-mm
// Check if logs look ok
tail -f -n30 /var/log/peti-rfq-mm/app/<log-file-with-start-time>.log
```

## Light MM application

The difference between Light MM and Default MM:
* Light MM will not actively send any request to RFQ server.
* Light MM will serve more api request.
* Light MM will not send any tx on chain by himself. Tx for `dstTransfer` and `srcRelease` will be sent by Relayer.
* Light MM will not charge tx gas cost and message fee. It will charged by Peti protocol.

### Light MM Installation

See [installation](#installation). 

Before building, switch to `light-mm` branch. If this branch has already been merged to `main`, then the switch operation
is not needed.

### Light MM Configuration

Follow the [configuration](#configuration) of Default MM, and make following changes:
* Set `lightMM` in `mm.toml` to `true`
* `rfqserver` in `mm.tonl` is no longer needed
* `fee.dstgascost`, `fee.srcgascost` and `fee.gasprices` in `fee.toml` are no longer needed
* (In need) If request signer is different with liquidity provider, then liquidity provider is required to call 
`RegisterAllowedSigner` at RFQ contract to make request signer's signature valid.

### Light MM Running

See [running](#running).

## Customize your own MM application

With the [SDK](./sdk.md#sdk), the way to customize your own MM application is totally up to yourself, as long as it meets
minimum requirements:
* Implement [ApiServer](./sdk.md#interface-apiserver)
* Utilize [RFQ Client](./sdk.md#type-client) to report supported tokens to RFQ Server
* Utilize [RFQ Client](./sdk.md#type-client) to get pending orders, process orders, and update orders

Besides, MM is suggested to have the ability of customizing fee and managing his liquidity on different chain, which includes but not limited
to:

- flexible fee configuration
- freeze and unfreeze requested token at appropriate time
- reuse the just released token for next swap orders.
- withdraw and add liquidity from/to remote liquidity pool if needed (It's not supported now by default MM application)

At last, do not forget to start serving requests, for example:
```go
yourMMApp := NewYourMMApp(...)
listener, err := net.Listen("tcp", fmt.Sprintf("localhost:%d", port))
if err != nil {
	panic(err)
}
grpcServer := grpc.NewServer(ops...)
rfqmmproto.RegisterApiServer(grpcServer, yourMMApp)
grpcServer.Serve(listener)
```

But if you think the struture of [Server](./sdk.md#type-server) is ok, then you can only customize its subcomponents that
you want to change. With this Server, you can still customize how does it serve price&quote requests and process orders.

### Customize subcomponents
There are four subcomponents, which are:
* [Chain Querier](./sdk.md#interface-chainquerier)
* [Liquidity Provider](./sdk.md#interface-liquidityprovider)
* [Amount Calculator](./sdk.md#interface-amountcalculator)
* [Request Signer](./sdk.md#interface-requestsigner)

Click on each to see the interface detail.

### Customize order processing
Requirements:
* [Validate Quote](./sdk.md#func-server-validatequote) for each order before processing it 
* double check any information comes from RFQ Server before transferring out token, especially for User has deposited
* stop processing order if there is any unhandled error
* as a consequence of RFQ Server help MMs maintain orders, an MM should timely update the status of an order through
  [UpdateOrders API](./sdk.md#func-client-updateorders). The most important times to call this api are:
  1. update to `OrderStatus_STATUS_MM_REJECTED` at any appropriate time when the MM thinks he should reject this order before any token transfer on dst chain.
  2. update to `OrderStatus_STATUS_MM_DST_EXECUTED` when the MM has sent a tx on dst chain for transferring token to the User, regardless of whether it's mined or not and its execution status.
  3. update to `OrderStatus_STATUS_MM_SRC_EXECUTED` when the MM has sent a tx on src chain for releasing token to himself, regardless of whether it's mined or not and its execution status.
  4. specially, update order from `OrderStatus_STATUS_SRC_DEPOSITED` to `OrderStatus_STATUS_REFUND_INITIATED` when it's a 
same chain swap `quote.GetSrcChainId() == quote.GetDstChainId()` and `quote.DstDeadline` has passed.


Example:
```go
// server := NewServer(...)
if server.Ctl == nil {
    log.Panicln("nil control channel")
}
ticker := time.NewTicker(time.Duration(server.Config.ProcessPeriod) * time.Second)
for {
    select {
    case <-ticker.C:
    // check component's functionality
    if server.LiquidityProvider.IsPaused() {
        server.StopProcessing("liquidity provider is paused in some reason")
        continue
    }
    resp, err := server.RfqClient.PendingOrders(context.Background(), &rfqproto.PendingOrdersRequest{})
    if err != nil {
        // handler err
        continue
    }
	// your customized processOrders
    // processOrders(server, resp.Orders)
    case <-server.Ctl:
        return
    }
}
```

### Customize request serving
If you want to customize request serving, you'd better package Server into a new structure. So that you can implement new
Price and Quote, and share the subcomponents of Server at the same time.

Example
```go
type YourMMApp struct {
    Server *rfqmm.Server
}
func (mm *YourMMApp) Price(ctx context.Context, request *proto.PriceRequest) (response *proto.PriceResponse, err error) {
    // todo, remove panic() and write your own implementation
    panic()
}
func (mm *YourMMApp) Quote(ctx context.Context, request *proto.QuoteRequest) (response *proto.QuoteResponse, err error) {
    // todo, remove panic() and write your own implementation
    panic()
}
func (mm *YourMMApp) Serve(ops ...grpc.ServerOption) {
    port := mm.Server.Config.PortListenOn
    listener, err := net.Listen("tcp", fmt.Sprintf("localhost:%d", port))
    if err != nil {
        panic(err)
    }
    grpcServer := grpc.NewServer(ops...)
    proto.RegisterApiServer(grpcServer, mm)
    grpcServer.Serve(listener)
}
```

## Support

### Information

#### Server
RFQ Server URL: `cbridge-v2-test.celer.network:9094`

#### Contract Address
RFQ contract
* Goerli: [0xfF3cf572D591391935EF45F379C6D83182Feff5C](https://goerli.etherscan.io/address/0xfF3cf572D591391935EF45F379C6D83182Feff5C)
* BSC Testnet: [0xdc23f4F3dFA283eD56F84d40F5AdE69dF16ec32E](https://testnet.bscscan.com/address/0xdc23f4F3dFA283eD56F84d40F5AdE69dF16ec32E)

Wrapped native contract
* Goerli: [0xB4FBF271143F4FBf7B91A5ded31805e42b2208d6](https://goerli.etherscan.io/address/0xB4FBF271143F4FBf7B91A5ded31805e42b2208d6)

### Debugging
