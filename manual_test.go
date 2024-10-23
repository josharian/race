package race_test

import (
	"testing"

	"github.com/josharian/race"
)

func TestEnabled(t *testing.T) {
	t.Logf("race enabled: %v", race.Enabled)
}
