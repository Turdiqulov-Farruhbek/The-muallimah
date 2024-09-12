DO $$
BEGIN
    IF NOT EXISTS (SELECT 1 FROM pg_type WHERE typname = 'progress_status_type') THEN
        CREATE TYPE progress_status_type AS ENUM ('not_started', 'in_progress', 'completed');
    END IF;

    IF NOT EXISTS (SELECT 1 FROM pg_type WHERE typname = 'transaction_type') THEN
        CREATE TYPE transaction_type AS ENUM ('debit', 'credit');
    END IF;

    IF NOT EXISTS (SELECT 1 FROM pg_type WHERE typname = 'certificate_status_type') THEN
        CREATE TYPE certificate_status_type AS ENUM ('pending', 'confirmed');
    END IF;

    IF NOT EXISTS (SELECT 1 FROM pg_type WHERE typname = 'order_item_type') THEN
        CREATE TYPE order_item_type AS ENUM ('book', 'product');
    END IF;

    IF NOT EXISTS (SELECT 1 FROM pg_type WHERE typname = 'order_status_type') THEN
        CREATE TYPE order_status_type AS ENUM ('pending', 'delivered');
    END IF;
END $$;

-- Create the necessary tables in the correct order

CREATE TABLE users (
    id UUID PRIMARY KEY DEFAULT gen_random_uuid(),
    first_name VARCHAR(200) NOT NULL,
    last_name VARCHAR(200) NOT NULL,
    dob DATE,
    phone_number VARCHAR(200) NOT NULL,
    email VARCHAR(200) UNIQUE NOT NULL,
    occupation VARCHAR(200),
    address VARCHAR(200),
    password VARCHAR(350),
    gender VARCHAR(200),
    photo_url VARCHAR(350),
    role VARCHAR(200) DEFAULT 'user',
    is_confirmed BOOLEAN DEFAULT FALSE,
    created_at TIMESTAMP DEFAULT CURRENT_TIMESTAMP,
    updated_at TIMESTAMP DEFAULT CURRENT_TIMESTAMP,
    deleted_at INT DEFAULT 0
);

CREATE TABLE categories (
    id UUID PRIMARY KEY DEFAULT gen_random_uuid(),
    name VARCHAR(200) NOT NULL
);

CREATE TABLE courses (
    id UUID PRIMARY KEY DEFAULT gen_random_uuid(),
    name VARCHAR(200) NOT NULL,
    description TEXT,
    price FLOAT DEFAULT 0.0,
    image_url VARCHAR(225) NOT NULL DEFAULT '',
    category_id UUID NOT NULL references categories(id), 
    created_at TIMESTAMP DEFAULT CURRENT_TIMESTAMP,
    updated_at TIMESTAMP DEFAULT CURRENT_TIMESTAMP,
    deleted_at INT DEFAULT 0
);

CREATE TABLE lessons (
    id UUID PRIMARY KEY DEFAULT gen_random_uuid(),
    course_id UUID NOT NULL references courses(id),
    name VARCHAR(200) NOT NULL,
    description TEXT,
    video_url VARCHAR(325) NOT NULL,
    created_at TIMESTAMP DEFAULT CURRENT_TIMESTAMP,
    updated_at TIMESTAMP DEFAULT CURRENT_TIMESTAMP,
    deleted_at INT DEFAULT 0
);

CREATE TABLE additional_materials
(
    id UUID PRIMARY KEY DEFAULT gen_random_uuid(),
    lesson_id UUID references lessons(id),
    resource_url VARCHAR(325),
    title VARCHAR,
    description TEXT
);

CREATE TABLE user_courses 
(
    id UUID PRIMARY KEY DEFAULT gen_random_uuid(),
    user_id UUID NOT NULL references users(id),
    course_id UUID NOT NULL references courses(id),
    start_date TIMESTAMP DEFAULT CURRENT_TIMESTAMP,
    end_date TIMESTAMP DEFAULT NULL,
    status progress_status_type  
);
 
CREATE TABLE evaluations(
    id UUID PRIMARY KEY DEFAULT gen_random_uuid(),
    user_course_id UUID references user_courses(id),
    score INT DEFAULT 0 CHECK (score >= 0 AND score <= 5),
    feedback TEXT,
    created_at TIMESTAMP DEFAULT CURRENT_TIMESTAMP,     
    updated_at TIMESTAMP DEFAULT CURRENT_TIMESTAMP,
    deleted_at INT DEFAULT 0
);

CREATE TABLE user_lessons (
    id UUID PRIMARY KEY DEFAULT gen_random_uuid(),
    lesson_id UUID NOT NULL references lessons(id),
    user_course_id UUID NOT NULL references user_courses(id),
    status progress_status_type DEFAULT 'not_started'
);

-- Create the orders table before it's referenced in transactions
CREATE TABLE orders(
    id UUID PRIMARY KEY DEFAULT gen_random_uuid(),
    user_id UUID references users(id),
    item_id UUID,
    type order_item_type,
    quantity INT DEFAULT 1,
    total_price DECIMAL(10,2),
    status order_status_type
);

CREATE TABLE transactions(
    id UUID PRIMARY KEY DEFAULT gen_random_uuid(),
    user_course_id UUID references user_courses(id),
    order_id UUID references orders(id),
    user_id UUID references users(id),
    amount DECIMAL(10,2),
    type transaction_type,
    created_at TIMESTAMP DEFAULT CURRENT_TIMESTAMP
);

CREATE TABLE certificates 
(
    id UUID PRIMARY KEY DEFAULT gen_random_uuid(),
    user_course_id UUID references user_courses(id),
    created_at TIMESTAMP DEFAULT CURRENT_TIMESTAMP,
    status certificate_status_type,
    file_url VARCHAR(255)
);
