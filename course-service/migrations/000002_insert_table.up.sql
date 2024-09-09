
-- Mock Data for Users
INSERT INTO users (first_name, last_name, dob, phone_number, email, occupation, address, password, gender, photo_url, role, is_confirmed) VALUES
('John', 'Doe', '1990-01-15', '+15551234567', 'john.doe@example.com', 'Software Engineer', '123 Main St', 'hashed_password', 'Male', 'https://example.com/john.jpg', 'user', TRUE),
('Jane', 'Smith', '1985-05-20', '+15559876543', 'jane.smith@example.com', 'Data Scientist', '456 Oak Ave', 'hashed_password', 'Female', 'https://example.com/jane.jpg', 'user', TRUE),
('Alice', 'Johnson', '1992-11-10', '+15551357924', 'alice.johnson@example.com', 'Product Manager', '789 Pine Ln', 'hashed_password', 'Female', 'https://example.com/alice.jpg', 'user', FALSE);

-- Mock Data for Categories
INSERT INTO categories (name) VALUES
('Web Development'),
('Data Science'),
('Business');

-- Mock Data for Courses
INSERT INTO courses (name, description, price, image_url, category_id) VALUES
('Introduction to Python', 'Learn the basics of Python programming.', 49.99, 'https://example.com/python.jpg', (SELECT id FROM categories WHERE name = 'Web Development')),
('Machine Learning Fundamentals', 'Explore the fundamentals of machine learning.', 79.99, 'https://example.com/machine_learning.jpg', (SELECT id FROM categories WHERE name = 'Data Science')),
('Project Management Essentials', 'Master the essentials of project management.', 59.99, 'https://example.com/project_management.jpg', (SELECT id FROM categories WHERE name = 'Business'));

-- Mock Data for Lessons
INSERT INTO lessons (course_id, name, description, video_url) VALUES
((SELECT id FROM courses WHERE name = 'Introduction to Python'), 'Variables and Data Types', 'Introduction to variables and data types in Python.', 'https://example.com/python_variables.mp4'),
((SELECT id FROM courses WHERE name = 'Introduction to Python'), 'Control Flow', 'Learn about control flow statements in Python.', 'https://example.com/python_control_flow.mp4'),
((SELECT id FROM courses WHERE name = 'Machine Learning Fundamentals'), 'Supervised Learning', 'Introduction to supervised learning algorithms.', 'https://example.com/supervised_learning.mp4'),
((SELECT id FROM courses WHERE name = 'Machine Learning Fundamentals'), 'Unsupervised Learning', 'Explore unsupervised learning techniques.', 'https://example.com/unsupervised_learning.mp4'),
((SELECT id FROM courses WHERE name = 'Project Management Essentials'), 'Project Initiation', 'Learn how to initiate a project effectively.', 'https://example.com/project_initiation.mp4'),
((SELECT id FROM courses WHERE name = 'Project Management Essentials'), 'Project Planning', 'Master the art of project planning.', 'https://example.com/project_planning.mp4');

-- Mock Data for Additional Materials
INSERT INTO additional_materials (lesson_id, resource_url, title, description) VALUES
((SELECT id FROM lessons WHERE name = 'Variables and Data Types'), 'https://example.com/python_variables_cheatsheet.pdf', 'Python Variables Cheatsheet', 'A handy cheatsheet for Python variables and data types.'),
((SELECT id FROM lessons WHERE name = 'Control Flow'), 'https://example.com/python_control_flow_examples.zip', 'Python Control Flow Examples', 'Example code for different control flow statements in Python.');

-- Mock Data for User Courses
INSERT INTO user_courses (user_id, course_id, status) VALUES
((SELECT id FROM users WHERE email = 'john.doe@example.com'), (SELECT id FROM courses WHERE name = 'Introduction to Python'), 'in_progress'),
((SELECT id FROM users WHERE email = 'jane.smith@example.com'), (SELECT id FROM courses WHERE name = 'Machine Learning Fundamentals'), 'completed'),
((SELECT id FROM users WHERE email = 'alice.johnson@example.com'), (SELECT id FROM courses WHERE name = 'Project Management Essentials'), 'not_started');

-- Mock Data for Evaluations
INSERT INTO evaluations (user_course_id, score, feedback) VALUES
((SELECT id FROM user_courses WHERE user_id = (SELECT id FROM users WHERE email = 'jane.smith@example.com') AND course_id = (SELECT id FROM courses WHERE name = 'Machine Learning Fundamentals')), 5, 'Great course! I learned a lot about machine learning.'),
((SELECT id FROM user_courses WHERE user_id = (SELECT id FROM users WHERE email = 'john.doe@example.com') AND course_id = (SELECT id FROM courses WHERE name = 'Introduction to Python')), 4, 'Good introduction to Python programming.');

-- Mock Data for User Lessons
INSERT INTO user_lessons (lesson_id, user_course_id, status) VALUES
((SELECT id FROM lessons WHERE name = 'Variables and Data Types'), (SELECT id FROM user_courses WHERE user_id = (SELECT id FROM users WHERE email = 'john.doe@example.com') AND course_id = (SELECT id FROM courses WHERE name = 'Introduction to Python')), 'completed'),
((SELECT id FROM lessons WHERE name = 'Control Flow'), (SELECT id FROM user_courses WHERE user_id = (SELECT id FROM users WHERE email = 'john.doe@example.com') AND course_id = (SELECT id FROM courses WHERE name = 'Introduction to Python')), 'in_progress'),
((SELECT id FROM lessons WHERE name = 'Supervised Learning'), (SELECT id FROM user_courses WHERE user_id = (SELECT id FROM users WHERE email = 'jane.smith@example.com') AND course_id = (SELECT id FROM courses WHERE name = 'Machine Learning Fundamentals')), 'completed'),
((SELECT id FROM lessons WHERE name = 'Unsupervised Learning'), (SELECT id FROM user_courses WHERE user_id = (SELECT id FROM users WHERE email = 'jane.smith@example.com') AND course_id = (SELECT id FROM courses WHERE name = 'Machine Learning Fundamentals')), 'completed');

-- Mock Data for Certificates
INSERT INTO certificates (user_course_id, status, file_url) VALUES
((SELECT id FROM user_courses WHERE user_id = (SELECT id FROM users WHERE email = 'jane.smith@example.com') AND course_id = (SELECT id FROM courses WHERE name = 'Machine Learning Fundamentals')), 'confirmed', 'https://example.com/certificates/jane_smith_ml.pdf');

-- Mock Data for Orders
INSERT INTO orders (user_id, item_id, type, quantity, total_price, status) VALUES
((SELECT id FROM users WHERE email = 'john.doe@example.com'), gen_random_uuid(), 'book', 1, 19.99, 'delivered'),
((SELECT id FROM users WHERE email = 'jane.smith@example.com'), gen_random_uuid(), 'product', 1, 99.00, 'pending');

