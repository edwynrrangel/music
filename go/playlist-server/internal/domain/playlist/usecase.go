package playlist

import (
	"context"
	"fmt"
	"log"

	"github.com/edwynrrangel/grpc/go/playlist-server/internal/shared"
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
		Mode:   data.Mode,
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
	playlist, err := u.playlistRepo.Get(ctx, userId, playlistId)
	if err != nil {
		return nil, err
	}

	if playlist == nil {
		return nil, fmt.Errorf("playlist not found")
	}

	contentExists := false
	for i, c := range playlist.Data {
		if c.ID == content.ID {
			contentExists = true
			playlist.Data[i] = content
			break
		}
	}

	if !contentExists {
		playlist.Data = append(playlist.Data, content)
	}

	err = u.playlistRepo.Update(ctx, userId, playlistId, playlist)
	if err != nil {
		return nil, err
	}

	return playlist, nil
}

func (u *usecase) RemoveContent(ctx context.Context, userId string, playlistId string, contentId string) (*Playlist, error) {
	log.Printf("RemoveContent received userId: %s playlistId: %s contentId: %s", userId, playlistId, contentId)
	err := u.playlistRepo.RemoveContent(ctx, userId, playlistId, contentId)
	if err != nil {
		return nil, err
	}

	return u.playlistRepo.Get(ctx, userId, playlistId)
}

func (u *usecase) PartyMode(ctx context.Context, data *PartyModeRequest) (*Playlist, error) {
	log.Printf("PartyMode received request: %+v", data)

	switch {
	case data.Action.AddContent != nil:
		return u.processAddContent(ctx, data.PlaylistId, data.Action.AddContent)
	case data.Action.SkipContent != nil:
		return u.processSkipContent(ctx, data.Action.SkipContent)
	default:
		return nil, fmt.Errorf("invalid action")
	}
}

func (u *usecase) processAddContent(ctx context.Context, contentId string, data *AddContentEvent) (*Playlist, error) {
	log.Printf("PartyMode AddContent received request: %+v", data)
	playlist, err := u.playlistRepo.GetByType(ctx, "party", contentId)
	if err != nil {
		return nil, err
	}

	if playlist == nil {
		return nil, fmt.Errorf("playlist not found")
	}

	contentExists := false
	for i, c := range playlist.Data {
		if c.ID == data.ContentId {
			contentExists = true
			playlist.Data[i].Order = data.Order
			break
		}
	}

	if !contentExists {
		playlist.Data = append(playlist.Data, Content{
			ID:    data.ContentId,
			Order: data.Order,
		})
	}

	err = u.playlistRepo.AddContent(ctx, data.ContentId, playlist.Data)
	if err != nil {
		return nil, err
	}

	return playlist, nil
}

func (u *usecase) processSkipContent(ctx context.Context, data *SkipContentEvent) (*Playlist, error) {
	log.Printf("PartyMode SkipContent received request: %+v", data)
	return nil, fmt.Errorf("not implemented")
}
