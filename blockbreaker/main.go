// Copyright (c) 2020 Davi Villalva.
// a licensa pode ser encontrada no arquivo LICENSE na raíz do repositório.
// license can be found at the root of the repository in the LICENSE file.

package main

import (
	"log"

	"github.com/hajimehoshi/ebiten"
)

type Objeto2d interface {
	Desenha(tela *ebiten.Image)
}

type Cena struct {
	objetos []Objeto2d
}

func (c *Cena) AdicionaObjeto(objeto Objeto2d) {
	c.objetos = append(c.objetos, objeto)
}

func (c *Cena) Desenha(tela *ebiten.Image) {
	for _, objeto := range c.objetos {
		objeto.Desenha(tela)
	}
}

type Vet2 struct {
	X, Y float32
}

type Jogador struct {
	Limites
}

type Jogo struct {
	cena Cena
}

func (j *Jogo) Update() error {
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
	if err := ebiten.RunGame(&Jogo{}); err != nil {
		log.Fatal(err)
	}
}
