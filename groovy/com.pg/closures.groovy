package com.pg
def a = "there"
def closure1 = {name -> println "Hello world $a $name"}
closure1.call("shehzin")
def fPassClosure(closal,a){ ///passing a closure to function
    closal.call("valhalla")
    println a
}
fPassClosure(closure1,"blablabla")
def closure2 ={ d,b,c -> 
return (d+b+c) }

println closure2.call(2,3,4) 
def list1=[3,4,5]
println list1.each { it }//builtin list closure each
println list1.find { item -> item==4} //closure find 
println list1.findAll { item -> item>3}
println list1.any { item -> item>4}
println list1.collect { item -> item*4}
println list1.every { item -> item>3} //checs if all items match condition
