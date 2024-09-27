package main

import (
	"testing"

	"github.com/NhutHuyDev/gfind/core"
)

func TestBuildOption(t *testing.T) {
	args := []string{"", "-v", "-n", "-c"}

	got, _ := core.BuildOptions(args)

	if got.FindDontContain != true {
		t.Errorf("FindDontContain: got %+v, wanted %v", false, true)
	}

	if got.CountMode != true {
		t.Errorf("CountMode: got %+v, wanted %v", false, true)
	}

	if got.ShowLineNumber != true {
		t.Errorf("ShowLineNumber: got %+v, wanted %v", false, true)
	}

	if got.IsCaseSensitive != true {
		t.Errorf("IsCaseSensitive: got %+v, wanted %v", false, true)
	}

	if got.SkipOfflineFiles != true {
		t.Errorf("SkipOfflineFiles: got %+v, wanted %v", false, true)
	}

	if got.HelpMode != false {
		t.Errorf("HelpMode: got %+v, wanted %v", true, false)
	}
}
