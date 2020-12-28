// Copyright (c) 2020 Davi Villalva.
// a licensa pode ser encontrada no arquivo LICENSE na raíz do repositório.
// license can be found at the root of the repository in the LICENSE file.

package main

import "github.com/hajimehoshi/ebiten"

type Jogo struct{}

func (j *Jogo) Update() error {
	return nil
}

func (j *Jogo) Draw(screen *ebiten.Image) {
	//
}

func (j *Jogo) Layout(outsideWidth, outsideHeight int) (screenWidth, screenHeight int) {
	return outsideWidth, outsideHeight
}
