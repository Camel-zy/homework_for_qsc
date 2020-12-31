# ROP NEO Backend

[![Quality gate](https://sonarqube.zjuqsc.com/api/project_badges/quality_gate?project=rop-back-neo)](https://sonarqube.zjuqsc.com/dashboard?id=rop-back-neo)

[![Reliability Rating](https://sonarqube.zjuqsc.com/api/project_badges/measure?project=rop-back-neo&metric=reliability_rating)](https://sonarqube.zjuqsc.com/dashboard?id=rop-back-neo)
[![Security Rating](https://sonarqube.zjuqsc.com/api/project_badges/measure?project=rop-back-neo&metric=security_rating)](https://sonarqube.zjuqsc.com/dashboard?id=rop-back-neo)
[![Maintainability Rating](https://sonarqube.zjuqsc.com/api/project_badges/measure?project=rop-back-neo&metric=sqale_rating)](https://sonarqube.zjuqsc.com/dashboard?id=rop-back-neo)

[![Coverage](https://sonarqube.zjuqsc.com/api/project_badges/measure?project=rop-back-neo&metric=coverage)](https://sonarqube.zjuqsc.com/dashboard?id=rop-back-neo)
[![Bugs](https://sonarqube.zjuqsc.com/api/project_badges/measure?project=rop-back-neo&metric=bugs)](https://sonarqube.zjuqsc.com/dashboard?id=rop-back-neo)
[![Technical Debt](https://sonarqube.zjuqsc.com/api/project_badges/measure?project=rop-back-neo&metric=sqale_index)](https://sonarqube.zjuqsc.com/dashboard?id=rop-back-neo)
[![Vulnerabilities](https://sonarqube.zjuqsc.com/api/project_badges/measure?project=rop-back-neo&metric=vulnerabilities)](https://sonarqube.zjuqsc.com/dashboard?id=rop-back-neo)



## Overview
| Web Framework | ORM | Database | Config | Auth |
| :---------: | :---: | :------: | :-----: | :-----------: |
| echo | GORM 2.0 | PostgreSQL | Viper | jwt-go |

port: `:1323`



## Directories
`database` stores CRUD functions and models of the tables.  
`web` stores functions that handle requests and perform responses, and also includes custom middlewares.  
`conf` stores functions related to the configuration file, and it is also the default directory of configuration file.



## Configuration
Please create a configuration under file `./conf` directory before running.  
The configuration file can be in JSON, YAML, TOML or INI format. This means the file name might be `conf.json`, `conf.yaml`, `conf.toml` or `conf.ini`. Just pick a format you like best.  
**Please make sure the extension of configuration file is correct!**   
Here is a sample of a YAML format configuration file:  

```yaml
# conf/conf.yaml
login:
  user: rop
  password: rop_pass
  host: localhost
  port: 5432
  db_name: rop
passport:
  is_secure_mode: false
  app_id: rop
  app_secret: ?????????????
  api_name: https://api.zjuqsc.com/passport/get_member_by_token?
jwt:
  secret_key: ?????????????
```

When you deploy this service, please change `is_secure_mode` to `true`. For more information about this, you are required to read the documentation of *Passport API v2*.  
To get the `app_secret`, consult the admin.   
The `secrut_key` of JWT can be created by your own. It can be literally anything, just make sure it is hard for others to guess.  

The value of these variables depends on the configuration of your PostgreSQL database. **Please don't just simply copy and paste it without thinking about any possible modification.**  
This is only a short-term solution. Configuration solutions like `Viper` are considered to be used in the future.



# Authenticate by JWT
This feature is still under development, and might be removed permanently in the future.  



# Testing
The value of key `is_secure_mode` in the configuration file is expected to be set `false` during the testing period.  

Before trying to send a request to this service, you need to set at least one cookie `qp2gl_sesstok` to the request header. You can also add another cookie `qp2gl_sesstok_secure` at the same time if you want, for the program can handle this situation properly.  
For more information of these to cookies, you are *strongly* suggested reading the documentation of *Passport API v2*



# Errors

### API
| Code | Description |
| :---: | :-------- |
| SUCCESS | No error occurred |
| BAD_REQUEST | The format of request contains error |
| USR_NOT_FOUND | Requested user not found |
| ORG_NOT_FOUND | Requested organization not found |
| DEP_NOT_FOUND | Requested department not found |

### Authentication
| Code | Description |
| :---: | :-------- |
| COOKIE_NOT_FOUND | You are required to provide a QSC passport cookie. |
| AUTH_SERVICE_ERROR | Authentication services of QSC Passport is unreachable due to some reason. If this keeps happening, please consult backend developers. |
| AUTH_FAILED | The authentication failed according to the response of QSC passport authentication service. |



# API in brief

### `/api`
| URL | Method | Required Parameter | Response | Type |
| --- | :---: | ---- | -------------| :----: |
| `/`                 | GET |     | The version of the current APIs | string |
| `/user`             | GET | uid | A user's metadata | JSON |
| `/user/all`         | GET |     | A list of all users' metadata | JSON |
| `/organization`     | GET | oid | A organization's metadata | JSON |
| `/organization/all` | GET |     | A list of all organizations' metadata | JSON |
| `/organization/department`       | GET | oid did | A department's metadata under a specific organization | JSON |
| `/organization/department/all`   | GET | oid | A list of all departments' metadata under a specific organization | JSON |
