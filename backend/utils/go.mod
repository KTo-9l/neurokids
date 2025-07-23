module dikobra3/utils

go 1.24.4

replace dikobra3/mongoApi => ../mongo/mongoApi

replace github.com/big-larry/mgo => ../mongo/mgo@v1.0.0

require dikobra3/mongoApi v0.0.0-00010101000000-000000000000

require github.com/big-larry/mgo v1.0.0
