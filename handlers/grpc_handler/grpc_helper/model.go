package grpc_helper

import (
	grpcpublic "x-msa-auth/handlers/grpc_handler/grpc_public"
	corehelper "x-msa-core/helper"
)

type Handler interface {
	Close()

	AddService(key, addr string, group corehelper.GroupsType)
	DeleteService(key string, group corehelper.GroupsType) error

	grpcpublic.AuthServices
}

type MSA interface {
	AddService(key, addr string, group corehelper.GroupsType)
	DeleteService(key string, group corehelper.GroupsType) error

	grpcpublic.AuthServices
}
