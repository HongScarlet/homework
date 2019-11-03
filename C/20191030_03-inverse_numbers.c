#include<stdio.h>
int main()
{
  int array[1000];
  int i=0;
  char c;
  printf("Example>>imput 10 20 30\n");
  printf("Input 10 20 30\n");
  printf("Output>>imput 30 20 10\n");
  printf("Please enter numbers\n");
   do
   {
    scanf("%d",&array[i]);
    i++;
   }while(c=getchar()!='\n');         /* 判斷是否按了enter */

  printf("Output\n");

  for(int j=i-1;j>=0;j--)  /* 從array的第i-1項印回第0項 */
  {
       printf("%d ",array[j]);
  }
   return 0;
}


