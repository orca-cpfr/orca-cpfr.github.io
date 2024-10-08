package generator_test

import (
	"testing"

	"github.com/orca-cpfr/orca-cpfr.github.io/app/internal/generator"
	"github.com/stretchr/testify/require"
)

func TestSetLevel(t *testing.T) {
	require.Equal(t, 1, generator.GetPageLevel("index.html"))
	require.Equal(t, 2, generator.GetPageLevel("parent/index.html"))
	require.Equal(t, 5, generator.GetPageLevel("1/2/3/4/index.html"))
}

func TestSiteData_RelativePath(t *testing.T) {
	testCases := []struct {
		Name      string
		PageLevel int
		Path      string
		Expected  string
	}{
		{
			Name:      "level 1",
			PageLevel: 1,
			Path:      "index.html",
			Expected:  "index.html",
		},
		{
			Name:      "level 2",
			PageLevel: 2,
			Path:      "index.html",
			Expected:  "../index.html",
		},
		{
			Name:      "level 2",
			PageLevel: 5,
			Path:      "index.html",
			Expected:  "../../../../index.html",
		},
	}

	for _, tt := range testCases {
		t.Run(tt.Name, func(t *testing.T) {
			siteData := generator.SiteData{PageLevel: tt.PageLevel}
			require.Equal(t, tt.Expected, siteData.RelativePath(tt.Path))
		})
	}

}
