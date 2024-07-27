package playlist

import (
	"context"
	"log"

	"github.com/edwynrrangel/grpc/go/playlist_server/internal/domain/content"
)

type usecase struct {
	playListRepository Repository
	contentClient      content.Adapter
}

func NewUseCase(
	playListRepository Repository,
	contentClient content.Adapter,
) UseCase {
	return &usecase{
		playListRepository: playListRepository,
		contentClient:      contentClient,
	}
}

func (u *usecase) Add(ctx context.Context, request *PlaylistRequest) error {
	log.Printf("Add received request: %+v", request)
	content, err := u.contentClient.Get(ctx, request.ContentID)
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

func (u *usecase) Get(ctx context.Context, request *PlaylistRequest) (*PlaylistResponse, error) {
	log.Printf("Get received request: %+v", request)
	got, err := u.playListRepository.Get(ctx, request.ID, request.UserID)
	if err != nil {
		return nil, err
	}
	if got == nil {
		return nil, nil
	}

	return got.toDTO(), nil
}

func (u *usecase) RemoveContent(ctx context.Context, request *PlaylistRequest) error {
	log.Printf("Remove content received request: %+v", request)
	return u.playListRepository.RemoveContent(ctx, request.ID, request.UserID, request.ContentID)
}

func (u *usecase) List(userID string) ([]Playlist, error) {
	log.Printf("List received request: %s", userID)
	return u.playListRepository.List(context.Background(), userID)
}

func (u *usecase) Remove(ctx context.Context, request *RemovePlaylistRequest) error {
	log.Printf("Remove received request: %+v", request)
	return u.playListRepository.Remove(ctx, request.ID, request.UserID)
}
