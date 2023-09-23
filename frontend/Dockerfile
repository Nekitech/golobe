FROM node:18

WORKDIR /frontend
COPY package*.json ./
RUN npm i

COPY . .

CMD ["npm", "run", "dev"]
#ENTRYPOINT ["top", "-b"]