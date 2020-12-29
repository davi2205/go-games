// Copyright (c) 2020 Davi Villalva.
// a licensa pode ser encontrada no arquivo LICENSE na raíz do repositório.
// license can be found at the root of the repository in the LICENSE file.

package main

import "github.com/hajimehoshi/ebiten"

type objeto2d interface {
	inicia()
	executaLogica()
	testaColisao(objeto objeto2d)
	desenha(tela *ebiten.Image)
}
