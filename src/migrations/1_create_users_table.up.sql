CREATE TABLE "users" (
    "id" INTEGER PRIMARY KEY AUTOINCREMENT NOT NULL,
    "created" INTEGER NOT NULL,
    "last_login" INTEGER NULL,
    "name" TEXT NOT NULL,
    "password" TEXT NOT NULL,
    "is_admin" INTEGER NOT NULL,
    "is_active" INTEGER NOT NULL
)
