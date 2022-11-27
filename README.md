## How to HTTP/1 works

- fakta bahwa HTTP/1 lebih lambat karena dia membuka koneksi TCP per request
- HTTP/1 tidak mengcompress header dan header plain text akan membuat payload request anda lebih besar dan
  konsekuensinya adalah latensi yang lebih besar untuk setiap call atau panggilan
- terkait point 1, HTTP/1 sebenarnya hanya menerima pola request/response dan jika semua request dan response
  ini berbagi koneksi yang berumur panjang, kita tidak akan memiliki banyak overhead tapi faktanya HTTP/1
  membuka 1 koneksi TCP per request

### example

kita memiliki client dan server, dan secara total client ingin membuat 3 permintaan

1. untuk mendapatkan file css
2. untuk mendapatkan 1 set gambar
3. untuk mendapatkan beberapa code javascript
   jadi untuk permintaan pertama kita perlu membuka koneksi TCP, membuat request dan mendapatkan file css
   dan untuk setelahnya kita melakukan proses yang sama untuk ke2 permintaan tersebut

## How to HTTP/2 works

- Jadi, daripada kita membuka 1 koneksi per request yang diperkenalkan oleh HTTP/1, HTTP/2 menyediakan koneksi
  yang lebih lama yang akan digunakan bersama oleh beberapa request dan beberapa response.
- Jadi jauh lebih efesien, khususnya saat ini, yang dimana kita membuat lebih banyak request dan lebih sering.
- HTTP/2 pada dasarnya mendukung server push dan itu berarti bahwa server dapat mendorong banyak pesan dari 1 request
  dari client dan client tidak perlu meminta lebih banyak data. Bisa saja menunggu server untuk mendorong data secara
  langsung ketika data sudah siap.
- HTTP/2 mendukung multiplexing dan itu berarti bahwa server dan client dapat mendorong beberapa pesan secara paralel
  melalui koneksi TCP yang sama. Ini juga lebih efesien karena sekarang kita dapat memproses request dan response lebih cepat
  dan dengan demikian kita memiliki latensi yang lebih sedikit.
- Dalam HTTP/1, kita melihat bahwa header adalah header plain text, tetapi dalam HTTP/2 muatannya jauh lebih ringan karena header
  dan datanya sama-sama dikompresi menjadi data binary
- HTTP/2 lebih aman secara default karena jika terhubung dari browser melalui HTTP/2 koneksi SSL akan diperlukan secara default.

### example

kita memiliki client dan server, dan secara total client ingin membuat 3 permintaan

1. untuk mendapatkan file css
2. untuk mendapatkan 1 set gambar
3. untuk mendapatkan beberapa code javascript
   HTTP/2 akan membuka 1 koneksi TCP, membuat request untuk aset dan kemudian server dapat mendorong 3 aset kepada kita

## Conclusion

HTTP/2 akan membuat 1 request dan mendapatkan banyak response ini membuat lebih efesien dan bandwith yang digunakan akan lebih sedikit
dan memiliki security yang meningkat karena SSL akan diperlukan

## Types of API in gRPC

1. API Unary
   client akan membuat 1 request dan server akan mengembalikan 1 response
2. API Server Streaming
   client akan mengirim 1 request dan server akan mengembalikan 1 atau lebih response tergantung kebutuhan
3. API Client Streaming
   client akan mengirim 1 atau lebih request dan server akan mengembalikan 1 response
4. API Bi directional streaming
   client akan mengirim beberapa request dan server akan mengambalikan beberapa repsonse juga

```
service GreetService {
  // Unary
  rpc Greet(GreetRequest) returns (GreetResponse) {};

  // Server Streaming
  rpc GreetManyTimes(GreetRequest) returns (stream GreetReponse) {};

  // Client Streaming
  rpc LongGreet(stream GreetRequest) returns (GreetResponse) {};

  // Bi Directional Streaming
  rpc GreetEveryone(stream GreetRequest) returns (stream GreetReponse) {};
}
```

## Scalability in gRPC

Di sisi server, semuanya asynchronous
Di sisis client, kita memiliki kebebasan memilih antara asynchronous dan synchronous

## Security in gRPC

1. Bahwa serialisasi bebasis skema akan memberikan elemen keamanan pertama
2. Karena datanya biner jadi ini tidak bisa dapat dibaca oleh manusia
3. gRPC sangat menganjurkan enkripsi SSL, dan sebagian besar kita dapat melihatnya dengan fakta bahwa mudah untuk menginisialisasi koneksi
   TLS antara client dan server dan bahwa implementasi gRPC akan menyediakan inisialisasi TLS yang mudah ini
4. Menggunakan intreceptor untuk menyediakan fitur authentication ke API kita

## gRPC VS REST

|      gRPC       |        REST         |
| :-------------: | :-----------------: |
| Protocol Buffer |        JSON         |
|     HTTP/2      |       HTTP/1        |
|    Streaming    |        Unary        |
| Bi Directional  |  Client -> Server   |
|   Free Design   | GET/POST/PUT/DELETE |
