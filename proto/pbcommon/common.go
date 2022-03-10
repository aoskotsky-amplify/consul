package pbcommon

import (
	"time"

	"github.com/hashicorp/consul/agent/structs"
)

// IsRead is always true for QueryOption
func (q *QueryOptions) IsRead() bool {
	return true
}

// AllowStaleRead returns whether a stale read should be allowed
func (q *QueryOptions) AllowStaleRead() bool {
	return q.AllowStale
}

func (q *QueryOptions) TokenSecret() string {
	return q.Token
}

func (q *QueryOptions) SetTokenSecret(s string) {
	q.Token = s
}

// SetToken is needed to implement the structs.QueryOptionsCompat interface
func (q *QueryOptions) SetToken(token string) {
	q.Token = token
}

// SetMinQueryIndex is needed to implement the structs.QueryOptionsCompat interface
func (q *QueryOptions) SetMinQueryIndex(minQueryIndex uint64) {
	q.MinQueryIndex = minQueryIndex
}

// SetMaxQueryTime is needed to implement the structs.QueryOptionsCompat interface
func (q *QueryOptions) SetMaxQueryTime(maxQueryTime time.Duration) {
	q.MaxQueryTime = maxQueryTime
}

// SetAllowStale is needed to implement the structs.QueryOptionsCompat interface
func (q *QueryOptions) SetAllowStale(allowStale bool) {
	q.AllowStale = allowStale
}

// SetRequireConsistent is needed to implement the structs.QueryOptionsCompat interface
func (q *QueryOptions) SetRequireConsistent(requireConsistent bool) {
	q.RequireConsistent = requireConsistent
}

// SetUseCache is needed to implement the structs.QueryOptionsCompat interface
func (q *QueryOptions) SetUseCache(useCache bool) {
	q.UseCache = useCache
}

// SetMaxStaleDuration is needed to implement the structs.QueryOptionsCompat interface
func (q *QueryOptions) SetMaxStaleDuration(maxStaleDuration time.Duration) {
	q.MaxStaleDuration = maxStaleDuration
}

// SetMaxAge is needed to implement the structs.QueryOptionsCompat interface
func (q *QueryOptions) SetMaxAge(maxAge time.Duration) {
	q.MaxAge = maxAge
}

// SetMustRevalidate is needed to implement the structs.QueryOptionsCompat interface
func (q *QueryOptions) SetMustRevalidate(mustRevalidate bool) {
	q.MustRevalidate = mustRevalidate
}

// SetStaleIfError is needed to implement the structs.QueryOptionsCompat interface
func (q *QueryOptions) SetStaleIfError(staleIfError time.Duration) {
	q.StaleIfError = staleIfError
}

func (q QueryOptions) HasTimedOut(start time.Time, rpcHoldTimeout, maxQueryTime, defaultQueryTime time.Duration) bool {
	o := structs.QueryOptions{
		MaxQueryTime:  q.MaxQueryTime,
		MinQueryIndex: q.MinQueryIndex,
	}
	return o.HasTimedOut(start, rpcHoldTimeout, maxQueryTime, defaultQueryTime)
}

// SetFilter is needed to implement the structs.QueryOptionsCompat interface
func (q *QueryOptions) SetFilter(filter string) {
	q.Filter = filter
}

// GetMaxQueryTime is required to implement blockingQueryOptions
func (m *QueryOptions) GetMaxQueryTime() time.Duration {
	if m != nil {
		return m.MaxQueryTime
	}
	return 0
}

// GetMinQueryIndex is required to implement blockingQueryOptions
func (m *QueryOptions) GetMinQueryIndex() uint64 {
	if m != nil {
		return m.MinQueryIndex
	}
	return 0
}

// GetRequireConsistent is required to implement blockingQueryOptions
func (m *QueryOptions) GetRequireConsistent() bool {
	if m != nil {
		return m.RequireConsistent
	}
	return false
}

// GetToken is required to implement blockingQueryOptions
func (m *QueryOptions) GetToken() string {
	if m != nil {
		return m.Token
	}
	return ""
}

// GetAllowStale is required to implement structs.QueryOptionsCompat
func (m *QueryOptions) GetAllowStale() bool {
	if m != nil {
		return m.AllowStale
	}
	return false
}

// GetFilter is required to implement structs.QueryOptionsCompat
func (m *QueryOptions) GetFilter() string {
	if m != nil {
		return m.Filter
	}
	return ""
}

// GetResultsFilteredByACLs is required to implement structs.QueryOptionsCompat
func (m *QueryOptions) GetMaxAge() time.Duration {
	if m != nil {
		return m.MaxAge
	}
	return 0
}

// GetMaxStaleDuration is required to implement structs.QueryOptionsCompat
func (m *QueryOptions) GetMaxStaleDuration() time.Duration {
	if m != nil {
		return m.MaxStaleDuration
	}
	return 0
}

// GetMustRevalidate is required to implement structs.QueryOptionsCompat
func (m *QueryOptions) GetMustRevalidate() bool {
	if m != nil {
		return m.MustRevalidate
	}
	return false
}

// GetStaleIfError is required to implement structs.QueryOptionsCompat
func (m *QueryOptions) GetStaleIfError() time.Duration {
	if m != nil {
		return m.StaleIfError
	}
	return 0
}

// GetUseCache is required to implement structs.QueryOptionsCompat
func (m *QueryOptions) GetUseCache() bool {
	if m != nil {
		return m.UseCache
	}
	return false
}

// SetLastContact is needed to implement the structs.QueryMetaCompat interface
func (q *QueryMeta) SetLastContact(lastContact time.Duration) {
	q.LastContact = lastContact
}

// SetKnownLeader is needed to implement the structs.QueryMetaCompat interface
func (q *QueryMeta) SetKnownLeader(knownLeader bool) {
	q.KnownLeader = knownLeader
}

// SetIndex is needed to implement the structs.QueryMetaCompat interface
func (q *QueryMeta) SetIndex(index uint64) {
	q.Index = index
}

// SetConsistencyLevel is needed to implement the structs.QueryMetaCompat interface
func (q *QueryMeta) SetConsistencyLevel(consistencyLevel string) {
	q.ConsistencyLevel = consistencyLevel
}

func (q *QueryMeta) GetBackend() structs.QueryBackend {
	return structs.QueryBackend(0)
}

// SetResultsFilteredByACLs is needed to implement the structs.QueryMetaCompat interface
func (q *QueryMeta) SetResultsFilteredByACLs(v bool) {
	q.ResultsFilteredByACLs = v
}

// GetIndex is required to implement blockingQueryResponseMeta
func (m *QueryMeta) GetIndex() uint64 {
	if m != nil {
		return m.Index
	}
	return 0
}

// GetConsistencyLevel is required to implement structs.QueryMetaCompat
func (m *QueryMeta) GetConsistencyLevel() string {
	if m != nil {
		return m.ConsistencyLevel
	}
	return ""
}

// GetKnownLeader is required to implement structs.QueryMetaCompat
func (m *QueryMeta) GetKnownLeader() bool {
	if m != nil {
		return m.KnownLeader
	}
	return false
}

// GetLastContact is required to implement structs.QueryMetaCompat
func (m *QueryMeta) GetLastContact() time.Duration {
	if m != nil {
		return m.LastContact
	}
	return 0
}

// GetResultsFilteredByACLs is required to implement structs.QueryMetaCompat
func (m *QueryMeta) GetResultsFilteredByACLs() bool {
	if m != nil {
		return m.ResultsFilteredByACLs
	}
	return false
}

// WriteRequest only applies to writes, always false
//
// IsRead implements structs.RPCInfo
func (w WriteRequest) IsRead() bool {
	return false
}

// SetTokenSecret implements structs.RPCInfo
func (w WriteRequest) TokenSecret() string {
	return w.Token
}

// SetTokenSecret implements structs.RPCInfo
func (w *WriteRequest) SetTokenSecret(s string) {
	w.Token = s
}

// AllowStaleRead returns whether a stale read should be allowed
//
// AllowStaleRead implements structs.RPCInfo
func (w WriteRequest) AllowStaleRead() bool {
	return false
}

// HasTimedOut implements structs.RPCInfo
func (w WriteRequest) HasTimedOut(start time.Time, rpcHoldTimeout, _, _ time.Duration) bool {
	return time.Since(start) > rpcHoldTimeout
}

// IsRead implements structs.RPCInfo
func (r *ReadRequest) IsRead() bool {
	return true
}

// AllowStaleRead implements structs.RPCInfo
func (r *ReadRequest) AllowStaleRead() bool {
	// TODO(partitions): plumb this?
	return false
}

// TokenSecret implements structs.RPCInfo
func (r *ReadRequest) TokenSecret() string {
	return r.Token
}

// SetTokenSecret implements structs.RPCInfo
func (r *ReadRequest) SetTokenSecret(token string) {
	r.Token = token
}

// HasTimedOut implements structs.RPCInfo
func (r *ReadRequest) HasTimedOut(start time.Time, rpcHoldTimeout, maxQueryTime, defaultQueryTime time.Duration) bool {
	return time.Since(start) > rpcHoldTimeout
}

// RequestDatacenter implements structs.RPCInfo
func (td TargetDatacenter) RequestDatacenter() string {
	return td.Datacenter
}
