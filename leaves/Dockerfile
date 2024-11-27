FROM node:22 AS deps
WORKDIR /app

COPY package.json package-lock.json ./
RUN npm install --production=false

FROM deps AS build

COPY tsconfig.json tailwind.config.cjs postcss.config.cjs app.config.ts ./
COPY src src
COPY public public

RUN npm run build

FROM node:22-slim
WORKDIR /app

COPY --from=build /app/.output /app

CMD ["node", "server/index.mjs"]
