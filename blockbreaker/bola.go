// Copyright (c) 2020 Davi Villalva.
// a licensa pode ser encontrada no arquivo LICENSE na raíz do repositório.
// license can be found at the root of the repository in the LICENSE file.

package main

import (
	"math"
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
	b.velocidade.setTamanho(7.0)
}

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

func (b *bola) testaColisao(objetoGenerico objeto2d) {
	var (
		pontoDeContato vet2
		normal         vet2
		distancia      float32
	)

	switch objeto := objetoGenerico.(type) {
	case *jogador:
		pontoDeContato = vet2{
			limita(b.centro.x, objeto.posicao.x, objeto.posicao.x+objeto.tamanho.x),
			limita(b.centro.y, objeto.posicao.y, objeto.posicao.y+objeto.tamanho.y),
		}

		direcao, tamanho, ok := b.centro.menos(pontoDeContato).direcaoETamanho()

		if !ok {
			return
		}

		normal = direcao
		distancia = tamanho
	case *bola:
		direcao, tamanho, ok := b.centro.menos(objeto.centro).direcaoETamanho()

		if !ok {
			return
		}

		pontoDeContato = direcao.vezesEscalar(objeto.raio).mais(objeto.centro)
		normal = direcao
		distancia = float32(math.Abs(float64(tamanho - objeto.raio)))
	default:
		return
	}

	if distancia > b.raio {
		return
	}

	respostaColisao := pontoDeContato.menos(b.centro.mais(normal.vezesEscalar(-b.raio)))

	b.centro = b.centro.mais(respostaColisao)
	b.reflete(normal)
}

func (b *bola) desenha(tela *ebiten.Image) {
	desenhaBola(tela, b.centro.x, b.centro.y, b.raio, 16)
}
