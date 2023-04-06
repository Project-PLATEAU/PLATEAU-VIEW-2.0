package sdk

import (
	"github.com/eukarya-inc/reearth-plateauview/server/cms"
	"github.com/eukarya-inc/reearth-plateauview/server/fme"
)

type Config struct {
	CMSBase        string
	CMSToken       string
	CMSIntegration string
	FMEBaseURL     string
	FMEToken       string
	FMEResultURL   string
	Secret         string
}

type Services struct {
	CMS cms.Interface
	FME fme.Interface
}

func NewServices(conf Config) (*Services, error) {
	cms, err := cms.New(conf.CMSBase, conf.CMSToken)
	if err != nil {
		return nil, err
	}

	fme, err := fme.New(conf.FMEBaseURL, conf.FMEToken, conf.FMEResultURL)
	if err != nil {
		return nil, err
	}

	return &Services{CMS: cms, FME: fme}, nil
}
