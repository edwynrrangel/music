package multimedia

import (
	"google.golang.org/grpc"

	"github.com/edwynrrangel/grpc/go/multimedia/config"
	"github.com/edwynrrangel/grpc/go/multimedia/internal/multimedia"
	"github.com/edwynrrangel/grpc/go/multimedia/internal/proto"
)

func Register(s *grpc.Server, config *config.Config) {
	proto.RegisterMultimediaServiceServer(
		s,
		NewAdapter(
			multimedia.NewUseCase(
				multimedia.NewRepository(),
			),
		),
	)
}
