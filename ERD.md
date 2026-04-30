erDiagram
    users {
        int id PK
        varchar email UK
        text password
        varchar full_name
        enum role
        enum provider
        varchar provider_id
        text avatar_url
        timestamp created_at
        timestamp updated_at
        timestamp deleted_at
    }

    user_addresses {
        int id PK
        int user_id FK
        varchar receiver_name
        varchar phone
        text address
        varchar city
        varchar province
        varchar postal_code
        boolean is_default
        timestamp created_at
    }

    shops {
        int id PK
        int user_id FK "UK"
        varchar name UK
        text description
        text logo_url
        boolean is_active
        timestamp created_at
        timestamp updated_at
        timestamp deleted_at
    }

    categories {
        int id PK
        varchar name
        varchar slug UK
        timestamp deleted_at
    }

    products {
        int id PK
        int shop_id FK
        int category_id FK
        varchar name
        varchar slug UK
        text description
        decimal price
        int weight
        boolean is_active
        decimal rating_avg
        int rating_count
        timestamp created_at
        timestamp updated_at
        timestamp deleted_at
    }

    product_variants {
        int id PK
        int product_id FK
        varchar sku UK
        decimal price_extra
        timestamp created_at
        timestamp updated_at
        timestamp deleted_at
    }

    product_images {
        int id PK
        int product_id FK
        text image_url
        boolean is_primary
        timestamp created_at
    }

    attributes {
        int id PK
        varchar name
        int category_id FK
    }

    variant_attribute_values {
        int id PK
        int variant_id FK
        int attribute_id FK
        varchar value
    }

    carts {
        int id PK
        int user_id FK
        int variant_id FK
        int quantity
        timestamp created_at
    }

    promotions {
        int id PK
        varchar code UK
        enum discount_type
        decimal discount_value
        decimal min_spend
        decimal max_discount
        timestamp start_date
        timestamp end_date
        boolean is_active
    }

    transactions {
        int id PK
        int user_id FK
        int promotion_id FK
        decimal grand_total
        varchar payment_method
        enum payment_status
        text snap_token
        timestamp payment_expired_at
        timestamp created_at
    }

    orders {
        int id PK
        int transaction_id FK
        int shop_id FK
        decimal total_price
        decimal shipping_cost
        enum status
        varchar tracking_number
        varchar receiver_name
        varchar receiver_phone
        text shipping_address
        timestamp created_at
    }

    order_items {
        int id PK
        int order_id FK
        int variant_id FK
        int qty
        decimal price_at_buy
    }

    inventory {
        int id PK
        int variant_id FK "UK"
        int total_stock
        int reserved_stock
        int version
        timestamp created_at
        timestamp updated_at
    }

    payment_logs {
        int id PK
        int transaction_id FK
        jsonb payload
        timestamp created_at
    }

    reviews {
        int id PK
        int order_item_id FK "UK"
        int user_id FK
        int product_id FK
        int rating
        text comment
        timestamp created_at
    }

    %% Relationships %%
    users ||--o{ user_addresses : "has"
    users ||--o| shops : "owns (1:1)"
    users ||--o{ carts : "has"
    users ||--o{ transactions : "makes"
    users ||--o{ reviews : "writes"

    shops ||--o{ products : "sells"
    shops ||--o{ orders : "receives"

    categories ||--o{ products : "contains"
    categories ||--o{ attributes : "has"

    products ||--o{ product_variants : "has"
    products ||--o{ product_images : "has"
    products ||--o{ reviews : "gets"

    product_variants ||--o{ variant_attribute_values : "has"
    product_variants ||--o{ carts : "added to"
    product_variants ||--o| inventory : "tracked by (1:1)"
    product_variants ||--o{ order_items : "included in"

    attributes ||--o{ variant_attribute_values : "defines"

    promotions |o--o{ transactions : "applied to"

    transactions ||--o{ orders : "splits into"
    transactions ||--o{ payment_logs : "has"

    orders ||--o{ order_items : "contains"

    order_items ||--o| reviews : "receives (1:1)"