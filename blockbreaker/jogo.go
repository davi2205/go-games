// Copyright (c) 2020 Davi Villalva.
// a licensa pode ser encontrada no arquivo LICENSE na raíz do repositório.
// license can be found at the root of the repository in the LICENSE file.

package main

import "github.com/hajimehoshi/ebiten"

type colisao struct {
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
	colidiuCom(colisao colisao)
	desenha(tela *ebiten.Image)
}

type jogo struct {
	objetos         []objeto2d
	testaColisao    func(objeto2d, objeto2d) (colisao, bool)
	colisoes        []colisao
	indicesAExcluir []int
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
	j.indicesAExcluir = j.indicesAExcluir[:0]

	for indice, objeto := range j.objetos {
		objeto.executaLogica()

		if objeto.deveTestarColisao() {
			for _, outroObjeto := range j.objetos {
				if outroObjeto != objeto {
					colisao, ok := j.testaColisao(objeto, outroObjeto)
					if ok {
						j.colisoes = append(j.colisoes, colisao)
					}
				}
			}
		}

		if !objeto.estaVivo() {
			j.indicesAExcluir = append(j.indicesAExcluir, indice)
		}
	}

	for _, colisao := range j.colisoes {
		colisao.sujeito.colidiuCom(colisao)
	}

	for indice, indiceEmObjetos := range j.indicesAExcluir {
		ultimoIndiceEmObjetos := len(j.objetos) - (1 + indice)
		j.objetos[indiceEmObjetos] = j.objetos[ultimoIndiceEmObjetos]
		j.objetos[ultimoIndiceEmObjetos] = nil
	}
	j.objetos = j.objetos[:len(j.objetos)-len(j.indicesAExcluir)]

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
