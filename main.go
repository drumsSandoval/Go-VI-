package main

import (
	"fmt"
	mult "./multimedia"
)


type contentWeb struct {
	multimedia [] mult.Multimedia
}

func (this *contentWeb) append(item mult.Multimedia) {
	this.multimedia = append(this.multimedia, item)
}

func (this *contentWeb) toString()  {
	for _, value := range this.multimedia {
		value.Show()
		fmt.Printf("\n")
	}
}

func main() {
	opc := ""
	bd := contentWeb{}
	for opc != "5" {
		fmt.Printf("\t\tMenu\n\t1.- Carturar Imagen\n\t2.- Capturar Audio\n\t3.- Capturar Video\n\t" +
			"4.- Mostrar Array\n\t5.- Salir\n\topc: ")
		fmt.Scanln(&opc)
		switch opc {
			case "1":
				img := mult.Image{}
				fmt.Print("\t\tCapturar Imagen\n\tTitulo: ")
				fmt.Scanln(&img.Title)
				fmt.Print("\n\tFormat: ")
				fmt.Scanln(&img.Format)
				fmt.Print("\n\tCanales: ")
				fmt.Scanln(&img.Channels)
				bd.append(&img)
			case "2":
				audio := mult.Audio{}
				fmt.Print("\t\tCapturar Audio\n\tTitulo: ")
				fmt.Scanln(&audio.Title)
				fmt.Print("\n\tFormat: ")
				fmt.Scanln(&audio.Format)
				fmt.Print("\n\tDuracion: ")
				fmt.Scanln(&audio.Duration)
				bd.append(&audio)
			case "3":
				video := mult.Video{}
				fmt.Print("\t\tCapturar Video\n\tTitulo: ")
				fmt.Scanln(&video.Title)
				fmt.Print("\n\tFormat: ")
				fmt.Scanln(&video.Format)
				fmt.Print("\n\tFrames: ")
				fmt.Scanln(&video.Frames)
				bd.append(&video)
			case "4":
				fmt.Println("\t\tMostrar Array")
				bd.toString()
			case "5":
				fmt.Println("Hasta pronto")
			default:
				fmt.Println("Opcion no valida\nIntentalo de nuevo")
		}
	}
}
