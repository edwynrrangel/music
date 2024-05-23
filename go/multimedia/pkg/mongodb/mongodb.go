package mongodb

import (
	"context"
	"crypto/tls"
	"crypto/x509"
	"encoding/base64"
	"fmt"

	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"

	"github.com/edwynrrangel/grpc/go/multimedia/config"
)

type builder struct {
	host     string
	port     string
	username string
	password string
	tls      bool
	ca       string
	params   map[string]string
	options  *options.ClientOptions
}

func NewBuilder(config config.Config) *builder {
	return &builder{
		host:     config.MongoDB.Host,
		port:     config.MongoDB.Port,
		username: config.MongoDB.Username,
		password: config.MongoDB.Password,
		params:   make(map[string]string),
	}
}

// WithTLS sets the TLS configuration for the MongoDB connection.
// The CA certificate is expected to be base64 encoded.
func (b *builder) WithTLS(ca string) *builder {
	b.tls = true
	b.ca = ca
	b.params["tls"] = "true"
	return b
}

// WithRetryWrites sets the retryWrites parameter for the MongoDB connection.
// The value should be a string representation of a boolean value (true or false).
func (b *builder) WithRetryWrites(value string) *builder {
	b.params["retryWrites"] = value
	return b
}

func (b *builder) uri() string {
	var params string
	uri := fmt.Sprintf("mongodb://%s:%s", b.host, b.port)
	if len(b.params) > 0 {
		for k, v := range b.params {
			params += fmt.Sprintf("&%s=%s", k, v)
		}
		uri = fmt.Sprintf("%s/?%s", uri, params[1:])
	}
	return uri
}

func (b *builder) setOptionURI() *builder {
	b.options = options.Client().ApplyURI(b.uri())
	return b
}

func (b *builder) setOptionAuth() *builder {
	if b.username != "" && b.password != "" {
		b.options = b.options.SetAuth(options.Credential{
			Username: b.username,
			Password: b.password,
		})
	}
	return b
}

func (b *builder) getTLSConfig() (*tls.Config, error) {
	caCert, err := base64.StdEncoding.DecodeString(b.ca)
	if err != nil {
		return nil, err
	}
	roots := x509.NewCertPool()
	if ok := roots.AppendCertsFromPEM(caCert); !ok {
		return nil, fmt.Errorf("failed to parse root certificate")
	}
	return &tls.Config{
		RootCAs: roots,
	}, nil
}

func (b *builder) setOptionTLS() (*builder, error) {
	tlsConfig, err := b.getTLSConfig()
	if err != nil {
		return nil, err
	}
	b.options = b.options.SetTLSConfig(tlsConfig)
	return b, nil
}

// GetClient returns a new MongoDB client.
// The client is not connected to the MongoDB server.
func (b *builder) GetClient(ctx context.Context) (*mongo.Client, error) {
	b.setOptionURI().setOptionAuth()
	if b.tls {
		_, err := b.setOptionTLS()
		if err != nil {
			return nil, err
		}
	}
	return mongo.Connect(ctx, b.options)
}
