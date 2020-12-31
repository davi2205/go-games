// Copyright (c) 2020 Davi Villalva.
// a licensa pode ser encontrada no arquivo LICENSE na raíz do repositório.
// license can be found at the root of the repository in the LICENSE file.

package main

import (
	"log"

	"github.com/hajimehoshi/ebiten"
)

var mapa = [11][11]int{
	{10, 10, 10, 10, 10, 10, 10, 10, 10, 10, 10},
	{9, 9, 9, 9, 9, 9, 9, 9, 9, 9, 9},
	{8, 8, 8, 8, 8, 8, 8, 8, 8, 8, 8},
	{7, 7, 7, 7, 7, 7, 7, 7, 7, 7, 7},
	{6, 6, 6, 6, 6, 6, 6, 6, 6, 6, 6},
	{5, 5, 5, 5, 5, 5, 5, 5, 5, 5, 5},
	{4, 4, 4, 4, 4, 4, 4, 4, 4, 4, 4},
	{3, 3, 3, 3, 3, 3, 3, 3, 3, 3, 3},
	{2, 2, 2, 2, 2, 2, 2, 2, 2, 2, 2},
	{1, 1, 1, 1, 1, 1, 1, 1, 1, 1, 1},
	{0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0},
}

func main() {
	ebiten.SetWindowSize(telaLargura, telaAltura)
	ebiten.SetWindowTitle(tituloJanela)

	jogo := &jogo{
		testaColisao: testaColisao,
	}

	jogador := &jogador{
		posicao: vet2{telaLargura/2 - 40, telaAltura - 32},
		tamanho: vet2{80, 10},
	}
	jogo.adicionaObjeto(jogador)

	bola := &bola{
		centro: vet2{telaLargura / 2, telaAltura - 64},
		raio:   8,
	}
	jogo.adicionaObjeto(bola)

	var (
		margemTijolos float32 = 60.0
		margemTijolo  float32 = 4.0
		tijoloTamanho         = vet2{
			(float32(telaLargura) - margemTijolos*2) / float32(len(mapa[0])),
			(float32(telaAltura)/1.5 - margemTijolos*2) / float32(len(mapa)),
		}
	)

	for j := 0; j < len(mapa); j++ {
		for i := 0; i < len(mapa[j]); i++ {
			tijolo := &tijolo{
				posicao: vet2{
					margemTijolos + float32(margemTijolo/2) + float32(i)*tijoloTamanho.x,
					margemTijolos + float32(margemTijolo/2) + float32(j)*tijoloTamanho.y,
				},
				tamanho: vet2{
					tijoloTamanho.x - float32(margemTijolo),
					tijoloTamanho.y - float32(margemTijolo),
				},
				vida: mapa[j][i],
			}
			jogo.adicionaObjeto(tijolo)
		}
	}

	jogo.inicia()
	if err := ebiten.RunGame(jogo); err != nil {
		log.Fatal(err)
	}
}
