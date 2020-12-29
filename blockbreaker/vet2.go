// Copyright (c) 2020 Davi Villalva.
// a licensa pode ser encontrada no arquivo LICENSE na raíz do repositório.
// license can be found at the root of the repository in the LICENSE file.

package main

import "math"

type vet2 struct{ x, y float32 }

func (v vet2) mais(other vet2) vet2 {
	return vet2{v.x + other.x, v.y + other.y}
}

func (v vet2) menos(other vet2) vet2 {
	return vet2{v.x - other.x, v.y - other.y}
}

func (v vet2) vezesEscalar(scalar float32) vet2 {
	return vet2{v.x * scalar, v.y * scalar}
}

func (v vet2) produtoEscalar(other vet2) float32 {
	return v.x*other.x + v.y*other.y
}

// mesmo que: módulo
func (v vet2) tamanho() float32 {
	return float32(math.Sqrt(float64(v.x*v.x + v.y*v.y)))
}

func (v *vet2) setTamanho(tamanho float32) {
	direcao, _, ok := v.direcaoETamanho()

	if !ok {
		return
	}

	v.x = direcao.x * tamanho
	v.y = direcao.y * tamanho
}

func (v vet2) direcao() (vet2, bool) {
	tamanho := v.tamanho()

	if tamanho < 0.0001 {
		return vet2{}, false
	}

	return vet2{v.x / tamanho, v.y / tamanho}, true
}

func (v vet2) direcaoETamanho() (direcao vet2, tamanho float32, ok bool) {
	tamanho = v.tamanho()

	if tamanho < 0.0001 {
		return
	}

	direcao = vet2{v.x / tamanho, v.y / tamanho}
	ok = true
	return
}
