.PHONY: protos

protos:
	# protoc -I rates/ rates/rate.proto --go-grpc_out=grpc:rates/rate
	protoc --go_out=. --go-grpc_out=. rates/rate.proto
gen:
	protoc --proto_path=bookshop bookshop/*.proto --go_out=. --go-grpc_out=.