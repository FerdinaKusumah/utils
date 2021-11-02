package url

import (
	"github.com/stretchr/testify/assert"
	"testing"
)

// TestEscapeStringFromUrl ...
func TestEscapeStringFromUrl(t *testing.T) {
	t.Log(`Test escape url "https://foo.bar?name=what is your name"`)
	{
		text := "https://foo.bar?name=what is your name"
		res := EscapeStringFromUrl(text)
		assert.Equal(t, res, "https%3A%2F%2Ffoo.bar%3Fname%3Dwhat+is+your+name", "value is not equals")
	}
}

// TestUnescapeStringFromUrl ...
func TestUnescapeStringFromUrl(t *testing.T) {
	t.Log(`Test unescape url "https://foo.bar?name=what is your name"`)
	{
		text := "https%3A%2F%2Ffoo.bar%3Fname%3Dwhat+is+your+name"
		res := UnescapeStringFromUrl(text)
		assert.Equal(t, res, "https://foo.bar?name=what is your name", "value is not equals")
	}
}

func TestIsUrlValid(t *testing.T) {
	t.Log(`Test valid url for: "http://www.example.com:80/path/to/myfile.html?key1=value1&key2=value2#somehwereInTheDocument"`)
	{
		text := "http://www.example.com:80/path/to/myfile.html?key1=value1&key2=value2#somehwereInTheDocument"
		res := IsUrlValid(text)
		assert.Equal(t, res, true, "value is not equals")
	}

	t.Log(`Test valid url for: "https://www.example.com:80/path/to/myfile.html?key1=value1&key2=value2#somehwereInTheDocument"`)
	{
		text := "https://www.example.com:80/path/to/myfile.html?key1=value1&key2=value2#somehwereInTheDocument"
		res := IsUrlValid(text)
		assert.Equal(t, res, true, "value is not equals")
	}

	t.Log(`Test valid url for: "https://www.example.com:80/path/to/myfile.html?key1=value1&key2=value2#somehwereInTheDocument"`)
	{
		text := "https:/www.example.com:80/path/to/myfile.html?key1=value1&key2=value2#somehwereInTheDocument"
		res := IsUrlValid(text)
		assert.Equal(t, res, true, "value is not equals")
	}

	t.Log(`Test valid url for: "https//www.example.com:80/path/to/myfile.html?key1=value1&key2=value2#somehwereInTheDocument"`)
	{
		text := "https://www.example.com:80/path/to/myfile.html?key1=value1&key2=value2#somehwereInTheDocument"
		res := IsUrlValid(text)
		assert.Equal(t, res, true, "value is not equals")
	}

	t.Log(`Test valid url for: "//www.example.com:80/path/to/myfile.html?key1=value1&key2=value2#somehwereInTheDocument"`)
	{
		text := "//www.example.com:80/path/to/myfile.html?key1=value1&key2=value2#somehwereInTheDocument"
		res := IsUrlValid(text)
		assert.Equal(t, res, true, "value is not equals")
	}
}

func TestGetHostFromUrl(t *testing.T) {
	t.Log(`Test get host url for: "http://www.example.com:80/path/to/myfile.html?key1=value1&key2=value2#somehwereInTheDocument"`)
	{
		text := "http://www.example.com:80/path/to/myfile.html?key1=value1&key2=value2#somehwereInTheDocument"
		res := GetHostFromUrl(text)
		assert.Equal(t, res, "www.example.com:80", "value is not equals")
	}

	t.Log(`Test get host url for: "https://www.example.com:80/path/to/myfile.html?key1=value1&key2=value2#somehwereInTheDocument"`)
	{
		text := "https://www.example.com:80/path/to/myfile.html?key1=value1&key2=value2#somehwereInTheDocument"
		res := GetHostFromUrl(text)
		assert.Equal(t, res, "www.example.com:80", "value is not equals")
	}

	t.Log(`Test get host url for: "https:/www.example.com:80/path/to/myfile.html?key1=value1&key2=value2#somehwereInTheDocument"`)
	{
		text := "https:/www.example.com:80/path/to/myfile.html?key1=value1&key2=value2#somehwereInTheDocument"
		res := GetHostFromUrl(text)
		assert.Equal(t, res, "", "value is not equals")
	}

	t.Log(`Test get host url for: "https//www.example.com:80/path/to/myfile.html?key1=value1&key2=value2#somehwereInTheDocument"`)
	{
		text := "https//www.example.com:80/path/to/myfile.html?key1=value1&key2=value2#somehwereInTheDocument"
		res := GetHostFromUrl(text)
		assert.Equal(t, res, "", "value is not equals")
	}

	t.Log(`Test get host url for: "//www.example.com:80/path/to/myfile.html?key1=value1&key2=value2#somehwereInTheDocument"`)
	{
		text := "//www.example.com:80/path/to/myfile.html?key1=value1&key2=value2#somehwereInTheDocument"
		res := GetHostFromUrl(text)
		assert.Equal(t, res, "www.example.com:80", "value is not equals")
	}
}

func TestGetSchemeFromUrl(t *testing.T) {
	t.Log(`Test get scheme url for: "http://www.example.com:80/path/to/myfile.html?key1=value1&key2=value2#somehwereInTheDocument"`)
	{
		text := "http://www.example.com:80/path/to/myfile.html?key1=value1&key2=value2#somehwereInTheDocument"
		res := GetSchemeFromUrl(text)
		assert.Equal(t, res, "http", "value is not equals")
	}

	t.Log(`Test get scheme url for: "https://www.example.com:80/path/to/myfile.html?key1=value1&key2=value2#somehwereInTheDocument"`)
	{
		text := "https://www.example.com:80/path/to/myfile.html?key1=value1&key2=value2#somehwereInTheDocument"
		res := GetSchemeFromUrl(text)
		assert.Equal(t, res, "https", "value is not equals")
	}

	t.Log(`Test get scheme url for: "https:/www.example.com:80/path/to/myfile.html?key1=value1&key2=value2#somehwereInTheDocument"`)
	{
		text := "https:/www.example.com:80/path/to/myfile.html?key1=value1&key2=value2#somehwereInTheDocument"
		res := GetSchemeFromUrl(text)
		assert.Equal(t, res, "https", "value is not equals")
	}

	t.Log(`Test get scheme url for: "https//www.example.com:80/path/to/myfile.html?key1=value1&key2=value2#somehwereInTheDocument"`)
	{
		text := "https//www.example.com:80/path/to/myfile.html?key1=value1&key2=value2#somehwereInTheDocument"
		res := GetSchemeFromUrl(text)
		assert.Equal(t, res, "", "value is not equals")
	}

	t.Log(`Test get scheme url for: "//www.example.com:80/path/to/myfile.html?key1=value1&key2=value2#somehwereInTheDocument"`)
	{
		text := "//www.example.com:80/path/to/myfile.html?key1=value1&key2=value2#somehwereInTheDocument"
		res := GetSchemeFromUrl(text)
		assert.Equal(t, res, "", "value is not equals")
	}
}

func TestGetPathFromUrl(t *testing.T) {
	t.Log(`Test get scheme url for: "http://www.example.com:80/path/to/myfile.html?key1=value1&key2=value2#somehwereInTheDocument"`)
	{
		text := "http://www.example.com:80/path/to/myfile.html?key1=value1&key2=value2#somehwereInTheDocument"
		res := GetPathFromUrl(text)
		assert.Equal(t, res, "/path/to/myfile.html", "value is not equals")
	}

	t.Log(`Test get scheme url for: "https://www.example.com:80/path/to/myfile.html?key1=value1&key2=value2#somehwereInTheDocument"`)
	{
		text := "https://www.example.com:80/path/to/myfile.html?key1=value1&key2=value2#somehwereInTheDocument"
		res := GetPathFromUrl(text)
		assert.Equal(t, res, "/path/to/myfile.html", "value is not equals")
	}

	t.Log(`Test get scheme url for: "https:/www.example.com:80/path/to/myfile.html?key1=value1&key2=value2#somehwereInTheDocument"`)
	{
		text := "https:/www.example.com:80/path/to/myfile.html?key1=value1&key2=value2#somehwereInTheDocument"
		res := GetPathFromUrl(text)
		assert.Equal(t, res, "/www.example.com:80/path/to/myfile.html", "value is not equals")
	}

	t.Log(`Test get scheme url for: "https//www.example.com:80/path/to/myfile.html?key1=value1&key2=value2#somehwereInTheDocument"`)
	{
		text := "https//www.example.com:80/path/to/myfile.html?key1=value1&key2=value2#somehwereInTheDocument"
		res := GetPathFromUrl(text)
		assert.Equal(t, res, "https//www.example.com:80/path/to/myfile.html", "value is not equals")
	}

	t.Log(`Test get scheme url for: "//www.example.com:80/path/to/myfile.html?key1=value1&key2=value2#somehwereInTheDocument"`)
	{
		text := "//www.example.com:80/path/to/myfile.html?key1=value1&key2=value2#somehwereInTheDocument"
		res := GetPathFromUrl(text)
		assert.Equal(t, res, "/path/to/myfile.html", "value is not equals")
	}
}

func TestGetRawPathFromUrl(t *testing.T) {
	t.Log(`Test get scheme url for: "http://www.example.com:80/path/to/myfile.html?key1=value1&key2=value2#somehwereInTheDocument"`)
	{
		text := "http://www.example.com:80/path/to/myfile.html?key1=value1&key2=value2#somehwereInTheDocument"
		res := GetRawPathFromUrl(text)
		assert.Equal(t, res, "/path/to/myfile.html", "value is not equals")
	}

	t.Log(`Test get scheme url for: "https://www.example.com:80/path/to/myfile.html?key1=value1&key2=value2#somehwereInTheDocument"`)
	{
		text := "https://www.example.com:80/path/to/myfile.html?key1=value1&key2=value2#somehwereInTheDocument"
		res := GetRawPathFromUrl(text)
		assert.Equal(t, res, "/path/to/myfile.html", "value is not equals")
	}

	t.Log(`Test get scheme url for: "https:/www.example.com:80/path/to/myfile.html?key1=value1&key2=value2#somehwereInTheDocument"`)
	{
		text := "https:/www.example.com:80/path/to/myfile.html?key1=value1&key2=value2#somehwereInTheDocument"
		res := GetRawPathFromUrl(text)
		assert.Equal(t, res, "/www.example.com:80/path/to/myfile.html", "value is not equals")
	}

	t.Log(`Test get scheme url for: "https//www.example.com:80/path/to/myfile.html?key1=value1&key2=value2#somehwereInTheDocument"`)
	{
		text := "https//www.example.com:80/path/to/myfile.html?key1=value1&key2=value2#somehwereInTheDocument"
		res := GetRawPathFromUrl(text)
		assert.Equal(t, res, "https//www.example.com:80/path/to/myfile.html", "value is not equals")
	}

	t.Log(`Test get scheme url for: "//www.example.com:80/path/to/myfile.html?key1=value1&key2=value2#somehwereInTheDocument"`)
	{
		text := "//www.example.com:80/path/to/myfile.html?key1=value1&key2=value2#somehwereInTheDocument"
		res := GetRawPathFromUrl(text)
		assert.Equal(t, res, "/path/to/myfile.html", "value is not equals")
	}
}

func TestGetRawQueryFromUrl(t *testing.T) {
	t.Log(`Test get scheme url for: "http://www.example.com:80/path/to/myfile.html?key1=value1&key2=value2#somehwereInTheDocument"`)
	{
		text := "http://www.example.com:80/path/to/myfile.html?key1=value1&key2=value2#somehwereInTheDocument"
		res := GetRawQueryFromUrl(text)
		assert.Equal(t, res, "key1=value1&key2=value2", "value is not equals")
	}

	t.Log(`Test get scheme url for: "https://www.example.com:80/path/to/myfile.html?key1=value1&key2=value2#somehwereInTheDocument"`)
	{
		text := "https://www.example.com:80/path/to/myfile.html?key1=value1&key2=value2#somehwereInTheDocument"
		res := GetRawQueryFromUrl(text)
		assert.Equal(t, res, "key1=value1&key2=value2", "value is not equals")
	}

	t.Log(`Test get scheme url for: "https:/www.example.com:80/path/to/myfile.html?key1=value1&key2=value2#somehwereInTheDocument"`)
	{
		text := "https:/www.example.com:80/path/to/myfile.html?key1=value1&key2=value2#somehwereInTheDocument"
		res := GetRawQueryFromUrl(text)
		assert.Equal(t, res, "key1=value1&key2=value2", "value is not equals")
	}

	t.Log(`Test get scheme url for: "https//www.example.com:80/path/to/myfile.html?key1=value1&key2=value2#somehwereInTheDocument"`)
	{
		text := "https//www.example.com:80/path/to/myfile.html?key1=value1&key2=value2#somehwereInTheDocument"
		res := GetRawQueryFromUrl(text)
		assert.Equal(t, res, "key1=value1&key2=value2", "value is not equals")
	}

	t.Log(`Test get scheme url for: "//www.example.com:80/path/to/myfile.html?key1=value1&key2=value2#somehwereInTheDocument"`)
	{
		text := "//www.example.com:80/path/to/myfile.html?key1=value1&key2=value2#somehwereInTheDocument"
		res := GetRawQueryFromUrl(text)
		assert.Equal(t, res, "key1=value1&key2=value2", "value is not equals")
	}
}
