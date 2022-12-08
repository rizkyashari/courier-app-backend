-- Active: 1668865884394@@127.0.0.1@5432@courierdb

INSERT INTO
    add_ons (id, name, description, price)
VALUES (
        1,
        'Safe package',
        'Add more protection for your package',
        50000
    ), (
        2,
        'Cooler',
        'Keep your food or drink cool',
        20000
    ), (
        3,
        'Heatkeeper',
        'Keep your food or drink warm',
        15000
    );

INSERT INTO
    sizes (id, name, description, price)
VALUES (
        1,
        'Large',
        'VOLUME > 50x50x50 cm OR WEIGHT > 3 kg',
        150000
    ), (
        2,
        'Medium',
        '25x25x25 cm < VOLUME <= 50x50x50 cm OR 2 kg < WEIGHT <= 3 kg',
        100000
    ), (
        3,
        'Small',
        'VOLUME <= 25x25x525 cm OR WEIGHT <= 2 kg',
        75000
    );

INSERT INTO
    categories (id, name, description, price)
VALUES (
        1,
        'Food and Beverages',
        'Send your food or beverages safely with us',
        30000
    ), (
        2,
        'Fragile',
        'Send your fragile package safely with us',
        25000
    );

INSERT INTO
    promos (
        id,
        name,
        min_fee,
        max_discount,
        discount
    )
VALUES (
        1,
        '40% discount vouchers',
        20000,
        15000,
        40
    ), (
        2,
        '60% discount vouchers',
        20000,
        15000,
        60
    ), (
        3,
        '80% discount vouchers',
        20000,
        15000,
        80
    );

INSERT INTO
    source_of_funds (id, name)
VALUES (1, 'Bank'), (2, 'Cash');

INSERT INTO
    users (
        id,
        name,
        email,
        password,
        phone_number,
        role,
        referral_code,
        balance,
        photos,
        updated_at
    )
VALUES (
        1,
        'Admin',
        'admin@gmail.com',
        '$2a$04$yFEYXpkrZ5omhgRCru7amevoaVFWNVtc3gkEWjVLDY41fxfTXYqjy',
        6285722320621,
        'Admin',
        null,
        null,
        'https://res.cloudinary.com/dmlzx9yxe/image/upload/v1669964733/wazdcvhu1lzoqp6v6lqo.jpg',
        '2022-11-29 15:25:03.01474+07'
    ), (
        2,
        'Rizky Ashari',
        'rizkyashari@email.com',
        '$2a$04$yFEYXpkrZ5omhgRCru7amevoaVFWNVtc3gkEWjVLDY41fxfTXYqjy',
        6285722320622,
        'User',
        'RIZK3804',
        250000,
        'https://res.cloudinary.com/dmlzx9yxe/image/upload/v1669964733/wazdcvhu1lzoqp6v6lqo.jpg',
        '2022-11-29 15:25:03.01474+07'
    ), (
        3,
        'Dedy Irawan',
        'dedy@email.com',
        '$2a$04$yFEYXpkrZ5omhgRCru7amevoaVFWNVtc3gkEWjVLDY41fxfTXYqjy',
        6285722320623,
        'User',
        'DEDY8190',
        250000,
        'https://res.cloudinary.com/dmlzx9yxe/image/upload/v1669964733/wazdcvhu1lzoqp6v6lqo.jpg',
        '2022-11-29 15:25:03.01474+07'
    ), (
        4,
        'Ahmad Dhani',
        'ahmadhani@email.com',
        '$2a$04$yFEYXpkrZ5omhgRCru7amevoaVFWNVtc3gkEWjVLDY41fxfTXYqjy',
        6285722320624,
        'User',
        'AHMA8013',
        250000,
        'https://res.cloudinary.com/dmlzx9yxe/image/upload/v1669964733/wazdcvhu1lzoqp6v6lqo.jpg',
        '2022-11-29 15:25:03.01474+07'
    ), (
        5,
        'Sammy Manuaggal',
        'sammy@email.com',
        '$2a$04$yFEYXpkrZ5omhgRCru7amevoaVFWNVtc3gkEWjVLDY41fxfTXYqjy',
        6285722320625,
        'User',
        'SAMM7242',
        250000,
        'https://res.cloudinary.com/dmlzx9yxe/image/upload/v1669964733/wazdcvhu1lzoqp6v6lqo.jpg',
        '2022-11-29 15:25:03.01474+07'
    );

INSERT INTO
    addresses (
        user_id,
        full_address,
        recipient_name,
        recipient_phone_number,
        updated_at
    )
VALUES (
        2,
        'Gang Rawa Simprug VIII No. 43',
        'Euis Nilam',
        6285722320626,
        '2022-11-15 15:28:03.179661+07'
    ), (
        2,
        'Jl. Kebayoran VI No. 25',
        'Erwin Fernanda',
        6285722320627,
        '2022-11-16 15:28:03.179661+07'
    ), (
        3,
        'Jl. Rawa Buaya V No. 20',
        'Rizki Kurniawan',
        6285722320628,
        '2022-11-17 15:28:03.179661+07'
    ), (
        4,
        'Jl. Pegangsaan Timur No. 19',
        'Irvan Jayadi',
        6285722320629,
        '2022-11-18 15:28:03.179661+07'
    ), (
        5,
        'Jl. Kuningan Barat No. 10',
        'Fajrul Ahlam',
        6285722320630,
        '2022-11-19 15:28:03.179661+07'
    );

INSERT INTO
    payments (
        id,
        user_id,
        payment_status,
        total_cost,
        promo_id,
        updated_at
    )
VALUES (
        1,
        2,
        'Paid',
        230000,
        null,
        '2022-11-25 15:28:50.177986+07'
    ), (
        2,
        2,
        'Paid with Promo',
        175000,
        1,
        '2022-11-26 15:28:50.177986+07'
    ), (
        3,
        3,
        'Paid',
        140000,
        null,
        '2022-11-27 15:28:50.177986+07'
    ), (
        4,
        4,
        'Paid with Promo',
        105000,
        2,
        '2022-11-28 15:28:50.177986+07'
    ), (
        5,
        5,
        'Paid',
        150000,
        null,
        '2022-11-29 15:28:50.177986+07'
    );

INSERT INTO
    shippings (
        user_id,
        size_id,
        address_id,
        payment_id,
        category_id,
        add_on_id,
        shipping_status,
        review,
        updated_at
    )
VALUES (
        2,
        1,
        1,
        1,
        1,
        1,
        'Paid: delivery is being prepared',
        null,
        '2022-11-20 15:28:50.177986+07'
    ), (
        2,
        1,
        2,
        2,
        2,
        3,
        'Paid: delivery is being prepared',
        null,
        '2022-11-21 15:28:50.177986+07'
    ), (
        3,
        2,
        3,
        3,
        2,
        3,
        'Paid: delivery is being prepared',
        null,
        '2022-11-22 15:28:50.177986+07'
    ), (
        4,
        3,
        4,
        4,
        1,
        2,
        'Paid: delivery is being prepared',
        null,
        '2022-11-23 15:28:50.177986+07'
    ), (
        5,
        3,
        5,
        5,
        2,
        1,
        'Paid: delivery is being prepared',
        null,
        '2022-11-24 15:28:50.177986+07'
    );

INSERT INTO
    transactions (
        source_of_fund_id,
        user_id,
        destination_id,
        amount,
        description,
        category,
        updated_at
    )
VALUES (
        1,
        2,
        2,
        250000,
        'Top Up from Bank',
        'Top Up',
        '2022-11-20 15:28:50.177986+07'
    ), (
        1,
        2,
        2,
        250000,
        'Top Up from Bank',
        'Top Up',
        '2022-11-21 15:28:50.177986+07'
    ), (
        2,
        3,
        3,
        250000,
        'Top Up from Cash',
        'Top Up',
        '2022-11-22 15:28:50.177986+07'
    ), (
        null,
        4,
        4,
        105000,
        'Send package',
        'Send package',
        '2022-11-23 15:28:50.177986+07'
    ), (
        null,
        5,
        5,
        150000,
        'Send package',
        'Send package',
        '2022-11-24 15:28:50.177986+07'
    );

INSERT INTO
    user_promos (
        user_id,
        promo_id,
        status,
        updated_at
    )
VALUES (
        2,
        2,
        0,
        '2022-11-25 15:28:50.177986+07'
    ), (
        2,
        2,
        1,
        '2022-11-26 15:28:50.177986+07'
    ), (
        3,
        3,
        0,
        '2022-11-27 15:28:50.177986+07'
    ), (
        4,
        1,
        1,
        '2022-11-28 15:28:50.177986+07'
    ), (
        5,
        2,
        0,
        '2022-11-29 15:28:50.177986+07'
    );

INSERT INTO
    add_on_shippings (
        promotion,
        add_on_id,
        shipping_id
    )
VALUES ('40% discount vouchers', 1, 1), ('60% discount vouchers', 2, 2), ('80% discount vouchers', 3, 3), ('40% discount vouchers', 1, 2), ('60% discount vouchers', 2, 1);