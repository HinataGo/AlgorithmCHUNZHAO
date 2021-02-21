package mon

//
// func fib(n int) int {
// 	a, b, c := 0, 1, 0
// 	for i := 0; i < n; i++ {
// 		c = a + b
// 		a = b
// 		b = c
// 	}
// 	return a
// }
func fib(n int) int {
	a, b := 0, 1
	for i := 0; i < n; i++ {
		tmp := a + b
		a = b
		b = tmp
	}
	return b
}
