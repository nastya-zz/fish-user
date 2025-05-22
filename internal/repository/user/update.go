package user

import (
	"context"
	"errors"
	"fmt"
	sq "github.com/Masterminds/squirrel"
	"github.com/jackc/pgx/v4"
	"log"
	"time"
	"user/internal/client/db"
	"user/internal/model"
)

func (r repo) UpdateProfile(ctx context.Context, updateUser *model.UpdateProfile) (*model.Profile, error) {
	const op = "user.UpdateProfile"

	log.Printf("updating user %+v", updateUser)
	uuId, err := model.GetUuid(updateUser.ID)

	returning := fmt.Sprintf("RETURNING %s, %s, %s, %s, %s, %s, %s, %s ", idColumn, nameColumn, emailColumn, isPublicColumn, createdAtColumn, experienceLevelColumn, isVerifiedColumn, avatarPathColumn)
	builder := sq.Update(tableName).
		PlaceholderFormat(sq.Dollar).
		Set(updatedAtColumn, time.Now()).
		Where(sq.Eq{idColumn: uuId}).
		Suffix(returning)

	keys := [4]string{"Name", "AvatarPath", "Bio", "IsPublic"}
	fields := updateUser.GetFieldMap(*updateUser)

	for _, key := range keys {
		if val, ok := fields[key]; ok && val != nil {
			builder = builder.Set(column[key], fields[key])
		}
	}

	query, args, err := builder.ToSql()
	if err != nil {
		return nil, fmt.Errorf("cannot build query update user with id: %s", updateUser.ID)
	}

	q := db.Query{
		Name:     op,
		QueryRaw: query,
	}
	var profile model.Profile
	err = r.db.DB().QueryRowContext(ctx, q, args...).Scan(&profile.ID, &profile.Name, &profile.Email, &profile.IsPublic, &profile.CreatedAt, &profile.ExperienceLevel, &profile.IsVerified, &profile.AvatarPath)
	if errors.Is(err, pgx.ErrNoRows) {
		log.Printf("error in update user with id: %s %+v", err, profile)

		return nil, fmt.Errorf("cannot update user with id: %s", updateUser.ID)
	}
	if err != nil {
		log.Printf("error in update user with id: %s", err)
		return nil, fmt.Errorf("cannot update user %w", err)
	}

	log.Printf("updating user %+v", updateUser)
	return &profile, nil
}
