

[Back](../README.md)


# Petdemo

## Integration Diagram
![](integration.svg)







## Application Index


| Application Name | Method | Source Location |
|----|----|----|
| Petdemo | [GET /random-pet-pic](#Petdemo-GETrandom-pet-pic) | [https://github.com/anz-bank/sysl-template/blob/master/specs/frontend/petdemo/petdemo.yaml](https://github.com/anz-bank/sysl-template/blob/master/specs/frontend/petdemo/petdemo.yaml)|  




## Type Index


| Application Name | Type Name | Source Location |
|----|----|----|
| Petdemo | [GenericError](#Petdemo.GenericError) | [https://github.com/anz-bank/sysl-template/blob/master/specs/frontend/petdemo/petdemo.yaml](https://github.com/anz-bank/sysl-template/blob/master/specs/frontend/petdemo/petdemo.yaml)|
| Petdemo | [GenericErrorStatus](#Petdemo.GenericErrorStatus) | [https://github.com/anz-bank/sysl-template/blob/master/specs/frontend/petdemo/petdemo.yaml](https://github.com/anz-bank/sysl-template/blob/master/specs/frontend/petdemo/petdemo.yaml)|
| Petdemo | [PetResponse](#Petdemo.PetResponse) | [https://github.com/anz-bank/sysl-template/blob/master/specs/frontend/petdemo/petdemo.yaml](https://github.com/anz-bank/sysl-template/blob/master/specs/frontend/petdemo/petdemo.yaml)|








# Applications





## Application Petdemo



- Pet service to demo the sysl codegen capabilties.











### <a name=Petdemo-GETrandom-pet-pic></a>Petdemo GET /random-pet-pic


<details>
<summary>Sequence Diagram</summary>

![](Petdemo/getrandom-pet-pic.svg)
</details>

<details>
<summary>Request types</summary>



<span style="color:grey">No Request types</span>







</details>

<details>
<summary>Response types</summary>






![](Petdemo/genericerror.svg)





![](Petdemo/genericerror.svg)





![](Petdemo/genericerror.svg)





![](Petdemo/petresponse.svg)




</details>


---





# Types







<a name=Petdemo.GenericError></a><details>
<summary>Petdemo.GenericError</summary>

### Petdemo.GenericError



![](Petdemo/genericerrorsimple.svg)

[Full Diagram](Petdemo/genericerror.svg)


#### Fields

| Field name | Type | Description |
|----|----|----|
| status | GenericErrorStatus | |


</details>
<a name=Petdemo.GenericErrorStatus></a><details>
<summary>Petdemo.GenericErrorStatus</summary>

### Petdemo.GenericErrorStatus



![](Petdemo/genericerrorstatussimple.svg)

[Full Diagram](Petdemo/genericerrorstatus.svg)


#### Fields

| Field name | Type | Description |
|----|----|----|
| code | string | |
| description | string | |


</details>
<a name=Petdemo.PetResponse></a><details>
<summary>Petdemo.PetResponse</summary>

### Petdemo.PetResponse



![](Petdemo/petresponsesimple.svg)

[Full Diagram](Petdemo/petresponse.svg)


#### Fields

| Field name | Type | Description |
|----|----|----|
| name | string | |
| uri | string | |


</details>


<div class="footer">

