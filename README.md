# Authin

Testing out the knowledge provided in Alex Mux's Golang Secure Login Portal video.   
https://www.youtube.com/watch?v=OmLdoEMcr_Y  

To run the project: 
```
go run *.go
```


Endpoint structure: 
```
http://localhost:4242/register
```

Sample usage:
1. Register a user  
2. Login using that same username and password:  
```
$ curl -X POST http://localhost:4242/login -d "username=issausername&password=issapassword" -c cookies.txt
Login successful!
```
3. Visit the protected endpoint using CSRF token from the cookie and user:  
```
curl -X POST http://localhost:4242/protected -b cookies.txt -H "X-CSRF-Token: F4jwCQw5dS_NA7c6X9bIUUQBnMLsI_zsHHa-QtaNEA8=" -d "username=issausername"
CSRF validation successful! Welcome, issausername
```
