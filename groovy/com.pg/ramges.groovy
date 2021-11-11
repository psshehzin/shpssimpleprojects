
def a='a'..'z'
println a.getFrom()
println a.getTo()
println a.size()
assert a.from=='a'
assert a.to=='z'
println a[4]
def b=1..<10 //exclusive wont include 10
println b
println a.contains('d')
println a.isReverse()
def bc=a.subList(3,7)
println bc
for(i in bc){
    println i
}
bc.each {i -> println "$i"}