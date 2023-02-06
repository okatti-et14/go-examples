```
# 日本語フォント取得
apt update
apt install -y fonts-ipafont
apt install -y fontconfig
fc-cache -fv
# serverless-libreofficeを取得
weget https://github.com/vladgolubev/serverless-libreoffice/releases#:~:text=Jan%2010%2C%202020-,lo.tar.gz,-135%20MB
tar -xf lo.tar.gz
# 実行
## 2回実行しないとPDFが生成されない https://github.com/vladgolubev/serverless-libreoffice/issues/31
./instdir/program/soffice.bin --headless --invisible --nodefault --nofirststartwizard --nolockcheck --nologo --norestore --convert-to pdf --outdir ./ tmp.xlsx
./instdir/program/soffice.bin --headless --invisible --nodefault --nofirststartwizard --nolockcheck --nologo --norestore --convert-to pdf --outdir ./ tmp.xlsx
```
