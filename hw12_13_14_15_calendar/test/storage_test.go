package test

import (
	"context"
	"testing"

	"github.com/SomchaiSPB/otus-golang-hw/hw12_13_14_15_calendar/internal/storage"
	memorystorage "github.com/SomchaiSPB/otus-golang-hw/hw12_13_14_15_calendar/internal/storage/memory"
	"github.com/stretchr/testify/require"
)

func TestMemoryStorage(t *testing.T) {
	t.Run("storage create/list/get/delete test", func(t *testing.T) {
		sut := memorystorage.New()

		expected := storage.Event{
			Title:       "test",
			Duration:    1,
			Description: "description",
			NotifyIn:    2,
		}

		ctx := context.Background()

		actual, err := sut.CreateEvent(expected, &ctx)

		require.NoError(t, err)
		require.Equal(t, expected.Title, actual.Title)
		require.Equal(t, expected.Description, actual.Description)
		require.NotEmpty(t, actual.ID)

		actual = sut.GetEvent(actual.ID)

		require.NotNil(t, actual)
		require.Equal(t, expected.Title, actual.Title)

		updatedExpected := storage.Event{
			ID:          actual.ID,
			Title:       "new title",
			Duration:    10,
			Description: "new description",
			NotifyIn:    20,
			UserID:      actual.UserID,
		}

		updatedActual, err := sut.UpdateEvent(updatedExpected)

		require.NoError(t, err)
		require.Equal(t, updatedExpected.Title, updatedActual.Title)
		require.Equal(t, updatedExpected.Description, updatedActual.Description)
		require.Equal(t, updatedExpected.Duration, updatedActual.Duration)
		require.Equal(t, updatedExpected.NotifyIn, updatedActual.NotifyIn)

		events := sut.GetEvents()

		require.Equal(t, 1, len(events))
		require.NotEmpty(t, events)

		actualDel := sut.DeleteEvent(actual.ID)
		events = sut.GetEvents()

		require.NoError(t, actualDel)
		require.Equal(t, 0, len(events))

		err = sut.DeleteEvent(actual.ID)
		require.NoError(t, err)
	})
}

func TestSQLStorage(t *testing.T) {
	t.Run("storage create/list/get/delete test", func(t *testing.T) {
		//
	})
}