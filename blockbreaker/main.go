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

	"net/http"
	_ "net/http/pprof"
)

const (
	telaLargura  = 640
	telaAltura   = 480
	tituloJanela = "BlockBreaker by Davi Villalva"
)

var mapa = [9][9]uint8{
	{7, 7, 7, 7, 7, 7, 7, 7, 7},
	{7, 7, 7, 7, 7, 7, 7, 7, 7},
	{7, 7, 7, 7, 7, 7, 7, 7, 7},
	{7, 7, 7, 7, 7, 7, 7, 7, 7},
	{7, 7, 7, 7, 7, 7, 7, 7, 7},
	{7, 7, 7, 7, 7, 7, 7, 7, 7},
	{7, 7, 7, 7, 7, 7, 7, 7, 7},
	{7, 7, 7, 7, 7, 7, 7, 7, 7},
	{7, 7, 7, 7, 7, 7, 7, 7, 7},
}

var emptyImage = ebiten.NewImage(3, 3)

func limita(valor, minimo, maximo float32) float32 {
	if valor < minimo {
		return minimo
	} else if valor > maximo {
		return maximo
	} else {
		return valor
	}
}

type Vet2 struct{ X, Y float32 }

func (v Vet2) Mais(other Vet2) Vet2 {
	return Vet2{v.X + other.X, v.Y + other.Y}
}

func (v Vet2) Menos(other Vet2) Vet2 {
	return Vet2{v.X - other.X, v.Y - other.Y}
}

func (v Vet2) VezesEscalar(scalar float32) Vet2 {
	return Vet2{v.X * scalar, v.Y * scalar}
}

func (v Vet2) ProdutoEscalar(other Vet2) float32 {
	return v.X*other.X + v.Y*other.Y
}

func (v Vet2) Tamanho() float32 {
	return float32(math.Sqrt(float64(v.X*v.X + v.Y*v.Y)))
}

func (v Vet2) FormaNormal() Vet2 {
	tamanho := v.Tamanho()
	return Vet2{v.X / tamanho, v.Y / tamanho}
}

type Objeto2d interface {
	Inicia(jogo *Jogo)
	ExecutaLogica(jogo *Jogo)
	Desenha(tela *ebiten.Image)
}

type Jogador struct {
	Posicao Vet2
	Tamanho float32
}

func (j *Jogador) Inicia(jogo *Jogo) {}

func (j *Jogador) ExecutaLogica(jogo *Jogo) {
	if ebiten.IsKeyPressed(ebiten.KeyLeft) {
		j.Posicao.X -= 5
	}
	if ebiten.IsKeyPressed(ebiten.KeyRight) {
		j.Posicao.X += 5
	}

	metade := j.Tamanho / 2

	if j.Posicao.X-metade < 0 {
		j.Posicao.X = metade
	}
	if j.Posicao.X+metade > float32(jogo.telaLargura) {
		j.Posicao.X = float32(jogo.telaLargura) - metade
	}
}

func (j *Jogador) Desenha(tela *ebiten.Image) {
	ebitenutil.DrawRect(
		tela,
		float64(j.Posicao.X-j.Tamanho/2.0),
		float64(j.Posicao.Y),
		float64(j.Tamanho),
		20.0,
		&color.White,
	)
}

type GraficoBola struct {
	indices               []uint16
	vertices              []ebiten.Vertex
	verticesTransformados []ebiten.Vertex
}

func NovoGraficoBola(numVertices int, raio float32) *GraficoBola {
	var (
		indices               = make([]uint16, 0, numVertices)
		vertices              = make([]ebiten.Vertex, 0, numVertices+1)
		verticesTransformados = make([]ebiten.Vertex, 0, numVertices+1)
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
	Posicao     Vet2
	Velocidade  Vet2
	Raio        float32
	graficoBola *GraficoBola
}

func (b *Bola) Reflete(normal Vet2) {
	prodEscVelNormal := b.Velocidade.ProdutoEscalar(normal)

	if prodEscVelNormal >= 0.0 {
		return
	}

	b.Velocidade = b.Velocidade.Menos(normal.VezesEscalar(2.0 * prodEscVelNormal))
}

func (b *Bola) Inicia(jogo *Jogo) {
	grafico, ok := jogo.recursos["graficoBola"]
	if !ok {
		grafico = NovoGraficoBola(16, b.Raio)
		jogo.recursos["graficoBola"] = grafico
	}

	b.graficoBola = grafico.(*GraficoBola)
}

func (b *Bola) ExecutaLogica(jogo *Jogo) {
	if b.Posicao.X+b.Raio > float32(jogo.telaLargura) {
		b.Reflete(Vet2{-1.0, 0.0})
	}
	if b.Posicao.X-b.Raio < 0.0 {
		b.Reflete(Vet2{1.0, 0.0})
	}
	if b.Posicao.Y+b.Raio > float32(jogo.telaAltura) {
		b.Reflete(Vet2{0.0, -1.0})
	}
	if b.Posicao.Y-b.Raio < 0.0 {
		b.Reflete(Vet2{0.0, 1.0})
	}

	for _, objetoGenerico := range jogo.cena.objetos {
		switch objeto := objetoGenerico.(type) {
		case *Jogador:
			metadeTamanhoJogador := objeto.Tamanho / 2

			pontoDeContato := Vet2{
				limita(b.Posicao.X, objeto.Posicao.X-metadeTamanhoJogador, objeto.Posicao.X+metadeTamanhoJogador),
				limita(b.Posicao.Y, objeto.Posicao.Y, objeto.Posicao.Y+20.0),
			}

			var (
				delta        = b.Posicao.Menos(pontoDeContato)
				deltaTamanho = delta.Tamanho()
			)

			if deltaTamanho > b.Raio || math.Abs(float64(deltaTamanho)) < 0.0001 {
				break
			}

			deltaNormalizado := delta.FormaNormal()

			b.Posicao = pontoDeContato.Mais(deltaNormalizado.VezesEscalar(b.Raio))
			b.Reflete(deltaNormalizado)
		}
	}

	b.Posicao = b.Posicao.Mais(b.Velocidade)
}

func (b *Bola) Desenha(tela *ebiten.Image) {
	b.graficoBola.Desenha(tela, b.Posicao.X, b.Posicao.Y)
}

type Tijolo struct {
	Posicao Vet2
	Tamanho Vet2
	Vida    uint8
}

func (t *Tijolo) Inicia(jogo *Jogo) {}

func (t *Tijolo) ExecutaLogica(jogo *Jogo) {}

func (t *Tijolo) Desenha(tela *ebiten.Image) {}

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

func main() {
	go func() {
		log.Println(http.ListenAndServe("localhost:6060", nil))
	}()

	ebiten.SetWindowSize(telaLargura, telaAltura)
	ebiten.SetWindowTitle(tituloJanela)

	emptyImage.Fill(color.White)

	jogo := &Jogo{
		telaLargura: telaLargura,
		telaAltura:  telaAltura,
	}

	jogador := &Jogador{
		Posicao: Vet2{telaLargura / 2, telaAltura - 64},
		Tamanho: 80,
	}
	jogo.cena.AdicionaObjeto(jogador)

	bola := &Bola{
		Posicao:    Vet2{telaLargura / 2, telaAltura / 2},
		Velocidade: Vet2{4.0, 4.0},
		Raio:       12,
	}
	jogo.cena.AdicionaObjeto(bola)

	jogo.Inicia()
	if err := ebiten.RunGame(jogo); err != nil {
		log.Fatal(err)
	}
}
