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
| echo | GORM 2.0 | PostgreSQL | ? |

## Before you start
Please create `conf/login.json` from the root of your project directory.  
The configurations related to *database connection* are stored in this file, with the format shown below:  
```json
{
  "user":   "rop",
  "passwd": "rop_pass",
  "host":   "localhost",
  "port":   "5432",
  "dbName": "rop"
}
```

The value of these variables depends on the configuration of your PostgreSQL database. **Please don't just simply copy and paste it without thinking about any possible modification.**  
This is only a short-term solution. Configuration solutions like `Viper` are considered to be used in the future.  

## Directories
`database` stores CRUD functions and models of the tables.  
`web` stores functions that handle requests and perform responses.  
