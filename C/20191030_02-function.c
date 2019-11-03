/* 加1的函數 */
void F_20191030(int *x)
{
   int tmp;
   tmp = *x;    /* 保存x */
   *x = tmp+1;    /* tmp+1之後把值給予x */
  
   return;
}