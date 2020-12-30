// Copyright (c) 2020 Davi Villalva.
// a licensa pode ser encontrada no arquivo LICENSE na raíz do repositório.
// license can be found at the root of the repository in the LICENSE file.

package main

import "math"

func testaColisao(sujeitoGenerico, objetoGenerico objeto2d) (colisao colisao, colidiu bool) {
	sujeito, ok := sujeitoGenerico.(*bola)

	if !ok {
		return
	}

	switch objeto := objetoGenerico.(type) {
	case *jogador:
		pontoDeContato := vet2{
			limita(sujeito.centro.x, objeto.posicao.x, objeto.posicao.x+objeto.tamanho.x),
			limita(sujeito.centro.y, objeto.posicao.y, objeto.posicao.y+objeto.tamanho.y),
		}

		direcao, tamanho, ok := sujeito.centro.menos(pontoDeContato).direcaoETamanho()

		if !ok {
			return
		}

		colisao.sujeito = sujeito
		colisao.objeto = objeto
		colisao.pontoDeContato = pontoDeContato
		colisao.normal = direcao
		colisao.distancia = tamanho
		colidiu = colisao.distancia <= sujeito.raio
	case *tijolo:
		pontoDeContato := vet2{
			limita(sujeito.centro.x, objeto.posicao.x, objeto.posicao.x+objeto.tamanho.x),
			limita(sujeito.centro.y, objeto.posicao.y, objeto.posicao.y+objeto.tamanho.y),
		}

		direcao, tamanho, ok := sujeito.centro.menos(pontoDeContato).direcaoETamanho()

		if !ok {
			return
		}

		colisao.sujeito = sujeito
		colisao.objeto = objeto
		colisao.pontoDeContato = pontoDeContato
		colisao.normal = direcao
		colisao.distancia = tamanho
		colidiu = colisao.distancia <= sujeito.raio
	case *bola:
		direcao, tamanho, ok := sujeito.centro.menos(objeto.centro).direcaoETamanho()

		if !ok {
			return
		}

		colisao.sujeito = sujeito
		colisao.objeto = objeto
		colisao.pontoDeContato = direcao.vezesEscalar(objeto.raio).mais(objeto.centro)
		colisao.normal = direcao
		colisao.distancia = float32(math.Abs(float64(tamanho - objeto.raio)))
		colidiu = colisao.distancia <= sujeito.raio
	}
	return
}
