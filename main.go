package main
import ("fmt")

func main() {
  fmt.Println("Hello World!")


//package main; import ("fmt"); func main() { fmt.Println("Hello World!");}









  var student1 string = "John" //type is string
  var student2 = "Jane" //type is inferred
  x := 2 //type is inferred

  fmt.Println(student1)
  fmt.Println(student2)
  fmt.Println(x)











  var a string // return " "
  var b int//return 0
  var c bool//return false

  fmt.Println(a)
  fmt.Println(b)
  fmt.Println(c)








  var s string
  s = "John"
  fmt.Println(s)



  
  
  
  
  

   /*
   a := 1
   
  func main() {
    fmt.Println(a)
  }
   *///it make a err
  



/*var a int
var b int = 2
var c = 3
 
func main() {
  a = 1
  fmt.Println(a)
  fmt.Println(b)
  fmt.Println(c)
}*///it dont make a err







var aa, ba, ca, da int = 1, 3, 5, 7

fmt.Println(aa)
fmt.Println(ba)
fmt.Println(ca)
fmt.Println(da)








/*
یک متغیر می تواند یک نام کوتاه (مانند x و y) یا یک نام توصیفی تر (سن، قیمت، نام کار و غیره) داشته باشد.

قوانین نامگذاری متغیر برو:

نام متغیر باید با حرف یا نویسه زیرخط (_) شروع شود.
نام متغیر نمی تواند با یک رقم شروع شود
نام متغیر فقط می‌تواند شامل نویسه‌های الف-عددی و زیرخط ( a-z, A-Zو 0-9، و _) باشد.
نام متغیرها به حروف بزرگ و کوچک حساس هستند (age ,  Age , AGE سه متغیر متفاوت هستند)
هیچ محدودیتی در طول نام متغیر وجود ندارد
نام متغیر نمی تواند حاوی فاصله باشد
باشد نام متغیر نمی تواند هیچ کلمه کلیدی Go 

*/






const PI = 3.14
fmt.Println(PI)






/*
const A int = 1//Typed Constants
const A = 1//Untyped Constants
*/










const (
  A int = 1
  B = 3.14
  C = "Hi!"
)

fmt.Println(A)
fmt.Println(B)
fmt.Println(C)









/*
Print()
Println()
Printf()
*/







/*         \n        */ //همون <br>










var i, j string = "Hello","World"
fmt.Print(i)
fmt.Print(j)










var ii, jj string = "Hello","World"
fmt.Print(ii, "\n", jj)











var i1,j2 string = "Hello","World"
fmt.Println(i1,j2)









var i3 string = "Hello"
var j4 int = 15

fmt.Printf("i has value: %v and type: %T\n", i3, i3)
fmt.Printf("j has value: %v and type: %T", j4, j4) 
// return 
/*
i has value: Hello and type: string
j has value: 15 and type: int
*/









/////////////////////////////////General Formatting Verbs///////////////////////////////////////


/*
%v	Prints the value in the default format
%#v	Prints the value in Go-syntax format
%T	Prints the type of the value
%%	Prints the % sign
*/


var ia = 15.5
var txt = "Hello World!"

fmt.Printf("%v\n", ia)
fmt.Printf("%#v\n", ia)
fmt.Printf("%v%%\n", ia)
fmt.Printf("%T\n", ia)

fmt.Printf("%v\n", txt)
fmt.Printf("%#v\n", txt)
fmt.Printf("%T\n", txt)





/////////////////////////////Integer Formatting Verbs///////////////////////////////////////




/*
%b	Base 2
%d	Base 10
%+d	Base 10 and always show sign
%o	Base 8
%O	Base 8, with leading 0o
%x	Base 16, lowercase
%X	Base 16, uppercase
%#x	Base 16, with leading 0x
%4d	Pad with spaces (width 4, right justified)
%-4d	Pad with spaces (width 4, left justified)
%04d	Pad with zeroes (width 4
*/


var iw = 15
fmt.Printf("%b\n", iw)
fmt.Printf("%d\n", iw)
fmt.Printf("%+d\n", iw)
fmt.Printf("%o\n", iw)
fmt.Printf("%O\n", iw)
fmt.Printf("%x\n", iw)
fmt.Printf("%X\n", iw)
fmt.Printf("%#x\n", iw)
fmt.Printf("%4d\n", iw)
fmt.Printf("%-4d\n", iw)
fmt.Printf("%04d\n", iw)

/*
1111
15
+15
17
0o17
f
F
0xf
  15
15
0015
*/




//////////////////////////String Formatting Verbs/////////////////////////////////////////
/*
%s	Prints the value as plain string
%q	Prints the value as a double-quoted string
%8s	Prints the value as plain string (width 8, right justified)
%-8s	Prints the value as plain string (width 8, left justified)
%x	Prints the value as hex dump of byte values
% x	Prints the value as hex dump with spaces
*/



var txt1 = "Hello"
 
fmt.Printf("%s\n", txt1)
fmt.Printf("%q\n", txt1)
fmt.Printf("%8s\n", txt1)
fmt.Printf("%-8s\n", txt1)
fmt.Printf("%x\n", txt1)
fmt.Printf("% x\n", txt1)
/*
Hello
"Hello"
   Hello
Hello
48656c6c6f
48 65 6c 6c 6f
*/


//////////////////////////Boolean Formatting Verbs////////////////////////////////////
/*
%t	Value of the boolean operator in true or false format (same as using %v)
*/



var i4 = true
var j3= false
fmt.Printf("%t\n", i4)
fmt.Printf("%t\n", j3)
/*
true
false
*/


//////////////////////////////Float Formatting Verbs////////////////////////////////////
/*
%e	Scientific notation with 'e' as exponent
%f	Decimal point, no exponent
%.2f	Default width, precision 2
%6.2f	Width 6, precision 2
%g	Exponent as needed, only necessary digits
*/





var i5 = 3.141
fmt.Printf("%e\n", i5)
fmt.Printf("%f\n", i5)
fmt.Printf("%.2f\n", i5)
fmt.Printf("%6.2f\n", i5)
fmt.Printf("%g\n", i5)
/*

3.141000e+00
3.141000
3.14
  3.14
3.141
*/




//////////////////////////////////////////////////////////////////////////////////////////////////////////////////////////////////////////////////////////////





var a2 bool = true     // Boolean
var b2 int = 5         // Integer
var c2 float32 = 3.14  // Floating point number
var d string = "Hi!"  // String

fmt.Println("Boolean: ", a2)
fmt.Println("Integer: ", b2)
fmt.Println("Float:   ", c2)
fmt.Println("String:  ", d)








var b1 bool = true // typed declaration with initial value
var b5 = true // untyped declaration with initial value
var b3 bool // typed declaration without initial value 
b4 := true // untyped declaration with initial value

fmt.Println(b1)
fmt.Println(b5)
fmt.Println(b3)
fmt.Println(b4)
















}