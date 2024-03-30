package main

import "testing"

func TestContent(t *testing.T) {
	ContentCheck("https://www.hearkenundermoon.net/", []string{
		"182.150.5.184",
		"119.84.76.6",
	})
}
