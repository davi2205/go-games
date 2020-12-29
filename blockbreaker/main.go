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

	jogo := &jogo{}

	jogador := &jogador{
		posicao: vet2{telaLargura/2 - 40, telaAltura - 64},
		tamanho: vet2{80, 20},
	}
	jogo.adicionaObjeto(jogador)

	for i := 0; i < 2; i++ {
		bola := &bola{
			centro: vet2{30 + float32(i*20), telaAltura / 2},
			raio:   12,
		}
		jogo.adicionaObjetoATestarColisao(bola)
	}

	jogo.inicia()
	if err := ebiten.RunGame(jogo); err != nil {
		log.Fatal(err)
	}
}
