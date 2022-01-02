package resources

import (
	"bufio"
	"embed"
	"io/ioutil"
	"log"

	"fyne.io/fyne/v2"
)

//go:embed assets
var Assets embed.FS

type Resource string

const (
	LogoIcon Resource = "assets/logo.png"
)

func GetResource(res Resource, name string) *fyne.StaticResource {
	f, err := Assets.Open(string(res))
	r := bufio.NewReader(f)

	b, err := ioutil.ReadAll(r)
	if err != nil {
		log.Fatal(err)
	}
	return fyne.NewStaticResource(name, b)
}
