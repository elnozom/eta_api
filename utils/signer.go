package utils

import "C"
import (
	"fmt"
	"syscall"
)

func signDocument() {
	tokenDll, _ := syscall.NewLazyDLL("/home/ahmed/ahmedashrafdevv/nozom/eta/api/assets/dlls/ElnozomEtoken.dll")
	fmt.Println(tokenDll)

}
