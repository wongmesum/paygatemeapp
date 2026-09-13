# PayGateMe Platform

Monorepo gateway pembayaran yang menghubungkan aplikasi merchant ke akun merchant Shopee melalui SDK komunitas.

> Peringatan: integrasi Shopee bersifat tidak resmi dan reverse-engineered. Jangan gunakan untuk uang nyata sebelum uji staging, peninjauan hukum/kepatuhan, dan rencana fallback tersedia.

## Struktur

- `paygateme/` — SDK provider, QRIS, alokasi nominal unik, dan pencocokan pembayaran.
- `paygatemeapp/` — API gateway, dashboard, PostgreSQL, webhook, dan notifikasi.

## Build

```bash
cd paygatemeapp/web
npm ci
npm run build
cd ../..
go test ./paygateme/...
go test ./paygatemeapp/...
go build ./paygatemeapp/cmd/server
```

Atau gunakan container dari root monorepo:

```bash
docker build -t paygatemeapp .
```

## Environment wajib

- `DATABASE_URL`
- `ADMIN_PASSWORD` — hash bcrypt
- `ADMIN_JWT_SECRET` — minimal 32 karakter acak
- `SESSION_ENCRYPT_KEY` — tepat 64 karakter hex

Jangan commit password, OTP, cookie Shopee, token Telegram, QRIS mentah, atau session provider.

## Health

`GET /healthz` mengembalikan status database dan autentikasi provider tanpa membuka secret.

## Deployment

Jalankan pada VPS/container yang selalu aktif. Proses ini membutuhkan polling feed, retry webhook, dan refresh session sehingga tidak cocok untuk runtime stateless/on-demand.
