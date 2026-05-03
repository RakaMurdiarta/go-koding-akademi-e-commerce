-- +goose Up
-- +goose StatementBegin

-- ==========================================================
-- 0. ENUM TYPES (PostgreSQL)
-- ==========================================================
DO $do$
BEGIN
    IF NOT EXISTS (SELECT 1 FROM pg_type WHERE typname = 'payment_status_enum') THEN
        CREATE TYPE payment_status_enum AS ENUM ('pending', 'paid', 'expired', 'cancelled');
    END IF;

    IF NOT EXISTS (SELECT 1 FROM pg_type WHERE typname = 'order_status_enum') THEN
        CREATE TYPE order_status_enum AS ENUM ('waiting_payment', 'processed', 'shipped', 'delivered', 'cancelled', 'completed');
    END IF;

    IF NOT EXISTS (SELECT 1 FROM pg_type WHERE typname = 'discount_type_enum') THEN
        CREATE TYPE discount_type_enum AS ENUM ('fixed', 'percentage');
    END IF;

    IF NOT EXISTS (SELECT 1 FROM pg_type WHERE typname = 'user_role_enum') THEN
        CREATE TYPE user_role_enum AS ENUM ('guest', 'admin');
    END IF;

    IF NOT EXISTS (SELECT 1 FROM pg_type WHERE typname = 'provider_enum') THEN
        CREATE TYPE provider_enum AS ENUM ('local', 'google', 'github');
    END IF;
END $do$;

-- ==========================================================
-- 1. MODUL: AUTH & USERS (Supports OAuth2)
-- ==========================================================
CREATE TABLE IF NOT EXISTS users (
    id SERIAL PRIMARY KEY,
    email VARCHAR(100) UNIQUE NOT NULL,
    password TEXT, -- NULL jika login via OAuth
    full_name VARCHAR(100),
    role user_role_enum DEFAULT 'guest',
    provider provider_enum DEFAULT 'local',
    provider_id VARCHAR(255),
    avatar_url TEXT,
    created_at TIMESTAMP DEFAULT CURRENT_TIMESTAMP,
    updated_at TIMESTAMP DEFAULT CURRENT_TIMESTAMP,
    deleted_at TIMESTAMP
);

-- Index umum untuk soft delete
CREATE INDEX IF NOT EXISTS idx_users_deleted_at ON users(deleted_at);

-- ==========================================================
-- 1.1 USER ADDRESSES (Snapshot + multi address)
-- ==========================================================
CREATE TABLE IF NOT EXISTS user_addresses (
    id SERIAL PRIMARY KEY,
    user_id INT REFERENCES users(id) ON DELETE CASCADE,

    receiver_name VARCHAR(100),
    phone VARCHAR(20),

    address TEXT,
    city VARCHAR(100),
    province VARCHAR(100),
    postal_code VARCHAR(10),

    is_default BOOLEAN DEFAULT false,
    created_at TIMESTAMP DEFAULT CURRENT_TIMESTAMP
);

CREATE INDEX IF NOT EXISTS idx_user_addresses_user_id ON user_addresses(user_id);

-- ==========================================================
-- 2. MODUL: SHOP (Marketplace Feature)
-- ==========================================================
CREATE TABLE IF NOT EXISTS shops (
    id SERIAL PRIMARY KEY,
    user_id INT UNIQUE REFERENCES users(id) ON DELETE CASCADE,
    name VARCHAR(100) UNIQUE NOT NULL,
    description TEXT,
    logo_url TEXT,
    is_active BOOLEAN DEFAULT true,
    created_at TIMESTAMP DEFAULT CURRENT_TIMESTAMP,
    updated_at TIMESTAMP DEFAULT CURRENT_TIMESTAMP,
    deleted_at TIMESTAMP
);

CREATE INDEX IF NOT EXISTS idx_shops_deleted_at ON shops(deleted_at);

-- ==========================================================
-- 3. MODUL: PRODUCTS & INVENTORY
-- ==========================================================
CREATE TABLE IF NOT EXISTS categories (
    id SERIAL PRIMARY KEY,
    name VARCHAR(50) NOT NULL,
    slug VARCHAR(50) UNIQUE NOT NULL,
    parent_id INT REFERENCES categories(id) ON DELETE SET NULL,
    created_at TIMESTAMP DEFAULT CURRENT_TIMESTAMP,
    updated_at TIMESTAMP DEFAULT CURRENT_TIMESTAMP,
    deleted_at TIMESTAMP
);

CREATE INDEX IF NOT EXISTS idx_categories_deleted_at ON categories(deleted_at);

CREATE TABLE IF NOT EXISTS products (
    id SERIAL PRIMARY KEY,
    shop_id INT REFERENCES shops(id) ON DELETE CASCADE,
    category_id INT REFERENCES categories(id),

    name VARCHAR(255) NOT NULL,
    slug VARCHAR(255) UNIQUE NOT NULL,
    description TEXT,

    price DECIMAL(15, 2) NOT NULL,
    weight INT DEFAULT 100,
    unit VARCHAR(255),
    thumbnail TEXT NOT NULL,

    -- tambahan production-ready
    is_active BOOLEAN DEFAULT true,
    rating_avg DECIMAL(2, 1) DEFAULT 0,
    rating_count INT DEFAULT 0,

    created_at TIMESTAMP DEFAULT CURRENT_TIMESTAMP,
    updated_at TIMESTAMP DEFAULT CURRENT_TIMESTAMP,
    deleted_at TIMESTAMP
);

CREATE INDEX IF NOT EXISTS idx_products_shop_id ON products(shop_id);
CREATE INDEX IF NOT EXISTS idx_products_category_id ON products(category_id);
CREATE INDEX IF NOT EXISTS idx_products_deleted_at ON products(deleted_at);

CREATE TABLE IF NOT EXISTS product_variants (
    id SERIAL PRIMARY KEY,
    product_id INT REFERENCES products(id) ON DELETE CASCADE,
    price_extra DECIMAL(15, 2) DEFAULT 0,
    created_at TIMESTAMP DEFAULT CURRENT_TIMESTAMP,
    updated_at TIMESTAMP DEFAULT CURRENT_TIMESTAMP,
    deleted_at TIMESTAMP
);

CREATE INDEX IF NOT EXISTS idx_product_variants_product_id ON product_variants(product_id);
CREATE INDEX IF NOT EXISTS idx_product_variants_deleted_at ON product_variants(deleted_at);

-- Product Images
CREATE TABLE IF NOT EXISTS product_images (
    id SERIAL PRIMARY KEY,
    product_id INT REFERENCES products(id) ON DELETE CASCADE,
    image_url TEXT NOT NULL,
    is_primary BOOLEAN DEFAULT false,
    created_at TIMESTAMP DEFAULT CURRENT_TIMESTAMP
);

CREATE INDEX IF NOT EXISTS idx_product_images_product_id ON product_images(product_id);

CREATE TABLE IF NOT EXISTS attributes (
    id SERIAL PRIMARY KEY,
    name VARCHAR(255) NOT NULL,
    CONSTRAINT uq_attribute_name UNIQUE ( name)
);

CREATE INDEX IF NOT EXISTS idx_attributes_name_id ON attributes(name);

CREATE TABLE IF NOT EXISTS attribute_values (
    id SERIAL PRIMARY KEY,
    attribute_id INT REFERENCES attributes(id) ON DELETE CASCADE,
    value VARCHAR(255) NOT NULL -- Example: '16GB', 'Blue'
);

-- EXPLICIT JOIN TABLE
CREATE TABLE IF NOT EXISTS variant_attribute_values (
    variant_id INT REFERENCES product_variants(id) ON DELETE CASCADE,
    attribute_value_id INT REFERENCES attribute_values(id) ON DELETE CASCADE,
    PRIMARY KEY (variant_id, attribute_value_id)
);

-- ==========================================================
-- 4. MODUL: CART & PROMOTION
-- ==========================================================
CREATE TABLE IF NOT EXISTS carts (
    id SERIAL PRIMARY KEY,
    user_id INT REFERENCES users(id) ON DELETE CASCADE,
    variant_id INT REFERENCES product_variants(id) ON DELETE CASCADE,
    quantity INT NOT NULL CHECK (quantity > 0),
    created_at TIMESTAMP DEFAULT CURRENT_TIMESTAMP
);

-- prevent duplicate cart line
CREATE UNIQUE INDEX IF NOT EXISTS idx_unique_user_variant
ON carts(user_id, variant_id);

CREATE INDEX IF NOT EXISTS idx_carts_user_id ON carts(user_id);
CREATE INDEX IF NOT EXISTS idx_carts_variant_id ON carts(variant_id);

CREATE TABLE IF NOT EXISTS promotions (
    id SERIAL PRIMARY KEY,
    code VARCHAR(20) UNIQUE NOT NULL,
    discount_type discount_type_enum,
    discount_value DECIMAL(15, 2) NOT NULL,
    min_spend DECIMAL(15, 2) DEFAULT 0,
    max_discount DECIMAL(15, 2),
    start_date TIMESTAMP,
    end_date TIMESTAMP,
    is_active BOOLEAN DEFAULT true
);

CREATE INDEX IF NOT EXISTS idx_promotions_code ON promotions(code);
CREATE INDEX IF NOT EXISTS idx_promotions_is_active ON promotions(is_active);

-- ==========================================================
-- 5. MODUL: TRANSACTIONS & ORDER SPLITTING
-- ==========================================================
CREATE TABLE IF NOT EXISTS transactions (
    id SERIAL PRIMARY KEY,
    user_id INT REFERENCES users(id),
    promotion_id INT REFERENCES promotions(id),

    grand_total DECIMAL(15, 2) NOT NULL,
    payment_method VARCHAR(50),
    payment_status payment_status_enum DEFAULT 'pending',
    snap_token TEXT,

    -- tambahan penting
    payment_expired_at TIMESTAMP,

    created_at TIMESTAMP DEFAULT CURRENT_TIMESTAMP
);

CREATE INDEX IF NOT EXISTS idx_transactions_user_id ON transactions(user_id);
CREATE INDEX IF NOT EXISTS idx_transactions_promotion_id ON transactions(promotion_id);
CREATE INDEX IF NOT EXISTS idx_transactions_payment_status ON transactions(payment_status);

CREATE TABLE IF NOT EXISTS orders (
    id SERIAL PRIMARY KEY,
    transaction_id INT REFERENCES transactions(id) ON DELETE CASCADE,
    shop_id INT REFERENCES shops(id),

    total_price DECIMAL(15, 2) NOT NULL,
    shipping_cost DECIMAL(15, 2) DEFAULT 0,
    status order_status_enum DEFAULT 'waiting_payment',
    tracking_number VARCHAR(100),

    -- snapshot shipping
    receiver_name VARCHAR(100),
    receiver_phone VARCHAR(20),
    shipping_address TEXT,

    created_at TIMESTAMP DEFAULT CURRENT_TIMESTAMP
);

CREATE INDEX IF NOT EXISTS idx_orders_transaction_id ON orders(transaction_id);
CREATE INDEX IF NOT EXISTS idx_orders_shop_id ON orders(shop_id);
CREATE INDEX IF NOT EXISTS idx_orders_status ON orders(status);

CREATE TABLE IF NOT EXISTS order_items (
    id SERIAL PRIMARY KEY,
    order_id INT REFERENCES orders(id) ON DELETE CASCADE,
    variant_id INT REFERENCES product_variants(id),
    qty INT NOT NULL CHECK (qty > 0),
    price_at_buy DECIMAL(15, 2) NOT NULL
);

CREATE INDEX IF NOT EXISTS idx_order_items_order_id ON order_items(order_id);
CREATE INDEX IF NOT EXISTS idx_order_items_variant_id ON order_items(variant_id);

CREATE TABLE IF NOT EXISTS inventory (
    id SERIAL PRIMARY KEY,
    variant_id INT UNIQUE REFERENCES product_variants(id) ON DELETE CASCADE,

    total_stock INT DEFAULT 0 CHECK (total_stock >= 0),
    reserved_stock INT DEFAULT 0 CHECK (reserved_stock >= 0),
    CONSTRAINT chk_stock_integrity CHECK (total_stock >= reserved_stock),

    -- optimistic locking (optional but recommended)
    version INT DEFAULT 1,

    created_at TIMESTAMP DEFAULT CURRENT_TIMESTAMP,
    updated_at TIMESTAMP DEFAULT CURRENT_TIMESTAMP
);

CREATE INDEX IF NOT EXISTS idx_inventory_variant_id ON inventory(variant_id);


-- ==========================================================
-- 6. MODUL: REVIEWS
-- ==========================================================
CREATE TABLE IF NOT EXISTS reviews (
    id SERIAL PRIMARY KEY,
    order_item_id INT UNIQUE REFERENCES order_items(id),
    user_id INT REFERENCES users(id),
    product_id INT REFERENCES products(id),

    rating INT CHECK (rating >= 1 AND rating <= 5),
    comment TEXT,
    created_at TIMESTAMP DEFAULT CURRENT_TIMESTAMP
);

CREATE INDEX IF NOT EXISTS idx_reviews_user_id ON reviews(user_id);
CREATE INDEX IF NOT EXISTS idx_reviews_product_id ON reviews(product_id);

-- +goose StatementEnd


-- +goose Down
-- +goose StatementBegin

-- Drop tables (reverse dependency order)
DROP TABLE IF EXISTS reviews CASCADE;

DROP TABLE IF EXISTS inventory CASCADE;

DROP TABLE IF EXISTS order_items CASCADE;
DROP TABLE IF EXISTS orders CASCADE;
DROP TABLE IF EXISTS transactions CASCADE;

DROP TABLE IF EXISTS promotions CASCADE;
DROP TABLE IF EXISTS carts CASCADE;

DROP TABLE IF EXISTS attributes CASCADE;
DROP TABLE IF EXISTS attribute_values CASCADE;
DROP TABLE IF EXISTS variant_attribute_values CASCADE;

DROP TABLE IF EXISTS product_images CASCADE;
DROP TABLE IF EXISTS product_variants CASCADE;
DROP TABLE IF EXISTS products CASCADE;
DROP TABLE IF EXISTS categories CASCADE;

DROP TABLE IF EXISTS shops CASCADE;
DROP TABLE IF EXISTS user_addresses CASCADE;
DROP TABLE IF EXISTS users CASCADE;

-- Drop types
DO $do$
BEGIN
    IF EXISTS (SELECT 1 FROM pg_type WHERE typname = 'order_status_enum') THEN
        DROP TYPE order_status_enum;
    END IF;

    IF EXISTS (SELECT 1 FROM pg_type WHERE typname = 'payment_status_enum') THEN
        DROP TYPE payment_status_enum;
    END IF;

    IF EXISTS (SELECT 1 FROM pg_type WHERE typname = 'discount_type_enum') THEN
        DROP TYPE discount_type_enum;
    END IF;

    IF EXISTS (SELECT 1 FROM pg_type WHERE typname = 'provider_enum') THEN
        DROP TYPE provider_enum;
    END IF;

    IF EXISTS (SELECT 1 FROM pg_type WHERE typname = 'user_role_enum') THEN
        DROP TYPE user_role_enum;
    END IF;
END $do$;

-- +goose StatementEnd