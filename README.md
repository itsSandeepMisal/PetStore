# PetStore
Come up with standalone and/or code that can be deployed to server code along with unit test included to call all of the following PETSTORE APIs and have the information displayed to the caller.

**GET /pets:** The backend returns a page of available pets in the PetStore. This is an example of the HTTP integration type. The URL of the integration endpoint is http://petstore-demo-endpoint.execute-api.com/petstore/pets.

**POST /pets:** for write access to the API's /pets resource that is integrated with the backend /petstore/pets resource. Upon receiving a correct request, the backend adds the specified pet to the PetStore and returns the result to the caller. The integration is also HTTP.

**GET /pets/{petId}:** for read access to a pet as identified by a petId value as specified as a path variable of the incoming request URL. The backend returns the specified pet found in the PetStore. The URL of the backend HTTP endpoint is http://petstore-demo-endpoint.execute-api.com/petstore/pets/n, where n is an integer as the identifier of the queried pet


**To run this project execute below commands on terminal,**
-----------------------------------------------------------

**$ go get -u github.com/gorilla/mux**

**$ go build**

**$ ./petstore**



**Sample API requests and responses.**
--------------------------------------

**1. To fetch all Pets**

GET: localhost:8080/pets

Response:
[
    {
        "id": 1,
        "type": "dog",
        "price": 249.99
    },
    {
        "id": 2,
        "type": "cat",
        "price": 124.99
    },
    {
        "id": 3,
        "type": "fish",
        "price": 0.99
    }
]


**2. To add new Pet**

POST: localhost:8080/pets

Request Body:
{
    "id": 4,
    "type": "elephant",
    "price": 10000.49
}

Response:
{
    "pet": {
        "id": 4,
        "type": "elephant",
        "price": 10000.49
    },
    "message": "success"
}


**3. To fetch pet by id**

GET: localhost:8080/pets/2

Response:
{
    "id": 2,
    "type": "cat",
    "price": 124.99
}


**Local terminal output for reference:**
-------------------------------------------

sandeep@sandeep:~/Desktop/petstore$ ./petstore

Endpoint: fetchAllPets

Calling API... http://petstore-demo-endpoint.execute-api.com/petstore/pets/

API Response as struct [{Id:1 Type:dog Price:249.99} {Id:2 Type:cat Price:124.99} {Id:3 Type:fish Price:0.99}]



Endpoint: fetchPetById

Calling API... http://petstore-demo-endpoint.execute-api.com/petstore/pets/2

API Response as struct {Id:2 Type:cat Price:124.99}



Endpoint: createPet

Calling API... http://petstore-demo-endpoint.execute-api.com/petstore/pets/

Request Data {Id:4 Type:elephant Price:10000.49}

API Response as struct {Pet:{Id:4 Type:elephant Price:10000.49} Message:success}




