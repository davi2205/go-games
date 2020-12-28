// Copyright (c) 2020 Davi Villalva.
// a licensa pode ser encontrada no arquivo LICENSE na raíz do repositório.
// license can be found at the root of the repository in the LICENSE file.

package main

import (
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
	X, Y    float32
	Tamanho float32
}

func (j *Jogador) ExecutaLogica(cena *Cena) {
	if ebiten.IsKeyPressed(ebiten.KeyLeft) {
		j.X -= 5
	}
	if ebiten.IsKeyPressed(ebiten.KeyRight) {
		j.X += 5
	}
}

func (j *Jogador) Desenha(tela *ebiten.Image) {
	ebitenutil.DrawRect(
		tela,
		float64(j.X-j.Tamanho/2.0),
		float64(j.Y),
		float64(j.Tamanho),
		20.0,
		color.RGBA{255, 255, 255, 255},
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

	jogador := &Jogador{
		X:       telaLargura / 2,
		Y:       telaAltura - 64,
		Tamanho: 80,
	}

	jogo := new(Jogo)
	jogo.cena.AdicionaObjeto(jogador)

	if err := ebiten.RunGame(jogo); err != nil {
		log.Fatal(err)
	}
}
