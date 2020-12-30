// Copyright (c) 2020 Davi Villalva.
// a licensa pode ser encontrada no arquivo LICENSE na raíz do repositório.
// license can be found at the root of the repository in the LICENSE file.

package main

import (
	"image/color"

	"github.com/hajimehoshi/ebiten"
	"github.com/hajimehoshi/ebiten/ebitenutil"
)

var coresTijolo = []color.RGBA{
	{255, 0, 0, 255},
	{255, 255, 0, 255},
	{0, 255, 0, 255},
	{0, 255, 255, 255},
	{0, 0, 255, 255},
	{255, 0, 255, 255},
}

type tijolo struct {
	posicao, tamanho vet2
	vida             int
}

func (t *tijolo) inicia() {}

func (t *tijolo) executaLogica() {}

func (t *tijolo) estaVivo() bool {
	return t.vida >= 0
}

func (t *tijolo) testaColisao(objeto objeto2d) {}

func (t *tijolo) desenha(tela *ebiten.Image) {
	if !t.estaVivo() {
		return
	}

	ebitenutil.DrawRect(
		tela,
		float64(t.posicao.x), float64(t.posicao.y),
		float64(t.tamanho.x), float64(t.tamanho.y),
		coresTijolo[t.vida],
	)
}
