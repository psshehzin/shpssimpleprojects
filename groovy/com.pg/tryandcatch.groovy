package com.pg
try{
    def int i=4/0
}catch(Exception e){  //Exception is the super class we can catch just Arithmetic exception in this case
    println e.getMessage()
    println e.getCause()
    println e.printStackTrace() // prints actual error
}//we can have multiple catch blocks and also finally which will always get executed
//finally block will get executed even if no exceptions occur