package bot_libs

import "testing"

func TestOutput(t *testing.T) {
	Output("s")
}

func TestStrip(t *testing.T) {
	if out, _ := stripCmd(".stock GOOG"); out != "GOOG" {
		t.Logf("Stock does not equal")
		t.Fail()
	}
	if out, _ := stripCmd(".calc 1 + 1"); out != "1 + 1" {
		t.Logf("Calc does not equal")
		t.Fail()
	}
	if _, err := stripCmd(".calc"); err == nil {
		t.Logf("stripCmd did not fail on no arguments")
		t.Fail()
	}
	if _, err := stripCmd(".test     "); err == nil {
		t.Logf("stripCmd did not fail on no arguments")
		t.Fail()
	}
}

func TestStock(t *testing.T) {
	if _, err := Stock(".stock GOOG"); err != nil {
		t.Logf("Stock returned error: %s", err)
		t.Fail()
	}
	if ticker, _ := Stock(".stock GOOG"); ticker == "" {
		t.Logf("Stock could not get ticker")
		t.Fail()
	}
}
