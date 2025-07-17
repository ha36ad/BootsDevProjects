package main

import (
	"fmt"
	"testing"
	"time"

	"github.com/ha36ad/BootsDevProjects/pokedexcli/internal/pokecache"
)

func TestCleanInput(t *testing.T) {
	cases := []struct {
		input    string
		expected []string
	}{
		{
			input:    "  hello  world  ",
			expected: []string{"hello", "world"},
		},
		{
			input:    "go  is  cool",
			expected: []string{"go", "is", "cool"},
		},
		{
			input:    "   trim   and split  ",
			expected: []string{"trim", "and", "split"},
		},
		{
			input:    "",
			expected: []string{},
		},
	}

	for _, c := range cases {
		actual := cleanInput(c.input)

		if len(actual) != len(c.expected) {
			t.Errorf("For input %q: expected %d items, got %d", c.input, len(c.expected), len(actual))
			continue
		}

		for i := range actual {
			if actual[i] != c.expected[i] {
				t.Errorf("For input %q: expected %q, got %q", c.input, c.expected[i], actual[i])
			}
		}
	}
}
func TestAddGet(t *testing.T) {
	const interval = 5 * time.Second
	cache := pokecache.NewCache(interval)

	cases := []struct {
		key string
		val []byte
	}{
		{
			key: "https://example.com",
			val: []byte("testdata"),
		},
		{
			key: "https://example.com/path",
			val: []byte("moretestdata"),
		},
	}

	for i, c := range cases {
		t.Run(fmt.Sprintf("Test case %v", i), func(t *testing.T) {
			cache.Add(c.key, c.val)
			val, ok := cache.Get(c.key)
			if !ok {
				t.Errorf("expected to find key")
				return
			}
			if string(val) != string(c.val) {
				t.Errorf("expected value %s, got %s", string(c.val), string(val))
				return
			}
		})
	}
}

func TestReapLoop(t *testing.T) {
	const baseTime = 5 * time.Millisecond
	const waitTime = baseTime + 10*time.Millisecond
	cache := pokecache.NewCache(baseTime)

	cache.Add("https://example.com", []byte("testdata"))
	if _, ok := cache.Get("https://example.com"); !ok {
		t.Errorf("expected to find key")
		return
	}

	time.Sleep(waitTime)

	if _, ok := cache.Get("https://example.com"); ok {
		t.Errorf("expected key to be reaped")
	}
}
