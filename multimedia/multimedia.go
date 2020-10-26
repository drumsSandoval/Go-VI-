package multimedia

import "fmt"

type Multimedia interface {
	Show()
}

type Image struct {
	Title string
	Format string
	Channels string
}

func (this *Image) Show() {
	fmt.Printf("\t\t IMAGEN\n\tTitulo: %s \n\tFormato: %s \n\tCanales: %s", this.Title,this.Format,this.Channels)
}

type Audio struct {
	Title string
	Format string
	Duration string
}

func (this *Audio) Show() {
	fmt.Printf("\t\t AUDIO\n\tTitulo: %s \n\tFormato: %s \n\tDuracion: %s", this.Title,this.Format,this.Duration)
}

type Video struct {
	Title string
	Format string
	Frames int
}

func (this *Video) Show() {
	fmt.Printf("\t\t VIDEO\n\tTitulo: %s \n\tFormato: %s \n\tFRAMES: %s", this.Title,this.Format,this.Frames)
}