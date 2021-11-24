package com.pg
def a=["hi":4,3:5]
println a[3]
println a.get("hi")
println a.size()
a.put("name","shehzin")
println a
println a.containsKey(3)
println a.containsValue(4)
println a.getClass()//class from where maps are implemented
def ab=a.clone()
println ab
a.each {key,val -> println "$key has a value $val"}
a.eachWithIndex {key,val,i -> println "at $i $key has a value $val"}
def abc=a.entrySet()
println abc.getClass()
a.each { entry ->
            assert entry.key in ["hi",3,"name"]//it will give error if any of the keys are from out of the list
            }
ab.clear()
print(ab)