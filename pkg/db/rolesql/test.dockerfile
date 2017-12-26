FROM influx6/squarel-0.0.1
MAINTAINER GOKIT(gitbub.com/gokit) <trinoxf@gmail.com>

# This is a test image, dont do this for any production
# system please, am begging please. Instead load secrets 
# through the --env-file flag for docker run .
ENV MYSQL_ROOT_PASSWORD 123456
ENV MYSQL_DATABASE test_db
ENV MYSQL_USER test_user
ENV MYSQL_USER_PASSWORD test_user_pass
ENV PKG_SQL_TEST_DB test_db
ENV PKG_SQL_TEST_USER test_user
ENV PKG_SQL_TEST_ADDR 0.0.0.0
ENV PKG_SQL_TEST_PORT 3306
ENV PKG_SQL_TEST_Driver mysq
ENV PKG_SQL_TEST_PASSWORD test_user_pass

CMD /bin/sh -c $fork=true bootsql && go test -v ./...