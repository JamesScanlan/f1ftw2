package json_converter

import "testing"

// TestReader tests the Reader function.
// It creates a Reader instance and checks if the Read method returns the expected result.
// If the result is not as expected, it reports an error using the Testing.T interface.
func TestReader(t *testing.T) {
	got := readJSONasText()
	want := `{"Grands_Prix":[{"Year":"2024","Races":[{"Grand_Prix":"Netherlands","Predictions":[{"Predictor":"Andrew Chadwick","Qualifying":"Lewis Hamilton","Race":"Carlos Sainz Jnr","Progression":"Williams"},{"Predictor":"James Scanlan","Qualifying":"Charles Leclerc","Race":"Max Verstappen","Progression":"Mercedes"},{"Predictor":"Matt Potter","Qualifying":"Alexander Albon","Race":"George Russell","Progression":"Alpine"}]}]}]}`

	if got != want {
		t.Errorf("got %q, wanted %q", got, want)
	}
}
