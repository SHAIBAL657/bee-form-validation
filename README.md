# bee-form-validation

### Prerequisites

    ~ Install go and beego
    ~ Set GOPATH
    ~ Download https://github.com/SHAIBAL657/bee-form-validation and https://github.com/SHAIBAL657/beego-rest-api/sara project.
    ~ Both project should be run on individual port

### How to test
    ~ go mod tidy
    ~ bee run
    ~ go get <library-name> (if not available)
    ~ Run api project first on port 8080
    ~ This run on port 8081 (http://localhost:8081/)
    ~ Enter valid data to create a user. This give response back accordingly.

### Features 
    ~ First name and Last name can't contain admin.
    ~ Valid email,phone,date should be entered.
    ~ Validate data before sending to api (http://localhost:8080/v1/object)
    ~ Validate data in api for duplicate data and database response,then on success or error response back to http://localhost:8081/
    
### Create DATABASE postgresql (PGADMIN)

    createdb:=CREATE TABLE IF NOT EXISTS public."USER"
    (
    firstname character varying(255) COLLATE pg_catalog."default" NOT NULL,
    lastname character varying(255) COLLATE pg_catalog."default" NOT NULL,
    phone character varying(255) COLLATE pg_catalog."default" NOT NULL,
    password character varying(255) COLLATE pg_catalog."default" NOT NULL,
    email character varying(255) COLLATE pg_catalog."default" NOT NULL,
    dob character varying(255) COLLATE pg_catalog."default" NOT NULL,
    CONSTRAINT "USER_pkey" PRIMARY KEY (email)
    )
    
    db.Exec(createdb)

