#### Nginx

Using this type, you can `restart` or `stop` or `start` or `reload` Nginx HTTP Server:

#### Reset

```yaml
- name: "Nginx Service"
  type: "nginx"
  command: "restart"
```
#### Stop

```yaml
- name: "Nginx Service"
  type: "nginx"
  command: "stop"
```
#### Start

```yaml
- name: "Nginx Service"
  type: "nginx"
  command: "start"
```

#### Reload

```yaml
- name: "Nginx Service"
  type: "nginx"
  command: "reload"
```
 
