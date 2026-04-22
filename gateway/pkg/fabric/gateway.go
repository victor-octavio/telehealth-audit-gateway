package fabric

import (
	"fmt"
	"os"
	"time"

	"github.com/hyperledger/fabric-gateway/pkg/client"
	"github.com/hyperledger/fabric-gateway/pkg/identity"
	"google.golang.org/grpc"
	"google.golang.org/grpc/credentials/insecure"
)

type Gateway struct {
	client   *client.Gateway
	network  *client.Network
	contract *client.Contract
	conn     *grpc.ClientConn
}

func NewGateway(cfg FabricConfig) (*Gateway, error) {
	conn, err := newGrpcConnection(cfg)
	if err != nil {
		return nil, fmt.Errorf("erro ao conectar via gRPC: %w", err)
	}

	id, err := newIdentity(cfg)
	if err != nil {
		conn.Close()
		return nil, fmt.Errorf("erro ao carregar identidade: %w", err)
	}

	sign, err := newSigner(cfg)
	if err != nil {
		conn.Close()
		return nil, fmt.Errorf("erro ao carregar assinador: %w", err)
	}

	gw, err := client.Connect(
		id,
		client.WithSign(sign),
		client.WithClientConnection(conn),
		client.WithEvaluateTimeout(5*time.Second),
		client.WithEndorseTimeout(15*time.Second),
		client.WithSubmitTimeout(5*time.Second),
		client.WithCommitStatusTimeout(60*time.Second),
	)
	if err != nil {
		conn.Close()
		return nil, fmt.Errorf("erro ao criar gateway: %w", err)
	}

	network := gw.GetNetwork(cfg.ChannelName)
	contract := network.GetContract(cfg.ChaincodeName)

	return &Gateway{
		client:   gw,
		network:  network,
		contract: contract,
		conn:     conn,
	}, nil
}

func (g *Gateway) Contract() *client.Contract {
	return g.contract
}

func (g *Gateway) Close() {
	g.client.Close()
	g.conn.Close()
}

func newGrpcConnection(cfg FabricConfig) (*grpc.ClientConn, error) {
	// passa o authority header para o proxy do Microfab rotear corretamente
	//md := metadata.Pairs(":authority", cfg.GatewayPeer)
	return grpc.NewClient(
		cfg.PeerEndpoint,
		grpc.WithTransportCredentials(insecure.NewCredentials()),
		grpc.WithAuthority(cfg.GatewayPeer),
	)
}

func newIdentity(cfg FabricConfig) (*identity.X509Identity, error) {
	certPEM, err := os.ReadFile(cfg.CertPath)
	if err != nil {
		return nil, fmt.Errorf("erro ao ler certificado: %w", err)
	}

	cert, err := identity.CertificateFromPEM(certPEM)
	if err != nil {
		return nil, fmt.Errorf("erro ao parsear certificado: %w", err)
	}

	return identity.NewX509Identity(cfg.MspID, cert)
}

func newSigner(cfg FabricConfig) (identity.Sign, error) {
	keyPEM, err := os.ReadFile(cfg.KeyPath)
	if err != nil {
		return nil, fmt.Errorf("erro ao ler chave privada: %w", err)
	}

	privateKey, err := identity.PrivateKeyFromPEM(keyPEM)
	if err != nil {
		return nil, fmt.Errorf("erro ao parsear chave privada: %w", err)
	}

	return identity.NewPrivateKeySign(privateKey)
}
