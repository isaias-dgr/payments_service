echo "🛠️ Load Var Env"
export PROJ_NAME="payments"
export PROJ_ENV="dev"

export POSTGRES_DATABASE="${PROJ_ENV}_${PROJ_NAME}"
export POSTGRES_HOST="${PROJ_NAME}-db"
export POSTGRES_PORT="5432"
export POSTGRES_USER="user_${PROJ_ENV}_${PROJ_NAME}"
export POSTGRES_PASSWORD="ab22cd66-56d9-4b65-80d2-f675c0afba49"
export POSTGRES_CONN="postgresql://${POSTGRES_USER}:${POSTGRES_PASSWORD}@${POSTGRES_HOST}:${POSTGRES_PORT}/${POSTGRES_DATABASE}"

echo "🏁 End Load Var Env"