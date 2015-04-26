package structure

// CollectorDetails type holds collector data
type CollectorDetails struct {
	CollectorName string `json:"collectorName"`
	CashBalance   int    `json:"cashBalance"`
	DreamBalance  int    `json:"dreamBalance"`
}

// CollectorTier type holds collector tier
type CollectorTier struct {
	CollectorTier string `json:"tier"`
}
