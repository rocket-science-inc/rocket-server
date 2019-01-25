FROM node:11

# setup environment
ENV VERSION             "0.1.0"
ENV SERVICE_PATH        rocket-server/serverless

LABEL serviceless.version=${VERSION}

# create workspace and get project sources
WORKDIR ${SERVICE_PATH}
ADD . .

# install project dependecies
RUN yarn
# run
CMD yarn start

# network
EXPOSE 4000