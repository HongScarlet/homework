/*20191023 Homework */

#include <stdio.h>
#include <stdlib.h>    /* rand() */

int main(void) {
    int min,max,x;

    printf("Please enter min:\n");
    scanf("%d", &min);
    printf("min = %d\n", min);
    printf("Please enter max:\n");
    scanf("%d", &max);
    printf("max = %d\n", max);
    
    if (max>=min)
    {
        x=(rand()%(max-min+1))+min;
        printf("Random Number = %d",x);
    }
    else
    {
        printf("ERROR");
    }
    return 0;
}

/* NOTE */

/* scanf("%d", &min)*/
/* 讓使用者輸入變數並存入min中 */

/* (rand()%(max-min+1))+min */
/* 在 min 至 max 之間隨機取一個數*/
/* EX : rand()%10)+1 => 在1~10之中取一個隨機數*/
/* EX : rand()%5)+2 => 在2~6之中取一個隨機數*/
