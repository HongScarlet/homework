<strong>bash code<strong>
```bash
#!/bin/bash

DB_FILE=test.sql


#!/bin/bash
TITLE="Select file menu"
PROMPT="Pick a task:"
OPTIONS=("show table form" "add" "del" "list data")

show_table_form () {
echo "----------------------------------------"
sqlite3 test.sql << EOF
.schema customers
.quit
EOF
echo "----------------------------------------"
}

add () {
echo "Please enter Name:"
read -p "Name = " name
echo "Please enter Phone:"
read -p "Phone = " phone

sqlite3 test.sql << EOF
INSERT INTO customers (Name,Phone)
VALUES ('$name','$phone');
.quit
EOF
echo "----------------------------------------"

}

del () {
echo "Please enter Id you want to delete:"
read -p "Id = " id

sqlite3 test.sql << EOF
DELETE FROM customers WHERE Id=$id;
.quit
EOF
echo "----------------------------------------"

}

list_data () {
echo "----------------------------------------"
sqlite3 test.sql << EOF
.header on
.mode column
SELECT * FROM customers;
.quit
EOF
echo "----------------------------------------"
}


function main_menu() {
  echo "${TITLE}"
  PS3="${PROMPT} "
  select OPT in "${OPTIONS[@]}" "quit"; do
    case "$REPLY" in
      1 )
        show_table_form
        main_menu
      ;;
      2 )
        add
        main_menu
      ;;
      3 )
        del
        main_menu
      ;;
      4 )
        list_data
        main_menu
      ;;
      $(( ${#OPTIONS[@]}+1 )) ) echo "Exiting!"; break;;
      *) echo "Invalid option. Try another one.";continue;;
    esac
  done
}

main_menu
                                      

```

說明

<pre><code>定義4個function "show table form" "add" "del" "list data"
分別執行 顯示表格格式 新增資料 刪除資料 顯示表格所有資料
add 新增資料時 會要求操作者輸入 Name 及 Phone
del 刪除資料時 會要求操作者輸入 Id 來刪除
另外參考課本的範例，製作了簡易的選單功能

本次的測試過程為：
顯示表格格式 > 顯示表格所有資料 > 新增一筆 Name='AAA' Phone='123456'的資料 >
顯示表格所有資料 > 刪除 Id=6 的資料 > 顯示表格所有資料
</code></pre>

測試結果

![image](https://github.com/HongScarlet/homework/blob/master/bash/img/11databasemenu.png)


***
