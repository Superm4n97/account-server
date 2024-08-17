# account-server
`Cobra CLI` `gin` `MongoDB`

## RUN APP


## APIs
## user
**CREATE**

`[POST]`: `/api/v1/user` 
<br>
Body Data:
```json
{
  "name": "admin",
  "email": "admin@xyz",
  "password": "1234"
}
```

**GET**

`[GET]`: `/api/v1/user`<br>
`[GET]`: `/api/v1/user/:userid`
<br>

**DELETE**

`[DELETE]`: `/api/v1/user/:userid`
<br>