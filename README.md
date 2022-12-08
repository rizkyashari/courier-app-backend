# final-project-backend

This is a backend repository for final project of Courier App using Golang.

## Problem Description

### Admin Story:

- Auth & Profile Management:

  - [✓] Admin can login and logout.
  - [✓] On this page, users can edit email addresses, long names, phone numbers, and add profile photos.
  - [✓] Profile photos can be saved in the database as BLOBs.

- View Addresses List

  - [✓] View a list of addresses.

- Manage Shippings:

  - [1/2] View a list of shippings. Can be sorted by size, category, payment and status.
  - [✓] If the user clicks on a shipping, the user can see detailed information about that shipping. Users can also see a review of the shipping if any.
  - [1/2] View earnings reports by month.
  - [✓] Update shipping status.

- Manage Promos:

  - [✓] View the list of promos.
  - [✓] Can be sorted by quota and expiration date.
  - [✓] Doing an update on a promo.

- Additionals:
  - [1/2] Add search and pagination for possible pages.

### User Story:

- Auth & Profile Management:

  - [✓] User can login and logout.
  - [✓] Profile page. On this page, users can edit email addresses, long names, phone numbers, and add profile photos. Profile photos can be saved in the database as BLOBs. On this page users can also see a referral code that can be shared with other potential users.
  - [✓] The registration page will have input for: Email, Passwords, Full names, Phone number, and Referral codes.

- Referral Systems:

  - [1/2] A new user who registers using a referral code will get an additional balance of Rp. 50,000 after completing the cumulative transaction amounting to Rp. 350,000.
  - [0] Existing users whose referral code is used are entitled to an additional balance of Rp. 25,000 when a new user has completed a cumulative transaction of Rp. 500,000.

- Create Shipping:

  - [✓] User can choose a size.
  - [✓] User can choose a category.
  - [1/2] Choose an add-on (can be more than one, but 1 type of add-on can only be selected once).
  - [✓] Choose a shipping address.

- Shippings List:

  - [✓] User can view a list of shipping that has been created by the user, including the status of each shipping.
  - [✓] When the user clicks on a shipping, the user can view the shipping details and make a payment.
  - [✓] When the user clicks on a shipping that already has the status done, the user can leave a comment in the form of text. These comments are used as service satisfaction feedback.

- Create Address:

  - [✓] This page is used to enter a new shipping address. This address can later be used on the create shipping page.

- Address List:

  - [✓] This page will display a list of addresses that have been created by the user.

- Edit addresses:

  - [✓] This page is used to edit the address that has been saved by the user.

- Payments:

  - [✓] On this page, users can make payments for a shipping. Users can also choose which promotions to use (if available).

- Top Up:

  - [✓] On this page, users can make payments for a shipping. Users can also choose which promotions to use (if available).

- Games:

  - [1/2] For every shipment that is done, the user has one chance to play the gacha game. The prizes that users might get are 40% discount vouchers (min. spending Rp. 20,000, max. Rp. 20,000), 60% discount vouchers (min. spending Rp. 20,000, max. Rp. 20,000), 80% discount vouchers (min. spending IDR 20,000, max IDR 20,000).

- Additional:
  - [✓] Add search and pagination for possible pages.

## Technical Requirement

- [✓] Make a REST API server using Go. Add authentication and authorization to secure access to these endpoints.
- [✓] Determine what endpoints must be available to fulfill the features mentioned above. Also think about the URL parameters, URL query, header, request body, response status and response body for each endpoint created.
- [✓] Implement a middleware error handler and the right message/status code for each successful and failed request.
- [✓] All delete features will use the concept of soft delete.
- [✓] API Documentation (Swagger) for each endpoint that has been created, each endpoint must have the following details:

  - URL and URL Parameters (if any)
  - Example payload
  - Example of a successful response and a failed response

- [✓] To save the image, you can upload the image to the image sharing platform and save the image URL to the database. (example: Cloudinary)
- [0] Deploy backend to hosting service

## Prerequisites to run the application

- DB

  - Setup db using gorm table auto-migration for model & seed.sql for the data sample **Note:the data seeding must be executed sequentially**
  - ENV already setup using local configuration, feel free to update the .env because datebase-seeding.sql didn't define db name. After creating the table, seed.sql can be executed
  - Default admin & user that already defined
    - admin@gmail.com / password
    - user@gmail.com / password

- Development

  - Golang
  - PostgreSQL
  - Gin :
    > `go get github.com/gin-gonic/gin`
  - Gorm :
    > `go get gorm.io/gorm`
  - Postgres Driver :
    > `go get gorm.io/driver/postgres`
  - Godotenv :
    > `go get github.com/joho/godotenv`
  - JWT :
    > `go get github.com/golang-jwt/jwt/v4`

- Testing
  - Testify:
    > `go get github.com/stretchr/testify`
  - Mockery:
    > `go install github.com/vektra/mockery/v2@latest`
  - GoSQL Mock:
    > `go get github.com/DATA-DOG/go-sqlmock`

## ERD

## ![ERD Diagram](https://res.cloudinary.com/dmlzx9yxe/image/upload/v1670238593/fnccsudshvbm2c4lsoki.png "erd")

## Testing Coverage

![Testing Coverage](https://res.cloudinary.com/dmlzx9yxe/image/upload/v1670258604/fjglkpwoh9blkkg3piel.png "coverage")

## API Documentation

More about API Implementation, explained in swagger docs:
[API Documentation](http://localhost:8080/dist/)

---

## Run Application & Test Suite

- Application

  - Start Application :
    > `go run .`

- Test Suite

  - Run All Test :

    > `make test`
