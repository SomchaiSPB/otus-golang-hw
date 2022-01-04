package test

import (
	"github.com/SomchaiSPB/otus-golang-hw/hw12_13_14_15_calendar/internal/storage"
	memorystorage "github.com/SomchaiSPB/otus-golang-hw/hw12_13_14_15_calendar/internal/storage/memory"
	"github.com/stretchr/testify/require"
	"testing"
)

func TestStorage(t *testing.T) {

	t.Run("storage create/list/get/delete test", func(t *testing.T) {
		sut := memorystorage.New()

		expected := storage.Event{
			Title:       "test",
			Duration:    1,
			Description: "description",
			NotifyIn:    2,
			UserId:      "123",
		}

		actual, err := sut.CreateEvent(expected)

		require.NoError(t, err)
		require.Equal(t, expected.Title, actual.Title)
		require.Equal(t, expected.Description, actual.Description)
		require.NotEmpty(t, actual.ID)

		updatedExpected := storage.Event{
			ID:          actual.ID,
			Title:       "new title",
			Duration:    10,
			Description: "new description",
			NotifyIn:    20,
			UserId:      "123",
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
	})
}
