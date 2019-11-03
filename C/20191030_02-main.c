#include <stdio.h>
#include <stdlib.h>    /* DEL */   

void F_20191030(int *x);
 
int main ()
{
   int a = 100;
 
   printf("a = %d\n", a );

   F_20191030(&a);
 
   printf("a+1 by function =%d\n", a );

   system("PAUSE");    /* DEL */
   return 0;
}