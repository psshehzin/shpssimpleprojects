def filepath= "E:/MISC/mycoderepo/vscodes/groovy/com.pg/groovyfile"
File file1=new File(filepath) //File is a class creating an object file1 now
//read enterfile as string
println file1.text
//collect lines into a list
def list1=file1.collect {it}
println list1
//collect into array of string
def ar=file1 as String[]//or
def st=file1.readLines()
println ar
println st
//reading line by lien
file1.eachLine {line -> println "lineno : $line"}
file1.eachLine {line,lineno -> println "$lineno : $line"}
def line1
file1.withReader { reader->
line1= reader.readLine()
println "$line1"
}
file1.withReader { reader->
line1= reader.readLine()
println reader.readLine()
println "$line1"
}
//reading and keeping file open 
def ofile="activefilegroovy"
def reader=file1.newReader()//file2.append(reader)
File file2=new File(ofile)
file2.append(reader.readLine())
file2.append("\n")

file2.append(reader)
reader.close()
println file1.bytes
//size in bytes
println file2.length()
//is file
println file1.isFile()
println file1.isDirectory()
//get list of files in a directory
new File("E:/MISC/mycoderepo/vscodes/groovy/com.pg").eachFileRecurse{
    file ->println file.getAbsolutePath()
}
//copying file
def file=new File("data.txt")
file<<file1.text
//to delete a file file.delete

file<<file2.text//appending is happening