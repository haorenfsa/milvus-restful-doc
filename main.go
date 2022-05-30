package main

import (
	"github.com/milvus.io/milvus-restful-doc/internal/proto/common"
	"github.com/milvus.io/milvus-restful-doc/internal/proto/schema"
	"github.com/milvus.io/milvus-restful-doc/internal/proto/server"
)

// @title           Milvus RESTful API
// @version         v2.1
// @description     The RESTful API Document for MilvusV2.
// @host      milvus-proxy:8080
// @BasePath  /api/v1

// @license.name  Apache 2.0
// @license.url   http://www.apache.org/licenses/LICENSE-2.0.html
func main() {

}

type CreateCollectionRequest struct {
	// Not useful for now
	Base *common.MsgBase `protobuf:"bytes,1,opt,name=base,proto3" json:"base,omitempty"`
	// Not useful for now
	DbName string `protobuf:"bytes,2,opt,name=db_name,json=dbName,proto3" json:"db_name,omitempty"`
	// The unique collection name in milvus.(Required)
	CollectionName string `protobuf:"bytes,3,opt,name=collection_name,json=collectionName,proto3" json:"collection_name,omitempty"`
	// The serialized `schema.CollectionSchema`(Required)
	Schema schema.CollectionSchema `protobuf:"bytes,4,opt,name=schema,proto3" json:"schema,omitempty"`
	// Once set, no modification is allowed (Optional)
	// https://github.com/milvus-io/milvus/issues/6690
	ShardsNum int32 `protobuf:"varint,5,opt,name=shards_num,json=shardsNum,proto3" json:"shards_num,omitempty"`
	// The consistency level that the collection used, modification is not supported now.
	ConsistencyLevel common.ConsistencyLevel `protobuf:"varint,6,opt,name=consistency_level,json=consistencyLevel,proto3,enum=milvus.proto.common.ConsistencyLevel" json:"consistency_level,omitempty"`
}

// CreateCollection
// @Summary		CreateCollection
// @Description Create a collection
// @Tags		Collection
// @Param       req		body		CreateCollectionRequest	true	"CreateCollectionRequest"
// @Success 	200 	{object}	common.Status	"Request accepted"
// @Failure 	400 	{object}	common.Status	"Bad request format"
// @Router 		/collection [post]
func CreateCollection(req *CreateCollectionRequest) *common.Status {
	return nil
}

// DescribeCollection
// @Summary		DescribeCollection
// @Description Describe a collection
// @Tags		Collection
// @Param       req		body		server.DescribeCollectionRequest	true	"DescribeCollectionRequest"
// @Success 	200 	{object}	server.DescribeCollectionResponse	"Request accepted"
// @Failure 	400 	{object}	common.Status	"Bad request format"
// @Router 		/collection [get]
func DescribeCollection(req *server.DescribeCollectionRequest) *server.DescribeCollectionResponse {
	return nil
}

// CreateCollection
// @Summary		DropCollection
// @Description Drop a collection
// @Tags		Collection
// @Param       req		body		server.DropCollectionRequest	true	"DropCollectionRequest"
// @Success 	200 	{object}	common.Status	"Request accepted"
// @Failure 	400 	{object}	common.Status	"Bad request format"
// @Router 		/collection [delete]
func DropCollection(req *server.DropCollectionRequest) *common.Status {
	return nil
}

// HasCollection
// @Summary		HasCollection
// @Description Get if a collection's existence
// @Tags		Collection
// @Param       req		body		server.HasCollectionRequest	true	"HasCollectionRequest"
// @Success 	200 	{object}	server.BoolResponse	"Request accepted"
// @Failure 	400 	{object}	common.Status	"Bad request format"
// @Router 		/collection/existence [get]
func HasCollection(req *server.HasCollectionRequest) *server.BoolResponse {
	return nil
}

// LoadCollection
// @Summary		LoadCollection
// @Description Load a collection for search
// @Tags		Collection
// @Param       req		body		server.HasCollectionRequest	true	"LoadCollectionRequest"
// @Success 	200 	{object}	common.Status	"Request accepted"
// @Failure 	400 	{object}	common.Status	"Bad request format"
// @Router 		/collection/load [post]
func LoadCollection(req *server.LoadCollectionRequest) *common.Status {
	return nil
}

// ReleaseCollection
// @Summary		ReleaseCollection
// @Description Release a collection loaded before
// @Tags		Collection
// @Param       req		body		server.HasCollectionRequest	true	"ReleaseCollectionRequest"
// @Success 	200 	{object}	common.Status	"Request accepted"
// @Failure 	400 	{object}	common.Status	"Bad request format"
// @Router 		/collection/load [delete]
func ReleaseCollection(req *server.ReleaseCollectionRequest) *common.Status {
	return nil
}

// GetCollectionStatistics
// @Summary		GetCollectionStatistics
// @Description Get a collection's statistics
// @Tags		Collection
// @Param       req		body		server.GetCollectionStatisticsRequest	true	"GetCollectionStatisticsRequest"
// @Success 	200 	{object}	server.GetCollectionStatisticsResponse	"Request accepted"
// @Failure 	400 	{object}	common.Status	"Bad request format"
// @Router 		/collection/statistics [get]
func GetCollectionStatistics(req *server.GetCollectionStatisticsRequest) *server.GetCollectionStatisticsResponse {
	return nil
}

// ShowCollections
// @Summary		ShowCollections
// @Description Show all collections
// @Tags		Collection
// @Param       req		body		server.ShowCollectionsRequest	true	"ShowCollectionsRequest"
// @Success 	200 	{object}	server.ShowCollectionsResponse	"Request accepted"
// @Failure 	400 	{object}	common.Status	"Bad request format"
// @Router 		/collections [get]
func ShowCollections(req *server.ShowCollectionsRequest) *server.ShowCollectionsResponse {
	return nil
}

// CreatePartition
// @Summary		CreatePartition
// @Description Create a partition
// @Tags		Partition
// @Param       req		body		server.CreatePartitionRequest	true	"CreatePartitionRequest"
// @Success 	200 	{object}	common.Status	"Request accepted"
// @Failure 	400 	{object}	common.Status	"Bad request format"
// @Router 		/partition [post]
func CreatePartition(req *server.CreatePartitionRequest) *common.Status {
	return nil
}

// DropPartition
// @Summary		DropPartition
// @Description Delete a partition
// @Tags		Partition
// @Param       req		body		server.DropPartitionRequest	true	"DropPartitionRequest"
// @Success 	200 	{object}	common.Status	"Request accepted"
// @Failure 	400 	{object}	common.Status	"Bad request format"
// @Router 		/partition [delete]
func DropPartition(req *server.DropPartitionRequest) *common.Status {
	return nil
}

// HasPartition
// @Summary		HasPartition
// @Description Get if a partition exists
// @Tags		Partition
// @Param       req		body		server.HasPartitionRequest	true	"HasPartitionRequest"
// @Success 	200 	{object}	server.BoolResponse	"Request accepted"
// @Failure 	400 	{object}	common.Status	"Bad request format"
// @Router 		/partition/existence [get]
func HasPartition(req *server.HasPartitionRequest) *server.BoolResponse {
	return nil
}

// LoadPartitions
// @Summary		LoadPartitions
// @Description Load a group of paritions for search
// @Tags		Partition
// @Param       req		body		server.LoadPartitionsRequest	true	"LoadPartitionsRequest"
// @Success 	200 	{object}	common.Status	"Request accepted"
// @Failure 	400 	{object}	common.Status	"Bad request format"
// @Router 		/partitions/load [post]
func LoadPartitions(req *server.LoadPartitionsRequest) *common.Status {
	return nil
}

// ReleasePartitions
// @Summary		ReleasePartitions
// @Description Release a group of loaded paritions
// @Tags		Partition
// @Param       req		body		server.ReleasePartitionsRequest	true	"ReleasePartitionsRequest"
// @Success 	200 	{object}	common.Status	"Request accepted"
// @Failure 	400 	{object}	common.Status	"Bad request format"
// @Router 		/partitions/load [delete]
func ReleasePartitions(req *server.ReleasePartitionsRequest) *common.Status {
	return nil
}

// GetPartitionStatistics
// @Summary		GetPartitionStatistics
// @Description Get a partition's statistics
// @Tags		Partition
// @Param       req		body		server.GetPartitionStatisticsRequest	true	"GetPartitionStatisticsRequest"
// @Success 	200 	{object}	server.GetPartitionStatisticsResponse	"Request accepted"
// @Failure 	400 	{object}	common.Status	"Bad request format"
// @Router 		/partition/statistics [get]
func GetPartitionStatistics(req *server.GetPartitionStatisticsRequest) *server.GetPartitionStatisticsResponse {
	return nil
}

// ShowPartitions
// @Summary		ShowPartitions
// @Description Show all partitions
// @Tags		Partition
// @Param       req		body		server.ShowPartitionsRequest	true	"ShowPartitionsRequest"
// @Success 	200 	{object}	server.ShowPartitionsResponse	"Request accepted"
// @Failure 	400 	{object}	common.Status	"Bad request format"
// @Router 		/partitions [get]
func ShowPartitions(req *server.ShowPartitionsRequest) *server.ShowPartitionsResponse {
	return nil
}

// CreateAlias
// @Summary		CreateAlias
// @Description Create an alias for a collection name
// @Tags		Alias
// @Param       req		body		server.CreateAliasRequest	true	"CreateAliasRequest"
// @Success 	200 	{object}	common.Status	"Request accepted"
// @Failure 	400 	{object}	common.Status	"Bad request format"
// @Router 		/alias [post]
func CreateAlias(req *server.CreateAliasRequest) *common.Status {
	return nil
}

// DropAlias
// @Summary		DropAlias
// @Description Delete an Alias
// @Tags		Alias
// @Param       req		body		server.DropAliasRequest	true	"DropAliasRequest"
// @Success 	200 	{object}	common.Status	"Request accepted"
// @Failure 	400 	{object}	common.Status	"Bad request format"
// @Router 		/alias [delete]
func DropAlias(req *server.DropAliasRequest) *common.Status {
	return nil
}

// AlterAlias
// @Summary		AlterAlias
// @Description Alter an alias
// @Tags		Alias
// @Param       req		body		server.AlterAliasRequest	true	"AlterAliasRequest"
// @Success 	200 	{object}	common.Status	"Request accepted"
// @Failure 	400 	{object}	common.Status	"Bad request format"
// @Router 		/alias [patch]
func AlterAlias(req *server.AlterAliasRequest) *common.Status {
	return nil
}

// CreateIndex
// @Summary		CreateIndex
// @Description Create an Index
// @Tags		Index
// @Param       req		body		server.CreateIndexRequest	true	"CreateIndexRequest"
// @Success 	200 	{object}	common.Status	"Request accepted"
// @Failure 	400 	{object}	common.Status	"Bad request format"
// @Router 		/index [post]
func CreateIndex(req *server.CreateIndexRequest) *common.Status {
	return nil
}

// DescribeIndex
// @Summary		DescribeIndex
// @Description Describe an index
// @Tags		Index
// @Param       req		body		server.DescribeIndexRequest	true	"DescribeIndexRequest"
// @Success 	200 	{object}	server.DescribeIndexResponse	"Request accepted"
// @Failure 	400 	{object}	common.Status	"Bad request format"
// @Router 		/index [get]
func DescribeIndex(req *server.DescribeIndexRequest) *server.DescribeIndexResponse {
	return nil
}

// DropIndex
// @Summary		DropIndex
// @Description Drop an index
// @Tags		Index
// @Param       req		body		server.DropIndexRequest	true	"DropIndexRequest"
// @Success 	200 	{object}	common.Status	"Request accepted"
// @Failure 	400 	{object}	common.Status	"Bad request format"
// @Router 		/index [delete]
func DropIndex(req *server.DropIndexRequest) *common.Status {
	return nil
}

// GetIndexState
// @Summary		GetIndexState
// @Description Get the state of an index
// @Tags		Index
// @Param       req		body		server.GetIndexStateRequest	true	"GetIndexStateRequest"
// @Success 	200 	{object}	server.GetIndexStateResponse	"Request accepted"
// @Failure 	400 	{object}	common.Status	"Bad request format"
// @Router 		/index/state [get]
func GetIndexState(req *server.GetIndexStateRequest) *server.GetIndexStateResponse {
	return nil
}

// GetIndexBuildProgress
// @Summary		GetIndexBuildProgress
// @Description Get the build progress of an index
// @Tags		Index
// @Param       req		body		server.GetIndexBuildProgressRequest	true	"GetIndexBuildProgressRequest"
// @Success 	200 	{object}	server.GetIndexBuildProgressResponse	"Request accepted"
// @Failure 	400 	{object}	common.Status	"Bad request format"
// @Router 		/index/progress [get]
func GetIndexBuildProgress(req *server.GetIndexBuildProgressRequest) *server.GetIndexBuildProgressResponse {
	return nil
}

// Insert
// @Summary		Insert
// @Description Insert rows of data entities into a collection
// @Tags		Entity
// @Param       req		body		server.InsertRequest	true	"InsertRequest"
// @Success 	200 	{object}	common.Status	"Request accepted"
// @Failure 	400 	{object}	common.Status	"Bad request format"
// @Router 		/entities [post]
func Insert(req *server.InsertRequest) *common.Status {
	return nil
}

// Delete
// @Summary		Delete
// @Description Delete rows of data entities from a collection by given expresssion
// @Tags		Entity
// @Param       req		body		server.DeleteRequest	true	"DeleteRequest"
// @Success 	200 	{object}	common.Status	"Request accepted"
// @Failure 	400 	{object}	common.Status	"Bad request format"
// @Router 		/entities [delete]
func Delete(req *server.DeleteRequest) *common.Status {
	return nil
}

// Search
// @Summary		Search
// @Description Do a k nearest neighbors search with bool expression
// @Tags		Entity
// @Param       req		body		server.SearchRequest	true	"SearchRequest"
// @Success 	200 	{object}	server.SearchResults	"Request accepted"
// @Failure 	400 	{object}	common.Status	"Bad request format"
// @Router 		/search [post]
func Search(req *server.SearchRequest) *server.SearchResults {
	return nil
}

// Query
// @Summary		Query
// @Description do a explicit record query by given expression. For example when you want to query by primary key.
// @Tags		Entity
// @Param       req		body		server.QueryRequest	true	"QueryRequest"
// @Success 	200 	{object}	server.QueryResults	"Request accepted"
// @Failure 	400 	{object}	common.Status	"Bad request format"
// @Router 		/query [post]
func Query(req *server.QueryRequest) *server.QueryResults {
	return nil
}

// Flush
// @Summary		Flush
// @Description Flush a collection's data to disk. Milvus's data will be auto flushed. Flush is only required when you want to get up to date entities numbers in statistics due to some internal mechanism. It will be removed in the future.
// @Tags		Entity
// @Param       req		body		server.FlushRequest	true	"FlushRequest"
// @Success 	200 	{object}	server.FlushResponse	"Request accepted"
// @Failure 	400 	{object}	common.Status	"Bad request format"
// @Router 		/persist [post]
func Flush(req *server.FlushRequest) *server.FlushResponse {
	return nil
}

// CalcDistance
// @Summary		CalcDistance
// @Description Calculate distance between specified vectors
// @Tags		Entity
// @Param       req		body		server.CalcDistanceRequest	true	"CalcDistanceRequest"
// @Success 	200 	{object}	server.CalcDistanceResults	"Request accepted"
// @Failure 	400 	{object}	common.Status	"Bad request format"
// @Router 		/distance [post]
func CalcDistance(req *server.CalcDistanceRequest) *server.CalcDistanceResults {
	return nil
}

// GetFlushState
// @Summary		GetFlushState
// @Description Get the flush state of multiple segments
// @Tags		Entity
// @Param       req		body		server.GetFlushStateRequest	true	"GetFlushStateRequest"
// @Success 	200 	{object}	server.GetFlushStateResponse	"Request accepted"
// @Failure 	400 	{object}	common.Status	"Bad request format"
// @Router 		/persist/state [get]
func GetFlushState(req *server.GetFlushStateRequest) *server.GetFlushStateResponse {
	return nil
}

// GetPersistentSegmentInfo
// @Summary		GetPersistentSegmentInfo
// @Description Returns sealed segments's information of a collection
// @Tags		Entity
// @Param       req		body		server.GetPersistentSegmentInfoRequest	true	"GetPersistentSegmentInfoRequest"
// @Success 	200 	{object}	server.GetPersistentSegmentInfoResponse	"Request accepted"
// @Failure 	400 	{object}	common.Status	"Bad request format"
// @Router 		/persist/segment-info [get]
func GetPersistentSegmentInfo(req *server.GetPersistentSegmentInfoRequest) *server.GetPersistentSegmentInfoResponse {
	return nil
}

// GetQuerySegmentInfo
// @Summary		GetQuerySegmentInfo
// @Description Returns growing segments's information of a collection
// @Tags		Entity
// @Param       req		body		server.GetQuerySegmentInfoRequest	true	"GetQuerySegmentInfoRequest"
// @Success 	200 	{object}	server.GetQuerySegmentInfoResponse	"Request accepted"
// @Failure 	400 	{object}	common.Status	"Bad request format"
// @Router 		/query-segment-info [get]
func GetQuerySegmentInfo(req *server.GetQuerySegmentInfoRequest) *server.GetQuerySegmentInfoResponse {
	return nil
}

// GetReplicasRequest
// @Summary		GetReplicas
// @Description GetReplicas info of a collection
// @Tags		Ops
// @Param       req		body		server.GetReplicasRequest	true	"GetReplicasRequest"
// @Success 	200 	{object}	server.GetReplicasResponse	"Request accepted"
// @Failure 	400 	{object}	common.Status	"Bad request format"
// @Router 		/replicas [get]
func GetReplicas(req *server.GetReplicasRequest) *server.GetReplicasResponse {
	return nil
}

// GetMetrics
// @Summary		GetMetrics
// @Description Get metrics
// @Tags		Metrics
// @Param       req		body		server.GetMetricsRequest	true	"GetMetricsRequest"
// @Success 	200 	{object}	server.GetMetricsResponse	"Request accepted"
// @Failure 	400 	{object}	common.Status	"Bad request format"
// @Router 		/metrics [get]
func GetMetrics(req *server.GetMetricsRequest) *server.GetMetricsResponse {
	return nil
}

// LoadBalance
// @Summary		LoadBalance
// @Description Do a load balancing operation between query nodes
// @Tags		Ops
// @Param       req		body		server.LoadBalanceRequest	true	"LoadBalanceRequest"
// @Success 	200 	{object}	common.Status	"Request accepted"
// @Failure 	400 	{object}	common.Status	"Bad request format"
// @Router 		/load-balance [post]
func LoadBalance(req *server.LoadBalanceRequest) *common.Status {
	return nil
}

// ManualCompaction
// @Summary		ManualCompaction
// @Description Do a mannual compaction
// @Tags		Ops
// @Param       req		body		server.ManualCompactionRequest	true	"ManualCompactionRequest"
// @Success 	200 	{object}	server.ManualCompactionResponse	"Request accepted"
// @Failure 	400 	{object}	common.Status	"Bad request format"
// @Router 		/compaction [post]
func ManualCompaction(req *server.ManualCompactionRequest) *server.ManualCompactionResponse {
	return nil
}

// GetCompactionState
// @Summary		GetCompactionState
// @Description Get the state of a compaction
// @Tags		Ops
// @Param       req		body		server.GetCompactionStateRequest	true	"GetCompactionStateRequest"
// @Success 	200 	{object}	server.GetCompactionStateResponse	"Request accepted"
// @Failure 	400 	{object}	common.Status	"Bad request format"
// @Router 		/compaction/state [get]
func GetCompactionState(req *server.GetCompactionStateRequest) *server.GetCompactionStateResponse {
	return nil
}

// GetCompactionPlans
// @Summary		GetCompactionPlans
// @Description Get the plans of a compaction
// @Tags		Ops
// @Param       req		body		server.GetCompactionPlansRequest	true	"GetCompactionPlansRequest"
// @Success 	200 	{object}	server.GetCompactionPlansResponse	"Request accepted"
// @Failure 	400 	{object}	common.Status	"Bad request format"
// @Router 		/compaction/plans [get]
func GetCompactionPlans(req *server.GetCompactionPlansRequest) *server.GetCompactionPlansResponse {
	return nil
}

// Import
// @Summary		Import
// @Description Import data files(json, numpy, etc.) on MinIO/S3 storage, read and parse them into sealed segments
// @Tags		Import
// @Param       req		body		server.ImportRequest	true	"ImportRequest"
// @Success 	200 	{object}	server.ImportResponse	"Request accepted"
// @Failure 	400 	{object}	common.Status	"Bad request format"
// @Router 		/import [post]
func Import(req *server.ImportRequest) *server.ImportResponse {
	return nil
}

// GetImportState
// @Summary		GetImportState
// @Description Get the state of a import task
// @Tags		Import
// @Param       req		body		server.GetImportStateRequest	true	"GetImportStateRequest"
// @Success 	200 	{object}	server.GetImportStateResponse	"Request accepted"
// @Failure 	400 	{object}	common.Status	"Bad request format"
// @Router 		/compaction/state [get]
func GetImportState(req *server.GetImportStateRequest) *server.GetImportStateResponse {
	return nil
}

// ListImportTasks
// @Summary		ListImportTasks
// @Description List all import tasks
// @Tags		Import
// @Param       req		body		server.ListImportTasksRequest	true	"ListImportTasksRequest"
// @Success 	200 	{object}	server.ListImportTasksResponse	"Request accepted"
// @Failure 	400 	{object}	common.Status	"Bad request format"
// @Router 		/compaction/tasks [get]
func ListImportTasks(req *server.ListImportTasksRequest) *server.ListImportTasksResponse {
	return nil
}

// CreateCredential
// @Summary		CreateCredential
// @Description Create a new user and password
// @Tags		Credential
// @Param       req		body		server.CreateCredentialRequest	true	"CreateCredentialRequest"
// @Success 	200 	{object}	common.Status	"Request accepted"
// @Failure 	400 	{object}	common.Status	"Bad request format"
// @Router 		/credential [post]
func CreateCredential(req *server.CreateCredentialRequest) *common.Status {
	return nil
}

// UpdateCredential
// @Summary		UpdateCredential
// @Description Update password for a user
// @Tags		Credential
// @Param       req		body		server.UpdateCredentialRequest	true	"UpdateCredential"
// @Success 	200 	{object}	common.Status	"Request accepted"
// @Failure 	400 	{object}	common.Status	"Bad request format"
// @Router 		/credential [patch]
func UpdateCredential(req *server.UpdateCredentialRequest) *common.Status {
	return nil
}

// DeleteCredential
// @Summary		DeleteCredential
// @Description Delete a Credential
// @Tags		Credential
// @Param       req		body		server.DeleteCredentialRequest	true	"DeleteCredentialRequest"
// @Success 	200 	{object}	common.Status	"Request accepted"
// @Failure 	400 	{object}	common.Status	"Bad request format"
// @Router 		/credential [delete]
func DeleteCredential(req *server.DeleteCredentialRequest) *common.Status {
	return nil
}

// ListCredUsers
// @Summary		ListCredUsers
// @Description List all users
// @Tags		Credential
// @Param       req		body		server.ListCredUsersRequest	true	"ListCredUsersRequest"
// @Success 	200 	{object}	server.ListCredUsersResponse	"Request accepted"
// @Failure 	400 	{object}	common.Status	"Bad request format"
// @Router 		/credential/users [get]
func ListCredUsers(req *server.ListCredUsersRequest) *server.ListCredUsersResponse {
	return nil
}
