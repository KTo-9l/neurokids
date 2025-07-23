module course/createCourse

go 1.24.4

replace course/models => ../models

replace dikobra3/mongoApi => ../../mongo/mongoApi

replace dikobra3/utils => ../../utils

replace github.com/big-larry/mgo => ../../mongo/mgo@v1.0.0

replace github.com/okonma-violet/services => /home/andrey/go/pkg/mod/github.com/okonma-violet/services@v0.0.0-20250303113135-3d036f9a188f

replace github.com/big-larry/suckutils => /home/andrey/go/pkg/mod/github.com/big-larry/suckutils@v0.0.0-20231029230114-645d5d858694

require (
	course/models v0.0.0-00010101000000-000000000000
	dikobra3/mongoApi v0.0.0-00010101000000-000000000000
	dikobra3/utils v0.0.0-00010101000000-000000000000
	github.com/big-larry/mgo v1.0.0
	github.com/big-larry/suckhttp v0.0.0-20250417113412-b2d284b11f53
	github.com/okonma-violet/services v0.0.0-00010101000000-000000000000
)

require (
	github.com/big-larry/suckutils v0.0.0-20231029230114-645d5d858694 // indirect
	github.com/okonma-violet/confdecoder v0.0.0-20230926094403-7e3eab7eff29 // indirect
	github.com/okonma-violet/dynamicworkerspool v0.0.0-20240317115954-810a6361f715 // indirect
)

