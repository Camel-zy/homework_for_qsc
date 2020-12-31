# ROP NEO Backend

[![Quality gate](https://sonarqube.zjuqsc.com/api/project_badges/quality_gate?project=rop-back-neo)](https://sonarqube.zjuqsc.com/dashboard?id=rop-back-neo)

[![Reliability Rating](https://sonarqube.zjuqsc.com/api/project_badges/measure?project=rop-back-neo&metric=reliability_rating)](https://sonarqube.zjuqsc.com/dashboard?id=rop-back-neo)
[![Security Rating](https://sonarqube.zjuqsc.com/api/project_badges/measure?project=rop-back-neo&metric=security_rating)](https://sonarqube.zjuqsc.com/dashboard?id=rop-back-neo)
[![Maintainability Rating](https://sonarqube.zjuqsc.com/api/project_badges/measure?project=rop-back-neo&metric=sqale_rating)](https://sonarqube.zjuqsc.com/dashboard?id=rop-back-neo)

[![Coverage](https://sonarqube.zjuqsc.com/api/project_badges/measure?project=rop-back-neo&metric=coverage)](https://sonarqube.zjuqsc.com/dashboard?id=rop-back-neo)
[![Bugs](https://sonarqube.zjuqsc.com/api/project_badges/measure?project=rop-back-neo&metric=bugs)](https://sonarqube.zjuqsc.com/dashboard?id=rop-back-neo)
[![Technical Debt](https://sonarqube.zjuqsc.com/api/project_badges/measure?project=rop-back-neo&metric=sqale_index)](https://sonarqube.zjuqsc.com/dashboard?id=rop-back-neo)
[![Vulnerabilities](https://sonarqube.zjuqsc.com/api/project_badges/measure?project=rop-back-neo&metric=vulnerabilities)](https://sonarqube.zjuqsc.com/dashboard?id=rop-back-neo)


## Used packages
| Web Framework | ORM | Database | Config |
| :---------: | :---: | :------: | :-----: |
| echo | GORM 2.0 | PostgreSQL | Viper |

## Directories
`database` stores CRUD functions and models of the tables.  
`web` stores functions that handle requests and perform responses, and also includes custom middlewares.  
`conf` stores functions related to the configuration file, and it is also the default directory of configuration file.

## Before you start
Please create `conf/login.json` from the root of your project directory.  
The configurations related to *database connection* are stored in this file, with the format shown below:  
```json
{
  "login": {
    "user":   "rop",
    "password": "rop_pass",
    "host":   "localhost",
    "port":   "5432",
    "db_name": "rop"
  },
  "passport": {
    "is_secure_mode": false,
    "app_id": "rop",
    "app_secret": "???????????",
    "api_name": "https://api.zjuqsc.com/passport/get_member_by_token?"
  }
}
```
When you want to deploy this service, please change `is_secure_mode` to `true`. For more information about this, you are required to read the documentation of *Passport API v2*.  
To get the `app_secret`, consult the admin. 

The value of these variables depends on the configuration of your PostgreSQL database. **Please don't just simply copy and paste it without thinking about any possible modification.**  
This is only a short-term solution. Configuration solutions like `Viper` are considered to be used in the future.

# Testing
The value of key `is_secure_mode` in the configuration file is expected to be set `false` during the testing period.  

Before trying to send a request to this service, you need to set at least one cookie `qp2gl_sesstok` to the request header. You can also add another cookie `qp2gl_sesstok_secure` at the same time if you want, for the program can handle this situation properly.  
For more information of these to cookies, you are *strongly* suggested reading the documentation of *Passport API v2*
