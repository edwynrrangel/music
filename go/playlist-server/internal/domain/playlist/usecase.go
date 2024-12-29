package playlist

import (
	"context"
	"log"

	"github.com/edwynrrangel/grpc/go/playlist_server/internal/shared"
)

type usecase struct {
	playlistRepo Repository
}

func NewUseCase(
	playlistRepo Repository,
) UseCase {
	return &usecase{
		playlistRepo: playlistRepo,
	}
}

func (u *usecase) Create(ctx context.Context, data *CreateRequest) (*Playlist, error) {
	log.Printf("Create received request: %+v", data)
	if data.Name == "" {
		data.Name = "New Playlist"
	}
	playlist := &Playlist{
		UserId: data.UserId,
		Name:   data.Name,
	}
	err := u.playlistRepo.Create(ctx, playlist)
	if err != nil {
		return nil, err
	}

	return playlist, nil
}

func (u *usecase) List(ctx context.Context, data ListRequest) (*shared.Paginated[Playlist], error) {
	total, err := u.playlistRepo.Count(ctx, data.UserId, data.Query)
	if err != nil {
		return nil, err
	}

	if total == 0 {
		return nil, nil
	}

	playlists, err := u.playlistRepo.List(ctx, data.UserId, data.Query)
	if err != nil {
		return nil, err
	}

	return &shared.Paginated[Playlist]{
		Page:  data.Page,
		Limit: data.Limit,
		Total: total,
		Data:  playlists,
	}, nil
}

func (u *usecase) Get(ctx context.Context, userId string, playlistId string) (*Playlist, error) {
	log.Printf("Get received userId: %s playlistId: %s", userId, playlistId)
	got, err := u.playlistRepo.Get(ctx, userId, playlistId)
	if err != nil {
		return nil, err
	}

	return got, nil
}

func (u *usecase) AddContent(ctx context.Context, userId string, playlistId string, content Content) (*Playlist, error) {
	log.Printf("AddContent received userId: %s playlistId: %s contentId: %s", userId, playlistId, content.ID)
	err := u.playlistRepo.Update(ctx, userId, playlistId, content)
	if err != nil {
		return nil, err
	}

	return u.playlistRepo.Get(ctx, userId, playlistId)
}
