package reflection

func Print(x any) {
	t := reflect.TypeOf(x)
	switch t.Kind() {
	case reflect.Int:
		fmt.Println("It's an integer!")
	case reflect.Float64:
		fmt.Println("It's a float!")
	case reflect.String:
		fmt.Println("It's a string!")
	default:
		fmt.Println("I don't recognize the type.")
	}
}
