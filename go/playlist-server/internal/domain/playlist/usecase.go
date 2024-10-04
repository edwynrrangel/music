package playlist

import (
	"context"
	"errors"
	"log"
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

func (u *usecase) Get(ctx context.Context, userID string, playlistID string) (*Playlist, error) {
	log.Printf("Get received userID: %s playlistID: %s", userID, playlistID)
	got, err := u.playlistRepo.Get(ctx, userID, playlistID)
	if err != nil {
		return nil, err
	}

	return got, nil
}

func (u *usecase) List(ctx context.Context, userID string) ([]Playlist, error) {
	log.Printf("List received request: %s", userID)
	return u.playlistRepo.List(ctx, userID)
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
}
