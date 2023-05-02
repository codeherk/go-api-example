-- Create Tasks table
USE sample;

CREATE TABLE Tasks (
    id INT NOT NULL AUTO_INCREMENT,
    description VARCHAR(255),
    created_timestamp TIMESTAMP NOT NULL DEFAULT CURRENT_TIMESTAMP,
    updated_timestamp TIMESTAMP NOT NULL DEFAULT CURRENT_TIMESTAMP ON UPDATE CURRENT_TIMESTAMP,
    PRIMARY KEY (id)
);

-- Populate 20 tasks
INSERT INTO Tasks (description) VALUES 
('Write a poem about the changing seasons'),
('Research the history of the Eiffel Tower'),
('Create a new recipe for a vegetarian lasagna'),
('Design a logo for a new business'),
('Plan a weekend road trip through the mountains'),
('Read a book in a new genre'),
('Learn a new skill on Skillshare'),
('Organize your closet using the KonMari method'),
('Write a short story about a time traveler'),
('Paint a picture of your favorite place'),
('Try a new workout routine for a week'),
('Take a photography course online'),
('Research and book a volunteer trip abroad'),
('Make homemade soap from scratch'),
('Write a letter to your future self'),
('Build a piece of furniture from reclaimed wood'),
('Create a vision board for your goals'),
('Organize a fundraiser for a local charity'),
('Learn to play a new instrument'),
('Spend a day without technology');