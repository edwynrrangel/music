package playlist

import (
	"context"
	"log"

	"github.com/edwynrrangel/grpc/go/multimedia/internal/domain/content"
)

type usecase struct {
	playListRepository Repository
	contentRepository  content.Repository
}

func NewUseCase(
	playListRepository Repository,
	contentRepository content.Repository,
) UseCase {
	return &usecase{
		playListRepository: playListRepository,
		contentRepository:  contentRepository,
	}
}

func (u *usecase) Add(ctx context.Context, request *PlayListRequest) error {
	log.Printf("Received request: %+v", request)
	content, err := u.contentRepository.Get(ctx, request.ContentID)
	if err != nil {
		return err
	}
	if content == nil {
		return nil
	}

	if request.ID != "" {
		return u.playListRepository.Update(ctx, request.ID, request.UserID, (Content)(*content))
	}

	if request.Name == "" {
		request.Name = "new playlist"
	}
	playList := &Playlist{
		Name:   request.Name,
		UserID: request.UserID,
		Content: []Content{
			(Content)(*content),
		},
	}

	if err = u.playListRepository.Add(ctx, playList); err != nil {
		return err
	}
	request.ID = playList.ID
	return nil
}

func (u *usecase) Get(ctx context.Context, request *PlayListRequest) (*Playlist, error) {
	log.Printf("Received request: %+v", request)
	return u.playListRepository.Get(ctx, request.ID, request.UserID)
}

func (u *usecase) Remove(ctx context.Context, request *PlayListRequest) error {
	log.Printf("Received request: %+v", request)
	return u.playListRepository.RemoveContent(ctx, request.ID, request.UserID, request.ContentID)
}

func (u *usecase) List(userID string) ([]Playlist, error) {
	log.Printf("Received request: %s", userID)
	return u.playListRepository.List(context.Background(), userID)
}
