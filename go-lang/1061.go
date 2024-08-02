package main
import "fmt"
func main() {
	var dIni, hIni, mIni, sIni int
	var dFim, hFim, mFim, sFim int
	var  dia, hora, min, seg int
	fmt.Scanf("Dia %d", &dIni)
	fmt.Scanf("%d : %d : %d",&hIni, &mIni, &sIni)
	fmt.Scanf("Dia %d", &dFim)
	fmt.Scanf("%d : %d : %d",&hFim, &mFim, &sFim)

	dia = dFim - dIni
	hora = hFim - hIni
	min = mFim - mIni
	seg = sFim - sIni

	if(seg<0){
		seg+=60;
		min--;
	}
	
	if(min<0){
		min+=60;
		hora--;
	}
	
	if(hora<0){
		hora+=24;
		dia--;
	}

	fmt.Printf("%d dia(s)\n", dia);
    fmt.Printf("%d hora(s)\n", hora);
    fmt.Printf("%d minuto(s)\n", min);
    fmt.Printf("%d segundo(s)\n", seg);

}