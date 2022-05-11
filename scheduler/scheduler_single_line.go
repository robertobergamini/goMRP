package scheduler

// var _permutation := map[string]int

var testarray = []string{"A", "B"}

// func generatePermutations (n int, arr []string) {

// 	// calculate permutaion
// 	//var permutations := [][]string

// 	if n == 1 {
// 		fmt.Printf("%v\n", arr)
// 	} else {

// 		generatePermutations(n-1, arr)

// 		for i := 0; i<n-1; i++ {
// 			if (n%2==0) {
// 				swap(arr[i], arr[n-1])
// 				fmt.Printf("%v\n", arr)
// 			} else {
// 				swap(arr[0], arr[n-1])
// 				fmt.Printf("%v\n", arr)
// 			}

// 			generatePermutations(n - 1, arr)
// 		}
// 	}

// }

func swap(values []string, i int, j int) {

	var tmp string = values[i]
	values[i] = values[j]
	values[j] = tmp
}
