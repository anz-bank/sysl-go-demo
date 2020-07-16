

[Back](../README.md)


# petstore

## Integration Diagram
![](integration.svg)







## Application Index


| Application Name | Method | Source Location |
|----|----|----|
| Petstore | [GET /pet](#Petstore-GETpet) | [https://github.com/anz-bank/sysl-template/blob/master/specs/backend/petstore/petstore.yaml](https://github.com/anz-bank/sysl-template/blob/master/specs/backend/petstore/petstore.yaml)|  




## Type Index


| Application Name | Type Name | Source Location |
|----|----|----|
| Petstore | [Error](#Petstore.Error) | [https://github.com/anz-bank/sysl-template/blob/master/specs/backend/petstore/petstore.yaml](https://github.com/anz-bank/sysl-template/blob/master/specs/backend/petstore/petstore.yaml)|
| Petstore | [Pet](#Petstore.Pet) | [https://github.com/anz-bank/sysl-template/blob/master/specs/backend/petstore/petstore.yaml](https://github.com/anz-bank/sysl-template/blob/master/specs/backend/petstore/petstore.yaml)|








# Applications





## Application Petstore



- No description.











### <a name=Petstore-GETpet></a>Petstore GET /pet


<details>
<summary>Sequence Diagram</summary>

![](Petstore/getpet.svg)
</details>

<details>
<summary>Request types</summary>



<span style="color:grey">No Request types</span>







</details>

<details>
<summary>Response types</summary>






![](Petstore/error.svg)





![](Petstore/pet.svg)




</details>


---





# Types







<a name=Petstore.Error></a><details>
<summary>Petstore.Error</summary>

### Petstore.Error



![](Petstore/errorsimple.svg)

[Full Diagram](Petstore/error.svg)


#### Fields

| Field name | Type | Description |
|----|----|----|
| code | int | |
| message | string | |


</details>
<a name=Petstore.Pet></a><details>
<summary>Petstore.Pet</summary>

### Petstore.Pet



![](Petstore/petsimple.svg)

[Full Diagram](Petstore/pet.svg)


#### Fields

| Field name | Type | Description |
|----|----|----|
| id | int | |
| name | string | |
| tag | string | |


</details>


<div class="footer">

