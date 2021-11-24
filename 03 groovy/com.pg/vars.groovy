package com.pg
def name="shehzin"
String hello="Hello"
//the above are the two methods to declare vars either def or data type and s i capital
println "$hello name is $name"
println '$hello name is $name'
// the above illustartes difference between single and double quotes
//groovy the typing is dynamic  eg below
hello=55
println "$hello name is $name"
// initialising multiple vars
def (String a,int b,Double c)=[30,6,7]
println("$a")
println("$b")
println("$c")