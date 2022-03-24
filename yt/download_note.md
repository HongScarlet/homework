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
# yt-dlp
1. 可下載已完成直播之存檔


```

## test
