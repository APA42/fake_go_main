package main

import (
	"fmt"
	"github.com/APA42/fake_go_component_1"
	"github.com/aleasoluciones/private_fake_go_component_1"
)

func main(){
	fmt.Println("Fake Main Compoment -> Executable")

	var firstFakeComponent = fake_go_component_1.FakeComponentStruct{}

	fmt.Println(firstFakeComponent.DoSomething())

	var privateFakeComponent = private_fake_go_component_1.PrivateFakeComponentStruct{}
	fmt.Println(privateFakeComponent.DoSomething())
}
