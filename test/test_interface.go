package main

func (d *DD) GetValue1() []Digit {
	var df []Digit
	d1 := new(DD)
	d1.Value = 234
	df = append(df,d1)
	return df
}

type Digit interface {
	GetValue() int
	GetPast() int
}

type DD struct {
	Value int
}

func (d *DD) GetValue() int{
	return d.Value
}

func (d *DD) GetPast() int{
	return d.Value + 1
}

func main() {
	d1 := new(DD)
	d1.Value = 1
	result:=d1.GetValue1()
	println(result[0].GetValue())
}