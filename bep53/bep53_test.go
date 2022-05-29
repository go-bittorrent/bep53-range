package bep53_test

import (
	"testing"

	"github.com/stretchr/testify/require"

	"github.com/go-bittorrent/bep53-range/bep53"
)

func TestParse(t *testing.T) {
	t.Parallel()

	t.Run("MultipleRanges", func(t *testing.T) {
		t.Parallel()

		parsed, err := bep53.Parse([]string{"1-4", "5", "3-3", "8-12"})
		require.NoError(t, err)
		require.EqualValues(t, []int{1, 2, 3, 4, 5, 3, 8, 9, 10, 11, 12}, parsed)
	})
	t.Run("OneValueFailed", func(t *testing.T) {
		t.Parallel()

		_, err := bep53.Parse([]string{"5%"})
		require.Error(t, err)
	})
	t.Run("RangeFailed", func(t *testing.T) {
		t.Parallel()

		_, err := bep53.Parse([]string{"5%-10"})
		require.Error(t, err)
		_, err = bep53.Parse([]string{"5-10%"})
		require.Error(t, err)
	})
	t.Run("ParseFailed", func(t *testing.T) {
		t.Parallel()

		_, err := bep53.Parse([]string{"5--10"})
		require.Error(t, err)
		_, err = bep53.Parse([]string{"-5-10"})
		require.Error(t, err)
	})
}

func TestCompose(t *testing.T) {
	t.Parallel()

	t.Run("OneValue", func(t *testing.T) {
		t.Parallel()

		composed := bep53.Compose([]int{1})
		require.EqualValues(t, []string{"1"}, composed)
	})
	t.Run("MultipleRanges", func(t *testing.T) {
		t.Parallel()

		composed := bep53.Compose([]int{1, 2, 3, 4, 5, 3, 8, 9, 10, 11, 12})
		require.EqualValues(t, []string{"1-5", "3", "8-12"}, composed)
	})
}
