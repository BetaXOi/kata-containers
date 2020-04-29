/*
 * Cloud Hypervisor API
 *
 * Local HTTP based API for managing and inspecting a cloud-hypervisor virtual machine.
 *
 * API version: 0.3.0
 * Generated by: OpenAPI Generator (https://openapi-generator.tech)
 */

package openapi
// PmemConfig struct for PmemConfig
type PmemConfig struct {
	File string `json:"file"`
	Size int64 `json:"size"`
	Iommu bool `json:"iommu,omitempty"`
	Mergeable bool `json:"mergeable,omitempty"`
	DiscardWrites bool `json:"discard_writes,omitempty"`
}
