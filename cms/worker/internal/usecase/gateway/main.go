package gateway

type Container struct {
	File File
	CMS  CMS
}

func NewGateway(f File, cms CMS) *Container {
	return &Container{
		File: f,
		CMS:  cms,
	}
}
