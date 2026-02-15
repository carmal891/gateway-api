package suite

// Conformance suite shared constants
const (
	// InfrastructureNamespace is the main conformance infra namespace for Gateways and core resources
	InfrastructureNamespace = "gateway-conformance-infra"
	// AppBackendNamespace is the namespace for app backends
	AppBackendNamespace = "gateway-conformance-app-backend"
	// WebBackendNamespace is the namespace for web backends
	WebBackendNamespace = "gateway-conformance-web-backend"
	// MeshNamespace is the namespace for mesh conformance tests
	MeshNamespace = "gateway-conformance-mesh"
	// MeshConsumerNamespace is the namespace for mesh consumer in conformance tests
	MeshConsumerNamespace = "gateway-conformance-mesh-consumer"
	// InfrastructureGatewayName is the default Gateway name in the infra namespace
	InfrastructureGatewayName = "gateway-conformance-infra-test"

	// undefinedKeyword is set in the ConformanceReport "GatewayAPIVersion" and
	// "GatewayAPIChannel" fields in case it's not possible to figure out the actual
	// values in the cluster, due to multiple versions of CRDs installed.
	UndefinedKeyword = "UNDEFINED"
)
