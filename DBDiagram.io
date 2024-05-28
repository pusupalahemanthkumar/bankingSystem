Table "users" {
    "username" varchar [pk]
    "hashed_password" varchar [not null]
    "full_name" varchar [not null]
    "email" varchar [unique, not null]
    "password_changed_at" timestamptz [not null, default: `0001-01-01 00:00:00Z` ]
    "created_at" timestamptz [not null, default: `now()`]

    Indexes {
        "username"
        "email"
    }
}

Table  "accounts" {
    "id" bigserial [pk, increment]
    "owner" varchar [not null]
    "balance" DECIMAL(20, 5)   [not null]
    "currency" varchar [not null]
    "created_at" timestamptz [not null, default: `now()`]
    "updated_at" timestamptz [not null, default: `now()`]

    Indexes {
        "owner"
        ("owner", "currency") [unique]
    }
}

Table "entries" {
    "id" bigserial [pk, increment]
    "account_id" bigint [not null]
    "amount" DECIMAL(20, 5) [not null]
    "created_at" timestamptz [not null, default: `now()`]

    Indexes {
        "account_id"
    }
}

Table "transfers"{
    "id" bigserial [pk, increment]
    "from_account_id" bigint [not null]
    "to_account_id" bigint [not null]
    "amount" DECIMAL(20, 5)  [not null]
    "created_at" timestamptz [not null, default: `now()`]

    Indexes {
        "from_account_id" 
        "to_account_id"
        ("from_account_id", "to_account_id")
    }

}

Ref:"accounts"."id" < "entries"."account_id"

Ref:"accounts"."id" < "transfers"."from_account_id"

Ref:"accounts"."id" < "transfers"."to_account_id"

Ref:"users"."username" < "accounts"."owner"