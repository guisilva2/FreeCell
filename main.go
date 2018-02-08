package main

import (
	"fmt"
	"os"
	"os/exec"

	"github.com/fatih/color"
)

func clear() {
	cmd := exec.Command("clear")
	cmd.Stdout = os.Stdout
	cmd.Run()
}

type carta struct {
	valor int
	naipe string
}

func NewCard(valor int, naipe string) carta {
	card := carta{valor, naipe}
	return card
}

var c = color.RedString("♥ ")
var o = color.RedString("♦ ")
var e = color.BlackString("♠ ")
var p = color.BlackString("♣ ")

func CriarCartas() {
	//naipes := []string{"♥", "♠", "♦", "♣"}

	naipes := []string{c, e, o, p}
	for _, naipe := range naipes {
		for i := 1; i <= 13; i++ {
			baralho = append(baralho, NewCard(i, naipe))
		}
	}
}

func EmbaralhaCartas() {
	//> 4 e <
	for i := 0; i < len(baralho); i++ {
		for j := 0; j < len(baralho); j += 4 {
			aux = baralho[i]
			baralho[i] = baralho[j]
			baralho[j] = aux
		}
	}
}

func DistribuiCartas() {
	for i := 0; i < 7; i++ {
		pilha1 = append(pilha1, baralho[i])
		pilha2 = append(pilha2, baralho[i+7])
		pilha3 = append(pilha3, baralho[i+14])
		pilha4 = append(pilha4, baralho[i+21])
		if i < 6 {
			pilha5 = append(pilha5, baralho[i+28])
			pilha6 = append(pilha6, baralho[i+34])
			pilha7 = append(pilha7, baralho[i+40])
			pilha8 = append(pilha8, baralho[i+46])
		}
	}
}

func exibeMesa() {
	fCyan := color.New(color.BgCyan)
	fWhite := color.New(color.BgHiWhite)
	fCyan.Println(color.BlackString(" |  1  |   |  2  |   |  3  |   |  4  |  |  A  | |  B  | |  C  | |  D  | |  E  | |  F  | |  G  | |  H  |  |  5  | |  6  | |  7  | |  8  |"))
	for i := 0; i < 15; i++ {
		if i < 1 {
			fmt.Printf("%2v %2v %2v %2v", celula1, celula2, celula3, celula4)
		} else {
			fmt.Printf("                                       ")
		}
		if i < len(pilha1) {
			fWhite.Printf(color.BlackString(" %2v"), pilha1[i])
		} else {
			fmt.Print("        ")
		}
		if i < len(pilha2) {
			fWhite.Printf(color.BlackString(" %2v"), pilha2[i])
		} else {
			fmt.Print("        ")
		}
		if i < len(pilha3) {
			fWhite.Printf(color.BlackString(" %2v"), pilha3[i])
		} else {
			fmt.Print("        ")
		}
		if i < len(pilha4) {
			fWhite.Printf(color.BlackString(" %2v"), pilha4[i])
		} else {
			fmt.Print("        ")
		}
		if i < len(pilha5) {
			fWhite.Printf(color.BlackString(" %2v"), pilha5[i])
		} else {
			fmt.Print("        ")
		}
		if i < len(pilha6) {
			fWhite.Printf(color.BlackString(" %2v"), pilha6[i])
		} else {
			fmt.Print("        ")
		}
		if i < len(pilha7) {
			fWhite.Printf(color.BlackString(" %2v"), pilha7[i])
		} else {
			fmt.Print("        ")
		}
		if i < len(pilha8) {
			fWhite.Printf(color.BlackString(" %2v"), pilha8[i])
		} else {
			fmt.Print("        ")
		}
		if i < len(naipe1) {
			///ALT
			fmt.Printf("  %2v", naipe1[i])
		} else {
			fmt.Printf("        ")
		}
		if i < len(naipe2) {
			///ALT
			fmt.Printf(" %2v", naipe2[i])
		} else {
			fmt.Printf("        ")
		}
		if i < len(naipe3) {
			///ALT
			fmt.Printf(" %2v", naipe3[i])
		} else {
			fmt.Printf("        ")
		}
		if i < len(naipe4) {
			///ALT
			fmt.Printf(" %v\n", naipe4[i])
		} else {
			fmt.Printf("        \n")
		}

		//////
	}
	fCyan.Println(color.BlackString(" |  1  |   |  2  |   |  3  |   |  4  |  |  A  | |  B  | |  C  | |  D  | |  E  | |  F  | |  G  | |  H  |  |  5  | |  6  | |  7  | |  8  |"))
}

func CorDiferente(cartaA carta, cartaB carta) bool {
	vermelho := 0
	preto := 0
	if cartaA.naipe == c || cartaA.naipe == o {
		vermelho++
	}
	if cartaA.naipe == e || cartaA.naipe == p {
		preto++
	}
	if cartaB.naipe == c || cartaB.naipe == o {
		vermelho++
	}
	if cartaB.naipe == e || cartaB.naipe == p {
		preto++
	}
	if vermelho == 1 && preto == 1 {
		return true
	} else {
		return false
	}
}

func NaipesIguais(cartaA carta, cartaB carta) bool {
	if cartaA.naipe == cartaB.naipe {
		return true
	} else {
		return false
	}
}

func MoveCarta(x1 string, auxPilha []carta, pilha2 []carta) (string, int, []carta) {
	var count int
	if len(pilha2) != 0 {
		//Alterar caso tenha apenas a sequência de 2 cartas
		for i := len(auxPilha) - 1; i > 0; i-- {
			if auxPilha[i].valor+1 == auxPilha[i-1].valor && CorDiferente(auxPilha[i], auxPilha[i-1]) {
				count++
				if i > 0 { //alteração
					if auxPilha[i].valor+1 == pilha2[len(pilha2)-1].valor && CorDiferente(auxPilha[i], pilha2[len(pilha2)-1]) {
						break
					}
				}
				fmt.Println("Contagem de baixo pra cima: ", count)
			} else {
				break
			}
		}
		if len(auxPilha) > 0 {
			if auxPilha[len(auxPilha)-1].valor+1 == pilha2[len(pilha2)-1].valor {
				count = 0
			}
		}
		if pilha2[len(pilha2)-1].valor-1 == auxPilha[len(auxPilha)-1-count].valor && CorDiferente(pilha2[len(pilha2)-1], auxPilha[len(auxPilha)-1-count]) {
			if count == 0 {
				pilha2 = append(pilha2, auxPilha[len(auxPilha)-1])
			} else {
				for i := 0; i <= count; i++ {
					pilha2 = append(pilha2, auxPilha[len(auxPilha)-1-count+i])
				}
			}
			return x1, count, pilha2
		}

	} else {
		pilha2 = append(pilha2, auxPilha[len(auxPilha)-1])
		return x1, count, pilha2
	}
	return "", count, pilha2
}
func AlteraPilha(count int, pilha []carta) []carta {
	for i := 0; i <= count; i++ {
		pilha = append(pilha[:len(pilha)-1], pilha[len(pilha):]...)
	}
	return pilha
}
func VerificaVitoria(naipe1 []carta, naipe2 []carta, naipe3 []carta, naipe4 []carta) int {
	var fundacoesCompletas int
	if len(naipe1) == 13 && len(naipe2) == 13 && len(naipe3) == 13 && len(naipe4) == 13 {
		fundacoesCompletas = 4
	}
	return fundacoesCompletas
}

//Variáveis
var baralho = []carta{}
var pilha1 = []carta{}
var pilha2 = []carta{}
var pilha3 = []carta{}
var pilha4 = []carta{}
var pilha5 = []carta{}
var pilha6 = []carta{}
var pilha7 = []carta{}
var pilha8 = []carta{}
var aux = carta{1, "aux"}
var auxPilha = []carta{}
var celula1 = []carta{{1, ""}}
var celula2 = []carta{{2, ""}}
var celula3 = []carta{{3, ""}}
var celula4 = []carta{{4, ""}}
var naipe1 = []carta{}
var naipe2 = []carta{}
var naipe3 = []carta{}
var naipe4 = []carta{}

func main() {
	var x, y, aviso string
	var count, fundacoesCompletas int
	//agora := time.Now().Second()
	//random := rand.Intn((agora / 2))

	CriarCartas()
	EmbaralhaCartas()
	DistribuiCartas()

	for x != "sair" {
		//clear()
		switch aviso {
		case "1":
			celula1[0].valor = 1
			celula1[0].naipe = ""
			aviso = ""
		case "2":
			celula2[0].valor = 2
			celula2[0].naipe = ""
			aviso = ""
		case "3":
			celula3[0].valor = 3
			celula3[0].naipe = ""
			aviso = ""
		case "4":
			celula4[0].valor = 4
			celula4[0].naipe = ""
			aviso = ""
		case "a":
			pilha1 = AlteraPilha(count, pilha1)
			aviso = ""
		case "b":
			pilha2 = AlteraPilha(count, pilha2)
			aviso = ""
		case "c":
			pilha3 = AlteraPilha(count, pilha3)
			aviso = ""
		case "d":
			pilha4 = AlteraPilha(count, pilha4)
			aviso = ""
		case "e":
			pilha5 = AlteraPilha(count, pilha5)
			aviso = ""
		case "f":
			pilha6 = AlteraPilha(count, pilha6)
			aviso = ""
		case "g":
			pilha7 = AlteraPilha(count, pilha7)
			aviso = ""
		case "h":
			pilha8 = AlteraPilha(count, pilha8)
			aviso = ""
		case "5":
			naipe1 = AlteraPilha(count, naipe1)
			aviso = ""
		case "6":
			naipe2 = AlteraPilha(count, naipe2)
			aviso = ""
		case "7":
			naipe3 = AlteraPilha(count, naipe3)
			aviso = ""
		case "8":
			naipe4 = AlteraPilha(count, naipe4)
			aviso = ""
		}

		exibeMesa()
		if aviso != x {
			fmt.Println(aviso)
		} else {
			fmt.Println("")
		}

		fmt.Println("Mova a carta:")
		fmt.Scan(&x, &y)
		if x == y {
			x = ""
			y = ""
			aviso = "Por favor, selecione o destino diferente da origem."
		}

		count = 0
		switch x {
		case "1":
			auxPilha = celula1
		case "2":
			auxPilha = celula2
		case "3":
			auxPilha = celula3
		case "4":
			auxPilha = celula4
		case "a":
			auxPilha = pilha1

		case "b":
			auxPilha = pilha2

		case "c":
			auxPilha = pilha3

		case "d":
			auxPilha = pilha4
		case "e":
			auxPilha = pilha5
		case "f":
			auxPilha = pilha6
		case "g":
			auxPilha = pilha7
		case "h":
			auxPilha = pilha8
		case "5":
			auxPilha = naipe1
		case "6":
			auxPilha = naipe2
		case "7":
			auxPilha = naipe3
		case "8":
			auxPilha = naipe4
		}

		switch y {
		case "1":
			if len(auxPilha) > 0 {
				if celula1[0].naipe == "" {
					celula1[0] = auxPilha[len(auxPilha)-1]
					aviso = x
				}
			}
		case "2":
			if len(auxPilha) > 0 {
				if celula2[0].naipe == "" {
					celula2[0] = auxPilha[len(auxPilha)-1]
					aviso = x
				}
			}
		case "3":
			if len(auxPilha) > 0 {
				if celula3[0].naipe == "" {
					celula3[0] = auxPilha[len(auxPilha)-1]
					aviso = x
				}
			}
		case "4":
			if len(auxPilha) > 0 {
				if celula4[0].naipe == "" {
					celula4[0] = auxPilha[len(auxPilha)-1]
					aviso = x
				}
			}
		case "a":
			MoveCarta(x, auxPilha, pilha1)
			aviso, count, pilha1 = MoveCarta(x, auxPilha, pilha1)
		case "b":
			MoveCarta(x, auxPilha, pilha2)
			aviso, count, pilha2 = MoveCarta(x, auxPilha, pilha2)
		case "c":
			MoveCarta(x, auxPilha, pilha3)
			aviso, count, pilha3 = MoveCarta(x, auxPilha, pilha3)
		case "d":
			MoveCarta(x, auxPilha, pilha4)
			aviso, count, pilha4 = MoveCarta(x, auxPilha, pilha4)
		case "e":
			MoveCarta(x, auxPilha, pilha5)
			aviso, count, pilha5 = MoveCarta(x, auxPilha, pilha5)
		case "f":
			MoveCarta(x, auxPilha, pilha6)
			aviso, count, pilha6 = MoveCarta(x, auxPilha, pilha6)
		case "g":
			MoveCarta(x, auxPilha, pilha7)
			aviso, count, pilha7 = MoveCarta(x, auxPilha, pilha7)
		case "h":
			MoveCarta(x, auxPilha, pilha8)
			aviso, count, pilha8 = MoveCarta(x, auxPilha, pilha8)
		case "5":
			fundacoesCompletas = VerificaVitoria(naipe1, naipe2, naipe3, naipe4)
			if auxPilha[len(auxPilha)-1].valor == 1 && len(naipe1) == 0 {
				naipe1 = append(naipe1, auxPilha[len(auxPilha)-1])
				aviso = x
			} else if len(naipe1) > 0 {
				if naipe1[len(naipe1)-1].valor+1 == auxPilha[len(auxPilha)-1].valor && NaipesIguais(naipe1[len(naipe1)-1], auxPilha[len(auxPilha)-1]) {
					naipe1 = append(naipe1, auxPilha[len(auxPilha)-1])
					aviso = x
				}
			}
		case "6":
			fundacoesCompletas = VerificaVitoria(naipe1, naipe2, naipe3, naipe4)
			if auxPilha[len(auxPilha)-1].valor == 1 && len(naipe2) == 0 {
				naipe2 = append(naipe2, auxPilha[len(auxPilha)-1])
				aviso = x
			} else if len(naipe2) > 0 {
				if naipe2[len(naipe2)-1].valor+1 == auxPilha[len(auxPilha)-1].valor && NaipesIguais(naipe2[len(naipe2)-1], auxPilha[len(auxPilha)-1]) {
					naipe2 = append(naipe2, auxPilha[len(auxPilha)-1])
					aviso = x
				}
			}
		case "7":
			fundacoesCompletas = VerificaVitoria(naipe1, naipe2, naipe3, naipe4)
			if auxPilha[len(auxPilha)-1].valor == 1 && len(naipe3) == 0 {
				naipe3 = append(naipe3, auxPilha[len(auxPilha)-1])
				aviso = x
			} else if len(naipe3) > 0 {
				if naipe3[len(naipe3)-1].valor+1 == auxPilha[len(auxPilha)-1].valor && NaipesIguais(naipe3[len(naipe3)-1], auxPilha[len(auxPilha)-1]) {
					naipe3 = append(naipe3, auxPilha[len(auxPilha)-1])
					aviso = x
				}
			}
		case "8":
			fundacoesCompletas = VerificaVitoria(naipe1, naipe2, naipe3, naipe4)
			if auxPilha[len(auxPilha)-1].valor == 1 && len(naipe4) == 0 {
				naipe4 = append(naipe4, auxPilha[len(auxPilha)-1])
				aviso = x
			} else if len(naipe4) > 0 {
				if naipe4[len(naipe4)-1].valor+1 == auxPilha[len(auxPilha)-1].valor && NaipesIguais(naipe4[len(naipe4)-1], auxPilha[len(auxPilha)-1]) {
					naipe4 = append(naipe4, auxPilha[len(auxPilha)-1])
					aviso = x
				}
			}
		}
		if fundacoesCompletas == 4 {
			aviso = "Parabéns, você ganhou!"
			break
		}
		if aviso != x {
			fRed := color.New(color.BgRed).Add(color.FgBlack)
			fRed.Println("   Movimento inválido, favor verificar!   ")
		} else {
			fGreen := color.New(color.BgGreen).Add(color.FgBlack)
			fGreen.Println("                Muito Bem!                 ")
		}
	}
	fmt.Println(aviso)
}
