CREATE TABLE orders (
    id varchar(255) NOT NULL,
    price float NOT NULL,
    tax float NOT NULL,
    final_price float NOT NULL,
    PRIMARY KEY (id)
);


CREATE TABLE orders_test (
                        id varchar(255) NOT NULL,
                        price float NOT NULL,
                        tax float NOT NULL,
                        final_price float NOT NULL,
                        PRIMARY KEY (id)
);

--{"id":"abc","price":10.0,"tax":1.0}
