// int funcion(int* v){
//     int n=sizeof(v)/sizeof(v[0]);
//     int prom=0;
//     for(int i=0;i<n;i++){
//     prom=prom+v[i];
//     }
//     return prom/n;
// }

package main

import (
	"errors"
	"fmt"
)

func main() {
	prom, err := promedio([]int{1, 2, 3, 3, 4})
	if err != nil {
		fmt.Println("Error: ", err)
		return
	}
	fmt.Println("Promedio: ", prom)
}

func promedio(v []int) (int, error) {
	n := len(v)
	if n == 0 {
		e := errors.New("empty array")
		return 0, e
	}
	prom := 0
	for _, x := range v {
		prom = prom + x
	}
	return prom / n, nil
}
