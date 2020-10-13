

[Back](../README.md)


# flickr

## Integration Diagram
![](integration.svg)







## Application Index


| Application Name | Method | Source Location |
|----|----|----|
| Flickr | [GET /rest](#Flickr-GETrest) | [https://github.com/anz-bank/sysl-template/blob/master/specs/backend/flickr/flickr.yaml](https://github.com/anz-bank/sysl-template/blob/master/specs/backend/flickr/flickr.yaml)|  




## Type Index


| Application Name | Type Name | Source Location |
|----|----|----|
| Flickr | [Note](#Flickr.Note) | [https://github.com/anz-bank/sysl-template/blob/master/specs/backend/flickr/flickr.yaml](https://github.com/anz-bank/sysl-template/blob/master/specs/backend/flickr/flickr.yaml)|
| Flickr | [Owner](#Flickr.Owner) | [https://github.com/anz-bank/sysl-template/blob/master/specs/backend/flickr/flickr.yaml](https://github.com/anz-bank/sysl-template/blob/master/specs/backend/flickr/flickr.yaml)|
| Flickr | [Photo](#Flickr.Photo) | [https://github.com/anz-bank/sysl-template/blob/master/specs/backend/flickr/flickr.yaml](https://github.com/anz-bank/sysl-template/blob/master/specs/backend/flickr/flickr.yaml)|
| Flickr | [PhotoResource](#Flickr.PhotoResource) | [https://github.com/anz-bank/sysl-template/blob/master/specs/backend/flickr/flickr.yaml](https://github.com/anz-bank/sysl-template/blob/master/specs/backend/flickr/flickr.yaml)|
| Flickr | [Photo_comments](#Flickr.Photo_comments) | [https://github.com/anz-bank/sysl-template/blob/master/specs/backend/flickr/flickr.yaml](https://github.com/anz-bank/sysl-template/blob/master/specs/backend/flickr/flickr.yaml)|
| Flickr | [Photo_dates](#Flickr.Photo_dates) | [https://github.com/anz-bank/sysl-template/blob/master/specs/backend/flickr/flickr.yaml](https://github.com/anz-bank/sysl-template/blob/master/specs/backend/flickr/flickr.yaml)|
| Flickr | [Photo_description](#Flickr.Photo_description) | [https://github.com/anz-bank/sysl-template/blob/master/specs/backend/flickr/flickr.yaml](https://github.com/anz-bank/sysl-template/blob/master/specs/backend/flickr/flickr.yaml)|
| Flickr | [Photo_editability](#Flickr.Photo_editability) | [https://github.com/anz-bank/sysl-template/blob/master/specs/backend/flickr/flickr.yaml](https://github.com/anz-bank/sysl-template/blob/master/specs/backend/flickr/flickr.yaml)|
| Flickr | [Photo_notes](#Flickr.Photo_notes) | [https://github.com/anz-bank/sysl-template/blob/master/specs/backend/flickr/flickr.yaml](https://github.com/anz-bank/sysl-template/blob/master/specs/backend/flickr/flickr.yaml)|
| Flickr | [Photo_people](#Flickr.Photo_people) | [https://github.com/anz-bank/sysl-template/blob/master/specs/backend/flickr/flickr.yaml](https://github.com/anz-bank/sysl-template/blob/master/specs/backend/flickr/flickr.yaml)|
| Flickr | [Photo_permissions](#Flickr.Photo_permissions) | [https://github.com/anz-bank/sysl-template/blob/master/specs/backend/flickr/flickr.yaml](https://github.com/anz-bank/sysl-template/blob/master/specs/backend/flickr/flickr.yaml)|
| Flickr | [Photo_publiceditability](#Flickr.Photo_publiceditability) | [https://github.com/anz-bank/sysl-template/blob/master/specs/backend/flickr/flickr.yaml](https://github.com/anz-bank/sysl-template/blob/master/specs/backend/flickr/flickr.yaml)|
| Flickr | [Photo_tags](#Flickr.Photo_tags) | [https://github.com/anz-bank/sysl-template/blob/master/specs/backend/flickr/flickr.yaml](https://github.com/anz-bank/sysl-template/blob/master/specs/backend/flickr/flickr.yaml)|
| Flickr | [Photo_title](#Flickr.Photo_title) | [https://github.com/anz-bank/sysl-template/blob/master/specs/backend/flickr/flickr.yaml](https://github.com/anz-bank/sysl-template/blob/master/specs/backend/flickr/flickr.yaml)|
| Flickr | [Photo_urls](#Flickr.Photo_urls) | [https://github.com/anz-bank/sysl-template/blob/master/specs/backend/flickr/flickr.yaml](https://github.com/anz-bank/sysl-template/blob/master/specs/backend/flickr/flickr.yaml)|
| Flickr | [Photo_usage](#Flickr.Photo_usage) | [https://github.com/anz-bank/sysl-template/blob/master/specs/backend/flickr/flickr.yaml](https://github.com/anz-bank/sysl-template/blob/master/specs/backend/flickr/flickr.yaml)|
| Flickr | [Photo_visibility](#Flickr.Photo_visibility) | [https://github.com/anz-bank/sysl-template/blob/master/specs/backend/flickr/flickr.yaml](https://github.com/anz-bank/sysl-template/blob/master/specs/backend/flickr/flickr.yaml)|
| Flickr | [Tag](#Flickr.Tag) | [https://github.com/anz-bank/sysl-template/blob/master/specs/backend/flickr/flickr.yaml](https://github.com/anz-bank/sysl-template/blob/master/specs/backend/flickr/flickr.yaml)|
| Flickr | [URL](#Flickr.URL) | [https://github.com/anz-bank/sysl-template/blob/master/specs/backend/flickr/flickr.yaml](https://github.com/anz-bank/sysl-template/blob/master/specs/backend/flickr/flickr.yaml)|








# Applications





## Application Flickr



- A subset of Flickr's API defined in Swagger format.











### <a name=Flickr-GETrest></a>Flickr GET /rest


<details>
<summary>Sequence Diagram</summary>

![](Flickr/getrest.svg)
</details>

<details>
<summary>Request types</summary>



<span style="color:grey">No Request types</span>










#### Query Parameter

![](primitive/stringmethod.svg)



#### Query Parameter

![](primitive/stringtext.svg)



#### Query Parameter

![](primitive/stringtags.svg)



#### Query Parameter

![](primitive/stringuser_id.svg)



#### Query Parameter

![](primitive/stringmin_upload_date.svg)



#### Query Parameter

![](primitive/stringmax_upload_date.svg)



#### Query Parameter

![](primitive/stringmin_taken_date.svg)



#### Query Parameter

![](primitive/stringmax_taken_date.svg)



#### Query Parameter

![](primitive/stringlicense.svg)



#### Query Parameter

![](primitive/stringsort.svg)



#### Query Parameter

![](primitive/stringbbox.svg)



#### Query Parameter

![](primitive/stringaccuracy.svg)



#### Query Parameter

![](primitive/stringmachine_tags.svg)



#### Query Parameter

![](primitive/stringmachine_tag_mode.svg)



#### Query Parameter

![](primitive/stringgroup_id.svg)



#### Query Parameter

![](primitive/stringcontacts.svg)



#### Query Parameter

![](primitive/stringwoe_id.svg)



#### Query Parameter

![](primitive/stringplace_id.svg)



#### Query Parameter

![](primitive/stringmedia.svg)



#### Query Parameter

![](primitive/stringhas_geo.svg)



#### Query Parameter

![](primitive/stringgeo_context.svg)



#### Query Parameter

![](primitive/stringlat.svg)



#### Query Parameter

![](primitive/stringlon.svg)



#### Query Parameter

![](primitive/stringradius_units.svg)



#### Query Parameter

![](primitive/boolis_commons.svg)



#### Query Parameter

![](primitive/boolin_gallery.svg)



#### Query Parameter

![](primitive/boolis_getty.svg)

</details>

<details>
<summary>Response types</summary>






![](Flickr/photoresource.svg)




</details>


---





# Types







<a name=Flickr.Note></a><details>
<summary>Flickr.Note</summary>

### Flickr.Note



![](Flickr/notesimple.svg)

[Full Diagram](Flickr/note.svg)


#### Fields

| Field name | Type | Description |
|----|----|----|
| description | string | |


</details>
<a name=Flickr.Owner></a><details>
<summary>Flickr.Owner</summary>

### Flickr.Owner



![](Flickr/ownersimple.svg)

[Full Diagram](Flickr/owner.svg)


#### Fields

| Field name | Type | Description |
|----|----|----|
| iconfarm | string | |
| iconserver | string | |
| is_ad_free | bool | |
| ispro | bool | |
| location | string | |
| noindexfollow | bool | |
| nsid | string | |
| path_alias | string | |
| realname | string | |
| username | string | |


</details>
<a name=Flickr.Photo></a><details>
<summary>Flickr.Photo</summary>

### Flickr.Photo



![](Flickr/photosimple.svg)

[Full Diagram](Flickr/photo.svg)


#### Fields

| Field name | Type | Description |
|----|----|----|
| comments | Photo_comments | |
| dates | Photo_dates | |
| dateuploaded | string | |
| description | Photo_description | |
| editability | Photo_editability | |
| farm | string | |
| id | string | |
| isfavorite | bool | |
| license | string | |
| media | string | |
| notes | Photo_notes | |
| originalsecret | string | |
| owner | Owner | |
| people | Photo_people | |
| permissions | Photo_permissions | |
| publiceditability | Photo_publiceditability | |
| rotation | string | |
| safe | bool | |
| safety_level | string | |
| secret | string | |
| server | string | |
| tags | Photo_tags | |
| title | Photo_title | |
| urls | Photo_urls | |
| usage | Photo_usage | |
| views | string | |
| visibility | Photo_visibility | |


</details>
<a name=Flickr.PhotoResource></a><details>
<summary>Flickr.PhotoResource</summary>

### Flickr.PhotoResource



![](Flickr/photoresourcesimple.svg)

[Full Diagram](Flickr/photoresource.svg)


#### Fields

| Field name | Type | Description |
|----|----|----|
| page | float | |
| pages | float | |
| perpage | float | |
| photos | sequence of Photo | |
| total | float | |


</details>
<a name=Flickr.Photo_comments></a><details>
<summary>Flickr.Photo_comments</summary>

### Flickr.Photo_comments



![](Flickr/photo_commentssimple.svg)

[Full Diagram](Flickr/photo_comments.svg)


#### Fields

| Field name | Type | Description |
|----|----|----|
| _content | string | |


</details>
<a name=Flickr.Photo_dates></a><details>
<summary>Flickr.Photo_dates</summary>

### Flickr.Photo_dates



![](Flickr/photo_datessimple.svg)

[Full Diagram](Flickr/photo_dates.svg)


#### Fields

| Field name | Type | Description |
|----|----|----|
| lastupdate | string | |
| posted | string | |
| taken | string | |
| takengranularity | string | |
| takenunknown | bool | |


</details>
<a name=Flickr.Photo_description></a><details>
<summary>Flickr.Photo_description</summary>

### Flickr.Photo_description



![](Flickr/photo_descriptionsimple.svg)

[Full Diagram](Flickr/photo_description.svg)


#### Fields

| Field name | Type | Description |
|----|----|----|
| _content | string | |


</details>
<a name=Flickr.Photo_editability></a><details>
<summary>Flickr.Photo_editability</summary>

### Flickr.Photo_editability



![](Flickr/photo_editabilitysimple.svg)

[Full Diagram](Flickr/photo_editability.svg)


#### Fields

| Field name | Type | Description |
|----|----|----|
| canaddmeta | bool | |
| cancomment | bool | |


</details>
<a name=Flickr.Photo_notes></a><details>
<summary>Flickr.Photo_notes</summary>

### Flickr.Photo_notes



![](Flickr/photo_notessimple.svg)

[Full Diagram](Flickr/photo_notes.svg)


#### Fields

| Field name | Type | Description |
|----|----|----|
| note | sequence of Note | |


</details>
<a name=Flickr.Photo_people></a><details>
<summary>Flickr.Photo_people</summary>

### Flickr.Photo_people



![](Flickr/photo_peoplesimple.svg)

[Full Diagram](Flickr/photo_people.svg)


#### Fields

| Field name | Type | Description |
|----|----|----|
| haspeople | bool | |


</details>
<a name=Flickr.Photo_permissions></a><details>
<summary>Flickr.Photo_permissions</summary>

### Flickr.Photo_permissions



![](Flickr/photo_permissionssimple.svg)

[Full Diagram](Flickr/photo_permissions.svg)


#### Fields

| Field name | Type | Description |
|----|----|----|
| permaddmeta | string | |
| permcomment | string | |


</details>
<a name=Flickr.Photo_publiceditability></a><details>
<summary>Flickr.Photo_publiceditability</summary>

### Flickr.Photo_publiceditability



![](Flickr/photo_publiceditabilitysimple.svg)

[Full Diagram](Flickr/photo_publiceditability.svg)


#### Fields

| Field name | Type | Description |
|----|----|----|
| canaddmeta | bool | |
| cancomment | bool | |


</details>
<a name=Flickr.Photo_tags></a><details>
<summary>Flickr.Photo_tags</summary>

### Flickr.Photo_tags



![](Flickr/photo_tagssimple.svg)

[Full Diagram](Flickr/photo_tags.svg)


#### Fields

| Field name | Type | Description |
|----|----|----|
| tag | sequence of Tag | |


</details>
<a name=Flickr.Photo_title></a><details>
<summary>Flickr.Photo_title</summary>

### Flickr.Photo_title



![](Flickr/photo_titlesimple.svg)

[Full Diagram](Flickr/photo_title.svg)


#### Fields

| Field name | Type | Description |
|----|----|----|
| _content | string | |


</details>
<a name=Flickr.Photo_urls></a><details>
<summary>Flickr.Photo_urls</summary>

### Flickr.Photo_urls



![](Flickr/photo_urlssimple.svg)

[Full Diagram](Flickr/photo_urls.svg)


#### Fields

| Field name | Type | Description |
|----|----|----|
| url | sequence of URL | |


</details>
<a name=Flickr.Photo_usage></a><details>
<summary>Flickr.Photo_usage</summary>

### Flickr.Photo_usage



![](Flickr/photo_usagesimple.svg)

[Full Diagram](Flickr/photo_usage.svg)


#### Fields

| Field name | Type | Description |
|----|----|----|
| canblog | bool | |
| candownload | bool | |
| canprint | bool | |
| canshare | bool | |


</details>
<a name=Flickr.Photo_visibility></a><details>
<summary>Flickr.Photo_visibility</summary>

### Flickr.Photo_visibility



![](Flickr/photo_visibilitysimple.svg)

[Full Diagram](Flickr/photo_visibility.svg)


#### Fields

| Field name | Type | Description |
|----|----|----|
| isfamily | bool | |
| isfriend | bool | |
| ispublic | bool | |


</details>
<a name=Flickr.Tag></a><details>
<summary>Flickr.Tag</summary>

### Flickr.Tag



![](Flickr/tagsimple.svg)

[Full Diagram](Flickr/tag.svg)


#### Fields

| Field name | Type | Description |
|----|----|----|
| _content | string | |
| author | string | |
| authorname | string | |
| id | string | |
| machine_tag | bool | |
| raw | string | |


</details>
<a name=Flickr.URL></a><details>
<summary>Flickr.URL</summary>

### Flickr.URL



![](Flickr/urlsimple.svg)

[Full Diagram](Flickr/url.svg)


#### Fields

| Field name | Type | Description |
|----|----|----|
| _content | string | |
| type | string | |


</details>


<div class="footer">

