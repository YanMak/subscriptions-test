CREATE TABLE subscription (
    id serial PRIMARY KEY,
    service_name text NOT NULL,
    price integer NOT NULL,
    user_id uuid NOT NULL,
    start_date date NOT NULL,
    end_date date
);

CREATE INDEX idx_subscription_user_id ON subscription(user_id);