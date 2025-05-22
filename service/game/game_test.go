package game

import (
	"fmt"
	"testing"
	"time"

	"github.com/stretchr/testify/assert"
)

type mockService struct {
	Service
}

func TestCalculatePrize(t *testing.T) {
	s := &mockService{}

	tests := []struct {
		number   int
		expected float32
	}{
		{954, float32(954 * 0.7)}, // >900
		{851, float32(851 * 0.5)}, // >600
		{449, float32(449 * 0.3)}, // >300
		{257, float32(257 * 0.1)}, // <=300
	}

	for _, tt := range tests {
		t.Run(
			prizeLabel(tt.number), func(t *testing.T) {
				got := s.calculatePrize(tt.number)
				assert.Equal(t, tt.expected, got)
			})
	}
}

func prizeLabel(number int) string {
	return fmt.Sprintf("number=%d", number)
}

func TestGetGameResult(t *testing.T) {
	s := &mockService{}

	tests := []struct {
		name      string
		userID    string
		number    int
		wantWin   bool
		wantPrize float32
	}{
		{"Even number, >900", "user1", 902, true, float32(902 * 0.7)},
		{"Even number, >600", "user2", 811, false, float32(811 * 0.5)},
		{"Even number, >300", "user3", 420, true, float32(420 * 0.3)},
		{"Even number, <=300", "user4", 210, true, float32(210 * 0.1)},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			result := s.getGameResult(tt.userID, tt.number)

			assert.Equal(t, tt.userID, result.UserID)
			assert.Equal(t, tt.wantWin, result.IsWin)
			assert.Equal(t, tt.number, result.Number)
			assert.Equal(t, tt.wantPrize, result.Prize)
			assert.WithinDuration(t, time.Now(), result.CreatedAt, time.Second)
		})
	}
}
