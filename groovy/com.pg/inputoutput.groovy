def name=System.console().readLine()
println "hello $name"
print "enter the first number: "
def a=System.console().readLine().toInteger()
print "enter the second number: "
def b=System.console().readLine().toInteger()
println "$a + $b = "+(a+b)
printf "sum is %d\n",a+b
printf "%d,%s,%-6s,%6s,%.2f\n",[3,"hi","hi","hi",10.2763]
//%-6s will pat to rightand%6swill pat to left