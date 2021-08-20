#### Folder

Using this type, you can `create` or `delete` or `force_delete` folders:

#### Create

```yaml
- name: "create the folder"
  type: "folder"
  command: "create"
  path: "/var/www/html/test"
```

#### Delete

```yaml
- name: "delete the folder"
  type: "folder"
  command: "delete"
  path: "/var/www/html/test"
```

#### Force Delete

This will delete the folder and its contents.

```yaml
- name: "delete the folder and it's contents"
  type: "folder"
  command: "force_delete"
  path: "/var/www/html/test"
```

