package utils

import (
	"github.com/zairza-cetb/bench-routes/tsdb"
)

const (
	// ConfigurationFilePath is the constant path to the configuration file needed to start the application
	// written from root file since the application starts from `make run`
	ConfigurationFilePath = "local-config.yml"
	// PathPing stores the defualt address of storage directory of ping data
	PathPing = "storage/ping"
	// PathJitter stores the defualt address of storage directory of jitter data
	PathJitter = "storage/jitter"
	// PathFloodPing stores the defualt address of storage directory of flood ping data
	PathFloodPing = "storage/flood-ping"
	// PathReqResDelayMonitoring stores the defualt address of storage directory of req-res and monitoring data
	PathReqResDelayMonitoring = "storage/req-res-delay-monitoring"
)

var (
	// PingDBNames contains the name of the database corresponding to the unique config url
	PingDBNames = make(map[string]string)
	// FloodPingDBNames contains the name of the flood ping corresponding to the unique config url
	FloodPingDBNames = make(map[string]string)
	// GlobalPingChain contains chains of all the pings operating in bench-routes which has to be globally accessed
	// This is necessary as it helps to retain the parent values which are required for concurreny
	GlobalPingChain []*tsdb.Chain

	// GlobalChain is the global chain array which can be used to maintain a list of chains that represent
	// the time-series values
	GlobalChain []*tsdb.Chain

	// GlobalFloodPingChain contains chains of flood ping operations in bench-routes which has to be globally accessed
	// This is necessary as it helps to retain the parent values which are required for concurreny
	GlobalFloodPingChain []*tsdb.Chain

	// GlobalReqResDelChain stands for Request-Response-Delay
	GlobalReqResDelChain []*tsdb.Chain
)

// TypePingScrap as datatype for ping outputs
type TypePingScrap struct {
	Min, Avg, Max, Mdev float64
}

// TypeFloodPingScrap as datatype for flood ping outputs
type TypeFloodPingScrap struct {
	Min, Avg, Max, Mdev, PacketLoss float64
}
