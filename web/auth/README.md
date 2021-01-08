# Documentation of authentication

As a service, we have to judge the request's validity.  
This project uses `QSC Passport API` to take control of the access of the users.  
Every request sent to the ROP service is supposed to attach two cookies, `qp2gl_sesstok` and `qp2gl_sesstok_secure`.  
The value of one of these two cookies will be sent to the Passport API, and that API will respond the status of this user, including login status and some basic information like UID, name, student ID, etc.  
According to the documentation of Passport API, these two cookies are being used under different situations. When the ROP is using HTTPS, then `qp2gl_sesstok_secure` needs to be sent to Passport API while authentication. Otherwise, `qp2gl_sesstok` will be sent.  
The format of the requests mentioned above are shown as below:  
```shell
https://api.zjuqsc.com/passport/get_member_by_token?appid=rop&appsecret=PleaseConsultTheAdminForThis&token=${qp2gl_sesstok}
https://api.zjuqsc.com/passport/get_member_by_token?appid=rop&appsecret=PleaseConsultTheAdminForThis&token_secure=${qp2gl_sesstok_secure}
```
In testing environment, we usually use `qp2gl_sesstok`, since we do not have a valid certificate in this environment. However, in production environment, we will use HTTPS, so `qp2gl_sesstok_secure` will be used.  
To solve this problem properly, a value in the configuration file is required to tell something about the current environment, that is, whether it is production environment or not. The name of this value is `is_secure_mode`, and you can see it in the example of configuration file.  

**You can find all the procedures mentioned above in `./middleware.go`**

### Improvements
It's really weird to send an authentication request to the Passport API each time we receive a request from ROP frontend. The thing is, **we need to reduce the frequency of sending authentication requests**.  
To do this, we need to store the status of a user in some sort of way. Maybe we can cache it in the memory. Maybe we can store it in the database. Maybe we can make good use of the `payload` part of JWT. 
Since the payload of JWT is not encrypted, maybe we need to encrypt it or use `JWE` instead. For encryption, we may simply use `XOR`, or maybe we can also use the `aes` package.  
The final implementation is still under discussion. Currently, the generation and parsing functions of JWT has been implemented.  

### Draft of encryption implementation

| Sequence | Layer name | Implementation |
| :--- | :--- | :------------ |
| 1 | Padding | PKCS #5 |
| 2 | Block cipher mode | CBC |
| 2 | Encryption algorithm | AES-256 |
| 3 | Binary-to-text encoding | Base64 |

This seems complicated, although in fact it is quite easy to understand.  
I'm afraid this will puzzle other maintainers...  
Maybe `JWE` is better than this...  
