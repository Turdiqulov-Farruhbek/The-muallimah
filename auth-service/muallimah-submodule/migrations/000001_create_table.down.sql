DROP TABLE IF EXISTS orders;

DROP TABLE IF EXISTS certificates;

DROP TABLE IF EXISTS transactions;

DROP TABLE IF EXISTS user_lessons;

DROP TABLE IF EXISTS evaluations;

DROP TABLE IF EXISTS user_courses;

DROP TABLE IF EXISTS additional_materials;

DROP TABLE IF EXISTS lessons;

DROP TABLE IF EXISTS courses;

DROP TABLE IF EXISTS categories;

DROP TABLE IF EXISTS users;

DO $$
BEGIN
    IF EXISTS (SELECT 1 FROM pg_type WHERE typname = 'order_status_type') THEN
        DROP TYPE order_status_type;
    END IF;

    IF EXISTS (SELECT 1 FROM pg_type WHERE typname = 'order_item_type') THEN
        DROP TYPE order_item_type;
    END IF;

    IF EXISTS (SELECT 1 FROM pg_type WHERE typname = 'certificate_status_type') THEN
        DROP TYPE certificate_status_type;
    END IF;

    IF EXISTS (SELECT 1 FROM pg_type WHERE typname = 'transaction_type') THEN
        DROP TYPE transaction_type;
    END IF;

    IF EXISTS (SELECT 1 FROM pg_type WHERE typname = 'progress_status_type') THEN
        DROP TYPE progress_status_type;
    END IF;
END $$;
