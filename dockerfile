FROM --platform=$BUILDPLATFORM golang:1.23 AS builder

ARG TARGETPLATFORM
ARG BUILDPLATFORM

# Gerekli araçları yükle
RUN apt-get update && apt-get install -y \
  build-essential \
  libgtk-3-dev \
  libwebkit2gtk-4.0-dev

# Wails'i yükle
RUN go install github.com/wailsapp/wails/v2/cmd/wails@latest

WORKDIR /app

# Proje dosyalarını kopyala
COPY . .

# Wails build komutunu çalıştır
RUN CGO_ENABLED=1 GOOS=linux GOARCH=amd64 wails build

# Çıktıyı host sistemine kopyalamak için boş bir aşama
FROM scratch
COPY --from=builder /app/build/apiredator /apiredator