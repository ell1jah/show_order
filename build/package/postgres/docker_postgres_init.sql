CREATE TABLE IF NOT EXISTS delivery (
    id      SERIAL PRIMARY KEY,
    name    VARCHAR(255),
    phone   VARCHAR(20),
    zip     VARCHAR(10),
    city    VARCHAR(255),
    address VARCHAR(255),
    region  VARCHAR(255),
    email   VARCHAR(255)
);

CREATE TABLE IF NOT EXISTS payment (
    id      SERIAL PRIMARY KEY,
    transaction  VARCHAR(255),
    request_id   INTEGER,
    currency     VARCHAR(10),
    provider     VARCHAR(255),
    amount       INTEGER,
    payment_dt   INTEGER,
    bank         VARCHAR(255),
    delivery_cost INTEGER,
    goods_total   INTEGER,
    custom_fee    INTEGER
);

CREATE TABLE IF NOT EXISTS order (
    order_uid          VARCHAR(255),
    track_number       VARCHAR(255),
    entry              VARCHAR(255),
    locale             VARCHAR(255),
    internal_signature VARCHAR(255),
    customer_id        VARCHAR(255),
    delivery_service   VARCHAR(255),
    shardkey           INTEGER,
    sm_id              INTEGER,
    date_created       TIMESTAMP,
    oof_shard          INTEGER,
    delivery_id        INTEGER REFERENCES delivery(id),
    payment_id         INTEGER REFERENCES payment(id)
);

CREATE TABLE IF NOT EXISTS item (
    id           SERIAL PRIMARY KEY,
    order_uid    VARCHAR(255), REFERENCES order(order_uid),
    chrt_id      INTEGER,
    track_number VARCHAR(255),
    price        INTEGER,
    rid          VARCHAR(255),
    name         VARCHAR(255),
    sale         INTEGER,
    size         INTEGER,
    total_price  INTEGER,
    nm_id        INTEGER,
    brand        VARCHAR(255),
    status       INTEGER
);
