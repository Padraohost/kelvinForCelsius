package main

import "fmt"

const ebulicaoK float64 = 373.0


func main(){

	tempK:= ebulicaoK
	tempKE:= (tempK - 273)

	fmt.Printf("O ponto de ebulição da água em CELSIUS é %g", tempKE)

}