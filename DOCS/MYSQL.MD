#### Mysql

Using this type, you can `restart` or `stop` or `start` Mysql db:

#### Reset

```yaml
- name: "Mysql Service"
  type: "mysql"
  command: "restart"
```
#### Stop

```yaml
- name: "Mysql Service"
  type: "mysql"
  command: "stop"
```
#### Start

```yaml
- name: "Mysql Service"
  type: "mysql"
  command: "start"
```
 
