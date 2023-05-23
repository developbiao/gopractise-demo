package tdd

// Test-Driven Development is a method that can increase the quality of your code and reduce defects .
// The unit test is written before actual implementation.
// The unit test fails first; the aim is then to make it pass.
import "testing"

func TestVowelCount(t *testing.T) {
	expected := uint(5)
	actual := VowelCount("I love you")
	if actual != expected {
		t.Errorf("actual is %d, expected %d", actual, expected)
	}
}
