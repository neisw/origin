// Code generated by applyconfiguration-gen. DO NOT EDIT.

package v1

// IBMCloudPlatformSpecApplyConfiguration represents a declarative configuration of the IBMCloudPlatformSpec type for use
// with apply.
type IBMCloudPlatformSpecApplyConfiguration struct {
	ServiceEndpoints []IBMCloudServiceEndpointApplyConfiguration `json:"serviceEndpoints,omitempty"`
}

// IBMCloudPlatformSpecApplyConfiguration constructs a declarative configuration of the IBMCloudPlatformSpec type for use with
// apply.
func IBMCloudPlatformSpec() *IBMCloudPlatformSpecApplyConfiguration {
	return &IBMCloudPlatformSpecApplyConfiguration{}
}

// WithServiceEndpoints adds the given value to the ServiceEndpoints field in the declarative configuration
// and returns the receiver, so that objects can be build by chaining "With" function invocations.
// If called multiple times, values provided by each call will be appended to the ServiceEndpoints field.
func (b *IBMCloudPlatformSpecApplyConfiguration) WithServiceEndpoints(values ...*IBMCloudServiceEndpointApplyConfiguration) *IBMCloudPlatformSpecApplyConfiguration {
	for i := range values {
		if values[i] == nil {
			panic("nil value passed to WithServiceEndpoints")
		}
		b.ServiceEndpoints = append(b.ServiceEndpoints, *values[i])
	}
	return b
}
