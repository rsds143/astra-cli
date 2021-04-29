/*
 * Astra DevOps API
 *
 * Use this REST API to perform lifecycle actions for DataStax Astra databases.</br> </br> To get started, get your application token from your Astra database. You can then create, terminate, resize, park, and unpark databases using the DevOps API. You cannot park, unpark, or resize serverless databases.  # Authentication  <!-- ReDoc-Inject: <security-definitions> -->
 *
 * API version: 2.0.0
 * Contact: ad-astra@datastax.com
 * Generated by: Swagger Codegen (https://github.com/swagger-api/swagger-codegen.git)
 */
package swagger

// Storage contains the information about how much storage space a cluster has available
type Storage struct {
	// NodeCount for the cluster
	NodeCount int32 `json:"nodeCount"`
	// ReplicationFactor is the number of nodes storing a piece of data
	ReplicationFactor int32 `json:"replicationFactor"`
	// TotalStorage of the cluster in GB
	TotalStorage int32 `json:"totalStorage"`
	// UsedStorage in GB
	UsedStorage int32 `json:"usedStorage,omitempty"`
}