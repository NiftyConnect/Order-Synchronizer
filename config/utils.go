package config

import "strings"

func FormatBlockchainName(blockchain string) string {
	return strings.ToUpper(blockchain)
}
