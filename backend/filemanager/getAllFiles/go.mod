module filesharing/getAllFiles

go 1.24.6

replace dikobra3/mongoApi => ../../mongo/mongoApi

replace dikobra3/utils => ../../utils

replace github.com/big-larry/mgo => ../../mongo/mgo@v1.0.0

require (
	dikobra3/mongoApi v0.0.0-00010101000000-000000000000
	dikobra3/utils v0.0.0-00010101000000-000000000000
)

require (
	github.com/big-larry/mgo v1.0.0
	github.com/big-larry/suckhttp v0.0.0-20250813123807-92fba52c65cb
	github.com/okonma-violet/services/logs/logger v0.0.0-20250813191921-b2f7a4bb2ca5
	github.com/okonma-violet/services/universalservice_nonepoll v0.0.0-20250813141351-a2866ef3e748
)

require (
	github.com/big-larry/suckutils v0.0.0-20250812153223-6165c58dc43b // indirect
	github.com/okonma-violet/confdecoder v0.0.0-20230926094403-7e3eab7eff29 // indirect
	github.com/okonma-violet/dynamicworkerspool v0.0.0-20240317115954-810a6361f715 // indirect
	github.com/okonma-violet/services/basicmessage v0.0.0-20250813141351-a2866ef3e748 // indirect
	github.com/okonma-violet/services/basicmessage/basicmessagetypes v0.0.0-20250813141351-a2866ef3e748 // indirect
	github.com/okonma-violet/services/connector_nonepoll v0.0.0-20250813141351-a2866ef3e748 // indirect
	github.com/okonma-violet/services/logs/encode v0.0.0-20250813191921-b2f7a4bb2ca5 // indirect
	github.com/okonma-violet/services/types/configuratortypes v0.0.0-20250813191921-b2f7a4bb2ca5 // indirect
	github.com/okonma-violet/services/types/netprotocol v0.0.0-20250813191921-b2f7a4bb2ca5 // indirect
)
