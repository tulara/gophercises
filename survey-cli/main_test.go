package main_test

import (
	"os/exec"
	"testing"

	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/require"
)

func TestCLI_HappyPath(t *testing.T) {
	// survey-1-responses.csv has 5 submitted rows and 1 unsubmitted row.
	// go test sets the working directory to the package directory, so the
	// relative path resolves correctly.
	cmd := exec.Command("go", "run", ".", "survey-1-responses.csv")
	out, err := cmd.CombinedOutput()

	require.NoError(t, err, "program exited with error:\n%s", string(out))
	assert.Contains(t, string(out), "Total participants: 5")
	assert.Contains(t, string(out), "Participation percentage: 83.33%")
}
