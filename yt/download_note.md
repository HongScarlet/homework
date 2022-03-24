## 目錄
* [目錄](#目錄)
* [前置作業](#前置作業)
* [兩者差異](#兩者差異)
* [test](#test)
---

## 前置作業
```bash
# 1. 下載yt-dlp>>>[yt-dlp下載連結](https://github.com/yt-dlp/yt-dlp) 
# 2. 下載ytarchive>>>[ytarchive下載連結](https://github.com/Kethsar/ytarchive) 
# 3. 下載ffmpeg>>>[ffmpeg下載連結](https://ffmpeg.org/download.html)
# 4. 將上述執行檔路徑加入環境變數中(PATH)

```
## 兩者差異
```bash
# 1. 下載yt-dlp>>>[yt-dlp下載連結](https://github.com/yt-dlp/yt-dlp) 
# 2. 下載ytarchive>>>[ytarchive下載連結](https://github.com/Kethsar/ytarchive) 
# 3. 下載ffmpeg>>>[ffmpeg下載連結](https://ffmpeg.org/download.html)
# 4. 將上述執行檔路徑加入環境變數中(PATH)

```

Part A
```bash
# 編譯vasp前準備
server :~ # tar -zxvf vasp.6.1.1.tgz
server :~ # cd vasp.6.1.1/arch
server :~ # cp makefile.include.linux_intel ./makefile.include   # 把makefile.include.linux_intel複製到上一層並重新命名
server :~ # cd ..
```

```bash
# 記得source compliter (XXX為compliter安裝路徑)
server :~ # source  XXX/compilers_and_libraries/linux/bin/compilervars.sh intel64
server :~ # source  XXX/compilers_and_libraries/linux/mpi/intel64/bin/mpivars.sh
server :~ # source  XXX/compilers_and_libraries/linux/mkl/bin/mklvars.sh intel64

# 檢查是否source成功
server :~ # which icc
server :~ # which ifort
server :~ # which mpicc

● 順便檢查source的編譯器版本是否正確
● 如果出現 which no icc in XXX 等訊息代表沒有source成功 (每次ssh登入皆需要重新source)

# 開始編譯
server :~ # make std


● 編譯時間可能較長，請耐心等候
```

Part B
```bash
# 編譯vasp前準備
● 大致步驟與先前相同，但需要對部分檔案進行修改
● 可以參考vtst官網>>>[vtstcode安裝教學](http://theory.cm.utexas.edu/vtsttools/installation.html) 
● 由於vasp/src內 與 vtstcode內皆有chain.F 可以先將其備份
server :~ # tar -zxvf vasp.6.1.1.tgz
server :~ # tar -zxvf vtstcode-180.tgz         #解壓縮vtstcode
server :~ # cp vtstcode-180/* vasp.6.1.1/src/. #把vtstcode複製到vasp/src內
```

```bash
# 修改src/main.F
server :~ # vi main.F
將
CALL CHAIN_FORCE(T_INFO%NIONS,DYN%POSION,TOTEN,TIFOR, &
     LATT_CUR%A,LATT_CUR%B,IO%IU6)
改為
CALL CHAIN_FORCE(T_INFO%NIONS,DYN%POSION,TOTEN,TIFOR, &
     TSIF,LATT_CUR%A,LATT_CUR%B,IO%IU6)
```


```bash
# 修改src/.objects
server :~ # vi .objects
找到chain.o 並在前面加入
bfgs.o dynmat.o  instanton.o  lbfgs.o sd.o   cg.o dimer.o bbm.o \
fire.o lanczos.o neb.o  qm.o opt.o

例如
bfgs.o dynmat.o  instanton.o  lbfgs.o sd.o   cg.o dimer.o bbm.o \
fire.o lanczos.o neb.o  qm.o opt.o \
chain.o

```

```bash
# 增加solvent
server :~ # vi makefile.include
CPP_OPTIONS 加入 -Dsol_compat

例如

```

---

## test
