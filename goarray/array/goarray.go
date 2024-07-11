package array

import "fmt"

//  7. Bonus: Create a "Product" struct with title, id, price and create a
//     dynamic list of products (at least 2 products).
type product struct {
	title string
	id    string
	price float64
}

func array() {
	listproducts := []product{{"firstproduct", "productname", 1299}, {"secondproduct", "productname", 1098}}
	fmt.Println(listproducts)
	listproducts = append(listproducts, product{"thirdproduct", "productname", 1066})
	fmt.Println(listproducts)
	// Time to practice what you learned!

	//  1. Create a new array (!) that contains three hobbies you have
	//     Output (print) that array in the command line.
	hobbies := [3]string{"SWIM", "CRIC", "GOLF"}
	fmt.Println(hobbies[0:])
	//  2. Also output more data about that array:
	//     - The first element (standalone)
	//     - The second and third element combined as a new list
	fmt.Println(hobbies[0])
	fmt.Println(hobbies[1:])
	//  3. Create a slice based on the first element that contains
	//     the first and second elements.
	//     Create that slice in two different ways (i.e. create two slices in the end)
	updatehobbies := hobbies[0:2]
	fmt.Println(updatehobbies)
	//  4. Re-slice the slice from (3) and change it to contain the second
	//     and last element of the original array.
	updatehobbies = hobbies[1:]
	fmt.Println(updatehobbies)
	//  5. Create a "dynamic array" that contains your course goals (at least 2 goals)
	dynamicarray := []string{"business", "fame"}
	fmt.Println(dynamicarray)
	//  6. Set the second goal to a different one AND then add a third goal to that existing dynamic array
	dynamicarray[1] = "holder"
	fmt.Println(dynamicarray)
	dynamicarray = append(dynamicarray, "fame")
	fmt.Println(dynamicarray)

}
