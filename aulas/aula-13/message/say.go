package message

//Para que um método seja exportado, é necessário que 
//a primeira letra da assinatura maiúscula 
func SayHello(name string)(){
	show("Olá, "+ name)
}