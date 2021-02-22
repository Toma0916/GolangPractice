package main

type Vertex struct {
	X, Y float64
}

func (v *Vertex) Scale(f float64) {
	v.X = v.X * f
	v.Y = v.Y * f
}

func ScaleFunc(v *Vertex, f float64) {
	v.X = v.X * f
	v.Y = v.Y * f
}

func main() {
	v := Vertex{3, 4}
	v.Scale(2) // ポインタレシーバであれば呼び出し時に変数またはポインタいずれかのレシーバとして用いることができる
	p := &v
	p.Scale(19)
	ScaleFunc(&v, 10)
}
