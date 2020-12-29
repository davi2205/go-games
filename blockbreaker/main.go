// Copyright (c) 2020 Davi Villalva.
// a licensa pode ser encontrada no arquivo LICENSE na raíz do repositório.
// license can be found at the root of the repository in the LICENSE file.

package main

import (
	"log"

	"github.com/hajimehoshi/ebiten"
)

func main() {
	ebiten.SetWindowSize(telaLargura, telaAltura)
	ebiten.SetWindowTitle(tituloJanela)
	ebiten.SetWindowResizable(true)

	jogo := &jogo{}

	jogador := &jogador{
		posicao: vet2{telaLargura/2 - 40, telaAltura - 64},
		tamanho: vet2{80, 20},
	}
	jogo.adicionaObjeto(jogador)

	bola := &bola{
		posicao:    vet2{telaLargura / 2, telaAltura / 2},
		velocidade: vet2{6.0, -6.0},
		raio:       8,
	}
	jogo.adicionaObjetoATestarColisao(bola)

	if err := ebiten.RunGame(jogo); err != nil {
		log.Fatal(err)
	}
}
