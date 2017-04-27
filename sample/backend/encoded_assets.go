package backend

func GetAsset(path string) *EncodedAsset {
	switch path {
	case "assets/avatars/justice.png":
		return &EncodedAsset{
			Path:     "assets/avatars/justice.png",
			Filename: "justice.png",
			MimeType: "image/png",
			Width:    72,
			Height:   72,
			Data: func() []byte {
				asset, _ := _bindata["assets/assets/avatars/justice.png"]()
				if asset == nil {
					return nil
				}
				return asset.bytes
			}}
	}
	return nil
}
