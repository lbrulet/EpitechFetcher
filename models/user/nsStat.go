package user

type NsStatInfo struct {
	Active    float32 `json:"active"`
	Idle      int     `json:"idle"`
	OutActive int     `json:"out_active"`
	OutIdle   int     `json:"out_idle"`
	NsLogNorm int     `json:"nslog_norm"`
}
