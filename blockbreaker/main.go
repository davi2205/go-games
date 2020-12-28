// Copyright (c) 2020 Davi Villalva.
// a licensa pode ser encontrada no arquivo LICENSE na raíz do repositório.
// license can be found at the root of the repository in the LICENSE file.

package main

import (
	"image/color"
	"log"
	"math"

	"github.com/hajimehoshi/ebiten"
	"github.com/hajimehoshi/ebiten/ebitenutil"
)

var emptyImage = ebiten.NewImage(3, 3)

type Objeto2d interface {
	Inicia(jogo *Jogo)
	ExecutaLogica(jogo *Jogo)
	Desenha(tela *ebiten.Image)
}

type Jogador struct {
	X, Y    float32
	Tamanho float32
}

func (j *Jogador) Inicia(jogo *Jogo) {}

func (j *Jogador) ExecutaLogica(jogo *Jogo) {
	if ebiten.IsKeyPressed(ebiten.KeyLeft) {
		j.X -= 5
	}
	if ebiten.IsKeyPressed(ebiten.KeyRight) {
		j.X += 5
	}

	metade := j.Tamanho / 2

	if j.X-metade < 0 {
		j.X = metade
	}
	if j.X+metade > float32(jogo.telaLargura) {
		j.X = float32(jogo.telaLargura) - metade
	}
}

func (j *Jogador) Desenha(tela *ebiten.Image) {
	ebitenutil.DrawRect(
		tela,
		float64(j.X-j.Tamanho/2.0),
		float64(j.Y),
		float64(j.Tamanho),
		20.0,
		color.RGBA{255, 255, 255, 255},
	)
}

type GraficoBola struct {
	indices               []uint16
	vertices              []ebiten.Vertex
	verticesTransformados []ebiten.Vertex
}

func NovoGraficoBola(numVertices int, raio float32) *GraficoBola {
	var (
		indices               = []uint16{}
		vertices              = []ebiten.Vertex{}
		verticesTransformados = []ebiten.Vertex{}
	)

	for i := 0; i < numVertices; i++ {
		racio := float64(i) / float64(numVertices)
		cr := 0.0
		cg := 0.0
		cb := 0.0
		if racio < 1.0/3.0 {
			cb = 2 - 2*(racio*3)
			cr = 2 * (racio * 3)
		}
		if 1.0/3.0 <= racio && racio < 2.0/3.0 {
			cr = 2 - 2*(racio-1.0/3.0)*3
			cg = 2 * (racio - 1.0/3.0) * 3
		}
		if 2.0/3.0 <= racio {
			cg = 2 - 2*(racio-2.0/3.0)*3
			cb = 2 * (racio - 2.0/3.0) * 3
		}

		indices = append(indices, uint16(i), uint16(i+1)%uint16(numVertices), uint16(numVertices))

		vertice := ebiten.Vertex{
			DstX:   float32(float64(raio) * math.Cos(2*math.Pi*racio)),
			DstY:   float32(float64(raio) * math.Sin(2*math.Pi*racio)),
			SrcX:   0,
			SrcY:   0,
			ColorR: float32(cr),
			ColorG: float32(cg),
			ColorB: float32(cb),
			ColorA: 1,
		}
		vertices = append(vertices, vertice)
		verticesTransformados = append(verticesTransformados, vertice)
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
	vertices = append(vertices, vertice)
	verticesTransformados = append(verticesTransformados, vertice)

	return &GraficoBola{
		indices:               indices,
		vertices:              vertices,
		verticesTransformados: verticesTransformados,
	}
}

func (g *GraficoBola) Desenha(tela *ebiten.Image, x, y float32) {
	for i, vertice := range g.vertices {
		g.verticesTransformados[i].DstX = x + vertice.DstX
		g.verticesTransformados[i].DstY = y + vertice.DstY
	}
	tela.DrawTriangles(g.verticesTransformados, g.indices, emptyImage, nil)
}

type Bola struct {
	X, Y    float32
	Raio    float32
	grafico *GraficoBola
}

func (b *Bola) Inicia(jogo *Jogo) {
	grafico, ok := jogo.recursos["graficoBola"]
	if !ok {
		grafico = NovoGraficoBola(10, b.Raio)
		jogo.recursos["graficoBola"] = grafico
	}

	b.grafico = grafico.(*GraficoBola)
}

func (b *Bola) ExecutaLogica(jogo *Jogo) {}

func (b *Bola) Desenha(tela *ebiten.Image) {
	b.grafico.Desenha(tela, b.X, b.Y)
}

type Cena struct {
	objetos []Objeto2d
}

func (c *Cena) AdicionaObjeto(objeto Objeto2d) {
	c.objetos = append(c.objetos, objeto)
}

func (c *Cena) ExecutaLogica(jogo *Jogo) {
	for _, objeto := range c.objetos {
		objeto.ExecutaLogica(jogo)
	}
}

func (c *Cena) Desenha(tela *ebiten.Image) {
	for _, objeto := range c.objetos {
		objeto.Desenha(tela)
	}
}

type Jogo struct {
	cena        Cena
	recursos    map[string]interface{}
	telaLargura int
	telaAltura  int
}

func (j *Jogo) Inicia() {
	j.recursos = make(map[string]interface{})
	for _, objeto := range j.cena.objetos {
		objeto.Inicia(j)
	}
}

func (j *Jogo) Update() error {
	j.cena.ExecutaLogica(j)
	return nil
}

func (j *Jogo) Draw(screen *ebiten.Image) {
	j.cena.Desenha(screen)
}

func (j *Jogo) Layout(outsideWidth, outsideHeight int) (screenWidth, screenHeight int) {
	return outsideWidth, outsideHeight
}

const (
	telaLargura  = 640
	telaAltura   = 480
	tituloJanela = "BlockBreaker by Davi Villalva"
)

func main() {
	ebiten.SetWindowSize(telaLargura, telaAltura)
	ebiten.SetWindowTitle(tituloJanela)

	emptyImage.Fill(color.White)

	jogador := &Jogador{
		X:       telaLargura / 2,
		Y:       telaAltura - 64,
		Tamanho: 80,
	}

	bola := &Bola{
		X:    telaLargura / 2,
		Y:    telaAltura / 2,
		Raio: 16,
	}

	jogo := &Jogo{
		telaLargura: telaLargura,
		telaAltura:  telaAltura,
	}
	jogo.cena.AdicionaObjeto(jogador)
	jogo.cena.AdicionaObjeto(bola)

	jogo.Inicia()
	if err := ebiten.RunGame(jogo); err != nil {
		log.Fatal(err)
	}
}
