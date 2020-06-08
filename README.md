##### - to run a server you can go backend folder and type :
```sh
go run .
```
##### - or you can use a executable file :
```
./start
```

## Schema url dan response nya :
##### URL Get all genre : 

```text
http://localhost:8000/artist/genre
```
##### response : 

```json
{
  "success": true,
  "message": "SUCCESS",
  "data": [
    {
      "id": 1,
      "name": "Reggae"
    },
    {
      "id": 2,
      "name": "Jazz"
    },
    {
      "id": 3,
      "name": "Pop"
    },
    {
      "id": 4,
      "name": "Rock"
    },
    {
      "id": 5,
      "name": "Ska"
    }
  ]
}
```

##### URL Get all artist by genre : 

```text
http://localhost:8000/artist/genre/5
```
##### response : 

```json
{
  "success": true,
  "message": "SUCCESS",
  "data": [
    {
      "id": 3,
      "name": "tipe-x",
      "debut": "2008-09-16",
      "category": "Group",
      "genre": "Ska",
      "image": "https://i.ytimg.com/vi/wRZzI6bxAb0/hqdefault.jpg"
    }
  ]
}
```

##### URL Get all song by artist : 

```text
http://localhost:8000/artist/song/3
```
##### response : 

```json
{
  "success": true,
  "message": "SUCCESS",
  "data": [
    {
      "id": 1,
      "artist_name": "tipe-x",
      "song_name": "selamat jalan",
      "image": "https://i.ytimg.com/vi/YwvJJBEHUGc/maxresdefault.jpg"
    },
    {
      "id": 2,
      "artist_name": "tipe-x",
      "song_name": "mawar hitam",
      "image": "https://i.ytimg.com/vi/M-ENv1kSkWM/maxresdefault.jpg"
    },
    {
      "id": 3,
      "artist_name": "tipe-x",
      "song_name": "salam rindu",
      "image": "http://i.ytimg.com/vi/mYk-ElV6vnA/maxresdefault.jpg"
    }
  ]
}
```
