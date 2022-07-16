raftnode.pb.go: raftnode.proto
	protoc raftnode.proto --go_out=plugins=grpc:. --go_opt=paths=source_relative

force:
	rm -f raftnode.pb.go
	make raftnode.pb.go