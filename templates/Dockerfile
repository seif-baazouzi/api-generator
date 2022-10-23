FROM node

RUN useradd -ms /bin/bash user
USER user

WORKDIR /home/user

COPY package.json .
RUN npm install

COPY --chown=user:user src src

CMD npm start
