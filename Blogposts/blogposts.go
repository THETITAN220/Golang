package blogposts

import (
	blogposts "github.com/quii/learn-go-with-tests/tree/main/reading-files"
	"testing"
	"testing/fstest"
)

type Post struct {
}

func NewPostsFromFS(fileSystem fstest.MapFS) []Post {
	return nil
}
