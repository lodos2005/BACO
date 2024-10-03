# BACO (Bruteforce Algorithm Complexity Optimization)

**BACO**, kullanıcının girdiği kişisel bilgilere dayalı olarak özelleştirilmiş parola listeleri oluşturan bir Go uygulamasıdır. Bu araç, brute-force saldırılarının etkinliğini artırmak için tasarlanmıştır.

## Özellikler

- Kullanıcı profiline dayalı parola listesi oluşturma
- Özel karakterler ve sayılarla kombinasyonlar ekleme
- Parola listesini metin dosyasına kaydetme

## Kurulum

### Gereksinimler

- Go 1.16 veya üzeri

### Adımlar

- Depoyu klonlayın

`git clone https://github.com/lodos2005/BACO.git`

- Proje dizinine gidin

`cd BACO`

- Uygulamayı derleyin

`go build -o BACO main.go generator.go profile.go`

- Uygulamayı çalıştırın

`./BACO`

- YADA Uygulamayı kodundan çalıştırın

`go run *.go`
