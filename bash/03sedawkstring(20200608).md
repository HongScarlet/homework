<strong>bash code<strong>
```bash
#!/usr/bin/env bash

STR="/usr/lib/python/site-package/xxx-1.0/yyy.zz"

echo "str: "$STR
echo
echo "Q1 : /usr/lib/python/site-package/xxx-1.0/yyy"
echo "Q2 : /usr/lib/python/site-package/xxx-1"
echo "Q3 : usr/lib/python/site-package/xxx-1.0/yyy.zz"
echo "Q4 : yyy.zz"
echo
echo -n "ans1 : "
echo $STR | sed 's/...$//'
echo -n "ans2 : "
echo $STR | sed 's/\./ /' | awk '{print $1}'
echo $STR | awk 'BEGIN {FS="."}; {print $1}'
echo -n "ans3 : "
echo $STR | sed 's/^.//'
echo -n "ans4 : "
echo $STR | awk 'BEGIN {FS="/"}; {print $NF}'
 
read -p "Press enter to end..." 

```

說明

<pre><code>第一題sed 's/...$//' 由後方刪除3個字元(每多一個"."則多一個字) ，例如 sed s/.$// 為刪除最後一個字
第二題先利用 sed 's/\./ / 將"."取代成空格(這時字串被空格分成三份)，再利用awk 抓第一個部分
其實也可以直接將awk的分隔符號改為"."，再取第一部分
第三題直接刪除第一個字即可
第四題修改awk的分隔符號為"/"，再取最後一行{print $NF}
上面awk 當中 BEGIN的用途是讓這個指令先執行 (先去修改分隔符號)

</code></pre>

測試結果

![image](https://github.com/HongScarlet/homework/blob/master/bash/img/03sedawkstring.png)


***
