

[Back](../README.md)


# jsonplaceholder

## Integration Diagram
![](integration.svg)







## Application Index
| Application Name | Method | Source Location |
----|----|----
jsonplaceholder | [GET /todos/{id}](#jsonplaceholder-GETtodos{id}) | [../../api/jsonplaceholder.json](../../api/jsonplaceholder.json)|  

## Type Index
| Application Name | Type Name | Source Location |
----|----|----
jsonplaceholder | [todosResponse](#jsonplaceholder.todosResponse) | [../../api/jsonplaceholder.json](../../api/jsonplaceholder.json)|




# Applications





## Application jsonplaceholder

- No description.





### jsonplaceholder GETtodos{id}


<details>
<summary>Sequence Diagram</summary>

![](jsonplaceholder/gettodos{id}.svg)
</details>

<details>
<summary>Request types</summary>

#### Request types








#### Path Parameter

![](primitive/intsimple.svg)



</details>
<details>
<summary>Response types</summary>

#### Response types





![](jsonplaceholder/todosresponse.svg)



</details>

---




# Types




<details>
<summary>jsonplaceholder.todosResponse</summary>

### jsonplaceholder.todosResponse

- 

![](jsonplaceholder/todosresponsesimple.svg)

[Full Diagram](jsonplaceholder/todosresponse.svg)

#### Fields

| Field name | Type | Description |
|----|----|----|
| completed | bool | |
| id | float | |
| title | string | |
| userId | float | |

</details>

<div class="footer">

