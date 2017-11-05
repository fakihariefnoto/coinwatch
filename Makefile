usage: FORCE
	# See targets in Makefile (e.g. "buildlet.darwin-amd64")
	exit 1

FORCE:

coin: FORCE
	@echo " >> building binaries..."
	@go build -o coinwatch
	@echo " >> coinwatch app builded."
	@echo " >> building config binaries..."
	@go build -o configcoin cmd/config/config.go
	@echo " >> config app builded."
	@echo " >> call coinwatch app ...."
	@./coinwatch
	@echo " >> coinwatch running.. <<"
