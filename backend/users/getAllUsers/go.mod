module users/getAllUsers

go 1.24.6

replace dikobra3/utils => ../../utils

replace authorization/database => ../../userDatabase

replace authorization/auth_helpers => ../../auth_helpers

require (
	authorization/auth_helpers v0.0.0-00010101000000-000000000000
	authorization/database v0.0.0-00010101000000-000000000000
	dikobra3/utils v0.0.0-00010101000000-000000000000
	github.com/big-larry/suckhttp v0.0.0-20250813123807-92fba52c65cb
	github.com/okonma-violet/services/logs/logger v0.0.0-20250813191921-b2f7a4bb2ca5
	github.com/okonma-violet/services/universalservice_nonepoll v0.0.0-20250813141351-a2866ef3e748
)

require (
	github.com/big-larry/suckutils v0.0.0-20250812153223-6165c58dc43b // indirect
	github.com/golang-jwt/jwt v3.2.2+incompatible // indirect
	github.com/jackc/pgpassfile v1.0.0 // indirect
	github.com/jackc/pgservicefile v0.0.0-20240606120523-5a60cdf6a761 // indirect
	github.com/jackc/pgx/v5 v5.7.5 // indirect
	github.com/jackc/puddle/v2 v2.2.2 // indirect
	github.com/okonma-violet/confdecoder v0.0.0-20230926094403-7e3eab7eff29 // indirect
	github.com/okonma-violet/dynamicworkerspool v0.0.0-20240317115954-810a6361f715 // indirect
	github.com/okonma-violet/services/basicmessage v0.0.0-20250813141351-a2866ef3e748 // indirect
	github.com/okonma-violet/services/basicmessage/basicmessagetypes v0.0.0-20250813141351-a2866ef3e748 // indirect
	github.com/okonma-violet/services/connector_nonepoll v0.0.0-20250813141351-a2866ef3e748 // indirect
	github.com/okonma-violet/services/logs/encode v0.0.0-20250813191921-b2f7a4bb2ca5 // indirect
	github.com/okonma-violet/services/types/configuratortypes v0.0.0-20250813191921-b2f7a4bb2ca5 // indirect
	github.com/okonma-violet/services/types/netprotocol v0.0.0-20250813191921-b2f7a4bb2ca5 // indirect
	golang.org/x/crypto v0.41.0 // indirect
	golang.org/x/sync v0.16.0 // indirect
	golang.org/x/text v0.28.0 // indirect
)
