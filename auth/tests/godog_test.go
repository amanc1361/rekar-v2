package tests

import (
	"testing"

	"github.com/cucumber/godog"
	"rekar.ir/v2/auth/tests/step_definitions"
)

func TestFeatures(t *testing.T) {
	status := godog.TestSuite{
		Name:                "godogs",
		ScenarioInitializer: step_definitions.FeatureContext,
		Options: &godog.Options{
			Format: "pretty",
			Paths:  []string{"./features"},
		},
	}.Run()

	if status != 0 {
		t.Fatalf("non-zero status: %d", status)
	}
}
