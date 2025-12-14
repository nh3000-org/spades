package getcards

import (
	"embed"

	"log"
)

//go:embed *.png
var CardsFS embed.FS

// Implement Fyne's fyne.Resource, loading from embedded files
type EmbededResource struct {
	path  string
	bytes []byte
}

// id is 2C.png for two of clubs, JH.png for jack of hearts...
func NewEmbeddedResource(path string) *EmbededResource {
	// translate
	bytes, err := CardsFS.ReadFile(path)
	if err != nil {
		log.Println("CardFS Not found", err, path)
	}
	return &EmbededResource{
		path:  path,
		bytes: bytes,
	}
}

func (resource *EmbededResource) Name() string {
	return resource.path
}

func (resource *EmbededResource) Content() []byte {
	return resource.bytes
}

/* func GetImage(id string) canvas.Image {

	img := canvas.NewImageFromResource(NewEmbeddedResource(id))

	return *img
} */
