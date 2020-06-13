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
echo -n "ans3 : "
echo $STR | sed 's/^.//'
echo -n "ans4 : "
echo $STR | awk 'BEGIN {FS="/"}; {print $NF}'
 
read -p "Press enter to end..." 

```

說明

<pre><code>sed 'sed s/...$//' 由後方刪除3個字元(每多一個"."則多一個字) ，例如 sed s/.$// 為刪除最後一個字
將輸入的MAC進行處理(除去: 轉為十進制)
再利用迴圈產生指令數量的連續MAC並轉回十六進制
最後依照特定格式輸出
</code></pre>

測試結果

![image](https://github.com/HongScarlet/homework/blob/master/bash/img/02createMAC.png)


***
