package backend

type EncodedAsset struct {
	Path     string
	Filename string
	MimeType string
	Data     (func() []byte)
	Width    int
	Height   int
}

func (decodedAsset *EncodedAsset) Url() string {
	return "http://localhost:3000/assets/" + decodedAsset.Path
}
