# ⚽ Terminal Bounce Game (Go)

Bu küçük Go projesi, terminal ekranında sürekli hareket eden ve
duvarlardan seken bir top animasyonu gösterir.\
Terminalde grafik olmadan, sadece karakterler kullanarak basit bir
animasyon yapılmıştır.

------------------------------------------------------------------------

## 🚀 Nasıl Çalıştırılır?

### Gereksinimler

-   [Go](https://go.dev/dl/) kurulu olmalı.

### Adımlar

``` bash
# Repoyu klonla
git clone https://github.com/karagultm/ball-game-example.git
cd ball-game-example

# Programı çalıştır
go run main.go
```

------------------------------------------------------------------------

## 🎮 Nasıl Çalışır?

-   Terminal ekranı temizlenir (`github.com/inancgumus/screen` paketi
    kullanılıyor).\
-   Top (`⚽` karakteri) belirli bir hızla hareket eder.\
-   Duvarlara çarpınca yön değiştirir (x ve y hızları tersine
    çevrilir).\
-   Sonsuz döngü halinde animasyon devam eder.

------------------------------------------------------------------------

## 📌 Kullanılan Teknolojiler

-   Go dilinde yazılmıştır.\
-   `inancgumus/screen` paketi ile terminal temizleme ve cursor kontrolü
    yapılır.
