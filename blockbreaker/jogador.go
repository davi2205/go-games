// Copyright (c) 2020 Davi Villalva.
// a licensa pode ser encontrada no arquivo LICENSE na raíz do repositório.
// license can be found at the root of the repository in the LICENSE file.

package main

import (
	"image/color"

	"github.com/hajimehoshi/ebiten"
	"github.com/hajimehoshi/ebiten/ebitenutil"
)

type jogador struct {
	posicao, tamanho vet2
}

func (j *jogador) inicia() {}

func (j *jogador) estaVivo() bool { return true }

func (j *jogador) deveTestarColisao() bool { return false }

func (j *jogador) executaLogica() {
	if ebiten.IsKeyPressed(ebiten.KeyLeft) {
		j.posicao.x -= 10
	}
	if ebiten.IsKeyPressed(ebiten.KeyRight) {
		j.posicao.x += 10
	}

	if j.posicao.x < 0 {
		j.posicao.x = 0
	}
	if j.posicao.x+j.tamanho.x > float32(telaLargura) {
		j.posicao.x = float32(telaLargura) - j.tamanho.x
	}
}

func (j *jogador) colidiuCom(colisao colisao) {}

func (j *jogador) desenha(tela *ebiten.Image) {
	ebitenutil.DrawRect(
		tela,
		float64(j.posicao.x), float64(j.posicao.y),
		float64(j.tamanho.x), float64(j.tamanho.y),
		color.White,
	)
}
