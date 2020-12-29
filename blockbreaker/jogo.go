// Copyright (c) 2020 Davi Villalva.
// a licensa pode ser encontrada no arquivo LICENSE na raíz do repositório.
// license can be found at the root of the repository in the LICENSE file.

package main

import "github.com/hajimehoshi/ebiten"

type jogo struct {
	objetos               []objeto2d
	objetosATestarColisao []objeto2d
}

func (j *jogo) inicia() {
	for _, objeto := range j.objetos {
		objeto.inicia()
	}
}

func (j *jogo) adicionaObjeto(objeto objeto2d) {
	j.objetos = append(j.objetos, objeto)
}

func (j *jogo) adicionaObjetoATestarColisao(objeto objeto2d) {
	j.objetos = append(j.objetos, objeto)
	j.objetosATestarColisao = append(j.objetosATestarColisao, objeto)
}

func (j *jogo) Update() error {
	for _, objeto := range j.objetos {
		objeto.executaLogica()
	}
	for _, objetoATestarColisao := range j.objetosATestarColisao {
		// Jeito mais tosco de testar colisão, mas serve pra esse caso de uso
		for _, outroObjeto := range j.objetos {
			if outroObjeto != objetoATestarColisao {
				objetoATestarColisao.testaColisao(outroObjeto)
			}
		}
	}
	return nil
}

func (j *jogo) Draw(screen *ebiten.Image) {
	for _, objeto := range j.objetos {
		objeto.desenha(screen)
	}
}

func (j *jogo) Layout(outsideWidth, outsideHeight int) (screenWidth, screenHeight int) {
	return telaLargura, telaAltura
}
