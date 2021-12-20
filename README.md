# bee-form-validation

### Prerequisites

    ~ Install go and beego
    ~ Set GOPATH
    ~ Download https://github.com/SHAIBAL657/bee-form-validation and https://github.com/SHAIBAL657/beego-rest-api project.
    ~ Both project should be run on individual port

### How to test
    ~ go mod tidy
    ~ bee run
    ~ go get <library-name> (if not available)
    ~ This run on port 8081 (http://localhost:8081/)
    ~ Enter valid data to create a user. This give response back accordingly.

### Features 
    ~ First name and Last name can't contain admin.
    ~ Valid email,phone,date should be entered.
    ~ Validate data before sending to api (http://localhost:8080/v1/object)
    ~ Validate data in api for duplicate data and database response,then on success or error response back to http://localhost:8081/
