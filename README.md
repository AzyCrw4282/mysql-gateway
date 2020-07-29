This project allows interaction with the MySQL database over HTTP


## What is it?

To interact wit the `mysql` database it requires you to either access it via the 
work-bench or on the `cmd-line/terminal` through `mysql` cli. This for simple operation
is cumbersome and so this was born to solve that, e.g.

To `SELECT * from users`, you simply do
```
curl http://localhost:8080/getTable=users
```


## Intended Features

- get the table, 
- get a row by id, 
- get required coloumns
- update where field=<value>,
- delete where field=<value>,
- insert operation

-- For all above, be able to limit results

- get filtered values from multiple columns

-- For all above, perform aggregate function

- Perform operations on relational table 


## Usage | Instructions

TBA











#### Acknowledgements

This project was inspired by the works from user @just1689 and his project, `pg-gatway` where postgresql db can be interacted 
over HTTP using module such as `curl` .



#### Packages used