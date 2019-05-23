
// GO中使用 C 的三种方式： https://blog.csdn.net/yongyu_it/article/details/81167958
package main
 
/*
#include <stdio.h>
#include <stdlib.h>
#include <string.h>
char* test_hello(const char* name){
    const char* hello=" -> hello";
    char* result=(char*)malloc(sizeof(char)*(strlen(name))+strlen(hello));
    strcpy(result,name);
    strcat(result,hello);
    return result;
}
*/
import "C"
import "fmt"
 
func main() {
	fmt.Println(C.GoString(C.test_hello(C.CString("singcl"))));
}