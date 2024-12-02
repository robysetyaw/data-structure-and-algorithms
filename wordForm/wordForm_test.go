package wordform

import "testing"

func TestWordForm(t *testing.T) {
	runTestCase(t, "peluang", "ulang", true, "should return true when 'ulang' can be formed from 'peluang'")
	runTestCase(t, "peluang", "gula", true, "should return true when 'gula' can be formed from 'peluang'")
	runTestCase(t, "peluang", "gulung", false, "should return false when 'gulung' cannot be formed from 'peluang'")
	runTestCase(t, "peluang", "ulangan", false, "should return false when 'ulangan' cannot be formed due to insufficient 'n'")
	runTestCase(t, "peluang", "xyz", false, "should return false when 'xyz' cannot be formed from 'peluang'")
	runTestCase(t, "", "ulang", false, "should return false when 'ulang' cannot be formed from an empty string")
	runTestCase(t, "ulang", "", true, "should return true when an empty string can be formed from 'ulang'")
	runTestCase(t, "ulang", "ulang", true, "should return true when 'ulang' matches exactly with 'ulang'")
}

func runTestCase(t *testing.T, kata1, kata2 string, want bool, description string) {
	t.Run(description, func(t *testing.T) {
		got := formWord(kata1, kata2)
		if got != want {
			t.Errorf("%s. canFormWord(%q, %q) = %v; want %v", description, kata1, kata2, got, want)
		}
	})
}