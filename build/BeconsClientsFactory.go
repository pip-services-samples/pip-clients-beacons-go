package build

import (
	bclients "github.com/pip-services-samples/pip-clients-beacons-go/clients/version1"
	cref "github.com/pip-services3-go/pip-services3-commons-go/refer"
	cbuild "github.com/pip-services3-go/pip-services3-components-go/build"
)

type BeaconsClientFactory struct {
	cbuild.Factory
	NullClientDescriptor    *cref.Descriptor
	DirectClientDescriptor  *cref.Descriptor
	HttpClientDescriptor    *cref.Descriptor
	GrpcClientDescriptor    *cref.Descriptor
	CmdHttpClientDescriptor *cref.Descriptor
	CmdGrpcClientDescriptor *cref.Descriptor
	MemoryClientDescriptor  *cref.Descriptor
}

func NewBeaconsClientFactory() *BeaconsClientFactory {

	bcf := BeaconsClientFactory{}
	bcf.Factory = *cbuild.NewFactory()

	bcf.NullClientDescriptor = cref.NewDescriptor("pip-services-beacons", "client", "null", "*", "1.0")
	bcf.DirectClientDescriptor = cref.NewDescriptor("pip-services-beacons", "client", "direct", "*", "1.0")
	bcf.HttpClientDescriptor = cref.NewDescriptor("pip-services-beacons", "client", "http", "*", "1.0")
	bcf.GrpcClientDescriptor = cref.NewDescriptor("pip-services-beacons", "client", "grpc", "*", "1.0")
	bcf.CmdHttpClientDescriptor = cref.NewDescriptor("pip-services-beacons", "client", "commandable-http", "*", "1.0")
	bcf.CmdGrpcClientDescriptor = cref.NewDescriptor("pip-services-beacons", "client", "commandable-grpc", "*", "1.0")
	bcf.MemoryClientDescriptor = cref.NewDescriptor("pip-services-beacons", "client", "memory", "*", "1.0")

	bcf.RegisterType(bcf.NullClientDescriptor, bclients.NewBeaconsNullClientV1)
	bcf.RegisterType(bcf.DirectClientDescriptor, bclients.NewBeaconsDirectClientV1)
	bcf.RegisterType(bcf.HttpClientDescriptor, bclients.NewBeaconsRestClientV1)
	bcf.RegisterType(bcf.GrpcClientDescriptor, bclients.NewBeaconsGrpcClientV1)
	bcf.RegisterType(bcf.CmdGrpcClientDescriptor, bclients.NewBeaconsCommandableGrpcClientV1)
	bcf.RegisterType(bcf.CmdHttpClientDescriptor, bclients.NewBeaconsCommandableHttpClientV1)
	bcf.RegisterType(bcf.MemoryClientDescriptor, bclients.NewBeaconsMemoryClientV1)

	return &bcf
}
