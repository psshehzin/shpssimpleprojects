package com.pg
def a=[3,4,"5sasd",6]
println a.get(1)
println a[2][3]
println a[0..2]
println a.contains(3)
println a.size()
a.add(5)
a << 10
a.add(1,"jhasduaysgdug")
println a
a.remove(3)
a.pop()//removeslasttelement
println a.intersect([3,4,"jhasduaysgdug","blabla"])//will return those match
println a.reverse()
println a.sort()
a.clear()
println a.isEmpty()