// Copyright (c) 2020 Davi Villalva.
// a licensa pode ser encontrada no arquivo LICENSE na raíz do repositório.
// license can be found at the root of the repository in the LICENSE file.

package main

import (
	"math/rand"

	"github.com/hajimehoshi/ebiten"
)

type bola struct {
	centro     vet2
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

func (b *bola) inicia() {
	b.velocidade = vet2{rand.Float32(), -rand.Float32()}
	b.velocidade.setTamanho(6.0)
}

func (b *bola) estaVivo() bool { return true }

func (b *bola) deveTestarColisao() bool { return true }

func (b *bola) executaLogica() {
	if b.centro.x+b.raio > float32(telaLargura) {
		b.reflete(vet2{-1.0, 0.0})
	}
	if b.centro.x-b.raio < 0.0 {
		b.reflete(vet2{1.0, 0.0})
	}
	if b.centro.y+b.raio > float32(telaAltura) {
		b.reflete(vet2{0.0, -1.0})
	}
	if b.centro.y-b.raio < 0.0 {
		b.reflete(vet2{0.0, 1.0})
	}

	b.centro = b.centro.mais(b.velocidade)
}

func (b *bola) colidiuCom(colisao colisao) {
	delta := colisao.pontoDeContato.menos(b.centro.mais(colisao.normal.vezesEscalar(-b.raio)))
	b.centro = b.centro.mais(delta)
	b.reflete(colisao.normal)

	if tijolo, ok := colisao.objeto.(*tijolo); ok {
		tijolo.vida--
	}
}

func (b *bola) desenha(tela *ebiten.Image) {
	desenhaBola(tela, b.centro.x, b.centro.y, b.raio, 16)
}
