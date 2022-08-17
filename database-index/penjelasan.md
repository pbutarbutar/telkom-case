#Bagaimana cara index bekerja pada sebuah database?

```
Pengindeksan adalah cara untuk mendapatkan tabel yang tidak berurutan ke dalam urutan yang akan memaksimalkan efisiensi kueri saat mencari.

Saat tabel tidak diindeks, urutan baris kemungkinan tidak akan terlihat oleh kueri karena dioptimalkan dengan cara apa pun, dan karena itu kueri Anda harus menelusuri baris secara linier. Dengan kata lain, kueri harus menelusuri setiap baris untuk menemukan baris yang cocok dengan kondisi. Ini bisa memakan waktu lama. Melihat melalui setiap baris tidak terlalu efisien.

```

```
Bagaimana Pengindeksan Bekerja?
Pada kenyataannya tabel database tidak menyusun ulang sendiri setiap kali kondisi kueri berubah untuk mengoptimalkan kinerja kueri: itu tidak realistis. Pada kenyataannya, yang terjadi adalah indeks menyebabkan database membuat struktur data. Tipe struktur data kemungkinan besar adalah B-Tree. Meskipun keuntungan dari B-Tree sangat banyak, keuntungan utama untuk tujuan kami adalah dapat disortir. Ketika struktur data diurutkan, itu membuat pencarian kita lebih efisien.

Ketika indeks membuat struktur data pada kolom tertentu, penting untuk dicatat bahwa tidak ada kolom lain yang disimpan dalam struktur data. Struktur data kami untuk tabel di atas hanya akan berisi nomor company_id. Unit dan unit_cost tidak akan disimpan dalam struktur data.
```
