# Tugas Deployment in Heroku

## Cara setup
- Buat file Procfile dan lakukan setting `web: <nama-project>`
- Lalu tambahkan `// +heroku goVersion go<1.19>` di dalam go.mod di atas versi golangnya atau dibawah nama modulenya. untuk versi disesuaikan dengan versi golang yang dibuat.
- lalu push project ke github repository.
- setelah itu buat project di heroku
- bagian tab resource pilih addon database yang sesuai
- bagian tab deploy pilih method connect to github lalu cari repo yang sesuai
- terus centang automatic deploy
- sebelum klik `Deploy Branch` pergi ke tab settings
- bagian tab settings di bagian config vars isi sesuai .env untuk keynya dan valuenya disesuaikan
- pada bagian buildpacks pilih bahasa go
- setelah itu kembali lagi ke tab deploy lalu klik `Deploy Branch`.
- Selesai

## Endpoint
- Restful API - https://alterra-agmc.herokuapp.com/
- Repository - https://github.com/dakasakti/deploy-apps-hexagonal

## Contributing
- Mahmuda Karima