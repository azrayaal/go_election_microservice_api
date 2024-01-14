# How To Use Election Microservice Backend Using Postman

## User API

### Get all Users

`No Authorization`

- URL: http://localhost:8080/api/v1/users
- Method: `GET`
- Request Body:

```
{
    "id": 1,
    "full name": "admAzra",
    "email": "admAzra@gmail.com",
    "password": "rahasia",
    "address": "jalan jalan setapak dihiraukan saja",
    "gender": "male",
    "userName": "admAzra"
}
```

## Article API

### Get All Articles

`No Authorization`

- URL: http://localhost:8080/api/v1/articles
- Method: `GET`
- Request Body:

```
{
    "article": [
        {
            "id": 1,
            "title": "ARTICLE 1",
            "image": "impg.jpg",
            "content": "INI ISI DARI CONTENT ARTICLE 1",
            "date": "2024-01-09",
            "author": "admAzra",
            "user_id": 1,
            "author detail": {
                "id": 1,
                "full name": "admAzra",
                "email": "admAzra@gmail.com",
                "password": "rahasia",
                "address": "jalan jalan setapak dihiraukan saja",
                "gender": "male",
                "userName": "admAzra"
            }
        }
    ]
}
```

## Party (partai) API

### Get all Parties

`No Authorization`

- URL: http://localhost:8080/api/v1/parties
- Method: `GET`
- Request Body:

```
{
    "party": [
        {
            "id": 1,
            "Party Name": "JAMIHA",
            "Image": "jpg.jpg",
            "Chairman": "DADANG SUGENEP",
            "Vision Mission": "MELANJUTKAN WARISAN BAPAK",
            "Address": "LAMPUNG"
        },
        {
            "id": 2,
            "Party Name": "AGENI",
            "Image": "gambar.jpg",
            "Chairman": "AGUNI",
            "Vision Mission": "MEMUSNAHKAN SEMUA IBLIS",
            "Address": "Di tengah pulau Dewata"
        },
        {
            "id": 3,
            "Party Name": "PARTAI 3",
            "Image": "3.jpg",
            "Chairman": "ULTRAMAN TIGA",
            "Vision Mission": "MENGALAHKAN SEMuA KAIJU",
            "Address": "Gunung Salak"
        }
    ]
}
```

## Candidate(paslon) API

### Get all Candidates

`No Authorization`

- URL: http://localhost:8080/api/v1/candidates
- Method: `GET`
- Request Body:

```
{
    "candidate": [
        {
            "id": 1,
            "Candidate Name": "JAJANG MIHARJANG",
            "Image": "img.jpg",
            "Number": 1,
            "Vision Mission": "MENCARI KEBENARAN TENTANG KESALAHAN",
            "party_id": 1,
            "Party Data": {
                "id": 1,
                "Party Name": "JAMIHA",
                "Image": "jpg.jpg",
                "Chairman": "DADANG SUGENEP",
                "Vision Mission": "MELANJUTKAN WARISAN BAPAK",
                "Address": "LAMPUNG"
            }
        }
    ]
}
```

## Voter API

### Get all voters

`Authorization`

- URL: http://localhost:4000/api/v1/voters
- Method: `GET`
- Request Body:

```
[
    {
        "userid": 1,
        "voter_name": "admazra",
        "voter_address": "gang masjid",
        "voter_gender": "male",
        "candidateid": 1,
        "candidatename": "Mimiko Hasaba"
    },
    {
        "userid": 2,
        "voter_name": "admGuest",
        "voter_address": "namek",
        "voter_gender": "male",
        "candidateid": 2,
        "candidatename": "IPPO"
    }
]
```

### Create new voter

`Authorization`

- URL: http://localhost:4000/api/v1/voter
- Method: `POST`
- Request Body:

```
[
    {
        "candidateId": 1,
        "voter_name": "admazra",
        "voter_address": "gang masjid",
        "voter_gender": "male",
        "candidateid": 1,
        "candidatename": "Mimiko Hasaba"
    }
]
```

`The voter data, including voter name, voter address, etc., is retrieved from the user data obtained through user login authorization`
