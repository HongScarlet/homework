## 目錄
* [目錄](#目錄)
* [前置作業](#前置作業)
* [兩者差異](#兩者差異)
* [yt-dlp](#yt-dlp)
* [ytarchive](#ytarchive)
* [test](#test)
---

## 前置作業

1. 下載yt-dlp>>>[yt-dlp下載連結](https://github.com/yt-dlp/yt-dlp) 
2. 下載ytarchive>>>[ytarchive下載連結](https://github.com/Kethsar/ytarchive) 
3. 下載ffmpeg>>>[ffmpeg下載連結](https://ffmpeg.org/download.html)
4. 將上述執行檔路徑加入環境變數中(PATH)

## 兩者差異
```
# yt-dlp
1. 可下載已完成直播之存檔
2. 可下載會限內容 (需要cookies)
3. 可下載直播中的存檔 (但會由執行command當下開始記錄)
4. 無法預約下載
5. 不可連續載6小時以上
7. 若下載途中影片被轉為私人，會無法下載完全
```
```
# ytarchive
1. 無法下載已完成直播之存檔
2. 可下載會限內容 (需要cookies)
3. 可下載直播中的存檔 (即使直播設置為不可回放，也可以從直播開頭開始下載)
4. 可以預約下載
```
```
# 結論
1. 下載影片 or 直播已結束，使用 yt-dlp
2. 直播中 or 直播尚未開始(預約直播)，使用 ytarchive
```

## yt-dlp
```
# 常用參數
  -F                  列出可下載格式
  -o                  輸出路徑及檔案名稱
  --format            選用的下載格式
  --cookies=          cookies的檔案位置
  --write-thumbnail   下載直播縮圖
  --throttled-rate    下載速度
```
```
# example
yt-dlp -F [yt_url]
  ● 列出可供下載之格式
  
yt-dlp --format bestvideo+bestaudio --write-thumbnail [yt_url]
  ● 下載最高品質影片+音訊，並同時下載縮圖
  
yt-dlp --format bestvideo+bestaudio --write-thumbnail -o "(path_to_your_foldor)\[%(upload_date)s]%(title)s-%(id)s.%(ext)s" [yt_url]
  ● 下載至(path_to_your_foldor)，並且命名成 [上傳日期]標題-影片ID.副檔名
  ● 詳細output格式請參照原作者github
  
yt-dlp --format bestvideo+bestaudio --write-thumbnail --cookies=(path_to_your_cookies_file) -o "(path_to_your_foldor)\[%(upload_date)s]%  (title)s-%(id)s.%(ext)s" --throttled-rate 2M [yt_url]
  ● 使用cookies，會限影片/直播需要cookies下載
  ● 若失敗可以嘗試重新撈cookie (ex：加新會員/帳號有重新登入過等等)
  
yt-dlp --format 301 --write-thumbnail --cookies=(path_to_your_cookies_file) -o "(path_to_your_foldor)\[%(upload_date)s]%  (title)s-%(id)s.%(ext)s" --throttled-rate 2M [yt_url]
  ● --format 301 僅限於直播中的下載 (m3u8格式)
  ● 注意，此方法僅能從下command地當下開始進行記錄，直播中下載建議使用ytarchive
  ● 刪檔直播建議使用 ytarchive 或者是 yt-dlp 由直播開始時就用m3u8下載，若直播結束後才使用yt-dlp下載，有機會載到一半時轉私人
  
```

## test
