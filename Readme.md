# Golang
- Golang does not contains class, It only hace packages.
- golang feels us more connectivity with OS.

## Datatypes
- There is no need of defining datatypes, The compiler is smart enough to determin the datatype automatically.
And also at automatically remove semicolons.
- When we save the file it comppiles and check the syntex.

### Shortcut to assign a variable:
 before :
 var person ="Rio"

 after:
 person := "Rio"

 summary: 
 var person ="Rio" == person :="Rio"

 ## Note:
 If we assign any variable starting with capital letter (Person), It can directly export to another file but cannot eccess outside if assign with small starting letter(Person).
 Same is APPLICABLE FOR FUNCTIONS ALSO.

### Format Specifier:
%s -> string, %d -> int , %.3f -> float
eg: fmt.Printf("Name is %s age is %d Height is %.3f", name, age, height)

### Note2 : 
The curly braces will starts where function is declared not from the next line.
eg,

- correct:
func main(){

}

- wrong:
func main()
{

}

### Array and Slices:
- Array ->static 
- Slices -> dynamic (It's like vector)

declaration :

Array :
- var name[5] string
- var name =[5] int{1,2,3,4,5}

Slices :
- numbers :=[] int {1,2,3}
It has length,capacity, Pointer.


### Map

- unordered
- Dynamic
- Key and values
- Initialization
 
 ### Time
- In go, The time format is fixed. i.e, 2006-01-02 15:04:05
- YYYY-MM-DD hh:mm:ss


