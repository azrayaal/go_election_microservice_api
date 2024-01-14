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

- URL: http://localhost:4000/api/v1/parties
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

`Candidate data will be shown after the admin creates a candidate(paslon)`

### Create new Party

`Authorization`

- URL: http://localhost:4000/api/v1/party
- Method: `POST`
- Request Body:

```
{
  "name": "ASTA",
  "image": "1704718751443.png",
  "number": 3,
  "vision_mission": "MENGHANCURKAN NEGARA TANAH",
  "address": "KERAJAAN SIHIR"
}
```

### Get detail Party

`No Authorization`

- URL: http://localhost:4000/api/v1/party/:id
- Method: `GET`
- Request Body:

```
{
    "id": 1,
    "name": "GETO UHUY",
    "image": "412520687_829427478869516_382939006133828220_n-1704695129411.png",
    "chairman": "Suguru Geto",
    "vision_mission": "In publishing and graphic design",
    "address": "Jalan Veteran, Jakarta Pusat",
    "candidates": [
        {
            "id": 1,
            "name": "Mimiko Hasaba",
            "image": "LOGO-1704700723367.png",
            "number": 1,
            "vision_mission": "MENGHENTIKAN INVASI NEGARA API"
        }
    ]
}
```

### Update Party

`Authorization`

- URL: http://localhost:4000/api/v1/party/:id
- Method: `PUT`
- Request Body:

```
{
  "name": "Party UPDATE",
  "image": "1704718751443.png",
  "chairman": "BULALA",
  "vision_mission": "MENGHANCURKAN NEGARA TANAH",
  "address": "KERAJAAN SIHIR"
}
```

### Delete Party

`Authorization`

- URL: http://localhost:4000/api/v1/party/:id
- Method: `DELETE`
- Request Body:

```
{
  message: "Party has been removed"
}
```

## Candidate(paslon) API

### Get all Candidates

`No Authorization`

- URL: http://localhost:4000/api/v1/candidates
- Method: `GET`
- Request Body:

```
[
    {
        "candidate_id": 1,
        "candidate_name": "Mimiko Hasaba",
        "candidate_image": "FIFA LOGO-1704700723367.png",
        "candidate_vision_mission": "MENGHENTIKAN INVASI NEGARA API",
        "partyid": 1,
        "partyname": "GETO UHUY"
    },
    {
        "candidate_id": 2,
        "candidate_name": "IPPO",
        "candidate_image": "FIFA LOGO-1704702416331.png",
        "candidate_vision_mission": "MENINJU SEMUA ORG YG ADA",
        "partyid": 2,
        "partyname": "KAMOGAWA GYM"
    }
]
```

### Create new Candidate

`Authorization`

- URL: http://localhost:4000/api/v1/candidate
- Method: `POST`
- Request Body:

```
   {
        "name": "IPPO",
        "image": "FIFA LOGO-1704702416331.png",
        "number": 3,
        "vision_mission": "MENINJU SEMUA ORG YG ADA",
        "partyid": 2
    }
```

`Make sure to send the partyId using the id from the party to associate the candidate with the correct party`

### Get detail Candidate

`No Authorization`

- URL: http://localhost:4000/api/v1/candidate/:id
- Method: `GET`
- Request Body:

```
   {
    "id": 1,
    "name": "AANG",
    "image": "FIFA LOGO-1704700723367.png",
    "vision_mission": "MENGHENTIKAN INVASI NEGARA API",
    "party": {
        "id": 1,
        "name": "GETO UHUY",
        "vision_mission": "In publishing and graphic design, Lorem ipsum is a placeholder text commonly used to demonstrate the visual form of a webpage or publication,"
        }
    }
```

### Update Candidate

`Authorization`

- URL: http://localhost:4000/api/v1/candidate/:id
- Method: `PUT`
- Request Body:

```
   {
        "name": "PASLON UPDATE",
        "image": "FIFA LOGO-1704702416331.png",
        "number": 3,
        "vision_mission": "MENINJU SEMUA ORG YG ADA",
        "partyid": 2
    }
```

### Delete Candidate

`Authorization`

- URL: http://localhost:4000/api/v1/candidate/:id
- Method: `DELETE`
- Request Body:

```
   {
        message: "Candidate has been removed"
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
