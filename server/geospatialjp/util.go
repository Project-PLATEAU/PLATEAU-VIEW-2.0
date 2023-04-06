package geospatialjp

import (
	"fmt"
	"net/http"
	"net/url"
	"path"
	"regexp"
	"strconv"

	"github.com/eukarya-inc/reearth-plateauview/server/geospatialjp/ckan"
	"github.com/pkg/errors"
	"github.com/samber/lo"
	"github.com/vincent-petithory/dataurl"
)

const (
	ResourceNameCityGML = "CityGML"
	ResourceNameAll     = "3D Tiles, MVT"
	ResourceNameCatalog = "データ目録"
	licenseDefaultID    = "plateau"
	licenseDefaultTitle = "PLATEAU Site Policy 「３．著作権について」に拠る"
	licenseDefaultURL   = "https://www.mlit.go.jp/plateau/site-policy/"
)

var reFileName = regexp.MustCompile(`^([0-9]+?)_(.+?)_([0-9]+?)_`)

func findResource(pkg *ckan.Package, name, format, desc, url string) (_ ckan.Resource, needUpdate bool) {
	r, found := lo.Find(pkg.Resources, func(r ckan.Resource) bool {
		return r.Name == name
	})
	if !found {
		r = ckan.Resource{
			PackageID:   pkg.ID,
			Name:        name,
			Format:      format,
			Description: desc,
		}
		needUpdate = true
	}
	if url != "" && r.URL != url {
		r.URL = url
		needUpdate = true
	}
	return r, needUpdate
}

func extractCityName(fn string) (string, string, int, error) {
	u, err := url.Parse(fn)
	if err != nil {
		return "", "", 0, err
	}

	base := path.Base(u.Path)
	s := reFileName.FindStringSubmatch(base)
	if s == nil {
		return "", "", 0, errors.New("invalid file name")
	}

	y, _ := strconv.Atoi(s[3])

	return s[1], s[2], y, nil
}

func packageFromCatalog(c *Catalog, org, pkgName string, private bool) ckan.Package {
	var thumbnailURL string
	if c.Thumbnail != nil {
		thumbnailURL = dataurl.New(c.Thumbnail, http.DetectContentType(c.Thumbnail)).String()
	}

	return ckan.Package{
		Name:            pkgName,
		Title:           c.Title,
		Private:         private || c.Public != "パブリック",
		Author:          c.Author,
		AuthorEmail:     c.AuthorEmail,
		Maintainer:      c.Maintainer,
		MaintainerEmail: c.MaintainerEmail,
		Notes:           c.Notes,
		Version:         c.Version,
		Tags: lo.Map(c.Tags, func(t string, _ int) ckan.Tag {
			return ckan.Tag{Name: t}
		}),
		OwnerOrg:         org,
		Restriction:      c.Restriction,
		Charge:           c.Charge,
		RegisterdDate:    c.RegisteredDate,
		LicenseAgreement: c.LicenseAgreement,
		LicenseTitle:     licenseDefaultTitle,
		LicenseURL:       licenseDefaultURL,
		LicenseID:        licenseDefaultID,
		Fee:              c.Fee,
		Area:             c.Area,
		Quality:          c.Quality,
		Emergency:        c.Emergency,
		URL:              c.Source,
		Spatial:          c.Spatial,
		ThumbnailURL:     thumbnailURL,
		// unused: URL: c.URL (empty), 組織: c.Organization (no field)
	}
}

func suffixFromSpec(s float64) string {
	vi := int(s)
	if vi <= 1 {
		return ""
	}

	return fmt.Sprintf("（v%d）", vi)
}

func datasetName(cityCode, cityName string, year int) string {
	datasetName := ""
	if isTokyo23ku(cityName) {
		if year <= 2020 {
			datasetName = fmt.Sprintf("plateau-%s", gspatialjpTokyo23ku)
		} else {
			datasetName = fmt.Sprintf("plateau-%s-%d", gspatialjpTokyo23ku, year)
		}
	} else {
		datasetName = fmt.Sprintf("plateau-%s-%s-%d", cityCode, cityName, year)
	}
	return datasetName
}

func isTokyo23ku(cityName string) bool {
	return cityName == citygmlTokyo23ku || cityName == citygmlTokyo23ku2 || cityName == gspatialjpTokyo23ku
}
