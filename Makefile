.PHONY: protos

protos:
	# protoc -I rates/ rates/rate.proto --go-grpc_out=grpc:rates/rate
	# protoc --go_out=. --go-grpc_out=. rates/rate.proto
	protoc --go_out=. --go-grpc_out=. order/order.proto
gen:
	protoc --proto_path=messages/createorder messages/createorder/*.proto --go_out=. --go-grpc_out=.