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


CREATE TABLE IF NOT EXISTS users (
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

CREATE TABLE IF NOT EXISTS categories (
    id UUID PRIMARY KEY DEFAULT gen_random_uuid(),
    name VARCHAR(200) NOT NULL
);

CREATE TABLE IF NOT EXISTS courses (
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

CREATE TABLE IF NOT EXISTS lessons (
    id UUID PRIMARY KEY DEFAULT gen_random_uuid(),
    course_id UUID NOT NULL references courses(id),
    name VARCHAR(200) NOT NULL,
    description TEXT,
    video_url VARCHAR(325) NOT NULL,
    created_at TIMESTAMP DEFAULT CURRENT_TIMESTAMP,
    updated_at TIMESTAMP DEFAULT CURRENT_TIMESTAMP,
    deleted_at INT DEFAULT 0
);

CREATE TABLE IF NOT EXISTS additional_materials
(
    id UUID PRIMARY KEY DEFAULT gen_random_uuid(),
    lesson_id UUID references lessons(id),
    resource_url VARCHAR(325),
    title VARCHAR,
    description TEXT
);

CREATE TABLE IF NOT EXISTS user_courses 
(
    id UUID PRIMARY KEY DEFAULT gen_random_uuid(),
    user_id UUID NOT NULL references users(id),
    course_id UUID NOT NULL references courses(id),
    start_date TIMESTAMP DEFAULT CURRENT_TIMESTAMP,
    end_date TIMESTAMP DEFAULT NULL,
    status progress_status_type  
);
 
CREATE TABLE IF NOT EXISTS evaluations(
    id UUID PRIMARY KEY DEFAULT gen_random_uuid(),
    user_course_id UUID references user_courses(id),
    score INT DEFAULT 0 CHECK (score >= 0 AND score <= 5),
    feedback TEXT,
    created_at TIMESTAMP DEFAULT CURRENT_TIMESTAMP,     
    updated_at TIMESTAMP DEFAULT CURRENT_TIMESTAMP,
    deleted_at INT DEFAULT 0
);

CREATE TABLE IF NOT EXISTS user_lessons (
    id UUID PRIMARY KEY DEFAULT gen_random_uuid(),
    lesson_id UUID NOT NULL references lessons(id),
    user_course_id UUID NOT NULL references user_courses(id),
    status progress_status_type DEFAULT 'not_started'
);

CREATE TABLE IF NOT EXISTS certificates 
(
    id UUID PRIMARY KEY DEFAULT gen_random_uuid(), 
    user_course_id UUID references user_courses(id),
    created_at TIMESTAMP DEFAULT CURRENT_TIMESTAMP,
    status certificate_status_type,
    file_url VARCHAR(255)
);

CREATE TABLE IF NOT EXISTS orders(
    id UUID PRIMARY KEY DEFAULT gen_random_uuid(),
    user_id UUID references users(id),
    item_id UUID,
    type order_item_type,
    quantity INT DEFAULT 1,
    total_price DECIMAL(10,2),
    status order_status_type
);

CREATE TABLE IF NOT EXISTS transactions(
    id UUID PRIMARY KEY DEFAULT gen_random_uuid(),
    user_course_id UUID references user_courses(id) DEFAULT NULL,
    order_id UUID references orders(id) DEFAULT NULL,
    user_id UUID references users(id) DEFAULT NULL,
    amount DECIMAL(10,2),
    type transaction_type,
    created_at TIMESTAMP DEFAULT CURRENT_TIMESTAMP
);
create TABLE IF NOT EXISTS notifications (
    id UUID PRIMARY KEY DEFAULT gen_random_uuid(),
    reciever_id UUID NOT NULL,
    sender_id UUID,
    message TEXT NOT NULL,
    status  VARCHAR(50) DEFAULT 'pending',
    created_at TIMESTAMP DEFAULT NOW(),
    updated_at TIMESTAMP DEFAULT NOW(),
    deleted_at INT DEFAULT 0
)