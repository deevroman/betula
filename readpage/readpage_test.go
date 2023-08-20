package readpage

import (
	"embed"
	"golang.org/x/net/html"
	"net/url"
	"reflect"
	"testing"
)

//go:embed testdata/*
var testdata embed.FS

func TestTitles(t *testing.T) {
	table := map[string]string{
		"title outside head": "A title!",
		"title none":         "",
	}
	for name, expectedTitle := range table {
		file, _ := testdata.Open("testdata/" + name + ".html")
		doc, _ := html.Parse(file)
		data, err := findData("https://bouncepaw.com", titleWorkers, doc)
		if data.title != expectedTitle {
			t.Errorf("In ‘%s’, expected title ‘%s’, got ‘%s’. Error value is ‘%v’.",
				name, expectedTitle, data.title, err)
		}
	}
}

func TestHEntries(t *testing.T) {
	gutenberg := func() *url.URL {
		u, _ := url.ParseRequestURI("https://www.gutenberg.org/files/2701/2701-h/2701-h.htm#link2HCH0001")
		return u
	}()

	table := map[string]FoundData{
		"h-entry with p-name": {
			PostName:   "CHAPTER 1. Loomings.",
			BookmarkOf: nil,
			Tags:       nil,
			Mycomarkup: "",
			IsHFeed:    false,
		},

		"h-entry with p-name u-bookmark-of": {
			PostName:   "CHAPTER 1. Loomings.",
			BookmarkOf: gutenberg,
			Tags:       nil,
			Mycomarkup: "",
			IsHFeed:    false,
		},

		"h-feed with h-entries": {
			PostName:   "CHAPTER 1. Loomings.",
			BookmarkOf: nil,
			Tags:       nil,
			Mycomarkup: "",
			IsHFeed:    true,
		},
	}

	for name, expectedData := range table {
		file, _ := testdata.Open("testdata/" + name + ".html")
		doc, _ := html.Parse(file)
		data, err := findData("https://bouncepaw.com", repostWorkers, doc)
		data.docurl = nil
		if !reflect.DeepEqual(data, expectedData) {
			t.Errorf("In ‘%s’,\nwant %v,\ngot %v. Error value is ‘%v’.",
				name, expectedData, data, err)
		}
	}
}
