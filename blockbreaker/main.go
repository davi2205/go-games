// Copyright (c) 2020 Davi Villalva.
// a licensa pode ser encontrada no arquivo LICENSE na raíz do repositório.
// license can be found at the root of the repository in the LICENSE file.

package main

import (
	"image"
	"image/color"
	"log"

	"github.com/hajimehoshi/ebiten"
	"github.com/hajimehoshi/ebiten/ebitenutil"
)

type Objeto2d interface {
	ExecutaLogica(cena *Cena)
	Desenha(tela *ebiten.Image)
}

type Cena struct {
	objetos []Objeto2d
}

func (c *Cena) AdicionaObjeto(objeto Objeto2d) {
	c.objetos = append(c.objetos, objeto)
}

func (c *Cena) ExecutaLogica() {
	for _, objeto := range c.objetos {
		objeto.ExecutaLogica(c)
	}
}

func (c *Cena) Desenha(tela *ebiten.Image) {
	for _, objeto := range c.objetos {
		objeto.Desenha(tela)
	}
}

type Jogador struct {
	Limites image.Rectangle
}

func (j *Jogador) ExecutaLogica(cena *Cena) {

}

func (j *Jogador) Desenha(tela *ebiten.Image) {
	ebitenutil.DrawRect(
		tela,
		float64(j.Limites.Min.X),
		float64(j.Limites.Min.Y),
		float64(j.Limites.Size().X),
		float64(j.Limites.Size().Y),
		color.RGBA{255, 0, 0, 255},
	)
}

type Jogo struct {
	cena Cena
}

func (j *Jogo) Update() error {
	j.cena.ExecutaLogica()
	return nil
}

func (j *Jogo) Draw(screen *ebiten.Image) {
	j.cena.Desenha(screen)
}

func (j *Jogo) Layout(outsideWidth, outsideHeight int) (screenWidth, screenHeight int) {
	return outsideWidth, outsideHeight
}

const (
	telaLargura  = 640
	telaAltura   = 480
	tituloJanela = "BlockBreaker by Davi Villalva"
)

func main() {
	ebiten.SetWindowSize(telaLargura, telaAltura)
	ebiten.SetWindowTitle(tituloJanela)

	jogo := new(Jogo)
	jogo.cena.AdicionaObjeto(new(Jogador))

	if err := ebiten.RunGame(jogo); err != nil {
		log.Fatal(err)
	}
}
