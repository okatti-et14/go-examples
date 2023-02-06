```
apt update
# LibreOfficeインストール
## 印刷範囲設定が反映される最低限の構成をインストールするように調査した方が良い
apt install -y libreoffice
# 日本語フォント取得
apt install -y fontconfig
apt install -y fonts-ipafont
fc-cache -fv
soffice --headless --invisible --nodefault --nofirststartwizard --nolockcheck --nologo --norestore --convert-to pdf --outdir . *.xlsx
```
