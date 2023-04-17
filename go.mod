module github.com/greyvar/datlint

go 1.18

replace github.com/greyvar/datlib => ../datlib/

require (
	github.com/greyvar/datlib v0.0.0-20220723191212-08d2466064bf
	github.com/sirupsen/logrus v1.9.0
)

require (
	golang.org/x/sys v0.0.0-20220715151400-c0bba94af5f8 // indirect
	gopkg.in/yaml.v2 v2.4.0 // indirect
)
