package main

import "testing"

func Test_gitProject(t *testing.T) {
	if userProject := gitProject("https://github.com/Jukie/cart"); userProject != "Jukie/cart" {
		t.Errorf("Expected %q, got %q", "Jukie/cart", userProject)
	}
	if userProject := gitProject("git@github.com:Jukie/cart.git"); userProject != "Jukie/cart" {
		t.Errorf("Expected %q, got %q", "Jukie/cart", userProject)
	}
	// TODO: recognize other Git hosts
}
