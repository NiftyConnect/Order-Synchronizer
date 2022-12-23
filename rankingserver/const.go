package rankingserver

const (
	defaultListenAddr = "0.0.0.0:8080"
	twitterPointsFile = "twitterPoints.cfg"
)

type GeneralNFT struct {
	Key   string `json:"key"`
	Value int    `json:"value"`
}

type PremiumNFT struct {
	Key     string `json:"key"`
	Value   int    `json:"value"`
	Ranking int    `json:"ranking"`
}
