FROM influx6/mongrel-0.0.1
MAINTAINER GOKIT(gitbub.com/gokit) <trinoxf@gmail.com>

# This is a test image, dont do this for any production
# system please, am begging please. Instead load secrets
# through the --env-file flag for docker run .
ENV USERCLAIMJWT_MONGO_TEST_HOST 0.0.0.0:27017
ENV USERCLAIMJWT_MONGO_TEST_DB test_db
ENV USERCLAIMJWT_MONGO_TEST_AUTHDB test_db

CMD [/bin/sh -c "/bin/bootmgo --no-auth --fork && /usr/bin/mongo --file /mnt/db/mongodb/db.js && go test -v ./..."]