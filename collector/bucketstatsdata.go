// Copyright 2019 Adel Abdelhak.
// Use of this source code is governed by the Apache
// license that can be found in the LICENSE.txt file.

package collector

import (
	"encoding/json"

	p "github.com/prometheus/client_golang/prometheus"
	log "github.com/sirupsen/logrus"
)

// BucketStatsData (/pools/default/buckets/<bucket_name>/stats)
type BucketStatsData struct {
	Op struct {
		Samples struct {
			CouchTotalDiskSize                   []float64 `json:"couch_total_disk_size"`
			CouchDocsFragmentation               []float64 `json:"couch_docs_fragmentation"`
			CouchViewsFragmentation              []float64 `json:"couch_views_fragmentation"`
			HitRatio                             []float64 `json:"hit_ratio"`
			EpCacheMissRate                      []float64 `json:"ep_cache_miss_rate"`
			EpResidentItemsRate                  []float64 `json:"ep_resident_items_rate"`
			VbAvgActiveQueueAge                  []float64 `json:"vb_avg_active_queue_age"`
			VbAvgReplicaQueueAge                 []float64 `json:"vb_avg_replica_queue_age"`
			VbAvgPendingQueueAge                 []float64 `json:"vb_avg_pending_queue_age"`
			VbAvgTotalQueueAge                   []float64 `json:"vb_avg_total_queue_age"`
			VbActiveResidentItemsRatio           []float64 `json:"vb_active_resident_items_ratio"`
			VbReplicaResidentItemsRatio          []float64 `json:"vb_replica_resident_items_ratio"`
			VbPendingResidentItemsRatio          []float64 `json:"vb_pending_resident_items_ratio"`
			AvgDiskUpdateTime                    []float64 `json:"avg_disk_update_time"`
			AvgDiskCommitTime                    []float64 `json:"avg_disk_commit_time"`
			AvgBgWaitTime                        []float64 `json:"avg_bg_wait_time"`
			AvgActiveTimestampDrift              []float64 `json:"avg_active_timestamp_drift"`  // couchbase 5.1.1
			AvgReplicaTimestampDrift             []float64 `json:"avg_replica_timestamp_drift"` // couchbase 5.1.1
			EpDcpViewsIndexesCount               []float64 `json:"ep_dcp_views+indexes_count"`
			EpDcpViewsIndexesItemsRemaining      []float64 `json:"ep_dcp_views+indexes_items_remaining"`
			EpDcpViewsIndexesProducerCount       []float64 `json:"ep_dcp_views+indexes_producer_count"`
			EpDcpViewsIndexesTotalBacklogSize    []float64 `json:"ep_dcp_views+indexes_total_backlog_size"`
			EpDcpViewsIndexesItemsSent           []float64 `json:"ep_dcp_views+indexes_items_sent"`
			EpDcpViewsIndexesTotalBytes          []float64 `json:"ep_dcp_views+indexes_total_bytes"`
			EpDcpViewsIndexesBackoff             []float64 `json:"ep_dcp_views+indexes_backoff"`
			BgWaitCount                          []float64 `json:"bg_wait_count"`
			BgWaitTotal                          []float64 `json:"bg_wait_total"`
			BytesRead                            []float64 `json:"bytes_read"`
			BytesWritten                         []float64 `json:"bytes_written"`
			CasBadval                            []float64 `json:"cas_badval"`
			CasHits                              []float64 `json:"cas_hits"`
			CasMisses                            []float64 `json:"cas_misses"`
			CmdGet                               []float64 `json:"cmd_get"`
			CmdSet                               []float64 `json:"cmd_set"`
			CouchDocsActualDiskSize              []float64 `json:"couch_docs_actual_disk_size"`
			CouchDocsDataSize                    []float64 `json:"couch_docs_data_size"`
			CouchDocsDiskSize                    []float64 `json:"couch_docs_disk_size"`
			CouchSpatialDataSize                 []float64 `json:"couch_spatial_data_size"`
			CouchSpatialDiskSize                 []float64 `json:"couch_spatial_disk_size"`
			CouchSpatialOps                      []float64 `json:"couch_spatial_ops"`
			CouchViewsActualDiskSize             []float64 `json:"couch_views_actual_disk_size"`
			CouchViewsDataSize                   []float64 `json:"couch_views_data_size"`
			CouchViewsDiskSize                   []float64 `json:"couch_views_disk_size"`
			CouchViewsOps                        []float64 `json:"couch_views_ops"`
			CurrConnections                      []float64 `json:"curr_connections"`
			CurrItems                            []float64 `json:"curr_items"`
			CurrItemsTot                         []float64 `json:"curr_items_tot"`
			DecrHits                             []float64 `json:"decr_hits"`
			DecrMisses                           []float64 `json:"decr_misses"`
			DeleteHits                           []float64 `json:"delete_hits"`
			DeleteMisses                         []float64 `json:"delete_misses"`
			DiskCommitCount                      []float64 `json:"disk_commit_count"`
			DiskCommitTotal                      []float64 `json:"disk_commit_total"`
			DiskUpdateCount                      []float64 `json:"disk_update_count"`
			DiskUpdateTotal                      []float64 `json:"disk_update_total"`
			DiskWriteQueue                       []float64 `json:"disk_write_queue"`
			EpActiveAheadExceptions              []float64 `json:"ep_active_ahead_exceptions"` // couchbase 5.1.1
			EpActiveHlcDrift                     []float64 `json:"ep_active_hlc_drift"`        // couchbase 5.1.1
			EpActiveHlcDriftCount                []float64 `json:"ep_active_hlc_drift_count"`  // couchbase 5.1.1
			EpBgFetched                          []float64 `json:"ep_bg_fetched"`
			EpClockCasDriftThresholdExceeded     []float64 `json:"ep_clock_cas_drift_threshold_exceeded"` // couchbase 5.1.1
			EpDcp2IBackoff                       []float64 `json:"ep_dcp_2i_backoff"`
			EpDcp2ICount                         []float64 `json:"ep_dcp_2i_count"`
			EpDcp2IItemsRemaining                []float64 `json:"ep_dcp_2i_items_remaining"`
			EpDcp2IItemsSent                     []float64 `json:"ep_dcp_2i_items_sent"`
			EpDcp2IProducerCount                 []float64 `json:"ep_dcp_2i_producer_count"`
			EpDcp2ITotalBacklogSize              []float64 `json:"ep_dcp_2i_total_backlog_size"`
			EpDcp2ITotalBytes                    []float64 `json:"ep_dcp_2i_total_bytes"`
			EpDcpFtsBackoff                      []float64 `json:"ep_dcp_fts_backoff"`
			EpDcpFtsCount                        []float64 `json:"ep_dcp_fts_count"`
			EpDcpFtsItemsRemaining               []float64 `json:"ep_dcp_fts_items_remaining"`
			EpDcpFtsItemsSent                    []float64 `json:"ep_dcp_fts_items_sent"`
			EpDcpFtsProducerCount                []float64 `json:"ep_dcp_fts_producer_count"`
			EpDcpFtsTotalBacklogSize             []float64 `json:"ep_dcp_fts_total_backlog_size"`
			EpDcpFtsTotalBytes                   []float64 `json:"ep_dcp_fts_total_bytes"`
			EpDcpOtherBackoff                    []float64 `json:"ep_dcp_other_backoff"`
			EpDcpOtherCount                      []float64 `json:"ep_dcp_other_count"`
			EpDcpOtherItemsRemaining             []float64 `json:"ep_dcp_other_items_remaining"`
			EpDcpOtherItemsSent                  []float64 `json:"ep_dcp_other_items_sent"`
			EpDcpOtherProducerCount              []float64 `json:"ep_dcp_other_producer_count"`
			EpDcpOtherTotalBacklogSize           []float64 `json:"ep_dcp_other_total_backlog_size"`
			EpDcpOtherTotalBytes                 []float64 `json:"ep_dcp_other_total_bytes"`
			EpDcpReplicaBackoff                  []float64 `json:"ep_dcp_replica_backoff"`
			EpDcpReplicaCount                    []float64 `json:"ep_dcp_replica_count"`
			EpDcpReplicaItemsRemaining           []float64 `json:"ep_dcp_replica_items_remaining"`
			EpDcpReplicaItemsSent                []float64 `json:"ep_dcp_replica_items_sent"`
			EpDcpReplicaProducerCount            []float64 `json:"ep_dcp_replica_producer_count"`
			EpDcpReplicaTotalBacklogSize         []float64 `json:"ep_dcp_replica_total_backlog_size"`
			EpDcpReplicaTotalBytes               []float64 `json:"ep_dcp_replica_total_bytes"`
			EpDcpViewsBackoff                    []float64 `json:"ep_dcp_views_backoff"`
			EpDcpViewsCount                      []float64 `json:"ep_dcp_views_count"`
			EpDcpViewsItemsRemaining             []float64 `json:"ep_dcp_views_items_remaining"`
			EpDcpViewsItemsSent                  []float64 `json:"ep_dcp_views_items_sent"`
			EpDcpViewsProducerCount              []float64 `json:"ep_dcp_views_producer_count"`
			EpDcpViewsTotalBacklogSize           []float64 `json:"ep_dcp_views_total_backlog_size"`
			EpDcpViewsTotalBytes                 []float64 `json:"ep_dcp_views_total_bytes"`
			EpDcpXdcrBackoff                     []float64 `json:"ep_dcp_xdcr_backoff"`
			EpDcpXdcrCount                       []float64 `json:"ep_dcp_xdcr_count"`
			EpDcpXdcrItemsRemaining              []float64 `json:"ep_dcp_xdcr_items_remaining"`
			EpDcpXdcrItemsSent                   []float64 `json:"ep_dcp_xdcr_items_sent"`
			EpDcpXdcrProducerCount               []float64 `json:"ep_dcp_xdcr_producer_count"`
			EpDcpXdcrTotalBacklogSize            []float64 `json:"ep_dcp_xdcr_total_backlog_size"`
			EpDcpXdcrTotalBytes                  []float64 `json:"ep_dcp_xdcr_total_bytes"`
			EpDiskqueueDrain                     []float64 `json:"ep_diskqueue_drain"`
			EpDiskqueueFill                      []float64 `json:"ep_diskqueue_fill"`
			EpDiskqueueItems                     []float64 `json:"ep_diskqueue_items"`
			EpFlusherTodo                        []float64 `json:"ep_flusher_todo"`
			EpItemCommitFailed                   []float64 `json:"ep_item_commit_failed"`
			EpKvSize                             []float64 `json:"ep_kv_size"`
			EpMaxSize                            []float64 `json:"ep_max_size"`
			EpMemHighWat                         []float64 `json:"ep_mem_high_wat"`
			EpMemLowWat                          []float64 `json:"ep_mem_low_wat"`
			EpMetaDataMemory                     []float64 `json:"ep_meta_data_memory"`
			EpNumNonResident                     []float64 `json:"ep_num_non_resident"`
			EpNumOpsDelMeta                      []float64 `json:"ep_num_ops_del_meta"`
			EpNumOpsDelRetMeta                   []float64 `json:"ep_num_ops_del_ret_meta"`
			EpNumOpsGetMeta                      []float64 `json:"ep_num_ops_get_meta"`
			EpNumOpsSetMeta                      []float64 `json:"ep_num_ops_set_meta"`
			EpNumOpsSetRetMeta                   []float64 `json:"ep_num_ops_set_ret_meta"`
			EpNumValueEjects                     []float64 `json:"ep_num_value_ejects"`
			EpOomErrors                          []float64 `json:"ep_oom_errors"`
			EpOpsCreate                          []float64 `json:"ep_ops_create"`
			EpOpsUpdate                          []float64 `json:"ep_ops_update"`
			EpOverhead                           []float64 `json:"ep_overhead"`
			EpQueueSize                          []float64 `json:"ep_queue_size"`
			EpReplicaAheadExceptions             []float64 `json:"ep_replica_ahead_exceptions"`              // couchbase 5.1.1
			EpReplicaHlcDrift                    []float64 `json:"ep_replica_hlc_drift"`                     // couchbase 5.1.1
			EpReplicaHlcDriftCount               []float64 `json:"ep_replica_hlc_drift_count"`               // couchbase 5.1.1
			EpTapRebalanceCount                  []float64 `json:"ep_tap_rebalance_count"`                   // couchbase 4.5.1
			EpTapRebalanceQlen                   []float64 `json:"ep_tap_rebalance_qlen"`                    // couchbase 4.5.1
			EpTapRebalanceQueueBackfillremaining []float64 `json:"ep_tap_rebalance_queue_backfillremaining"` // couchbase 4.5.1
			EpTapRebalanceQueueBackoff           []float64 `json:"ep_tap_rebalance_queue_backoff"`           // couchbase 4.5.1
			EpTapRebalanceQueueDrain             []float64 `json:"ep_tap_rebalance_queue_drain"`             // couchbase 4.5.1
			EpTapRebalanceQueueFill              []float64 `json:"ep_tap_rebalance_queue_fill"`              // couchbase 4.5.1
			EpTapRebalanceQueueItemondisk        []float64 `json:"ep_tap_rebalance_queue_itemondisk"`        // couchbase 4.5.1
			EpTapRebalanceTotalBacklogSize       []float64 `json:"ep_tap_rebalance_total_backlog_size"`      // couchbase 4.5.1
			EpTapReplicaCount                    []float64 `json:"ep_tap_replica_count"`                     // couchbase 4.5.1
			EpTapReplicaQlen                     []float64 `json:"ep_tap_replica_qlen"`                      // couchbase 4.5.1
			EpTapReplicaQueueBackfillremaining   []float64 `json:"ep_tap_replica_queue_backfillremaining"`   // couchbase 4.5.1
			EpTapReplicaQueueBackoff             []float64 `json:"ep_tap_replica_queue_backoff"`             // couchbase 4.5.1
			EpTapReplicaQueueDrain               []float64 `json:"ep_tap_replica_queue_drain"`               // couchbase 4.5.1
			EpTapReplicaQueueFill                []float64 `json:"ep_tap_replica_queue_fill"`                // couchbase 4.5.1
			EpTapReplicaQueueItemondisk          []float64 `json:"ep_tap_replica_queue_itemondisk"`          // couchbase 4.5.1
			EpTapReplicaTotalBacklogSize         []float64 `json:"ep_tap_replica_total_backlog_size"`        // couchbase 4.5.1
			EpTapTotalCount                      []float64 `json:"ep_tap_total_count"`                       // couchbase 4.5.1
			EpTapTotalQlen                       []float64 `json:"ep_tap_total_qlen"`                        // couchbase 4.5.1
			EpTapTotalQueueBackfillremaining     []float64 `json:"ep_tap_total_queue_backfillremaining"`     // couchbase 4.5.1
			EpTapTotalQueueBackoff               []float64 `json:"ep_tap_total_queue_backoff"`               // couchbase 4.5.1
			EpTapTotalQueueDrain                 []float64 `json:"ep_tap_total_queue_drain"`                 // couchbase 4.5.1
			EpTapTotalQueueFill                  []float64 `json:"ep_tap_total_queue_fill"`                  // couchbase 4.5.1
			EpTapTotalQueueItemondisk            []float64 `json:"ep_tap_total_queue_itemondisk"`            // couchbase 4.5.1
			EpTapTotalTotalBacklogSize           []float64 `json:"ep_tap_total_total_backlog_size"`          // couchbase 4.5.1
			EpTapUserCount                       []float64 `json:"ep_tap_user_count"`                        // couchbase 4.5.1
			EpTapUserQlen                        []float64 `json:"ep_tap_user_qlen"`                         // couchbase 4.5.1
			EpTapUserQueueBackfillremaining      []float64 `json:"ep_tap_user_queue_backfillremaining"`      // couchbase 4.5.1
			EpTapUserQueueBackoff                []float64 `json:"ep_tap_user_queue_backoff"`                // couchbase 4.5.1
			EpTapUserQueueDrain                  []float64 `json:"ep_tap_user_queue_drain"`                  // couchbase 4.5.1
			EpTapUserQueueFill                   []float64 `json:"ep_tap_user_queue_fill"`                   // couchbase 4.5.1
			EpTapUserQueueItemondisk             []float64 `json:"ep_tap_user_queue_itemondisk"`             // couchbase 4.5.1
			EpTapUserTotalBacklogSize            []float64 `json:"ep_tap_user_total_backlog_size"`           // couchbase 4.5.1
			EpTmpOomErrors                       []float64 `json:"ep_tmp_oom_errors"`
			EpVbTotal                            []float64 `json:"ep_vb_total"`
			Evictions                            []float64 `json:"evictions"`
			GetHits                              []float64 `json:"get_hits"`
			GetMisses                            []float64 `json:"get_misses"`
			IncrHits                             []float64 `json:"incr_hits"`
			IncrMisses                           []float64 `json:"incr_misses"`
			MemUsed                              []float64 `json:"mem_used"`
			Misses                               []float64 `json:"misses"`
			Ops                                  []float64 `json:"ops"`
			VbActiveEject                        []float64 `json:"vb_active_eject"`
			VbActiveItmMemory                    []float64 `json:"vb_active_itm_memory"`
			VbActiveMetaDataMemory               []float64 `json:"vb_active_meta_data_memory"`
			VbActiveNum                          []float64 `json:"vb_active_num"`
			VbActiveNumNonResident               []float64 `json:"vb_active_num_non_resident"`
			VbActiveOpsCreate                    []float64 `json:"vb_active_ops_create"`
			VbActiveOpsUpdate                    []float64 `json:"vb_active_ops_update"`
			VbActiveQueueAge                     []float64 `json:"vb_active_queue_age"`
			VbActiveQueueDrain                   []float64 `json:"vb_active_queue_drain"`
			VbActiveQueueFill                    []float64 `json:"vb_active_queue_fill"`
			VbActiveQueueSize                    []float64 `json:"vb_active_queue_size"`
			VbPendingCurrItems                   []float64 `json:"vb_pending_curr_items"`
			VbPendingEject                       []float64 `json:"vb_pending_eject"`
			VbPendingItmMemory                   []float64 `json:"vb_pending_itm_memory"`
			VbPendingMetaDataMemory              []float64 `json:"vb_pending_meta_data_memory"`
			VbPendingNum                         []float64 `json:"vb_pending_num"`
			VbPendingNumNonResident              []float64 `json:"vb_pending_num_non_resident"`
			VbPendingOpsCreate                   []float64 `json:"vb_pending_ops_create"`
			VbPendingOpsUpdate                   []float64 `json:"vb_pending_ops_update"`
			VbPendingQueueAge                    []float64 `json:"vb_pending_queue_age"`
			VbPendingQueueDrain                  []float64 `json:"vb_pending_queue_drain"`
			VbPendingQueueFill                   []float64 `json:"vb_pending_queue_fill"`
			VbPendingQueueSize                   []float64 `json:"vb_pending_queue_size"`
			VbReplicaCurrItems                   []float64 `json:"vb_replica_curr_items"`
			VbReplicaEject                       []float64 `json:"vb_replica_eject"`
			VbReplicaItmMemory                   []float64 `json:"vb_replica_itm_memory"`
			VbReplicaMetaDataMemory              []float64 `json:"vb_replica_meta_data_memory"`
			VbReplicaNum                         []float64 `json:"vb_replica_num"`
			VbReplicaNumNonResident              []float64 `json:"vb_replica_num_non_resident"`
			VbReplicaOpsCreate                   []float64 `json:"vb_replica_ops_create"`
			VbReplicaOpsUpdate                   []float64 `json:"vb_replica_ops_update"`
			VbReplicaQueueAge                    []float64 `json:"vb_replica_queue_age"`
			VbReplicaQueueDrain                  []float64 `json:"vb_replica_queue_drain"`
			VbReplicaQueueFill                   []float64 `json:"vb_replica_queue_fill"`
			VbReplicaQueueSize                   []float64 `json:"vb_replica_queue_size"`
			VbTotalQueueAge                      []float64 `json:"vb_total_queue_age"`
			XdcOps                               []float64 `json:"xdc_ops"`
			CPUIdleMs                            []float64 `json:"cpu_idle_ms"`
			CPULocalMs                           []float64 `json:"cpu_local_ms"`
			CPUUtilizationRate                   []float64 `json:"cpu_utilization_rate"`
			HibernatedRequests                   []float64 `json:"hibernated_requests"`
			HibernatedWaked                      []float64 `json:"hibernated_waked"`
			MemActualFree                        []float64 `json:"mem_actual_free"`
			MemActualUsed                        []float64 `json:"mem_actual_used"`
			MemFree                              []float64 `json:"mem_free"`
			MemTotal                             []float64 `json:"mem_total"`
			MemUsedSys                           []float64 `json:"mem_used_sys"`
			RestRequests                         []float64 `json:"rest_requests"`
			SwapTotal                            []float64 `json:"swap_total"`
			SwapUsed                             []float64 `json:"swap_used"`
		} `json:"samples"`
	} `json:"op"`
}

// BucketStatsExporter encapsulates bucket stats and context.
type BucketStatsExporter struct {
	context Context
	route   string
	metrics map[string]*CustomDesc
}

// NewBucketStatsExporter creates the BucketStatsExporter and fill it with metrics metadata from the metrics file.
func NewBucketStatsExporter(context Context) (*BucketStatsExporter, error) {
	bucketStatsMetrics, err := GetMetricsFromFile("bucketstats")
	if err != nil {
		return &BucketStatsExporter{}, err
	}
	// metrics is a map where the key is the metric ID and the value is a Prometheus Descriptor for that metric.
	metrics := make(map[string]*CustomDesc, len(bucketStatsMetrics.List))
	for _, metric := range bucketStatsMetrics.List {
		fqName := p.BuildFQName("cb", bucketStatsMetrics.Name, metric.Name)
		metrics[metric.ID] = newCustomDesc(fqName, metric.Description, metric.Labels, metric.Type)
	}
	return &BucketStatsExporter{
		context: context,
		route:   bucketStatsMetrics.Route,
		metrics: metrics,
	}, nil
}

// Describe describes exported metrics.
func (e *BucketStatsExporter) Describe(ch chan<- *p.Desc) {
	for _, metric := range e.metrics {
		ch <- metric.pDesc
	}
}

// Collect fetches data for each exported metric.
func (e *BucketStatsExporter) Collect(ch chan<- p.Metric) {
	body, err := Fetch(e.context, e.route)
	if err != nil {
		log.Error("Error when retrieving bucketstats data. Bucketstats metrics won't be scraped")
		return
	}
	var buckets []BucketData
	err = json.Unmarshal(body, &buckets)
	if err != nil {
		log.Error("Could not unmarshal buckets data")
		return
	}

	// Each bucket has its own API route.
	var routes []string
	for _, bucket := range buckets {
		routes = append(routes, e.route+"/"+bucket.Name+"/stats")
	}

	bodies := MultiFetch(e.context, routes)

	for _, bucket := range buckets {
		var bucketStats BucketStatsData
		err = json.Unmarshal(bodies[e.route+"/"+bucket.Name+"/stats"], &bucketStats)
		if err != nil {
			log.Error("Could not unmarshal bucketstats data for bucket " + bucket.Name)
			return
		}
		flat := FlattenStruct(bucketStats)
		for id, metric := range e.metrics {
			if array, ok := flat[id].([]float64); ok {
				lenArray := len(array)
				if lenArray != 0 {
					ch <- p.MustNewConstMetric(metric.pDesc, getValueType(metric.mType), array[lenArray-1], bucket.Name)
				}
			}
		}
	}
}
