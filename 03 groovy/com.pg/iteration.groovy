package com.pg
// for and while are exact relicas from c so ignoring
for (a in 1..5){
    println a
}
map1=["name":"shehzi","vali":"vasu"]
for (a in map1){
    println a.key
    println a
}
1.upto(5){println "$it"}//it is a special var
1.step(10,2){println "$it"}
5.times{println "$it"}