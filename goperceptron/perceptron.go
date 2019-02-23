package main


//Use github.com/fogleman/gg next time!! 

import (
	"math/rand"
	"gonum.org/v1/plot"
	"gonum.org/v1/plot/plotter"
	"gonum.org/v1/plot/vg/draw"
	"gonum.org/v1/plot/vg" 
	"os"
	"image/color"
	"time"
)



func main() {

	trainSize := 1500

	dataset := randomXYData(trainSize)
	p, err := plot.New()
	rand.Seed(time.Now().UTC().UnixNano())
	if err != nil {
		panic(err)
	}

	p.Title.Text = "DataSet"
	p.X.Label.Text = "X"
	p.Y.Label.Text = "Y"
	m_orig := rand.Float64()
	c_orig := rand.Float64()

	createClassifiedScatter(p, dataset,m_orig, c_orig)


	//wt.WriteTo(os.Stdout)

	pData := make(plotter.XYs, len(dataset))

	for i, xy := range dataset {
		pData[i].X = xy.X
		pData[i].Y = xy.Y
	}


	trainData := label(dataset,m_orig, c_orig)

	wt, bias := perceptronTrain(trainData, 500, 0.05)

	m_pred := (-1*wt.X/(wt.Y))
	c_pred := (-1*bias)/ (wt.Y)





	//testError := predErr(tData, wt, bias)


	orig_l := drawline(m_orig,c_orig, draw.LineStyle{Color: color.RGBA{R: 196, A: 255},})

	p.Add(orig_l)

	p.Legend.Add("class_line", orig_l)

	pLine := drawline(m_pred,c_pred, draw.LineStyle{Color: color.RGBA{G: 196, A: 255},})
	p.Add(pLine)
	p.Legend.Add("pred_line", pLine)

	wt2, e := p.WriterTo(350,350,"png")
	if e != nil {
    	panic(err)
	}

	img, f_err := os.Create("img.png")
    if err != nil{
    	panic(f_err)
    }

    defer img.Close()

	wt2.WriteTo(img)
}

func drawline(m,c float64, colorScheme draw.LineStyle) (*plotter.Line){

	var line = make(plotter.XYs,2)

	line[0].X = 0
	line[0].Y = c
	line[1].X = 5
	line[1].Y = m*(5) + c


	l,err := plotter.NewLine(line)
	l.LineStyle = colorScheme
	l.LineStyle.Width = 1
	//l.FillColor = color.RGBA{R: 196, G: 255, B: 196, A: 255}
	if err != nil {
    	panic(err)
	}
	return l
}

func predErr(data []labeledPoint, wt plotter.XY, bias float64) float64{
		//tot := len(data)
		wrong := 0.0

		for _, point := range data {
			r := wt.X*point.xy.X + wt.Y*point.xy.Y + bias 
			switch {
				case r>0:
					if point.label!=1 {
						wrong+=1
					} 
				case r<0: 
					if point.label!=0 {
						wrong+=1
					}
			}
		}
		return wrong

}


type perceptron struct{
	wt plotter.XY
	bias float64
}

func (p *perceptron) init(){
	p.wt.X = rand.Float64()
	p.wt.Y = rand.Float64()
	p.bias = rand.Float64()
}

func (p *perceptron) adjust(points plotter.XY, delta float64, lRate float64){
	p.wt.X += points.X * delta * lRate
	p.wt.Y += points.Y * delta * lRate
	p.bias += delta * lRate

}

func (p *perceptron) predict(x ,y float64) float64{
	 if (p.wt.X*x) + (p.wt.Y*y) + (2*p.bias) > 0 {
	 	return 1
	 } else {
	 	return 0
	 }
}

func perceptronTrain(dataset []labeledPoint, iter int, lRate float64)(plotter.XY, float64){
	var p perceptron
	p.init()
	for i := 0; i < iter; i++ {
		for _, point := range dataset{

			x := point.xy.X
			y := point.xy.Y
			r := p.predict(x,y)

			delta := point.label - r
			p.adjust(point.xy, delta, 0.05)
		}

	}
	return p.wt,p.bias
}

type labeledPoint struct{
	xy plotter.XY
	label float64
}

func label(data []plotter.XY, m,c float64) []labeledPoint{
	ret := make([]labeledPoint,len(data) )
	for i, xy := range data{
		ret[i].xy = xy
		if xy.Y < m*xy.X + c {
			ret[i].label = 0
		} else {
			ret[i].label = 1
		} // y= mx+c
	} 
	return ret
}

func randomXYData(size int) ([]plotter.XY) {
	var data []plotter.XY
	
	rand.Seed(time.Now().UTC().UnixNano())

	for i := 0; i < size; i++ {
		data = append(data, plotter.XY{
			5*rand.Float64(),
			5*rand.Float64(),
			})
	}
	return data
}

func createClassifiedScatter(p *plot.Plot, data []plotter.XY, m_orig,c_orig float64) (*plot.Plot) {
	
	pData := make(plotter.XYs, len(data))

	for i, xy := range data {
		pData[i].X = xy.X
		pData[i].Y = xy.Y
	}

	s, err := plotter.NewScatter(pData)
	if err != nil {
    	panic(err)
	}
	

	/*s.GlyphStyle.Shape = draw.CrossGlyph{}*/
	s.GlyphStyleFunc = func(i int) draw.GlyphStyle {
    	x, y  := pData.XY(i)
    	var c color.RGBA
    	if y < m_orig*x+c_orig{
    		c = color.RGBA{B:255 , A: 255}
    	} else {
    		c = color.RGBA{G:255 , A: 255}
    	}
   	 	if err != nil {
     	   panic(err)
    	}
    	return draw.GlyphStyle{Color: c, Radius: vg.Points(3), Shape: draw.CircleGlyph{}}
	}


	p.Add(s)

	return p
}



