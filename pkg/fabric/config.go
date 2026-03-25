package fabric

import "os"

type FabricConfig struct {
	PeerEndpoint  string
	GatewayPeer   string
	MspID         string
	CertPath      string
	KeyPath       string
	TLSCertPath   string
	ChannelName   string
	ChaincodeName string
}

func Load() FabricConfig {
	return FabricConfig{
		PeerEndpoint:  getEnv("PEER_ENDPOINT", "localhost:8080"),
		GatewayPeer:   getEnv("GATEWAY_PEER", "org1peer-api.127-0-0-1.nip.io:8080"),
		MspID:         getEnv("MSP_ID", "Org1MSP"),
		CertPath:      getEnv("CERT_PATH", "../../../fabric-telehealth-audit/network/wallet/org1admin/msp/signcerts/cert.pem"),
		KeyPath:       getEnv("KEY_PATH", "../../../fabric-telehealth-audit/network/wallet/org1admin/msp/keystore/key.pem"),
		TLSCertPath:   getEnv("TLS_CERT_PATH", ""),
		ChannelName:   getEnv("CHANNEL_NAME", "mychannel"),
		ChaincodeName: getEnv("CHAINCODE_NAME", "diagnostic-log"),
	}
}

func getEnv(key, fallback string) string {
	if val := os.Getenv(key); val != "" {
		return val
	}
	return fallback
}
