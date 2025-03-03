package buildah

import (
	"testing"

	"github.com/containers/storage"
	"github.com/openshift/imagebuilder"
	"github.com/stretchr/testify/assert"
)

func TestGetImageName(t *testing.T) {
	t.Parallel()
	tt := []struct {
		caseName string
		name     string
		names    []string
		expected string
	}{
		{"tagged image", "busybox1", []string{"docker.io/library/busybox:latest", "docker.io/library/busybox1:latest"}, "docker.io/library/busybox1:latest"},
		{"image name not in the resolved image names", "image1", []string{"docker.io/library/busybox:latest", "docker.io/library/busybox1:latest"}, "docker.io/library/busybox:latest"},
		{"resolved image with empty name list", "image1", []string{}, "image1"},
	}

	for _, tc := range tt {
		img := &storage.Image{Names: tc.names}
		res := getImageName(tc.name, img)
		if res != tc.expected {
			t.Errorf("test case '%s' failed: expected %#v but got %#v", tc.caseName, tc.expected, res)
		}
	}
}

func TestNoBaseImageSpecifierIsScratch(t *testing.T) {
	t.Parallel()
	assert.Equal(t, "scratch", imagebuilder.NoBaseImageSpecifier) // juuuuust in case
	assert.Equal(t, "scratch", BaseImageFakeName)
}
