#include<stdio.h>
int main() {
    int a,b,i,j,k,x,y;
  printf("Enter side length of triangle: ");
  scanf("%d",&a);

  /* 決定loop次數 */
  /* a=1>>1行,a=2.3>>2行,a=4.5,三行... */
  if (a%2==0)
  {
      b=(a+2)/2;
  }
  else
  {
      b=(a+1)/2;
  }
  
  k=1;     /* 第一行要印1個* */  
  y=b-1;   /* 第一行要印b-1個空白 */  
  /* loop次 */
  for(i=1;i<=b;i++){

      /* 印出空白 */
      for(x=1;x<=y;x++){
          printf(" ");
      }
      
      /* 印出* */
      for(j=1;j<=k;j++){
          printf("*");
      }
      printf("\n");
      k=k+2;     /* 每一行多印兩個* */            
      y=y-1;     /* 每一行少印一個空白 */    
  }
 return 0;
} 
  
