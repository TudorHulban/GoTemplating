package blogfile

import (
	"testing"

	"github.com/TudorHulban/GoTemplating/internal/articles"
	"github.com/stretchr/testify/require"
)

// test load article
// test load articles

// test save article
// test save articles

func TestBlogArticles(t *testing.T) {
	b, err := NewBlogFromArticles(articles.DefaultArticles()...)
	require.Nil(t, err)
	require.Nil(t, b.SaveBlogArticles())
}

func TestBlogFiles(t *testing.T) {
	files := make([]string, len(articles.DefaultArticles()))

	for i, art := range articles.DefaultArticles() {
		files[i] = art.SaveToFile
	}

	_, err := NewBlogFromFiles(files...)
	require.Nil(t, err)
}
