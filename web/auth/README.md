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

The thing is, sending authentication requests frequently is pretty annoying. 
In this case, this ROP Backend service has an "independent" way of authentication, and it is implemented by using JWT.  

If the cookie `qsc_rop_jwt` doesn't exist, or if this cookie is invalid for some reason, then this program will send an authentication request to the QSC Passport services, with the value stored in `qp2gl_sesstok` or `qp2gl_sesstok_secure`. 
According to this response, if the user is authorized, then a `qsc_rop_jwt` cookie will be set, and before this cookie expires, the user can quickly access the ROP Backend service.

**You can find all the procedures mentioned above in the source code files under the current directory.**



### Improvements
Since the payload of JWT is not encrypted, maybe we need to encrypt it or use `JWE` instead. For encryption, we may simply use `XOR`, or maybe we can also use the `aes` package.  
The final implementation is still under discussion. Currently, the generation and parsing functions of JWT has been implemented.  

### Draft of encryption implementation
**AES128-CBC-Base64**  

| Sequence | Layer name | Implementation |
| :--- | :--- | :------------ |
| 1 | Padding | PKCS 5 |
| 2 | Block cipher mode | CBC |
| 2 | Encryption algorithm | AES-128 |
| 3 | Binary-to-text encoding | Base64 |

This seems complicated, although in fact it is quite easy to understand.  
I'm afraid this will puzzle other maintainers...  
Maybe `JWE` is better than this...  
