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
	playlist := &Playlist{
		UserID: data.UserID,
		Name:   data.Name,
	}
	err := u.playlistRepo.Create(ctx, playlist)
	if err != nil {
		return nil, err
	}

	return playlist, nil
}

func (u *usecase) List(ctx context.Context, data ListRequest) (*shared.Paginated[Playlist], error) {
	total, err := u.playlistRepo.Count(ctx, data.UserID, data.Query)
	if err != nil {
		return nil, err
	}

	if total == 0 {
		return nil, nil
	}

	playlists, err := u.playlistRepo.List(ctx, data.UserID, data.Query)
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

func (u *usecase) Get(ctx context.Context, userID string, playlistID string) (*Playlist, error) {
	log.Printf("Get received userID: %s playlistID: %s", userID, playlistID)
	got, err := u.playlistRepo.Get(ctx, userID, playlistID)
	if err != nil {
		return nil, err
	}

	return got, nil
}

/*
func (u *usecase) Manage(ctx context.Context, request *PlaylistRequest) (*Playlist, error) {
	log.Printf("Manage received request: %+v", request)
	switch request.Operation {
	case OperationKeyAdd:
		return u.add(ctx, request.ToEntity())
	case OperationKeyRemove:
		return u.removeContent(ctx, request.ToEntity())
	case OperationKeyGet:
		return u.Get(ctx, request.UserID, request.ID)
	default:
		return nil, errors.New("invalid operation")
	}
}





func (u *usecase) Remove(ctx context.Context, userID, playlistID string) error {
	log.Printf("Remove received userID: %s playlist: %s", userID, playlistID)
	return u.playlistRepo.Remove(ctx, userID, playlistID)
}

func (u *usecase) add(ctx context.Context, data *Playlist) (*Playlist, error) {
	log.Printf("Add received request: %+v", data)
	if data.ID != "" {
		if err := u.playlistRepo.Update(ctx, data.UserID, data.ID, data.Contents); err != nil {
			return nil, err
		}
		return u.playlistRepo.Get(ctx, data.UserID, data.ID)
	}

	if data.Name == "" {
		data.Name = "New Playlist"
	}

	if err := u.playlistRepo.Add(ctx, data); err != nil {
		return nil, err
	}
	return data, nil
}

func (u *usecase) removeContent(ctx context.Context, data *Playlist) (*Playlist, error) {
	log.Printf("Remove content received data: %+v", data)
	contentIDs := make([]string, 0)
	for _, content := range data.Contents {
		contentIDs = append(contentIDs, content.ID)
	}
	if err := u.playlistRepo.RemoveContent(ctx, data.UserID, data.ID, contentIDs); err != nil {
		return nil, err
	}
	return u.playlistRepo.Get(ctx, data.UserID, data.ID)
} */
