package custompackage

//simple package and function to reverse a string
//**** as its not a main package, we cannot build it in the standard format as an executable
func Reverse(s string) string{  //captialise all functions(sort of like java)
	result := ""
	for _,v := range s{
		result = string(v) + result
	}
	return result
}