# ROP NEO Backend

[![Quality gate](https://sonarqube.zjuqsc.com/api/project_badges/quality_gate?project=rop-back-neo)](https://sonarqube.zjuqsc.com/dashboard?id=rop-back-neo)

[![Reliability Rating](https://sonarqube.zjuqsc.com/api/project_badges/measure?project=rop-back-neo&metric=reliability_rating)](https://sonarqube.zjuqsc.com/dashboard?id=rop-back-neo)
[![Security Rating](https://sonarqube.zjuqsc.com/api/project_badges/measure?project=rop-back-neo&metric=security_rating)](https://sonarqube.zjuqsc.com/dashboard?id=rop-back-neo)
[![Maintainability Rating](https://sonarqube.zjuqsc.com/api/project_badges/measure?project=rop-back-neo&metric=sqale_rating)](https://sonarqube.zjuqsc.com/dashboard?id=rop-back-neo)

[![Coverage](https://sonarqube.zjuqsc.com/api/project_badges/measure?project=rop-back-neo&metric=coverage)](https://sonarqube.zjuqsc.com/dashboard?id=rop-back-neo)
[![Bugs](https://sonarqube.zjuqsc.com/api/project_badges/measure?project=rop-back-neo&metric=bugs)](https://sonarqube.zjuqsc.com/dashboard?id=rop-back-neo)
[![Technical Debt](https://sonarqube.zjuqsc.com/api/project_badges/measure?project=rop-back-neo&metric=sqale_index)](https://sonarqube.zjuqsc.com/dashboard?id=rop-back-neo)
[![Vulnerabilities](https://sonarqube.zjuqsc.com/api/project_badges/measure?project=rop-back-neo&metric=vulnerabilities)](https://sonarqube.zjuqsc.com/dashboard?id=rop-back-neo)



## External Packages 
|  Web Framework   |     ORM      |     Database Driver     | Object Storage Driver |
| :--------------: | :----------: | :---------------------: | :-------------------: |
| labstack/echo/v4 | gorm.io/gorm | gorm.io/driver/postgres |   minio/minio-go/v7   |

| Configuration Manager |   Log Manager   |  Authentication  |
| :-------------------: | :-------------: | :--------------: |
|      spf13/viper      | sirupsen/logrus | dgrijalva/jwt-go |

service port: `:1323`

### Unit test
|      Assert      | In-Memory Database Driver |
| :--------------: | :-----------------------: |
| stretchr/testify |   gorm.io/driver/sqlite   |

## Directories
`conf` contain functions related to the configuration file, and it is also the default directory of configuration file.  
`model` contain CRUD functions and models of the tables.  
`test` contain functions only called by the unit test procedures.  
`utils` contain functions that could not be classified into any other directory.  
`web` contain functions that handle requests and perform responses, and also includes custom middlewares.

## Quick Start  
1. Clone repository
2. Install PostgreSQL and MinIO
3. Create database and bucket
4. Set configuration file properly (**IMPORTANT!!!**)
5. Run!

If you cannot understand the processes mentioned above, 
then you need to find documentations in our GitLab Wiki for help.  

## How to set configuration file correctly
Please create a configuration file under the project root directory before running.  
The configuration file can be in JSON, YAML, TOML or INI format.
This means the file name could be `conf.json`, `conf.yaml`, `conf.toml` or `conf.ini`.
Just pick a format you like best.  
**Please make sure the extension of configuration file is correct!**   
Here is a sample of a YAML format configuration file while deployment (**NOT FOR TESTING**).
If you want to make some testings on your own device, some values in the configuration file needs to be changed.
For more information, please seek for the related documentation on GitLab Wiki.

```yaml
# conf.yaml
rop:
  api_version: 0.0
  allow_origins:  # simply set a '*' in this field while testing
    - example.com
    - deploy.environment.zjuqsc.com
sql:           # please set these values according to your psql configuration
  user: rop
  password: rop_pass
  host: localhost
  port: 5432
  db_name: rop
minio:          # please set these values according to your MinIO configuration
  enable: true  # set this to false if you haven't installed MinIO on your device
  endpoint: 127.0.0.1:9000
  id: minioadmin
  secret: minioadmin
  secure: false
  bucket_name: rop
passport:
  enable: true   # this must be set true in production environment!!!
  is_secure_mode: true       # this determines which Passport cookie would be processed
  app_id: rop
  app_secret:                # consult the admin for this
  api_name: https://api.zjuqsc.com/passport/get_member_by_token?
jwt:
  issuer: rop                # note: you can freely change this
  max_age: 600               # seconds
  secret_key:                # set this by yourself
rpc:
  enable: true   # this must be set true in production environment!!!
  endpoint: 127.0.0.1:50051
  timeout: 1000              # milliseconds, including executing time
  app_id:                    # consult the admin for this
  app_key:                   # consult the admin for this
message:
  test: false
  base_url: https://rop.zjuqsc.com/form?UUID= # url for interview selection form
```

If you set `passport.enable` false, you will be in a superuser testing mode. 
This means that if you set the session token of QSC Passport as "MockToken" in the Cookie field of your request, 
you can quickly pass the authentication of a mocked QSC Passport, while using the identity `UID=0`.  
This field is designed for testing purpose, which means you *must* set it true in production environment.  


For more information about this, you are required to read the documentation of *Passport API v2*.  
To get the `passport.app_secret`, consult the admin.   
The `jwt.secret_key` can be set by your own.
It can be literally anything, just make sure it is hard for others to guess.

The value of these variables depends on the configuration of your PostgreSQL database.
**Please don't just simply copy and paste it without thinking about any possible modification.**  
