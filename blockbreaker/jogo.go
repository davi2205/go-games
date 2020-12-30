// Copyright (c) 2020 Davi Villalva.
// a licensa pode ser encontrada no arquivo LICENSE na raíz do repositório.
// license can be found at the root of the repository in the LICENSE file.

package main

import "github.com/hajimehoshi/ebiten"

type colisao struct {
	ocorreu        bool
	sujeito        objeto2d
	objeto         objeto2d
	pontoDeContato vet2
	normal         vet2
	distancia      float32
}

type objeto2d interface {
	inicia()
	estaVivo() bool
	deveTestarColisao() bool
	executaLogica()
	colidiuCom(objeto objeto2d, colisao colisao)
	desenha(tela *ebiten.Image)
}

type jogo struct {
	objetos      []objeto2d
	testaColisao func(a, b objeto2d) colisao
	colisoes     []colisao
}

func (j *jogo) inicia() {
	j.colisoes = make([]colisao, 0, 10)

	for _, objeto := range j.objetos {
		objeto.inicia()
	}
}

func (j *jogo) adicionaObjeto(objeto objeto2d) {
	j.objetos = append(j.objetos, objeto)
}

func (j *jogo) Update() error {
	j.colisoes = j.colisoes[:0]
	for _, objeto := range j.objetos {
		objeto.executaLogica()

		if objeto.deveTestarColisao() {
			for _, outroObjeto := range j.objetos {
				if outroObjeto != objeto {
					colisao := j.testaColisao(objeto, outroObjeto)

					if colisao.ocorreu {
						j.colisoes = append(j.colisoes, colisao)
					}
				}
			}
		}
	}
	for _, colisao := range j.colisoes {
		colisao.sujeito.colidiuCom(colisao.objeto, colisao)
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
