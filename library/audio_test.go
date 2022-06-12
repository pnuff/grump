package library_test

import (
	"testing"

	"github.com/dhulihan/grump/library"
)

func TestShouldInclude(t *testing.T) {
	s, _ := library.NewLocalAudioShelf("empty")

	var tests = []struct {
		path     string
		expected bool
	}{
		{".", false},
		{"..", false},
		{"my-dir/", false},
		{"my-dir/99-11-tame_impala-glimmer.mp3", true},
		{"my-dir/99-11-tame_impala-glimmer.MP3", true},
		{"my-dir/99-11-tame_impala-glimmer.flac", true},
		{"my-dir/99-11-tame_impala-glimmer.wav", true},
		{"my-dir/foo.zip", false},
		{"my-dir/foo.pdf", false},
		{"my-dir/Cover.jpg", false},
	}

	for _, test := range tests {
		ret := s.ShouldInclude(test.path)
		if ret != test.expected {
			t.Errorf("for [%s] wanted [%t], got [%t]", test.path, test.expected, ret)
		}
	}
}

func TestIsPlaylist(t *testing.T) {
	var tests = []struct {
		path     string
		expected bool
	}{
		{".", false},
		{"..", false},
		{"my-dir/", false},
		{"my-dir/99-11-tame_impala-glimmer.m3u", true},
		{"my-dir/99-11-tame_impala-glimmer.M3U", true},
		{"my-dir/99-11-tame_impala-glimmer.flac", false},
		{"my-dir/99-11-tame_impala-glimmer.wav", false},
		{"my-dir/foo.zip", false},
		{"my-dir/foo.pdf", false},
		{"my-dir/Cover.jpg", false},
	
	}

	for _, test := range tests {
		ret := library.IsPlaylist(test.path)
		if ret != test.expected {
			t.Errorf("for [%s] wanted [%t], got [%t]", test.path, test.expected, ret)
		}
	}
}
