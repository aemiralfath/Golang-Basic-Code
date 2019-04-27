package main
import "fmt"

func perulangan(){
	//tipe 1
	for i:=0;i<5;i++{
		fmt.Println("Angka", i)
	}

	//tipe 2
	var i = 0
	for i<5{
		fmt.Println("Angka", i)
		i++
	}

	//tipe 3
	i = 0
	for {
		fmt.Println("Angka",i)
		i++
		if i == 5{
			break
		}
	}

	//tipe 4
	for i := 1; i<10; i++{
		if i%2 == 1{
			continue
		}
		if i > 8{
			break
		}
		fmt.Println("Angka", i)
	}

	//tipe 5
	for i:=0;i<5;i++{
		for j:=i;j<5;j++{
			fmt.Print(j," ")
		}
		fmt.Println()
	}

	//tipe 6
	outerLoop:
	for i:=0;i<5;i++{
		for j:=0;j<5;j++{
			if i==3{
				break outerLoop
			}
			fmt.Print("matriks [",i,"][",j,"]\n")
		}
	}
}
