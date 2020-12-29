// Copyright (c) 2020 Davi Villalva.
// a licensa pode ser encontrada no arquivo LICENSE na raíz do repositório.
// license can be found at the root of the repository in the LICENSE file.

package main

import (
	"image/color"
	"math"

	"github.com/hajimehoshi/ebiten"
)

func limita(valor, minimo, maximo float32) float32 {
	if valor < minimo {
		return minimo
	} else if valor > maximo {
		return maximo
	} else {
		return valor
	}
}

var (
	imagemVazia               = ebiten.NewImage(2, 2)
	bolaNumVertices           = 16
	bolaRaio                  = 8
	bolaIndices               = make([]uint16, 0, bolaNumVertices)
	bolaVertices              = make([]ebiten.Vertex, 0, bolaNumVertices+1)
	bolaVerticesTransformados = make([]ebiten.Vertex, 0, bolaNumVertices+1)
)

func init() {
	imagemVazia.Fill(color.White)
	reconstroiBola()
}

func reconstroiBola() {
	bolaIndices = bolaIndices[:0]
	bolaVertices = bolaVertices[:0]
	bolaVerticesTransformados = bolaVerticesTransformados[:0]

	for i := 0; i < bolaNumVertices; i++ {
		proporcao := float64(i) / float64(bolaNumVertices)
		cr := 0.0
		cg := 0.0
		cb := 0.0
		if proporcao < 1.0/3.0 {
			cb = 2 - 2*(proporcao*3)
			cr = 2 * (proporcao * 3)
		}
		if 1.0/3.0 <= proporcao && proporcao < 2.0/3.0 {
			cr = 2 - 2*(proporcao-1.0/3.0)*3
			cg = 2 * (proporcao - 1.0/3.0) * 3
		}
		if 2.0/3.0 <= proporcao {
			cg = 2 - 2*(proporcao-2.0/3.0)*3
			cb = 2 * (proporcao - 2.0/3.0) * 3
		}

		bolaIndices = append(
			bolaIndices,
			uint16(i),
			uint16(i+1)%uint16(bolaNumVertices),
			uint16(bolaNumVertices),
		)

		vertice := ebiten.Vertex{
			DstX:   float32(math.Cos(2 * math.Pi * proporcao)),
			DstY:   float32(math.Sin(2 * math.Pi * proporcao)),
			SrcX:   0,
			SrcY:   0,
			ColorR: float32(cr),
			ColorG: float32(cg),
			ColorB: float32(cb),
			ColorA: 1,
		}
		bolaVertices = append(bolaVertices, vertice)
		bolaVerticesTransformados = append(bolaVerticesTransformados, vertice)
	}

	vertice := ebiten.Vertex{
		DstX:   0,
		DstY:   0,
		SrcX:   0,
		SrcY:   0,
		ColorR: 1,
		ColorG: 1,
		ColorB: 1,
		ColorA: 1,
	}
	bolaVertices = append(bolaVertices, vertice)
	bolaVerticesTransformados = append(bolaVerticesTransformados, vertice)
}

func desenhaBola(tela *ebiten.Image, x, y, raio float32, numVertices int) {
	if numVertices != bolaNumVertices {
		bolaNumVertices = numVertices
		reconstroiBola()
	}
	for i, vertice := range bolaVertices {
		bolaVerticesTransformados[i].DstX = x + vertice.DstX*raio
		bolaVerticesTransformados[i].DstY = y + vertice.DstY*raio
	}
	tela.DrawTriangles(bolaVerticesTransformados, bolaIndices, imagemVazia, nil)
}
