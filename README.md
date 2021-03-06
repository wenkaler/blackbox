# BLACK BOX

#### Описание
Программа предназначена для хранения настроек любого проекта в JSON формате. Основная функция для каждого проекта - создавать отдельный экземпляр настроек, который будет доступен по token ключу и, соответственно, для каждой копии, можно изменить настройки и применить отдельно, не затронув остальные копии проекта.

## Project

### 1. Create project.
**ROUTE**  /api/v1/project/create

**Method** POST

**Пример запроса:**
```json
{
	"various": "proxy",
	"setting" : {"ServiceName": "proxy"}
}
```
**Пример ответа:**
```json
    {
        "id": 1,
        "name": "proxy",
        "token": "H8FADPXw7W13gFmte4JYVitvbxwLAfMRnbb129x8fzCMFxGs2Z",
        "basic_setting": {
            "ServiceName": "proxy"
        } 
    }   
```
* name - уникальное имя проекта 
* token - уникальный токен ключ проекта по которому будут осуществляться все запросы. 
* basic_setting - основные настройки проекта, c них будет браться копия для экзепляров проекта.

### 2. Get project.
**ROUTER** /api/v1/project/get

**Method** GET

**Пример запроса:**
* **Header:**
    * **token**: *H8FADPXw7W13gFmte4JYVitvbxwLAfMRnbb129x8fzCMFxGs2Z*

**Пример ответа:**
```json
    {
        "id": 1,
        "name": "proxy",
        "token": "H8FADPXw7W13gFmte4JYVitvbxwLAfMRnbb129x8fzCMFxGs2Z",
        "basic_setting": {
            "ServiceName": "proxy"
        } 
    }   
```


### 3. Update project.
**ROUTER** /api/v1/project/update

**Method** POST

**Пример запроса:**
```json
    {
        "name": "proxy1",
        "token": "H8FADPXw7W13gFmte4JYVitvbxwLAfMRnbb129x8fzCMFxGs2Z",
        "basic_setting": {
            "ServiceName": "proxy1"
        } 
    }  
```
**Пример ответа:**
```json
    {
        "id": 1,
        "name": "proxy1",
        "token": "H8FADPXw7W13gFmte4JYVitvbxwLAfMRnbb129x8fzCMFxGs2Z",
        "basic_setting": {
            "ServiceName": "proxy1"
        } 
    }   
```

> Можно изменить основные настройки для экзепляров проекта и имя самого проекта.

### 4. List project.
**ROUTER** /api/v1/project/list

**Method** GET

**Пример запроса:**
* **Header:**
    * **token**: *H8FADPXw7W13gFmte4JYVitvbxwLAfMRnbb129x8fzCMFxGs2Z*

**Пример ответа:**
```json

```

### 5. Remove project.
**ROUTER** /api/v1/project/remove

**Method** GET

**Пример запроса:**
* **Header:**
    * **token**: *H8FADPXw7W13gFmte4JYVitvbxwLAfMRnbb129x8fzCMFxGs2Z*




## Setting

### 1. Init setting
**ROUTER** /api/v1/setting/init

**Method** GET

**Пример запроса:**
* **Header:**
    * **token**: *H8FADPXw7W13gFmte4JYVitvbxwLAfMRnbb129x8fzCMFxGs2Z*
    * **various**: *10.10.20.20*
**Пример ответа**
```json
    {
        "id": 1,
        "id_project": 1,
        "indexer": "10.10.20.20",
        "token": "LGW3ypQGacomjj2XlxLKUZB4GReTlwZ7E69DvaSXDHJw1LUmJp",
        "status": "new",
        "setting": {
            "ServiceName": "proxy"
        },
        "last_update_date": "2019-04-29T10:42:44.621471"
    }
```

* indexer - индетификатор экземпляра настроек. 
* token - уникальный ключ для отдельного экземпляра настроек.
* setting - настройки который должен применять данный экзепляр проекта. - изначально берется из основных настроек проэкьа.
* last_update_date - последний раз обращения к настройкам.
* status - определяет в каком сейчас состоянии находятся настройки. существует 3 состояния. (new, pending, done)

### 2. Clean setting
**ROUTER** /api/v1/setting/clean

**Method** GET

**Пример запроса:**
* **Header:**
    * **token**: *H8FADPXw7W13gFmte4JYVitvbxwLAfMRnbb129x8fzCMFxGs2Z* 
    >необходимо использовать токен проекта.
    * **various**: *24 hour*

**Пример ответа**
```json

```

### 3. Get setting
**ROUTER** /api/v1/setting/get

**Method** GET

**Пример запроса:**
* **Header:**
    * **token**: *LGW3ypQGacomjj2XlxLKUZB4GReTlwZ7E69DvaSXDHJw1LUmJp*

**Пример ответа**
```json
    {
        "id": 1,
        "id_project": 1,
        "indexer": "10.10.20.20",
        "token": "LGW3ypQGacomjj2XlxLKUZB4GReTlwZ7E69DvaSXDHJw1LUmJp",
        "status": "pending",
        "setting": {
            "ServiceName": "proxy"
        },
        "last_update_date": "2019-04-29T10:42:44.621471"
    }
```

### 4. Confirm setting
**ROUTER** /api/v1/setting/confirm

**Method** GET

**Пример запроса:**
* **Header:**
    * **token**: *LGW3ypQGacomjj2XlxLKUZB4GReTlwZ7E69DvaSXDHJw1LUmJp*

**Пример ответа**
```json
    {
        "id": 1,
        "id_project": 1,
        "indexer": "10.10.20.20",
        "token": "LGW3ypQGacomjj2XlxLKUZB4GReTlwZ7E69DvaSXDHJw1LUmJp",
        "status": "done",
        "setting": {
            "ServiceName": "proxy"
        },
        "last_update_date": "2019-04-29T10:48:44.878789"
    }
```

### 5. Update setting
**ROUTER** /api/v1/setting/update

**Method** POST

**Пример запроса:**
```json
{
	"various": "LGW3ypQGacomjj2XlxLKUZB4GReTlwZ7E69DvaSXDHJw1LUmJp",
	"setting": {
		"ServiceName": "proxy1"
	}
}
```
**Пример ответа**
```json
    {
        "id": 1,
        "id_project": 1,
        "indexer": "10.10.20.20",
        "token": "LGW3ypQGacomjj2XlxLKUZB4GReTlwZ7E69DvaSXDHJw1LUmJp",
        "status": "done",
        "setting": {
            "ServiceName": "proxy1"
        },
        "last_update_date": "2019-04-29T10:48:44.878789"
    }
```