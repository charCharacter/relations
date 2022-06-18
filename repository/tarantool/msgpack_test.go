package tarantool

import (
	"bytes"
	"testing"
	"time"

	"github.com/google/uuid"
	"github.com/stretchr/testify/require"
	"gopkg.in/vmihailenco/msgpack.v2"
)

func TestAlbum(t *testing.T) {
	t.Parallel()
	var b bytes.Buffer
	var want, got *album

	t.Run("EncodeMsgpack", func(t *testing.T) {
		enc := msgpack.NewEncoder(&b)
		now := time.Now()
		id, fileID_1, fileID_2 := newUUID(t), newUUID(t), newUUID(t)
		files := []uuid.UUID{fileID_1, fileID_2}
		want = &album{
			ID:        id,
			ChangedAt: now,
			Name:      "Name",
			Icon:      "Icon",
			UserID:    2000,
			Files:     files,
		}
		require.NoError(t, want.EncodeMsgpack(enc))
	})
	t.Run("DecodeMsgpack", func(t *testing.T) {
		dec := msgpack.NewDecoder(&b)
		got = new(album)
		require.NoError(t, got.DecodeMsgpack(dec))
	})

	require.WithinDuration(t, want.ChangedAt, got.ChangedAt, time.Second)
	want.ChangedAt, got.ChangedAt = time.Time{}, time.Time{}
	require.Equal(t, want, got)
}

func newUUID(t *testing.T) uuid.UUID {
	id, err := uuid.NewUUID()
	require.NoError(t, err)
	return id
}
