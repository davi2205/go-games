// Copyright (c) 2020 Davi Villalva.
// a licensa pode ser encontrada no arquivo LICENSE na raíz do repositório.
// license can be found at the root of the repository in the LICENSE file.

package main

import (
	"math"

	"github.com/hajimehoshi/ebiten"
)

type bola struct {
	posicao    vet2
	velocidade vet2
	raio       float32
}

func (b *bola) reflete(normal vet2) {
	prodEscVelNormal := b.velocidade.produtoEscalar(normal)

	if prodEscVelNormal >= 0.0 {
		return
	}

	b.velocidade = b.velocidade.menos(normal.vezesEscalar(2.0 * prodEscVelNormal))
}

func (b *bola) inicia() {}

func (b *bola) executaLogica() {
	if b.posicao.x+b.raio > float32(telaLargura) {
		b.reflete(vet2{-1.0, 0.0})
	}
	if b.posicao.x-b.raio < 0.0 {
		b.reflete(vet2{1.0, 0.0})
	}
	if b.posicao.y+b.raio > float32(telaAltura) {
		b.reflete(vet2{0.0, -1.0})
	}
	if b.posicao.y-b.raio < 0.0 {
		b.reflete(vet2{0.0, 1.0})
	}

	b.posicao = b.posicao.mais(b.velocidade)
}

func (b *bola) testaColisao(objetoGenerico objeto2d) {
	switch objeto := objetoGenerico.(type) {
	case *jogador:
		pontoDeContato := vet2{
			limita(b.posicao.x, objeto.posicao.x, objeto.posicao.x+objeto.tamanho.x),
			limita(b.posicao.y, objeto.posicao.y, objeto.posicao.y+objeto.tamanho.y),
		}

		var (
			delta        = b.posicao.menos(pontoDeContato)
			deltaTamanho = delta.tamanho()
		)

		if deltaTamanho > b.raio || math.Abs(float64(deltaTamanho)) < 0.0001 {
			break
		}

		deltaNormalizado := delta.formaNormal()

		b.posicao = pontoDeContato.mais(deltaNormalizado.vezesEscalar(b.raio))
		b.reflete(deltaNormalizado)
	}
}

func (b *bola) desenha(tela *ebiten.Image) {
	desenhaBola(tela, b.posicao.x, b.posicao.y, b.raio, 16)
}
