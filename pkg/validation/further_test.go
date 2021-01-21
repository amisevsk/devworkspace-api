package validation

import (
	"io/ioutil"
	"path/filepath"
	"testing"

	"github.com/devfile/api/v2/pkg/apis/workspaces/v1alpha2"
	"github.com/stretchr/testify/assert"

	"sigs.k8s.io/yaml"
)

type testCase struct {
	Name              string               `json:"name"`
	InputCommands     []v1alpha2.Command   `json:"commands"`
	InputComponents   []v1alpha2.Component `json:"components"`
	ExpectedErrRegexp *string              `json:"errRegexp"`
}

func loadTestCaseOrPanic(t *testing.T, testFilename string) testCase {
	testPath := filepath.Join("./testdata", testFilename)
	bytes, err := ioutil.ReadFile(testPath)
	if err != nil {
		t.Fatal(err)
	}
	var test testCase
	if err := yaml.Unmarshal(bytes, &test); err != nil {
		t.Fatal(err)
	}
	return test
}

func loadAllTestsOrPanic(t *testing.T) []testCase {
	files, err := ioutil.ReadDir("./testdata")
	if err != nil {
		t.Fatal(err)
	}
	var tests []testCase
	for _, file := range files {
		if file.IsDir() {
			continue
		}
		tests = append(tests, loadTestCaseOrPanic(t, file.Name()))
	}
	return tests
}

func TestValidateCommandsPartTwo(t *testing.T) {
	tests := loadAllTestsOrPanic(t)
	for _, tt := range tests {
		t.Run(tt.Name, func(t *testing.T) {
			// sanity check: input defines components
			assert.True(t, len(tt.InputCommands) > 0, "Test case defines workspace with no commands")
			err := ValidateCommands(tt.InputCommands, tt.InputComponents)
			if tt.ExpectedErrRegexp != nil && assert.Error(t, err) {
				assert.Regexp(t, *tt.ExpectedErrRegexp, err.Error(), "Error message should match")
			} else {
				assert.NoError(t, err, "Expected error to be nil")
			}
		})
	}
}
