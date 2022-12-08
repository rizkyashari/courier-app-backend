CREATE TABLE
    "users" (
        "user_id" int PRIMARY KEY,
        "email" varchar,
        "password" varchar,
        "long_name" varchar,
        "phone_number" int,
        "role" varchar,
        "referral_code" varchar,
        "created_at" timestamp,
        "updated_at" timestamp,
        "deleted_at" timestamp
    );

CREATE TABLE
    "addresses" (
        "address_id" int PRIMARY KEY,
        "user_id" int,
        "full_address" varchar,
        "recipient_name" varchar,
        "recipient_phone_number" int,
        "created_at" timestamp,
        "updated_at" timestamp,
        "deleted_at" timestamp
    );

CREATE TABLE
    "shippings" (
        "shipping_id" int PRIMARY KEY,
        "user_id" int,
        "size_id" int,
        "address_id" int,
        "payment_id" int,
        "category_id" int,
        "shipping_status" varchar,
        "review" varchar,
        "created_at" timestamp,
        "updated_at" timestamp,
        "deleted_at" timestamp
    );

CREATE TABLE
    "categories" (
        "category_id" int PRIMARY KEY,
        "name" varchar,
        "description" varchar,
        "price" number,
        "created_at" timestamp,
        "updated_at" timestamp,
        "deleted_at" timestamp
    );

CREATE TABLE
    "sizes" (
        "size_id" int PRIMARY KEY,
        "name" varchar,
        "description" varchar,
        "price" number,
        "created_at" timestamp,
        "updated_at" timestamp,
        "deleted_at" timestamp
    );

CREATE TABLE
    "payments" (
        "payment_id" int PRIMARY KEY,
        "payment_status" varchar,
        "total_cost" varchar,
        "promo_id" int,
        "created_at" timestamp,
        "updated_at" timestamp,
        "deleted_at" timestamp
    );

CREATE TABLE
    "promos" (
        "promo_id" int PRIMARY KEY,
        "name" varchar,
        "min_fee" int,
        "max_discount" int,
        "discount" int,
        "quota" int,
        "exp_date" timestamp,
        "created_at" timestamp,
        "updated_at" timestamp,
        "deleted_at" timestamp
    );

CREATE TABLE
    "add_on_shippings" (
        "promotion" varchar,
        "shipping_id" int,
        "add_on_id" int
    );

CREATE TABLE
    "add_ons" (
        "add_on_id" int PRIMARY KEY,
        "name" varchar,
        "description" varchar,
        "price" int,
        "created_at" timestamp,
        "updated_at" timestamp,
        "deleted_at" timestamp
    );

ALTER TABLE "addresses"
ADD
    FOREIGN KEY ("user_id") REFERENCES "users" ("user_id");

ALTER TABLE "shippings"
ADD
    FOREIGN KEY ("user_id") REFERENCES "users" ("user_id");

ALTER TABLE "shippings"
ADD
    FOREIGN KEY ("address_id") REFERENCES "addresses" ("address_id");

ALTER TABLE "shippings"
ADD
    FOREIGN KEY ("category_id") REFERENCES "categories" ("category_id");

ALTER TABLE "shippings"
ADD
    FOREIGN KEY ("size_id") REFERENCES "sizes" ("size_id");

ALTER TABLE "shippings"
ADD
    FOREIGN KEY ("payment_id") REFERENCES "payments" ("payment_id");

ALTER TABLE "payments"
ADD
    FOREIGN KEY ("promo_id") REFERENCES "promos" ("promo_id");

ALTER TABLE "add_on_shippings"
ADD
    FOREIGN KEY ("shipping_id") REFERENCES "shippings" ("shipping_id");

ALTER TABLE "add_on_shippings"
ADD
    FOREIGN KEY ("add_on_id") REFERENCES "add_ons" ("add_on_id");