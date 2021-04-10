all:
	go build
	cd examples/profile && go build -o ../../bin/profile
	cd examples/teams && go build -o ../../bin/teams