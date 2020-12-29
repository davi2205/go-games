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

func (v vet2) tamanho() float32 {
	return float32(math.Sqrt(float64(v.x*v.x + v.y*v.y)))
}

func (v vet2) formaNormal() vet2 {
	tamanho := v.tamanho()
	return vet2{v.x / tamanho, v.y / tamanho}
}
