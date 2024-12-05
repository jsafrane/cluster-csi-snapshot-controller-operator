// Code generated by applyconfiguration-gen. DO NOT EDIT.

package v1

import (
	v1 "github.com/openshift/api/operator/v1"
)

// HTTPCompressionPolicyApplyConfiguration represents a declarative configuration of the HTTPCompressionPolicy type for use
// with apply.
type HTTPCompressionPolicyApplyConfiguration struct {
	MimeTypes []v1.CompressionMIMEType `json:"mimeTypes,omitempty"`
}

// HTTPCompressionPolicyApplyConfiguration constructs a declarative configuration of the HTTPCompressionPolicy type for use with
// apply.
func HTTPCompressionPolicy() *HTTPCompressionPolicyApplyConfiguration {
	return &HTTPCompressionPolicyApplyConfiguration{}
}

// WithMimeTypes adds the given value to the MimeTypes field in the declarative configuration
// and returns the receiver, so that objects can be build by chaining "With" function invocations.
// If called multiple times, values provided by each call will be appended to the MimeTypes field.
func (b *HTTPCompressionPolicyApplyConfiguration) WithMimeTypes(values ...v1.CompressionMIMEType) *HTTPCompressionPolicyApplyConfiguration {
	for i := range values {
		b.MimeTypes = append(b.MimeTypes, values[i])
	}
	return b
}