# go-library-API-pr3
This project is based a video from Akhil Sharma. (go-bookstore-API)

This API performs basic CRUD(Create, Read, Update, Delete) operations on a MYSQL DB.

The API has been developed based on a video from Akhil Sharma. 

In this repositary i have kept the source code of the API & dockerfile script to deploy this API in a docker container.

The .env file basically contains mysql password and library name

## To Do LIST

- [x] Code Changes in Dockerfile to copy all the source file to the docker.
- [x] Code Changes in Dockerfile to give the MySQL config.
- [x] Code Changes in Dockerfile to configure the database in the MySQL and set up login credentials
- [x] Create a Docker Compose file(YAML) to implement the connection between golang and MySQL (port details, networking, MySQL passwords)
- [ ] Change the underlying user to non-root with minimum privilaged access.
- [ ] Change the MySQL user to non-root user.
- [ ] Create a Front-End Application for this project. 
