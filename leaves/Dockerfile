FROM node:22 as deps
WORKDIR /app

# COPY package.json package-lock.json ./
# RUN npm install --production=false

# # Build the app.
# FROM deps as build

# COPY tsconfig.json tailwind.config.cjs postcss.config.cjs app.config.ts Makefile ./
# COPY src src
# COPY public public

# RUN npm run build

# # A production dependencies image.
# # FROM deps as production-deps

# # ENV NODE_ENV=production
# # RUN npm prune --omit=dev --force

# # Right now, there's no production dependencies.
# FROM node:22-slim
# WORKDIR /app

# COPY --from=build /app/.output /app

CMD ["node", "server/index.mjs"]
