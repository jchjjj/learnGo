/*************************************************************************
	> File Name: hello.c
	> Author: ma6174
	> Mail: ma6174@163.com 
	> Created Time: 2013年10月14日 星期一 16时50分49秒
 ************************************************************************/

#include<stdio.h>
int main()
{
    long i,sum;
    for(i=0;i<1000000000;i++){
        sum +=i;
    }
    printf("sum :%lx\n,\t%ld\n",sum);
    return 0;
}
